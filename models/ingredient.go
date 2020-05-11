package models

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/dgraph-io/dgo/v2"
	"github.com/dgraph-io/dgo/v2/protos/api"
)

// IngredientResponse is a struct that represents a single ingredient. It is used exclusively
// for marshalling responses back to API clients.
type IngredientResponse struct {
	ID     string `json:"uid,omitempty"`
	Name   string `json:"name,omitempty" validate:"required"`
	Amount string `json:"amount",omitempty`

	Categories []NestedCategory `json:"categories,omitempty"`

	DType []string `json:"dgraph.type,omitempty"`
}

// NestedIngredient is a stripped down struct used when an Ingredient is nested
// within a parent struct in an API response
type NestedIngredient struct {
	ID     string `json:"uid,omitempty"`
	Name   string `json:"name,omitempty" validate:"required"`
	Amount string `json:"amount",omitempty`

	DType []string `json:"dgraph.type,omitempty"`
}

// Ingredient is a struct that represents a single ingredient
type Ingredient struct {
	ID     string `json:"uid,omitempty"`
	Name   string `json:"name,omitempty" validate:"required"`
	Amount string `json:"ingredients|amount",omitempty`

	Categories []Category `json:"categories,omitempty"`

	DType []string `json:"dgraph.type,omitempty"`
}

// ManyIngredients is a struct that represents multiple ingredients
type ManyIngredients struct {
	Ingredients []Ingredient `json:"ingredients"`
}

// parent struct for dgraph responses
type rootIngredient struct {
	Ingredients []Ingredient `json:"root"`
}

// GetAllIngredients will fetch all ingredients
func GetAllIngredients(c *dgo.Dgraph) (*ManyIngredients, error) {
	txn := c.NewReadOnlyTxn()

	const q = `
		{
			root(func: type(Ingredient)) {
				uid
				name
				categories {
					uid
					name
				}
				dgraph.type
			}
		}
	`
	resp, err := txn.Query(context.Background(), q)
	if err != nil {
		return nil, err
	}

	root := rootIngredient{}
	err = json.Unmarshal(resp.Json, &root)
	if err != nil {
		return nil, err
	}

	return &ManyIngredients{root.Ingredients}, nil
}

// GetIngredient will fetch an ingredient via a given ID
func (ingredient *Ingredient) GetIngredient(c *dgo.Dgraph) error {
	txn := c.NewReadOnlyTxn()

	variables := map[string]string{"$id": ingredient.ID}
	const q = `
		query all($id: string) {
			root(func: uid($id)) @filter(type(Ingredient)) {
				uid
				name
				categories {
					uid
					name
				}
				dgraph.type
			}
		}
	`
	resp, err := txn.QueryWithVars(context.Background(), q, variables)
	if err != nil {
		return err
	}

	root := rootIngredient{}
	err = json.Unmarshal(resp.Json, &root)
	if err != nil {
		return err
	}

	*ingredient = root.Ingredients[0]

	return nil
}

// CreateIngredient will create a new ingredient from the given Ingredient struct
func (ingredient *Ingredient) CreateIngredient(c *dgo.Dgraph) error {
	fmt.Println("CreateIngredient() start")

	txn := c.NewTxn()
	defer txn.Discard(context.Background())

	// assign an alias ID that can be ref'd out of the response's uid []string map
	ingredient.ID = "_:ingredient"
	ingredient.DType = []string{"Ingredient"}

	pb, err := json.Marshal(ingredient)
	if err != nil {
		return err
	}

	mu := &api.Mutation{
		CommitNow: true,
	}
	mu.SetJson = pb
	res, err := txn.Mutate(context.Background(), mu)
	if err != nil {
		return err
	}

	fmt.Println("CreateIngredient mutation resp: ")
	fmt.Printf("%+v\n", res)

	ingredient.ID = res.Uids["ingredient"]

	return nil
}

// UpdateIngredient will update an ingredient via a given ID
func (ingredient *Ingredient) UpdateIngredient(c *dgo.Dgraph) error {
	fmt.Println("UpdateIngredient() start")

	txn := c.NewTxn()
	defer txn.Discard(context.Background())

	ingredient.DType = []string{"Ingredient"}

	pb, err := json.Marshal(ingredient)
	if err != nil {
		return err
	}

	mu := &api.Mutation{
		CommitNow: true,
	}
	mu.SetJson = pb
	res, err := txn.Mutate(context.Background(), mu)
	if err != nil {
		return err
	}

	fmt.Println("CreateIngredient mutation resp: ")
	fmt.Printf("%+v\n", res)

	ingredient.ID = res.Uids["ingredient"]

	return nil
}

// DeleteIngredient will delete an ingredient via a given by ID
func (ingredient *Ingredient) DeleteIngredient(c *dgo.Dgraph) error {
	return nil
}

package models

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/dgraph-io/dgo/v2"
	"github.com/dgraph-io/dgo/v2/protos/api"
)

// Ingredient is a struct that represents a single ingredient
type Ingredient struct {
	ID         string     `json:"uid,omitempty"`
	Name       string     `json:"name,omitempty" validate:"required"`
	Categories []Category `json:"categories,omitempty"`
	DType      []string   `json:"dgraph.type,omitempty"`

	Amount string `json:"ingredients|amount",omitempty`
}

// ManyIngredients is a struct that represents multiple ingredients
type ManyIngredients struct {
	Ingredients []Ingredient
}

// parent struct for dgraph responses
type rootIngredient struct {
	Ingredient []Ingredient `json:"root"`
}

// GetIngredient will fetch an ingredient via a given ID
func (i *Ingredient) GetIngredient(c *dgo.Dgraph) error {
	txn := c.NewReadOnlyTxn()

	variables := map[string]string{"$id": i.ID}
	const q = `
		query all($id: string) {
			root(func: uid($id)) {
				uid
				name
				ingredient_categories
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

	// this works fine for just a single field, but should use some reflection to
	// copy all fields from the temp struct to the calling one
	i.Name = root.Ingredient[0].Name

	return nil
}

// UpdateIngredient will update the name of an ingredient via a given by ID
func (i *Ingredient) UpdateIngredient(c *dgo.Dgraph) error {
	fmt.Println("UpdateIngredient() start")

	txn := c.NewTxn()
	defer txn.Discard(context.Background())

	i.DType = []string{"Ingredient"}

	pb, err := json.Marshal(i)
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

	i.ID = res.Uids["ingredient"]

	return nil
}

// DeleteIngredient will delete an ingredient via a given by ID
func (i *Ingredient) DeleteIngredient(c *dgo.Dgraph) error {
	return errors.New("Not implemented")
}

// CreateIngredient will create a new ingredient from the given Ingredient struct
func (i *Ingredient) CreateIngredient(c *dgo.Dgraph) error {
	fmt.Println("CreateIngredient() start")

	txn := c.NewTxn()
	defer txn.Discard(context.Background())

	// assign an alias ID that can be ref'd out of the response's uid []string map
	i.ID = "_:ingredient"
	i.DType = []string{"Ingredient"}

	pb, err := json.Marshal(i)
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

	i.ID = res.Uids["ingredient"]

	return nil
}

// GetAllIngredients will fetch all ingredients
func GetAllIngredients(c *dgo.Dgraph) (*ManyIngredients, error) {
	txn := c.NewReadOnlyTxn()

	const q = `
		{
			q(func: type(Ingredient)) {
				uid
				name
				ingredient_categories
			}
		}
	`
	resp, err := txn.Query(context.Background(), q)
	if err != nil {
		return nil, err
	}

	i := ManyIngredients{}
	err = json.Unmarshal(resp.Json, &i)
	if err != nil {
		return nil, err
	}

	return &i, nil
}

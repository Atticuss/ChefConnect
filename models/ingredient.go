package models

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/dgraph-io/dgo/v2"
	"github.com/dgraph-io/dgo/v2/protos/api"
)

// Ingredient is a struct that represents a single ingredient
type Ingredient struct {
	ID    string   `json:"uid"`
	Name  string   `json:"ingredient_name" validate:"required"`
	DType []string `json:"dgraph.type,omitempty"`
}

// ManyIngredients is a struct that represents multiple ingredients
type ManyIngredients struct {
	Ingredients []Ingredient
}

// cheap hack to get around how dgraph returns data
type singleIngredient struct {
	Ingredient []Ingredient
}

// GetIngredient will fetch an ingredient via a given ID
func (i *Ingredient) GetIngredient(c *dgo.Dgraph) error {
	txn := c.NewReadOnlyTxn()

	variables := map[string]string{"$id": i.ID}
	const q = `
		query all($id: string) {
			ingredient(func: uid($id)) {
				uid
				ingredient_name
			}
		}
	`
	resp, err := txn.QueryWithVars(context.Background(), q, variables)
	if err != nil {
		return err
	}

	single := singleIngredient{}
	err = json.Unmarshal(resp.Json, &single)
	if err != nil {
		return err
	}

	i.Name = single.Ingredient[0].Name

	return nil
}

// UpdateIngredient will update the name of an ingredient via a given by ID
func (i *Ingredient) UpdateIngredient(c *dgo.Dgraph) error {
	return errors.New("Not implemented")
}

// DeleteIngredient will delete an ingredient via a given by ID
func (i *Ingredient) DeleteIngredient(c *dgo.Dgraph) error {
	return errors.New("Not implemented")
}

// CreateIngredient will create a new ingredient from the given Ingredient struct
func (i *Ingredient) CreateIngredient(c *dgo.Dgraph) error {
	txn := c.NewTxn()
	defer txn.Discard(context.Background())

	i.ID = "_:ingredient"

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

	i.ID = res.Uids["ingredient"]

	return nil
}

// GetAllIngredients will fetch all ingredients
func GetAllIngredients(c *dgo.Dgraph) (*ManyIngredients, error) {
	txn := c.NewReadOnlyTxn()

	const q = `
		{
			ingredients(func: has(ingredient_name)) {
				uid
				ingredient_name
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

package models

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/dgraph-io/dgo/v2"
	"github.com/dgraph-io/dgo/v2/protos/api"
)

// Category is a struct that represents a single category
type Category struct {
	ID          string       `json:"uid,omitempty"`
	Name        string       `json:"name,omitempty" validate:"required"`
	Recipes     []Recipe     `json:"recipe_categories,omitempty"`
	Ingredients []Ingredient `json:"ingredient_categories,omitempty"`
	DType       []string     `json:"dgraph.type,omitempty"`
}

// ManyCategories is a struct that represents multiple categories
type ManyCategories struct {
	ManyCategories []Category
}

// parent struct for dgraph responses
type rootCategory struct {
	Category []Category `json:root`
}

// CreateCategory will create a new ingredient from the given Ingredient struct
func (cat *Category) CreateCategory(c *dgo.Dgraph) error {
	fmt.Println("CreateCategory() start")

	txn := c.NewTxn()
	defer txn.Discard(context.Background())

	// assign an alias ID that can be ref'd out of the response's uid []string map
	cat.ID = "_:category"
	cat.DType = []string{"Category"}

	pb, err := json.Marshal(cat)
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

	fmt.Println("CreateCategory mutation resp: ")
	fmt.Printf("%+v\n", res)

	cat.ID = res.Uids["category"]

	return nil
}

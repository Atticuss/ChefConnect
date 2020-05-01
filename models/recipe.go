package models

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/dgraph-io/dgo/v2"
	"github.com/dgraph-io/dgo/v2/protos/api"
)

// Recipe is a struct that represents a single recipe
type Recipe struct {
	ID            string       `json:"uid,omitempty"`
	Name          string       `json:"name,omitempty" validate:"required"`
	URL           string       `json:"url,omitempty"`
	Domain        string       `json:"domain,omitempty"`
	Directions    string       `json:"directions,omitempty"`
	Ingredients   []Ingredient `json:"ingredients,omitempty"`
	PrepTime      int          `json:"prep_time,omitempty"`
	CookTime      int          `json:"cook_time,omitempty"`
	TotalServings int          `json:"total_servings,omitempty"`
	Categories    []Category   `json:"categories,omitempty"`
	HasBeenTried  bool         `json:"has_been_tried,omitempty"`

	RatedBy        []User   `json:"~ratings,omitempty"`
	RatingScore    int      `json:"ratings|score,omitempty"`
	FavoritedBy    []User   `json:"~favorites,omitempty"`
	RelatedRecipes []Recipe `json:"related_recipes,omitempty"`
	Notes          []Note   `json:"~recipe,omitempty"`
	DType          []string `json:"dgraph.type,omitempty"`
}

// ManyRecipes is a struct that represents multiple recipes
type ManyRecipes struct {
	Recipes []Recipe
}

// cheap hack to get around how dgraph returns data
type singleRecipe struct {
	Recipe []Recipe
}

// GetRecipe will get a recipe via a given by ID
func (r *Recipe) GetRecipe(c *dgo.Dgraph) error {
	return errors.New("Not implemented")
}

// UpdateRecipe will update a recipe via a given by ID
func (r *Recipe) UpdateRecipe(c *dgo.Dgraph) error {
	return errors.New("Not implemented")
}

// DeleteRecipe will delete a recipe via a given by ID
func (r *Recipe) DeleteRecipe(c *dgo.Dgraph) error {
	return errors.New("Not implemented")
}

// CreateRecipe will create a new recipe from the given Recipe struct
func (r *Recipe) CreateRecipe(c *dgo.Dgraph) error {
	fmt.Println("CreateRecipe() start")

	txn := c.NewTxn()
	defer txn.Discard(context.Background())

	// assign an alias ID that can be ref'd out of the response's uid []string map
	r.ID = "_:recipe"
	r.DType = []string{"Recipe"}

	pb, err := json.Marshal(r)
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

	fmt.Println("CreateRecipe mutation resp: ")
	fmt.Printf("%+v\n", res)

	r.ID = res.Uids["recipe"]

	return nil
}

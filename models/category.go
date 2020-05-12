package models

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/dgraph-io/dgo/v2"
	"github.com/dgraph-io/dgo/v2/protos/api"
)

// CategoryResponse is a struct that represents a single category. It is used exclusively
// for marshalling responses back to API clients.
type CategoryResponse struct {
	ID   string `json:"uid,omitempty"`
	Name string `json:"name,omitempty" validate:"required"`

	Recipes     []NestedRecipe     `json:"recipes,omitempty"`
	Ingredients []NestedIngredient `json:"ingredients,omitempty"`
}

// NestedCategory is a stripped down struct used when a Category is nested
// within a parent struct in an API response
type NestedCategory struct {
	ID   string `json:"uid,omitempty"`
	Name string `json:"name,omitempty" validate:"required"`
}

// ManyCategoriesResponse is a struct that represents multiple categories. It is used
// exclusively for marshalling responsesback to API clients.
type ManyCategoriesResponse struct {
	Categories []CategoryResponse `json:"categories"`
}

// Category is a struct that represents a single category
type Category struct {
	ID   string `json:"uid,omitempty"`
	Name string `json:"name,omitempty" validate:"required"`

	Recipes     []Recipe     `json:"recipe_categories,omitempty"`
	Ingredients []Ingredient `json:"ingredient_categories,omitempty"`

	DType []string `json:"dgraph.type,omitempty"`
}

// ManyCategories is a struct that represents multiple categories
type ManyCategories struct {
	ManyCategories []Category `json:"categories"`
}

// parent struct for dgraph responses
type rootCategory struct {
	Categories []Category `json:"root"`
}

// GetAllCategories will fetch all categories
func GetAllCategories(c *dgo.Dgraph) (*ManyCategories, error) {
	txn := c.NewReadOnlyTxn()

	const q = `
		{
			categories(func: type(Category)) {
				uid
				name
				dgraph.type
			}
		}
	`
	resp, err := txn.Query(context.Background(), q)
	if err != nil {
		return nil, err
	}

	categories := ManyCategories{}
	err = json.Unmarshal(resp.Json, &categories)
	if err != nil {
		return nil, err
	}

	return &categories, nil
}

// GetCategory will fetch a category via a given ID
func (category *Category) GetCategory(c *dgo.Dgraph) error {
	txn := c.NewReadOnlyTxn()

	variables := map[string]string{"$id": category.ID}
	const q = `
		query all($id: string) {
			root(func: uid($id)) @filter(type(Category)) {
				uid
				name
				dgraph.type

				recipe_categories {
					uid
					name
				}

				ingredient_categories {
					uid
					name
				}
			}
		}
	`
	resp, err := txn.QueryWithVars(context.Background(), q, variables)
	if err != nil {
		return err
	}

	root := rootCategory{}
	err = json.Unmarshal(resp.Json, &root)
	if err != nil {
		return err
	}

	*category = root.Categories[0]

	return nil
}

// CreateCategory will create a new ingredient from the given Ingredient struct
func (category *Category) CreateCategory(c *dgo.Dgraph) error {
	fmt.Println("CreateCategory() start")

	txn := c.NewTxn()
	defer txn.Discard(context.Background())

	// assign an alias ID that can be ref'd out of the response's uid []string map
	category.ID = "_:category"
	category.DType = []string{"Category"}

	pb, err := json.Marshal(category)
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

	category.ID = res.Uids["category"]

	return nil
}

// UpdateCategory will update an category via a given ID
func (category *Category) UpdateCategory(c *dgo.Dgraph) error {
	return errors.New("Not implemented")
}

// DeleteCategory will delete a category via a given ID
func (category *Category) DeleteCategory(c *dgo.Dgraph) error {
	return errors.New("Not implemented")
}

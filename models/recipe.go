package models

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/dgraph-io/dgo/v2"
	"github.com/dgraph-io/dgo/v2/protos/api"
)

// RecipeResponse is a struct that represents a single recipe. It is used exclusively
// for marshalling responses back to API clients.
type RecipeResponse struct {
	ID            string `json:"uid,omitempty"`
	Name          string `json:"name,omitempty"`
	URL           string `json:"url,omitempty"`
	Domain        string `json:"domain,omitempty"`
	Directions    string `json:"directions,omitempty"`
	PrepTime      int    `json:"prep_time,omitempty"`
	CookTime      int    `json:"cook_time,omitempty"`
	TotalServings int    `json:"total_servings,omitempty"`
	HasBeenTried  bool   `json:"has_been_tried,omitempty"`

	Ingredients    []IngredientResponse `json:"ingredients,omitempty"`
	Categories     []NestedCategory     `json:"categories,omitempty"`
	RatedBy        []NestedUser         `json:"rated_by,omitempty"`
	RatingScore    int                  `json:"rating_score,omitempty"`
	FavoritedBy    []NestedUser         `json:"favorited_by,omitempty"`
	RelatedRecipes []NestedRecipe       `json:"related_recipes,omitempty"`
	Notes          []NestedNote         `json:"notes,omitempty"`
}

// NestedRecipe is a stripped down struct used when a Recipe is nested
// within a parent struct in an API response
type NestedRecipe struct {
	ID   string `json:"uid,omitempty"`
	Name string `json:"name,omitempty" validate:"required"`
}

// ManyRecipesResponse is a struct that represents multiple recipes. It is used
// exclusively for marshalling responsesback to API clients.
type ManyRecipesResponse struct {
	Recipes []RecipeResponse `json:"recipes"`
}

// Recipe is a struct that represents a single recipe. It is used exclusively
// for unmarshalling responses from dgraph
type Recipe struct {
	ID            string `json:"uid,omitempty"`
	Name          string `json:"name,omitempty" validate:"required"`
	URL           string `json:"url,omitempty"`
	Domain        string `json:"domain,omitempty"`
	Directions    string `json:"directions,omitempty"`
	PrepTime      int    `json:"prep_time,omitempty"`
	CookTime      int    `json:"cook_time,omitempty"`
	TotalServings int    `json:"total_servings,omitempty"`
	HasBeenTried  bool   `json:"has_been_tried,omitempty"`

	Ingredients    []Ingredient `json:"ingredients,omitempty"`
	Categories     []Category   `json:"categories,omitempty"`
	RatedBy        []User       `json:"~ratings,omitempty"`
	RatingScore    int          `json:"ratings|score,omitempty"`
	FavoritedBy    []User       `json:"~favorites,omitempty"`
	RelatedRecipes []Recipe     `json:"related_recipes,omitempty"`
	Notes          []Note       `json:"~recipe,omitempty"`

	DType []string `json:"dgraph.type,omitempty"`
}

// ManyRecipes is a struct that represents multiple recipes
type ManyRecipes struct {
	Recipes []Recipe `json:"recipes"`
}

// parent struct for dgraph responses
type rootRecipe struct {
	Recipe []Recipe `json:"root"`
}

// GetAllRecipes will get all recipes
func GetAllRecipes(c *dgo.Dgraph) (*[]Recipe, error) {
	return nil, errors.New("Not implemented")
}

// GetRecipe will get a recipe via a given by ID
func (r *Recipe) GetRecipe(c *dgo.Dgraph) error {
	txn := c.NewReadOnlyTxn()

	variables := map[string]string{"$id": r.ID}
	const q = `
		query all($id: string) {
			root(func: uid($id)) @filter(type(Recipe)) {
				uid
				name
				url
				domain
				directions
				prep_time
				cook_time
				total_servings
				has_been_tried
				dgraph.type
				
				ingredients @facets {
					uid
					name
				}
				categories {
					uid
					name
				}
				~ratings @facets {
					uid
					name
				}
				~favorites {
					uid
					name
				}
				~recipe {
					uid
					text
					author {
						uid
						name
					}
				}
			}
		}
	`
	resp, err := txn.QueryWithVars(context.Background(), q, variables)
	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", resp)

	root := rootRecipe{}
	err = json.Unmarshal(resp.Json, &root)
	if err != nil {
		return err
	}

	*r = root.Recipe[0]

	return nil
}

// CreateRecipe will create a new recipe from the given Recipe struct
func (r *Recipe) CreateRecipe(c *dgo.Dgraph) error {
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

	r.ID = res.Uids["recipe"]

	return nil
}

// UpdateRecipe will update a recipe via a given by ID
func (r *Recipe) UpdateRecipe(c *dgo.Dgraph) error {
	return errors.New("Not implemented")
}

// DeleteRecipe will delete a recipe via a given by ID
func (r *Recipe) DeleteRecipe(c *dgo.Dgraph) error {
	return errors.New("Not implemented")
}

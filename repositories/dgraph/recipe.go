package dgraph

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"

	"github.com/dgraph-io/dgo/v2"
	"github.com/dgraph-io/dgo/v2/protos/api"
	"github.com/jinzhu/copier"
	"google.golang.org/grpc"

	"github.com/atticuss/chefconnect/models"
	"github.com/atticuss/chefconnect/repositories"
)

type dgraphRecipeRepo struct {
	Client *dgo.Dgraph
}

// NewDgraphRecipeRepository configures a dgraph repository for accessing
// recipe data
func NewDgraphRecipeRepository(config *Config) repositories.RecipeRepository {
	conn, _ := grpc.Dial(config.Host, grpc.WithInsecure())
	client := dgo.NewDgraphClient(api.NewDgraphClient(conn))

	return &dgraphRecipeRepo{
		Client: client,
	}
}

type manyDgraphRecipes struct {
	Recipes []dgraphRecipe `json:"recipes"`
}

type dgraphRecipe struct {
	ID            string `json:"uid,omitempty"`
	Name          string `json:"name,omitempty" validate:"required"`
	URL           string `json:"url,omitempty"`
	Domain        string `json:"domain,omitempty"`
	Directions    string `json:"directions,omitempty"`
	PrepTime      int    `json:"prep_time,omitempty"`
	CookTime      int    `json:"cook_time,omitempty"`
	TotalServings int    `json:"total_servings,omitempty"`
	HasBeenTried  bool   `json:"has_been_tried,omitempty"`

	Ingredients       []dgraphIngredient `json:"ingredients,omitempty"`
	IngredientAmounts map[string]string  `json:"ingredients|amount,omitempty"`
	Tags              []dgraphTag        `json:"tags,omitempty"`
	RatedBy           []dgraphUser       `json:"~ratings,omitempty"`
	RatingScore       map[string]string  `json:"~ratings|score,omitempty"`
	FavoritedBy       []dgraphUser       `json:"~favorites,omitempty"`
	RelatedRecipes    []dgraphRecipe     `json:"related_recipes,omitempty"`
	Notes             []models.Note      `json:"~recipe,omitempty"`

	DType []string `json:"dgraph.type,omitempty"`
}

// Due to how dgraph returns facet data, we have to do more than just copy
// between two structs. a facet is data associated with a particular edge. For
// example, that's where ingredient amounts are stored. The dgraph query for a
// recipe will return this information in the form of:
// {
//	  "name": "My Recipe",
//    "ingredients": [
//        {"uid": "0xf00", "name": "Black Beans"}
//    ],
//    "ingredients|amount" : {
//	    {"0": "1 cup, presoaked"}
//    }
// }
// In order to restructure this in a sane manner, we move the "ingredients|amount"
// field values over into each element of "ingredient". Also need to convert str
// values into ints when appropriate, as `json.Unmarshal()` refuses to cast for you
func (dRecipe *dgraphRecipe) dgraphToModel(recipe *models.Recipe) error {
	copier.Copy(&recipe, &dRecipe)

	for s_idx, value := range dRecipe.IngredientAmounts {
		i_idx, err := strconv.Atoi(s_idx)

		if err != nil {
			return err
		}

		recipe.Ingredients[i_idx].Amount = value
	}

	for s_idx, s_value := range dRecipe.RatingScore {
		i_idx, err := strconv.Atoi(s_idx)
		if err != nil {
			return err
		}

		i_value, err := strconv.Atoi(s_value)
		if err != nil {
			return err
		}

		recipe.RatedBy[i_idx].RatingScore = i_value
	}

	return nil
}

// GetAll recipes out of dgraph
func (d *dgraphRecipeRepo) GetAll() (*models.ManyRecipes, error) {
	recipes := models.ManyRecipes{}
	dRecipes := manyDgraphRecipes{}
	txn := d.Client.NewReadOnlyTxn()
	defer txn.Discard(context.Background())

	const q = `
		{
			recipes(func: type(Recipe)) {
				uid
				name

				dgraph.type
			}
		}
	`

	resp, err := txn.Query(context.Background(), q)
	if err != nil {
		return &recipes, err
	}

	err = json.Unmarshal(resp.Json, &dRecipes)
	if err != nil {
		return &recipes, err
	}

	copier.Copy(&recipes, &dRecipes)
	return &recipes, nil
}

// Get a recipe out of dgraph by ID
func (d *dgraphRecipeRepo) Get(id string) (*models.Recipe, error) {
	recipe := models.Recipe{}
	dRecipes := manyDgraphRecipes{}
	txn := d.Client.NewReadOnlyTxn()
	defer txn.Discard(context.Background())

	variables := map[string]string{"$id": id}
	const q = `
		query all($id: string) {
			recipes(func: uid($id)) @filter(type(Recipe)) {
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
				tags {
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
		return &recipe, err
	}

	err = json.Unmarshal(resp.Json, &dRecipes)
	if err != nil {
		return &recipe, err
	}

	if len(dRecipes.Recipes) > 0 {
		err := dRecipes.Recipes[0].dgraphToModel(&recipe)
		if err != nil {
			return &recipe, err
		}
	}

	return &recipe, nil
}

// Create a recipe within dgraph
func (d *dgraphRecipeRepo) Create(recipe *models.Recipe) (*models.Recipe, error) {
	dRecipe := dgraphRecipe{}
	txn := d.Client.NewTxn()
	defer txn.Discard(context.Background())

	copier.Copy(&dRecipe, recipe)
	dRecipe.ID = "_:recipe"
	dRecipe.DType = []string{"Recipe"}

	pb, err := json.Marshal(dRecipe)
	if err != nil {
		return recipe, err
	}

	mu := &api.Mutation{
		CommitNow: true,
		SetJson:   pb,
	}

	res, err := txn.Mutate(context.Background(), mu)
	if err != nil {
		return recipe, err
	}

	recipe.ID = res.Uids["recipe"]

	return recipe, nil
}

// Update a recipe within dgraph
func (d *dgraphRecipeRepo) Update(recipe *models.Recipe) (*models.Recipe, error) {
	dRecipe := dgraphRecipe{}
	txn := d.Client.NewTxn()
	defer txn.Discard(context.Background())

	copier.Copy(&dRecipe, recipe)
	dRecipe.DType = []string{"Recipe"}

	pb, err := json.Marshal(dRecipe)
	if err != nil {
		return recipe, err
	}

	mu := &api.Mutation{
		CommitNow: true,
		SetJson:   pb,
	}

	_, err = txn.Mutate(context.Background(), mu)
	if err != nil {
		return recipe, err
	}

	return recipe, nil
}

// SetTags for a recipe in drgraph
func (d *dgraphRecipeRepo) SetTags(recipe *models.Recipe) (*models.Recipe, error) {
	dRecipe := dgraphRecipe{}
	txn := d.Client.NewTxn()
	defer txn.Discard(context.Background())

	copier.Copy(&dRecipe, recipe)
	dRecipe.DType = []string{"Recipe"}

	mu := &api.Mutation{}
	dgo.DeleteEdges(mu, dRecipe.ID, "tags")
	mu.CommitNow = true

	_, err := d.Client.NewTxn().Mutate(context.Background(), mu)
	if err != nil {
		return recipe, err
	}

	pb, err := json.Marshal(dRecipe)
	if err != nil {
		return recipe, err
	}

	mu = &api.Mutation{
		CommitNow: true,
		SetJson:   pb,
	}

	_, err = txn.Mutate(context.Background(), mu)
	if err != nil {
		return recipe, err
	}

	return recipe, nil
}

// Delete a recipe from dgraph
func (d *dgraphRecipeRepo) Delete(id string) error {
	return errors.New("not implemented")
}

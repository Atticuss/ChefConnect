package dgraph

import (
	"context"
	"encoding/json"
	"errors"

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
	IngredientAmounts map[int]string     `json:"ingredients|amount,omitempty"`
	Categories        []dgraphCategory   `json:"categories,omitempty"`
	RatedBy           []dgraphUser       `json:"~ratings,omitempty"`
	RatingScore       map[int]int        `json:"~ratings|score,omitempty"`
	FavoritedBy       []dgraphUser       `json:"~favorites,omitempty"`
	RelatedRecipes    []dgraphRecipe     `json:"related_recipes,omitempty"`
	Notes             []models.Note      `json:"~recipe,omitempty"`

	DType []string `json:"dgraph.type,omitempty"`
}

func (dRecipe *dgraphRecipe) dgraphToModel(recipe *models.Recipe) {
	copier.Copy(&recipe, &dRecipe)

	for idx, value := range dRecipe.IngredientAmounts {
		recipe.Ingredients[idx].Amount = value
	}

	for idx, value := range dRecipe.RatingScore {
		recipe.RatedBy[idx].RatingScore = value
	}
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
		return &recipe, err
	}

	err = json.Unmarshal(resp.Json, &dRecipes)
	if err != nil {
		return &recipe, err
	}

	if len(dRecipes.Recipes) > 0 {
		dRecipes.Recipes[0].dgraphToModel(&recipe)
		return &recipe, nil
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

// Delete a recipe from dgraph
func (d *dgraphRecipeRepo) Delete(id string) error {
	return errors.New("Not implemented")
}

package dgraph

import (
	"context"
	"encoding/json"

	"github.com/dgraph-io/dgo/v2"
	"github.com/dgraph-io/dgo/v2/protos/api"
	"github.com/jinzhu/copier"
	"google.golang.org/grpc"

	"github.com/atticuss/chefconnect/models"
	"github.com/atticuss/chefconnect/repositories"
)

type dgraphIngredientRepo struct {
	Client *dgo.Dgraph
}

// NewDgraphIngredientRepository configures a dgraph repository for accessing
// ingredient data
func NewDgraphIngredientRepository(config *Config) repositories.IngredientRepository {
	conn, _ := grpc.Dial(config.Host, grpc.WithInsecure())
	client := dgo.NewDgraphClient(api.NewDgraphClient(conn))

	return &dgraphIngredientRepo{
		Client: client,
	}
}

type manyDgraphIngredients struct {
	Ingredients []dgraphIngredient `json:"ingredients"`
}

type dgraphIngredient struct {
	ID   string `json:"uid,omitempty"`
	Name string `json:"name,omitempty" validate:"required"`

	Recipes []models.Recipe `json:"~ingredients,omitempty"`
	Tags    []dgraphTag     `json:"tags,omitempty"`

	DType []string `json:"dgraph.type,omitempty"`
}

// GetAll ingredients out of dgraph
func (d *dgraphIngredientRepo) GetAll() (*models.ManyIngredients, error) {
	ingredients := models.ManyIngredients{}
	dIngredients := manyDgraphIngredients{}
	txn := d.Client.NewReadOnlyTxn()
	defer txn.Discard(context.Background())

	const q = `
		{
			ingredients(func: type(Ingredient)) {
				uid
				name

				dgraph.type
			}
		}
	`

	resp, err := txn.Query(context.Background(), q)
	if err != nil {
		return &ingredients, err
	}

	err = json.Unmarshal(resp.Json, &dIngredients)
	if err != nil {
		return &ingredients, err
	}

	copier.Copy(&ingredients, &dIngredients)
	return &ingredients, nil
}

// Get an ingredient out of dgraph by ID
func (d *dgraphIngredientRepo) Get(id string) (*models.Ingredient, error) {
	ingredient := models.Ingredient{}
	dIngredients := manyDgraphIngredients{}
	txn := d.Client.NewReadOnlyTxn()
	defer txn.Discard(context.Background())

	variables := map[string]string{"$id": id}
	const q = `
		query all($id: string) {
			ingredients(func: uid($id)) @filter(type(Ingredient)) {
				uid
				name
				dgraph.type

				tags {
					uid
					name
				}

				~ingredients {
					uid
					name
				}
			}
		}
	`

	resp, err := txn.QueryWithVars(context.Background(), q, variables)
	if err != nil {
		return &ingredient, err
	}

	err = json.Unmarshal(resp.Json, &dIngredients)
	if err != nil {
		return &ingredient, err
	}

	if len(dIngredients.Ingredients) > 0 {
		copier.Copy(&ingredient, &dIngredients.Ingredients[0])
		return &ingredient, nil
	}

	return &ingredient, nil
}

// Create an ingredient within dgraph
func (d *dgraphIngredientRepo) Create(ingredient *models.Ingredient) (*models.Ingredient, error) {
	dIngredient := dgraphIngredient{}
	txn := d.Client.NewTxn()
	defer txn.Discard(context.Background())

	copier.Copy(&dIngredient, ingredient)
	dIngredient.ID = "_:ingredient"
	dIngredient.DType = []string{"Ingredient"}

	pb, err := json.Marshal(dIngredient)
	if err != nil {
		return ingredient, err
	}

	mu := &api.Mutation{
		CommitNow: true,
		SetJson:   pb,
	}

	res, err := txn.Mutate(context.Background(), mu)
	if err != nil {
		return ingredient, err
	}

	ingredient.ID = res.Uids["ingredient"]

	return ingredient, nil
}

// Update an ingredient within dgraph
func (d *dgraphIngredientRepo) Update(ingredient *models.Ingredient) (*models.Ingredient, error) {
	dIngredient := dgraphIngredient{}
	txn := d.Client.NewTxn()
	defer txn.Discard(context.Background())

	copier.Copy(&dIngredient, ingredient)
	dIngredient.DType = []string{"Ingredient"}

	mu := &api.Mutation{
		CommitNow: true,
	}
	dgo.DeleteEdges(mu, dIngredient.ID, "tags")

	_, err := d.Client.NewTxn().Mutate(context.Background(), mu)
	if err != nil {
		return ingredient, err
	}

	pb, err := json.Marshal(dIngredient)
	if err != nil {
		return ingredient, err
	}

	mu = &api.Mutation{
		CommitNow: true,
		SetJson:   pb,
	}

	_, err = txn.Mutate(context.Background(), mu)
	if err != nil {
		return ingredient, err
	}

	return ingredient, nil
}

// Delete an ingredient from dgraph
func (d *dgraphIngredientRepo) Delete(id string) error {
	txn := d.Client.NewTxn()
	defer txn.Discard(context.Background())

	m := map[string]string{"uid": id}
	pb, err := json.Marshal(m)
	if err != nil {
		return err
	}

	mu := &api.Mutation{
		CommitNow:  true,
		DeleteJson: pb,
	}

	_, err = txn.Mutate(context.Background(), mu)
	if err != nil {
		return err
	}

	return nil
}

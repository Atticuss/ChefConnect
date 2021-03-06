package dgraph

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/dgraph-io/dgo/v2"
	"github.com/dgraph-io/dgo/v2/protos/api"
	"github.com/jinzhu/copier"

	"github.com/atticuss/chefconnect/models"
)

type manyDgraphIngredients struct {
	Ingredients []dgraphIngredient `json:"ingredients"`
}

type dgraphIngredient struct {
	ID   string `json:"uid,omitempty"`
	Name string `json:"name,omitempty" validate:"required"`

	Recipes []dgraphRecipe `json:"~ingredients,omitempty"`
	Tags    []dgraphTag    `json:"ingredient_tags,omitempty"`

	DType []string `json:"dgraph.type,omitempty"`
}

// GetAllIngredients out of dgraph
func (d *dgraphRepo) GetAllIngredients() (*models.ManyIngredients, error) {
	ingredients := models.ManyIngredients{}
	dIngredients := manyDgraphIngredients{}
	ctx := d.buildAuthContext(context.Background())
	txn := d.Client.NewReadOnlyTxn()
	defer txn.Discard(ctx)

	const q = `
		{
			ingredients(func: type(Ingredient)) {
				uid
				name

				dgraph.type
			}
		}
	`

	resp, err := txn.Query(ctx, q)
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

// GetIngredient out of dgraph by ID
func (d *dgraphRepo) GetIngredient(id string) (*models.Ingredient, error) {
	ingredient := models.Ingredient{}
	dIngredients := manyDgraphIngredients{}
	ctx := d.buildAuthContext(context.Background())
	txn := d.Client.NewReadOnlyTxn()
	defer txn.Discard(ctx)

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

	resp, err := txn.QueryWithVars(ctx, q, variables)
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

// CreateIngredient within dgraph
func (d *dgraphRepo) CreateIngredient(ingredient *models.Ingredient) (*models.Ingredient, error) {
	dIngredient := dgraphIngredient{}
	ctx := d.buildAuthContext(context.Background())
	txn := d.Client.NewTxn()
	defer txn.Discard(ctx)

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

	res, err := txn.Mutate(ctx, mu)
	if err != nil {
		return ingredient, err
	}

	ingredient.ID = res.Uids["ingredient"]

	return ingredient, nil
}

// UpdateIngredient within dgraph
func (d *dgraphRepo) UpdateIngredient(ingredient *models.Ingredient) (*models.Ingredient, error) {
	dIngredient := dgraphIngredient{}
	ctx := d.buildAuthContext(context.Background())
	txn := d.Client.NewTxn()
	defer txn.Discard(ctx)

	copier.Copy(&dIngredient, ingredient)
	dIngredient.DType = []string{"Ingredient"}

	mu := &api.Mutation{
		CommitNow: true,
	}
	dgo.DeleteEdges(mu, dIngredient.ID, "tags")

	_, err := d.Client.NewTxn().Mutate(ctx, mu)
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

	_, err = txn.Mutate(ctx, mu)
	if err != nil {
		return ingredient, err
	}

	return ingredient, nil
}

// DeleteIngredient from dgraph
func (d *dgraphRepo) DeleteIngredient(id string) error {
	ctx := d.buildAuthContext(context.Background())
	txn := d.Client.NewTxn()
	defer txn.Discard(ctx)

	m := map[string]string{"uid": id}
	pb, err := json.Marshal(m)
	if err != nil {
		return err
	}

	mu := &api.Mutation{
		CommitNow:  true,
		DeleteJson: pb,
	}

	_, err = txn.Mutate(ctx, mu)
	if err != nil {
		return err
	}

	return nil
}

// SearchIngredientByName from dgraph
func (d *dgraphRepo) SearchIngredientByName(searchTerm string) (*models.ManyIngredients, error) {
	ingredients := models.ManyIngredients{}
	dIngredients := manyDgraphIngredients{}
	ctx := d.buildAuthContext(context.Background())
	txn := d.Client.NewReadOnlyTxn()
	defer txn.Discard(ctx)

	if len(searchTerm) < 3 {
		return &ingredients, errors.New("`searchTerm` must be at least 3 characters long")
	}

	searchTerm = fmt.Sprintf("%s%s%s", "/.*", searchTerm, ".*/i")
	variables := map[string]string{"$searchTerm": searchTerm}
	const q = `
		query all($searchTerm: string) {
			ingredients(func: regexp(name, $searchTerm)) @filter(type(Ingredient)) {
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

	resp, err := txn.QueryWithVars(ctx, q, variables)
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

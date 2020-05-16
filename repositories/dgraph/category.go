package dgraph

import (
	"context"
	"encoding/json"

	"github.com/dgraph-io/dgo/v2"
	"github.com/dgraph-io/dgo/v2/protos/api"
	"github.com/jinzhu/copier"

	"github.com/atticuss/chefconnect/models"
	"github.com/atticuss/chefconnect/repositories"
)

type dgraphCategoryRepo struct {
	Client *dgo.Dgraph
}

// NewDgraphCategoryRepository configures a dgraph repository for accessing
// category data
func NewDgraphCategoryRepository(db *dgo.Dgraph) repositories.CategoryRepository {
	return &dgraphCategoryRepo{
		Client: db,
	}
}

type manyDgraphCategories struct {
	Categories []dgraphCategory `json:"categories"`
}

type dgraphCategory struct {
	ID   string `json:"uid,omitempty"`
	Name string `json:"name,omitempty" validate:"required"`

	Recipes     []models.Recipe     `json:"~recipe_categories,omitempty"`
	Ingredients []models.Ingredient `json:"~ingredient_categories,omitempty"`

	DType []string `json:"dgraph.type,omitempty"`
}

// GetAll categories out of dgraph
func (d *dgraphCategoryRepo) GetAll() (*models.ManyCategories, error) {
	dCategories := manyDgraphCategories{}
	categories := models.ManyCategories{}
	txn := d.Client.NewReadOnlyTxn()
	defer txn.Discard(context.Background())

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
		return &categories, err
	}

	err = json.Unmarshal(resp.Json, &dCategories)
	if err != nil {
		return &categories, err
	}

	copier.Copy(&categories, &dCategories)

	return &categories, nil
}

// Get a category out of dgraph by ID
func (d *dgraphCategoryRepo) Get(id string) (*models.Category, error) {
	dCategories := manyDgraphCategories{}
	category := models.Category{}
	txn := d.Client.NewReadOnlyTxn()
	defer txn.Discard(context.Background())

	variables := map[string]string{"$id": id}
	const q = `
		query all($id: string) {
			categories(func: uid($id)) @filter(type(Category)) {
				uid
				name
				dgraph.type

				~recipe_categories {
					uid
					name
				}

				~ingredient_categories {
					uid
					name
				}
			}
		}
	`

	resp, err := txn.QueryWithVars(context.Background(), q, variables)
	if err != nil {
		return &category, err
	}

	err = json.Unmarshal(resp.Json, &dCategories)
	if err != nil {
		return &category, err
	}

	if len(dCategories.Categories) > 0 {
		copier.Copy(&category, &dCategories.Categories[0])
		return &category, nil
	}

	return &category, nil
}

// Create a category within dgraph
func (d *dgraphCategoryRepo) Create(category *models.Category) (*models.Category, error) {
	dCategory := dgraphCategory{}
	txn := d.Client.NewTxn()
	defer txn.Discard(context.Background())

	copier.Copy(&dCategory, category)

	// assign an alias ID that can be ref'd out of the response's uid map[string]string
	dCategory.ID = "_:category"
	dCategory.DType = []string{"Category"}

	pb, err := json.Marshal(dCategory)
	if err != nil {
		return category, err
	}

	mu := &api.Mutation{
		CommitNow: true,
		SetJson:   pb,
	}

	res, err := txn.Mutate(context.Background(), mu)
	if err != nil {
		return category, err
	}

	category.ID = res.Uids["category"]

	return category, nil
}

// Update a category within dgraph by ID
func (d *dgraphCategoryRepo) Update(category *models.Category) (*models.Category, error) {
	dCategory := dgraphCategory{}
	txn := d.Client.NewTxn()
	defer txn.Discard(context.Background())

	copier.Copy(&dCategory, category)

	dCategory.DType = []string{"Category"}

	pb, err := json.Marshal(dCategory)
	if err != nil {
		return category, err
	}

	mu := &api.Mutation{
		CommitNow: true,
		SetJson:   pb,
	}

	_, err = txn.Mutate(context.Background(), mu)
	if err != nil {
		return category, err
	}

	return category, nil
}

// Delete a category from dgraph by ID
func (d *dgraphCategoryRepo) Delete(id string) error {
	txn := d.Client.NewTxn()

	variables := map[string]string{"uid": id}
	pb, err := json.Marshal(variables)
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

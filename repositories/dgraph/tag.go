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

type dgraphTagRepo struct {
	Client *dgo.Dgraph
}

// NewDgraphTagRepository configures a dgraph repository for accessing
// tag data
func NewDgraphTagRepository(config *Config) repositories.TagRepository {
	conn, _ := grpc.Dial(config.Host, grpc.WithInsecure())
	client := dgo.NewDgraphClient(api.NewDgraphClient(conn))

	return &dgraphTagRepo{
		Client: client,
	}
}

type manyDgraphTags struct {
	Tags []dgraphTag `json:"tags"`
}

type dgraphTag struct {
	ID   string `json:"uid,omitempty"`
	Name string `json:"name,omitempty" validate:"required"`

	Recipes     []models.Recipe     `json:"~recipe_tags,omitempty"`
	Ingredients []models.Ingredient `json:"~ingredient_tags,omitempty"`

	DType []string `json:"dgraph.type,omitempty"`
}

// GetAll tags out of dgraph
func (d *dgraphTagRepo) GetAll() (*models.ManyTags, error) {
	dTags := manyDgraphTags{}
	tags := models.ManyTags{}
	txn := d.Client.NewReadOnlyTxn()
	defer txn.Discard(context.Background())

	const q = `
		{
			tags(func: type(Tag)) {
				uid
				name
				dgraph.type
			}
		}
	`

	resp, err := txn.Query(context.Background(), q)
	if err != nil {
		return &tags, err
	}

	err = json.Unmarshal(resp.Json, &dTags)
	if err != nil {
		return &tags, err
	}

	copier.Copy(&tags, &dTags)

	return &tags, nil
}

// Get a tag out of dgraph by ID
func (d *dgraphTagRepo) Get(id string) (*models.Tag, error) {
	dTags := manyDgraphTags{}
	tag := models.Tag{}
	txn := d.Client.NewReadOnlyTxn()
	defer txn.Discard(context.Background())

	variables := map[string]string{"$id": id}
	const q = `
		query all($id: string) {
			tags(func: uid($id)) @filter(type(Tag)) {
				uid
				name
				dgraph.type

				~recipe_tags {
					uid
					name
				}

				~ingredient_tags {
					uid
					name
				}
			}
		}
	`

	resp, err := txn.QueryWithVars(context.Background(), q, variables)
	if err != nil {
		return &tag, err
	}

	err = json.Unmarshal(resp.Json, &dTags)
	if err != nil {
		return &tag, err
	}

	if len(dTags.Tags) > 0 {
		copier.Copy(&tag, &dTags.Tags[0])
		return &tag, nil
	}

	return &tag, nil
}

// Create a tag within dgraph
func (d *dgraphTagRepo) Create(tag *models.Tag) (*models.Tag, error) {
	dTag := dgraphTag{}
	txn := d.Client.NewTxn()
	defer txn.Discard(context.Background())

	copier.Copy(&dTag, tag)

	// assign an alias ID that can be ref'd out of the response's uid map[string]string
	dTag.ID = "_:tag"
	dTag.DType = []string{"Tag"}

	pb, err := json.Marshal(dTag)
	if err != nil {
		return tag, err
	}

	mu := &api.Mutation{
		CommitNow: true,
		SetJson:   pb,
	}

	res, err := txn.Mutate(context.Background(), mu)
	if err != nil {
		return tag, err
	}

	tag.ID = res.Uids["tag"]

	return tag, nil
}

// Update a tag within dgraph by ID
func (d *dgraphTagRepo) Update(tag *models.Tag) (*models.Tag, error) {
	dTag := dgraphTag{}
	txn := d.Client.NewTxn()
	defer txn.Discard(context.Background())

	copier.Copy(&dTag, tag)

	dTag.DType = []string{"Tag"}

	pb, err := json.Marshal(dTag)
	if err != nil {
		return tag, err
	}

	mu := &api.Mutation{
		CommitNow: true,
		SetJson:   pb,
	}

	_, err = txn.Mutate(context.Background(), mu)
	if err != nil {
		return tag, err
	}

	return tag, nil
}

// Delete a tag from dgraph by ID
func (d *dgraphTagRepo) Delete(id string) error {
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

package dgraph

import (
	"context"
	"encoding/json"

	"github.com/dgraph-io/dgo/v2/protos/api"
	"github.com/jinzhu/copier"

	"github.com/atticuss/chefconnect/models"
)

type manyDgraphTags struct {
	Tags []dgraphTag `json:"tags"`
}

type dgraphTag struct {
	ID   string `json:"uid,omitempty"`
	Name string `json:"name,omitempty" validate:"required"`

	Recipes     []dgraphRecipe     `json:"~recipe_tags,omitempty"`
	Ingredients []dgraphIngredient `json:"~ingredient_tags,omitempty"`

	DType []string `json:"dgraph.type,omitempty"`
}

// GetAllTags out of dgraph
func (d *dgraphRepo) GetAllTags() (*models.ManyTags, error) {
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

// GetTag out of dgraph by ID
func (d *dgraphRepo) GetTag(id string) (*models.Tag, error) {
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

// CreateTag within dgraph
func (d *dgraphRepo) CreateTag(tag *models.Tag) (*models.Tag, error) {
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

// UpdateTag within dgraph by ID
func (d *dgraphRepo) UpdateTag(tag *models.Tag) (*models.Tag, error) {
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

// DeleteTag from dgraph by ID
func (d *dgraphRepo) DeleteTag(id string) error {
	txn := d.Client.NewTxn()
	defer txn.Discard(context.Background())

	readOnlyTxn := d.Client.NewReadOnlyTxn()
	defer txn.Discard(context.Background())

	// Nuke all our reverse edges by the parent node
	dTags := models.ManyTags{}
	const q = `
		query all($id: string) {
			tags(func: uid($id)) @filter(type(Tag)) {
				~recipe_tags {
					uid
				}
				~ingredient_tags {
					uid
				}
			}
		}
	`

	resp, err := readOnlyTxn.Query(context.Background(), q)
	if err != nil {
		return err
	}

	err = json.Unmarshal(resp.Json, &dTags)
	if err != nil {
		return err
	}

	// Doesn't exist, just return now
	if len(dTags.Tags) == 0 {
		return nil
	}

	for _, dRecipe := range dTags.Tags[0].Recipes {
		mu := &api.Mutation{
			Del: []*api.NQuad{
				{
					Subject:   dRecipe.ID,
					Predicate: "recipe_tags",
					ObjectId:  id,
				},
			},
		}

		_, err = txn.Mutate(context.Background(), mu)
		if err != nil {
			return err
		}
	}

	for _, dIngredient := range dTags.Tags[0].Ingredients {
		mu := &api.Mutation{
			Del: []*api.NQuad{
				{
					Subject:   dIngredient.ID,
					Predicate: "ingredient_tags",
					ObjectId:  id,
				},
			},
		}

		_, err = txn.Mutate(context.Background(), mu)
		if err != nil {
			return err
		}
	}

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

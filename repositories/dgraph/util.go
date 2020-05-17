package dgraph

import (
	"context"

	"github.com/dgraph-io/dgo/v2"
	"github.com/dgraph-io/dgo/v2/protos/api"

	"github.com/atticuss/chefconnect/repositories"
)

type dgraphUtilRepo struct {
	Client *dgo.Dgraph
}

// NewDgraphRepositoryUtility configures a dgraph repository for accessing
// various utility functions, typically leveraged during testing and
// application initialization
func NewDgraphRepositoryUtility(db *dgo.Dgraph) repositories.RepositoryUtility {
	return &dgraphUtilRepo{
		Client: db,
	}
}

func (d *dgraphUtilRepo) InitializeSchema() error {
	op := &api.Operation{}

	op.Schema = `
		name: string @index(term) .
		url: string .
		domain: string .
		directions: string .
		ingredients: [uid] @reverse .
		ingredient_categories: [uid] @reverse .
		recipe_categories: [uid] @reverse .
		prep_time: int @index(int) .
		cook_time: int @index(int) .
		total_servings: int .
		related_recipes: [uid] @reverse .
		ratings: [uid] @reverse .
		score: int @index(int) .
		username: string @index(exact) .
		password: string .
		favorites: [uid] @reverse .
		user_notes: [uid] @reverse .
		recipe_notes: [uid] @reverse .
		has_been_tried: bool @index(bool) .
		text: string .
		index: int .
		amount: string .

		recipe: [uid] @reverse .
		author: [uid] @reverse .

		type Ingredient {
			name
			<~ingredients>
			ingredient_categories

			amount
		}

		type Category {
			name
			<~ingredient_categories>
			<~recipe_categories>
		}

		type Recipe {
			name
			url
			domain
			directions
			ingredients
			prep_time
			cook_time
			total_servings
			related_recipes
			recipe_categories
			has_been_tried
			
			<~recipe>
			<~ratings>
			<~favorites>
			<~related_recipes>
		}

		type Note {
			text
			author
			recipe
		}

		type User {
			name
			username
			password
			favorites
			<~author>
			ratings
		}
	`

	if err := d.Client.Alter(context.Background(), op); err != nil {
		return err
	}

	return nil
}

func (d *dgraphUtilRepo) ClearDatastore() error {
	op := &api.Operation{DropAll: true}

	if err := d.Client.Alter(context.Background(), op); err != nil {
		return err
	}

	return nil
}

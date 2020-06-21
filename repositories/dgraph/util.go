package dgraph

import (
	"context"
	"fmt"

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
		roles: [uid] @reverse .
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
			roles
			favorites
			<~author>
			ratings
		}

		type Role {
			name
			<~roles>
		}
	`

	if err := d.Client.Alter(context.Background(), op); err != nil {
		return err
	}

	return nil
}

func (d *dgraphUtilRepo) InitializeBaseData() error {
	txn := d.Client.NewTxn()
	defer txn.Discard(context.Background())

	//$2a$14$zR/r6hmGbPk1mh1G8fsvJOE/iKfhosK5YjVoiA51zgKmDnp6lETja -> Password1!
	nquads := `
		_:role_admin <name> "SiteAdmin" .
		_:role_admin <dgraph.type> "Role" .

		_:role_user <name> "User" .
		_:role_user <dgraph.type> "Role" .

		_:user_jay <name> "Jay Sea" .
		_:user_jay <username> "jay.sea" .
		_:user_jay <password> "$2a$14$zR/r6hmGbPk1mh1G8fsvJOE/iKfhosK5YjVoiA51zgKmDnp6lETja" .
		_:user_jay <roles> _:role_user .
		_:user_jay <dgraph.type> "User" .

		_:user_el <name> "El Dubs" .
		_:user_el <username> "el.dubs" .
		_:user_el <password> "$2a$14$zR/r6hmGbPk1mh1G8fsvJOE/iKfhosK5YjVoiA51zgKmDnp6lETja" .
		_:user_el <roles> _:role_admin .
		_:user_el <dgraph.type> "User" .
	`

	mu := &api.Mutation{
		CommitNow: true,
		SetNquads: []byte(nquads),
	}

	res, err := txn.Mutate(context.Background(), mu)
	fmt.Printf("result: %+v\n", res)
	if err != nil {
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

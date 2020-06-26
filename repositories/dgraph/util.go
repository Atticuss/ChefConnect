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
		score: int .
		owner: [uid] @reverse .
		username: string @index(exact) .
		password: string .
		roles: [uid] @reverse .
		favorites: [uid] @reverse .
		user_notes: [uid] @reverse .
		recipe_notes: [uid] @reverse .
		has_been_tried: bool @index(bool) .
		text: string .
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
			owner
			
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
			<~owner>
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

		_:cat_fake_meat <name> "Fake Meat" .
		_:cat_fake_meat <dgraph.type> "Category" .
		
		_:cat_condiment <name> "Condiment" .
		_:cat_condiment <dgraph.type> "Category" .

		_:ing_soy_curls <name> "Soy Curls" .
		_:ing_soy_curls <ingredient_categories> _:cat_fake_meat .
		_:ing_soy_curls <dgraph.type> "Ingredient" .

		_:ing_buffalo <name> "Buffalo Sauce" .
		_:ing_buffalo <ingredient_categories> _:cat_condiment .
		_:ing_buffalo <dgraph.type> "Ingredient" .

		_:ing_black_beans <name> "Black Beans" .
		_:ing_black_beans <dgraph.type> "Ingredient" .

		_:ing_pasta <name> "Chickpea Pasta" .
		_:ing_pasta <dgraph.type> "Ingredient" .

		_:rec_soy_bowl <name> "Soy Curl Bowl" .
		_:rec_soy_bowl <owner> _:user_jay .
		_:rec_soy_bowl <url> "https://some.bullshit/terrible_recipe.pdf" .
		_:rec_soy_bowl <domain> "some.bullshit" .
		_:rec_soy_bowl <directions> "Presoak the soy curls for 10 min then sautee with buffalo sauce. Make black beans and chickpea pasta. Mix that ish together and devour, bonus points if you eat it quicker than you prepared it." .
		_:rec_soy_bowl <ingredients> _:ing_soy_curls (amount="1 cup, presoaked") .
		_:rec_soy_bowl <ingredients> _:ing_black_beans (amount="1/2 cup, presoaked") .
		_:rec_soy_bowl <ingredients> _:ing_pasta (amount="1 cup, presoaked") .
		_:rec_soy_bowl <ingredients> _:ing_buffalo (amount="3 tbsp") .
		_:rec_soy_bowl <prep_time> "10" .
		_:rec_soy_bowl <cook_time> "15" .
		_:rec_soy_bowl <total_servings> "2" .
		_:rec_soy_bowl <has_been_tried> "False" .
		_:rec_soy_bowl <dgraph.type> "Recipe" .

		_:note_jay_soy <text> "pretty damn good" .
		_:note_jay_soy <author> _:user_jay .
		_:note_jay_soy <recipe> _:rec_soy_bowl .
		_:note_jay_soy <dgraph.type> "Note" .

		_:user_jay <favorites> _:rec_soy_bowl .
		_:user_jay <ratings> _:rec_soy_bowl (score="4") .
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

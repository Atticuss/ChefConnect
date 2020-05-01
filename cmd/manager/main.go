package main

import (
	"context"
	"fmt"
	"log"

	//"encoding/json"

	"github.com/atticuss/chefconnect/models"
	"github.com/dgraph-io/dgo/v2"
	"github.com/dgraph-io/dgo/v2/protos/api"
	"google.golang.org/grpc"
)

func buildSchema(c *dgo.Dgraph) {
	ctx := context.Background()
	op := &api.Operation{}

	fmt.Println("rebuilding schema")

	op.Schema = `
		name: string @index(term) .
		url: string .
		domain: string .
		directions: string .
		ingredients: [uid] @reverse .
		categories: [uid] @reverse .
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
			categories

			amount
		}

		type Category {
			name
			<~categories>
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
			categories
			ratings
			has_been_tried
			
			<~recipe>

			score

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

	if err := c.Alter(ctx, op); err != nil {
		log.Fatal(err)
	}
}

func clear(c *dgo.Dgraph) {
	ctx := context.Background()
	op := &api.Operation{DropAll: true}

	fmt.Println("nuking db")

	if err := c.Alter(ctx, op); err != nil {
		log.Fatal(err)
	}
}

func initCategories(c *dgo.Dgraph) *[]models.Category {
	fmt.Println("init: categories")

	cat1 := models.Category{
		ID:   "_:breakfast",
		Name: "Breakfast",
	}

	cat2 := models.Category{
		ID:   "_:condiment",
		Name: "Condiment",
	}

	cat3 := models.Category{
		ID:   "_:fakemeat",
		Name: "Fake Meat",
	}

	categories := []models.Category{cat1, cat2, cat3}

	for idx, cat := range categories {
		if err := cat.CreateCategory(c); err != nil {
			log.Fatal(err)
		}
		categories[idx] = cat
	}

	return &categories
}

func initIngredients(c *dgo.Dgraph, categories *[]models.Category) *[]models.Ingredient {
	fmt.Println("init: ingredients")
	fmt.Printf("using category data: %+v\n", categories)

	ing1 := models.Ingredient{
		Name: "Black Beans",
	}

	ing2 := models.Ingredient{
		Name:       "Soy Curls",
		Categories: []models.Category{{ID: (*categories)[2].ID}},
	}

	ing3 := models.Ingredient{
		Name: "Chickpea Pasta",
	}

	ing4 := models.Ingredient{
		Name:       "Buffalo Sauce",
		Categories: []models.Category{{ID: (*categories)[1].ID}},
	}

	ingredients := []models.Ingredient{ing1, ing2, ing3, ing4}

	for idx, ing := range ingredients {
		if err := ing.CreateIngredient(c); err != nil {
			log.Fatal(err)
		}
		ingredients[idx] = ing
	}

	return &ingredients
}

func initRecipes(c *dgo.Dgraph, categories *[]models.Category, ingredients *[]models.Ingredient) *[]models.Recipe {
	fmt.Println("init: recipes")

	recipe1 := models.Recipe{
		Name:       "Soy Curl Bowl",
		URL:        "https://foo.com/some-bullshit",
		Domain:     "foo.com",
		Directions: "Prepare the chickpea pasta as directed. Soak the soy curls for 10 minutes. Drain and stir fry in an oiled pan for several minutes. Once most of the water has evaporated, mix in the buffalo sauce and sautee for a few more minutes. Mix with beans and pasta. Devour immediately.",
		Ingredients: []models.Ingredient{
			{ID: (*ingredients)[0].ID, Amount: "1 cup", DType: []string{"Ingredient"}},
			{ID: (*ingredients)[1].ID, Amount: "1 cup, presoaked", DType: []string{"Ingredient"}},
			{ID: (*ingredients)[2].ID, Amount: "1 cup, precooked", DType: []string{"Ingredient"}},
			{ID: (*ingredients)[3].ID, Amount: "3 tbsp", DType: []string{"Ingredient"}},
		},
		PrepTime:      10,
		CookTime:      12,
		TotalServings: 2,
	}

	recipes := []models.Recipe{recipe1}

	for idx, rec := range recipes {
		if err := rec.CreateRecipe(c); err != nil {
			log.Fatal(err)
		}
		recipes[idx] = rec
	}

	return &recipes
}

func initUsers(c *dgo.Dgraph, recipes *[]models.Recipe) *[]models.User {
	fmt.Println("init: users")

	user1 := models.User{
		Name:      "Jay Sea",
		Username:  "jay.sea@gmail.com",
		Password:  "Password1!",
		Favorites: *recipes,
		Ratings:   []models.Recipe{{ID: (*recipes)[0].ID, RatingScore: 4}},
	}

	users := []models.User{user1}

	for idx, u := range users {
		if err := u.CreateUser(c); err != nil {
			log.Fatal(err)
		}
		users[idx] = u
	}

	return &users
}

func initNotes(c *dgo.Dgraph, recipes *[]models.Recipe, users *[]models.User) *[]models.Note {
	fmt.Println("init: notes")

	note1 := models.Note{
		Text:   "this shit dank, yo",
		Recipe: models.Recipe{ID: (*recipes)[0].ID},
		User:   models.User{ID: (*users)[0].ID},
		DType:  []string{"Note"},
	}

	notes := []models.Note{note1}

	for idx, n := range notes {
		if err := n.CreateNote(c); err != nil {
			log.Fatal(err)
		}
		notes[idx] = n
	}

	return &notes
}

func initData(c *dgo.Dgraph) {
	fmt.Println("init: data")

	categories := initCategories(c)
	ingredients := initIngredients(c, categories)
	recipes := initRecipes(c, categories, ingredients)
	users := initUsers(c, recipes)
	initNotes(c, recipes, users)
}

func main() {
	conn, err := grpc.Dial("ec2-34-238-150-16.compute-1.amazonaws.com:9080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	dgraphClient := dgo.NewDgraphClient(api.NewDgraphClient(conn))
	clear(dgraphClient)
	buildSchema(dgraphClient)
	initData(dgraphClient)
}

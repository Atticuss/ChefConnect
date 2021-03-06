package dgraph

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/dgraph-io/dgo/v2"
	"github.com/dgraph-io/dgo/v2/protos/api"
	"github.com/jinzhu/copier"

	"github.com/atticuss/chefconnect/models"
)

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

	Ingredients          []dgraphIngredient `json:"ingredients,omitempty"`
	IngredientAmounts    map[string]string  `json:"ingredients|amount,omitempty"`
	Tags                 []dgraphTag        `json:"recipe_tags,omitempty"`
	RatedBy              []dgraphUser       `json:"~ratings,omitempty"`
	RatingScore          map[string]string  `json:"~ratings|score,omitempty"`
	FavoritedBy          []dgraphUser       `json:"~favorites,omitempty"`
	Owner                []dgraphUser       `json:"owner,omitempty"`
	RelatedRecipesParent []dgraphRecipe     `json:"related_recipes,omitempty"`
	RelatedRecipesChild  []dgraphRecipe     `json:"~related_recipes,omitempty"`
	Notes                []dgraphNote       `json:"~recipe,omitempty"`

	DType []string `json:"dgraph.type,omitempty"`
}

// Due to how dgraph returns facet data, we have to do more than just copy
// between two structs. A facet is data associated with a particular edge. For
// example, that's where ingredient amounts are stored. The dgraph query for a
// recipe will return this information in the form of:
// {
//	  "name": "My Recipe",
//    "ingredients": [
//        {"uid": "0xf00", "name": "Black Beans"}
//    ],
//    "ingredients|amount" : {
//	    {"0": "1 cup, presoaked"}
//    }
// }
// In order to restructure this in a sane manner, we move the "ingredients|amount"
// field values over into each element of "ingredient". Also need to convert str
// values into ints when appropriate, as `json.Unmarshal()` refuses to cast for
// you.
//
// Next, as dgraph returns all edges as a list, we must take special care
// to handle edges that are meant to only exist once, e.g. the User-type node which
// created a given recipe. Specifically, the `copier.Copy()` function does not play
// nicely with two fields of the same name when only one is an array.
//
// Finally, we need to aggregate both direct (child) and reverse (parent) edges into
// a single Recipe.RelatedRecipes struct field.
func (dRecipe *dgraphRecipe) dgraphToModel(recipe *models.Recipe) error {
	copier.Copy(&recipe, &dRecipe)

	// list to single value
	recipeCreator := models.User{}
	copier.Copy(&recipeCreator, dRecipe.Owner[0])
	recipe.CreatedBy = recipeCreator

	// merge facet data
	for s_idx, value := range dRecipe.IngredientAmounts {
		i_idx, err := strconv.Atoi(s_idx)

		if err != nil {
			return err
		}

		recipe.Ingredients[i_idx].Amount = value
	}

	for s_idx, s_value := range dRecipe.RatingScore {
		i_idx, err := strconv.Atoi(s_idx)
		if err != nil {
			return err
		}

		i_value, err := strconv.Atoi(s_value)
		if err != nil {
			return err
		}

		recipe.RatedBy[i_idx].RatingScore = i_value
	}

	// merge matching edges and reverse edges
	relatedRecipes := []models.Recipe{}
	dRelatedRecipes := append(dRecipe.RelatedRecipesParent, dRecipe.RelatedRecipesChild...)
	copier.Copy(&relatedRecipes, &dRelatedRecipes)
	recipe.RelatedRecipes = relatedRecipes

	return nil
}

// GetAllRecipes out of dgraph
func (d *dgraphRepo) GetAllRecipes() (*models.ManyRecipes, error) {
	recipes := models.ManyRecipes{}
	dRecipes := manyDgraphRecipes{}
	ctx := d.buildAuthContext(context.Background())
	txn := d.Client.NewReadOnlyTxn()
	defer txn.Discard(ctx)

	const q = `
		{
			recipes(func: type(Recipe)) {
				uid
				name

				dgraph.type

				recipe_tags {
					uid
					name
				}
			}
		}
	`

	resp, err := txn.Query(ctx, q)
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

// GetRecipe out of dgraph by ID
func (d *dgraphRepo) GetRecipe(id string) (*models.Recipe, error) {
	recipe := models.Recipe{}
	dRecipes := manyDgraphRecipes{}
	ctx := d.buildAuthContext(context.Background())
	txn := d.Client.NewReadOnlyTxn()
	defer txn.Discard(ctx)

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
				recipe_tags {
					uid
					name
				}
				owner {
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
				~related_recipes {
					uid
					name
				}
				related_recipes {
					uid
					name
				}
			}
		}
	`

	resp, err := txn.QueryWithVars(ctx, q, variables)
	if err != nil {
		return &recipe, err
	}

	err = json.Unmarshal(resp.Json, &dRecipes)
	if err != nil {
		return &recipe, err
	}

	if len(dRecipes.Recipes) > 0 {
		err := dRecipes.Recipes[0].dgraphToModel(&recipe)
		if err != nil {
			return &recipe, err
		}
	}

	return &recipe, nil
}

// CreateRecipe within dgraph
func (d *dgraphRepo) CreateRecipe(recipe *models.Recipe) (*models.Recipe, error) {
	dRecipe := dgraphRecipe{}
	ctx := d.buildAuthContext(context.Background())
	txn := d.Client.NewTxn()
	defer txn.Discard(ctx)

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

	res, err := txn.Mutate(ctx, mu)
	if err != nil {
		return recipe, err
	}

	recipe.ID = res.Uids["recipe"]

	return recipe, nil
}

// UpdateRecipe within dgraph
func (d *dgraphRepo) UpdateRecipe(recipe *models.Recipe) (*models.Recipe, error) {
	dRecipe := dgraphRecipe{}
	ctx := d.buildAuthContext(context.Background())
	txn := d.Client.NewTxn()
	defer txn.Discard(ctx)

	copier.Copy(&dRecipe, recipe)
	dRecipe.DType = []string{"Recipe"}

	mu := &api.Mutation{
		CommitNow: true,
	}
	dgo.DeleteEdges(mu, dRecipe.ID, "~related_recipes")

	pb, err := json.Marshal(dRecipe)
	if err != nil {
		return recipe, err
	}

	mu = &api.Mutation{
		CommitNow: true,
		SetJson:   pb,
	}

	_, err = txn.Mutate(ctx, mu)
	if err != nil {
		return recipe, err
	}

	return recipe, nil
}

// DeleteRecipe from dgraph
func (d *dgraphRepo) DeleteRecipe(id string) error {
	dRecipes := manyDgraphRecipes{}
	ctx := d.buildAuthContext(context.Background())
	txn := d.Client.NewReadOnlyTxn()
	defer txn.Discard(ctx)

	// First we need to grab all reverse edges as they must deleted by referencing
	// the parent node itself, and not the current child recipe being deleted. This
	// is a good sign that the schema isn't well designed, but that will be addressed
	// later on. This works Good Enough for now.
	variables := map[string]string{"$id": id}
	const q = `
		query all($id: string) {
			recipes(func: uid($id)) @filter(type(Recipe)) {
				uid
				~related_recipes {
					uid
				}
				~ratings {
					uid
				}
				~favorites {
					uid
				}
				~recipe {
					uid
				}
			}
		}
	`

	resp, err := txn.QueryWithVars(ctx, q, variables)
	if err != nil {
		return err
	}

	err = json.Unmarshal(resp.Json, &dRecipes)
	if err != nil {
		return err
	}

	// Doesn't exist, just return now
	if len(dRecipes.Recipes) == 0 {
		return nil
	}

	txn = d.Client.NewTxn()
	defer txn.Discard(ctx)

	// Now lets delete all our reverse edges by referencing the parent node as the subject
	for _, dRecipe := range dRecipes.Recipes[0].RelatedRecipesParent {
		mu := &api.Mutation{
			Del: []*api.NQuad{
				{
					Subject:   dRecipe.ID,
					Predicate: "related_recipes",
					ObjectId:  id,
				},
			},
		}

		_, err = txn.Mutate(ctx, mu)
		if err != nil {
			return err
		}
	}

	for _, dUser := range dRecipes.Recipes[0].RatedBy {
		mu := &api.Mutation{
			Del: []*api.NQuad{
				{
					Subject:   dUser.ID,
					Predicate: "ratings",
					ObjectId:  id,
				},
			},
		}

		_, err = txn.Mutate(ctx, mu)
		if err != nil {
			return err
		}
	}

	for _, dUser := range dRecipes.Recipes[0].FavoritedBy {
		mu := &api.Mutation{
			Del: []*api.NQuad{
				{
					Subject:   dUser.ID,
					Predicate: "favorites",
					ObjectId:  id,
				},
			},
		}

		_, err = txn.Mutate(ctx, mu)
		if err != nil {
			return err
		}
	}

	for _, dNote := range dRecipes.Recipes[0].Notes {
		mu := &api.Mutation{
			Del: []*api.NQuad{
				{
					Subject:   dNote.ID,
					Predicate: "recipes",
					ObjectId:  id,
				},
			},
		}

		_, err = txn.Mutate(ctx, mu)
		if err != nil {
			return err
		}
	}

	// Now lets delete the node itself
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

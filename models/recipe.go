package models

// APIRecipe is a struct that represents a single recipe. It is used exclusively
// for marshalling responses back to API clients.
type APIRecipe struct {
	ID            string `json:"uid,omitempty"`
	Name          string `json:"name,omitempty"`
	URL           string `json:"url,omitempty"`
	Domain        string `json:"domain,omitempty"`
	Directions    string `json:"directions,omitempty"`
	PrepTime      int    `json:"prep_time,omitempty"`
	CookTime      int    `json:"cook_time,omitempty"`
	TotalServings int    `json:"total_servings,omitempty"`
	HasBeenTried  bool   `json:"has_been_tried,omitempty"`

	Ingredients       []NestedIngredient `json:"ingredients,omitempty"`
	IngredientAmounts []string           `json:"ingredientAmounts,omitempty"`
	Tags              []NestedTag        `json:"categories,omitempty"`
	RatedBy           []NestedUser       `json:"rated_by,omitempty"`
	RatingScore       []int              `json:"rating_score,omitempty"`
	FavoritedBy       []NestedUser       `json:"favorited_by,omitempty"`
	RelatedRecipes    []NestedRecipe     `json:"related_recipes,omitempty"`
	Notes             []NestedNote       `json:"notes,omitempty"`
}

// NestedRecipe is a stripped down struct used when a Recipe is nested
// within a parent struct in an API response
type NestedRecipe struct {
	ID   string `json:"uid,omitempty"`
	Name string `json:"name,omitempty" validate:"required"`
}

// ManyAPIRecipes is a struct that represents multiple recipes. It is used
// exclusively for marshalling responsesback to API clients.
type ManyAPIRecipes struct {
	Recipes []APIRecipe `json:"recipes"`
}

// Recipe is a struct that represents a single recipe.
type Recipe struct {
	ID            string
	Name          string
	URL           string
	Domain        string
	Directions    string
	PrepTime      int
	CookTime      int
	TotalServings int
	HasBeenTried  bool

	Ingredients    []Ingredient
	Tags           []Tag
	RatedBy        []User
	FavoritedBy    []User
	RelatedRecipes []Recipe
	Notes          []Note
}

// ManyRecipes is a struct that represents multiple recipes
type ManyRecipes struct {
	Recipes []Recipe
}

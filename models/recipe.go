package models

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

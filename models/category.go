package models

// APICategory is a struct that represents a single category. It is used exclusively
// for interaction with clients.
type APICategory struct {
	ID   string `json:"uid,omitempty"`
	Name string `json:"name,omitempty" validate:"required"`

	Recipes     []NestedRecipe     `json:"recipes,omitempty"`
	Ingredients []NestedIngredient `json:"ingredients,omitempty"`
}

// NestedCategory is a stripped down struct used when a Category is nested
// within a parent struct in an API response
type NestedCategory struct {
	ID   string `json:"uid,omitempty"`
	Name string `json:"name,omitempty" validate:"required"`
}

// ManyAPICategories is a struct that represents multiple categories. It is used
// exclusively for interaction with clients.
type ManyAPICategories struct {
	Categories []APICategory `json:"categories"`
}

// Category is a struct that represents a single category
type Category struct {
	ID   string
	Name string

	Recipes     []Recipe
	Ingredients []Ingredient
}

// ManyCategories is a struct that represents multiple categories
type ManyCategories struct {
	Categories []Category
}

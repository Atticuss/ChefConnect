package models

// APIIngredient is a struct that represents a single ingredient. It is used exclusively
// for interaction with clients.
type APIIngredient struct {
	ID     string `json:"uid,omitempty"`
	Name   string `json:"name,omitempty" validate:"required"`
	Amount string `json:"amount",omitempty`

	IngredientCategories []NestedCategory `json:"ingredientCategories,omitempty"`
	RecipeCategories     []NestedCategory `json:"recipeCategories,omitempty"`
}

// NestedIngredient is a stripped down struct used when an Ingredient is nested
// within a parent struct in an API response
type NestedIngredient struct {
	ID     string `json:"uid,omitempty"`
	Name   string `json:"name,omitempty" validate:"required"`
	Amount string `json:"amount",omitempty`
}

// ManyAPIIngredients is a struct that represents multiple ingredients. It is used exclusively
// for interaction with clients.
type ManyAPIIngredients struct {
	Ingredients []APIIngredient `json:"ingredients"`
}

// Ingredient is a struct that represents a single ingredient
type Ingredient struct {
	ID     string `json:"uid,omitempty"`
	Name   string `json:"name,omitempty" validate:"required"`
	Amount string `json:"ingredients|amount",omitempty`

	IngredientCategories []Category `json:"ingredientCategories,omitempty"`
	RecipeCategories     []Category `json:"recipeCategories,omitempty"`
}

// ManyIngredients is a struct that represents multiple ingredients
type ManyIngredients struct {
	Ingredients []Ingredient `json:"ingredients"`
}

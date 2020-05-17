package models

// APIIngredient is a struct that represents a single ingredient. It is used exclusively
// for interaction with clients.
type APIIngredient struct {
	ID     string `json:"uid,omitempty"`
	Name   string `json:"name,omitempty" validate:"required"`
	Amount string `json:"amount,omitempty"`

	Categories []NestedCategory `json:"categories,omitempty"`
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
	ID     string
	Name   string
	Amount string

	Categories []Category
}

// ManyIngredients is a struct that represents multiple ingredients
type ManyIngredients struct {
	Ingredients []Ingredient
}

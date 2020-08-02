package models

// Ingredient is a struct that represents a single ingredient
type Ingredient struct {
	ID     string
	Name   string
	Amount string

	Tags []Tag
}

// ManyIngredients is a struct that represents multiple ingredients
type ManyIngredients struct {
	Ingredients []Ingredient
}

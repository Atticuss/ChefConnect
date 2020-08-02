package models

// Tag is a struct that represents a single tag
type Tag struct {
	ID   string
	Name string

	Recipes     []Recipe
	Ingredients []Ingredient
}

// ManyTags is a struct that represents multiple tags
type ManyTags struct {
	Tags []Tag
}

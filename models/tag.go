package models

// APITag is a struct that represents a single tag. It is used exclusively
// for interaction with clients.
type APITag struct {
	ID   string `json:"uid,omitempty"`
	Name string `json:"name,omitempty" validate:"required"`

	Recipes     []NestedRecipe     `json:"recipes,omitempty"`
	Ingredients []NestedIngredient `json:"ingredients,omitempty"`
}

// NestedTag is a stripped down struct used when a Tag is nested
// within a parent struct in an API response
type NestedTag struct {
	ID   string `json:"uid,omitempty"`
	Name string `json:"name,omitempty" validate:"required"`
}

// ManyAPITags is a struct that represents multiple tags. It is used
// exclusively for interaction with clients.
type ManyAPITags struct {
	Tags []APITag `json:"tags"`
}

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

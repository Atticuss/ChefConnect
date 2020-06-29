package models

// APIUser is a struct that represents a single user. It is used exclusively
// for marshalling responses back to API clients.
type APIUser struct {
	ID       string `json:"uid,omitempty"`
	Name     string `json:"name,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`

	Favorites []NestedRecipe `json:"favorites,omitempty"`
	Notes     []NestedNote   `json:"notes,omitempty"`
	Ratings   []NestedRecipe `json:"ratings,omitempty"`
	Roles     []NestedRole   `json:"roles,omitempty"`
}

// NestedUser is a stripped down struct used when a User is nested
// within a parent struct within an API response
type NestedUser struct {
	ID          string `json:"uid,omitempty"`
	Name        string `json:"name,omitempty"`
	Username    string `json:"username,omitempty"`
	RatingScore int    `json:"ratingScore,omitempty"`
}

// ManyAPIUsers is a struct that represents multiple users. It is used
// exclusively for marshalling responsesback to API clients.
type ManyAPIUsers struct {
	Users []APIUser `json:"users"`
}

// JwtUser is a struct holding the user data that should be present within an
// issued JWT's claims
type JwtUser struct {
	ID       string `json:"uid"`
	Name     string `json:"name"`
	Username string `json:"username"`

	Roles []NestedRole `json:"roles,omitempty"`
}

// User is a struct that represents a single user.
type User struct {
	ID          string
	Name        string
	Username    string
	Password    string
	RatingScore int

	Favorites []Recipe
	Notes     []Note
	Ratings   []Recipe
	Roles     []Role
}

// ManyUsers is a struct that represents multiple users
type ManyUsers struct {
	Users []User
}

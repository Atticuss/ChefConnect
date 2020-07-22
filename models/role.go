package models

// APIRole is a struct that represents a single role. It is used exclusively
// for interaction with clients.
type APIRole struct {
	ID   string `json:"uid,omitempty"`
	Name string `json:"name,omitempty"`

	Users []NestedUser `json:"users,omitempty"`
}

// NestedRole is a stripped down struct used when a Role is nested
// within a parent struct within an API response
type NestedRole struct {
	ID   string `json:"uid,omitempty"`
	Name string `json:"name,omitempty"`
}

// ManyAPIRoles is a struct that represents multiple roles. It is used
// exclusively for interaction with clients.
type ManyAPIRoles struct {
	Roles []APIRole `json:"roles"`
}

// Role is a struct that represents a single role.
type Role struct {
	ID   string
	Name string

	Users []User
}

// ManyRoles is a struct that represents multiple roles
type ManyRoles struct {
	Roles []Role
}

package models

// NestedRole is a stripped down struct used when a Role is nested
// within a parent struct within an API response
type NestedRole struct {
	ID   string `json:"uid,omitempty"`
	Name string `json:"name,omitempty"`
}

// Role is a struct that represents a single role.
type Role struct {
	ID   string
	Name string

	Users []User
}

// ManyRoles is a struct that represents multiple users
type ManyRoles struct {
	Roles []Role
}

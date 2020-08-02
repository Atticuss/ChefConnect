package models

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

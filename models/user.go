package models

// User is a struct that represents a single user.
type User struct {
	ID          string
	Name        string
	Username    string
	Password    string
	RatingScore int

	AuthToken            string
	RefreshToken         string
	RefreshTokenIssuedAt int64

	Favorites []Recipe
	Notes     []Note
	Ratings   []Recipe
	Roles     []Role
}

// ManyUsers is a struct that represents multiple users
type ManyUsers struct {
	Users []User
}

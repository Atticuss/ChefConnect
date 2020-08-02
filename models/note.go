package models

// Note is a struct that represents a single note
type Note struct {
	ID   string
	Text string

	User   []User
	Recipe []Recipe
}

// ManyNotes is a struct that represents multiple notes
type ManyNotes struct {
	Notes []Note
}

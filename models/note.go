package models

// APINote is a struct that represents a single note. It is used exclusively
// for marshalling responses back to API clients.
type APINote struct {
	ID   string `json:"uid,omitempty"`
	Text string `json:"text,omitempty"`

	User   []NestedUser   `json:"author,omitempty"`
	Recipe []NestedRecipe `json:"recipe,omitempty"`
}

// NestedNote is a stripped down struct used when a Note is nested
// within a parent struct in an API response
type NestedNote struct {
	ID   string `json:"uid,omitempty"`
	Text string `json:"text,omitempty"`
}

// ManyAPINotes is a struct that represents multiple notes. It is used
// exclusively for marshalling responsesback to API clients.
type ManyAPINotes struct {
	Notes []APINote `json:"notes"`
}

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

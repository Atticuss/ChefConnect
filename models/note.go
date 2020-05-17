package models

// NoteResponse is a struct that represents a single note. It is used exclusively
// for marshalling responses back to API clients.
type NoteResponse struct {
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

// ManyNotesResponse is a struct that represents multiple notes. It is used
// exclusively for marshalling responsesback to API clients.
type ManyNotesResponse struct {
	Notes []NoteResponse `json:"notes"`
}

// Note is a struct that represents a single note
type Note struct {
	ID   string `json:"uid,omitempty"`
	Text string `json:"text,omitempty"`

	User   []User   `json:"author,omitempty"`
	Recipe []Recipe `json:"recipe,omitempty"`

	DType []string `json:"dgraph.type,omitempty"`
}

// ManyNotes is a struct that represents multiple notes
type ManyNotes struct {
	Notes []Note `json:"notes"`
}

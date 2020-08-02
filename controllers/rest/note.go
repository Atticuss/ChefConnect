package rest

type restNote struct {
	ID   string `json:"uid,omitempty"`
	Text string `json:"text,omitempty"`

	User   []nestedUser   `json:"author,omitempty"`
	Recipe []nestedRecipe `json:"recipe,omitempty"`
}

type nestedNote struct {
	ID   string `json:"uid,omitempty"`
	Text string `json:"text,omitempty"`
}

// ManyAPINotes is a struct that represents multiple notes. It is used
// exclusively for marshalling responsesback to API clients.
type manyRestNotes struct {
	Notes []nestedNote `json:"notes"`
}

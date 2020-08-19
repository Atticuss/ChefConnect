package dgraph

type manyDgraphNotes struct {
	Notes []dgraphNote `json:"notes"`
}

type dgraphNote struct {
	ID   string `json:"uid,omitempty"`
	Name string `json:"name,omitempty" validate:"required"`

	User   []dgraphUser   `json:"author,omitempty"`
	Recipe []dgraphRecipe `json:"recipe,omitempty"`

	DType []string `json:"dgraph.type,omitempty"`
}

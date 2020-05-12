package models

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/dgraph-io/dgo/v2"
	"github.com/dgraph-io/dgo/v2/protos/api"
)

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
	Recipes []Note `json:"notes"`
}

// parent struct for dgraph responses
type rootNote struct {
	Note []Note `json:"root"`
}

// GetNote will fetch a note via a given ID
func (n *Note) GetNote(c *dgo.Dgraph) error {
	return errors.New("Not implemented")
}

// UpdateNote will update the text of a Note via a given by ID
func (n *Note) UpdateNote(c *dgo.Dgraph) error {
	return errors.New("Not implemented")
}

// DeleteNote will delete a note via a given by ID
func (n *Note) DeleteNote(c *dgo.Dgraph) error {
	return errors.New("Not implemented")
}

// CreateNote will create a new note from the given Note struct
func (n *Note) CreateNote(c *dgo.Dgraph) error {
	fmt.Println("CreateNote() start")

	txn := c.NewTxn()
	defer txn.Discard(context.Background())

	// assign an alias ID that can be ref'd out of the response's uid []string map
	n.ID = "_:note"
	n.DType = []string{"Note"}

	pb, err := json.Marshal(n)
	if err != nil {
		return err
	}

	mu := &api.Mutation{
		CommitNow: true,
	}
	mu.SetJson = pb
	res, err := txn.Mutate(context.Background(), mu)
	if err != nil {
		return err
	}

	fmt.Println("CreateNote mutation resp: ")
	fmt.Printf("%+v\n", res)

	n.ID = res.Uids["note"]

	return nil
}

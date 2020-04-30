package models

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/dgraph-io/dgo/v2"
	"github.com/dgraph-io/dgo/v2/protos/api"
)

// User is a struct that represents a single user
type User struct {
	ID        string   `json:"uid,omitempty"`
	Name      string   `json:"name,omitempty"`
	Username  string   `json:"username,omitempty"`
	Password  string   `json:"password,omitempty"`
	Favorites []Recipe `json:"favorites,omitempty"`
	Notes     []Note   `json:"notes,omitempty"`
	Ratings   []Recipe `json:"ratings,omitempty"`
	DType     []string `json:"dgraph.type,omitempty"`
}

// ManyUsers is a struct that represents multiple users
type ManyUsers struct {
	Recipes []Note
}

// cheap hack to get around how dgraph returns data
type singleUser struct {
	Note []Note
}

// GetUser will fetch a user via a given ID
func (u *User) GetUser(c *dgo.Dgraph) error {
	return errors.New("Not implemented")
}

// UpdateUser will update a user via a given by ID
func (u *User) UpdateUser(c *dgo.Dgraph) error {
	return errors.New("Not implemented")
}

// Deleteuser will delete a user via a given by ID
func (u *User) Deleteuser(c *dgo.Dgraph) error {
	return errors.New("Not implemented")
}

// CreateUser will create a new user from the given User struct
func (u *User) CreateUser(c *dgo.Dgraph) error {
	fmt.Println("CreateUser() start")

	txn := c.NewTxn()
	defer txn.Discard(context.Background())

	// assign an alias ID that can be ref'd out of the response's uid []string map
	u.ID = "_:user"
	u.DType = []string{"User"}

	pb, err := json.Marshal(u)
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

	fmt.Println("CreateUser mutation resp: ")
	fmt.Printf("%+v\n", res)

	u.ID = res.Uids["note"]

	return nil
}

package models

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/dgraph-io/dgo/v2"
	"github.com/dgraph-io/dgo/v2/protos/api"
)

// UserResponse is a struct that represents a single user. It is used exclusively
// for marshalling responses back to API clients.
type UserResponse struct {
	ID       string `json:"uid,omitempty"`
	Name     string `json:"name,omitempty"`
	Username string `json:"username,omitempty"`

	Favorites []NestedRecipe `json:"favorites,omitempty"`
	Notes     []NestedNote   `json:"notes,omitempty"`
	Ratings   []NestedRecipe `json:"ratings,omitempty"`
}

// NestedUser is a stripped down struct used when a User is nested
// within a parent struct within an API response
type NestedUser struct {
	ID       string `json:"uid,omitempty"`
	Name     string `json:"name,omitempty"`
	Username string `json:"username,omitempty"`
}

// ManyUsersResponse is a struct that represents multiple users. It is used
// exclusively for marshalling responsesback to API clients.
type ManyUsersResponse struct {
	Users []UserResponse `json:"users"`
}

// User is a struct that represents a single user. It is used exclusively
// for unmarshalling responses from dgraph
type User struct {
	ID       string `json:"uid,omitempty"`
	Name     string `json:"name,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`

	Favorites []Recipe `json:"favorites,omitempty"`
	Notes     []Note   `json:"~author,omitempty"`
	Ratings   []Recipe `json:"ratings,omitempty"`

	DType []string `json:"dgraph.type,omitempty"`
}

// ManyUsers is a struct that represents multiple users
type ManyUsers struct {
	Users []User `json:"users"`
}

// parent struct for dgraph responses
type rootUser struct {
	Users []User `json:"root"`
}

//GetAllUsers will fetch all users
func GetAllUsers(c *dgo.Dgraph) (*ManyUsers, error) {
	txn := c.NewReadOnlyTxn()

	const q = `
		{
			users(func: type(User)) {
				uid
				name
				password
				dgraph.type

				favorites {
					uid
					name
				}
			}
		}
	`

	resp, err := txn.Query(context.Background(), q)
	if err != nil {
		return nil, err
	}

	users := ManyUsers{}
	err = json.Unmarshal(resp.Json, &users)
	if err != nil {
		return nil, err
	}

	return &users, nil
}

// GetUser will fetch a user via a given ID
func (user *User) GetUser(c *dgo.Dgraph) error {
	txn := c.NewReadOnlyTxn()

	variables := map[string]string{"$id": user.ID}
	const q = `
		query all($id: string) {
			users(func: uid($id)) @filter(type(User)) {
				uid
				name
				password
				dgraph.type

				favorites {
					uid
					name
				}
			}
		}
	`

	resp, err := txn.QueryWithVars(context.Background(), q, variables)
	if err != nil {
		return err
	}

	users := ManyUsers{}
	err = json.Unmarshal(resp.Json, &users)
	if err != nil {
		return err
	}

	*user = users.Users[0]

	return nil
}

// UpdateUser will update a user via a given by ID
func (user *User) UpdateUser(c *dgo.Dgraph) error {
	return errors.New("Not implemented")
}

// DeleteUser will delete a user via a given by ID
func (user *User) DeleteUser(c *dgo.Dgraph) error {
	return errors.New("Not implemented")
}

// CreateUser will create a new user from the given User struct
func (user *User) CreateUser(c *dgo.Dgraph) error {
	fmt.Println("CreateUser() start")

	txn := c.NewTxn()
	defer txn.Discard(context.Background())

	// assign an alias ID that can be ref'd out of the response's uid []string map
	user.ID = "_:user"
	user.DType = []string{"User"}

	pb, err := json.Marshal(user)
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

	user.ID = res.Uids["user"]

	return nil
}

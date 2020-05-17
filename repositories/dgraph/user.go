package dgraph

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/dgraph-io/dgo/v2"
	"github.com/dgraph-io/dgo/v2/protos/api"
	"github.com/jinzhu/copier"

	"github.com/atticuss/chefconnect/models"
	"github.com/atticuss/chefconnect/repositories"
)

type dgraphUserRepo struct {
	Client *dgo.Dgraph
}

// NewDgraphUserRepository configures a dgraph repository for accessing
// user data
func NewDgraphUserRepository(db *dgo.Dgraph) repositories.UserRepository {
	return &dgraphUserRepo{
		Client: db,
	}
}

type manyDgraphUsers struct {
	Users []dgraphUser `json:"users"`
}

type dgraphUser struct {
	ID       string `json:"uid,omitempty"`
	Name     string `json:"name,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`

	Favorites []dgraphRecipe `json:"favorites,omitempty"`
	Notes     []models.Note  `json:"~author,omitempty"`
	Ratings   []dgraphRecipe `json:"ratings,omitempty"`

	DType []string `json:"dgraph.type,omitempty"`
}

// GetAll users out of dgraph
func (d *dgraphUserRepo) GetAll() (*models.ManyUsers, error) {
	users := models.ManyUsers{}
	dUsers := manyDgraphUsers{}
	txn := d.Client.NewReadOnlyTxn()
	defer txn.Discard(context.Background())

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
		return &users, err
	}

	err = json.Unmarshal(resp.Json, &dUsers)
	if err != nil {
		return &users, err
	}

	copier.Copy(&users, &dUsers)
	return &users, nil
}

// Get a user out of dgraph by ID
func (d *dgraphUserRepo) Get(id string) (*models.User, error) {
	user := models.User{}
	dUsers := manyDgraphUsers{}
	txn := d.Client.NewReadOnlyTxn()
	defer txn.Discard(context.Background())

	variables := map[string]string{"$id": id}
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
		return &user, err
	}

	err = json.Unmarshal(resp.Json, &dUsers)
	if err != nil {
		return &user, err
	}

	if len(dUsers.Users) > 0 {
		copier.Copy(&user, &dUsers.Users[0])
		return &user, nil
	}

	return &user, nil
}

// Create a user within dgraph
func (d *dgraphUserRepo) Create(user *models.User) (*models.User, error) {
	dUser := dgraphUser{}
	txn := d.Client.NewTxn()
	defer txn.Discard(context.Background())

	copier.Copy(&dUser, user)
	dUser.ID = "_:user"
	dUser.DType = []string{"User"}

	pb, err := json.Marshal(dUser)
	if err != nil {
		return user, err
	}

	mu := &api.Mutation{
		CommitNow: true,
		SetJson:   pb,
	}

	res, err := txn.Mutate(context.Background(), mu)
	if err != nil {
		return user, err
	}

	user.ID = res.Uids["user"]

	return user, nil
}

// Update a user within dgraph
func (d *dgraphUserRepo) Update(user *models.User) (*models.User, error) {
	dUser := dgraphUser{}
	txn := d.Client.NewTxn()
	defer txn.Discard(context.Background())

	copier.Copy(&dUser, user)
	dUser.DType = []string{"User"}

	pb, err := json.Marshal(dUser)
	if err != nil {
		return user, err
	}

	mu := &api.Mutation{
		CommitNow: true,
		SetJson:   pb,
	}

	_, err = txn.Mutate(context.Background(), mu)
	if err != nil {
		return user, err
	}

	return user, nil
}

// Delete a user from dgraph
func (d *dgraphUserRepo) Delete(id string) error {
	return errors.New("Not implemented")
}

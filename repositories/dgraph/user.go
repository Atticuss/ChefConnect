package dgraph

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/dgraph-io/dgo/v2/protos/api"
	"github.com/jinzhu/copier"

	"github.com/atticuss/chefconnect/models"
)

type manyDgraphUsers struct {
	Users []dgraphUser `json:"users"`
}

type dgraphUser struct {
	ID       string `json:"uid,omitempty"`
	Name     string `json:"name,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`

	Favorites []dgraphRecipe `json:"favorites,omitempty"`
	Notes     []dgraphNote   `json:"~author,omitempty"`
	Ratings   []dgraphRecipe `json:"ratings,omitempty"`
	Roles     []dgraphRole   `json:"roles,omitempty"`

	DType []string `json:"dgraph.type,omitempty"`
}

// GetAll users out of dgraph
func (d *dgraphRepo) GetAllUsers() (*models.ManyUsers, error) {
	users := models.ManyUsers{}
	dUsers := manyDgraphUsers{}
	ctx := d.buildAuthContext(context.Background())
	txn := d.Client.NewReadOnlyTxn()
	defer txn.Discard(context.Background())

	const q = `
		{
			users(func: type(User)) {
				uid
				name
				username
				password
				dgraph.type

				favorites {
					uid
					name
				}

				roles {
					uid
					name
				}
			}
		}
	`

	resp, err := txn.Query(ctx, q)
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
func (d *dgraphRepo) GetUser(id string) (*models.User, error) {
	user := models.User{}
	dUsers := manyDgraphUsers{}
	ctx := d.buildAuthContext(context.Background())
	txn := d.Client.NewReadOnlyTxn()
	defer txn.Discard(context.Background())

	variables := map[string]string{"$id": id}
	const q = `
		query all($id: string) {
			users(func: uid($id)) @filter(type(User)) {
				uid
				name
				username
				password
				dgraph.type

				favorites {
					uid
					name
				}

				roles {
					uid
					name
				}
			}
		}
	`

	resp, err := txn.QueryWithVars(ctx, q, variables)
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

// Get a user out of dgraph by name
func (d *dgraphRepo) GetUserByUsername(username string) (*models.User, error) {
	fmt.Println("in GetUserByUsername()")
	user := models.User{}
	dUsers := manyDgraphUsers{}
	ctx := d.buildAuthContext(context.Background())
	txn := d.Client.NewReadOnlyTxn()
	defer txn.Discard(context.Background())

	variables := map[string]string{"$username": username}
	const q = `
		query all($username: string) {
			users(func: eq(username, $username)) @filter(type(User)) {
				uid
				name
				username
				password
				dgraph.type

				favorites {
					uid
					name
				}

				roles {
					uid
					name
				}
			}
		}
	`

	fmt.Println("starting txn")

	resp, err := txn.QueryWithVars(ctx, q, variables)
	if err != nil {
		fmt.Printf("error during txn: %+v\n", err)
		return &user, err
	}

	fmt.Println("json unmarshalling")

	err = json.Unmarshal(resp.Json, &dUsers)
	if err != nil {
		fmt.Printf("error when json unmarshalling: %+v\n", err)
		return &user, err
	}

	fmt.Printf("dusers: %+v\n", dUsers)

	if len(dUsers.Users) > 0 {
		fmt.Println("user not found")
		copier.Copy(&user, &dUsers.Users[0])
		return &user, nil
	}

	fmt.Println("returning user details")

	return &user, nil
}

// Create a user within dgraph
func (d *dgraphRepo) CreateUser(user *models.User) (*models.User, error) {
	dUser := dgraphUser{}
	ctx := d.buildAuthContext(context.Background())
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

	res, err := txn.Mutate(ctx, mu)
	if err != nil {
		return user, err
	}

	user.ID = res.Uids["user"]

	return user, nil
}

// Update a user within dgraph
func (d *dgraphRepo) UpdateUser(user *models.User) (*models.User, error) {
	dUser := dgraphUser{}
	ctx := d.buildAuthContext(context.Background())
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

	_, err = txn.Mutate(ctx, mu)
	if err != nil {
		return user, err
	}

	return user, nil
}

// Delete a user from dgraph
func (d *dgraphRepo) DeleteUser(id string) error {
	ctx := d.buildAuthContext(context.Background())
	txn := d.Client.NewTxn()
	defer txn.Discard(context.Background())

	readOnlyTxn := d.Client.NewReadOnlyTxn()
	defer readOnlyTxn.Discard(context.Background())

	dUsers := manyDgraphUsers{}
	variables := map[string]string{"$id": id}
	const q = `
		query all($id: string) {
			users(func: uid($id)) @filter(type(User)) {
				uid
				~author {
					uid
				}
			}
		}
	`

	resp, err := readOnlyTxn.QueryWithVars(ctx, q, variables)
	if err != nil {
		return err
	}

	err = json.Unmarshal(resp.Json, &dUsers)
	if err != nil {
		return err
	}

	// Doesn't exist, just return now
	if len(dUsers.Users) == 0 {
		return nil
	}

	// Once the note repo is implemented, we can just call that for deletion. For
	// now, we'll rely on this.
	for _, dNote := range dUsers.Users[0].Notes {
		mu := &api.Mutation{
			Del: []*api.NQuad{
				{
					Subject:   dNote.ID,
					Predicate: "author",
					ObjectId:  id,
				},
			},
		}

		_, err = txn.Mutate(ctx, mu)
		if err != nil {
			return err
		}
	}

	// Now lets delete the node itself
	variables = map[string]string{"uid": id}
	pb, err := json.Marshal(variables)
	if err != nil {
		return err
	}

	mu := &api.Mutation{
		CommitNow:  true,
		DeleteJson: pb,
	}

	_, err = txn.Mutate(ctx, mu)
	if err != nil {
		return err
	}

	return nil
}

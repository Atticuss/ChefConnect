package dgraph

import (
	"context"
	"encoding/json"

	"github.com/dgraph-io/dgo/v2"
	"github.com/dgraph-io/dgo/v2/protos/api"
	"github.com/jinzhu/copier"
	"google.golang.org/grpc"

	"github.com/atticuss/chefconnect/models"
	"github.com/atticuss/chefconnect/repositories"
)

type dgraphRoleRepo struct {
	Client *dgo.Dgraph
}

// NewDgraphTagRepository configures a dgraph repository for accessing
// tag data
func NewDgraphRoleRepository(config *Config) repositories.RoleRepository {
	conn, _ := grpc.Dial(config.Host, grpc.WithInsecure())
	client := dgo.NewDgraphClient(api.NewDgraphClient(conn))

	return &dgraphRoleRepo{
		Client: client,
	}
}

type manyDgraphRoles struct {
	Roles []dgraphRole `json:"roles"`
}

type dgraphRole struct {
	ID   string `json:"uid,omitempty"`
	Name string `json:"name,omitempty"`

	Users []dgraphUser `json:"~roles,omitempty"`

	DType []string `json:"dgraph.type,omitempty"`
}

// GetAll roles out of dgraph
func (d *dgraphRoleRepo) GetAll() (*models.ManyRoles, error) {
	dRoles := manyDgraphRoles{}
	roles := models.ManyRoles{}
	txn := d.Client.NewReadOnlyTxn()
	defer txn.Discard(context.Background())

	const q = `
		{
			roles(func: type(Role)) {
				uid
				name
				dgraph.type
			}
		}
	`

	resp, err := txn.Query(context.Background(), q)
	if err != nil {
		return &roles, err
	}

	err = json.Unmarshal(resp.Json, &dRoles)
	if err != nil {
		return &roles, err
	}

	copier.Copy(&roles, &dRoles)

	return &roles, nil
}

// Get a role out of dgraph by ID
func (d *dgraphRoleRepo) Get(id string) (*models.Role, error) {
	dRoles := manyDgraphRoles{}
	role := models.Role{}
	txn := d.Client.NewReadOnlyTxn()
	defer txn.Discard(context.Background())

	variables := map[string]string{"$id": id}
	const q = `
		query all($id: string) {
			roles(func: uid($id)) @filter(type(Role)) {
				uid
				name
				dgraph.type

				~roles {
					uid
					name
					username
					dgraph.type
				}
			}
		}
	`

	resp, err := txn.QueryWithVars(context.Background(), q, variables)
	if err != nil {
		return &role, err
	}

	err = json.Unmarshal(resp.Json, &dRoles)
	if err != nil {
		return &role, err
	}

	if len(dRoles.Roles) > 0 {
		copier.Copy(&role, &dRoles.Roles[0])
	}

	return &role, nil
}

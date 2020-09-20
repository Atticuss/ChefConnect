package dgraph

import (
	"context"
	"encoding/json"

	"github.com/jinzhu/copier"

	"github.com/atticuss/chefconnect/models"
)

type manyDgraphRoles struct {
	Roles []dgraphRole `json:"roles"`
}

type dgraphRole struct {
	ID   string `json:"uid,omitempty"`
	Name string `json:"name,omitempty"`

	Users []dgraphUser `json:"~roles,omitempty"`

	DType []string `json:"dgraph.type,omitempty"`
}

// GetAllRoles roles out of dgraph
func (d *dgraphRepo) GetAllRoles() (*models.ManyRoles, error) {
	dRoles := manyDgraphRoles{}
	roles := models.ManyRoles{}
	ctx := d.buildAuthContext(context.Background())
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

	resp, err := txn.Query(ctx, q)
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

// GetRole out of dgraph by ID
func (d *dgraphRepo) GetRole(id string) (*models.Role, error) {
	dRoles := manyDgraphRoles{}
	role := models.Role{}
	ctx := d.buildAuthContext(context.Background())
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

	resp, err := txn.QueryWithVars(ctx, q, variables)
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

package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

// body comment
// swagger:response Role
type role struct {
	// in:body
	Body restRole
}

// body comment
// swagger:response ManyRoles
type manyRoles struct {
	// in:body
	Body manyRestRoles `json:"roles"`
}

type restRole struct {
	ID   string `json:"uid,omitempty"`
	Name string `json:"name,omitempty"`

	Users []nestedUser `json:"users,omitempty"`
}

type nestedRole struct {
	ID   string `json:"uid,omitempty"`
	Name string `json:"name,omitempty"`
}

type manyRestRoles struct {
	Roles []nestedRole `json:"roles"`
}

func (restCtrl *restController) getAllRoles(c *gin.Context) {
	// swagger:route GET /roles roles getAllRoles
	// Fetch all roles
	// responses:
	//   200: ManyRoles

	callingUser, err := getUserFromContext(c)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	if role, sErr := restCtrl.Service.GetAllRoles(callingUser); sErr.Error != nil {
		respondWithServiceError(c, sErr)
	} else {
		roleResp := manyRestRoles{}
		copier.Copy(&roleResp, &role)
		c.JSON(http.StatusOK, roleResp)
	}
}

func (restCtrl *restController) getRole(c *gin.Context) {
	// swagger:route GET /roles/{id} roles getRole
	// Fetch a single role by ID
	// responses:
	//   200: Role

	id := c.Param("id")
	callingUser, err := getUserFromContext(c)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	if roles, sErr := restCtrl.Service.GetRole(callingUser, id); sErr.Error != nil {
		respondWithServiceError(c, sErr)
	} else {
		roleResp := manyRestRoles{}
		copier.Copy(&roleResp, &roles)
		c.JSON(http.StatusOK, roleResp)
	}
}

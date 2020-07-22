package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/atticuss/chefconnect/models"
)

// body comment
// swagger:response Role
type role struct {
	// in:body
	Body models.APIRole
}

// body comment
// swagger:response ManyRoles
type manyRoles struct {
	// in:body
	Body models.ManyAPIRoles `json:"roles"`
}

func (restCtrl *restController) getAllRoles(c *gin.Context) {
	// swagger:route GET /roles roles getAllRoles
	// Fetch all roles
	// responses:
	//   200: ManyRoles

	if resp, sErr := restCtrl.Service.GetAllRoles(); sErr.Error != nil {
		respondWithServiceError(c, sErr)
	} else {
		c.JSON(http.StatusOK, resp)
	}
}

func (restCtrl *restController) getRole(c *gin.Context) {
	// swagger:route GET /roles/{id} roles getRole
	// Fetch a single role by ID
	// responses:
	//   200: Role

	id := c.Param("id")

	if resp, sErr := restCtrl.Service.GetRole(id); sErr.Error != nil {
		respondWithServiceError(c, sErr)
	} else {
		c.JSON(http.StatusOK, resp)
	}
}

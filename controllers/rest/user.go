package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/atticuss/chefconnect/models"
)

// body comment
// swagger:parameters createUser updateUser
type userRequest struct {
	// in:body
	models.APIUser
}

// swagger:response User
type user struct {
	// in:body
	Body models.APIUser
}

// swagger:response ManyUsers
type manyUsers struct {
	// in:body
	Body models.ManyAPIUsers `json:"users"`
}

func (restCtrl *restController) getAllUsers(c *gin.Context) {
	// swagger:route GET /users users getAllUsers
	// Fetch all users
	// responses:
	//   200: ManyUsers

	if resp, sErr := restCtrl.Service.GetAllUsers(); sErr.Error != nil {
		respondWithServiceError(c, sErr)
	} else {
		c.JSON(http.StatusOK, resp)
	}
}

// GetUser handles the GET /users/{id} req for fetching a specific user
func (restCtrl *restController) getUser(c *gin.Context) {
	// swagger:route GET /users/{id} users getUser
	// Fetch all users
	// responses:
	//   200: User

	id := c.Param("id")

	if resp, sErr := restCtrl.Service.GetUser(id); sErr.Error != nil {
		respondWithServiceError(c, sErr)
	} else {
		c.JSON(http.StatusOK, resp)
	}
}

// TODO: prevent dupes - https://dgraph.io/docs/mutations/#example-of-conditional-upsert
func (restCtrl *restController) createUser(c *gin.Context) {
	// swagger:route POST /users users createUser
	// Create a new user
	// responses:
	//   200: User

	var user models.APIUser
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if resp, sErr := restCtrl.Service.CreateUser(user); sErr.Error != nil {
		respondWithServiceError(c, sErr)
	} else {
		c.JSON(http.StatusOK, resp)
	}
}

func (restCtrl *restController) updateUser(c *gin.Context) {
	// swagger:route PUT /users/{id} users updateUser
	// Update a user
	// responses:
	//   200: User

	var user models.APIUser
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.ID = c.Param("id")

	if resp, sErr := restCtrl.Service.UpdateUser(user); sErr.Error != nil {
		respondWithServiceError(c, sErr)
	} else {
		c.JSON(http.StatusOK, resp)
	}
}

func (restCtrl *restController) deleteUser(c *gin.Context) {
	// swagger:route DELETE /users/{id} users deleteUser
	// Delete a user
	// responses:
	//   200

	id := c.Param("id")

	if sErr := restCtrl.Service.DeleteUser(id); sErr.Error != nil {
		respondWithServiceError(c, sErr)
	} else {
		c.Status(http.StatusNoContent)
	}
}

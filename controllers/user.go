package controllers

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

// GetAllUsers handles the GET /users req for fetching all users
func (ctx *ControllerCtx) GetAllUsers(c *gin.Context) {
	// swagger:route GET /users users getAllUsers
	// Fetch all users
	// responses:
	//   200: ManyUsers

	if resp, sErr := ctx.Service.GetAllUsers(); sErr.Error != nil {
		respondWithServiceError(c, sErr)
	} else {
		c.JSON(http.StatusOK, resp)
	}
}

// GetUser handles the GET /users/{id} req for fetching a specific user
func (ctx *ControllerCtx) GetUser(c *gin.Context) {
	// swagger:route GET /users/{id} users getUser
	// Fetch all users
	// responses:
	//   200: User

	id := c.Param("id")

	if resp, sErr := ctx.Service.GetUser(id); sErr.Error != nil {
		respondWithServiceError(c, sErr)
	} else {
		c.JSON(http.StatusOK, resp)
	}
}

// CreateUser handles the POST /users req for creating a user
// TODO: prevent dupes - https://dgraph.io/docs/mutations/#example-of-conditional-upsert
func (ctx *ControllerCtx) CreateUser(c *gin.Context) {
	// swagger:route POST /users users createUser
	// Create a new user
	// responses:
	//   200: User

	var user models.APIUser
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if resp, sErr := ctx.Service.CreateUser(user); sErr.Error != nil {
		respondWithServiceError(c, sErr)
	} else {
		c.JSON(http.StatusOK, resp)
	}
}

// UpdateUser handles the PUT /users/{id} req for updating a user
func (ctx *ControllerCtx) UpdateUser(c *gin.Context) {
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

	if resp, sErr := ctx.Service.UpdateUser(user); sErr.Error != nil {
		respondWithServiceError(c, sErr)
	} else {
		c.JSON(http.StatusOK, resp)
	}
}

// DeleteUser handles the DELETE /users/{id} req for deleting a user
func (ctx *ControllerCtx) DeleteUser(c *gin.Context) {
	// swagger:route DELETE /users/{id} users deleteUser
	// Delete a user
	// responses:
	//   200

	id := c.Param("id")

	if sErr := ctx.Service.DeleteUser(id); sErr.Error != nil {
		respondWithServiceError(c, sErr)
	} else {
		c.Status(http.StatusNoContent)
	}
}

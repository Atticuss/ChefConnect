package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"

	"github.com/atticuss/chefconnect/models"
)

// body comment
// swagger:parameters createUser updateUser
type userRequest struct {
	// in:body
	Body restUser
}

// swagger:response User
type user struct {
	// in:body
	Body restUser
}

// swagger:response ManyUsers
type manyUsers struct {
	// in:body
	Body manyRestUsers `json:"users"`
}

type restUser struct {
	ID       string `json:"uid,omitempty"`
	Name     string `json:"name,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`

	Favorites []nestedRecipe `json:"favorites,omitempty"`
	Notes     []nestedNote   `json:"notes,omitempty"`
	Ratings   []nestedRecipe `json:"ratings,omitempty"`
	Roles     []nestedRole   `json:"roles,omitempty"`
}

type nestedUser struct {
	ID          string `json:"uid,omitempty"`
	Name        string `json:"name,omitempty"`
	Username    string `json:"username,omitempty"`
	RatingScore int    `json:"ratingScore,omitempty"`
}

type manyRestUsers struct {
	Users []nestedUser `json:"users"`
}

func (restCtrl *restController) getAllUsers(c *gin.Context) {
	// swagger:route GET /users users getAllUsers
	// Fetch all users
	// responses:
	//   200: ManyUsers

	callingUser, err := getUserFromContext(c)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	if users, sErr := restCtrl.Service.GetAllUsers(callingUser); sErr.Error != nil {
		respondWithServiceError(c, sErr)
	} else {
		usersResp := manyRestUsers{}
		copier.Copy(&usersResp, &users)
		c.JSON(http.StatusOK, usersResp)
	}
}

// GetUser handles the GET /users/{id} req for fetching a specific user
func (restCtrl *restController) getUser(c *gin.Context) {
	// swagger:route GET /users/{id} users getUser
	// Fetch all users
	// responses:
	//   200: User

	id := c.Param("id")
	callingUser, err := getUserFromContext(c)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	if user, sErr := restCtrl.Service.GetUser(callingUser, id); sErr.Error != nil {
		respondWithServiceError(c, sErr)
	} else {
		userResp := restUser{}
		copier.Copy(&userResp, &user)
		c.JSON(http.StatusOK, userResp)
	}
}

// TODO: prevent dupes - https://dgraph.io/docs/mutations/#example-of-conditional-upsert
func (restCtrl *restController) createUser(c *gin.Context) {
	// swagger:route POST /users users createUser
	// Create a new user
	// responses:
	//   200: User

	var userReq restUser
	if err := c.ShouldBindJSON(&userReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{}
	copier.Copy(&user, &userReq)

	callingUser, err := getUserFromContext(c)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	if user, sErr := restCtrl.Service.CreateUser(callingUser, &user); sErr.Error != nil {
		respondWithServiceError(c, sErr)
	} else {
		userResp := restUser{}
		copier.Copy(&userResp, &user)
		c.JSON(http.StatusOK, userResp)
	}
}

func (restCtrl *restController) updateUser(c *gin.Context) {
	// swagger:route PUT /users/{id} users updateUser
	// Update a user
	// responses:
	//   200: User

	var userReq restUser
	if err := c.ShouldBindJSON(&userReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{}
	user.ID = c.Param("id")
	copier.Copy(&user, &userReq)

	callingUser, err := getUserFromContext(c)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	if user, sErr := restCtrl.Service.UpdateUser(callingUser, &user); sErr.Error != nil {
		respondWithServiceError(c, sErr)
	} else {
		userResp := restUser{}
		copier.Copy(&userResp, &user)
		c.JSON(http.StatusOK, userResp)
	}
}

func (restCtrl *restController) deleteUser(c *gin.Context) {
	// swagger:route DELETE /users/{id} users deleteUser
	// Delete a user
	// responses:
	//   200

	id := c.Param("id")
	callingUser, err := getUserFromContext(c)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	if sErr := restCtrl.Service.DeleteUser(callingUser, id); sErr.Error != nil {
		respondWithServiceError(c, sErr)
	} else {
		c.Status(http.StatusNoContent)
	}
}

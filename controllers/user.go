package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

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
func (ctx *ControllerCtx) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	// swagger:route GET /users users getAllUsers
	// Fetch all users
	// responses:
	//   200: ManyUsers

	if resp, sErr := ctx.Service.GetAllUsers(); sErr.Error != nil {
		respondWithServiceError(w, sErr)
	} else {
		respondWithJSON(w, http.StatusOK, resp)
	}
}

// GetUser handles the GET /users/{id} req for fetching a specific user
func (ctx *ControllerCtx) GetUser(w http.ResponseWriter, r *http.Request) {
	// swagger:route GET /users/{id} users getUser
	// Fetch all users
	// responses:
	//   200: User

	vars := mux.Vars(r)
	id := vars["id"]

	if resp, sErr := ctx.Service.GetUser(id); sErr.Error != nil {
		respondWithServiceError(w, sErr)
	} else {
		respondWithJSON(w, http.StatusOK, resp)
	}
}

// CreateUser handles the POST /users req for creating a user
// TODO: prevent dupes - https://dgraph.io/docs/mutations/#example-of-conditional-upsert
func (ctx *ControllerCtx) CreateUser(w http.ResponseWriter, r *http.Request) {
	// swagger:route POST /users users createUser
	// Create a new user
	// responses:
	//   200: User

	var user models.APIUser
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if resp, sErr := ctx.Service.CreateUser(user); sErr.Error != nil {
		respondWithServiceError(w, sErr)
	} else {
		respondWithJSON(w, http.StatusOK, resp)
	}
}

// UpdateUser handles the PUT /users/{id} req for updating a user
func (ctx *ControllerCtx) UpdateUser(w http.ResponseWriter, r *http.Request) {
	// swagger:route PUT /users/{id} users updateUser
	// Update a user
	// responses:
	//   200: User

	var user models.APIUser
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if resp, sErr := ctx.Service.UpdateUser(user); sErr.Error != nil {
		respondWithServiceError(w, sErr)
	} else {
		respondWithJSON(w, http.StatusOK, resp)
	}
}

// DeleteUser handles the DELETE /users/{id} req for deleting a user
func (ctx *ControllerCtx) DeleteUser(w http.ResponseWriter, r *http.Request) {
	// swagger:route DELETE /users/{id} users deleteUser
	// Delete a user
	// responses:
	//   200

	vars := mux.Vars(r)
	id := vars["id"]

	if sErr := ctx.Service.DeleteUser(id); sErr.Error != nil {
		respondWithServiceError(w, sErr)
	} else {
		respondWithJSON(w, http.StatusOK, models.Ingredient{})
	}
}

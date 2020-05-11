package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/atticuss/chefconnect/models"
	"github.com/gorilla/mux"
)

// body comment
// swagger:parameters createUser updateUser
type userRequest struct {
	// in:body
	models.UserResponse
}

// swagger:response User
type user struct {
	// in:body
	Body models.UserResponse
}

// swagger:response ManyUsers
type manyUsers struct {
	// in:body
	Body []models.UserResponse
}

// GetAllUsers handles the GET /users req for fetching all users
func (ctx *ControllerCtx) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	// swagger:route GET /users users getAllUsers
	// Fetch all users
	// responses:
	//   200: ManyUsers

	resp, err := models.GetAllUsers(ctx.DgraphClient)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	//cleanResp := models.ManyUsersResponse{}
	//for _, user := range resp.Users {
	//	cleanResp.Users = append(cleanResp.Users, models.UserResponse(user))
	//}

	//respondWithJSON(w, http.StatusOK, cleanResp)
	respondWithJSON(w, http.StatusOK, resp)
}

// GetUser handles the GET /users/{id} req for fetching a specific user
func (ctx *ControllerCtx) GetUser(w http.ResponseWriter, r *http.Request) {
	// swagger:route GET /users/{id} users getUser
	// Fetch all users
	// responses:
	//   200: User

	vars := mux.Vars(r)
	id := vars["id"]

	user := models.User{ID: id}
	if err := user.GetUser(ctx.DgraphClient); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	//respondWithJSON(w, http.StatusOK, models.UserResponse(user))
	respondWithJSON(w, http.StatusOK, user)
}

// CreateUser handles the POST /users req for creating a user
// TODO: prevent dupes - https://dgraph.io/docs/mutations/#example-of-conditional-upsert
func (ctx *ControllerCtx) CreateUser(w http.ResponseWriter, r *http.Request) {
	// swagger:route POST /users users createUser
	// Create a new user
	// responses:
	//   200: User

	var user models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	err := ctx.Validator.Struct(user)
	if err != nil {
		respondWithValidationError(w, err, user)
		return
	}

	if err := user.CreateUser(ctx.DgraphClient); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, user)
}

// UpdateUser handles the PUT /users/{id} req for updating a user
func (ctx *ControllerCtx) UpdateUser(w http.ResponseWriter, r *http.Request) {
	// swagger:route PUT /users/{id} users updateUser
	// Update a user
	// responses:
	//   200: User

	var user models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	err := ctx.Validator.Struct(user)
	if err != nil {
		respondWithValidationError(w, err, user)
		return
	}

	if err := user.UpdateUser(ctx.DgraphClient); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, user)
}

// DeleteUser handles the DELETE /users/{id} req for deleting a user
func (ctx *ControllerCtx) DeleteUser(w http.ResponseWriter, r *http.Request) {
	// swagger:route DELETE /users/{id} users deleteUser
	// Delete a user
	// responses:
	//   200

	respondWithError(w, http.StatusNotImplemented, "Not implemented yet")
	return

	var user models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := user.DeleteUser(ctx.DgraphClient); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, user)
}

package controllers

import (
	"net/http"

	"github.com/atticuss/chefconnect/models"
	"github.com/gorilla/mux"
)

// swagger:response User
type user struct {
	// in:body
	Body models.User
}

// GetAllUsers handles the GET /users req for fetching all users
func (ctx *ControllerCtx) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	// swagger:operation GET /users users Users
	//
	// Fetch all users
	// ---
	// consumes:
	// - application/json
	// produces:
	// - application/json
	// responses:
	//   '200':
	//     description: Successfully fetched
	//     schema:
	//       type: object
	//       properties:
	//         users:
	//           type: array
	//           items:
	//             "$ref": "#/responses/User"

	resp, err := models.GetAllUsers(ctx.DgraphClient)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	cleanResp := models.ManyUsersResponse{}
	for _, user := range resp.Users {
		cleanResp.Users = append(cleanResp.Users, models.UserResponse(user))
	}

	respondWithJSON(w, http.StatusOK, cleanResp)
}

// GetUser handles the GET /users/{id} req for fetching a specific user
func (ctx *ControllerCtx) GetUser(w http.ResponseWriter, r *http.Request) {
	// swagger:operation GET /users/{id} users Users
	//
	// Fetches a single user by ID
	// ---
	// consumes:
	// - application/json
	// produces:
	// - application/json
	// parameters:
	// - name: id
	//   in: path
	//   description: ID of the user to be returned.
	//   required: true
	//   type: string
	// responses:
	//   '200':
	//     description: Successfully fetched
	//     schema:
	//       "$ref": "#/responses/User"
	vars := mux.Vars(r)
	id := vars["id"]

	user := models.User{ID: id}
	if err := user.GetUser(ctx.DgraphClient); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, models.UserResponse(user))
}

package controllers

import (
	"net/http"

	"github.com/atticuss/chefconnect/models"
	"github.com/gorilla/mux"
)

// swagger:response Category
type category struct {
	// in:body
	Body models.Category
}

// GetAllCategories handles the GET /categories req for fetching all categories
func (ctx *ControllerCtx) GetAllCategories(w http.ResponseWriter, r *http.Request) {
	// swagger:operation GET /categories categories Categories
	//
	// Fetch all categories
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
	//         categories:
	//           type: array
	//           items:
	//             "$ref": "#/responses/Category"

	resp, err := models.GetAllCategories(ctx.DgraphClient)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, resp)
}

// GetCategory handles the GET /categories/{id} req for fetching a specific user
func (ctx *ControllerCtx) GetCategory(w http.ResponseWriter, r *http.Request) {
	// swagger:operation GET /categories/{id} category Category
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
	//   description: ID of the category to be returned.
	//   required: true
	//   type: string
	// responses:
	//   '200':
	//     description: Successfully fetched
	//     schema:
	//       "$ref": "#/responses/Category"
	vars := mux.Vars(r)
	id := vars["id"]

	category := models.Category{ID: id}
	if err := category.GetCategory(ctx.DgraphClient); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, category)
}

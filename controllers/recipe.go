package controllers

import (
	"net/http"

	"github.com/atticuss/chefconnect/models"
	"github.com/gorilla/mux"
)

// swagger:response Recipe
type recipe struct {
	// in:body
	Body models.Recipe
}

// GetRecipe handles the GET /recipes/{id} req for fetching a specific recipes
func (ctx *ControllerCtx) GetRecipe(w http.ResponseWriter, r *http.Request) {
	// swagger:operation GET /recipes/{id} recipes Recipes
	//
	// Fetches a single recipe by ID
	// ---
	// consumes:
	// - application/json
	// produces:
	// - application/json
	// parameters:
	// - name: id
	//   in: path
	//   description: ID of the recipe to be returned.
	//   required: true
	//   type: string
	// responses:
	//   '200':
	//     description: Successfully fetched
	//     schema:
	//       "$ref": "#/responses/Recipe"
	vars := mux.Vars(r)
	id := vars["id"]

	recipe := models.Recipe{ID: id}
	if err := recipe.GetRecipe(ctx.DgraphClient); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, models.ResponseRecipe(recipe))
}

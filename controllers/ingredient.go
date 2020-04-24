package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/atticuss/chefconnect/models"
	"github.com/gorilla/mux"
)

// swagger:response Ingredient
type ingredient struct {
	// in:body
	Body models.Ingredient
}

// GetAllIngredients handles the GET /ingredients req for fetching all ingredients
func (ctx *ControllerCtx) GetAllIngredients(w http.ResponseWriter, r *http.Request) {
	// swagger:operation GET /ingredients ingredients Ingredients
	//
	// Fetch all ingredients
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
	//         ingredients:
	//           type: array
	//           items:
	//             "$ref": "#/responses/Ingredient"

	resp, err := models.GetAllIngredients(ctx.DgraphClient)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, resp)
}

// GetIngredient handles the GET /ingredients/{id} req for fetching a specific ingredient
func (ctx *ControllerCtx) GetIngredient(w http.ResponseWriter, r *http.Request) {
	// swagger:operation GET /ingredients/{id} ingredients Ingredients
	//
	// Fetches a single ingredient by ID
	// ---
	// consumes:
	// - application/json
	// produces:
	// - application/json
	// parameters:
	// - name: id
	//   in: path
	//   description: ID of the ingredient to be returned.
	//   required: true
	//   type: string
	// responses:
	//   '200':
	//     description: Successfully fetched
	//     schema:
	//       "$ref": "#/responses/Ingredient"
	vars := mux.Vars(r)
	id := vars["id"]

	i := models.Ingredient{ID: id}
	if err := i.GetIngredient(ctx.DgraphClient); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, i)
}

// CreateIngredient handles the POST /ingredients/{id} req for creating an ingredient
// TODO: prevent dupes - https://dgraph.io/docs/mutations/#example-of-conditional-upsert
func (ctx *ControllerCtx) CreateIngredient(w http.ResponseWriter, r *http.Request) {
	// swagger:operation POST /ingredients ingredients Ingredients
	//
	// Create a new ingredient
	// ---
	// consumes:
	// - application/json
	// produces:
	// - application/json
	// parameters:
	// - name: id
	//   in: body
	//   required: true
	//   schema:
	//     "$ref": "#/responses/Ingredient"
	// responses:
	//   '200':
	//     description: Successfully fetched
	//     schema:
	//       "$ref": "#/responses/Ingredient"

	var i models.Ingredient
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&i); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	err := ctx.Validator.Struct(i)
	if err != nil {
		respondWithValidationError(w, err, i)
		return
	}

	if err := i.CreateIngredient(ctx.DgraphClient); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, i)
}

// UpdateIngredient handles the PUT /ingredients/{id} req for updating an ingredient
func (ctx *ControllerCtx) UpdateIngredient(w http.ResponseWriter, r *http.Request) {
	// swagger:operation PUT /ingredients ingredients Ingredients
	//
	// Create a new ingredient
	// ---
	// consumes:
	// - application/json
	// produces:
	// - application/json
	// parameters:
	// - name: id
	//   in: body
	//   required: true
	//   schema:
	//     "$ref": "#/responses/Ingredient"
	// responses:
	//   '200':
	//     description: Successfully updated

	var i models.Ingredient
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&i); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	err := ctx.Validator.Struct(i)
	if err != nil {
		respondWithValidationError(w, err, i)
		return
	}

	if err := i.CreateIngredient(ctx.DgraphClient); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, i)
}

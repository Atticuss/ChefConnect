package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/atticuss/chefconnect/models"

	"github.com/gorilla/mux"
	"github.com/jinzhu/copier"
)

// body comment
// swagger:parameters createIngredient updateIngredient
type ingredientRequest struct {
	// in:body
	models.NestedIngredient
}

// swagger:response Ingredient
type ingredient struct {
	// in:body
	Body models.IngredientResponse
}

// swagger:response ManyIngredients
type manyIngredients struct {
	// in:body
	Body models.ManyIngredientsResponse `json:"ingredients"`
}

// GetAllIngredients handles the GET /ingredients req for fetching all ingredients
func (ctx *ControllerCtx) GetAllIngredients(w http.ResponseWriter, r *http.Request) {
	// swagger:route GET /ingredients ingredients getAllIngredients
	// Fetch all ingredients
	// responses:
	//   200: ManyIngredients

	manyIngredients, err := models.GetAllIngredients(ctx.DgraphClient)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	apiResp := models.ManyIngredientsResponse{}
	copier.Copy(&apiResp, &manyIngredients)
	respondWithJSON(w, http.StatusOK, apiResp)
}

// GetIngredient handles the GET /ingredients/{id} req for fetching a specific ingredient
func (ctx *ControllerCtx) GetIngredient(w http.ResponseWriter, r *http.Request) {
	// swagger:route GET /ingredients/{id} ingredients getIngredient
	// Fetches a single ingredient by ID
	// responses:
	//   200: Ingredient

	vars := mux.Vars(r)
	id := vars["id"]

	ingredient := models.Ingredient{ID: id}
	if err := ingredient.GetIngredient(ctx.DgraphClient); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	apiResp := models.IngredientResponse{}
	copier.Copy(&apiResp, &ingredient)
	respondWithJSON(w, http.StatusOK, apiResp)
}

// CreateIngredient handles the POST /ingredients req for creating an ingredient
// TODO: prevent dupes - https://dgraph.io/docs/mutations/#example-of-conditional-upsert
func (ctx *ControllerCtx) CreateIngredient(w http.ResponseWriter, r *http.Request) {
	// swagger:route POST /ingredients ingredients createIngredient
	// Create a new ingredient
	// responses:
	//   200: Ingredient

	var ingredient models.Ingredient
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&ingredient); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	err := ctx.Validator.Struct(ingredient)
	if err != nil {
		respondWithValidationError(w, err, ingredient)
		return
	}

	if err := ingredient.CreateIngredient(ctx.DgraphClient); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	apiResp := models.IngredientResponse{}
	copier.Copy(&apiResp, &ingredient)
	respondWithJSON(w, http.StatusOK, apiResp)
}

// UpdateIngredient handles the PUT /ingredients/{id} req for updating an ingredient
func (ctx *ControllerCtx) UpdateIngredient(w http.ResponseWriter, r *http.Request) {
	// swagger:route PUT /ingredients/{id} ingredients updateIngredient
	// Update an ingredient
	// responses:
	//   200: Ingredient

	var ingredient models.Ingredient
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&ingredient); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	err := ctx.Validator.Struct(ingredient)
	if err != nil {
		respondWithValidationError(w, err, ingredient)
		return
	}

	if err := ingredient.UpdateIngredient(ctx.DgraphClient); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	apiResp := models.IngredientResponse{}
	copier.Copy(&apiResp, &ingredient)
	respondWithJSON(w, http.StatusOK, apiResp)
}

// DeleteIngredient handles the DELETE /ingredients/{id} req for deleting an ingredient
func (ctx *ControllerCtx) DeleteIngredient(w http.ResponseWriter, r *http.Request) {
	// swagger:route DELETE /ingredients/{id} ingredients deleteIngredient
	// Delete an ingredient
	// responses:
	//   200

	respondWithError(w, http.StatusNotImplemented, "Not implemented yet")
	return

	var ingredient models.Ingredient
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&ingredient); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := ingredient.DeleteIngredient(ctx.DgraphClient); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, nil)
}

package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/atticuss/chefconnect/models"
)

// body comment
// swagger:parameters createIngredient updateIngredient
type ingredientRequest struct {
	// in:body
	models.APIIngredient
}

// swagger:response Ingredient
type ingredient struct {
	// in:body
	Body models.APIIngredient
}

// swagger:response ManyIngredients
type manyIngredients struct {
	// in:body
	Body models.ManyAPIIngredients `json:"ingredients"`
}

// GetAllIngredients handles the GET /ingredients req for fetching all ingredients
func (ctx *ControllerCtx) GetAllIngredients(w http.ResponseWriter, r *http.Request) {
	// swagger:route GET /ingredients ingredients getAllIngredients
	// Fetch all ingredients
	// responses:
	//   200: ManyIngredients

	if resp, sErr := ctx.Service.GetAllIngredients(); sErr.Error != nil {
		respondWithServiceError(w, sErr)
	} else {
		respondWithJSON(w, http.StatusOK, resp)
	}
}

// GetIngredient handles the GET /ingredients/{id} req for fetching a specific ingredient
func (ctx *ControllerCtx) GetIngredient(w http.ResponseWriter, r *http.Request) {
	// swagger:route GET /ingredients/{id} ingredients getIngredient
	// Fetches a single ingredient by ID
	// responses:
	//   200: Ingredient

	vars := mux.Vars(r)
	id := vars["id"]

	if resp, sErr := ctx.Service.GetIngredient(id); sErr.Error != nil {
		respondWithServiceError(w, sErr)
	} else {
		respondWithJSON(w, http.StatusOK, resp)
	}
}

// CreateIngredient handles the POST /ingredients req for creating an ingredient
// TODO: prevent dupes - https://dgraph.io/docs/mutations/#example-of-conditional-upsert
func (ctx *ControllerCtx) CreateIngredient(w http.ResponseWriter, r *http.Request) {
	// swagger:route POST /ingredients ingredients createIngredient
	// Create a new ingredient
	// responses:
	//   200: Ingredient

	var ingredient models.APIIngredient
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&ingredient); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if resp, sErr := ctx.Service.CreateIngredient(ingredient); sErr.Error != nil {
		respondWithServiceError(w, sErr)
	} else {
		respondWithJSON(w, http.StatusOK, resp)
	}
}

// UpdateIngredient handles the PUT /ingredients/{id} req for updating an ingredient
func (ctx *ControllerCtx) UpdateIngredient(w http.ResponseWriter, r *http.Request) {
	// swagger:route PUT /ingredients/{id} ingredients updateIngredient
	// Update an ingredient
	// responses:
	//   200: Ingredient

	var ingredient models.APIIngredient
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&ingredient); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	vars := mux.Vars(r)
	ingredient.ID = vars["id"]

	if resp, sErr := ctx.Service.UpdateIngredient(ingredient); sErr.Error != nil {
		respondWithServiceError(w, sErr)
	} else {
		respondWithJSON(w, http.StatusOK, resp)
	}
}

// DeleteIngredient handles the DELETE /ingredients/{id} req for deleting an ingredient
func (ctx *ControllerCtx) DeleteIngredient(w http.ResponseWriter, r *http.Request) {
	// swagger:route DELETE /ingredients/{id} ingredients deleteIngredient
	// Delete an ingredient
	// responses:
	//   200

	vars := mux.Vars(r)
	id := vars["id"]

	if sErr := ctx.Service.DeleteIngredient(id); sErr.Error != nil {
		respondWithServiceError(w, sErr)
	} else {
		respondWithJSON(w, http.StatusOK, models.Ingredient{})
	}
}

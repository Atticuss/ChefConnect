package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/atticuss/chefconnect/models"
	"github.com/gorilla/mux"
)

// GetAllIngredients handles the GET /ingredients req for fetching all ingredients
func (ctx *ControllerCtx) GetAllIngredients(w http.ResponseWriter, r *http.Request) {
	resp, err := models.GetAllIngredients(ctx.DgraphClient)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, resp)
}

// GetIngredient handles the GET /ingredients/{id} req for fetching a specific ingredient
func (ctx *ControllerCtx) GetIngredient(w http.ResponseWriter, r *http.Request) {
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
func (ctx *ControllerCtx) CreateIngredient(w http.ResponseWriter, r *http.Request) {
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

package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/atticuss/chefconnect/models"
)

// body comment
// swagger:parameters createRecipe udpateRecipe
type recipeRequest struct {
	// in:body
	models.APIRecipe
}

// swagger:response Recipe
type recipe struct {
	// in:body
	Body models.APIRecipe
}

// swagger:response ManyRecipes
type manyRecipes struct {
	// in:body
	Body models.ManyAPIRecipes `json:"recipes"`
}

// GetAllRecipes handles the GET /recipes req for fetching all recipes
func (ctx *ControllerCtx) GetAllRecipes(w http.ResponseWriter, r *http.Request) {
	// swagger:route GET /recipes recipes getAllRecipes
	// Fetch all recipes
	// responses:
	//   200: ManyRecipes

	if resp, sErr := ctx.ServiceCtx.GetAllRecipes(); sErr.Error != nil {
		respondWithServiceError(w, sErr)
	} else {
		respondWithJSON(w, http.StatusOK, resp)
	}
}

// GetRecipe handles the GET /recipes/{id} req for fetching a specific recipes
func (ctx *ControllerCtx) GetRecipe(w http.ResponseWriter, r *http.Request) {
	// swagger:route GET /recipes/{id} recipes getRecipe
	// Fetch a recipe by ID
	// responses:
	//   200: Recipe

	vars := mux.Vars(r)
	id := vars["id"]

	if resp, sErr := ctx.ServiceCtx.GetRecipe(id); sErr.Error != nil {
		respondWithServiceError(w, sErr)
	} else {
		respondWithJSON(w, http.StatusOK, resp)
	}
}

// CreateRecipe handles the POST /recipes req for creating a recipe
// TODO: prevent dupes - https://dgraph.io/docs/mutations/#example-of-conditional-upsert
func (ctx *ControllerCtx) CreateRecipe(w http.ResponseWriter, r *http.Request) {
	// swagger:route POST /recipes recipes createRecipe
	// Create a new recipe
	// responses:
	//   200: Recipe

	var recipe models.Recipe
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&recipe); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if resp, sErr := ctx.ServiceCtx.CreateRecipe(recipe); sErr.Error != nil {
		respondWithServiceError(w, sErr)
	} else {
		respondWithJSON(w, http.StatusOK, resp)
	}
}

// UpdateRecipe handles the PUT /recipes/{id} req for updating a recipe
func (ctx *ControllerCtx) UpdateRecipe(w http.ResponseWriter, r *http.Request) {
	// swagger:route PUT /recipes/{id} recipes updateRecipe
	// Update a recipe
	// responses:
	//   200: Recipe

	var recipe models.Recipe
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&recipe); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	vars := mux.Vars(r)
	recipe.ID = vars["id"]

	if resp, sErr := ctx.ServiceCtx.UpdateRecipe(recipe); sErr.Error != nil {
		respondWithServiceError(w, sErr)
	} else {
		respondWithJSON(w, http.StatusOK, resp)
	}
}

// DeleteRecipe handles the DELETE /recipes/{id} req for deleting a recipe
func (ctx *ControllerCtx) DeleteRecipe(w http.ResponseWriter, r *http.Request) {
	// swagger:route DELETE /recipes/{id} recipes deleteRecipe
	// Delete a recipe
	// responses:
	//   200

	respondWithError(w, http.StatusNotImplemented, "Not implemented yet")
	return

	vars := mux.Vars(r)
	id := vars["id"]

	if sErr := ctx.ServiceCtx.DeleteRecipe(id); sErr.Error != nil {
		respondWithServiceError(w, sErr)
	} else {
		respondWithJSON(w, http.StatusOK, models.Ingredient{})
	}
}

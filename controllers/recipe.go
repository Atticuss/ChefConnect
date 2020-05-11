package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/atticuss/chefconnect/models"
	"github.com/gorilla/mux"
)

// body comment
// swagger:parameters createRecipe udpateRecipe
type recipeRequest struct {
	// in:body
	models.RecipeResponse
}

// swagger:response Recipe
type recipe struct {
	// in:body
	Body models.RecipeResponse
}

// swagger:response ManyRecipes
type manyRecipes struct {
	// in:body
	Body []models.RecipeResponse
}

// GetAllRecipes handles the GET /recipes req for fetching all recipes
func (ctx *ControllerCtx) GetAllRecipes(w http.ResponseWriter, r *http.Request) {
	// swagger:route GET /recipes recipes getAllRecipes
	// Fetch all recipes
	// responses:
	//   200: ManyRecipes

	resp, err := models.GetAllRecipes(ctx.DgraphClient)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, resp)
}

// GetRecipe handles the GET /recipes/{id} req for fetching a specific recipes
func (ctx *ControllerCtx) GetRecipe(w http.ResponseWriter, r *http.Request) {
	// swagger:route GET /recipes/{id} recipes getRecipe
	// Fetch a recipe by ID
	// responses:
	//   200: Recipe

	vars := mux.Vars(r)
	id := vars["id"]

	recipe := models.Recipe{ID: id}
	if err := recipe.GetRecipe(ctx.DgraphClient); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	//respondWithJSON(w, http.StatusOK, models.RecipeResponse(recipe))
	respondWithJSON(w, http.StatusOK, recipe)
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

	err := ctx.Validator.Struct(recipe)
	if err != nil {
		respondWithValidationError(w, err, recipe)
		return
	}

	if err := recipe.CreateRecipe(ctx.DgraphClient); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, recipe)
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

	err := ctx.Validator.Struct(recipe)
	if err != nil {
		respondWithValidationError(w, err, recipe)
		return
	}

	if err := recipe.UpdateRecipe(ctx.DgraphClient); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, recipe)
}

// DeleteRecipe handles the DELETE /recipes/{id} req for deleting a recipe
func (ctx *ControllerCtx) DeleteRecipe(w http.ResponseWriter, r *http.Request) {
	// swagger:route DELETE /recipes/{id} recipes deleteRecipe
	// Delete a recipe
	// responses:
	//   200

	respondWithError(w, http.StatusNotImplemented, "Not implemented yet")
	return

	var recipe models.Recipe
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&recipe); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := recipe.DeleteRecipe(ctx.DgraphClient); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, recipe)
}

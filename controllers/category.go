package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/atticuss/chefconnect/models"

	"github.com/gorilla/mux"
)

// body comment
// swagger:parameters createCategory updateCategory
type categoryRequest struct {
	// in:body
	models.NestedCategory
}

// body comment
// swagger:response Category
type category struct {
	// in:body
	Body models.CategoryResponse
}

// body comment
// swagger:response ManyCategories
type manyCategories struct {
	// in:body
	Body []models.CategoryResponse
}

// GetAllCategories handles the GET /categories req for fetching all categories
func (ctx *ControllerCtx) GetAllCategories(w http.ResponseWriter, r *http.Request) {
	// swagger:route GET /categories categories getAllCategories
	// Fetch all categories
	// responses:
	//   200: ManyCategories

	resp, err := models.GetAllCategories(ctx.DgraphClient)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, resp)
}

// GetCategory handles the GET /categories/{id} req for fetching a specific user
func (ctx *ControllerCtx) GetCategory(w http.ResponseWriter, r *http.Request) {
	// swagger:route GET /categories/{id} categories getCategory
	// Fetch a single category by ID
	// responses:
	//   200: Category

	vars := mux.Vars(r)
	id := vars["id"]

	category := models.Category{ID: id}
	if err := category.GetCategory(ctx.DgraphClient); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, category)
}

// CreateCategory handles the POST /categories req for creating a category
// TODO: prevent dupes - https://dgraph.io/docs/mutations/#example-of-conditional-upsert
func (ctx *ControllerCtx) CreateCategory(w http.ResponseWriter, r *http.Request) {
	// swagger:route POST /categories categories createCategory
	// Create a new category
	// responses:
	//   200: Category

	var category models.Category
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&category); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	err := ctx.Validator.Struct(category)
	if err != nil {
		respondWithValidationError(w, err, category)
		return
	}

	if err := category.CreateCategory(ctx.DgraphClient); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, category)
}

// UpdateCategory handles the PUT /categories/{id} req for updating a category
func (ctx *ControllerCtx) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	// swagger:route PUT /categories/{id} categories updateCategory
	// Update a category
	// responses:
	//   200: Category

	var category models.Category
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&category); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	err := ctx.Validator.Struct(category)
	if err != nil {
		respondWithValidationError(w, err, category)
		return
	}

	if err := category.UpdateCategory(ctx.DgraphClient); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, category)
}

// DeleteCategory handles the DELETE /categories/{id} req for deleting a category
func (ctx *ControllerCtx) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	// swagger:route DELETE /categories/{id} categories deleteCategory
	// Delete a category
	// responses:
	//   200

	respondWithError(w, http.StatusNotImplemented, "Not implemented yet")
	return

	var category models.Category
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&category); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := category.DeleteCategory(ctx.DgraphClient); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, category)
}

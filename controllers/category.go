package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/atticuss/chefconnect/models"
)

// body comment
// swagger:parameters createCategory updateCategory
type categoryRequest struct {
	// in:body
	models.APICategory
}

// body comment
// swagger:response Category
type category struct {
	// in:body
	Body models.APICategory
}

// body comment
// swagger:response ManyCategories
type manyCategories struct {
	// in:body
	Body models.ManyAPICategories `json:"categories"`
}

// GetAllCategories handles the GET /categories req for fetching all categories
func (ctx *ControllerCtx) GetAllCategories(w http.ResponseWriter, r *http.Request) {
	// swagger:route GET /categories categories getAllCategories
	// Fetch all categories
	// responses:
	//   200: ManyCategories

	if resp, sErr := ctx.Service.GetAllCategories(); sErr.Error != nil {
		respondWithServiceError(w, sErr)
	} else {
		respondWithJSON(w, http.StatusOK, resp)
	}
}

// GetCategory handles the GET /categories/{id} req for fetching a specific user
func (ctx *ControllerCtx) GetCategory(w http.ResponseWriter, r *http.Request) {
	// swagger:route GET /categories/{id} categories getCategory
	// Fetch a single category by ID
	// responses:
	//   200: Category

	vars := mux.Vars(r)
	id := vars["id"]

	if resp, sErr := ctx.Service.GetCategory(id); sErr.Error != nil {
		respondWithServiceError(w, sErr)
	} else {
		respondWithJSON(w, http.StatusOK, resp)
	}
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

	if resp, sErr := ctx.Service.CreateCategory(category); sErr.Error != nil {
		respondWithServiceError(w, sErr)
	} else {
		respondWithJSON(w, http.StatusOK, resp)
	}
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

	vars := mux.Vars(r)
	category.ID = vars["id"]

	if resp, sErr := ctx.Service.UpdateCategory(category); sErr.Error != nil {
		respondWithServiceError(w, sErr)
	} else {
		respondWithJSON(w, http.StatusOK, resp)
	}
}

// DeleteCategory handles the DELETE /categories/{id} req for deleting a category
func (ctx *ControllerCtx) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	// swagger:route DELETE /categories/{id} categories deleteCategory
	// Delete a category
	// responses:
	//   200

	vars := mux.Vars(r)
	id := vars["id"]

	if sErr := ctx.Service.DeleteCategory(id); sErr.Error != nil {
		respondWithServiceError(w, sErr)
	} else {
		respondWithJSON(w, http.StatusOK, models.Category{})
	}
}

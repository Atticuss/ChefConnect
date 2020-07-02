package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

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
func (ctlr *v1Controller) GetAllIngredients(c *gin.Context) {
	// swagger:route GET /ingredients ingredients getAllIngredients
	// Fetch all ingredients
	// responses:
	//   200: ManyIngredients

	if resp, sErr := ctlr.Service.GetAllIngredients(); sErr.Error != nil {
		respondWithServiceError(c, sErr)
	} else {
		c.JSON(http.StatusOK, resp)
	}
}

// GetIngredient handles the GET /ingredients/{id} req for fetching a specific ingredient
func (ctlr *v1Controller) GetIngredient(c *gin.Context) {
	// swagger:route GET /ingredients/{id} ingredients getIngredient
	// Fetches a single ingredient by ID
	// responses:
	//   200: Ingredient

	id := c.Param("id")

	if resp, sErr := ctlr.Service.GetIngredient(id); sErr.Error != nil {
		respondWithServiceError(c, sErr)
	} else {
		c.JSON(http.StatusOK, resp)
	}
}

// CreateIngredient handles the POST /ingredients req for creating an ingredient
// TODO: prevent dupes - https://dgraph.io/docs/mutations/#example-of-conditional-upsert
func (ctlr *v1Controller) CreateIngredient(c *gin.Context) {
	// swagger:route POST /ingredients ingredients createIngredient
	// Create a new ingredient
	// responses:
	//   200: Ingredient

	var ingredient models.APIIngredient
	if err := c.ShouldBindJSON(&ingredient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if resp, sErr := ctlr.Service.CreateIngredient(ingredient); sErr.Error != nil {
		respondWithServiceError(c, sErr)
	} else {
		c.JSON(http.StatusOK, resp)
	}
}

// UpdateIngredient handles the PUT /ingredients/{id} req for updating an ingredient
func (ctlr *v1Controller) UpdateIngredient(c *gin.Context) {
	// swagger:route PUT /ingredients/{id} ingredients updateIngredient
	// Update an ingredient
	// responses:
	//   200: Ingredient

	var ingredient models.APIIngredient
	if err := c.ShouldBindJSON(&ingredient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ingredient.ID = c.Param("id")

	if resp, sErr := ctlr.Service.UpdateIngredient(ingredient); sErr.Error != nil {
		respondWithServiceError(c, sErr)
	} else {
		c.JSON(http.StatusOK, resp)
	}
}

// DeleteIngredient handles the DELETE /ingredients/{id} req for deleting an ingredient
func (ctlr *v1Controller) DeleteIngredient(c *gin.Context) {
	// swagger:route DELETE /ingredients/{id} ingredients deleteIngredient
	// Delete an ingredient
	// responses:
	//   200

	id := c.Param("id")

	if sErr := ctlr.Service.DeleteIngredient(id); sErr.Error != nil {
		respondWithServiceError(c, sErr)
	} else {
		//c.JSON(http.StatusOK, []string{})
		c.Status(http.StatusNoContent)
	}
}

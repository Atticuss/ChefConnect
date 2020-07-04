package rest

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

func (restCtrl *restController) getAllIngredients(c *gin.Context) {
	// swagger:route GET /ingredients ingredients getAllIngredients
	// Fetch all ingredients
	// responses:
	//   200: ManyIngredients

	if resp, sErr := restCtrl.Service.GetAllIngredients(); sErr.Error != nil {
		respondWithServiceError(c, sErr)
	} else {
		c.JSON(http.StatusOK, resp)
	}
}

func (restCtrl *restController) getIngredient(c *gin.Context) {
	// swagger:route GET /ingredients/{id} ingredients getIngredient
	// Fetches a single ingredient by ID
	// responses:
	//   200: Ingredient

	id := c.Param("id")

	if resp, sErr := restCtrl.Service.GetIngredient(id); sErr.Error != nil {
		respondWithServiceError(c, sErr)
	} else {
		c.JSON(http.StatusOK, resp)
	}
}

// TODO: prevent dupes - https://dgraph.io/docs/mutations/#example-of-conditional-upsert
func (restCtrl *restController) createIngredient(c *gin.Context) {
	// swagger:route POST /ingredients ingredients createIngredient
	// Create a new ingredient
	// responses:
	//   200: Ingredient

	var ingredient models.APIIngredient
	if err := c.ShouldBindJSON(&ingredient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if resp, sErr := restCtrl.Service.CreateIngredient(ingredient); sErr.Error != nil {
		respondWithServiceError(c, sErr)
	} else {
		c.JSON(http.StatusOK, resp)
	}
}

func (restCtrl *restController) updateIngredient(c *gin.Context) {
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

	if resp, sErr := restCtrl.Service.UpdateIngredient(ingredient); sErr.Error != nil {
		respondWithServiceError(c, sErr)
	} else {
		c.JSON(http.StatusOK, resp)
	}
}

func (restCtrl *restController) deleteIngredient(c *gin.Context) {
	// swagger:route DELETE /ingredients/{id} ingredients deleteIngredient
	// Delete an ingredient
	// responses:
	//   200

	id := c.Param("id")

	if sErr := restCtrl.Service.DeleteIngredient(id); sErr.Error != nil {
		respondWithServiceError(c, sErr)
	} else {
		//c.JSON(http.StatusOK, []string{})
		c.Status(http.StatusNoContent)
	}
}

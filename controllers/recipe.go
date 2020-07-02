package controllers

import (
	"net/http"

	"github.com/atticuss/chefconnect/models"
	"github.com/gin-gonic/gin"
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
func (ctx *ControllerCtx) GetAllRecipes(c *gin.Context) {
	// swagger:route GET /recipes recipes getAllRecipes
	// Fetch all recipes
	// responses:
	//   200: ManyRecipes

	if resp, sErr := ctx.Service.GetAllRecipes(); sErr.Error != nil {
		respondWithServiceError(c, sErr)
	} else {
		c.JSON(http.StatusOK, resp)
	}
}

// GetRecipe handles the GET /recipes/{id} req for fetching a specific recipes
func (ctx *ControllerCtx) GetRecipe(c *gin.Context) {
	// swagger:route GET /recipes/{id} recipes getRecipe
	// Fetch a recipe by ID
	// responses:
	//   200: Recipe

	id := c.Param("id")

	if resp, sErr := ctx.Service.GetRecipe(id); sErr.Error != nil {
		respondWithServiceError(c, sErr)
	} else {
		c.JSON(http.StatusOK, resp)
	}
}

// CreateRecipe handles the POST /recipes req for creating a recipe
// TODO: prevent dupes - https://dgraph.io/docs/mutations/#example-of-conditional-upsert
func (ctx *ControllerCtx) CreateRecipe(c *gin.Context) {
	// swagger:route POST /recipes recipes createRecipe
	// Create a new recipe
	// responses:
	//   200: Recipe

	var recipe models.APIRecipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if resp, sErr := ctx.Service.CreateRecipe(recipe); sErr.Error != nil {
		respondWithServiceError(c, sErr)
	} else {
		c.JSON(http.StatusOK, resp)
	}
}

// UpdateRecipe handles the PUT /recipes/{id} req for updating a recipe
func (ctx *ControllerCtx) UpdateRecipe(c *gin.Context) {
	// swagger:route PUT /recipes/{id} recipes updateRecipe
	// Update a recipe
	// responses:
	//   200: Recipe

	var recipe models.APIRecipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	recipe.ID = c.Param("id")

	if resp, sErr := ctx.Service.UpdateRecipe(recipe); sErr.Error != nil {
		respondWithServiceError(c, sErr)
	} else {
		c.JSON(http.StatusOK, resp)
	}
}

// DeleteRecipe handles the DELETE /recipes/{id} req for deleting a recipe
func (ctx *ControllerCtx) DeleteRecipe(c *gin.Context) {
	// swagger:route DELETE /recipes/{id} recipes deleteRecipe
	// Delete a recipe
	// responses:
	//   200

	id := c.Param("id")

	if sErr := ctx.Service.DeleteRecipe(id); sErr.Error != nil {
		respondWithServiceError(c, sErr)
	} else {
		c.Status(http.StatusNoContent)
	}
}

package rest

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

func (restCtrl *restController) getAllRecipes(c *gin.Context) {
	// swagger:route GET /recipes recipes getAllRecipes
	// Fetch all recipes
	// responses:
	//   200: ManyRecipes

	if resp, sErr := restCtrl.Service.GetAllRecipes(); sErr.Error != nil {
		respondWithServiceError(c, sErr)
	} else {
		c.JSON(http.StatusOK, resp)
	}
}

func (restCtrl *restController) getRecipe(c *gin.Context) {
	// swagger:route GET /recipes/{id} recipes getRecipe
	// Fetch a recipe by ID
	// responses:
	//   200: Recipe

	id := c.Param("id")

	if resp, sErr := restCtrl.Service.GetRecipe(id); sErr.Error != nil {
		respondWithServiceError(c, sErr)
	} else {
		c.JSON(http.StatusOK, resp)
	}
}

// TODO: prevent dupes - https://dgraph.io/docs/mutations/#example-of-conditional-upsert
func (restCtrl *restController) createRecipe(c *gin.Context) {
	// swagger:route POST /recipes recipes createRecipe
	// Create a new recipe
	// responses:
	//   200: Recipe

	var recipe models.APIRecipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if resp, sErr := restCtrl.Service.CreateRecipe(recipe); sErr.Error != nil {
		respondWithServiceError(c, sErr)
	} else {
		c.JSON(http.StatusOK, resp)
	}
}

func (restCtrl *restController) updateRecipe(c *gin.Context) {
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

	if resp, sErr := restCtrl.Service.UpdateRecipe(recipe); sErr.Error != nil {
		respondWithServiceError(c, sErr)
	} else {
		c.JSON(http.StatusOK, resp)
	}
}

func (restCtrl *restController) setRecipeTags(c *gin.Context) {
	// swagger:route PUT /recipes/{id}/tags recipes setRecipeTags
	// Create a new recipe
	// responses:
	//   200: Recipe

	var recipe models.APIRecipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	recipe.ID = c.Param("id")

	if resp, sErr := restCtrl.Service.SetRecipeTags(recipe); sErr.Error != nil {
		respondWithServiceError(c, sErr)
	} else {
		c.JSON(http.StatusOK, resp)
	}
}

func (restCtrl *restController) deleteRecipe(c *gin.Context) {
	// swagger:route DELETE /recipes/{id} recipes deleteRecipe
	// Delete a recipe
	// responses:
	//   200

	id := c.Param("id")

	if sErr := restCtrl.Service.DeleteRecipe(id); sErr.Error != nil {
		respondWithServiceError(c, sErr)
	} else {
		c.Status(http.StatusNoContent)
	}
}

package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"

	"github.com/atticuss/chefconnect/models"
)

// body comment
// swagger:parameters createIngredient updateIngredient
type ingredientRequest struct {
	// in:body
	restIngredient
}

// swagger:response Ingredient
type ingredient struct {
	// in:body
	Body restIngredient
}

type restIngredient struct {
	ID     string `json:"uid,omitempty"`
	Name   string `json:"name,omitempty" validate:"required"`
	Amount string `json:"amount,omitempty"`

	Tags []nestedTag `json:"tags,omitempty"`
}

type nestedIngredient struct {
	ID     string `json:"uid,omitempty"`
	Name   string `json:"name,omitempty" validate:"required"`
	Amount string `json:"amount,omitempty"`
}

// swagger:response ManyIngredients
type manyIngredients struct {
	// in:body
	Ingredients []nestedIngredient `json:"ingredients"`
}

func (restCtrl *restController) getAllIngredients(c *gin.Context) {
	// swagger:route GET /ingredients ingredients getAllIngredients
	// Fetch all ingredients
	// responses:
	//   200: ManyIngredients

	ingredientsResp := manyIngredients{}
	callingUser, err := getUserFromContext(c)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	if ingredients, sErr := restCtrl.Service.GetAllIngredients(callingUser); sErr.Error != nil {
		respondWithServiceError(c, sErr)
	} else {
		copier.Copy(&ingredientsResp, &ingredients)
		c.JSON(http.StatusOK, ingredientsResp)
	}
}

func (restCtrl *restController) getIngredient(c *gin.Context) {
	// swagger:route GET /ingredients/{id} ingredients getIngredient
	// Fetches a single ingredient by ID
	// responses:
	//   200: Ingredient

	id := c.Param("id")
	callingUser, err := getUserFromContext(c)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	ingredientResp := restIngredient{}
	if ingredient, sErr := restCtrl.Service.GetIngredient(callingUser, id); sErr.Error != nil {
		respondWithServiceError(c, sErr)
	} else {
		copier.Copy(&ingredientResp, &ingredient)
		c.JSON(http.StatusOK, ingredientResp)
	}
}

// TODO: prevent dupes - https://dgraph.io/docs/mutations/#example-of-conditional-upsert
func (restCtrl *restController) createIngredient(c *gin.Context) {
	// swagger:route POST /ingredients ingredients createIngredient
	// Create a new ingredient
	// responses:
	//   200: Ingredient

	var ingredientReq restIngredient
	if err := c.ShouldBindJSON(&ingredientReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ingredient := models.Ingredient{}
	copier.Copy(&ingredient, &ingredientReq)

	callingUser, err := getUserFromContext(c)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	if ingredient, sErr := restCtrl.Service.CreateIngredient(callingUser, &ingredient); sErr.Error != nil {
		respondWithServiceError(c, sErr)
	} else {
		ingredientResp := restIngredient{}
		copier.Copy(&ingredientResp, &ingredient)
		c.JSON(http.StatusOK, ingredientResp)
	}
}

func (restCtrl *restController) updateIngredient(c *gin.Context) {
	// swagger:route PUT /ingredients/{id} ingredients updateIngredient
	// Update an ingredient
	// responses:
	//   200: Ingredient

	var ingredientReq restIngredient
	if err := c.ShouldBindJSON(&ingredientReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ingredientReq.ID = c.Param("id")
	ingredient := models.Ingredient{}
	copier.Copy(&ingredient, &ingredientReq)

	callingUser, err := getUserFromContext(c)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	if ingredient, sErr := restCtrl.Service.UpdateIngredient(callingUser, &ingredient); sErr.Error != nil {
		respondWithServiceError(c, sErr)
	} else {
		ingredientResp := restIngredient{}
		copier.Copy(&ingredientResp, &ingredient)
		c.JSON(http.StatusOK, ingredientResp)
	}
}

func (restCtrl *restController) deleteIngredient(c *gin.Context) {
	// swagger:route DELETE /ingredients/{id} ingredients deleteIngredient
	// Delete an ingredient
	// responses:
	//   200

	id := c.Param("id")
	callingUser, err := getUserFromContext(c)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	if sErr := restCtrl.Service.DeleteIngredient(callingUser, id); sErr.Error != nil {
		respondWithServiceError(c, sErr)
	} else {
		c.JSON(http.StatusOK, map[string]string{})
	}
}

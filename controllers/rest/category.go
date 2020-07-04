package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"

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

func (restCtrl *restController) getAllCategories(c *gin.Context) {
	// swagger:route GET /categories categories getAllCategories
	// Fetch all categories
	// responses:
	//   200: ManyCategories

	if resp, sErr := restCtrl.Service.GetAllCategories(); sErr.Error != nil {
		respondWithServiceError(c, sErr)
	} else {
		c.JSON(http.StatusOK, resp)
	}
}

func (restCtrl *restController) getCategory(c *gin.Context) {
	// swagger:route GET /categories/{id} categories getCategory
	// Fetch a single category by ID
	// responses:
	//   200: Category

	id := c.Param("id")

	if resp, sErr := restCtrl.Service.GetCategory(id); sErr.Error != nil {
		respondWithServiceError(c, sErr)
	} else {
		c.JSON(http.StatusOK, resp)
	}
}

// TODO: prevent dupes - https://dgraph.io/docs/mutations/#example-of-conditional-upsert
func (restCtrl *restController) createCategory(c *gin.Context) {
	// swagger:route POST /categories categories createCategory
	// Create a new category
	// responses:
	//   200: Category

	var category models.APICategory
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if resp, sErr := restCtrl.Service.CreateCategory(category); sErr.Error != nil {
		respondWithServiceError(c, sErr)
	} else {
		c.JSON(http.StatusOK, resp)
	}
}

func (restCtrl *restController) updateCategory(c *gin.Context) {
	// swagger:route PUT /categories/{id} categories updateCategory
	// Update a category
	// responses:
	//   200: Category

	var category models.APICategory
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category.ID = c.Param("id")

	if resp, sErr := restCtrl.Service.UpdateCategory(category); sErr.Error != nil {
		respondWithServiceError(c, sErr)
	} else {
		c.JSON(http.StatusOK, resp)
	}
}

func (restCtrl *restController) deleteCategory(c *gin.Context) {
	// swagger:route DELETE /categories/{id} categories deleteCategory
	// Delete a category
	// responses:
	//   200

	id := c.Param("id")

	if sErr := restCtrl.Service.DeleteCategory(id); sErr.Error != nil {
		respondWithServiceError(c, sErr)
	} else {
		c.Status(http.StatusNoContent)
	}
}

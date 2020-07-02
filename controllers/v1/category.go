package v1

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

// GetAllCategories handles the GET /categories req for fetching all categories
func (ctlr *v1Controller) GetAllCategories(c *gin.Context) {
	// swagger:route GET /categories categories getAllCategories
	// Fetch all categories
	// responses:
	//   200: ManyCategories

	if resp, sErr := ctlr.Service.GetAllCategories(); sErr.Error != nil {
		respondWithServiceError(c, sErr)
	} else {
		c.JSON(http.StatusOK, resp)
	}
}

// GetCategory handles the GET /categories/{id} req for fetching a specific user
func (ctlr *v1Controller) GetCategory(c *gin.Context) {
	// swagger:route GET /categories/{id} categories getCategory
	// Fetch a single category by ID
	// responses:
	//   200: Category

	id := c.Param("id")

	if resp, sErr := ctlr.Service.GetCategory(id); sErr.Error != nil {
		respondWithServiceError(c, sErr)
	} else {
		c.JSON(http.StatusOK, resp)
	}
}

// CreateCategory handles the POST /categories req for creating a category
// TODO: prevent dupes - https://dgraph.io/docs/mutations/#example-of-conditional-upsert
func (ctlr *v1Controller) CreateCategory(c *gin.Context) {
	// swagger:route POST /categories categories createCategory
	// Create a new category
	// responses:
	//   200: Category

	var category models.APICategory
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if resp, sErr := ctlr.Service.CreateCategory(category); sErr.Error != nil {
		respondWithServiceError(c, sErr)
	} else {
		c.JSON(http.StatusOK, resp)
	}
}

// UpdateCategory handles the PUT /categories/{id} req for updating a category
func (ctlr *v1Controller) UpdateCategory(c *gin.Context) {
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

	if resp, sErr := ctlr.Service.UpdateCategory(category); sErr.Error != nil {
		respondWithServiceError(c, sErr)
	} else {
		c.JSON(http.StatusOK, resp)
	}
}

// DeleteCategory handles the DELETE /categories/{id} req for deleting a category
func (ctlr *v1Controller) DeleteCategory(c *gin.Context) {
	// swagger:route DELETE /categories/{id} categories deleteCategory
	// Delete a category
	// responses:
	//   200

	id := c.Param("id")

	if sErr := ctlr.Service.DeleteCategory(id); sErr.Error != nil {
		respondWithServiceError(c, sErr)
	} else {
		c.Status(http.StatusNoContent)
	}
}

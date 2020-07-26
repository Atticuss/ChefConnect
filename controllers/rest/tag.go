package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/atticuss/chefconnect/models"
)

// body comment
// swagger:parameters createTag updateTag
type tagRequest struct {
	// in:body
	models.APITag
}

// body comment
// swagger:response Tag
type tag struct {
	// in:body
	Body models.APITag
}

// body comment
// swagger:response ManyTags
type manytags struct {
	// in:body
	Body models.ManyAPITags `json:"tags"`
}

func (restCtrl *restController) getAllTags(c *gin.Context) {
	// swagger:route GET /tags tags getAllTags
	// Fetch all tags
	// responses:
	//   200: ManyTags

	if resp, sErr := restCtrl.Service.GetAllTags(); sErr.Error != nil {
		respondWithServiceError(c, sErr)
	} else {
		c.JSON(http.StatusOK, resp)
	}
}

func (restCtrl *restController) getTag(c *gin.Context) {
	// swagger:route GET /tags/{id} tags getTag
	// Fetch a single tag by ID
	// responses:
	//   200: Tag

	id := c.Param("id")

	if resp, sErr := restCtrl.Service.GetTag(id); sErr.Error != nil {
		respondWithServiceError(c, sErr)
	} else {
		c.JSON(http.StatusOK, resp)
	}
}

// TODO: prevent dupes - https://dgraph.io/docs/mutations/#example-of-conditional-upsert
func (restCtrl *restController) createTag(c *gin.Context) {
	// swagger:route POST /tags tags createTag
	// Create a new tag
	// responses:
	//   200: Tag

	var tag models.APITag
	if err := c.ShouldBindJSON(&tag); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if resp, sErr := restCtrl.Service.CreateTag(tag); sErr.Error != nil {
		respondWithServiceError(c, sErr)
	} else {
		c.JSON(http.StatusOK, resp)
	}
}

func (restCtrl *restController) updateTag(c *gin.Context) {
	// swagger:route PUT /tags/{id} tags updateTag
	// Update a tag
	// responses:
	//   200: Tag

	var tag models.APITag
	if err := c.ShouldBindJSON(&tag); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tag.ID = c.Param("id")

	if resp, sErr := restCtrl.Service.UpdateTag(tag); sErr.Error != nil {
		respondWithServiceError(c, sErr)
	} else {
		c.JSON(http.StatusOK, resp)
	}
}

func (restCtrl *restController) deleteTag(c *gin.Context) {
	// swagger:route DELETE /tags/{id} tags deleteTag
	// Delete a tag
	// responses:
	//   200

	id := c.Param("id")

	if sErr := restCtrl.Service.DeleteTag(id); sErr.Error != nil {
		respondWithServiceError(c, sErr)
	} else {
		c.Status(http.StatusNoContent)
	}
}

package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"

	"github.com/atticuss/chefconnect/models"
)

// body comment
// swagger:parameters createTag updateTag
type tagRequest struct {
	// in:body
	Body restTag
}

// body comment
// swagger:response Tag
type tag struct {
	// in:body
	Body restTag
}

// body comment
// swagger:response ManyTags
type manytags struct {
	// in:body
	Body manyRestTags `json:"tags"`
}

type restTag struct {
	ID   string `json:"uid,omitempty"`
	Name string `json:"name,omitempty" validate:"required"`

	Recipes     []nestedRecipe     `json:"recipes,omitempty"`
	Ingredients []nestedIngredient `json:"ingredients,omitempty"`
}

type nestedTag struct {
	ID   string `json:"uid,omitempty"`
	Name string `json:"name,omitempty" validate:"required"`
}

type manyRestTags struct {
	Tags []nestedTag `json:"tags"`
}

func (restCtrl *restController) getAllTags(c *gin.Context) {
	// swagger:route GET /tags tags getAllTags
	// Fetch all tags
	// responses:
	//   200: ManyTags

	if tags, sErr := restCtrl.Service.GetAllTags(); sErr.Error != nil {
		respondWithServiceError(c, sErr)
	} else {
		tagResp := manyRestTags{}
		copier.Copy(&tagResp, &tags)
		c.JSON(http.StatusOK, tagResp)
	}
}

func (restCtrl *restController) getTag(c *gin.Context) {
	// swagger:route GET /tags/{id} tags getTag
	// Fetch a single tag by ID
	// responses:
	//   200: Tag

	id := c.Param("id")

	if tag, sErr := restCtrl.Service.GetTag(id); sErr.Error != nil {
		respondWithServiceError(c, sErr)
	} else {
		tagResp := restTag{}
		copier.Copy(&tagResp, &tag)
		c.JSON(http.StatusOK, tagResp)
	}
}

// TODO: prevent dupes - https://dgraph.io/docs/mutations/#example-of-conditional-upsert
func (restCtrl *restController) createTag(c *gin.Context) {
	// swagger:route POST /tags tags createTag
	// Create a new tag
	// responses:
	//   200: Tag

	var reqTag restTag
	if err := c.ShouldBindJSON(&reqTag); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tag := models.Tag{}
	copier.Copy(&tag, &reqTag)

	if tag, sErr := restCtrl.Service.CreateTag(&tag); sErr.Error != nil {
		respondWithServiceError(c, sErr)
	} else {
		tagResp := restTag{}
		copier.Copy(&tagResp, &tag)
		c.JSON(http.StatusOK, tagResp)
	}
}

func (restCtrl *restController) updateTag(c *gin.Context) {
	// swagger:route PUT /tags/{id} tags updateTag
	// Update a tag
	// responses:
	//   200: Tag

	var tagReq restTag
	if err := c.ShouldBindJSON(&tagReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tag := models.Tag{}
	tagReq.ID = c.Param("id")
	copier.Copy(&tag, &tagReq)

	if tag, sErr := restCtrl.Service.UpdateTag(&tag); sErr.Error != nil {
		respondWithServiceError(c, sErr)
	} else {
		tagResp := restTag{}
		copier.Copy(&tagResp, &tag)
		c.JSON(http.StatusOK, tagResp)
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

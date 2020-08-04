package rest

import (
	"net/http"

	"github.com/atticuss/chefconnect/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

// body comment
// swagger:parameters createRecipe udpateRecipe
type swaggerRecipeRequest struct {
	// in:body
	Body restRequestRecipe
}

// swagger:response Recipe
type swaggerRecipeResponse struct {
	// in:body
	Body restResponseRecipe
}

// swagger:response ManyRecipes
type swaggerManyRecipes struct {
	// in:body
	Body restResponseRecipe `json:"recipes"`
}

type restRequestRecipe struct {
	ID            string `json:"uid,omitempty"`
	Name          string `json:"name,omitempty"`
	URL           string `json:"url,omitempty"`
	Domain        string `json:"domain,omitempty"`
	Directions    string `json:"directions,omitempty"`
	PrepTime      int    `json:"prep_time,omitempty"`
	CookTime      int    `json:"cook_time,omitempty"`
	TotalServings int    `json:"total_servings,omitempty"`
	HasBeenTried  bool   `json:"has_been_tried,omitempty"`

	Ingredients    []nestedIngredient `json:"ingredients,omitempty"`
	Tags           []nestedTag        `json:"tags,omitempty"`
	RelatedRecipes []nestedRecipe     `json:"related_recipes,omitempty"`
}

type restResponseRecipe struct {
	ID            string `json:"uid,omitempty"`
	Name          string `json:"name,omitempty"`
	URL           string `json:"url,omitempty"`
	Domain        string `json:"domain,omitempty"`
	Directions    string `json:"directions,omitempty"`
	PrepTime      int    `json:"prep_time,omitempty"`
	CookTime      int    `json:"cook_time,omitempty"`
	TotalServings int    `json:"total_servings,omitempty"`
	HasBeenTried  bool   `json:"has_been_tried,omitempty"`

	CreatedBy      nestedUser         `json:"created_by,omitempty"`
	Ingredients    []nestedIngredient `json:"ingredients,omitempty"`
	Tags           []nestedTag        `json:"tags,omitempty"`
	RatedBy        []nestedUser       `json:"rated_by,omitempty"`
	RatingScore    []int              `json:"rating_score,omitempty"`
	FavoritedBy    []nestedUser       `json:"favorited_by,omitempty"`
	RelatedRecipes []nestedRecipe     `json:"related_recipes,omitempty"`
	Notes          []nestedNote       `json:"notes,omitempty"`
}

type nestedRecipe struct {
	ID   string `json:"uid,omitempty"`
	Name string `json:"name,omitempty" validate:"required"`
}

type manyRecipes struct {
	Recipes []nestedRecipe `json:"recipes"`
}

func (restCtrl *restController) getAllRecipes(c *gin.Context) {
	// swagger:route GET /recipes recipes getAllRecipes
	// Fetch all recipes
	// responses:
	//   200: ManyRecipes

	recipesResp := manyRecipes{}
	user, err := getUserFromContext(c)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	if recipes, sErr := restCtrl.Service.GetAllRecipes(user); sErr.Error != nil {
		respondWithServiceError(c, sErr)
	} else {
		copier.Copy(&recipesResp, recipes)
		c.JSON(http.StatusOK, recipesResp)
	}
}

func (restCtrl *restController) getRecipe(c *gin.Context) {
	// swagger:route GET /recipes/{id} recipes getRecipe
	// Fetch a recipe by ID
	// responses:
	//   200: Recipe

	id := c.Param("id")
	user, err := getUserFromContext(c)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	recipeResp := restResponseRecipe{}
	if recipe, sErr := restCtrl.Service.GetRecipe(user, id); sErr.Error != nil {
		respondWithServiceError(c, sErr)
	} else {
		copier.Copy(&recipeResp, &recipe)
		c.JSON(http.StatusOK, recipeResp)
	}
}

// TODO: prevent dupes - https://dgraph.io/docs/mutations/#example-of-conditional-upsert
func (restCtrl *restController) createRecipe(c *gin.Context) {
	// swagger:route POST /recipes recipes createRecipe
	// Create a new recipe
	// responses:
	//   200: Recipe

	var recipeReq restRequestRecipe
	if err := c.ShouldBindJSON(&recipeReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	recipe := models.Recipe{}
	copier.Copy(&recipe, &recipeReq)

	if recipe, sErr := restCtrl.Service.CreateRecipe(&recipe); sErr.Error != nil {
		respondWithServiceError(c, sErr)
	} else {
		recipeResp := restResponseRecipe{}
		copier.Copy(&recipeResp, &recipe)
		c.JSON(http.StatusOK, recipeResp)
	}
}

func (restCtrl *restController) updateRecipe(c *gin.Context) {
	// swagger:route PUT /recipes/{id} recipes updateRecipe
	// Update a recipe
	// responses:
	//   200: Recipe

	var recipeReq restRequestRecipe
	if err := c.ShouldBindJSON(&recipeReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	recipe := models.Recipe{}
	recipeReq.ID = c.Param("id")
	copier.Copy(&recipe, &recipeReq)

	if recipe, sErr := restCtrl.Service.UpdateRecipe(&recipe); sErr.Error != nil {
		respondWithServiceError(c, sErr)
	} else {
		recipeResp := restResponseRecipe{}
		copier.Copy(&recipeResp, &recipe)
		c.JSON(http.StatusOK, recipeResp)
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

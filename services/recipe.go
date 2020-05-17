package services

import (
	"github.com/jinzhu/copier"

	"github.com/atticuss/chefconnect/models"
)

// GetAllRecipes handles the business logic when a client requests all recipes
func (ctx *ServiceCtx) GetAllRecipes() (models.ManyAPIRecipes, ServiceError) {
	recipesResp := models.ManyAPIRecipes{}
	recipes, err := ctx.RecipeRepository.GetAll()
	if err != nil {
		return recipesResp, ServiceError{Error: err}
	}

	copier.Copy(&recipesResp, &recipes)

	return recipesResp, nilErr
}

// GetRecipe handles the business logic when a client requests a specific recipe
func (ctx *ServiceCtx) GetRecipe(id string) (models.APIRecipe, ServiceError) {
	recipeResp := models.APIRecipe{}
	recipe, err := ctx.RecipeRepository.Get(id)
	if err != nil {
		return recipeResp, ServiceError{Error: err}
	}

	copier.Copy(&recipeResp, &recipe)

	return recipeResp, nilErr
}

// CreateRecipe handles the business logic when a client creates a new recipe
func (ctx *ServiceCtx) CreateRecipe(recipe models.Recipe) (models.APIRecipe, ServiceError) {
	recipeResp := models.APIRecipe{}

	repoRecipe, err := ctx.RecipeRepository.Create(&recipe)
	if err != nil {
		return recipeResp, ServiceError{Error: err}
	}

	copier.Copy(&recipeResp, &repoRecipe)

	return recipeResp, nilErr
}

// UpdateRecipe handles the business logic when a client updates a recipe
func (ctx *ServiceCtx) UpdateRecipe(recipe models.Recipe) (models.APIRecipe, ServiceError) {
	recipeResp := models.APIRecipe{}

	repoRecipe, err := ctx.RecipeRepository.Update(&recipe)
	if err != nil {
		return recipeResp, ServiceError{Error: err}
	}

	copier.Copy(&recipeResp, &repoRecipe)

	return recipeResp, nilErr
}

// DeleteRecipe handles the business logic when a client deletes a recipe
func (ctx *ServiceCtx) DeleteRecipe(id string) ServiceError {
	err := ctx.RecipeRepository.Delete(id)
	if err != nil {
		return ServiceError{Error: err}
	}

	return nilErr
}

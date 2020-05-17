package services

import (
	"github.com/jinzhu/copier"

	"github.com/atticuss/chefconnect/models"
)

// GetAllIngredients handles the business logic when a client requests all ingredients
func (ctx *ServiceCtx) GetAllIngredients() (models.ManyAPIIngredients, ServiceError) {
	ingredientsResp := models.ManyAPIIngredients{}
	ingredients, err := ctx.IngredientRepository.GetAll()
	if err != nil {
		return ingredientsResp, ServiceError{Error: err}
	}

	copier.Copy(&ingredientsResp, &ingredients)

	return ingredientsResp, nilErr
}

// GetIngredient handles the business logic when a client requests a specific ingredient
func (ctx *ServiceCtx) GetIngredient(id string) (models.APIIngredient, ServiceError) {
	ingredientResp := models.APIIngredient{}
	ingredient, err := ctx.IngredientRepository.Get(id)
	if err != nil {
		return ingredientResp, ServiceError{Error: err}
	}

	copier.Copy(&ingredientResp, &ingredient)

	return ingredientResp, nilErr
}

// CreateIngredient handles the business logic when a client creates a new ingredient
func (ctx *ServiceCtx) CreateIngredient(ingredient models.Ingredient) (models.APIIngredient, ServiceError) {
	ingredientResp := models.APIIngredient{}

	repoIngredient, err := ctx.IngredientRepository.Create(&ingredient)
	if err != nil {
		return ingredientResp, ServiceError{Error: err}
	}

	copier.Copy(&ingredientResp, &repoIngredient)

	return ingredientResp, nilErr
}

// UpdateIngredient handles the business logic when a client updates an ingredient
func (ctx *ServiceCtx) UpdateIngredient(ingredient models.Ingredient) (models.APIIngredient, ServiceError) {
	ingredientResp := models.APIIngredient{}

	repoIngredient, err := ctx.IngredientRepository.Update(&ingredient)
	if err != nil {
		return ingredientResp, ServiceError{Error: err}
	}

	copier.Copy(&ingredientResp, &repoIngredient)

	return ingredientResp, nilErr
}

// DeleteIngredient handles the business logic when a client deletes an ingredient
func (ctx *ServiceCtx) DeleteIngredient(id string) ServiceError {
	err := ctx.IngredientRepository.Delete(id)
	if err != nil {
		return ServiceError{Error: err}
	}

	return nilErr
}

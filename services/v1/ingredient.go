package v1

import (
	"github.com/jinzhu/copier"

	"github.com/atticuss/chefconnect/models"
	"github.com/atticuss/chefconnect/services"
)

// GetAllIngredients handles the business logic when a client requests all ingredients
func (s *v1Service) GetAllIngredients() (models.ManyAPIIngredients, services.ServiceError) {
	ingredientsResp := models.ManyAPIIngredients{}
	ingredients, err := s.IngredientRepository.GetAll()
	if err != nil {
		return ingredientsResp, services.ServiceError{Error: err}
	}

	copier.Copy(&ingredientsResp, &ingredients)

	return ingredientsResp, nilErr
}

// GetIngredient handles the business logic when a client requests a specific ingredient
func (s *v1Service) GetIngredient(id string) (models.APIIngredient, services.ServiceError) {
	ingredientResp := models.APIIngredient{}
	ingredient, err := s.IngredientRepository.Get(id)
	if err != nil {
		return ingredientResp, services.ServiceError{Error: err}
	}

	copier.Copy(&ingredientResp, &ingredient)

	return ingredientResp, nilErr
}

// CreateIngredient handles the business logic when a client creates a new ingredient
func (s *v1Service) CreateIngredient(ingredient models.Ingredient) (models.APIIngredient, services.ServiceError) {
	ingredientResp := models.APIIngredient{}

	repoIngredient, err := s.IngredientRepository.Create(&ingredient)
	if err != nil {
		return ingredientResp, services.ServiceError{Error: err}
	}

	copier.Copy(&ingredientResp, &repoIngredient)

	return ingredientResp, nilErr
}

// UpdateIngredient handles the business logic when a client updates an ingredient
func (s *v1Service) UpdateIngredient(ingredient models.Ingredient) (models.APIIngredient, services.ServiceError) {
	ingredientResp := models.APIIngredient{}

	repoIngredient, err := s.IngredientRepository.Update(&ingredient)
	if err != nil {
		return ingredientResp, services.ServiceError{Error: err}
	}

	copier.Copy(&ingredientResp, &repoIngredient)

	return ingredientResp, nilErr
}

// DeleteIngredient handles the business logic when a client deletes an ingredient
func (s *v1Service) DeleteIngredient(id string) services.ServiceError {
	err := s.IngredientRepository.Delete(id)
	if err != nil {
		return services.ServiceError{Error: err}
	}

	return nilErr
}

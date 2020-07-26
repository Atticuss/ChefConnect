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
func (s *v1Service) CreateIngredient(apiIngredient models.APIIngredient) (models.APIIngredient, services.ServiceError) {
	ingredient := models.Ingredient{}
	copier.Copy(&ingredient, &apiIngredient)

	repoIngredient, err := s.IngredientRepository.Create(&ingredient)
	copier.Copy(&apiIngredient, &repoIngredient)

	if err != nil {
		return apiIngredient, services.ServiceError{Error: err}
	}

	return apiIngredient, nilErr
}

// UpdateIngredient handles the business logic when a client updates an ingredient
func (s *v1Service) UpdateIngredient(apiIngredient models.APIIngredient) (models.APIIngredient, services.ServiceError) {
	ingredient := models.Ingredient{}
	copier.Copy(&ingredient, &apiIngredient)

	repoIngredient, err := s.IngredientRepository.Update(&ingredient)
	copier.Copy(&apiIngredient, &repoIngredient)

	if err != nil {
		return apiIngredient, services.ServiceError{Error: err}
	}

	return apiIngredient, nilErr
}

// SetIngredientTags handles the business logic when a client tags a recipe
func (s *v1Service) SetIngredientTags(apiIngredient models.APIIngredient) (models.APIIngredient, services.ServiceError) {
	ingredient := models.Ingredient{}
	copier.Copy(&ingredient, &apiIngredient)

	repoIngredient, err := s.IngredientRepository.SetTags(&ingredient)
	if err != nil {
		return apiIngredient, services.ServiceError{Error: err}
	}

	apiIngredient = models.APIIngredient{}
	copier.Copy(&apiIngredient, &repoIngredient)

	return apiIngredient, nilErr
}

// DeleteIngredient handles the business logic when a client deletes an ingredient
func (s *v1Service) DeleteIngredient(id string) services.ServiceError {
	err := s.IngredientRepository.Delete(id)
	if err != nil {
		return services.ServiceError{Error: err}
	}

	return nilErr
}

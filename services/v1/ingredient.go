package v1

import (
	"github.com/atticuss/chefconnect/models"
	"github.com/atticuss/chefconnect/services"
)

// GetAllIngredients handles the business logic when a client requests all ingredients
func (s *v1Service) GetAllIngredients() (*models.ManyIngredients, *services.ServiceError) {
	ingredients, err := s.IngredientRepository.GetAll()
	if err != nil {
		return ingredients, &services.ServiceError{Error: err}
	}

	return ingredients, &nilErr
}

// GetIngredient handles the business logic when a client requests a specific ingredient
func (s *v1Service) GetIngredient(id string) (*models.Ingredient, *services.ServiceError) {
	ingredient, err := s.IngredientRepository.Get(id)
	if err != nil {
		return ingredient, &services.ServiceError{Error: err}
	}

	return ingredient, &nilErr
}

// CreateIngredient handles the business logic when a client creates a new ingredient
func (s *v1Service) CreateIngredient(ingredient *models.Ingredient) (*models.Ingredient, *services.ServiceError) {
	ingredient, err := s.IngredientRepository.Create(ingredient)
	if err != nil {
		return ingredient, &services.ServiceError{Error: err}
	}

	return ingredient, &nilErr
}

// UpdateIngredient handles the business logic when a client updates an ingredient
func (s *v1Service) UpdateIngredient(ingredient *models.Ingredient) (*models.Ingredient, *services.ServiceError) {
	ingredient, err := s.IngredientRepository.Update(ingredient)
	if err != nil {
		return ingredient, &services.ServiceError{Error: err}
	}

	return ingredient, &nilErr
}

// DeleteIngredient handles the business logic when a client deletes an ingredient
func (s *v1Service) DeleteIngredient(id string) *services.ServiceError {
	err := s.IngredientRepository.Delete(id)
	if err != nil {
		return &services.ServiceError{Error: err}
	}

	return &nilErr
}

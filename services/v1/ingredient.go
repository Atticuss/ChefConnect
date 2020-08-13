package v1

import (
	"errors"

	"github.com/atticuss/chefconnect/models"
	"github.com/atticuss/chefconnect/services"
)

// GetAllIngredients handles the business logic when a client requests all ingredients
func (s *v1Service) GetAllIngredients(callingUser *models.User) (*models.ManyIngredients, *services.ServiceError) {
	ingredients, err := s.IngredientRepository.GetAll()
	if err != nil {
		return ingredients, &services.ServiceError{Error: err}
	}

	return ingredients, &nilErr
}

// GetIngredient handles the business logic when a client requests a specific ingredient
func (s *v1Service) GetIngredient(callingUser *models.User, id string) (*models.Ingredient, *services.ServiceError) {
	ingredient, err := s.IngredientRepository.Get(id)
	if err != nil {
		return ingredient, &services.ServiceError{Error: err}
	}

	return ingredient, &nilErr
}

// CreateIngredient handles the business logic when a client creates a new ingredient
func (s *v1Service) CreateIngredient(callingUser *models.User, ingredient *models.Ingredient) (*models.Ingredient, *services.ServiceError) {
	ingredient, err := s.IngredientRepository.Create(ingredient)
	if err != nil {
		return ingredient, &services.ServiceError{Error: err}
	}

	return ingredient, &nilErr
}

// UpdateIngredient handles the business logic when a client updates an ingredient
func (s *v1Service) UpdateIngredient(callingUser *models.User, ingredient *models.Ingredient) (*models.Ingredient, *services.ServiceError) {
	ingredient, err := s.IngredientRepository.Update(ingredient)
	if err != nil {
		return ingredient, &services.ServiceError{Error: err}
	}

	return ingredient, &nilErr
}

// DeleteIngredient handles the business logic when a client deletes an ingredient
func (s *v1Service) DeleteIngredient(callingUser *models.User, id string) *services.ServiceError {
	ingredient, err := s.IngredientRepository.Get(id)
	if err != nil {
		return &services.ServiceError{Error: err}
	}

	if len(ingredient.Recipes) > 0 {
		return &services.ServiceError{Error: errors.New("resource is in use"), ErrorCode: services.ResourceInUse}
	}

	err = s.IngredientRepository.Delete(id)
	if err != nil {
		return &services.ServiceError{Error: err}
	}

	return &nilErr
}

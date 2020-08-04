package v1

import (
	"github.com/atticuss/chefconnect/models"
	"github.com/atticuss/chefconnect/services"
)

func cleanUnauthorizedRecipeFields(callingUser *models.User, recipe *models.Recipe) {
	unauthorized := false
	if len(callingUser.ID) == 0 {
		unauthorized = true
	}
	if callingUser.ID != recipe.CreatedBy.ID {
		unauthorized = true
	}

	if unauthorized {
		recipe.Notes = []models.Note{}
		recipe.CreatedBy = models.User{}
		recipe.FavoritedBy = []models.User{}
		recipe.RatedBy = []models.User{}
	}
}

// GetAllRecipes handles the business logic when a client requests all recipes
func (s *v1Service) GetAllRecipes(callingUser *models.User) (*models.ManyRecipes, *services.ServiceError) {
	recipes, err := s.RecipeRepository.GetAll()
	if err != nil {
		return recipes, &services.ServiceError{Error: err}
	}

	for idx := range recipes.Recipes {
		cleanUnauthorizedRecipeFields(callingUser, &recipes.Recipes[idx])
	}

	return recipes, &nilErr
}

// GetRecipe handles the business logic when a client requests a specific recipe
func (s *v1Service) GetRecipe(callingUser *models.User, id string) (*models.Recipe, *services.ServiceError) {
	recipe, err := s.RecipeRepository.Get(id)
	if err != nil {
		return recipe, &services.ServiceError{Error: err}
	}

	cleanUnauthorizedRecipeFields(callingUser, recipe)

	return recipe, &nilErr
}

// CreateRecipe handles the business logic when a client creates a new recipe
func (s *v1Service) CreateRecipe(recipe *models.Recipe) (*models.Recipe, *services.ServiceError) {
	recipe, err := s.RecipeRepository.Create(recipe)

	if err != nil {
		return recipe, &services.ServiceError{Error: err}
	}

	return recipe, &nilErr
}

// UpdateRecipe handles the business logic when a client updates a recipe
func (s *v1Service) UpdateRecipe(recipe *models.Recipe) (*models.Recipe, *services.ServiceError) {
	recipe, err := s.RecipeRepository.Update(recipe)
	if err != nil {
		return recipe, &services.ServiceError{Error: err}
	}

	return recipe, &nilErr
}

// DeleteRecipe handles the business logic when a client deletes a recipe
func (s *v1Service) DeleteRecipe(id string) *services.ServiceError {
	err := s.RecipeRepository.Delete(id)
	if err != nil {
		return &services.ServiceError{Error: err}
	}

	return &nilErr
}

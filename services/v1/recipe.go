package v1

import (
	"github.com/jinzhu/copier"

	"github.com/atticuss/chefconnect/models"
	"github.com/atticuss/chefconnect/services"
)

// GetAllRecipes handles the business logic when a client requests all recipes
func (s *v1Service) GetAllRecipes() (models.ManyAPIRecipes, services.ServiceError) {
	recipesResp := models.ManyAPIRecipes{}
	recipes, err := s.RecipeRepository.GetAll()
	if err != nil {
		return recipesResp, services.ServiceError{Error: err}
	}

	copier.Copy(&recipesResp, &recipes)

	return recipesResp, nilErr
}

// GetRecipe handles the business logic when a client requests a specific recipe
func (s *v1Service) GetRecipe(id string) (models.APIRecipe, services.ServiceError) {
	recipeResp := models.APIRecipe{}
	recipe, err := s.RecipeRepository.Get(id)
	if err != nil {
		return recipeResp, services.ServiceError{Error: err}
	}

	copier.Copy(&recipeResp, &recipe)

	return recipeResp, nilErr
}

// CreateRecipe handles the business logic when a client creates a new recipe
func (s *v1Service) CreateRecipe(recipe models.Recipe) (models.APIRecipe, services.ServiceError) {
	recipeResp := models.APIRecipe{}

	repoRecipe, err := s.RecipeRepository.Create(&recipe)
	if err != nil {
		return recipeResp, services.ServiceError{Error: err}
	}

	copier.Copy(&recipeResp, &repoRecipe)

	return recipeResp, nilErr
}

// UpdateRecipe handles the business logic when a client updates a recipe
func (s *v1Service) UpdateRecipe(recipe models.Recipe) (models.APIRecipe, services.ServiceError) {
	recipeResp := models.APIRecipe{}

	repoRecipe, err := s.RecipeRepository.Update(&recipe)
	if err != nil {
		return recipeResp, services.ServiceError{Error: err}
	}

	copier.Copy(&recipeResp, &repoRecipe)

	return recipeResp, nilErr
}

// DeleteRecipe handles the business logic when a client deletes a recipe
func (s *v1Service) DeleteRecipe(id string) services.ServiceError {
	err := s.RecipeRepository.Delete(id)
	if err != nil {
		return services.ServiceError{Error: err}
	}

	return nilErr
}

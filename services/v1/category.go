package v1

import (
	"github.com/jinzhu/copier"

	"github.com/atticuss/chefconnect/models"
	"github.com/atticuss/chefconnect/services"
)

// GetAllCategories handles the business logic when a client requests all categories
func (s *v1Service) GetAllCategories() (models.ManyAPICategories, services.ServiceError) {
	categoriesResp := models.ManyAPICategories{}
	categories, err := s.CategoryRepository.GetAll()
	if err != nil {
		return categoriesResp, services.ServiceError{Error: err}
	}

	copier.Copy(&categoriesResp, &categories)

	return categoriesResp, nilErr
}

// GetCategory handles the business logic when a client requests a specific category
func (s *v1Service) GetCategory(id string) (models.APICategory, services.ServiceError) {
	categoryResp := models.APICategory{}
	category, err := s.CategoryRepository.Get(id)
	if err != nil {
		return categoryResp, services.ServiceError{Error: err}
	}

	copier.Copy(&categoryResp, &category)

	return categoryResp, nilErr
}

// CreateCategory handles the business logic when a client creates a new category
func (s *v1Service) CreateCategory(apiCategory models.APICategory) (models.APICategory, services.ServiceError) {
	category := models.Category{}
	copier.Copy(&category, &apiCategory)

	repoCategory, err := s.CategoryRepository.Create(&category)
	copier.Copy(&apiCategory, &repoCategory)

	if err != nil {
		return apiCategory, services.ServiceError{Error: err}
	}

	return apiCategory, nilErr
}

// UpdateCategory handles the business logic when a client updates a category
func (s *v1Service) UpdateCategory(apiCategory models.APICategory) (models.APICategory, services.ServiceError) {
	category := models.Category{}
	copier.Copy(&category, &apiCategory)

	repoCategory, err := s.CategoryRepository.Update(&category)
	copier.Copy(&apiCategory, &repoCategory)

	if err != nil {
		return apiCategory, services.ServiceError{Error: err}
	}

	return apiCategory, nilErr
}

// DeleteCategory handles the business logic when a client deletes a category
func (s *v1Service) DeleteCategory(id string) services.ServiceError {
	err := s.CategoryRepository.Delete(id)
	if err != nil {
		return services.ServiceError{Error: err}
	}

	return nilErr
}

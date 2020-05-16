package services

import (
	"github.com/jinzhu/copier"

	"github.com/atticuss/chefconnect/models"
)

// GetAllCategories handles the business logic when a client requests all categories
func (ctx *ServiceCtx) GetAllCategories() (models.ManyAPICategories, ServiceError) {
	categoriesResp := models.ManyAPICategories{}
	categories, err := ctx.CategoryRepository.GetAll()
	if err != nil {
		return categoriesResp, ServiceError{Error: err}
	}

	copier.Copy(&categoriesResp, &categories)

	return categoriesResp, nilErr
}

// GetCategory handles the business logic when a client requests a specific category
func (ctx *ServiceCtx) GetCategory(id string) (models.APICategory, ServiceError) {
	categoryResp := models.APICategory{}
	category, err := ctx.CategoryRepository.Get(id)
	if err != nil {
		return categoryResp, ServiceError{Error: err}
	}

	copier.Copy(&categoryResp, &category)

	return categoryResp, nilErr
}

// CreateCategory handles the business logic when a client creates a new category
func (ctx *ServiceCtx) CreateCategory(category models.Category) (models.APICategory, ServiceError) {
	categoryResp := models.APICategory{}

	repoCategory, err := ctx.CategoryRepository.Create(&category)
	if err != nil {
		return categoryResp, ServiceError{Error: err}
	}

	copier.Copy(&categoryResp, &repoCategory)

	return categoryResp, nilErr
}

// UpdateCategory handles the business logic when a client updates a category
func (ctx *ServiceCtx) UpdateCategory(category models.Category) (models.APICategory, ServiceError) {
	categoryResp := models.APICategory{}

	repoCategory, err := ctx.CategoryRepository.Update(&category)
	if err != nil {
		return categoryResp, ServiceError{Error: err}
	}

	copier.Copy(&categoryResp, &repoCategory)

	return categoryResp, nilErr
}

// DeleteCategory handles the business logic when a client deletes a category
func (ctx *ServiceCtx) DeleteCategory(id string) ServiceError {
	err := ctx.CategoryRepository.Delete(id)
	if err != nil {
		return ServiceError{Error: err}
	}

	return nilErr
}

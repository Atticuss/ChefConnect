package services

import (
	"github.com/atticuss/chefconnect/models"
)

type Service interface {
	GetAllCategories() (models.ManyAPICategories, ServiceError)
	GetCategory(id string) (models.APICategory, ServiceError)
	CreateCategory(category models.Category) (models.APICategory, ServiceError)
	UpdateCategory(category models.Category) (models.APICategory, ServiceError)
	DeleteCategory(id string) ServiceError

	GetAllIngredients() (models.ManyAPIIngredients, ServiceError)
	GetIngredient(id string) (models.APIIngredient, ServiceError)
	CreateIngredient(ingredient models.Ingredient) (models.APIIngredient, ServiceError)
	UpdateIngredient(ingredient models.Ingredient) (models.APIIngredient, ServiceError)
	DeleteIngredient(id string) ServiceError

	GetAllRecipes() (models.ManyAPIRecipes, ServiceError)
	GetRecipe(id string) (models.APIRecipe, ServiceError)
	CreateRecipe(recipe models.Recipe) (models.APIRecipe, ServiceError)
	UpdateRecipe(recipe models.Recipe) (models.APIRecipe, ServiceError)
	DeleteRecipe(id string) ServiceError

	GetAllUsers() (models.ManyAPIUsers, ServiceError)
	GetUser(id string) (models.APIUser, ServiceError)
	CreateUser(user models.User) (models.APIUser, ServiceError)
	UpdateUser(user models.User) (models.APIUser, ServiceError)
	DeleteUser(id string) ServiceError
}

type errorCode int

// ServiceError holds both an `errors` object and an int from the enum set defined
// in the previous block.
type ServiceError struct {
	Error     error
	ErrorCode errorCode
}

// Enum set for errors that can occur within the models package. These are
// mapped back to HTTP status codes via a map within the controllers package.
const (
	Unhandled      errorCode = 0
	NotImplemented errorCode = 1
	NotFound       errorCode = 2
)

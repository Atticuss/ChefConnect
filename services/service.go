package services

import (
	"github.com/atticuss/chefconnect/models"
)

type Service interface {
	Login(authnRequest models.AuthnRequest) (models.JwtUser, ServiceError)

	ClearDatastore()
	InitializeSchema()

	GetAllCategories() (models.ManyAPICategories, ServiceError)
	GetCategory(id string) (models.APICategory, ServiceError)
	CreateCategory(category models.APICategory) (models.APICategory, ServiceError)
	UpdateCategory(category models.APICategory) (models.APICategory, ServiceError)
	DeleteCategory(id string) ServiceError

	GetAllIngredients() (models.ManyAPIIngredients, ServiceError)
	GetIngredient(id string) (models.APIIngredient, ServiceError)
	CreateIngredient(ingredient models.APIIngredient) (models.APIIngredient, ServiceError)
	UpdateIngredient(ingredient models.APIIngredient) (models.APIIngredient, ServiceError)
	DeleteIngredient(id string) ServiceError

	GetAllRecipes() (models.ManyAPIRecipes, ServiceError)
	GetRecipe(id string) (models.APIRecipe, ServiceError)
	CreateRecipe(recipe models.APIRecipe) (models.APIRecipe, ServiceError)
	UpdateRecipe(recipe models.APIRecipe) (models.APIRecipe, ServiceError)
	DeleteRecipe(id string) ServiceError

	GetAllUsers() (models.ManyAPIUsers, ServiceError)
	GetUser(id string) (models.APIUser, ServiceError)
	CreateUser(user models.APIUser) (models.APIUser, ServiceError)
	UpdateUser(user models.APIUser) (models.APIUser, ServiceError)
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
	NotAuthorized  errorCode = 3
)

package services

import (
	"github.com/atticuss/chefconnect/models"
)

// Service is an interface that all service implementations must follow
type Service interface {
	Login(authnRequest models.AuthnRequest) (models.JwtUser, ServiceError)

	ClearDatastore()
	InitializeSchema()

	GetAllTags() (models.ManyAPITags, ServiceError)
	GetTag(id string) (models.APITag, ServiceError)
	CreateTag(models.APITag) (models.APITag, ServiceError)
	UpdateTag(models.APITag) (models.APITag, ServiceError)
	DeleteTag(id string) ServiceError

	GetAllIngredients() (models.ManyAPIIngredients, ServiceError)
	GetIngredient(id string) (models.APIIngredient, ServiceError)
	CreateIngredient(models.APIIngredient) (models.APIIngredient, ServiceError)
	UpdateIngredient(models.APIIngredient) (models.APIIngredient, ServiceError)
	DeleteIngredient(id string) ServiceError

	GetAllRecipes() (models.ManyAPIRecipes, ServiceError)
	GetRecipe(id string) (models.APIRecipe, ServiceError)
	CreateRecipe(models.APIRecipe) (models.APIRecipe, ServiceError)
	UpdateRecipe(models.APIRecipe) (models.APIRecipe, ServiceError)
	SetRecipeTags(models.APIRecipe) (models.APIRecipe, ServiceError)
	DeleteRecipe(id string) ServiceError

	GetAllUsers() (models.ManyAPIUsers, ServiceError)
	GetUser(id string) (models.APIUser, ServiceError)
	CreateUser(models.APIUser) (models.APIUser, ServiceError)
	UpdateUser(models.APIUser) (models.APIUser, ServiceError)
	DeleteUser(id string) ServiceError

	GetAllRoles() (models.ManyAPIRoles, ServiceError)
	GetRole(id string) (models.APIRole, ServiceError)
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

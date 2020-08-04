package services

import (
	"github.com/atticuss/chefconnect/models"
)

// Service is an interface that all service implementations must follow
type Service interface {
	ValidateCredentials(userReq *models.User) (*models.User, *ServiceError)

	GetAllTags() (*models.ManyTags, *ServiceError)
	GetTag(id string) (*models.Tag, *ServiceError)
	CreateTag(*models.Tag) (*models.Tag, *ServiceError)
	UpdateTag(*models.Tag) (*models.Tag, *ServiceError)
	DeleteTag(id string) *ServiceError

	GetAllIngredients() (*models.ManyIngredients, *ServiceError)
	GetIngredient(id string) (*models.Ingredient, *ServiceError)
	CreateIngredient(*models.Ingredient) (*models.Ingredient, *ServiceError)
	UpdateIngredient(*models.Ingredient) (*models.Ingredient, *ServiceError)
	DeleteIngredient(id string) *ServiceError

	GetAllRecipes(*models.User) (*models.ManyRecipes, *ServiceError)
	GetRecipe(*models.User, string) (*models.Recipe, *ServiceError)
	CreateRecipe(*models.Recipe) (*models.Recipe, *ServiceError)
	DeleteRecipe(id string) *ServiceError
	UpdateRecipe(*models.Recipe) (*models.Recipe, *ServiceError)

	GetAllUsers() (*models.ManyUsers, *ServiceError)
	GetUser(id string) (*models.User, *ServiceError)
	CreateUser(*models.User) (*models.User, *ServiceError)
	UpdateUser(*models.User) (*models.User, *ServiceError)
	DeleteUser(id string) *ServiceError

	GetAllRoles() (*models.ManyRoles, *ServiceError)
	GetRole(id string) (*models.Role, *ServiceError)
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

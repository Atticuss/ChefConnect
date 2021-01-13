package services

import (
	"github.com/atticuss/chefconnect/models"
)

// Service is an interface that all service implementations must follow
type Service interface {
	GenerateJwtTokens(user *models.User) (*models.User, *ServiceError)
	ExchangeRefreshToken(refreshToken string) (*models.User, *ServiceError)
	DeserializeJwt(jwtToken string) (*models.User, *ServiceError)

	GetAllTags(callingUser *models.User) (*models.ManyTags, *ServiceError)
	GetTag(callingUser *models.User, id string) (*models.Tag, *ServiceError)
	CreateTag(callingUser *models.User, tag *models.Tag) (*models.Tag, *ServiceError)
	UpdateTag(callingUser *models.User, tag *models.Tag) (*models.Tag, *ServiceError)
	DeleteTag(callingUser *models.User, id string) *ServiceError

	GetAllIngredients(callingUser *models.User) (*models.ManyIngredients, *ServiceError)
	GetIngredient(callingUser *models.User, id string) (*models.Ingredient, *ServiceError)
	CreateIngredient(callingUser *models.User, igredient *models.Ingredient) (*models.Ingredient, *ServiceError)
	UpdateIngredient(callingUser *models.User, ingredient *models.Ingredient) (*models.Ingredient, *ServiceError)
	DeleteIngredient(callingUser *models.User, id string) *ServiceError
	IngredientSearch(callingUser *models.User, searchTerm string) (*models.ManyIngredients, *ServiceError)

	GetAllRecipes(callingUser *models.User) (*models.ManyRecipes, *ServiceError)
	GetRecipe(callingUser *models.User, id string) (*models.Recipe, *ServiceError)
	CreateRecipe(callingUser *models.User, recipe *models.Recipe) (*models.Recipe, *ServiceError)
	UpdateRecipe(callingUser *models.User, recipe *models.Recipe) (*models.Recipe, *ServiceError)
	DeleteRecipe(callingUser *models.User, id string) *ServiceError

	GetAllUsers(callingUser *models.User) (*models.ManyUsers, *ServiceError)
	GetUser(callingUser *models.User, id string) (*models.User, *ServiceError)
	CreateUser(callingUser *models.User, user *models.User) (*models.User, *ServiceError)
	UpdateUser(callingUser *models.User, user *models.User) (*models.User, *ServiceError)
	DeleteUser(callingUser *models.User, id string) *ServiceError

	GetAllRoles(callingUser *models.User) (*models.ManyRoles, *ServiceError)
	GetRole(callingUser *models.User, id string) (*models.Role, *ServiceError)
}

//type errorCode int
//type roleCode int

// ServiceError holds both an `errors` object and an int from the enum set defined
// in the previous block.
type ServiceError struct {
	Error     error
	ErrorCode int
}

// Enum set for errors that can occur within the models package. These are
// mapped back to HTTP status codes via a map within the controllers package.
const (
	Unhandled      int = 0
	NotImplemented int = 1
	NotFound       int = 2
	NotAuthorized  int = 3
	ResourceInUse  int = 4
)

const (
	Guest string = "Guest"
	Admin string = "Admin"
	User  string = "User"
)

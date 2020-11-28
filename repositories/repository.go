package repositories

import (
	"github.com/atticuss/chefconnect/models"
)

type Repository interface {
	InitializeSchema() error
	ClearDatastore() error
	InitializeBaseData() error
	InitializeTestData() error

	GetTag(id string) (*models.Tag, error)
	GetAllTags() (*models.ManyTags, error)
	CreateTag(*models.Tag) (*models.Tag, error)
	UpdateTag(*models.Tag) (*models.Tag, error)
	DeleteTag(id string) error

	GetIngredient(id string) (*models.Ingredient, error)
	GetAllIngredients() (*models.ManyIngredients, error)
	CreateIngredient(*models.Ingredient) (*models.Ingredient, error)
	UpdateIngredient(*models.Ingredient) (*models.Ingredient, error)
	DeleteIngredient(id string) error

	GetRecipe(id string) (*models.Recipe, error)
	GetAllRecipes() (*models.ManyRecipes, error)
	CreateRecipe(*models.Recipe) (*models.Recipe, error)
	UpdateRecipe(*models.Recipe) (*models.Recipe, error)
	DeleteRecipe(id string) error

	GetUser(id string) (*models.User, error)
	GetAllUsers() (*models.ManyUsers, error)
	GetUserByUsername(username string) (*models.User, error)
	GetUserByRefreshToken(refreshToken string) (*models.User, error)
	CreateUser(*models.User) (*models.User, error)
	UpdateUser(*models.User) (*models.User, error)
	DeleteUser(id string) error

	GetRole(id string) (*models.Role, error)
	GetAllRoles() (*models.ManyRoles, error)
}

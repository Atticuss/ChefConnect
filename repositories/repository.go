package repositories

import (
	"github.com/atticuss/chefconnect/models"
)

type RepositoryUtility interface {
	InitializeSchema() error
	ClearDatastore() error
	InitializeBaseData() error
}

type CategoryRepository interface {
	Get(id string) (*models.Category, error)
	GetAll() (*models.ManyCategories, error)
	Create(*models.Category) (*models.Category, error)
	Update(*models.Category) (*models.Category, error)
	Delete(id string) error
}

type IngredientRepository interface {
	Get(id string) (*models.Ingredient, error)
	GetAll() (*models.ManyIngredients, error)
	Create(*models.Ingredient) (*models.Ingredient, error)
	Update(*models.Ingredient) (*models.Ingredient, error)
	Delete(id string) error
}

type RecipeRepository interface {
	Get(id string) (*models.Recipe, error)
	GetAll() (*models.ManyRecipes, error)
	Create(*models.Recipe) (*models.Recipe, error)
	Update(*models.Recipe) (*models.Recipe, error)
	Delete(id string) error
}

type UserRepository interface {
	Get(id string) (*models.User, error)
	GetAll() (*models.ManyUsers, error)
	GetByUsername(username string) (*models.User, error)
	Create(*models.User) (*models.User, error)
	Update(*models.User) (*models.User, error)
	Delete(id string) error
}

package repositories

import (
	"github.com/atticuss/chefconnect/models"
)

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

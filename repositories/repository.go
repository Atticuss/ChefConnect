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

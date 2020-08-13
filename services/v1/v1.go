package v1

import (
	"fmt"

	"github.com/go-playground/validator/v10"

	"github.com/atticuss/chefconnect/repositories"
	"github.com/atticuss/chefconnect/services"
)

type v1Service struct {
	Validator            *validator.Validate
	TagRepository        repositories.TagRepository
	IngredientRepository repositories.IngredientRepository
	RecipeRepository     repositories.RecipeRepository
	UserRepository       repositories.UserRepository
	RoleRepository       repositories.RoleRepository
	RepositoryUtility    repositories.RepositoryUtility
}

// NewV1Service configures a service for handling business logic
func NewV1Service(
	tag *repositories.TagRepository,
	ing *repositories.IngredientRepository,
	rec *repositories.RecipeRepository,
	user *repositories.UserRepository,
	role *repositories.RoleRepository,
	util *repositories.RepositoryUtility,
) services.Service {

	v := validator.New()
	_ = v.RegisterValidation("required-update", func(fl validator.FieldLevel) bool {
		fmt.Printf("inside 'required-update' check with value: %+v\n", fl.Field())
		fmt.Printf("kind is %+v\n", fl.Field().Kind())
		fmt.Printf("len is %+v\n", len(fl.Field().String()))
		return len(fl.Field().String()) > 0
	})
	_ = v.RegisterValidation("banned-create", func(fl validator.FieldLevel) bool {
		fmt.Printf("inside 'banned-create' check with value: %+v\n", fl.Field())
		fmt.Printf("kind is %+v\n", fl.Field().Kind())
		fmt.Printf("len is %+v\n", len(fl.Field().String()))
		return len(fl.Field().String()) == 0
	})

	svc := v1Service{
		Validator:            v,
		TagRepository:        *tag,
		IngredientRepository: *ing,
		RecipeRepository:     *rec,
		UserRepository:       *user,
		RoleRepository:       *role,
		RepositoryUtility:    *util,
	}

	return &svc
}

var nilErr = services.ServiceError{Error: nil}

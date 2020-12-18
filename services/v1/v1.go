package v1

import (
	"fmt"

	"github.com/go-playground/validator/v10"

	"github.com/atticuss/chefconnect/repositories"
	"github.com/atticuss/chefconnect/services"
)

type v1Service struct {
	Validator  *validator.Validate
	Repository repositories.Repository

	// SecretKey is the string used to sign JWTs
	SecretKey string

	// TokenExpiry determines how long a JWT remains valid, in seconds
	TokenExpiry int

	// RefreshTokenLength determines how many chars are in the refresh token
	RefreshTokenLength int
}

// NewV1Service configures a service for handling business logic
func NewV1Service(repo *repositories.Repository, secretKey string, tokenExpiry int, refreshTokenLen int) services.Service {
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
		Validator:          v,
		Repository:         *repo,
		SecretKey:          secretKey,
		TokenExpiry:        tokenExpiry,
		RefreshTokenLength: refreshTokenLen,
	}

	return &svc
}

var nilErr = services.ServiceError{Error: nil}

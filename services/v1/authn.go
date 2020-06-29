package v1

import (
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/copier"

	"github.com/atticuss/chefconnect/models"
	"github.com/atticuss/chefconnect/services"
)

var hmacSecret []byte

type jwtClaims struct {
	Roles []models.NestedRole `json:"roles"`
	jwt.StandardClaims
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func compareHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// GetAllCategories handles the business logic when a client requests all categories
func (s *v1Service) Login(authReq models.AuthnRequest) (models.JwtUser, services.ServiceError) {
	jwtUser := models.JwtUser{}

	user, err := s.UserRepository.GetByUsername(authReq.Username)
	if err != nil {
		return jwtUser, services.ServiceError{Error: err, ErrorCode: services.Unhandled}
	}

	if compareHash(authReq.Password, user.Password) {
		copier.Copy(&jwtUser, &user)
		fmt.Printf("user: %+v\n", user)
		fmt.Printf("jwtuser: %+v\n", jwtUser)
		return jwtUser, nilErr
	}

	return jwtUser, services.ServiceError{
		Error:     errors.New("Invalid credentials provided"),
		ErrorCode: services.NotAuthorized,
	}
}

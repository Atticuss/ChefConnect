package v1

import (
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"github.com/atticuss/chefconnect/models"
	"github.com/atticuss/chefconnect/services"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func compareHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// ValidateCredentials handles the business logic when a client passes in authn creds
func (s *v1Service) ValidateCredentials(userReq *models.User) (*models.User, *services.ServiceError) {
	fmt.Println("inside ValidateCredentials()")
	user, err := s.Repository.GetUserByUsername(userReq.Username)
	if err != nil {
		fmt.Println("error when calling GetUserByUsername")
		fmt.Printf("%+v\n", err)
		return user, &services.ServiceError{Error: err, ErrorCode: services.Unhandled}
	}

	if compareHash(userReq.Password, user.Password) {
		fmt.Println("valid password")
		return user, &nilErr
	}

	fmt.Println("invalid password")

	return user, &services.ServiceError{
		Error:     errors.New("invalid credentials provided"),
		ErrorCode: services.NotAuthorized,
	}
}

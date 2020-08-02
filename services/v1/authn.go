package v1

import (
	"errors"

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
	user, err := s.UserRepository.GetByUsername(userReq.Username)
	if err != nil {
		return user, &services.ServiceError{Error: err, ErrorCode: services.Unhandled}
	}

	if compareHash(userReq.Password, user.Password) {
		return user, &nilErr
	}

	return user, &services.ServiceError{
		Error:     errors.New("Invalid credentials provided"),
		ErrorCode: services.NotAuthorized,
	}
}

package v1

import (
	"github.com/atticuss/chefconnect/models"
	"github.com/atticuss/chefconnect/services"
)

// GetAllUsers handles the business logic when a client requests all users
func (s *v1Service) GetAllUsers(callingUser *models.User) (*models.ManyUsers, *services.ServiceError) {
	users, err := s.UserRepository.GetAll()
	if err != nil {
		return users, &services.ServiceError{Error: err}
	}

	return users, &nilErr
}

// GetUser handles the business logic when a client requests a specific user
func (s *v1Service) GetUser(callingUser *models.User, id string) (*models.User, *services.ServiceError) {
	user, err := s.UserRepository.Get(id)
	if err != nil {
		return user, &services.ServiceError{Error: err}
	}

	return user, &nilErr
}

// CreateUser handles the business logic when a client creates a new recipe
func (s *v1Service) CreateUser(callingUser *models.User, user *models.User) (*models.User, *services.ServiceError) {
	user, err := s.UserRepository.Create(user)
	if err != nil {
		return user, &services.ServiceError{Error: err}
	}

	return user, &nilErr
}

// UpdateUser handles the business logic when a client updates a user
func (s *v1Service) UpdateUser(callingUser *models.User, user *models.User) (*models.User, *services.ServiceError) {
	user, err := s.UserRepository.Update(user)
	if err != nil {
		return user, &services.ServiceError{Error: err}
	}

	return user, &nilErr
}

// DeleteUser handles the business logic when a client deletes a recipe
func (s *v1Service) DeleteUser(callingUser *models.User, id string) *services.ServiceError {
	err := s.UserRepository.Delete(id)
	if err != nil {
		return &services.ServiceError{Error: err}
	}

	return &nilErr
}

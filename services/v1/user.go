package v1

import (
	"github.com/jinzhu/copier"

	"github.com/atticuss/chefconnect/models"
	"github.com/atticuss/chefconnect/services"
)

// GetAllUsers handles the business logic when a client requests all users
func (s *v1Service) GetAllUsers() (models.ManyAPIUsers, services.ServiceError) {
	usersResp := models.ManyAPIUsers{}
	users, err := s.UserRepository.GetAll()
	if err != nil {
		return usersResp, services.ServiceError{Error: err}
	}

	copier.Copy(&usersResp, &users)

	return usersResp, nilErr
}

// GetUser handles the business logic when a client requests a specific user
func (s *v1Service) GetUser(id string) (models.APIUser, services.ServiceError) {
	userResp := models.APIUser{}
	user, err := s.UserRepository.Get(id)
	if err != nil {
		return userResp, services.ServiceError{Error: err}
	}

	copier.Copy(&userResp, &user)

	return userResp, nilErr
}

// CreateUser handles the business logic when a client creates a new recipe
func (s *v1Service) CreateUser(apiUser models.APIUser) (models.APIUser, services.ServiceError) {
	user := models.User{}
	copier.Copy(&user, &apiUser)

	repoUser, err := s.UserRepository.Create(&user)
	copier.Copy(&apiUser, &repoUser)

	if err != nil {
		return apiUser, services.ServiceError{Error: err}
	}

	return apiUser, nilErr
}

// UpdateUser handles the business logic when a client updates a user
func (s *v1Service) UpdateUser(apiUser models.APIUser) (models.APIUser, services.ServiceError) {
	user := models.User{}
	copier.Copy(&user, &apiUser)

	repoUser, err := s.UserRepository.Update(&user)
	copier.Copy(&apiUser, &repoUser)

	if err != nil {
		return apiUser, services.ServiceError{Error: err}
	}

	return apiUser, nilErr
}

// DeleteUser handles the business logic when a client deletes a recipe
func (s *v1Service) DeleteUser(id string) services.ServiceError {
	err := s.UserRepository.Delete(id)
	if err != nil {
		return services.ServiceError{Error: err}
	}

	return nilErr
}

package v1

import (
	"github.com/atticuss/chefconnect/models"
	"github.com/atticuss/chefconnect/services"
)

// GetAllRoles handles the business logic when a client requests all roles
func (s *v1Service) GetAllRoles(callingUser *models.User) (*models.ManyRoles, *services.ServiceError) {
	roles, err := s.RoleRepository.GetAll()
	if err != nil {
		return roles, &services.ServiceError{Error: err}
	}

	return roles, &nilErr
}

// GetRole handles the business logic when a client requests a specific role
func (s *v1Service) GetRole(callingUser *models.User, id string) (*models.Role, *services.ServiceError) {
	role, err := s.RoleRepository.Get(id)
	if err != nil {
		return role, &services.ServiceError{Error: err}
	}

	return role, &nilErr
}

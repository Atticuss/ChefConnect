package v1

import (
	"errors"

	"github.com/atticuss/chefconnect/models"
	"github.com/atticuss/chefconnect/services"
)

// GetAllRoles handles the business logic when a client requests all roles
func (s *v1Service) GetAllRoles(callingUser *models.User) (*models.ManyRoles, *services.ServiceError) {
	roles := &models.ManyRoles{}

	authorized := false
	for _, role := range callingUser.Roles {
		if role.Name == services.Admin {
			authorized = true
		}
	}

	if !authorized {
		return roles, &services.ServiceError{Error: errors.New("unathorized"), ErrorCode: services.NotAuthorized}
	}

	roles, err := s.Repository.GetAllRoles()
	if err != nil {
		return roles, &services.ServiceError{Error: err}
	}

	return roles, &nilErr
}

// GetRole handles the business logic when a client requests a specific role
func (s *v1Service) GetRole(callingUser *models.User, id string) (*models.Role, *services.ServiceError) {
	role := &models.Role{}

	authorized := false
	for _, role := range callingUser.Roles {
		if role.Name == services.Admin {
			authorized = true
		}
	}

	if !authorized {
		return role, &services.ServiceError{Error: errors.New("unathorized"), ErrorCode: services.NotAuthorized}
	}

	role, err := s.Repository.GetRole(id)
	if err != nil {
		return role, &services.ServiceError{Error: err}
	}

	return role, &nilErr
}

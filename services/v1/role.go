package v1

import (
	"github.com/jinzhu/copier"

	"github.com/atticuss/chefconnect/models"
	"github.com/atticuss/chefconnect/services"
)

// GetAllRoles handles the business logic when a client requests all roles
func (s *v1Service) GetAllRoles() (models.ManyAPIRoles, services.ServiceError) {
	rolesResp := models.ManyAPIRoles{}
	roles, err := s.RoleRepository.GetAll()
	if err != nil {
		return rolesResp, services.ServiceError{Error: err}
	}

	copier.Copy(&rolesResp, &roles)

	return rolesResp, nilErr
}

// GetRole handles the business logic when a client requests a specific role
func (s *v1Service) GetRole(id string) (models.APIRole, services.ServiceError) {
	roleResp := models.APIRole{}
	role, err := s.RoleRepository.Get(id)
	if err != nil {
		return roleResp, services.ServiceError{Error: err}
	}

	copier.Copy(&roleResp, &role)

	return roleResp, nilErr
}

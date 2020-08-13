package v1

import (
	"errors"

	"github.com/atticuss/chefconnect/models"
	"github.com/atticuss/chefconnect/services"
)

// GetAllTags handles the business logic when a client requests all tags
func (s *v1Service) GetAllTags(callingUser *models.User) (*models.ManyTags, *services.ServiceError) {
	tags, err := s.TagRepository.GetAll()
	if err != nil {
		return tags, &services.ServiceError{Error: err}
	}

	return tags, &nilErr
}

// GetTag handles the business logic when a client requests a specific tag
func (s *v1Service) GetTag(callingUser *models.User, id string) (*models.Tag, *services.ServiceError) {
	tag, err := s.TagRepository.Get(id)
	if err != nil {
		return tag, &services.ServiceError{Error: err}
	}

	return tag, &nilErr
}

// CreateTag handles the business logic when a client creates a new tag
func (s *v1Service) CreateTag(callingUser *models.User, tag *models.Tag) (*models.Tag, *services.ServiceError) {
	authorized := false
	for _, role := range callingUser.Roles {
		if role.Name == services.Admin {
			authorized = true
		}
	}

	if !authorized {
		return tag, &services.ServiceError{Error: errors.New("unathorized"), ErrorCode: services.NotAuthorized}
	}

	tag, err := s.TagRepository.Create(tag)

	if err != nil {
		return tag, &services.ServiceError{Error: err}
	}

	return tag, &nilErr
}

// UpdateTag handles the business logic when a client updates a tag
func (s *v1Service) UpdateTag(callingUser *models.User, tag *models.Tag) (*models.Tag, *services.ServiceError) {
	authorized := false
	for _, role := range callingUser.Roles {
		if role.Name == services.Admin {
			authorized = true
		}
	}

	if !authorized {
		return tag, &services.ServiceError{Error: errors.New("unathorized"), ErrorCode: services.NotAuthorized}
	}

	tag, err := s.TagRepository.Update(tag)

	if err != nil {
		return tag, &services.ServiceError{Error: err}
	}

	return tag, &nilErr
}

// DeleteTag handles the business logic when a client deletes a tag
func (s *v1Service) DeleteTag(callingUser *models.User, id string) *services.ServiceError {
	authorized := false
	for _, role := range callingUser.Roles {
		if role.Name == services.Admin {
			authorized = true
		}
	}

	if !authorized {
		return &services.ServiceError{Error: errors.New("unathorized"), ErrorCode: services.NotAuthorized}
	}

	err := s.TagRepository.Delete(id)
	if err != nil {
		return &services.ServiceError{Error: err}
	}

	return &nilErr
}

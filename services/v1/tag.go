package v1

import (
	"github.com/jinzhu/copier"

	"github.com/atticuss/chefconnect/models"
	"github.com/atticuss/chefconnect/services"
)

// GetAllTags handles the business logic when a client requests all tags
func (s *v1Service) GetAllTags() (models.ManyAPITags, services.ServiceError) {
	tagsResp := models.ManyAPITags{}
	tags, err := s.TagRepository.GetAll()
	if err != nil {
		return tagsResp, services.ServiceError{Error: err}
	}

	copier.Copy(&tagsResp, &tags)

	return tagsResp, nilErr
}

// GetTag handles the business logic when a client requests a specific tag
func (s *v1Service) GetTag(id string) (models.APITag, services.ServiceError) {
	tagResp := models.APITag{}
	tag, err := s.TagRepository.Get(id)
	if err != nil {
		return tagResp, services.ServiceError{Error: err}
	}

	copier.Copy(&tagResp, &tag)

	return tagResp, nilErr
}

// CreateTag handles the business logic when a client creates a new tag
func (s *v1Service) CreateTag(apiTag models.APITag) (models.APITag, services.ServiceError) {
	tag := models.Tag{}
	copier.Copy(&tag, &apiTag)

	repoTag, err := s.TagRepository.Create(&tag)
	copier.Copy(&apiTag, &repoTag)

	if err != nil {
		return apiTag, services.ServiceError{Error: err}
	}

	return apiTag, nilErr
}

// UpdateTag handles the business logic when a client updates a tag
func (s *v1Service) UpdateTag(apiTag models.APITag) (models.APITag, services.ServiceError) {
	tag := models.Tag{}
	copier.Copy(&tag, &apiTag)

	repoTag, err := s.TagRepository.Update(&tag)
	copier.Copy(&apiTag, &repoTag)

	if err != nil {
		return apiTag, services.ServiceError{Error: err}
	}

	return apiTag, nilErr
}

// DeleteTag handles the business logic when a client deletes a tag
func (s *v1Service) DeleteTag(id string) services.ServiceError {
	err := s.TagRepository.Delete(id)
	if err != nil {
		return services.ServiceError{Error: err}
	}

	return nilErr
}

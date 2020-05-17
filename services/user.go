package services

import (
	"github.com/jinzhu/copier"

	"github.com/atticuss/chefconnect/models"
)

// GetAllUsers handles the business logic when a client requests all users
func (ctx *ServiceCtx) GetAllUsers() (models.ManyAPIUsers, ServiceError) {
	usersResp := models.ManyAPIUsers{}
	users, err := ctx.UserRepository.GetAll()
	if err != nil {
		return usersResp, ServiceError{Error: err}
	}

	copier.Copy(&usersResp, &users)

	return usersResp, nilErr
}

// GetUser handles the business logic when a client requests a specific user
func (ctx *ServiceCtx) GetUser(id string) (models.APIUser, ServiceError) {
	userResp := models.APIUser{}
	user, err := ctx.UserRepository.Get(id)
	if err != nil {
		return userResp, ServiceError{Error: err}
	}

	copier.Copy(&userResp, &user)

	return userResp, nilErr
}

// CreateUser handles the business logic when a client creates a new recipe
func (ctx *ServiceCtx) CreateUser(user models.User) (models.APIUser, ServiceError) {
	userResp := models.APIUser{}

	repoUser, err := ctx.UserRepository.Create(&user)
	if err != nil {
		return userResp, ServiceError{Error: err}
	}

	copier.Copy(&userResp, &repoUser)

	return userResp, nilErr
}

// UpdateUser handles the business logic when a client updates a user
func (ctx *ServiceCtx) UpdateUser(user models.User) (models.APIUser, ServiceError) {
	userResp := models.APIUser{}

	repoUser, err := ctx.UserRepository.Update(&user)
	if err != nil {
		return userResp, ServiceError{Error: err}
	}

	copier.Copy(&userResp, &repoUser)

	return userResp, nilErr
}

// DeleteUser handles the business logic when a client deletes a recipe
func (ctx *ServiceCtx) DeleteUser(id string) ServiceError {
	err := ctx.UserRepository.Delete(id)
	if err != nil {
		return ServiceError{Error: err}
	}

	return nilErr
}

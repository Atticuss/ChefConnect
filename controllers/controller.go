package controllers

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type Controller interface {
	ConfigureMiddleware() (*jwt.GinJWTMiddleware, error)
	Login(c *gin.Context) (interface{}, error)

	GetAllCategories(c *gin.Context)
	GetCategory(c *gin.Context)
	CreateCategory(c *gin.Context)
	UpdateCategory(c *gin.Context)
	DeleteCategory(c *gin.Context)

	GetAllIngredients(c *gin.Context)
	GetIngredient(c *gin.Context)
	CreateIngredient(c *gin.Context)
	UpdateIngredient(c *gin.Context)
	DeleteIngredient(c *gin.Context)

	GetAllRecipes(c *gin.Context)
	GetRecipe(c *gin.Context)
	CreateRecipe(c *gin.Context)
	UpdateRecipe(c *gin.Context)
	DeleteRecipe(c *gin.Context)

	GetAllUsers(c *gin.Context)
	GetUser(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}

package rest

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func (restCtlr *restController) Start() error {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if gin.IsDebugging() {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	log.Logger = log.Output(
		zerolog.ConsoleWriter{
			Out:     os.Stderr,
			NoColor: false,
		},
	)

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(restCtlr.setLogger(restCtlr.Config))

	authMiddleware, err := restCtlr.configureMiddleware()
	if err != nil {
		return errors.New("error configuring gin-jwt:" + err.Error())
	}

	router.GET("/ping", healthCheck)
	router.GET("/swagger.json", swagger)

	authRouter := router.Group("/auth")
	{
		authRouter.POST("/login", authMiddleware.LoginHandler)
		authRouter.GET("/refresh-token", authMiddleware.RefreshHandler)
	}

	ingredientRouter := router.Group("/ingredients")
	ingredientRouter.Use(authMiddleware.MiddlewareFunc())
	{
		ingredientRouter.GET("/", restCtlr.getAllIngredients)
		ingredientRouter.POST("/", restCtlr.createIngredient)
		ingredientRouter.GET("/:id", restCtlr.getIngredient)
		ingredientRouter.PUT("/:id", restCtlr.updateIngredient)
		ingredientRouter.DELETE("/:id", restCtlr.deleteIngredient)
	}

	recipeRouter := router.Group("/recipes")
	recipeRouter.Use(authMiddleware.MiddlewareFunc())
	{
		recipeRouter.GET("/", restCtlr.getAllRecipes)
		recipeRouter.POST("/", restCtlr.createRecipe)
		recipeRouter.GET("/:id", restCtlr.getRecipe)
		recipeRouter.PUT("/:id", restCtlr.updateRecipe)
		recipeRouter.DELETE("/:id", restCtlr.deleteIngredient)
	}

	userRouter := router.Group("/users")
	userRouter.Use(authMiddleware.MiddlewareFunc())
	{
		userRouter.GET("/", restCtlr.getAllUsers)
		userRouter.POST("/", restCtlr.createUser)
		userRouter.GET("/:id", restCtlr.getUser)
		userRouter.PUT("/:id", restCtlr.updateUser)
		userRouter.DELETE("/:id", restCtlr.deleteUser)
	}

	tagRouter := router.Group("/tags")
	tagRouter.Use(authMiddleware.MiddlewareFunc())
	{
		tagRouter.GET("/", restCtlr.getAllTags)
		tagRouter.POST("/", restCtlr.createTag)
		tagRouter.GET("/:id", restCtlr.getTag)
		tagRouter.PUT("/:id", restCtlr.updateTag)
		tagRouter.DELETE("/:id", restCtlr.deleteTag)
	}

	fmt.Printf("binding to port: %s\n", restCtlr.Config.Port)
	router.Run(restCtlr.Config.Port)

	return nil
}

func (restCtrl *restController) Stop() error {
	return errors.New("Not implemented")
}

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, []string{})
}

func swagger(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	jsonFile, err := os.Open("swagger.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, jsonFile)
}

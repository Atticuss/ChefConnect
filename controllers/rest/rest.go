package rest

import (
	"errors"
	"net/http"
	"os"

	"github.com/atticuss/chefconnect/controllers"
	"github.com/atticuss/chefconnect/services"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type restController struct {
	Service services.Service
	Config  Config
	Router  *gin.Engine
}

// Config defines the... configuration? I guess for the REST controller itself.
type Config struct {
	Port   string
	Logger *zerolog.Logger
	// UTC a boolean stating whether to use UTC time zone or local.
	UTC bool
}

// NewRestController configures a controller for handling request/response logic as a REST API
func NewRestController(svc *services.Service, config *Config) controllers.Controller {
	rest := restController{
		Service: *svc,
		Config:  *config,
	}

	return &rest
}

func (restCtlr *restController) SetupController() (interface{}, error) {
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
		return router, errors.New("error configuring gin-jwt:" + err.Error())
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
		recipeRouter.PUT("/:id/tags", restCtlr.setRecipeTags)
		recipeRouter.DELETE("/:id", restCtlr.deleteRecipe)
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

	roleRouter := router.Group("/roles")
	roleRouter.Use(authMiddleware.MiddlewareFunc())
	{
		roleRouter.GET("/", restCtlr.getAllRoles)
		roleRouter.GET("/:id", restCtlr.getRole)
	}

	restCtlr.Router = router

	return router, nil
}

func (restCtlr *restController) Run() error {
	err := restCtlr.Router.Run(restCtlr.Config.Port)
	return err
}

func (restCtlr *restController) Stop() error {
	return errors.New("not implemented")
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

package rest

import (
	"errors"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/atticuss/chefconnect/controllers"
	"github.com/atticuss/chefconnect/services"
)

type restController struct {
	Service   services.Service
	Config    Config
	Router    *gin.Engine
	GinRouter *ginadapter.GinLambda
}

// Config defines the... configuration? I guess for the REST controller itself.
type Config struct {
	Port                string
	Domain              string
	AuthTokenHeaderName string
	Logger              *zerolog.Logger
	IsProd              bool

	// UTC a boolean stating whether to use UTC time zone or local.
	UTC bool

	// IsLambda a boolean that toggles between AWS serverless and local deployments
	IsLambda bool
}

// NewRestController configures a controller for handling request/response logic as a REST API
func NewRestController(svc *services.Service, config *Config) controllers.Controller {
	rest := restController{
		Service: *svc,
		Config:  *config,
	}

	return &rest
}

func (restCtlr *restController) SetupController() error {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if !restCtlr.Config.IsProd {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	log.Logger = log.Output(
		zerolog.ConsoleWriter{
			Out:     os.Stderr,
			NoColor: false,
		},
	)

	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(restCtlr.setLogger(restCtlr.Config))
	router.Use(restCtlr.jwtDeserializationMiddleware())

	router.Use(corsMiddleware())
	router.GET("/ping", healthCheck)

	authRouter := router.Group("/auth")
	{
		authRouter.POST("/login", restCtlr.LoginHandler)
		authRouter.POST("/refresh-token", restCtlr.RefreshHandler)
	}

	ingredientRouter := router.Group("/ingredients")
	{
		ingredientRouter.GET("/", restCtlr.getAllIngredients)
		ingredientRouter.POST("/", restCtlr.createIngredient)
		ingredientRouter.GET("/:id", restCtlr.getIngredient)
		ingredientRouter.PUT("/:id", restCtlr.updateIngredient)
		ingredientRouter.DELETE("/:id", restCtlr.deleteIngredient)
	}

	recipeRouter := router.Group("/recipes")
	{
		recipeRouter.GET("/", restCtlr.getAllRecipes)
		recipeRouter.POST("/", restCtlr.createRecipe)
		recipeRouter.GET("/:id", restCtlr.getRecipe)
		recipeRouter.PUT("/:id", restCtlr.updateRecipe)
		recipeRouter.DELETE("/:id", restCtlr.deleteRecipe)
	}

	userRouter := router.Group("/users")
	{
		userRouter.GET("/", restCtlr.getAllUsers)
		userRouter.POST("/", restCtlr.createUser)
		userRouter.GET("/:id", restCtlr.getUser)
		userRouter.PUT("/:id", restCtlr.updateUser)
		userRouter.DELETE("/:id", restCtlr.deleteUser)
	}

	tagRouter := router.Group("/tags")
	{
		tagRouter.GET("/", restCtlr.getAllTags)
		tagRouter.POST("/", restCtlr.createTag)
		tagRouter.GET("/:id", restCtlr.getTag)
		tagRouter.PUT("/:id", restCtlr.updateTag)
		tagRouter.DELETE("/:id", restCtlr.deleteTag)
	}

	roleRouter := router.Group("/roles")
	{
		roleRouter.GET("/", restCtlr.getAllRoles)
		roleRouter.GET("/:id", restCtlr.getRole)
	}

	if restCtlr.Config.IsLambda {
		restCtlr.GinRouter = ginadapter.New(router)
	} else {
		restCtlr.Router = router
	}

	return nil
}

func (restCtlr *restController) Run() error {
	if restCtlr.Config.IsLambda {
		lambda.Start(restCtlr.handler)
		return nil
	} else {
		err := restCtlr.Router.Run(restCtlr.Config.Port)
		return err
	}
}

func (restCtlr *restController) Stop() error {
	return errors.New("not implemented")
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func (restCtlr *restController) handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return restCtlr.GinRouter.Proxy(req)
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

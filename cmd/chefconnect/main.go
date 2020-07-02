// ChefConnect API Docs
//
// The ChefConnect app is built using a modern Angular + back end API architecture. All API endpoints are detailed here. Many endpoints can be called as both an authenticated or unauthenticated user, though the data returned may differ. For example, when pulling back recipe details, the notes and ratings associated with that recipe will not be included unless authenticated.
//
//     Schemes: http
//     Host: localhost:8080
//     Version: 0.0.1
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: atticuss<jonn.callahan@gmail.com>
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package main

import (
	"log"
	"net/http"
	"os"

	"github.com/dgraph-io/dgo/v2"
	"github.com/dgraph-io/dgo/v2/protos/api"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	v1Ctlr "github.com/atticuss/chefconnect/controllers/v1"
	"github.com/atticuss/chefconnect/repositories/dgraph"
	v1Svc "github.com/atticuss/chefconnect/services/v1"
)

type app struct {
	Router *gin.Engine
}

func main() {
	a := app{}
	a.initialize("ec2-34-238-150-16.compute-1.amazonaws.com:9080")
	a.run(":8000")
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

func (a *app) initialize(dgraphURL string) {
	conn, err := grpc.Dial(dgraphURL, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	client := dgo.NewDgraphClient(api.NewDgraphClient(conn))

	categoryRepo := dgraph.NewDgraphCategoryRepository(client)
	ingredientRepo := dgraph.NewDgraphIngredientRepository(client)
	recipeRepo := dgraph.NewDgraphRecipeRepository(client)
	userRepo := dgraph.NewDgraphUserRepository(client)
	utilRepo := dgraph.NewDgraphRepositoryUtility(client)

	service := v1Svc.NewV1Service(
		&categoryRepo,
		&ingredientRepo,
		&recipeRepo,
		&userRepo,
		&utilRepo,
	)

	controller := v1Ctlr.NewV1Controller(&service)

	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	authMiddleware, err := controller.ConfigureMiddleware()
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
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
		ingredientRouter.GET("/", controller.GetAllIngredients)
		ingredientRouter.POST("/", controller.CreateIngredient)
		ingredientRouter.GET("/:id", controller.GetIngredient)
		ingredientRouter.PUT("/:id", controller.UpdateIngredient)
		ingredientRouter.DELETE("/:id", controller.DeleteIngredient)
	}

	recipeRouter := router.Group("/recipes")
	recipeRouter.Use(authMiddleware.MiddlewareFunc())
	{
		recipeRouter.GET("/", controller.GetAllRecipes)
		recipeRouter.POST("/", controller.CreateRecipe)
		recipeRouter.GET("/:id", controller.GetRecipe)
		recipeRouter.PUT("/:id", controller.UpdateRecipe)
		recipeRouter.DELETE("/:id", controller.DeleteIngredient)
	}

	userRouter := router.Group("/users")
	userRouter.Use(authMiddleware.MiddlewareFunc())
	{
		userRouter.GET("/", controller.GetAllUsers)
		userRouter.POST("/", controller.CreateUser)
		userRouter.GET("/:id", controller.GetUser)
		userRouter.PUT("/:id", controller.UpdateUser)
		userRouter.DELETE("/:id", controller.DeleteUser)
	}

	tagRouter := router.Group("/tags")
	tagRouter.Use(authMiddleware.MiddlewareFunc())
	{
		tagRouter.GET("/", controller.GetAllCategories)
		tagRouter.POST("/", controller.CreateCategory)
		tagRouter.GET("/:id", controller.GetCategory)
		tagRouter.PUT("/:id", controller.UpdateCategory)
		tagRouter.DELETE("/:id", controller.DeleteCategory)
	}

	a.Router = router
}

func (a *app) run(addr string) {
	//defer a.DgraphClient.Close()
	//handler := cors.Default().Handler(a.Router)
	a.Router.Run(addr)
	//log.Fatal(http.ListenAndServe(addr, handler))
}

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

	"github.com/atticuss/chefconnect/controllers"
	"github.com/atticuss/chefconnect/repositories/dgraph"
	v1 "github.com/atticuss/chefconnect/services/v1"
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

	service := v1.NewV1Service(
		&categoryRepo,
		&ingredientRepo,
		&recipeRepo,
		&userRepo,
		&utilRepo,
	)

	controllerCtx := controllers.ControllerCtx{
		Service: service,
	}

	router := gin.Default()
	router.POST("/login", controllerCtx.Login)
	router.GET("/ping", healthCheck)
	router.GET("/swagger.json", swagger)

	ingredientRouter := router.Group("/ingredients")
	{
		ingredientRouter.GET("/", controllerCtx.GetAllIngredients)
		ingredientRouter.POST("/", controllerCtx.CreateIngredient)
		ingredientRouter.GET("/:id", controllerCtx.GetIngredient)
		ingredientRouter.PUT("/:id", controllerCtx.UpdateIngredient)
		ingredientRouter.DELETE("/:id", controllerCtx.DeleteIngredient)
	}

	recipeRouter := router.Group("/recipes")
	{
		recipeRouter.GET("/", controllerCtx.GetAllRecipes)
		recipeRouter.POST("/", controllerCtx.CreateRecipe)
		recipeRouter.GET("/:id", controllerCtx.GetRecipe)
		recipeRouter.PUT("/:id", controllerCtx.UpdateRecipe)
		recipeRouter.DELETE("/:id", controllerCtx.DeleteIngredient)
	}

	userRouter := router.Group("/users")
	{
		userRouter.GET("/", controllerCtx.GetAllUsers)
		userRouter.POST("/", controllerCtx.CreateUser)
		userRouter.GET("/:id", controllerCtx.GetUser)
		userRouter.PUT("/:id", controllerCtx.UpdateUser)
		userRouter.DELETE("/:id", controllerCtx.DeleteUser)
	}

	tagRouter := router.Group("/tags")
	{
		tagRouter.GET("/", controllerCtx.GetAllCategories)
		tagRouter.POST("/", controllerCtx.CreateCategory)
		tagRouter.GET("/:id", controllerCtx.GetCategory)
		tagRouter.PUT("/:id", controllerCtx.UpdateCategory)
		tagRouter.DELETE("/:id", controllerCtx.DeleteCategory)
	}

	/*
		router.HandleFunc("/ping", healthCheck).Methods("GET")
		router.HandleFunc("/swagger.json", swagger).Methods("GET")
	*/

	a.Router = router
}

func (a *app) run(addr string) {
	//defer a.DgraphClient.Close()
	//handler := cors.Default().Handler(a.Router)
	a.Router.Run(addr)
	//log.Fatal(http.ListenAndServe(addr, handler))
}

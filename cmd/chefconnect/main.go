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
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/dgraph-io/dgo/v2"
	"github.com/dgraph-io/dgo/v2/protos/api"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	"github.com/atticuss/chefconnect/controllers"
	"github.com/atticuss/chefconnect/models"
	"github.com/atticuss/chefconnect/repositories/dgraph"
	v1 "github.com/atticuss/chefconnect/services/v1"
)

type app struct {
	Router *gin.Engine
}

// shamelessly stolen from: https://gist.github.com/dopey/c69559607800d2f2f90b1b1ed4e550fb
func assertAvailablePRNG() error {
	// Assert that a cryptographically secure PRNG is available.
	// Panic otherwise.
	buf := make([]byte, 1)

	if _, err := io.ReadFull(rand.Reader, buf); err != nil {
		return err
	}

	return nil
}

func generateRandomBytes(n int) ([]byte, error) {
	if err := assertAvailablePRNG(); err != nil {
		return nil, fmt.Errorf("crypto/rand is unavailable: Read() failed with %#v", err)
	}

	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
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
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	var identityKey = "uid"
	secretKey, err := generateRandomBytes(100)
	if err != nil {
		log.Fatal(err)
	}

	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "chefconnect",
		Key:         secretKey,
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(models.JwtUser); ok {
				// this logic is for converting the JwtUser struct to a map[string]interface{}
				// https://stackoverflow.com/a/42849112/13203635
				var claims jwt.MapClaims
				inrec, _ := json.Marshal(v)
				json.Unmarshal(inrec, &claims)

				return claims
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			fmt.Printf("claims: %+v\n", claims)
			jwtUser := &models.JwtUser{
				Username: claims[identityKey].(string),
			}
			fmt.Printf("jwtUser: %+v\n", jwtUser)
			return jwtUser
		},
		Authenticator: controllerCtx.Login,
		Authorizator: func(data interface{}, c *gin.Context) bool {
			return true //push authorization off to the services layer
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup:   "header: Authorization, cookie: jwt",
		TokenHeadName: "Token",
		TimeFunc:      time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	router.POST("/login", authMiddleware.LoginHandler)
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

	a.Router = router
}

func (a *app) run(addr string) {
	//defer a.DgraphClient.Close()
	//handler := cors.Default().Handler(a.Router)
	a.Router.Run(addr)
	//log.Fatal(http.ListenAndServe(addr, handler))
}

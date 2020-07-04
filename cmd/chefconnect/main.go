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
	"os"

	"github.com/dgraph-io/dgo/v2"
	"github.com/dgraph-io/dgo/v2/protos/api"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"

	"github.com/atticuss/chefconnect/controllers/rest"
	"github.com/atticuss/chefconnect/repositories/dgraph"
	v1 "github.com/atticuss/chefconnect/services/v1"
)

type app struct {
	Router *gin.Engine
}

func main() {
	conn, _ := grpc.Dial("ec2-34-238-150-16.compute-1.amazonaws.com:9080", grpc.WithInsecure())

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

	log.Logger = log.Output(
		zerolog.ConsoleWriter{
			Out:     os.Stderr,
			NoColor: false,
		},
	)

	subLog := zerolog.New(os.Stdout).With().Logger()
	config := rest.Config{
		Port:   ":8000",
		Logger: &subLog,
		UTC:    true,
	}

	controller := rest.NewRestController(&service, &config)
	controller.Start()
}

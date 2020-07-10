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

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/atticuss/chefconnect/controllers/rest"
	"github.com/atticuss/chefconnect/repositories/dgraph"
	v1 "github.com/atticuss/chefconnect/services/v1"
)

func main() {
	log.Logger = log.Output(
		zerolog.ConsoleWriter{
			Out:     os.Stderr,
			NoColor: false,
		},
	)

	subLog := zerolog.New(os.Stdout).With().Logger()
	restConfig := rest.Config{
		Port:   ":8000",
		Logger: &subLog,
		UTC:    true,
	}

	dgraphConfig := dgraph.Config{
		Host: "ec2-34-238-150-16.compute-1.amazonaws.com:9080",
	}

	tagRepo := dgraph.NewDgraphTagRepository(&dgraphConfig)
	ingredientRepo := dgraph.NewDgraphIngredientRepository(&dgraphConfig)
	recipeRepo := dgraph.NewDgraphRecipeRepository(&dgraphConfig)
	userRepo := dgraph.NewDgraphUserRepository(&dgraphConfig)
	utilRepo := dgraph.NewDgraphRepositoryUtility(&dgraphConfig)

	service := v1.NewV1Service(
		&tagRepo,
		&ingredientRepo,
		&recipeRepo,
		&userRepo,
		&utilRepo,
	)

	controller := rest.NewRestController(&service, &restConfig)
	controller.SetupController()
	controller.Run()
}

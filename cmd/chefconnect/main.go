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
	"fmt"
	"log"
	"net/http"

	"github.com/dgraph-io/dgo/v2"
	"github.com/dgraph-io/dgo/v2/protos/api"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"google.golang.org/grpc"

	"github.com/atticuss/chefconnect/controllers"
	"github.com/atticuss/chefconnect/repositories/dgraph"
	"github.com/atticuss/chefconnect/services"
)

type app struct {
	Router *mux.Router
}

func main() {
	a := app{}
	a.initialize("ec2-34-238-150-16.compute-1.amazonaws.com:9080")
	a.run(":8000")
}

func (a *app) initialize(dgraphURL string) {
	controllerCtx := controllers.ControllerCtx{}

	conn, err := grpc.Dial(dgraphURL, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	controllerCtx.DgraphClient = dgo.NewDgraphClient(api.NewDgraphClient(conn))

	// no longer need func-specific validation. will leave for now in case i end up
	// needing it for other controllers down the line.
	v := validator.New()
	_ = v.RegisterValidation("required-update", func(fl validator.FieldLevel) bool {
		fmt.Printf("inside 'required-update' check with value: %+v\n", fl.Field())
		fmt.Printf("kind is %+v\n", fl.Field().Kind())
		fmt.Printf("len is %+v\n", len(fl.Field().String()))
		return len(fl.Field().String()) > 0
	})
	_ = v.RegisterValidation("banned-create", func(fl validator.FieldLevel) bool {
		fmt.Printf("inside 'banned-create' check with value: %+v\n", fl.Field())
		fmt.Printf("kind is %+v\n", fl.Field().Kind())
		fmt.Printf("len is %+v\n", len(fl.Field().String()))
		return len(fl.Field().String()) == 0
	})

	controllerCtx.Validator = v

	client := dgo.NewDgraphClient(api.NewDgraphClient(conn))
	categoryRepo := dgraph.NewDgraphCategoryRepository(client)

	serviceCtx := services.ServiceCtx{
		Validator:          v,
		CategoryRepository: categoryRepo,
	}

	controllerCtx.ServiceCtx = &serviceCtx

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/ingredients", controllerCtx.GetAllIngredients).Methods("GET")
	router.HandleFunc("/ingredients", controllerCtx.CreateIngredient).Methods("POST")
	router.HandleFunc("/ingredients/{id}", controllerCtx.GetIngredient).Methods("GET")
	router.HandleFunc("/ingredients/{id}", controllerCtx.UpdateIngredient).Methods("PUT")
	router.HandleFunc("/ingredients/{id}", controllerCtx.DeleteIngredient).Methods("DELETE")

	router.HandleFunc("/recipes", controllerCtx.GetAllRecipes).Methods("GET")
	router.HandleFunc("/recipes", controllerCtx.CreateRecipe).Methods("POST")
	router.HandleFunc("/recipes/{id}", controllerCtx.GetRecipe).Methods("GET")
	router.HandleFunc("/recipes/{id}", controllerCtx.UpdateRecipe).Methods("PUT")
	router.HandleFunc("/recipes/{id}", controllerCtx.DeleteRecipe).Methods("DELETE")

	router.HandleFunc("/users", controllerCtx.GetAllUsers).Methods("GET")
	router.HandleFunc("/users", controllerCtx.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", controllerCtx.GetUser).Methods("GET")
	router.HandleFunc("/users/{id}", controllerCtx.UpdateIngredient).Methods("PUT")
	router.HandleFunc("/users/{id}", controllerCtx.DeleteUser).Methods("DELETE")

	router.HandleFunc("/categories", controllerCtx.GetAllCategories).Methods("GET")
	router.HandleFunc("/categories", controllerCtx.CreateCategory).Methods("POST")
	router.HandleFunc("/categories/{id}", controllerCtx.GetCategory).Methods("GET")
	router.HandleFunc("/categories/{id}", controllerCtx.UpdateCategory).Methods("PUT")
	router.HandleFunc("/categories/{id}", controllerCtx.DeleteCategory).Methods("DELETE")

	router.HandleFunc("/ping", healthCheck).Methods("GET")
	router.HandleFunc("/swagger.json", swagger).Methods("GET")

	a.Router = router
}

func (a *app) run(addr string) {
	//defer a.DgraphClient.Close()
	handler := cors.Default().Handler(a.Router)
	log.Fatal(http.ListenAndServe(addr, handler))
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func swagger(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	http.ServeFile(w, r, "swagger.json")
}

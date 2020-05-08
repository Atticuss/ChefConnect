// Testing go-swagger generation
//
// The purpose of this application is to test go-swagger in a simple GET request.
//
//     Schemes: http
//     Host: localhost:8080
//     Version: 0.0.1
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: atticuss<jonn.callahan@gmail.com>
//
//     Consumes:
//     - text/plain
//
//     Produces:
//     - text/plain
//
// swagger:meta
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/atticuss/chefconnect/controllers"
	"github.com/dgraph-io/dgo/v2"
	"github.com/dgraph-io/dgo/v2/protos/api"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"google.golang.org/grpc"
)

type app struct {
	Router *mux.Router
	Ctx    *controllers.ControllerCtx
}

func main() {
	a := app{}
	a.initialize("ec2-34-238-150-16.compute-1.amazonaws.com:9080")
	a.run(":8010")
}

func (a *app) initialize(dgraphURL string) {
	ctx := controllers.ControllerCtx{}

	conn, err := grpc.Dial(dgraphURL, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	ctx.DgraphClient = dgo.NewDgraphClient(api.NewDgraphClient(conn))

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

	ctx.Validator = v
	a.Ctx = &ctx

	router := mux.NewRouter().StrictSlash(true)
	//this errors out
	router.HandleFunc("/ingredients", ctx.GetAllIngredients).Methods("GET")
	router.HandleFunc("/ingredients", ctx.CreateIngredient).Methods("POST")
	router.HandleFunc("/ingredients/{id}", ctx.GetIngredient).Methods("GET")
	router.HandleFunc("/ingredients/{id}", ctx.DeleteIngredient).Methods("DELETE")

	router.HandleFunc("/recipes/{id}", ctx.GetRecipe).Methods("GET")

	router.HandleFunc("/users", ctx.GetAllUsers).Methods("GET")
	router.HandleFunc("/users/{id}", ctx.GetUser).Methods("GET")

	//this errors out
	router.HandleFunc("/categories", ctx.GetAllCategories).Methods("GET")
	router.HandleFunc("/categories/{id}", ctx.GetCategory).Methods("GET")

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

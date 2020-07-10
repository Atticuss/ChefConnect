package integration_rest

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/atticuss/chefconnect/controllers/rest"
	"github.com/atticuss/chefconnect/repositories/dgraph"
	v1 "github.com/atticuss/chefconnect/services/v1"
)

var (
	router *gin.Engine
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

func setup() {
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
	r, _ := controller.SetupController()
	router, _ = r.(*gin.Engine)
}

func performRequest(method, path string, session *http.Cookie, body []byte) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, bytes.NewBuffer(body))
	req.Header.Add("Content-Type", "application/json")

	if session != nil {
		req.AddCookie(session)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	return w
}

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
	"os"
	"strings"

	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"

	"github.com/atticuss/chefconnect/controllers/rest"
	"github.com/atticuss/chefconnect/repositories/dgraph"
	v1 "github.com/atticuss/chefconnect/services/v1"
)

type configuration struct {
	Database struct {
		Host      string `envconfig:"DB_HOST"`
		Port      string `envconfig:"DB_PORT"`
		AuthToken string `envconfig:"DB_TOKEN"`
	}
	Server struct {
		Domain   string `envconfig:"DOMAIN"`
		Port     string `envconfig:"SERVER_PORT"`
		IsLambda bool   `envconfig:"IS_LAMBDA"`

		SecretKey           string `envconfig:"SECRET_KEY"`
		TokenExpiry         int    `envconfig:"TOKEN_EXPIRY"`
		RefreshTokenLength  int    `envconfig:"REFRESH_TOKEN_LEN"`
		AuthTokenHeaderName string `envconfig:"AUTH_TOKEN_HEADER_NAME"`
	}

	Environment string `envconfig:"ENVIRONMENT"`
}

func parseConfig() *configuration {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	var config configuration

	if err := viper.ReadInConfig(); err == nil {
		viper.Unmarshal(&config)
	}

	envconfig.Process("", &config)
	return &config
}

func main() {
	log.Logger = log.Output(
		zerolog.ConsoleWriter{
			Out:     os.Stderr,
			NoColor: false,
		},
	)

	config := parseConfig()

	subLog := zerolog.New(os.Stdout).With().Logger()
	restConfig := rest.Config{
		Port:                config.Server.Port,
		Logger:              &subLog,
		UTC:                 true,
		IsLambda:            config.Server.IsLambda,
		AuthTokenHeaderName: config.Server.AuthTokenHeaderName,
	}

	env := strings.ToLower(config.Environment)
	if env == "prod" || env == "production" {
		restConfig.IsProd = true
	}

	dgraphConfig := dgraph.Config{
		Host:      fmt.Sprintf("%s:%s", config.Database.Host, config.Database.Port),
		AuthToken: config.Database.AuthToken,
	}

	dgraphRepo := dgraph.NewDgraphRepository(&dgraphConfig)
	service := v1.NewV1Service(&dgraphRepo, config.Server.SecretKey, config.Server.TokenExpiry, config.Server.RefreshTokenLength)
	controller := rest.NewRestController(&service, &restConfig)
	if err := controller.SetupController(); err != nil {
		log.Fatal().Msg(err.Error())
	}

	if err := controller.Run(); err != nil {
		log.Fatal().Msg(err.Error())
	}
}

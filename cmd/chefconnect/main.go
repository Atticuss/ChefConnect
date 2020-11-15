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
	}

	Environment string `envconfig:"ENVIRONMENT"`
}

func parseConfig() (*configuration, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	var config configuration

	if err := viper.ReadInConfig(); err != nil {
		return &config, err
	}
	err := viper.Unmarshal(&config)
	if err != nil {
		return &config, err
	}

	err = envconfig.Process("", &config)
	if err != nil {
		return &config, err
	}

	return &config, nil
}

func main() {
	log.Logger = log.Output(
		zerolog.ConsoleWriter{
			Out:     os.Stderr,
			NoColor: false,
		},
	)

	config, err := parseConfig()
	fmt.Printf("config: %+v\n", config)
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	subLog := zerolog.New(os.Stdout).With().Logger()
	restConfig := rest.Config{
		Port:     config.Server.Port,
		Logger:   &subLog,
		UTC:      true,
		IsLambda: config.Server.IsLambda,
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
	service := v1.NewV1Service(&dgraphRepo)
	controller := rest.NewRestController(&service, &restConfig)
	if err := controller.SetupController(); err != nil {
		log.Fatal().Msg(err.Error())
	}

	if err := controller.Run(); err != nil {
		log.Fatal().Msg(err.Error())
	}
}

package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"

	"github.com/atticuss/chefconnect/controllers/liveview"
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
		Domain string `envconfig:"DOMAIN"`
		Port   string `envconfig:"SERVER_PORT"`

		SecretKey          string `envconfig:"SECRET_KEY"`
		TokenExpiry        int    `envconfig:"TOKEN_EXPIRY"`
		RefreshTokenLength int    `envconfig:"REFRESH_TOKEN_LEN"`
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
	lvConfig := liveview.Config{
		Port:   config.Server.Port,
		Logger: &subLog,
		UTC:    true,
	}

	env := strings.ToLower(config.Environment)
	if env == "prod" || env == "production" {
		lvConfig.IsProd = true
	}

	dgraphConfig := dgraph.Config{
		Host:      fmt.Sprintf("%s:%s", config.Database.Host, config.Database.Port),
		AuthToken: config.Database.AuthToken,
	}

	dgraphRepo := dgraph.NewDgraphRepository(&dgraphConfig)
	service := v1.NewV1Service(&dgraphRepo, config.Server.SecretKey, config.Server.TokenExpiry, config.Server.RefreshTokenLength)
	controller := liveview.NewLiveViewController(&service, &lvConfig)
	if err := controller.SetupController(); err != nil {
		log.Fatal().Msg(err.Error())
	}

	if err := controller.Run(); err != nil {
		log.Fatal().Msg(err.Error())
	}
}

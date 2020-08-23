package main

import (
	"fmt"
	"log"

	"github.com/atticuss/chefconnect/repositories/dgraph"
	"github.com/kelseyhightower/envconfig"
	"github.com/spf13/viper"
)

type configuration struct {
	Database struct {
		Host string `envconfig:"DB_HOST"`
		Port string `envconfig:"DB_PORT"`
	}
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
	config, err := parseConfig()
	if err != nil {
		log.Fatal(err.Error())
	}

	dgraphConfig := dgraph.Config{
		Host: fmt.Sprintf("%s:%s", config.Database.Host, config.Database.Port),
	}
	repo := dgraph.NewDgraphRepository(&dgraphConfig)

	if err := repo.ClearDatastore(); err != nil {
		log.Fatal(err)
	}

	if err := repo.InitializeSchema(); err != nil {
		log.Fatal(err)
	}

	if err := repo.InitializeBaseData(); err != nil {
		log.Fatal(err)
	}

	if err := repo.InitializeTestData(); err != nil {
		log.Fatal(err)
	}
}

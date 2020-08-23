package main

import (
	"fmt"
	"log"
	"runtime"
	"strings"

	"github.com/kelseyhightower/envconfig"
	"github.com/spf13/viper"
)

type databaseConfig struct {
	Host string `envconfig:"DB_HOST"`
	Port int    `envconfig:"DB_PORT"`
}

type serverConfig struct {
	Port     int  `envconfig:"SERVER_PORT"`
	IsLambda bool `envconfig:"IS_LAMBDA"`
}

type configuration struct {
	Database databaseConfig
	Server   serverConfig
}

// slight modification of: https://stackoverflow.com/a/35213181/13203635
func getFrame() (string, error) {
	targetFrameIndex := 4

	// Set size to targetFrameIndex+2 to ensure we have room for one more caller than we need
	programCounters := make([]uintptr, targetFrameIndex+2)
	n := runtime.Callers(0, programCounters)

	frame := runtime.Frame{Function: "unknown"}
	if n > 0 {
		frames := runtime.CallersFrames(programCounters[:n])
		for more, frameIndex := true, 0; more && frameIndex <= targetFrameIndex; frameIndex++ {
			var frameCandidate runtime.Frame
			frameCandidate, more = frames.Next()
			if frameIndex == targetFrameIndex {
				frame = frameCandidate
			}
		}
	}

	funcNames := strings.Split(frame.Function, ".")
	if len(funcNames) != 2 {
		return "", fmt.Errorf("unexpected function name: %s", frame.Function)
	}

	return funcNames[1], nil
}

func foo() {
	callerFunc, err := getFrame()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("i was called by: %s\n", callerFunc)
}

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	var config configuration

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	log.Printf("config: %+v\n", config)

	err = envconfig.Process("", &config)
	if err != nil {
		log.Fatalf("unable to read env:, %v", err)
	}
	log.Printf("config: %+v\n", config)
}

package config

import (
	"os"
	"strings"

	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/source/cli"
	"github.com/micro/go-micro/config/source/env"
	"github.com/micro/go-micro/config/source/file"
	"github.com/micro/go-micro/util/log"
	// "github.com/micro/go-plugins/config/source/configmap"
)

var (
	// IsProduction will have CurrentMode of the application
	IsProduction bool
)

// init tries to load and map the ServiceConfiguration struct
// the sources are sequentially loaded: config-file, config-map, environment-variables
func init() {
	if _, found := os.LookupEnv("APP_ENV"); found {
		IsProduction = true
	}

	InitConfig()
}

// InitConfig loads the configuration from file then from environment variables and then from cli flags
func InitConfig() {
	if err := config.Load(
		// base config from file. Default: config.json
		file.NewSource(file.WithPath("config.yaml")),
		// override file from configmap
		// configmap.NewSource(),
		// override configmap from env
		env.NewSource(),
		// override env with cli flags
		cli.NewSource(),
	); err != nil {
		if strings.Contains(err.Error(), "no such file") {
			log.Log("missing config.yaml, use environment variables")
		} else {
			log.Fatal(err.Error())
		}
	}
}

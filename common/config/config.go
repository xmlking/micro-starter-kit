package config

import (
	"strings"

	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/source/env"
	"github.com/micro/go-micro/config/source/file"
	"github.com/micro/go-micro/config/source/flag"
	"github.com/micro/go-micro/util/log"
	// "github.com/micro/go-plugins/config/source/configmap"
)

// init tries to load and map the ServiceConfiguration struct
// the sources are sequentially loaded: config-file, environment-variables
func init() {
	initConfig()
}

// InitConfig loads the configuration from file and then from environment variables
func initConfig() {
	if err := config.Load(
		// base config from file. Default: config.json
		file.NewSource(file.WithPath("config.yaml")),
		// override file from configmap
		// configmap.NewSource(),
		// override configmap from env
		env.NewSource(),
	); err != nil {
		if strings.Contains(err.Error(), "no such file") {
			log.Log("missing config.yaml, use environment variables")
		} else {
			log.Fatal(err.Error())
		}
	}
}

// InitConfigWithFlag loads the configuration from file, then from environment variables and then from flags
func InitConfigWithFlag() {
	if err := config.Load(
		// base config from file. Default: config.json
		file.NewSource(file.WithPath("config.yaml")),
		// override file from configmap
		// configmap.NewSource(),
		// override configmap from env
		env.NewSource(),
		// override env with flags
		flag.NewSource(flag.IncludeUnset(true)),
	); err != nil {
		if strings.Contains(err.Error(), "no such file") {
			log.Log("missing config.yaml, use environment variables")
		} else {
			log.Fatal(err.Error())
		}
	}
}

package config

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	microConfig "github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/source/cli"
	"github.com/micro/go-micro/config/source/env"
	"github.com/micro/go-micro/config/source/file"
	log "github.com/sirupsen/logrus"
	// "github.com/micro/go-plugins/config/source/configmap"
)

var (
	// IsProduction will have CurrentMode of the application
	IsProduction bool

	// Version is populated by govvv in compile-time.
	Version = "untouched"
	// BuildDate is populated by govvv.
	BuildDate string
	// GitCommit is populated by govvv.
	GitCommit string
	// GitBranch is populated by govvv.
	GitBranch string
	// GitState is populated by govvv.
	GitState string
	// GitSummary is populated by govvv.
	GitSummary string
)

// VersionMsg is the message that is shown after process started.
const versionMsg = `version     : %s
build date  : %s
go version  : %s
go compiler : %s
platform    : %s/%s
git commit  : %s
git branch  : %s
git state   : %s
git summary : %s
`

const (
	// DefaultConfigDir if o ConfigDir supplied
	DefaultConfigDir = "config"
	// DefaultConfigFile if o ConfigFile supplied
	DefaultConfigFile = "config.yaml"
)

// PrintBuildInfo print build info
func PrintBuildInfo() {
	log.Info(GetBuildInfo())
}

// GetBuildInfo get build info
func GetBuildInfo() string {
	return fmt.Sprintf(versionMsg, Version, BuildDate, runtime.Version(), runtime.Compiler, runtime.GOOS, runtime.GOARCH,
		GitCommit, GitBranch, GitState, GitSummary)
}

// init tries to load and map the ServiceConfiguration struct
// the sources are sequentially loaded: config-file, config-map, environment-variables
func init() {
	if _, found := os.LookupEnv("APP_ENV"); found {
		IsProduction = true
	}

	InitConfig("", "")
}

// InitConfig loads the configuration from file then from environment variables and then from cli flags
func InitConfig(configDir, configFile string) {
	if configDir == "" {
		configDir = DefaultConfigDir
	}
	if configFile == "" {
		configFile = DefaultConfigFile
	}
	configPath := filepath.Join(configDir, configFile)
	log.Infof("loading configuration from file: %s", configPath)

	if err := microConfig.Load(
		// base config from file. Default: config/config.yaml
		file.NewSource(file.WithPath(configPath)),
		// override file from configmap
		// configmap.NewSource(),
		// override configmap from env
		env.NewSource(),
		// override env with cli flags
		cli.NewSource(),
	); err != nil {
		if strings.Contains(err.Error(), "no such file") {
			log.Errorf(`missing config file at %s, fallback to default config path.
            you can set config path via: --configDir=path/to/my/configDir --configFile=config.yaml`, configPath)
		} else {
			log.Fatal(err.Error())
		}
	}
}

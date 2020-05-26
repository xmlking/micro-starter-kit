package config

import (
    "context"
    "crypto/tls"
    "fmt"
    "os"
    "path/filepath"
    "runtime"
    "strings"

    microConfig "github.com/micro/go-micro/v2/config"
    "github.com/micro/go-micro/v2/config/source/cli"
    "github.com/micro/go-micro/v2/config/source/env"
    "github.com/micro/go-plugins/config/source/pkger/v2"
    "github.com/rs/zerolog/log"

    "github.com/xmlking/micro-starter-kit/shared/util"
)

var (
	// Default Config
	DefaultConfig Config = NewConfig()
)

var (
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
const versionMsg = `
version     : %s
build date  : %s
go version  : %s
go compiler : %s
platform    : %s/%s
git commit  : %s
git branch  : %s
git state   : %s
git summary : %s
`

type Config interface {
	Init(options ...Option) error
	Options() Options
	GetBuildInfo() string
	GetServiceConfig() ServiceConfiguration
	CreateServerCerts() (*tls.Config, error)
	GetFeatureFlags() FeatureFlags
	IsProduction() bool
	String() string
}

type defaultConfig struct {
	opts         Options
	cfg          ServiceConfiguration
	isProduction bool
	ff           FeatureFlags
}

// Re-init means, loading extra configuration into MicroConfig, when new files are provided.
func (c *defaultConfig) Init(opts ...Option) (err error) {
	for _, o := range opts {
		o(&c.opts)
	}

	configPath := filepath.Join(c.opts.ConfigDir, c.opts.ConfigFile)
	log.Info().Msgf("loading configuration from file: %s", configPath)

	err = microConfig.Load(
		// base config from file. Default: config/config.yaml
		pkger.NewSource(pkger.WithPath(configPath)),
		// override file from configmap
		// configmap.NewSource(),
		// override configmap from env
		env.NewSource(),
		// override env with cli flags
		cli.NewSource(),
	)
	if err != nil {
		if strings.Contains(err.Error(), "no such file") {
			log.Error().Err(err).Msgf("missing config file at %s", configPath)
		} else {
			log.Fatal().Err(err).Msg("")
		}
		return
	}

	err = microConfig.Scan(&c.cfg)
	if err != nil {
		return
	}

    if c.cfg.Environment == "production" {
        c.isProduction = true
    } else if _, found := os.LookupEnv("APP_ENV"); found {
		c.isProduction = true
	}

	c.ff = &featureFlags{features: c.cfg.Features}
	return
}

func (c *defaultConfig) GetBuildInfo() string {
	return fmt.Sprintf(versionMsg, Version, BuildDate, runtime.Version(), runtime.Compiler, runtime.GOOS, runtime.GOARCH,
		GitCommit, GitBranch, GitState, GitSummary)
}

func (c *defaultConfig) CreateServerCerts() (tlsConf *tls.Config, err error) {
	tlsConf, err = util.GetTLSConfig(
		filepath.Join(c.opts.ConfigDir, microConfig.Get("features", "mtls", "certfile").String("")),
		filepath.Join(c.opts.ConfigDir, microConfig.Get("features", "mtls", "keyfile").String("")),
		filepath.Join(c.opts.ConfigDir, microConfig.Get("features", "mtls", "cafile").String("")),
		filepath.Join(c.opts.ConfigDir, microConfig.Get("features", "mtls", "servername").String("")),
	)
	if err != nil {
		tlsConf, err = util.GetSelfSignedTLSConfig("*")
	}
	return
}

func (c *defaultConfig) GetFeatureFlags() FeatureFlags {
	return c.ff
}

func (c *defaultConfig) IsProduction() bool {
	return c.isProduction
}

func (c *defaultConfig) GetServiceConfig() (cfg ServiceConfiguration) {
	return c.cfg
}

func (c *defaultConfig) Options() Options {
	return c.opts
}

func (c *defaultConfig) String() string {
	return "default"
}

func NewConfig(opts ...Option) Config {
	// Default options
	options := Options{
		ConfigDir:  "/config",
		ConfigFile: "config.yaml",
		Context:    context.Background(),
	}

	c := &defaultConfig{opts: options}
	_ = c.Init(opts...)
	return c
}

// Helper functions on DefaultConfig
func Init(options ...Option) error {
	return DefaultConfig.Init(options...)
}
func IsProduction() bool {
	return DefaultConfig.IsProduction()
}

func GetBuildInfo() string {
	return DefaultConfig.GetBuildInfo()
}
func GetServiceConfig() ServiceConfiguration {
	return DefaultConfig.GetServiceConfig()
}
func CreateServerCerts() (*tls.Config, error) {
	return DefaultConfig.CreateServerCerts()
}
func GetFeatureFlags() FeatureFlags {
	return DefaultConfig.GetFeatureFlags()
}



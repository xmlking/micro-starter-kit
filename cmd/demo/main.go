package main

import (
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/service/grpc"
	log "github.com/sirupsen/logrus"

	myConfig "github.com/xmlking/micro-starter-kit/shared/config"
	logger "github.com/xmlking/micro-starter-kit/shared/log"
	logWrapper "github.com/xmlking/micro-starter-kit/shared/wrapper/log"
)

const (
	serviceName = "account-cmd"
)

var (
	configDir  string
	configFile string
	cfg        myConfig.ServiceConfiguration
)

// init is called on package initialization and can therefore be used to initialize global stuff like logging, config, ..
func init() {
	config.Scan(&cfg)
}

func main() {
	// New Service
	service := grpc.NewService(
		// optional cli flag to override config.
		// comment out if you don't need to override any base config via CLI
		micro.Flags(
			cli.StringFlag{
				Name:        "configDir, d",
				Value:       "config",
				Usage:       "Path to the config directory. Defaults to 'config'",
				EnvVar:      "CONFIG_DIR",
				Destination: &configDir,
			},
			cli.StringFlag{
				Name:        "configFile, c",
				Value:       "config.yaml",
				Usage:       "Config file in configDir. Defaults to 'config.yaml'",
				EnvVar:      "CONFIG_FILE",
				Destination: &configFile,
			},
			cli.StringFlag{
				Name:   "database_host, dh",
				Value:  "127.0.0.1",
				Usage:  "Database hostname. Defaults to 127.0.0.1",
				EnvVar: "DATABASE_HOST",
			},
			cli.IntFlag{
				Name:   "database_port, dp",
				Value:  5432,
				Usage:  "Database port. Defaults to 5432",
				EnvVar: "DATABASE_PORT",
			},
		),
		micro.Name(serviceName),
		micro.Version(myConfig.Version),
		micro.WrapHandler(logWrapper.NewHandlerWrapper()),
	)

	// Initialize service
	service.Init(
		micro.Action(func(c *cli.Context) {
			// load config
			myConfig.InitConfig(configDir, configFile)
			config.Scan(&cfg)
			logger.InitLogger(cfg.Log)
		}),
	)

	log.Debugf("IsProduction? %v", myConfig.IsProduction)
	log.Debugf("environment: %v", cfg.Environment)
	log.Debug(config.Get("database", "dialect").String("postgres"))
	log.Debug(config.Get("database", "host").String("no-address"))
	log.Debug(config.Get("database", "port").Int(0000))
	log.Debug(config.Get("observability", "tracing", "flushInterval").Int(2000000000))
	log.Debug(cfg)
	log.Debugf("cfg is %v", cfg)
	log.Debug(configDir)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

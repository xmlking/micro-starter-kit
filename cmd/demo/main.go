package main

import (
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/util/log"

	myConfig "github.com/xmlking/micro-starter-kit/shared/config"
	_ "github.com/xmlking/micro-starter-kit/shared/log"
	"github.com/xmlking/micro-starter-kit/shared/wrapper"
)

var (
	configurationFile string
	cfg               myConfig.ServiceConfiguration
)

// init is called on package initialization and can therefore be used to initialize global stuff like logging, config, ..
func init() {
	config.Scan(&cfg)
}

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.cli.account"), // cfg.ServiceName
		micro.Version("latest"),            // cfg.Version
		// optional cli flag overrides. commit out if you don't need to override
		micro.Flags(
			cli.StringFlag{
				Name:        "config, c",
				Value:       "config.yml",
				Usage:       "Path to the configuration file to use. Defaults to config.yml",
				EnvVar:      "CONFIG_FILE",
				Destination: &configurationFile,
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
			}),
		micro.WrapHandler(wrapper.LogWrapper),
	)

	// Initialise service
	service.Init(
		micro.Action(func(c *cli.Context) {
			myConfig.InitConfig()
			config.Scan(&cfg)
		}),
	)

	log.Logf("IsProduction? %v", myConfig.IsProduction)
	log.Logf("environment: %v", cfg.Environment)
	log.Log(config.Get("database", "dialect").String("postgres"))
	log.Log(config.Get("database", "host").String("no-address"))
	log.Log(config.Get("database", "port").Int(0000))
	log.Log(config.Get("observability", "tracing", "flushInterval").Int(2000000000))
	log.Log(cfg)
	log.Logf("cfg is %v", cfg)
	log.Log(configurationFile)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

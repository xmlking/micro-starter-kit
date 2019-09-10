package main

import (
	"github.com/micro/cli"
	"github.com/micro/go-micro"

	// "github.com/micro/go-micro/server"
	"github.com/micro/go-micro/config"
	log "github.com/sirupsen/logrus"

	"github.com/xmlking/micro-starter-kit/srv/emailer/registry"
	"github.com/xmlking/micro-starter-kit/srv/emailer/subscriber"

	// "github.com/micro/go-micro/service/grpc"
	myConfig "github.com/xmlking/micro-starter-kit/shared/config"
	_ "github.com/xmlking/micro-starter-kit/shared/log"
)

const (
	serviceName = "emailer-srv"
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
	service := micro.NewService(
		// service := grpc.NewService(
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
				Name:        "configFile, f",
				Value:       "config.yaml",
				Usage:       "Config file in configDir. Defaults to 'config.yaml'",
				EnvVar:      "CONFIG_FILE",
				Destination: &configFile,
			}),
		micro.Name(serviceName),
		micro.Version(myConfig.Version),
	)

	// Initialise service
	service.Init(
		// TODO : implement graceful shutdown
		micro.Action(func(c *cli.Context) {
			// load config
			myConfig.InitConfig(configDir, configFile)
			config.Scan(&cfg)
		}),
	)

	// Initialise DI Container
	ctn, err := registry.NewContainer(cfg)
	defer ctn.Clean()
	if err != nil {
		log.Fatalf("failed to build container: %v", err)
	}

	emailSubscriber := ctn.Resolve("emailer-subscriber") //.(subscriber.EmailSubscriber)
	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.emailer", service.Server(), emailSubscriber)

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.emailer", service.Server(), subscriber.Handler)

	// register subscriber with queue, each message is delivered to a unique subscriber
	// micro.RegisterSubscriber("go.micro.srv.emailer.2", service.Server(), subscriber.Handler, server.SubscriberQueue("queue.pubsub"))

	myConfig.PrintBuildInfo()
	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

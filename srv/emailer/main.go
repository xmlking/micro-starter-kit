package main

import (
	"github.com/micro/cli"
	"github.com/micro/go-micro"

	// "github.com/micro/go-micro/server"
	"github.com/micro/go-micro/config"
	log "github.com/sirupsen/logrus"

	myConfig "github.com/xmlking/micro-starter-kit/shared/config"
	"github.com/xmlking/micro-starter-kit/srv/emailer/subscriber"
)

const (
	serviceName = "go.micro.srv.emailer"
)

var (
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
		// optional cli flag to override config. comment out if you don't need to override
		micro.Flags(
			cli.StringFlag{
				Name:        "config, c",
				Value:       "config/config.yaml",
				Usage:       "Path to the configuration file to use. Defaults to 'config/config.yaml'",
				EnvVar:      "CONFIG_FILE",
				Destination: &configFile,
			}),
		micro.Name(serviceName),
		micro.Version(myConfig.Version),
	)

	// Initialise service
	service.Init(
	// TODO : implement graceful shutdown
	)

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.emailer", service.Server(), new(subscriber.Emailer))

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

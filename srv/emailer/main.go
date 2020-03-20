package main

import (
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/transport"
	"github.com/xmlking/logger/log"

	// "github.com/micro/go-micro/v2/service/grpc"
	myConfig "github.com/xmlking/micro-starter-kit/shared/config"
	"github.com/xmlking/micro-starter-kit/shared/logger"
	"github.com/xmlking/micro-starter-kit/shared/util"
	logWrapper "github.com/xmlking/micro-starter-kit/shared/wrapper/log"
	transWrapper "github.com/xmlking/micro-starter-kit/shared/wrapper/transaction"
	"github.com/xmlking/micro-starter-kit/srv/emailer/registry"
	// "github.com/xmlking/micro-starter-kit/srv/emailer/subscriber"
)

const (
	serviceName = "emailersrv"
)

var (
	configDir  string
	configFile string
	cfg        myConfig.ServiceConfiguration
)

func main() {
	// New Service
	service := micro.NewService(
		// optional cli flag to override config.
		// comment out if you don't need to override any base config via CLI
		micro.Flags(
			&cli.StringFlag{
				Name:        "configDir",
				Aliases:     []string{"d"},
				Value:       "/config",
				Usage:       "Path to the config directory. Defaults to 'config'",
				EnvVars:     []string{"CONFIG_DIR"},
				Destination: &configDir,
			},
			&cli.StringFlag{
				Name:        "configFile",
				Aliases:     []string{"f"},
				Value:       "config.yaml",
				Usage:       "Config file in configDir. Defaults to 'config.yaml'",
				EnvVars:     []string{"CONFIG_FILE"},
				Destination: &configFile,
			},
		),
		micro.Name(serviceName),
		micro.Version(myConfig.Version),
	)

	// Initialize service
	service.Init(
		// TODO : implement graceful shutdown
		micro.Action(func(c *cli.Context) (err error) {
			// load config
			myConfig.InitConfig(configDir, configFile)
			err = config.Scan(&cfg)
			logger.InitLogger(cfg.Log)
			return
		}),
	)

	// Initialize Features
	var options []micro.Option
	if cfg.Features["mtls"].Enabled {
		if tlsConf, err := util.GetSelfSignedTLSConfig("localhost"); err != nil {
			log.WithError(err).Error("unable to load certs")
		} else {
			options = append(options,
				// https://github.com/ykumar-rb/ZTP/blob/master/pnp/server.go
				// grpc.WithTLS(tlsConf),
				micro.Transport(transport.NewTransport(transport.Secure(true))),
				micro.Transport(transport.NewTransport(transport.TLSConfig(tlsConf))),
			)
		}
	}
	// Wrappers are invoked in the order as they added
	if cfg.Features["reqlogs"].Enabled {
		options = append(options, micro.WrapSubscriber(logWrapper.NewSubscriberWrapper()))
	}
	if cfg.Features["translogs"].Enabled {
		topic := config.Get("features", "translogs", "topic").String("recordersrv")
		publisher := micro.NewEvent(topic, service.Client())
		options = append(options, micro.WrapSubscriber(transWrapper.NewSubscriberWrapper(publisher)))
	}

	// Initialize Features
	service.Init(
		options...,
	)

	// Initialize DI Container
	ctn, err := registry.NewContainer(cfg)
	defer ctn.Clean()
	if err != nil {
		log.Fatalf("failed to build container: %v", err)
	}

	emailSubscriber := ctn.Resolve("emailer-subscriber") //.(subscriber.EmailSubscriber)
	// Register Struct as Subscriber
	micro.RegisterSubscriber("emailersrv", service.Server(), emailSubscriber)

	// Register Function as Subscriber
	// micro.RegisterSubscriber("emailersrv", service.Server(), subscriber.Handler)

	// register subscriber with queue, each message is delivered to a unique subscriber
	// micro.RegisterSubscriber("emailersrv-2", service.Server(), subscriber.Handler, server.SubscriberQueue("queue.pubsub"))

	myConfig.PrintBuildInfo()
	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

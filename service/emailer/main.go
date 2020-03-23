package main

import (
	"path/filepath"

	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/config"
	"github.com/xmlking/logger/log"
	"github.com/xmlking/micro-starter-kit/shared/constants"

	// "github.com/micro/go-micro/v2/service/grpc"
	"github.com/xmlking/micro-starter-kit/service/emailer/registry"
	myConfig "github.com/xmlking/micro-starter-kit/shared/config"
	"github.com/xmlking/micro-starter-kit/shared/logger"
	"github.com/xmlking/micro-starter-kit/shared/util"
	logWrapper "github.com/xmlking/micro-starter-kit/shared/wrapper/log"
	transWrapper "github.com/xmlking/micro-starter-kit/shared/wrapper/transaction"
	// "github.com/xmlking/micro-starter-kit/service/emailer/subscriber"
)

const (
	serviceName = constants.EMAILER_SERVICE
	configDir   = "/config"
	configFile  = "config.yaml"
)

var (
	cfg myConfig.ServiceConfiguration
)

func main() {
	// New Service
	service := micro.NewService(
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
		// if tlsConf, err := util.GetSelfSignedTLSConfig("localhost"); err != nil {
		if tlsConf, err := util.GetTLSConfig(
			filepath.Join(configDir, config.Get("features", "mtls", "certfile").String("")),
			filepath.Join(configDir, config.Get("features", "mtls", "keyfile").String("")),
			filepath.Join(configDir, config.Get("features", "mtls", "cafile").String("")),
			filepath.Join(configDir, config.Get("features", "mtls", "servername").String("")),
		); err != nil {
			log.WithError(err).Error("unable to load certs")
		} else {
			println(tlsConf)
			options = append(options,
				util.WithTLS(tlsConf),
			)
		}
	}
	// Wrappers are invoked in the order as they added
	if cfg.Features["reqlogs"].Enabled {
		options = append(options, micro.WrapSubscriber(logWrapper.NewSubscriberWrapper()))
	}
	if cfg.Features["translogs"].Enabled {
		topic := config.Get("features", "translogs", "topic").String(constants.RECORDER_SERVICE)
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
	micro.RegisterSubscriber(constants.EMAILER_SERVICE, service.Server(), emailSubscriber)

	// Register Function as Subscriber
	// micro.RegisterSubscriber(constants.EMAILER_SERVICE, service.Server(), subscriber.Handler)

	// register subscriber with queue, each message is delivered to a unique subscriber
	// micro.RegisterSubscriber("mkit.service.emailer-2", service.Server(), subscriber.Handler, server.SubscriberQueue("queue.pubsub"))

	myConfig.PrintBuildInfo()
	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

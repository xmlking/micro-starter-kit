package main

import (
	"path/filepath"

	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/config"
	"github.com/xmlking/logger/log"

	transactionPB "github.com/xmlking/micro-starter-kit/service/recorder/proto/transaction"
	"github.com/xmlking/micro-starter-kit/service/recorder/registry"
	myConfig "github.com/xmlking/micro-starter-kit/shared/config"
	"github.com/xmlking/micro-starter-kit/shared/constants"
	"github.com/xmlking/micro-starter-kit/shared/logger"
	"github.com/xmlking/micro-starter-kit/shared/util"
	logWrapper "github.com/xmlking/micro-starter-kit/shared/wrapper/log"
)

const (
	serviceName = constants.RECORDER_SERVICE
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
		options = append(options,
			micro.WrapSubscriber(logWrapper.NewSubscriberWrapper()),
		)
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

	transactionSubscriber := ctn.Resolve("transaction-subscriber") //.(subscriber.TransactionSubscriber)
	recorderTopic := config.Get("recorder", "topic").String(constants.RECORDER_SERVICE)

	// Register Struct as Subscriber
	_ = micro.RegisterSubscriber(recorderTopic, service.Server(), transactionSubscriber)

	// register subscriber with queue, each message is delivered to a unique subscriber
	// _ = micro.RegisterSubscriber(recorderTopic, service.Server(), transactionSubscriber, server.SubscriberQueue("queue.pubsub"))

	transactionHandler := ctn.Resolve("transaction-handler").(transactionPB.TransactionServiceHandler)
	transactionPB.RegisterTransactionServiceHandler(service.Server(), transactionHandler)

	myConfig.PrintBuildInfo()
	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

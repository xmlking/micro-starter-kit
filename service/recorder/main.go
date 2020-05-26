package main

import (
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/config"
	"github.com/rs/zerolog/log"

	transactionPB "github.com/xmlking/micro-starter-kit/service/recorder/proto/transaction"
	"github.com/xmlking/micro-starter-kit/service/recorder/registry"
	myConfig "github.com/xmlking/micro-starter-kit/shared/config"
	"github.com/xmlking/micro-starter-kit/shared/constants"
	_ "github.com/xmlking/micro-starter-kit/shared/logger"
	"github.com/xmlking/micro-starter-kit/shared/util"
	logWrapper "github.com/xmlking/micro-starter-kit/shared/wrapper/log"
)

const (
	serviceName = constants.RECORDER_SERVICE
)

var (
    cfg = myConfig.GetServiceConfig()
    ff = myConfig.GetFeatureFlags()
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
			// do some life cycle actions
			return
		}),
	)

	// Initialize Features
	var options []micro.Option
	if ff.IsTLSEnabled() {
		if tlsConf, err := myConfig.CreateServerCerts(); err != nil {
			log.Error().Err(err).Msg("unable to load certs")
		} else {
            log.Info().Msg("TLS Enabled")
			options = append(options,
				util.WithTLS(tlsConf),
			)
		}
	}
	// Wrappers are invoked in the order as they added
    if ff.IsReqlogsEnabled() {
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
        log.Fatal().Msgf("failed to build container: %v", err)
	}

	transactionSubscriber := ctn.Resolve("transaction-subscriber") //.(subscriber.TransactionSubscriber)
	recorderTopic := config.Get("recorder", "topic").String(constants.RECORDER_SERVICE)

	// Register Struct as Subscriber
	_ = micro.RegisterSubscriber(recorderTopic, service.Server(), transactionSubscriber)

	// register subscriber with queue, each message is delivered to a unique subscriber
	// _ = micro.RegisterSubscriber(recorderTopic, service.Server(), transactionSubscriber, server.SubscriberQueue("queue.pubsub"))

	transactionHandler := ctn.Resolve("transaction-handler").(transactionPB.TransactionServiceHandler)
	transactionPB.RegisterTransactionServiceHandler(service.Server(), transactionHandler)

    println(myConfig.GetBuildInfo())
	// Run service
	if err := service.Run(); err != nil {
        log.Fatal().Err(err).Msg("")
	}
}

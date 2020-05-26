package main

import (
    "github.com/micro/cli/v2"
    "github.com/micro/go-micro/v2"
    "github.com/micro/go-micro/v2/config"
    "github.com/rs/zerolog/log"

    "github.com/xmlking/micro-starter-kit/shared/constants"

    // "github.com/micro/go-micro/v2/service/grpc"
    "github.com/xmlking/micro-starter-kit/service/emailer/registry"
    myConfig "github.com/xmlking/micro-starter-kit/shared/config"
    _ "github.com/xmlking/micro-starter-kit/shared/logger"
    "github.com/xmlking/micro-starter-kit/shared/util"
    logWrapper "github.com/xmlking/micro-starter-kit/shared/wrapper/log"
    transWrapper "github.com/xmlking/micro-starter-kit/shared/wrapper/transaction"
    // "github.com/xmlking/micro-starter-kit/service/emailer/subscriber"
)

const (
	serviceName = constants.EMAILER_SERVICE
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
		options = append(options, micro.WrapSubscriber(logWrapper.NewSubscriberWrapper()))
	}
    if ff.IsTranslogsEnabled() {
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
        log.Fatal().Msgf("failed to build container: %v", err)
	}

	emailSubscriber := ctn.Resolve("emailer-subscriber") //.(subscriber.EmailSubscriber)
	// Register Struct as Subscriber
	micro.RegisterSubscriber(constants.EMAILER_SERVICE, service.Server(), emailSubscriber)

	// Register Function as Subscriber
	// micro.RegisterSubscriber(constants.EMAILER_SERVICE, service.Server(), subscriber.Handler)

	// register subscriber with queue, each message is delivered to a unique subscriber
	// micro.RegisterSubscriber("mkit.service.emailer-2", service.Server(), subscriber.Handler, server.SubscriberQueue("queue.pubsub"))

	// PrintBuildInfo
    println(myConfig.GetBuildInfo())
	// Run service
	if err := service.Run(); err != nil {
        log.Fatal().Err(err).Msg("")
	}
}

package main

import (
    "github.com/micro/go-micro/v2"
    sgrpc "github.com/micro/go-micro/v2/server/grpc"

    "github.com/rs/zerolog/log"

    "github.com/xmlking/micro-starter-kit/shared/constants"

    "github.com/xmlking/micro-starter-kit/service/emailer/registry"
    "github.com/xmlking/micro-starter-kit/shared/config"
    "github.com/xmlking/micro-starter-kit/shared/util/tls"
    logWrapper "github.com/xmlking/micro-starter-kit/shared/wrapper/log"
    transWrapper "github.com/xmlking/micro-starter-kit/shared/wrapper/transaction"
    // "github.com/xmlking/micro-starter-kit/service/emailer/subscriber"
)

func main() {
    serviceName := constants.EMAILER_SERVICE
    cfg := config.GetConfig()

    lis, err := config.GetListener(cfg.Services.Emailer.Endpoint)
    if err != nil {
        log.Fatal().Msgf("failed to create listener: %v", err)
    }

    // New Service
    service := micro.NewService(
        micro.Server(sgrpc.NewServer(sgrpc.Listener(lis))), // KEEP-IT-FIRST
        micro.Name(serviceName),
        micro.Version(config.Version),
    )

    // Initialize Features
    var options []micro.Option

    if cfg.Features.Tls.Enabled {
        if tlsConf, err := config.CreateServerCerts(); err != nil {
            log.Error().Err(err).Msg("unable to load certs")
        } else {
            log.Info().Msg("TLS Enabled")
            options = append(options, tls.WithTLS(tlsConf))
        }
    }

    // Wrappers are invoked in the order as they added
    if cfg.Features.Reqlogs.Enabled {
        options = append(options, micro.WrapSubscriber(logWrapper.NewSubscriberWrapper()))
    }
    if cfg.Features.Translogs.Enabled {
        topic := cfg.Features.Translogs.Topic
        publisher := micro.NewEvent(topic, service.Client())
        options = append(options, micro.WrapSubscriber(transWrapper.NewSubscriberWrapper(publisher)))
    }

    // Adding some optional lifecycle actions
    options = append(options,
        micro.BeforeStart(func() (err error) {
            log.Debug().Msg("called BeforeStart")
            return
        }),
        micro.BeforeStop(func() (err error) {
            log.Debug().Msg("called BeforeStop")
            return
        }),
    )

    // Initialize service
    service.Init(options...)

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

    println(config.GetBuildInfo())

    // Run service
    if err := service.Run(); err != nil {
        log.Fatal().Err(err).Send()
    }
}

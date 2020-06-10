package main

import (
    "github.com/micro/go-micro/v2"
    "github.com/rs/zerolog/log"

    transactionPB "github.com/xmlking/micro-starter-kit/service/recorder/proto/transaction"
    "github.com/xmlking/micro-starter-kit/service/recorder/registry"
    "github.com/xmlking/micro-starter-kit/shared/config"
    "github.com/xmlking/micro-starter-kit/shared/constants"
    "github.com/xmlking/micro-starter-kit/shared/util/tls"
    logWrapper "github.com/xmlking/micro-starter-kit/shared/wrapper/log"
)

const (
    serviceName = constants.RECORDER_SERVICE
)

var (
    cfg = config.GetConfig()
)

func main() {
    // New Service
    service := micro.NewService(
        micro.Name(serviceName),
        micro.Version(config.Version),
    )

    // Initialize service
    service.Init(
        micro.BeforeStart(func() (err error) {
            return
        }),
        micro.BeforeStop(func() (err error) {
            return
        }),
    )

    // Initialize Features
    var options []micro.Option
    if cfg.Features.Tls.Enabled {
        if tlsConf, err := config.CreateServerCerts(); err != nil {
            log.Error().Err(err).Msg("unable to load certs")
        } else {
            log.Info().Msg("TLS Enabled")
            options = append(options,
                tls.WithTLS(tlsConf),
            )
        }
    }
    // Wrappers are invoked in the order as they added
    if cfg.Features.Reqlogs.Enabled {
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
    recorderTopic := cfg.Features.Translogs.Topic

    // Register Struct as Subscriber
    _ = micro.RegisterSubscriber(recorderTopic, service.Server(), transactionSubscriber)

    // register subscriber with queue, each message is delivered to a unique subscriber
    // _ = micro.RegisterSubscriber(recorderTopic, service.Server(), transactionSubscriber, server.SubscriberQueue("queue.pubsub"))

    transactionHandler := ctn.Resolve("transaction-handler").(transactionPB.TransactionServiceHandler)
    transactionPB.RegisterTransactionServiceHandler(service.Server(), transactionHandler)

    println(config.GetBuildInfo())

    // Run service
    if err := service.Run(); err != nil {
        log.Fatal().Err(err).Msg("")
    }
}

package main

import (
    "github.com/micro/go-micro/v2"
    sgrpc "github.com/micro/go-micro/v2/server/grpc"
    "github.com/rs/zerolog/log"

    transactionPB "github.com/xmlking/micro-starter-kit/service/recorder/proto/transaction"
    "github.com/xmlking/micro-starter-kit/service/recorder/registry"
    "github.com/xmlking/micro-starter-kit/shared/config"
    "github.com/xmlking/micro-starter-kit/shared/constants"
    "github.com/xmlking/micro-starter-kit/shared/util/tls"
    logWrapper "github.com/xmlking/micro-starter-kit/shared/wrapper/log"
)

func main() {
    serviceName := constants.RECORDER_SERVICE
    cfg := config.GetConfig()

    lis, err := config.GetListener(cfg.Services.Recorder.Endpoint)
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
        log.Fatal().Err(err).Send()
    }
}

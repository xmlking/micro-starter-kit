package main

import (
    "github.com/micro/go-micro/v2"
    "github.com/micro/go-micro/v2/client"
    "github.com/micro/go-micro/v2/server"
    "github.com/rs/zerolog/log"

    transactionPB "github.com/xmlking/micro-starter-kit/service/recorder/proto/transaction"
    "github.com/xmlking/micro-starter-kit/service/recorder/registry"
    "github.com/xmlking/micro-starter-kit/shared/config"
    "github.com/xmlking/micro-starter-kit/shared/constants"
    myMicro "github.com/xmlking/micro-starter-kit/shared/util/micro"
    logWrapper "github.com/xmlking/micro-starter-kit/shared/wrapper/log"
    validatorWrapper "github.com/xmlking/micro-starter-kit/shared/wrapper/validator"
)

func main() {
    cfg := config.GetConfig()

    // Initialize Features
    var clientWrappers []client.Wrapper
    var handlerWrappers []server.HandlerWrapper
    var subscriberWrappers []server.SubscriberWrapper

    // Wrappers are invoked in the order as they added
    if cfg.Features.Reqlogs.Enabled {
        clientWrappers = append(clientWrappers, logWrapper.NewClientWrapper())
        handlerWrappers = append(handlerWrappers, logWrapper.NewHandlerWrapper())
        subscriberWrappers = append(subscriberWrappers, logWrapper.NewSubscriberWrapper())
    }
    if cfg.Features.Validator.Enabled {
        handlerWrappers = append(handlerWrappers, validatorWrapper.NewHandlerWrapper())
        subscriberWrappers = append(subscriberWrappers, validatorWrapper.NewSubscriberWrapper())
    }

    service := micro.NewService(
        micro.Name(constants.RECORDER_SERVICE),
        micro.Version(config.Version),
        myMicro.WithTLS(),
        // Wrappers are applied in reverse order so the last is executed first.
        micro.WrapClient(clientWrappers...),
        micro.WrapHandler(handlerWrappers...),
        micro.WrapSubscriber(subscriberWrappers...),
        // Adding some optional lifecycle actions
        micro.BeforeStart(func() (err error) {
            log.Debug().Msg("called BeforeStart")
            return
        }),
        micro.BeforeStop(func() (err error) {
            log.Debug().Msg("called BeforeStop")
            return
        }),
    )

    service.Init()

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

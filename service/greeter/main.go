package main

import (
    "github.com/micro/go-micro/v2"
    "github.com/micro/go-micro/v2/client"
    "github.com/micro/go-micro/v2/server"
    "github.com/rs/zerolog/log"

    "github.com/xmlking/micro-starter-kit/service/greeter/handler"
    greeterPB "github.com/xmlking/micro-starter-kit/service/greeter/proto/greeter"
    healthPB "github.com/xmlking/micro-starter-kit/service/greeter/proto/health"
    "github.com/xmlking/micro-starter-kit/shared/config"
    "github.com/xmlking/micro-starter-kit/shared/constants"
    myMicro "github.com/xmlking/micro-starter-kit/shared/util/micro"
    logWrapper "github.com/xmlking/micro-starter-kit/shared/wrapper/log"
    transWrapper "github.com/xmlking/micro-starter-kit/shared/wrapper/transaction"
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
    //if cfg.Features.Translogs.Enabled {
    //    topic := cfg.Features.Translogs.Topic
    //    publisher := micro.NewEvent(topic, client.DefaultClient) // service.Client())
    //    handlerWrappers = append(handlerWrappers, transWrapper.NewHandlerWrapper(publisher))
    //    subscriberWrappers = append(subscriberWrappers, transWrapper.NewSubscriberWrapper(publisher))
    //}
    if cfg.Features.Validator.Enabled {
        handlerWrappers = append(handlerWrappers, validatorWrapper.NewHandlerWrapper())
        subscriberWrappers = append(subscriberWrappers, validatorWrapper.NewSubscriberWrapper())
    }

    service := micro.NewService(
        micro.Name(constants.GREETER_SERVICE),
        micro.Version(config.Version),
        myMicro.WithTLS(),
        // Wrappers are applied in reverse order so the last is executed first.
        micro.WrapClient(clientWrappers...),
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

    if cfg.Features.Translogs.Enabled {
        topic := cfg.Features.Translogs.Topic
        publisher := micro.NewEvent(topic, service.Client())
        handlerWrappers = append(handlerWrappers, transWrapper.NewHandlerWrapper(publisher))
        subscriberWrappers = append(subscriberWrappers, transWrapper.NewSubscriberWrapper(publisher))
    }

    service.Init(
        micro.WrapHandler(handlerWrappers...),
        micro.WrapSubscriber(subscriberWrappers...),
    )

    _ = healthPB.RegisterHealthHandler(service.Server(), handler.NewHealthHandler())
    _ = greeterPB.RegisterGreeterServiceHandler(service.Server(), handler.NewGreeterHandler())

    println(config.GetBuildInfo())

    // Run service
    if err := service.Run(); err != nil {
        log.Fatal().Err(err).Send()
    }
}

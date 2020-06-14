package main

import (
    "github.com/micro/go-micro/v2"
    "github.com/micro/go-micro/v2/client"
    "github.com/micro/go-micro/v2/server"
    "github.com/rs/zerolog/log"

    "github.com/xmlking/micro-starter-kit/service/account/handler"
    profilePB "github.com/xmlking/micro-starter-kit/service/account/proto/profile"
    userPB "github.com/xmlking/micro-starter-kit/service/account/proto/user"
    "github.com/xmlking/micro-starter-kit/service/account/registry"
    "github.com/xmlking/micro-starter-kit/service/account/repository"
    greeterPB "github.com/xmlking/micro-starter-kit/service/greeter/proto/greeter"
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
        micro.Name(constants.ACCOUNT_SERVICE),
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

    // Initialize DI Container
    ctn, err := registry.NewContainer(cfg)
    defer ctn.Clean()
    if err != nil {
        log.Fatal().Msgf("failed to build container: %v", err)
    }

    // Publisher publish to "mkit.service.emailer"
    publisher := micro.NewEvent(constants.EMAILER_SERVICE, service.Client())
    // greeterSrv Client to call "mkit.service.greeter"
    greeterSrvClient := greeterPB.NewGreeterService(constants.GREETER_SERVICE, service.Client())

    // // Handlers
    userHandler := handler.NewUserHandler(ctn.Resolve("user-repository").(repository.UserRepository), publisher, greeterSrvClient)
    profileHandler := ctn.Resolve("profile-handler").(profilePB.ProfileServiceHandler)

    // Register Handlers
    userPB.RegisterUserServiceHandler(service.Server(), userHandler)
    profilePB.RegisterProfileServiceHandler(service.Server(), profileHandler)

    println(config.GetBuildInfo())

    // Run service
    if err := service.Run(); err != nil {
        log.Fatal().Err(err).Send()
    }
}

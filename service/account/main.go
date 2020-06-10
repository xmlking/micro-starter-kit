package main

import (
    "github.com/micro/go-micro/v2"
    "github.com/rs/zerolog/log"

    // "github.com/micro/go-micro/v2/service/grpc"
    "github.com/xmlking/micro-starter-kit/service/account/handler"
    profilePB "github.com/xmlking/micro-starter-kit/service/account/proto/profile"
    userPB "github.com/xmlking/micro-starter-kit/service/account/proto/user"
    "github.com/xmlking/micro-starter-kit/service/account/registry"
    "github.com/xmlking/micro-starter-kit/service/account/repository"
    greeterPB "github.com/xmlking/micro-starter-kit/service/greeter/proto/greeter"
    "github.com/xmlking/micro-starter-kit/shared/config"
    "github.com/xmlking/micro-starter-kit/shared/constants"
    "github.com/xmlking/micro-starter-kit/shared/util/tls"
    logWrapper "github.com/xmlking/micro-starter-kit/shared/wrapper/log"
    transWrapper "github.com/xmlking/micro-starter-kit/shared/wrapper/transaction"
    validatorWrapper "github.com/xmlking/micro-starter-kit/shared/wrapper/validator"
)

const (
    serviceName = constants.ACCOUNT_SERVICE
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
            micro.WrapHandler(logWrapper.NewHandlerWrapper()),
            micro.WrapClient(logWrapper.NewClientWrapper()),
        )
    }
    if cfg.Features.Validator.Enabled {
        options = append(options,
            micro.WrapHandler(validatorWrapper.NewHandlerWrapper()),
        )
    }
    if cfg.Features.Translogs.Enabled {
        topic := cfg.Features.Translogs.Topic
        publisher := micro.NewEvent(topic, service.Client())
        options = append(options,
            micro.WrapHandler(transWrapper.NewHandlerWrapper(publisher)),
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

    log.Debug().Msgf("Client type: grpc or regular? %T\n", service.Client()) // FIXME: expected *grpc.grpcClient but got *micro.clientWrapper

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
        log.Fatal().Err(err).Msg("")
    }
}

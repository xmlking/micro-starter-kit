package main

import (
    "github.com/micro/go-micro/v2"
    "github.com/rs/zerolog/log"

    "github.com/xmlking/micro-starter-kit/service/greeter/handler"
    greeterPB "github.com/xmlking/micro-starter-kit/service/greeter/proto/greeter"
    healthPB "github.com/xmlking/micro-starter-kit/service/greeter/proto/health"
    "github.com/xmlking/micro-starter-kit/shared/config"
    "github.com/xmlking/micro-starter-kit/shared/constants"
    "github.com/xmlking/micro-starter-kit/shared/util/tls"
    logWrapper "github.com/xmlking/micro-starter-kit/shared/wrapper/log"
    transWrapper "github.com/xmlking/micro-starter-kit/shared/wrapper/transaction"
)

const (
    serviceName = constants.GREETER_SERVICE
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
        options = append(options, micro.WrapHandler(logWrapper.NewHandlerWrapper()))
    }
    if cfg.Features.Translogs.Enabled {
        topic := cfg.Features.Translogs.Topic
        publisher := micro.NewEvent(topic, service.Client())
        options = append(options, micro.WrapHandler(transWrapper.NewHandlerWrapper(publisher)))
    }

    // Initialize Features
    service.Init(
        options...,
    )

    // Register Handler
    _ = greeterPB.RegisterGreeterServiceHandler(service.Server(), handler.NewGreeterHandler())
    _ = healthPB.RegisterHealthHandler(service.Server(), handler.NewHealthHandler())

    println(config.GetBuildInfo())

    // Run service
    if err := service.Run(); err != nil {
        log.Fatal().Err(err).Msg("")
    }
}

package main

import (
    "github.com/micro/go-micro/v2"

    "github.com/rs/zerolog/log"

    "github.com/xmlking/micro-starter-kit/shared/config"
    _ "github.com/xmlking/micro-starter-kit/shared/logger"

    logWrapper "github.com/xmlking/micro-starter-kit/shared/wrapper/log"
)

const (
    serviceName = "account-cmd"
)

var (
    cfg = config.GetConfig()
)

func main() {
    // New Service
    service := micro.NewService(
        micro.Name(serviceName),
        micro.Version(config.Version),
        micro.WrapHandler(logWrapper.NewHandlerWrapper()),
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

    log.Debug().Msgf("IsProduction? %v", config.IsProduction())
    log.Debug().Interface("Dialect", cfg.Database.Dialect).Send()
    log.Debug().Msg(cfg.Database.Host)
    log.Debug().Uint32("Port", cfg.Database.Port).Send()
    log.Debug().Uint64("FlushInterval", cfg.Features.Tracing.FlushInterval).Send()
    log.Debug().Msgf("cfg is %v", cfg)

    // Run service
    if err := service.Run(); err != nil {
        log.Fatal().Err(err).Msg("")
    }
}

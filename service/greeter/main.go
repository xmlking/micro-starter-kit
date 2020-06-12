package main

import (
    "net/http"

    "github.com/micro/go-micro/v2"
    "github.com/rs/zerolog/log"
    "github.com/soheilhy/cmux"
    //"google.golang.org/grpc"
    //"google.golang.org/grpc/health"
    //"google.golang.org/grpc/health/grpc_health_v1"

    sgrpc "github.com/micro/go-micro/v2/server/grpc"

    "github.com/xmlking/micro-starter-kit/service/greeter/handler"
    greeterPB "github.com/xmlking/micro-starter-kit/service/greeter/proto/greeter"
    healthPB "github.com/xmlking/micro-starter-kit/service/greeter/proto/health"
    "github.com/xmlking/micro-starter-kit/shared/config"
    "github.com/xmlking/micro-starter-kit/shared/constants"
    "github.com/xmlking/micro-starter-kit/shared/util/tls"
    logWrapper "github.com/xmlking/micro-starter-kit/shared/wrapper/log"
    transWrapper "github.com/xmlking/micro-starter-kit/shared/wrapper/transaction"
)

func main() {
    serviceName := constants.GREETER_SERVICE
    cfg := config.GetConfig()

    lis, err := config.GetListener(cfg.Services.Greeter.Endpoint)
    if err != nil {
        log.Fatal().Msgf("failed to create listener: %v", err)
    }

    // Create a cmux.
    mux := cmux.New(lis)
    // Match connections in order:
    grpcL := mux.Match(cmux.HTTP2HeaderField("content-type", "application/grpc"))
    httpL := mux.Match(cmux.HTTP1Fast())

	// New Service
	service := micro.NewService(
        // Using grpc listener created by cmux
        micro.Server(sgrpc.NewServer(sgrpc.Listener(grpcL))), // KEEP-IT-FIRST
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
        options = append(options, micro.WrapHandler(logWrapper.NewHandlerWrapper()))
    }
    if cfg.Features.Translogs.Enabled {
        topic := cfg.Features.Translogs.Topic
        publisher := micro.NewEvent(topic, service.Client())
        options = append(options, micro.WrapHandler(transWrapper.NewHandlerWrapper(publisher)))
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
        micro.AfterStart(func() (err error) {
          log.Debug().Msg("called AfterStart")
          // Start cmux after go-micro starts
          return mux.Serve()
        }),
    )

    // Initialize service
    service.Init(options...)


    // Register http Handlers
    httpS := &http.Server{
        Handler: handler.NewHttpHandler(),
    }

    // Register grpc Handlers
    /**
    grpcS := grpc.NewServer()
    hsrv := health.NewServer()
    for name, _ := range grpcS.GetServiceInfo() {
       hsrv.SetServingStatus(name, grpc_health_v1.HealthCheckResponse_SERVING)
    }
    grpc_health_v1.RegisterHealthServer(grpcS, hsrv)
    **/
    _ = healthPB.RegisterHealthHandler(service.Server(), handler.NewHealthHandler())
    _ = greeterPB.RegisterGreeterServiceHandler(service.Server(), handler.NewGreeterHandler())

    println(config.GetBuildInfo())

    // Run http servers.
    go httpS.Serve(httpL)

    // Run grpc service
    if err := service.Run(); err != nil {
       log.Fatal().Err(err).Send()
    }
}

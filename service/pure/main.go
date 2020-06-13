package main
import (
    "github.com/rs/zerolog/log"
    "google.golang.org/grpc"

    "github.com/xmlking/micro-starter-kit/service/pure/handler"
    purePB "github.com/xmlking/micro-starter-kit/service/pure/proto/pure"
    "github.com/xmlking/micro-starter-kit/shared/config"
    "github.com/xmlking/micro-starter-kit/shared/constants"
)

func main() {
    serviceName := constants.GREETER_SERVICE
    cfg := config.GetConfig()

    println(serviceName)
    lis, err := config.GetListener(cfg.Services.Greeter.Endpoint)
    if err != nil {
        log.Fatal().Msgf("failed to create listener: %v", err)
    }

    // create a server instance
    s := handler.NewPureHandler()
    // create a gRPC server object
    grpcServer := grpc.NewServer()
    // attach the Ping service to the server
    purePB.RegisterPingServer(grpcServer, s)
    // start the server
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatal().Err(err).Send()
    }
}

package main

import (
    "context"

    "github.com/rs/zerolog/log"
    "google.golang.org/grpc/balancer/roundrobin"

    "google.golang.org/grpc"

    purePB "github.com/xmlking/micro-starter-kit/service/pure/proto/pure"
    "github.com/xmlking/micro-starter-kit/shared/config"
    "github.com/xmlking/micro-starter-kit/shared/constants"
)

func main() {
    serviceName := constants.GREETER_SERVICE
    cfg := config.GetConfig()

    println(serviceName)
    var conn *grpc.ClientConn

    conn, err := grpc.Dial(cfg.Services.Greeter.Endpoint, grpc.WithInsecure(), grpc.WithBalancerName(roundrobin.Name))
    if err != nil {
        log.Fatal().Msgf("did not connect: %s", err)
    }
    defer conn.Close()
    c := purePB.NewPingClient(conn)
    response, err := c.SayHello(context.Background(), &purePB.PingMessage{Greeting: "foo"})
    if err != nil {
        log.Fatal().Msgf("Error when calling SayHello: %s", err)
    }
    log.Printf("Response from server: %s", response.Greeting)
}

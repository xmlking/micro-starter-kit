package main

import (
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/server"
	"github.com/micro/micro/cmd"

	cli "github.com/micro/go-micro/client/grpc"
	srv "github.com/micro/go-micro/server/grpc"
	bkr "github.com/micro/go-plugins/broker/grpc"
)

func main() {
	// setup broker/client/server to use grpc
	broker.DefaultBroker = bkr.NewBroker()
	client.DefaultClient = cli.NewClient()
	server.DefaultServer = srv.NewServer()

	cmd.Init()
}

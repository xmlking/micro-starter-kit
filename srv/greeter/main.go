package main

import (
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro"
	"github.com/xmlking/micro-starter-kit/srv/greeter/handler"


	greeterPB "github.com/xmlking/micro-starter-kit/srv/greeter/proto/greeter"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.greeter"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	greeterPB.RegisterGreeterHandler(service.Server(), new(handler.Greeter))


	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro"
	"github.com/xmlking/micro-starter-kit/srv/echo/handler"
	"github.com/xmlking/micro-starter-kit/srv/echo/subscriber"

	echo "github.com/xmlking/micro-starter-kit/srv/echo/proto/echo"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.echo"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	echo.RegisterEchoHandler(service.Server(), new(handler.Echo))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.echo", service.Server(), new(subscriber.Echo))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.echo", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro"
	"github.com/xmlking/micro-starter-kit/srv/account/handler"
	"github.com/xmlking/micro-starter-kit/srv/account/subscriber"

	account "github.com/xmlking/micro-starter-kit/srv/account/proto/account"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.account"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	account.RegisterAccountHandler(service.Server(), new(handler.Account))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.account", service.Server(), new(subscriber.Account))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.account", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

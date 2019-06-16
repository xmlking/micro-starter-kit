package main

import (
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro"
	
	"github.com/xmlking/micro-starter-kit/srv/emailer/subscriber"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.emailer"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.emailer", service.Server(), new(subscriber.Emailer))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.emailer", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

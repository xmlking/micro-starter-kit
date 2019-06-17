package main

import (
	"github.com/micro/go-micro"
	// "github.com/micro/go-micro/server"
	"github.com/micro/go-micro/util/log"

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

	// register subscriber with queue, each message is delivered to a unique subscriber
	// micro.RegisterSubscriber("go.micro.srv.emailer.2", service.Server(), subscriber.Handler, server.SubscriberQueue("queue.pubsub"))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

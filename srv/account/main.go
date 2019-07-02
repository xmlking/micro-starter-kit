package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/config"
	log "github.com/sirupsen/logrus"

	"github.com/xmlking/micro-starter-kit/srv/account/handler"
	accountPB "github.com/xmlking/micro-starter-kit/srv/account/proto/account"
	"github.com/xmlking/micro-starter-kit/srv/account/registry"
	"github.com/xmlking/micro-starter-kit/srv/account/repository"
	"github.com/xmlking/micro-starter-kit/srv/account/subscriber"

	myConfig "github.com/xmlking/micro-starter-kit/shared/config"
	"github.com/xmlking/micro-starter-kit/shared/wrapper"
)

var (
	cfg myConfig.ServiceConfiguration
)

// init is called on package initialization and can therefore be used to initialize global stuff like logging, config, ..
func init() {
	config.Scan(&cfg)
}

func main() {

	// New Service
	service := micro.NewService(
		micro.Name(cfg.ServiceName),
		micro.Version(cfg.Version),
		micro.WrapHandler(wrapper.LogWrapper),
	)

	// Initialise service
	service.Init()

	// Initialise DI Container
	ctn, err := registry.NewContainer(cfg)
	defer ctn.Clean()
	if err != nil {
		log.Fatalf("failed to build container: %v", err)
	}

	// Publisher publish to "go.micro.srv.emailer"
	publisher := micro.NewPublisher("go.micro.srv.emailer", service.Client())

	// // Handlers
	userHandler := handler.NewUserHandler(ctn.Resolve("user-repository").(repository.UserRepository), publisher)
	profileHandler := ctn.Resolve("profile-handler").(accountPB.ProfileServiceHandler)

	// Register Handlers
	accountPB.RegisterUserServiceHandler(service.Server(), userHandler)
	accountPB.RegisterProfileServiceHandler(service.Server(), profileHandler)

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.account", service.Server(), new(subscriber.AccountSubscriber))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.account", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

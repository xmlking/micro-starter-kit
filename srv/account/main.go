package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/util/log"

	"github.com/xmlking/micro-starter-kit/shared/database"
	"github.com/xmlking/micro-starter-kit/srv/account/entity"
	"github.com/xmlking/micro-starter-kit/srv/account/handler"
	accountPB "github.com/xmlking/micro-starter-kit/srv/account/proto/account"
	"github.com/xmlking/micro-starter-kit/srv/account/repository"
	"github.com/xmlking/micro-starter-kit/srv/account/subscriber"

	myConfig "github.com/xmlking/micro-starter-kit/shared/config"
	// _ "github.com/xmlking/micro-starter-kit/shared/log"
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

	// Lets do DI manually
	//db
	db, err := database.GetDatabaseConnection(&cfg.Database)
	if err != nil {
		log.Fatalf("error getting database connection: %v", err)
	}
	defer db.Close()

	db.AutoMigrate(&entity.User{}, &entity.Profile{})

	// Repositories
	userRepository := repository.NewUserRepository(db)
	profileRepository := repository.NewProfileRepository(db)

	// Publisher publish to "go.micro.srv.emailer"
	publisher := micro.NewPublisher("go.micro.srv.emailer", service.Client())

	// Handlers
	userHandler := handler.NewUserHandler(userRepository, publisher)
	profileHandler := handler.NewProfileHandler(profileRepository)

	// Initialise service
	service.Init()

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

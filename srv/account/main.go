package main

import (
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/config"
	log "github.com/sirupsen/logrus"

	"github.com/xmlking/micro-starter-kit/srv/account/handler"
	accountPB "github.com/xmlking/micro-starter-kit/srv/account/proto/account"
	"github.com/xmlking/micro-starter-kit/srv/account/registry"
	"github.com/xmlking/micro-starter-kit/srv/account/repository"

	myConfig "github.com/xmlking/micro-starter-kit/shared/config"
	"github.com/xmlking/micro-starter-kit/shared/wrapper"
)

var (
	configFile string
	cfg        myConfig.ServiceConfiguration
)

// init is called on package initialization and can therefore be used to initialize global stuff like logging, config, ..
func init() {
	config.Scan(&cfg)
}

func main() {

	// New Service
	service := micro.NewService(
		// optional cli flag to override config. comment out if you don't need to override
		micro.Flags(
			cli.StringFlag{
				Name:        "config, c",
				Value:       "config/config.yaml",
				Usage:       "Path to the configuration file to use. Defaults to 'config/config.yaml'",
				EnvVar:      "CONFIG_FILE",
				Destination: &configFile,
			}),
		micro.Name(cfg.ServiceName),
		micro.Version(cfg.Version),
		micro.WrapHandler(wrapper.LogWrapper),
	)

	// Initialise service
	service.Init(
		micro.Action(func(c *cli.Context) {
			// this time, includes flag config overrides
			myConfig.InitConfig(configFile)
			config.Scan(&cfg)
		}),
	)

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

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

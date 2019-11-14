package main

import (
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/service/grpc"

	log "github.com/sirupsen/logrus"

	myConfig "github.com/xmlking/micro-starter-kit/shared/config"
	logger "github.com/xmlking/micro-starter-kit/shared/log"
	logWrapper "github.com/xmlking/micro-starter-kit/shared/wrapper/log"
	"github.com/xmlking/micro-starter-kit/srv/account/handler"
	accountPB "github.com/xmlking/micro-starter-kit/srv/account/proto/account"
	"github.com/xmlking/micro-starter-kit/srv/account/registry"
	"github.com/xmlking/micro-starter-kit/srv/account/repository"
	greeterPB "github.com/xmlking/micro-starter-kit/srv/greeter/proto/greeter"
)

const (
	serviceName = "accountsrv"
)

var (
	configDir  string
	configFile string
	cfg        myConfig.ServiceConfiguration
)

func main() {

	// New Service
	service := grpc.NewService(
		// optional cli flag to override config.
		// comment out if you don't need to override any base config via CLI
		micro.Flags(
			cli.StringFlag{
				Name:        "configDir, d",
				Value:       "/config",
				Usage:       "Path to the config directory. Defaults to 'config'",
				EnvVar:      "CONFIG_DIR",
				Destination: &configDir,
			},
			cli.StringFlag{
				Name:        "configFile, f",
				Value:       "config.yaml",
				Usage:       "Config file in configDir. Defaults to 'config.yaml'",
				EnvVar:      "CONFIG_FILE",
				Destination: &configFile,
			}),
		micro.Name(serviceName),
		micro.Version(myConfig.Version),
		micro.WrapHandler(logWrapper.NewHandlerWrapper()),
		micro.WrapClient(logWrapper.NewClientWrapper()),
	)

	// Initialize service
	service.Init(
		micro.Action(func(c *cli.Context) {
			// load config
			myConfig.InitConfig(configDir, configFile)
			config.Scan(&cfg)
			logger.InitLogger(cfg.Log)
		}),
	)

	// Initialize DI Container
	ctn, err := registry.NewContainer(cfg)
	defer ctn.Clean()
	if err != nil {
		log.Fatalf("failed to build container: %v", err)
	}

	log.Debugf("Client type: grpc or regular? %T\n", service.Client()) // FIXME: expected *grpc.grpcClient but got *micro.clientWrapper

	// Publisher publish to "emailersrv"
	emailerSrvEp := config.Get("emailersrv", "endpoint").String("emailersrv")
	publisher := micro.NewPublisher(emailerSrvEp, service.Client())
	// greeterSrv Client to call "greetersrv"
	greeterSrvEp := config.Get("greetersrv", "endpoint").String("greetersrv")
	greeterSrvClient := greeterPB.NewGreeterService(greeterSrvEp, service.Client())

	// // Handlers
	userHandler := handler.NewUserHandler(ctn.Resolve("user-repository").(repository.UserRepository), publisher, greeterSrvClient)
	profileHandler := ctn.Resolve("profile-handler").(accountPB.ProfileServiceHandler)

	// Register Handlers
	accountPB.RegisterUserServiceHandler(service.Server(), userHandler)
	accountPB.RegisterProfileServiceHandler(service.Server(), profileHandler)

	myConfig.PrintBuildInfo()
	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

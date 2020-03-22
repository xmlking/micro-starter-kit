package main

import (
	"path/filepath"

	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/config"
	"github.com/xmlking/logger/log"

	// "github.com/micro/go-micro/v2/service/grpc"
	"github.com/xmlking/micro-starter-kit/service/account/handler"
	profilePB "github.com/xmlking/micro-starter-kit/service/account/proto/profile"
	userPB "github.com/xmlking/micro-starter-kit/service/account/proto/user"
	"github.com/xmlking/micro-starter-kit/service/account/registry"
	"github.com/xmlking/micro-starter-kit/service/account/repository"
	greeterPB "github.com/xmlking/micro-starter-kit/service/greeter/proto/greeter"
	myConfig "github.com/xmlking/micro-starter-kit/shared/config"
	"github.com/xmlking/micro-starter-kit/shared/constants"
	"github.com/xmlking/micro-starter-kit/shared/logger"
	"github.com/xmlking/micro-starter-kit/shared/util"
	logWrapper "github.com/xmlking/micro-starter-kit/shared/wrapper/log"
	transWrapper "github.com/xmlking/micro-starter-kit/shared/wrapper/transaction"
	validatorWrapper "github.com/xmlking/micro-starter-kit/shared/wrapper/validator"
)

const (
	serviceName = constants.ACCOUNT_SERVICE
	configDir   = "/config"
	configFile  = "config.yaml"
)

var (
	cfg myConfig.ServiceConfiguration
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name(serviceName),
		micro.Version(myConfig.Version),
	)

	// Initialize service
	service.Init(
		micro.Action(func(c *cli.Context) (err error) {
			// load config
			myConfig.InitConfig(configDir, configFile)
			err = config.Scan(&cfg)
			logger.InitLogger(cfg.Log)
			return
		}),
	)

	// Initialize Features
	var options []micro.Option
	if cfg.Features["mtls"].Enabled {
		// if tlsConf, err := util.GetSelfSignedTLSConfig("localhost"); err != nil {
		if tlsConf, err := util.GetTLSConfig(
			filepath.Join(configDir, config.Get("features", "mtls", "certfile").String("")),
			filepath.Join(configDir, config.Get("features", "mtls", "keyfile").String("")),
			filepath.Join(configDir, config.Get("features", "mtls", "cafile").String("")),
			filepath.Join(configDir, config.Get("features", "mtls", "servername").String("")),
		); err != nil {
			log.WithError(err).Error("unable to load certs")
		} else {
			println(tlsConf)
			options = append(options,
				util.WithTLS(tlsConf),
			)
		}
	}
	// Wrappers are invoked in the order as they added
	if cfg.Features["reqlogs"].Enabled {
		options = append(options,
			micro.WrapHandler(logWrapper.NewHandlerWrapper()),
			micro.WrapClient(logWrapper.NewClientWrapper()),
		)
	}
	if cfg.Features["validator"].Enabled {
		options = append(options,
			micro.WrapHandler(validatorWrapper.NewHandlerWrapper()),
		)
	}
	if cfg.Features["translogs"].Enabled {
		topic := config.Get("features", "translogs", "topic").String("mkit.service.recorder")
		publisher := micro.NewEvent(topic, service.Client())
		options = append(options,
			micro.WrapHandler(transWrapper.NewHandlerWrapper(publisher)),
		)
	}

	// Initialize Features
	service.Init(
		options...,
	)

	// Initialize DI Container
	ctn, err := registry.NewContainer(cfg)
	defer ctn.Clean()
	if err != nil {
		log.Fatalf("failed to build container: %v", err)
	}

	log.Debugf("Client type: grpc or regular? %T\n", service.Client()) // FIXME: expected *grpc.grpcClient but got *micro.clientWrapper

	// Publisher publish to "mkit.service.emailer"
	publisher := micro.NewEvent(constants.EMAILER_SERVICE, service.Client())
	// greeterSrv Client to call "mkit.service.greeter"
	greeterSrvClient := greeterPB.NewGreeterService(constants.GREETER_SERVICE, service.Client())

	// // Handlers
	userHandler := handler.NewUserHandler(ctn.Resolve("user-repository").(repository.UserRepository), publisher, greeterSrvClient)
	profileHandler := ctn.Resolve("profile-handler").(profilePB.ProfileServiceHandler)

	// Register Handlers
	userPB.RegisterUserServiceHandler(service.Server(), userHandler)
	profilePB.RegisterProfileServiceHandler(service.Server(), profileHandler)

	myConfig.PrintBuildInfo()
	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

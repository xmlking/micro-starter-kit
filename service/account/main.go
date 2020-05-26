package main

import (
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/config"
	"github.com/rs/zerolog/log"

	// "github.com/micro/go-micro/v2/service/grpc"
	"github.com/xmlking/micro-starter-kit/service/account/handler"
	profilePB "github.com/xmlking/micro-starter-kit/service/account/proto/profile"
	userPB "github.com/xmlking/micro-starter-kit/service/account/proto/user"
	"github.com/xmlking/micro-starter-kit/service/account/registry"
	"github.com/xmlking/micro-starter-kit/service/account/repository"
	greeterPB "github.com/xmlking/micro-starter-kit/service/greeter/proto/greeter"
	myConfig "github.com/xmlking/micro-starter-kit/shared/config"
	"github.com/xmlking/micro-starter-kit/shared/constants"
	_ "github.com/xmlking/micro-starter-kit/shared/logger"
	"github.com/xmlking/micro-starter-kit/shared/util"
	logWrapper "github.com/xmlking/micro-starter-kit/shared/wrapper/log"
	transWrapper "github.com/xmlking/micro-starter-kit/shared/wrapper/transaction"
	validatorWrapper "github.com/xmlking/micro-starter-kit/shared/wrapper/validator"
)

const (
	serviceName = constants.ACCOUNT_SERVICE
)

var (
	cfg = myConfig.GetServiceConfig()
    ff = myConfig.GetFeatureFlags()
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
			// do some life cycle actions
			return
		}),
	)

	// Initialize Features
	var options []micro.Option
	if ff.IsTLSEnabled() {
		if tlsConf, err := myConfig.CreateServerCerts(); err != nil {
			log.Error().Err(err).Msg("unable to load certs")
		} else {
            log.Info().Msg("TLS Enabled")
			options = append(options,
				util.WithTLS(tlsConf),
			)
		}
	}
	// Wrappers are invoked in the order as they added
	if ff.IsReqlogsEnabled() {
		options = append(options,
			micro.WrapHandler(logWrapper.NewHandlerWrapper()),
			micro.WrapClient(logWrapper.NewClientWrapper()),
		)
	}
	if ff.IsValidatorEnabled() {
		options = append(options,
			micro.WrapHandler(validatorWrapper.NewHandlerWrapper()),
		)
	}
	if ff.IsTranslogsEnabled() {
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
        log.Fatal().Msgf("failed to build container: %v", err)
	}

	log.Debug().Msgf("Client type: grpc or regular? %T\n", service.Client()) // FIXME: expected *grpc.grpcClient but got *micro.clientWrapper

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

	println(myConfig.GetBuildInfo())
	// Run service
	if err := service.Run(); err != nil {
        log.Fatal().Err(err).Msg("")
	}
}

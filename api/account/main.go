package main

import (
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/config"

	// "github.com/micro/go-micro/service/grpc"
	log "github.com/sirupsen/logrus"

	"github.com/xmlking/micro-starter-kit/api/account/handler"
	accountPB "github.com/xmlking/micro-starter-kit/api/account/proto/account"
	"github.com/xmlking/micro-starter-kit/shared/wrapper"
	userPB "github.com/xmlking/micro-starter-kit/srv/account/proto/account"

	myConfig "github.com/xmlking/micro-starter-kit/shared/config"
	_ "github.com/xmlking/micro-starter-kit/shared/log"
)

const (
	// serviceName = "go.micro.api.account"
	// TODO: `micro api --handler=api` automatically add `go.micro.api` namespace
	// so I cannot use short serviceName
	serviceName = "account-api"
)

var (
	configDir  string
	configFile string
	cfg        myConfig.ServiceConfiguration
)

func main() {

	// New Service
	service := micro.NewService(
		// service := grpc.NewService(
		// optional cli flag to override config.
		// comment out if you don't need to override any base config via CLI
		micro.Flags(
			cli.StringFlag{
				Name:        "configDir, d",
				Value:       "config",
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
		micro.WrapHandler(wrapper.LogWrapper),
	)

	// Initialize service
	service.Init(
		micro.Action(func(c *cli.Context) {
			// load config
			myConfig.InitConfig(configDir, configFile)
			config.Scan(&cfg)
		}),
	)
	// retry client
	// cli := client.NewClient(
	// 	client.Retries(4),
	// 	client.Retry(func(ctx context.Context, req client.Request, retryCount int, err error) (b bool, e error) {
	// 		if err != nil {
	// 			log.Errorf("[ERR] , err: %s, %v", retryCount, err)
	// 			return true, nil
	// 		}

	// 		return false, nil
	// 	}),
	// )
	// userSrvClient := userPB.NewUserService("account-srv", cli)

	// NOTE: has to give `port` when using with k8s as `registry`
	// userSrvClient := userPB.NewUserService("account:8080", service.Client())
	userSrvClient := userPB.NewUserService("account-srv", service.Client())
	profSrvClient := userPB.NewProfileService("account-srv", service.Client()) // service.Client() or client.DefaultClient???
	accountHandler := handler.NewAccountHandler(userSrvClient, profSrvClient)

	// Register Handler
	accountPB.RegisterAccountServiceHandler(service.Server(), accountHandler)
	// service.Server().Handle(service.Server().NewHandler(accountHandler))

	myConfig.PrintBuildInfo()
	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

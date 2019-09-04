package main

import (
	"github.com/micro/go-micro"
	log "github.com/sirupsen/logrus"
	"github.com/xmlking/micro-starter-kit/api/account/handler"
	_ "github.com/xmlking/micro-starter-kit/shared/config"
	_ "github.com/xmlking/micro-starter-kit/shared/log"

	accountPB "github.com/xmlking/micro-starter-kit/api/account/proto/account"
	myConfig "github.com/xmlking/micro-starter-kit/shared/config"
	userPB "github.com/xmlking/micro-starter-kit/srv/account/proto/account"
)

const (
	serviceName = "go.micro.api.account"
)

func main() {

	// New Service
	service := micro.NewService(
		micro.Name(serviceName),
		micro.Version(myConfig.Version),
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
	// userSrvClient := userPB.NewUserService("go.micro.srv.account", cli)

	// NOTE: has to give `port` when using with k8s as `registry`
	// userSrvClient := userPB.NewUserService("account:8080", service.Client())
	userSrvClient := userPB.NewUserService("go.micro.srv.account", service.Client())
	profSrvClient := userPB.NewProfileService("go.micro.srv.account", service.Client()) // service.Client() or client.DefaultClient???
	accountHandler := handler.NewAccountHandler(userSrvClient, profSrvClient)

	// initialize service
	service.Init()

	// Register Handler
	accountPB.RegisterAccountServiceHandler(service.Server(), accountHandler)
	// service.Server().Handle(service.Server().NewHandler(accountHandler))

	myConfig.PrintBuildInfo()
	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}

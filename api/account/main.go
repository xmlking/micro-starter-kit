package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/util/log"
	"github.com/xmlking/micro-starter-kit/api/account/handler"
	accountPB "github.com/xmlking/micro-starter-kit/api/account/proto/account"
	userPB "github.com/xmlking/micro-starter-kit/srv/account/proto/account"
)

func main() {

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.api.account"),
		micro.Version("latest"),
	)

	userSrvClient := userPB.NewUserService("go.micro.srv.account", client.DefaultClient)
	profSrvClient := userPB.NewProfileService("go.micro.srv.account", client.DefaultClient)
	accountHandler := handler.NewAccountHandler(userSrvClient, profSrvClient)

	// Initialise service
	service.Init()

	// Register Handler
	accountPB.RegisterAccountServiceHandler(service.Server(), accountHandler)

	// Run service
	if err := service.Run(); err != nil {
		log.Log(err)
	}

}

package main

import (
    "context"

    "github.com/micro/go-micro/v2"
    "github.com/micro/go-micro/v2/client"
    "github.com/micro/go-micro/v2/metadata"
    "github.com/rs/zerolog/log"

    "github.com/golang/protobuf/ptypes/wrappers"

    userPB "github.com/xmlking/micro-starter-kit/service/account/proto/user"
    "github.com/xmlking/micro-starter-kit/shared/config"
    "github.com/xmlking/micro-starter-kit/shared/constants"
    _ "github.com/xmlking/micro-starter-kit/shared/logger"
    logWrapper "github.com/xmlking/micro-starter-kit/shared/wrapper/log"
)

var (
    cfg = config.GetConfig()
)

func main() {
    log.Debug().Msgf("IsProduction? %v", config.IsProduction())
    //log.Debug().Interface("Dialect", cfg.Database.Dialect).Send()
    //log.Debug().Msg(cfg.Database.Host)
    //log.Debug().Uint32("Port", cfg.Database.Port).Send()
    //log.Debug().Uint64("FlushInterval", cfg.Features.Tracing.FlushInterval).Send()
    //log.Debug().Msgf("cfg is %v", cfg)

    userService := userPB.NewUserService(constants.ACCOUNT_SERVICE, client.DefaultClient)

    if  _, err := userService.Create(context.TODO(), &userPB.CreateRequest{
        Username:  &wrappers.StringValue{Value: "sumo"},
        FirstName: &wrappers.StringValue{Value: "sumo"},
        LastName:  &wrappers.StringValue{Value: "demo"},
        Email:     &wrappers.StringValue{Value: "sumo@demo.com"},
    });  err != nil {
        log.Fatal().Err(err).Msg("Unable to create User")
    }

    getUserList(userService)
    getUserList2()
}

func getUserList(us userPB.UserService) {
    if  rsp, err := us.List(context.Background(), &userPB.ListRequest{Limit: &wrappers.UInt32Value{Value : 10}});  err != nil {
        log.Fatal().Err(err).Msg("Unable to List Users")
    } else {
        log.Info().Interface("listRsp", rsp).Send()
    }
}

// Just to showcase usage of generic micro client
func getUserList2() {
    service := micro.NewService(
        micro.Name("mkit.client.account"),
        micro.Version(config.Version),
    )
    service.Init(micro.WrapClient(logWrapper.NewClientWrapper())) // Showcase ClientWrapper

    cl := service.Client()

    // Create new request to service mkit.service.account, method UserService.List
    listReq := cl.NewRequest(constants.ACCOUNT_SERVICE, "UserService.List", &userPB.ListRequest{Limit: &wrappers.UInt32Value{Value : 10}})
    listRsp := &userPB.ListResponse{}

    // Create context with metadata - (Optional) Just for demonstration
    ctx := metadata.NewContext(context.Background(), map[string]string{
        "X-User-Id": "john",
        "X-From-Id": "script",
    })

    if err :=  cl.Call(ctx, listReq, listRsp); err != nil {
        log.Fatal().Err(err).Msg("Unable to List Users")
    } else {
        log.Info().Interface("listRsp",listRsp).Send()
    }
}

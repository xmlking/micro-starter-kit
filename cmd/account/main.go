package main

import (
    "context"
    "flag"

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

    username := flag.String("username", "sumo", "username of user to be create")
    email := flag.String("email", "sumo@demo.com", "email of user to be create")
    limit := flag.Uint64("limit", 10, "Limit number of results")
    flag.Parse()

    log.Debug().Str("username", *username).Str("email", *email).Uint64("limit", *limit).Msg("Flags Using:")

    userService := userPB.NewUserService(constants.ACCOUNT_SERVICE, client.DefaultClient)

    if _, err := userService.Create(context.TODO(), &userPB.CreateRequest{
        Username:  &wrappers.StringValue{Value: *username},
        FirstName: &wrappers.StringValue{Value: "sumo"},
        LastName:  &wrappers.StringValue{Value: "demo"},
        Email:     &wrappers.StringValue{Value: *email},
    }); err != nil {
        log.Fatal().Err(err).Msg("Unable to create User")
    }

    getUserList(userService, uint32(*limit))
    getUserList2(uint32(*limit))
}

func getUserList(us userPB.UserService, limit uint32) {
    if rsp, err := us.List(context.Background(), &userPB.ListRequest{Limit: &wrappers.UInt32Value{Value: limit}}); err != nil {
        log.Fatal().Err(err).Msg("Unable to List Users")
    } else {
        log.Info().Interface("listRsp", rsp).Send()
    }
}

// Just to showcase usage of generic micro client
func getUserList2(limit uint32) {

    // New Service
    service := micro.NewService(
        micro.Name(constants.ACCOUNT_CLIENT),
        micro.Version(config.Version),
        micro.WrapClient(logWrapper.NewClientWrapper()), // Showcase ClientWrapper usage
    )

    cl := service.Client()

    // Create new request to service mkit.service.account, method UserService.List
    listReq := cl.NewRequest(constants.ACCOUNT_SERVICE, "UserService.List", &userPB.ListRequest{
        Limit: &wrappers.UInt32Value{Value: limit},
    })
    listRsp := &userPB.ListResponse{}

    // Create context with metadata - (Optional) Just for demonstration
    ctx := metadata.NewContext(context.Background(), map[string]string{
        "X-User-Id": "john",
        "X-From-Id": "script",
    })

    if err := cl.Call(ctx, listReq, listRsp); err != nil {
        log.Fatal().Err(err).Msg("Unable to List Users")
    } else {
        log.Info().Interface("listRsp", listRsp).Send()
    }
}

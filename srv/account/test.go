package main

//https://github.com/uknth/go-base/blob/master/base/log/log.go

import (
	"fmt"
	"flag"
	
	"github.com/micro/go-micro"
	"github.com/micro/cli"
	"github.com/micro/go-micro/util/log"
	goConf "github.com/micro/go-micro/config"

 	"github.com/xmlking/micro-starter-kit/common/config"
)

var (
	srvCfg        config.ServiceConfiguration
	// baseLogger *logrus.Logger
)

// init is called on package initialization and can therefore be used to initialize global stuff like logging, config, ..
func init() {
	flag.String("database_host", "127.0.0.1", "the db host")
	flag.Int("database_port", 3306, "the db port")
	flag.Parse()
	config.InitConfigWithFlag()
	goConf.Scan(&srvCfg)
	// baseLogger = initLogging(cfg.Log)
}

func main() {
		// New Service
		service := micro.NewService(
			micro.Name("go.micro.srv.account"),
			micro.Version("latest"),
			micro.Flags(
			cli.StringFlag{
				Name:   "database_host",
				EnvVar: "DATABASE_POST",
				Usage:  " Invalid!!!",
			},
			cli.IntFlag{
				Name:   "database_port",
				EnvVar: "DATABASE_PORT",
				Usage:  " Invalid!!!",
			}),
		)
	
		// Initialise service
		service.Init()

		fmt.Println(goConf.Get("database", "dialect").String("postgres"))
		fmt.Println(goConf.Get("database", "host").String("no-address"))
		fmt.Println(goConf.Get("database", "port").Int(0000))
		fmt.Println(goConf.Get("observability", "tracing", "flushInterval").Int(2000000000))
		fmt.Println(srvCfg)

		// Run service
		if err := service.Run(); err != nil {
			log.Fatal(err)
		}
}
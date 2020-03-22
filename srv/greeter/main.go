package main

import (
	"path/filepath"

	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/config"
	"github.com/xmlking/logger/log"

	myConfig "github.com/xmlking/micro-starter-kit/shared/config"
	"github.com/xmlking/micro-starter-kit/shared/logger"
	"github.com/xmlking/micro-starter-kit/shared/util"
	logWrapper "github.com/xmlking/micro-starter-kit/shared/wrapper/log"
	transWrapper "github.com/xmlking/micro-starter-kit/shared/wrapper/transaction"
	"github.com/xmlking/micro-starter-kit/srv/greeter/handler"
	greeterPB "github.com/xmlking/micro-starter-kit/srv/greeter/proto/greeter"
)

const (
	serviceName = "greetersrv"
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
		options = append(options, micro.WrapHandler(logWrapper.NewHandlerWrapper()))
	}
	if cfg.Features["translogs"].Enabled {
		topic := config.Get("features", "translogs", "topic").String("recordersrv")
		publisher := micro.NewEvent(topic, service.Client())
		options = append(options, micro.WrapHandler(transWrapper.NewHandlerWrapper(publisher)))
	}

	// Initialize Features
	service.Init(
		options...,
	)

	// Register Handler
	_ = greeterPB.RegisterGreeterServiceHandler(service.Server(), handler.NewGreeterHandler())
	myConfig.PrintBuildInfo()
	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

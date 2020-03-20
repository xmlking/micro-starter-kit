package main

import (
	"crypto/tls"

	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	gc "github.com/micro/go-micro/v2/client/grpc"
	"github.com/micro/go-micro/v2/config"
	gs "github.com/micro/go-micro/v2/server/grpc"
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
)

var (
	configDir  string
	configFile string
	cfg        myConfig.ServiceConfiguration
)

func main() {
	// New Service
	service := micro.NewService(
		// optional cli flag to override config.
		// comment out if you don't need to override any base config via CLI
		micro.Flags(
			&cli.StringFlag{
				Name:        "configDir",
				Aliases:     []string{"d"},
				Value:       "/config",
				Usage:       "Path to the config directory. Defaults to 'config'",
				EnvVars:     []string{"CONFIG_DIR"},
				Destination: &configDir,
			},
			&cli.StringFlag{
				Name:        "configFile",
				Aliases:     []string{"f"},
				Value:       "config.yaml",
				Usage:       "Config file in configDir. Defaults to 'config.yaml'",
				EnvVars:     []string{"CONFIG_FILE"},
				Destination: &configFile,
			}),
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
		if tlsConf, err := util.GetSelfSignedTLSConfig("localhost"); err != nil {
			log.WithError(err).Error("unable to load certs")
		} else {
			options = append(options,
				// https://github.com/ykumar-rb/ZTP/blob/master/pnp/server.go
				WithTLS(tlsConf),
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
func WithTLS(t *tls.Config) micro.Option {
	return func(o *micro.Options) {
		o.Client.Init(
			gc.AuthTLS(t),
		)
		o.Server.Init(
			gs.AuthTLS(t),
		)
	}
}

package log

// TODO https://github.com/uknth/go-base/blob/master/base/log/log.go

import (
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/util/log"
	"github.com/sirupsen/logrus"
	micrologrus "github.com/tudurom/micro-logrus"
)

func init() {
	level := config.Get("log", "level").String("info")
	format := config.Get("log", "format").String("text")
	golog := logrus.New()

	if level == "debug" {
		golog.SetLevel(logrus.DebugLevel)
	}

	if format == "text" {
		golog.SetFormatter(&logrus.TextFormatter{
			ForceColors:   true,
			FullTimestamp: true,
		})
	} else {
		golog.SetFormatter(&logrus.JSONFormatter{})
	}

	log.SetLogger(micrologrus.NewMicroLogrus(golog))
}

package log

// TODO https://github.com/uknth/go-base/blob/master/base/log/log.go
// TODO https://github.com/pmker/vash/blob/master/common/log/log.go

import (
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/util/log"
	"github.com/sirupsen/logrus"
	micrologrus "github.com/tudurom/micro-logrus"
)

// log level: panic, fatal, error, warn, info, debug, trace
// log format: json, text
func init() {
	level := config.Get("log", "level").String("info")
	format := config.Get("log", "format").String("text")
	golog := logrus.New()

	logLevel, err := logrus.ParseLevel(level)
	if err != nil {
		logLevel = logrus.InfoLevel
	}
	golog.SetLevel(logLevel)

	if format == "json" {
		golog.SetFormatter(&logrus.JSONFormatter{})
	} else {
		golog.SetFormatter(&logrus.TextFormatter{
			ForceColors:   true,
			FullTimestamp: true,
		})
	}

	golog.WithFields(logrus.Fields{
		"level": logLevel,
		"format": format,
	  }).Debug("Logger set to Logrus with:")
 
	log.SetLogger(micrologrus.NewMicroLogrus(golog))
}

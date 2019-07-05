package log

import (
	"github.com/micro/go-micro/config"
	microlog "github.com/micro/go-micro/util/log"
	"github.com/sirupsen/logrus"
	micrologrus "github.com/tudurom/micro-logrus"
)

func init() {
	logrusLogger := newLogger()
	// also set same Formatter and Level for logrus's global logger
	logrus.SetFormatter(logrusLogger.Formatter)
	logrus.SetLevel(logrusLogger.Level)

	microLogger := micrologrus.NewMicroLogrus(logrusLogger)
	microlog.SetLogger(microLogger)
}

// newLogger create new logrus logger from config
// log level: panic, fatal, error, warn, info, debug, trace
// log format: json, text
func newLogger() *logrus.Logger {
	level := config.Get("log", "level").String("info")
	format := config.Get("log", "format").String("text")
	lslog := logrus.New()

	logLevel, err := logrus.ParseLevel(level)
	if err != nil {
		logLevel = logrus.InfoLevel
	}
	lslog.SetLevel(logLevel)

	if format == "json" {
		lslog.SetFormatter(&logrus.JSONFormatter{})
	} else {
		lslog.SetFormatter(&logrus.TextFormatter{
			ForceColors:   true,
			FullTimestamp: true,
		})
	}

	lslog.WithFields(logrus.Fields{
		"logLevel": logLevel,
		"format":   format,
	}).Info("Logger set to Logrus with:")

	return lslog
}

// NewLogger create new logrus logger from config and return FieldLogger interface
// log level: panic, fatal, error, warn, info, debug, trace
// log format: json, text
func NewLogger() logrus.FieldLogger {
	return newLogger()
}

package log

import (
	"os"

	microlog "github.com/micro/go-micro/util/log"
	"github.com/sirupsen/logrus"
	micrologrus "github.com/tudurom/micro-logrus"
	"github.com/xmlking/micro-starter-kit/shared/config"
)

func InitLogger(cfg config.LogConfiguration) {
	logrusLogger := newLogger(cfg)
	// also set same Formatter and Level for logrus's global logger
	logrus.SetFormatter(logrusLogger.Formatter)
	logrus.SetLevel(logrusLogger.Level)

	microLogger := micrologrus.NewMicroLogrus(logrusLogger)
	microlog.SetLogger(microLogger)
	// also set same log_level for go-micro
	// TODO: microlog.SetLevel(microlog.LevelDebug)
	os.Setenv("MICRO_LOG_LEVEL", logrusLogger.GetLevel().String())

	logrus.WithFields(logrus.Fields{
		"logLevel": cfg.Level,
		"format":   cfg.Format,
	}).Info("Logger set to Logrus with:")
}

// newLogger create new logrus logger from config
// log level: panic, fatal, error, warn, info, debug, trace
// log format: json, text
func newLogger(cfg config.LogConfiguration) *logrus.Logger {
	level := cfg.Level
	format := cfg.Format
	logger := logrus.New()

	logLevel, err := logrus.ParseLevel(level)
	if err != nil {
		logLevel = logrus.InfoLevel
	}
	logger.SetLevel(logLevel)

	if format == "json" {
		logger.SetFormatter(&logrus.JSONFormatter{})
	} else {
		logger.SetFormatter(&logrus.TextFormatter{
			ForceColors:   true,
			FullTimestamp: true,
		})
	}

	return logger
}

// NewLogger create new logrus logger from config and return FieldLogger interface
// log level: panic, fatal, error, warn, info, debug, trace
// log format: json, text
func NewLogger(cfg config.LogConfiguration) logrus.FieldLogger {
	return newLogger(cfg)
}

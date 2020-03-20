package logger

import (
	"github.com/xmlking/logger"
	"github.com/xmlking/logger/log"
	"github.com/xmlking/logger/zerolog"

	"github.com/xmlking/micro-starter-kit/shared/config"
)

func InitLogger(logConf config.LogConfiguration) {
	// set as default logger
	logger.DefaultLogger = newLogger(logConf)
	log.WithFields(map[string]interface{}{
		"logLevel": logConf.Level,
		"runtime":  logConf.Runtime,
	}).Info("Logger set to Zerolog with:")
}

// newLogger create new logger from config
// log level: panic, fatal, error, warn, info, debug, trace
func newLogger(cfg config.LogConfiguration) (mLogger logger.Logger) {
	level := cfg.Level
	runtime := cfg.Runtime
	logLevel, err := logger.ParseLevel(level)
	if err != nil {
		panic(err)
	}

	if runtime == "gcp" {
		mLogger = zerolog.NewLogger(
			logger.WithLevel(logLevel),
			zerolog.UseAsDefault(),
			zerolog.WithGCPMode(),
		)
	} else if config.IsProduction {
		mLogger = zerolog.NewLogger(
			logger.WithLevel(logLevel),
			zerolog.UseAsDefault(),
			zerolog.WithProductionMode(),
		)
	} else {
		mLogger = zerolog.NewLogger(
			logger.WithLevel(logLevel),
			zerolog.UseAsDefault(),
			zerolog.WithDevelopmentMode(),
		)
	}
	return
}

// NewLogger create new logger from config and return Logger interface
// log level: panic, fatal, error, warn, info, debug, trace
// log runtime: dev, prod, gcp, azure, aws
func NewLogger(cfg config.LogConfiguration) logger.Logger {
	return newLogger(cfg)
}
func NewLoggerWithFields(cfg config.LogConfiguration, fields map[string]interface{}) logger.Logger {
	logr := newLogger(cfg)
	_ = logr.Init(logger.WithFields(fields))
	return logr
}

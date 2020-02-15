package log

import (
	// "github.com/micro/go-micro/v2/logger"
	"github.com/xmlking/micro-starter-kit/shared/config"
	"github.com/xmlking/micro-starter-kit/shared/micro/logger"
	zero "github.com/xmlking/micro-starter-kit/shared/micro/logger/zerolog"
)

var Logger logger.Logger

func InitLogger(cfg config.LogConfiguration) {
	mLogger := newLogger(cfg)
	logger.Register(mLogger)
	mLogger.Fields(map[string]interface{}{
		"logLevel": cfg.Level,
		"format":   cfg.Format,
	}).Log(logger.InfoLevel, "Logger set to Zerolog with:")
}

// newLogger create new logger from config
// log level: panic, fatal, error, warn, info, debug, trace
func newLogger(cfg config.LogConfiguration) (mLogger logger.Logger) {
	level := cfg.Level
	logLevel, err := zero.ParseLevel(level)
	if err != nil {
		logLevel = logger.InfoLevel
	}

	if config.IsProduction {
		mLogger = zero.NewLogger(
			zero.WithLevel(logLevel),
			zero.UseAsDefault(),
			zero.WithProductionMode(),
		)
	} else {
		mLogger = zero.NewLogger(
			zero.WithLevel(logLevel),
			zero.UseAsDefault(),
			zero.WithDevelopmentMode(),
		)
	}
	return
}

// NewLogger create new logger from config and return Logger interface
// log level: panic, fatal, error, warn, info, debug, trace
func NewLogger(cfg config.LogConfiguration) logger.Logger {
	return newLogger(cfg)
}

func Info(msg string) {
	Logger.Log(logger.InfoLevel, msg)
}
func Infof(format string, args ...interface{}) {
	Logger.Logf(logger.InfoLevel, format, args...)
}

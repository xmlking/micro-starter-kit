package log

import (
	ml "github.com/micro/go-micro/v2/logger"
	"github.com/xmlking/micro-starter-kit/shared/config"
	zero "github.com/xmlking/micro-starter-kit/shared/micro/logger/zerolog"
)

var Logger ml.Logger

func InitLogger(cfg config.LogConfiguration) {
	logger := newLogger(cfg)
	ml.Register(logger)
	logger.Fields([]ml.Field{
		{Key: "logLevel", Value: cfg.Level},
		{Key: "format", Value: cfg.Format},
	}...).Log(ml.InfoLevel, "Logger set to Zerolog with:")
}

// newLogger create new logger from config
// log level: panic, fatal, error, warn, info, debug, trace
func newLogger(cfg config.LogConfiguration) (mLogger ml.Logger) {
	level := cfg.Level
	logLevel, err := zero.ParseLevel(level)
	if err != nil {
		logLevel = ml.InfoLevel
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
func NewLogger(cfg config.LogConfiguration) ml.Logger {
	return newLogger(cfg)
}

func Info(msg string) {
	Logger.Log(ml.InfoLevel, msg)
}
func Infof(format string, args ...interface{}) {
	Logger.Logf(ml.InfoLevel, format, args...)
}

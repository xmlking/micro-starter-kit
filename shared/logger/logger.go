package logger

import (
	"time"

	"github.com/micro/go-micro/v2/logger"
	zlog "github.com/micro/go-plugins/logger/zerolog/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"github.com/xmlking/micro-starter-kit/shared/config"
	log "github.com/xmlking/micro-starter-kit/shared/micro/logger"
	"github.com/xmlking/micro-starter-kit/shared/micro/logger/stackdriver"
)

func InitLogger(cfg config.LogConfiguration) {
	mLogger := newLogger(cfg)
	// Register with micro logger
	logger.Register(mLogger)
	// set as default logger
	log.Logger = mLogger
	mLogger.Fields(map[string]interface{}{
		"logLevel": cfg.Level,
		"runtime":  cfg.Runtime,
	}).Log(logger.InfoLevel, "Logger set to Zerolog with:")
}

// newLogger create new logger from config
// log level: panic, fatal, error, warn, info, debug, trace
func newLogger(cfg config.LogConfiguration) (mLogger logger.Logger) {
	level := cfg.Level
	runtime := cfg.Runtime
	logLevel, err := logger.GetLevel(level)
	if err != nil {
		panic(err)
	}

	if runtime == "gcp" {
		zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
		zerolog.LevelFieldName = "severity"
		zerolog.TimestampFieldName = "timestamp"
		mLogger = zlog.NewLogger(
			zlog.WithLevel(logLevel),
			zlog.UseAsDefault(),
			zlog.WithProductionMode(),
			zlog.WithHooks([]zerolog.Hook{stackdriver.StackdriverSeverityHook{}}),
			zlog.WithTimeFormat(time.RFC3339Nano),
		)
	} else if config.IsProduction {
		mLogger = zlog.NewLogger(
			zlog.WithLevel(logLevel),
			zlog.UseAsDefault(),
			zlog.WithProductionMode(),
		)
	} else {
		mLogger = zlog.NewLogger(
			zlog.WithLevel(logLevel),
			zlog.UseAsDefault(),
			zlog.WithDevelopmentMode(),
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

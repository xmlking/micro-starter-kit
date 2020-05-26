package logger_test

import (
    "testing"

    "github.com/rs/zerolog/log"
    // bootstrap config and logger
    "github.com/xmlking/micro-starter-kit/shared/logger"
)

func TestLogger(t *testing.T) {

	log.Info().Msg("Hello World")
	log.Info().Msgf("Hello %s", "Sumo")
}

func TestWithGcp(t *testing.T) {
    // logger.Init(logger.WithFormat())
    logger.Init(logger.WithTimeFormat("ddd"))
	log.Info().Msgf("testing: %s", "WithGcp")
	// reset `LevelFieldName` to make other tests pass.
	// logger.Init(logger.WithFields())
	log.Info().Msgf("testing: %s", "WithDevelopment")
}

// TODO: test Sub-loggers
// Sub-loggers let you chain default logger with additional context
// With creates a child logger with the field added to its context.
//sublogger := log.With().
//Str("component", "foo").
//Logger()
//sublogger.Info().Msg("hello world")

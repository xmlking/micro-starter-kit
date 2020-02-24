package logger_test

import (
    "testing"

    "github.com/xmlking/logger/log"

    "github.com/xmlking/micro-starter-kit/shared/config"
    "github.com/xmlking/micro-starter-kit/shared/logger"
)

func TestLogger(t *testing.T) {
    logger.InitLogger(config.LogConfiguration{Level: "debug", Runtime: "development"})
    log.Info("Hello World")
    log.Infof("Hello %s", "Sumo")
}

func TestWithGcp(t *testing.T) {
    logger.InitLogger(config.LogConfiguration{Level: "debug", Runtime: "gcp"})
    log.Infof("testing: %s", "WithGcp")
    // reset `LevelFieldName` to make other tests pass.
    logger.InitLogger(config.LogConfiguration{Level: "debug", Runtime: "development"})
    log.Infof("testing: %s", "WithDevelopment")
}

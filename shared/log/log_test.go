package log_test

import (
	"testing"

	// "github.com/micro/go-micro/v2/logger"
	"github.com/xmlking/micro-starter-kit/shared/config"
	"github.com/xmlking/micro-starter-kit/shared/log"
	"github.com/xmlking/micro-starter-kit/shared/micro/logger"
)

func TestName(t *testing.T) {
	log.Logger = log.NewLogger(config.LogConfiguration{Level: "debug", Format: "json"})
	log.Info("Hello World")
	log.Infof("Hello %s", "Sumo")
	// log.Logger.Logf(ml.InfoLevel, "name: %s", "SUMO")
	t.Logf("testing logger name: %s", log.Logger.String())
}

func TestRegister(t *testing.T) {
	log.InitLogger(config.LogConfiguration{Level: "debug", Format: "json"})
	logr, _ := logger.GetLogger("zerolog")
	logr.Log(logger.InfoLevel, "Hello World")
	logr.Logf(logger.InfoLevel, "Hello %s", "Sumo")
	t.Logf("testing logger name: %s", logr.String())
}

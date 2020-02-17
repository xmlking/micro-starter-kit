package logger_test

import (
	"os"
	"testing"

	mlog "github.com/micro/go-micro/v2/logger"
	zlog "github.com/micro/go-plugins/logger/zerolog/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"github.com/xmlking/micro-starter-kit/shared/config"
	"github.com/xmlking/micro-starter-kit/shared/logger"
	log "github.com/xmlking/micro-starter-kit/shared/micro/logger"
	"github.com/xmlking/micro-starter-kit/shared/micro/logger/stackdriver"
)

func TestLogger(t *testing.T) {
	logger.InitLogger(config.LogConfiguration{Level: "debug", Runtime: "development"})
	log.Info("Hello World")
	log.Infof("Hello %s", "Sumo")
}

func TestRegistation(t *testing.T) {
	logger.InitLogger(config.LogConfiguration{Level: "debug", Runtime: "development"})
	logr, _ := mlog.GetLogger("zerolog")
	logr.Log(mlog.InfoLevel, "Hello World")
	logr.Logf(mlog.InfoLevel, "Hello %s", "Sumo")
}

func TestWithGcp(t *testing.T) {
	logger.InitLogger(config.LogConfiguration{Level: "debug", Runtime: "gcp"})
	log.Infof("testing: %s", "WithGcp")
	// reset `LevelFieldName` to make other tests pass.
	logger.InitLogger(config.LogConfiguration{Level: "debug", Runtime: "development"})
	log.Infof("testing: %s", "WithDevelopment")
}

func ExampleInitLogger() {
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	zerolog.LevelFieldName = "severity"
	zerolog.TimestampFieldName = "timestamp"
	mLogger := zlog.NewLogger(
		zlog.WithOut(os.Stdout),
		zlog.UseAsDefault(),
		zlog.WithProductionMode(),
		zlog.WithHooks([]zerolog.Hook{stackdriver.StackdriverSeverityHook{}}),
		zlog.WithTimeFormat("aaa"),
	)
	mlog.Register(mLogger)
	logr, _ := mlog.GetLogger("zerolog")
	logr.Log(mlog.InfoLevel, "Hello World")
	// Output:
	// {"severity":"info","timestamp":"aaa","severity":"Info","message":"Hello World"}
}

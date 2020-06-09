package logger_test

import (
    "fmt"
    "os"
    "testing"
    "time"

    "github.com/pkg/errors"
    "github.com/rs/zerolog"
    "github.com/rs/zerolog/log"
    // bootstrap logger
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

//func ExampleWithOut() {
//    logger.Init(
//       logger.WithOutput(os.Stdout),
//       logger.WithFormat("json"),
//       logger.WithTimeFormat("ddd"),
//    )
//    log.Info().Msg("testing: Info")
//    log.Info().Msgf("testing: %s", "Info with Format")
//    log.Info().Msgf("testing: %s", "Info with Fields")
//    log.Warn().Fields(map[string]interface{}{
//        "sumo":  "demo",
//        "human": true,
//        "age":   99,
//    }).Msgf("testing: %s", "Warn with Fields")
//    // Output:
//    //{"level":"info","LogLevel":"","LogFormat":"json","time":"ddd","message":"Logger set to Zerolog with:"}
//    //{"level":"info","time":"ddd","message":"testing: Info"}
//    //{"level":"info","time":"ddd","message":"testing: Info with Format"}
//    //{"level":"info","time":"ddd","message":"testing: Info with Fields"}
//    //{"level":"warn","age":99,"human":true,"sumo":"demo","time":"ddd","message":"testing: Warn with Fields"}
//}

func TestSetLevel(t *testing.T) {
    logger.Init(logger.WithLevel(zerolog.DebugLevel))
    log.Debug().Msgf("test show debug: %s", "debug msg")

    logger.Init(logger.WithLevel(zerolog.InfoLevel))
    log.Debug().Msgf("test non-show  debug: %s", "debug msg")
}

func TestReportCaller(t *testing.T) {
    logger.Init(logger.ReportCaller())

    log.Info().Msgf("testing: %s", "ReportCaller")
}

func TestReportCallerWithDevelopmentMode(t *testing.T) {
    logger.Init(logger.ReportCaller(), logger.WithFormat("pretty"))

    log.Info().Msgf("testing: %s", "ReportCallerWithDevelopmentMode")
}

func TestWithOutput(t *testing.T) {
    logger.Init(logger.WithOutput(os.Stdout))

    log.Info().Msgf("testing: %s", "WithOutput")
}

func TestWithGCPMode(t *testing.T) {
    logger.Init(logger.ReportCaller(), logger.WithFormat("gcp"))

    log.Info().Msgf("testing: %s", "TestWithGCPMode Infof")
    log.Info().Fields(map[string]interface{}{
        "sumo":  "demo",
        "human": true,
        "age":   99,
    }).Msgf("testing: TestWithGCPMode Infow")

    logger.Init(logger.ReportCaller())
    log.Error().Err(fmt.Errorf("Error %v: %w", "nested", errors.New("root error message"))).Msg("TestWithGCPModeAndWithError")
    logger.Init(logger.WithTimeFormat(time.RFC3339Nano))

    log.Info().Msgf("testing: %s", "TestWithGCPMode")
    // reset `LevelFieldName` to make other tests pass.
    t.Cleanup(func() {
        logger.Init(logger.WithFormat("json"))
    })
}

func TestWithDevelopmentMode(t *testing.T) {
    logger.Init(logger.WithFormat("pretty"))

    log.Info().Msgf("testing: %s", "DevelopmentMode")
}

func TestSubLoggerWithMoreFields(t *testing.T) {
    logger.Init(logger.WithFields(map[string]interface{}{
        "component": "AccountHandler",
    }))

    log.Debug().Fields(map[string]interface{}{
        "name":  "demo",
        "human": true,
        "age":   77,
    }).Msg("testing: Infow with extra fields")
    log.Info().Msgf("testing: %s", "Infof with default fields")
    // Output:
    //{"level":"info","component":"AccountHandler","age":77,"human":true,"name":"demo","time":"2020-02-23T12:01:10-08:00","message":"testing: Infow with extra fields"}
    //{"level":"info","component":"AccountHandler","time":"2020-02-23T12:01:10-08:00","message":"testing: Infof with default fields"}
}

func TestWithError(t *testing.T) {
    log.Error().Msgf("testing: %s", "TestWithError")
    log.Error().Err(fmt.Errorf("Error %v: %w", "nested", errors.New("root error message"))).Msg("TestWithError")
}

func TestWithErrorAndDefaultFields(t *testing.T) {
    logger.Init(
        logger.WithFields(map[string]interface{}{
            "name":  "sumo",
            "age":   99,
            "alive": true,
        }))
    log.Error().Msg("TestWithErrorAndDefaultFields")
    log.Error().Msgf("testing: %s", "TestWithErrorAndDefaultFields")
    err := errors.Wrap(errors.New("error message"), "from TestWithErrorAndDefaultFields")
    log.Error().Err(err).Msgf("testing: %s", "WithErrorAndDefaultFields")
}

func TestLogStack(t *testing.T) {
    logger.Init(logger.WithFormat("pretty"))
    err := errors.Wrap(errors.New("error message"), "from TestLogStack")
    log.Error().Err(err).Msgf("testing: %s", "LogStack")
}

func TestWithHooks(t *testing.T) {
    simpleHook := zerolog.HookFunc(func(e *zerolog.Event, level zerolog.Level, msg string) {
        e.Bool("has_level", level != zerolog.NoLevel)
        e.Str("test", "logged")
    })

    log.Logger.Hook(simpleHook)

    log.Info().Msgf("testing: %s", "WithHooks")
}

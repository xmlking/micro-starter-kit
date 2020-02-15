package zerolog

import (
	"errors"
	"os"
	"testing"
	"time"

	// "github.com/micro/go-micro/v2/logger"
	"github.com/rs/zerolog"
	"github.com/xmlking/micro-starter-kit/shared/micro/logger"
)

func TestName(t *testing.T) {
	l := NewLogger()

	if l.String() != "zerolog" {
		t.Errorf("error: name expected 'zerolog' actual: %s", l.String())
	}

	t.Logf("testing logger name: %s", l.String())
}

func ExampleWithOut() {
	l := NewLogger(WithOut(os.Stdout), WithProductionMode())

	l.Logf(logger.InfoLevel, "testing: %s", "logf")

	// Output:
	// {"level":"info","message":"testing: logf"}
}

func TestSetLevel(t *testing.T) {
	l := NewLogger()

	l.SetLevel(logger.DebugLevel)
	l.Logf(logger.DebugLevel, "test show debug: %s", "debug msg")

	l.SetLevel(logger.InfoLevel)
	l.Logf(logger.DebugLevel, "test non-show debug: %s", "debug msg")
}

func TestWithReportCaller(t *testing.T) {
	l := NewLogger(ReportCaller())

	l.Logf(logger.InfoLevel, "testing: %s", "WithReportCaller")
}
func TestWithOut(t *testing.T) {
	l := NewLogger(WithOut(os.Stdout))

	l.Logf(logger.InfoLevel, "testing: %s", "WithOut")
}

func TestWithDevelopmentMode(t *testing.T) {
	l := NewLogger(WithDevelopmentMode(), WithTimeFormat(time.Kitchen))

	l.Logf(logger.InfoLevel, "testing: %s", "DevelopmentMode")
}
func TestWithLevelFieldName(t *testing.T) {
	l := NewLogger(WithGCPMode())

	l.Logf(logger.InfoLevel, "testing: %s", "WithLevelFieldName")
	// reset `LevelFieldName` to make other tests pass.
	NewLogger(WithProductionMode())
}

func TestWithFields(t *testing.T) {
	l := NewLogger()

	l.Fields(map[string]interface{}{
		"sumo":  "demo",
		"human": true,
		"age":   99,
	}).Logf(logger.InfoLevel, "testing: %s", "WithFields")
}

func TestWithError(t *testing.T) {
	l := NewLogger()

	l.Error(errors.New("I am Error")).Logf(logger.ErrorLevel, "testing: %s", "WithError")
}

func TestWithHooks(t *testing.T) {
	simpleHook := zerolog.HookFunc(func(e *zerolog.Event, level zerolog.Level, msg string) {
		e.Bool("has_level", level != zerolog.NoLevel)
		e.Str("test", "logged")
	})

	l := NewLogger(WithHooks([]zerolog.Hook{simpleHook}))

	l.Logf(logger.InfoLevel, "testing: %s", "WithHooks")
}

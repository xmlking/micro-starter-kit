# GCP

set log level name to `Severity` for __GCP__ `StackDriver`

## Usage

```go
import (
	"os"
	"testing"

	mlog "github.com/micro/go-micro/v2/logger"
	zlog "github.com/micro/go-plugins/logger/zerolog/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"github.com/xmlking/micro-starter-kit/shared/config"
	"github.com/xmlking/micro-starter-kit/shared/micro/logger/stackdriver"
)

func ExampleInitLogger_withGcp() {
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
	logr.Log(logger.InfoLevel, "Hello World")
	// Output:
	// {"severity":"info","timestamp":"aaa","severity":"Info","message":"Hello World"}
}
```

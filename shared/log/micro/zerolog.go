package microzerolog

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/rs/zerolog"
)

type microZeroLogger struct {
	l zerolog.Logger
}

func parentCaller() string {
	pc, _, _, ok := runtime.Caller(4)
	fn := runtime.FuncForPC(pc)
	if ok && fn != nil {
		return fn.Name()
	}
	return ""
}

func NewMicroZeroLogger(logger zerolog.Logger) microZeroLogger {
	return microZeroLogger{l: logger}
}

func (mzl microZeroLogger) Log(args ...interface{}) {
	pc := parentCaller()
	if strings.HasSuffix(pc, "Fatal") {
		mzl.l.Fatal().Msg(fmt.Sprint(args...))
	} else {
		mzl.l.Info().Msg(fmt.Sprint(args...))
	}
}

func (mzl microZeroLogger) Logf(format string, args ...interface{}) {
	pc := parentCaller()
	if strings.HasSuffix(pc, "Fatalf") {
		mzl.l.Fatal().Msg(fmt.Sprintf(format, args...))
	} else {
		mzl.l.Info().Msg(fmt.Sprintf(format, args...))
	}
}

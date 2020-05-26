package gcp

import (
    "fmt"
    "runtime"
    "strings"

    "github.com/rs/zerolog"
)

const (
    // CallerSkipFrameCount is the number of stack frames to skip to find the caller.
    CallerSkipFrameCount = 3
)

type StackdriverSeverityHook struct{}

func (h StackdriverSeverityHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
    e.Str("severity", LevelToSeverity(level))
}

// LevelToSeverity converts a zerolog level to the stackdriver severity
// Stackdriver has more levels than zerolog so we skip some severities.
// By default we set info when no level is provided.
var LevelToSeverity = func(level zerolog.Level) string {
    switch level {
    case zerolog.NoLevel:
        return "Default"
    case zerolog.DebugLevel:
        return "Debug"
    // Let info falls into the default
    case zerolog.WarnLevel:
        return "Warning"
    case zerolog.ErrorLevel:
        return "Error"
    case zerolog.FatalLevel:
        return "Alert"
    case zerolog.PanicLevel:
        return "Emergency"
    default:
        return "Info"
    }
}

// callerHook implements zerolog.Hook interface.
type CallerHook struct{}

// Run adds sourceLocation for the log to zerolog.Event.
func (h CallerHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
    var file, line, function string
    if pc, filePath, lineNum, ok := runtime.Caller(CallerSkipFrameCount); ok {
        if f := runtime.FuncForPC(pc); f != nil {
            function = f.Name()
        }
        line = fmt.Sprintf("%d", lineNum)
        parts := strings.Split(filePath, "/")
        file = parts[len(parts)-1]
    }
    e.Dict("logging.googleapis.com/sourceLocation",
        zerolog.Dict().Str("file", file).Str("line", line).Str("function", function))
}

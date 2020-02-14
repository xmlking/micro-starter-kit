package zero

import (
	"github.com/rs/zerolog"
)

type stackdriverSeverityHook struct{}

func (h stackdriverSeverityHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
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

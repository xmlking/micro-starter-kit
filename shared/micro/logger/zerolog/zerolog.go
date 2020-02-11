package zero

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/micro/go-micro/v2/logger"
	"github.com/rs/zerolog"
)

var (
	out   io.Writer = os.Stderr
	color           = false
	exit            = os.Exit
)

type zeroLogger struct {
	nativelogger zerolog.Logger
}

func (l *zeroLogger) Fields(fields ...logger.Field) logger.Logger {
	data := make(map[string]interface{}, len(fields))
	for _, f := range fields {
		data[f.Key] = f.GetValue()
	}
	return &zeroLogger{l.nativelogger.With().Fields(data).Logger()}
}

func (l *zeroLogger) Init(opts ...logger.Option) error {

	options := &Options{logger.Options{Context: context.Background()}}
	for _, o := range opts {
		o(&options.Options)
	}

	// Prepare Writer
	if useColor, ok := options.Context.Value(useColorKey{}).(bool); ok {
		color = useColor
	}
	if o, ok := options.Context.Value(outKey{}).(io.Writer); ok {
		out = o
	}
	if pretty, ok := options.Context.Value(prettyKey{}).(bool); ok && pretty {
		out = zerolog.NewConsoleWriter(
			func(w *zerolog.ConsoleWriter) {
				w.TimeFormat = time.RFC3339
				w.Out = out
				w.NoColor = !color
			},
		)
	}
	l.nativelogger = l.nativelogger.Output(out)

	if level, ok := options.Context.Value(levelKey{}).(logger.Level); ok {
		//zerolog.SetGlobalLevel(loggerToZerologLevel(level))
		l.nativelogger = l.nativelogger.Level(loggerToZerologLevel(level))
	} else {
		l.nativelogger = l.nativelogger.Level(zerolog.InfoLevel)
	}

	if caller, ok := options.Context.Value(reportCallerKey{}).(bool); ok && caller {
		l.nativelogger = l.nativelogger.With().Caller().Logger()
	}

	if levelFieldName, ok := options.Context.Value(levelFieldKey{}).(string); ok {
		zerolog.LevelFieldName = levelFieldName
	}

	if hooks, ok := options.Context.Value(hooksKey{}).([]zerolog.Hook); ok {
		for _, hook := range hooks {
			l.nativelogger = l.nativelogger.Hook(hook)
		}
	}

	if exitFunc, ok := options.Context.Value(exitKey{}).(func(int)); ok {
		exit = exitFunc
	}
	return nil
}

func (l *zeroLogger) SetLevel(level logger.Level) {
	//zerolog.SetGlobalLevel(loggerToZerologLevel(level))
	l.nativelogger = l.nativelogger.Level(loggerToZerologLevel(level))
}

func (l *zeroLogger) Level() logger.Level {
	return ZerologToLoggerLevel(l.nativelogger.GetLevel())
}

func (l *zeroLogger) Log(level logger.Level, args ...interface{}) {
	msg := fmt.Sprintf("%s", args)
	l.nativelogger.WithLevel(loggerToZerologLevel(level)).Msg(msg)
	// Invoke os.Exit because unlike zerolog.Logger.Fatal zerolog.Logger.WithLevel won't stop the execution.
	if level == logger.FatalLevel {
		exit(1)
	}
}

func (l *zeroLogger) Logf(level logger.Level, format string, args ...interface{}) {
	l.nativelogger.WithLevel(loggerToZerologLevel(level)).Msgf(format, args...)
	// Invoke os.Exit because unlike zerolog.Logger.Fatal zerolog.Logger.WithLevel won't stop the execution.
	if level == logger.FatalLevel {
		exit(1)
	}
}

func (l *zeroLogger) String() string {
	return "zerolog"
}

// NewLogger builds a new logger based on options
func NewLogger(opts ...logger.Option) logger.Logger {
	l := &zeroLogger{}
	_ = l.Init(opts...)
	return l
}

func loggerToZerologLevel(level logger.Level) zerolog.Level {
	switch level {
	case logger.TraceLevel:
		return zerolog.TraceLevel
	case logger.DebugLevel:
		return zerolog.DebugLevel
	case logger.InfoLevel:
		return zerolog.InfoLevel
	case logger.WarnLevel:
		return zerolog.WarnLevel
	case logger.ErrorLevel:
		return zerolog.ErrorLevel
	case logger.PanicLevel:
		return zerolog.PanicLevel
	case logger.FatalLevel:
		return zerolog.FatalLevel
	default:
		return zerolog.InfoLevel
	}
}

func ZerologToLoggerLevel(level zerolog.Level) logger.Level {
	switch level {
	case zerolog.TraceLevel:
		return logger.TraceLevel
	case zerolog.DebugLevel:
		return logger.DebugLevel
	case zerolog.InfoLevel:
		return logger.InfoLevel
	case zerolog.WarnLevel:
		return logger.WarnLevel
	case zerolog.ErrorLevel:
		return logger.ErrorLevel
	case zerolog.PanicLevel:
		return logger.PanicLevel
	case zerolog.FatalLevel:
		return logger.FatalLevel
	default:
		return logger.InfoLevel
	}
}

package zero

import (
	"context"
	"fmt"
	"io"
	"os"
	"runtime/debug"
	"time"

	"github.com/micro/go-micro/v2/logger"
	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

type Mode uint8

const (
	Production Mode = iota + 1
	Development
	GCP
)

var (
	//  It's common to set this to a file, or leave it default which is `os.Stderr`
	out io.Writer = os.Stderr
	// Function to exit the application, defaults to `os.Exit()`
	exitFunc = os.Exit
	// Flag for whether to log caller info (off by default)
	reportCaller = false
	// use this logger as system wide default logger
	useAsDefault = false
	// The logging level the logger should log at.
	// This defaults to 100 means not explicitly set by user
	level      logger.Level = 100
	fields     []logger.Field
	hooks      []zerolog.Hook
	timeFormat string
	// default Production (1)
	mode Mode = Production
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

func (l *zeroLogger) Error(err error) logger.Logger {
	return &zeroLogger{
		l.nativelogger.With().Fields(map[string]interface{}{zerolog.ErrorFieldName: err}).Logger(),
	}
}

func (l *zeroLogger) Init(opts ...logger.Option) error {

	options := &Options{logger.Options{Context: context.Background()}}
	for _, o := range opts {
		o(&options.Options)
	}

	if o, ok := options.Context.Value(outKey{}).(io.Writer); ok {
		out = o
	}
	if hs, ok := options.Context.Value(hooksKey{}).([]zerolog.Hook); ok {
		hooks = hs
	}
	if flds, ok := options.Context.Value(fieldsKey{}).([]logger.Field); ok {
		fields = flds
	}
	if lvl, ok := options.Context.Value(levelKey{}).(logger.Level); ok {
		level = lvl
	}
	if tf, ok := options.Context.Value(timeFormatKey{}).(string); ok {
		timeFormat = tf
	}
	if exitFunction, ok := options.Context.Value(exitKey{}).(func(int)); ok {
		exitFunc = exitFunction
	}
	if caller, ok := options.Context.Value(reportCallerKey{}).(bool); ok && caller {
		reportCaller = caller
	}
	if useDefault, ok := options.Context.Value(useAsDefaultKey{}).(bool); ok && useDefault {
		useAsDefault = useDefault
	}
	if devMode, ok := options.Context.Value(developmentModeKey{}).(bool); ok && devMode {
		mode = Development
	}
	if prodMode, ok := options.Context.Value(productionModeKey{}).(bool); ok && prodMode {
		mode = Production
	}
	if gcpMode, ok := options.Context.Value(gcpModeKey{}).(bool); ok && gcpMode {
		mode = GCP
	}

	switch mode {
	case Development:
		zerolog.ErrorStackMarshaler = func(err error) interface{} {
			fmt.Println(string(debug.Stack()))
			return nil
		}
		consOut := zerolog.NewConsoleWriter(
			func(w *zerolog.ConsoleWriter) {
				w.TimeFormat = time.RFC3339
				w.Out = out
				w.NoColor = false
			},
		)
		level = logger.DebugLevel
		l.nativelogger = zerolog.New(consOut).
			Level(zerolog.DebugLevel).
			With().Timestamp().Stack().Logger()
	case GCP:
		zerolog.LevelFieldName = "severity"
		zerolog.TimestampFieldName = "timestamp"
		timeFormat = time.RFC3339Nano
		// Adds stackdriver severity hook
		hooks = append(hooks, stackdriverSeverityHook{})

		l.nativelogger = zerolog.New(out).
			Level(zerolog.InfoLevel).
			With().Timestamp().Stack().Logger()
	default: // Production
		zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
		l.nativelogger = zerolog.New(out).
			Level(zerolog.InfoLevel).
			With().Timestamp().Stack().Logger()
	}

	// Change  Writer if not default
	if out != os.Stderr {
		l.nativelogger = l.nativelogger.Output(out)
	}

	// Set log Level if not default
	if level != 100 {
		//zerolog.SetGlobalLevel(loggerToZerologLevel(level))
		l.nativelogger = l.nativelogger.Level(loggerToZerologLevel(level))
	}

	// Adding hooks if exist
	if reportCaller {
		l.nativelogger = l.nativelogger.With().Caller().Logger()
	}
	for _, hook := range hooks {
		l.nativelogger = l.nativelogger.Hook(hook)
	}

	// Setting timeFormat
	if len(timeFormat) > 0 {
		zerolog.TimeFieldFormat = timeFormat
	}

	// Adding seed fields if exist
	if fields != nil {
		data := make(map[string]interface{}, len(fields))
		for _, f := range fields {
			data[f.Key] = f.GetValue()
		}
		l.nativelogger = l.nativelogger.With().Fields(data).Logger()
	}

	// Also set it as Default zerolog logger
	if useAsDefault {
		zlog.Logger = l.nativelogger
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
	l.nativelogger.WithLevel(loggerToZerologLevel(level)).Msg(msg[1 : len(msg)-1])
	// Invoke os.Exit because unlike zerolog.Logger.Fatal zerolog.Logger.WithLevel won't stop the execution.
	if level == logger.FatalLevel {
		exitFunc(1)
	}
}

func (l *zeroLogger) Logf(level logger.Level, format string, args ...interface{}) {
	l.nativelogger.WithLevel(loggerToZerologLevel(level)).Msgf(format, args...)
	// Invoke os.Exit because unlike zerolog.Logger.Fatal zerolog.Logger.WithLevel won't stop the execution.
	if level == logger.FatalLevel {
		exitFunc(1)
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

// ParseLevel converts a level string into a logger Level value.
// returns an error if the input string does not match known values.
func ParseLevel(levelStr string) (lvl logger.Level, err error) {
	if zLevel, err := zerolog.ParseLevel(levelStr); err == nil {
		return ZerologToLoggerLevel(zLevel), err
	} else {
		return lvl, fmt.Errorf("Unknown Level String: '%s' %w", levelStr, err)
	}
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

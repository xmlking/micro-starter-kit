package logger

import (
    "context"
    "fmt"
    "os"
    "runtime/debug"
    "time"

    "github.com/rs/zerolog"
    "github.com/rs/zerolog/log"
    "github.com/rs/zerolog/pkgerrors"

    mLogger "github.com/micro/go-micro/v2/logger"

    "github.com/xmlking/micro-starter-kit/shared/config"
    "github.com/xmlking/micro-starter-kit/shared/logger/gcp"
    zeroToMicroAdopter "github.com/xmlking/micro-starter-kit/shared/logger/micro"
)

var (
    // Default Logger
    DefaultLogger Logger = NewLogger()
)

type Logger interface {
    Init(options ...Option) error
    Options() Options
    String() string
}

type defaultLogger struct {
    opts Options
}

func (l *defaultLogger) Init(opts ...Option) error {
    for _, o := range opts {
        o(&l.opts)
    }

    // Reset to zerolog defaults
    zerolog.TimeFieldFormat = time.RFC3339
    zerolog.ErrorStackMarshaler = nil
    zerolog.LevelFieldName = "level"
    zerolog.TimestampFieldName = "time"
    zerolog.LevelFieldMarshalFunc = func(l zerolog.Level) string { return l.String() }

    var logr zerolog.Logger

    if l.opts.Format == config.GCP { // Only GCP Mode implemented

        zerolog.TimeFieldFormat = time.RFC3339Nano
        zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
        zerolog.LevelFieldName = "severity"
        // logr.Hook(gcp.StackdriverSeverityHook{})
        zerolog.TimestampFieldName = "timestamp"
        zerolog.LevelFieldMarshalFunc = gcp.LevelToSeverity

        logr = zerolog.New(l.opts.Out).
            Level(zerolog.InfoLevel).
            With().Timestamp().Stack().Logger()

    } else if l.opts.Format == config.JSON || config.IsProduction() { // Production Mode

        zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
        logr = zerolog.New(l.opts.Out).
            Level(zerolog.InfoLevel).
            With().Timestamp().Stack().Logger()

    } else { // Default  Development Mode

        zerolog.ErrorStackMarshaler = func(err error) interface{} {
            fmt.Println(string(debug.Stack()))
            return nil
        }
        consOut := zerolog.NewConsoleWriter(
            func(w *zerolog.ConsoleWriter) {
                if len(l.opts.TimeFormat) > 0 {
                    w.TimeFormat = l.opts.TimeFormat
                }
                w.Out = l.opts.Out
                w.NoColor = false
            },
        )
        logr = zerolog.New(consOut).
            Level(zerolog.DebugLevel).
            With().Timestamp().Stack().Logger()

    }

    // Set log Level if not default
    if l.opts.Level != zerolog.NoLevel {
        zerolog.SetGlobalLevel(l.opts.Level)
        logr = logr.Level(l.opts.Level)
    }

    // Adding ReportCaller hook
    if l.opts.ReportCaller {
        if l.opts.Format == config.GCP {
            logr.Hook(gcp.CallerHook{})
        } else {
            logr = logr.With().Caller().Logger()
        }
    }

    // Setting timeFormat
    if len(l.opts.TimeFormat) > 0 {
        zerolog.TimeFieldFormat = l.opts.TimeFormat
    }

    // Adding seed fields if exist
    if l.opts.Fields != nil {
        logr = logr.With().Fields(l.opts.Fields).Logger()
    }

    // Also set it as zerolog's Default logger
    log.Logger = logr

    // Also set it as micro's Default logger
    mLogger.DefaultLogger = zeroToMicroAdopter.Convert(logr)

    log.Info().
        Str("LogLevel", l.opts.Level.String()).
        Str("LogFormat", string(l.opts.Format)).
        Msg("Logger set to Zerolog with:")

    return nil
}

func (l *defaultLogger) Options() Options {
    return l.opts
}

func (l *defaultLogger) String() string {
    return "default"
}

func NewLogger(opts ...Option) Logger {
    logCfg := config.GetServiceConfig().Log
    level, err := logCfg.LogLevel()
    if err != nil {
        log.Err(err).Msg("")
    }

    // Set default options
    options := Options{
        Level:   level,
        Format:  logCfg.LogFormat(),
        Out:     os.Stderr,
        Context: context.Background(),
    }

    l := &defaultLogger{opts: options}
    _ = l.Init(opts...)
    return l
}

// Helper functions on DefaultLogger
func Init(options ...Option) error {
    return DefaultLogger.Init(options...)
}

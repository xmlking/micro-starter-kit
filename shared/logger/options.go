package logger

import (
    "context"

    "github.com/rs/zerolog"

    "github.com/xmlking/micro-starter-kit/shared/config"
)


type Option func(*Options)

type Options struct {
	// The logging level the logger should log at. default is `InfoLevel`
    Level zerolog.Level
    // Log format. default `json`
    Format config.Format
    // TimeFormat is one of time.RFC3339, time.RFC3339Nano, time.*
    TimeFormat string
    // Flag for whether to log caller info (off by default)
    ReportCaller bool
    // fields to always be logged
    Fields map[string]interface{}
	// Alternative options
	Context context.Context
}

// WithLevel set default level for the logger
func WithLevel(level zerolog.Level) Option {
	return func(args *Options) {
		args.Level = level
	}
}

// WithFormat set default log format for the logger
func WithFormat(format config.Format) Option {
    return func(args *Options) {
        args.Format = format
    }
}

// WithTimeFormat set default timeFormat for the logger
func WithTimeFormat(timeFormat string) Option {
    return func(args *Options) {
        args.TimeFormat = timeFormat
    }
}

// ReportCaller set value to `true`
func ReportCaller() Option {
    return func(args *Options) {
        args.ReportCaller = true
    }
}

// WithFields set default fields for the logger
func WithFields(fields map[string]interface{}) Option {
    return func(args *Options) {
        args.Fields = fields
    }
}

func SetOption(k, v interface{}) Option {
	return func(o *Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, k, v)
	}
}

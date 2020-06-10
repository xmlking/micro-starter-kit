package logger

import (
    "context"
    "fmt"
    "io"

    "github.com/rs/zerolog"
)


// Log format enum.
type Format string
const (
    PRETTY Format = "pretty"
    JSON Format = "json"
    GCP Format = "gcp"
    AZURE Format = "azure"
    AWS Format = "aws"
)
func ParseFormat(formatStr string) (Format, error)  {
    switch formatStr {
    case "pretty":
        return PRETTY, nil
    case "json":
        return JSON, nil
    case "gcp":
        return GCP, nil
    case "azure":
        return AZURE, nil
    case "aws":
        return AWS, nil
    }
    return JSON, fmt.Errorf("unknown log Format string: '%s', defaulting to JSON", formatStr)
}

type Option func(*Options)

type Options struct {
	// The logging level the logger should log at. default is `InfoLevel`
    Level zerolog.Level
    // Log format. default `json`
    Format Format
    // TimeFormat is one of time.RFC3339, time.RFC3339Nano, time.*
    TimeFormat string
    // Flag for whether to log caller info (off by default)
    ReportCaller bool
    // fields to always be logged
    Fields map[string]interface{}
    // It's common to set this to a file, or leave it default which is `os.Stderr`
    Out io.Writer
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
func WithFormat(format Format) Option {
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

// WithOutput set default output writer for the logger
func WithOutput(out io.Writer) Option {
    return func(args *Options) {
        args.Out = out
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

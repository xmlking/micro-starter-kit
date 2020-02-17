package log

import (
	"github.com/micro/go-micro/v2/logger"
	zlog "github.com/micro/go-plugins/logger/zerolog/v2"
)

// Logger is the global logger.
var Logger logger.Logger = zlog.NewLogger(zlog.WithDevelopmentMode())

func Log(level logger.Level, args ...interface{}) {
	Logger.Log(level, args)
}
func Logf(level logger.Level, format string, args ...interface{}) {
	Logger.Logf(level, format, args...)
}

func Info(args ...interface{}) {
	Logger.Log(logger.InfoLevel, args)
}
func Infof(format string, args ...interface{}) {
	Logger.Logf(logger.InfoLevel, format, args...)
}

func Trace(args ...interface{}) {
	Logger.Log(logger.TraceLevel, args)
}
func Tracef(format string, args ...interface{}) {
	Logger.Logf(logger.TraceLevel, format, args...)
}

func Debug(args ...interface{}) {
	Logger.Log(logger.DebugLevel, args)
}
func Debugf(format string, args ...interface{}) {
	Logger.Logf(logger.DebugLevel, format, args...)
}

func Warn(args ...interface{}) {
	Logger.Log(logger.WarnLevel, args)
}
func Warnf(format string, args ...interface{}) {
	Logger.Logf(logger.WarnLevel, format, args...)
}

func WithError(err error, args ...interface{}) {
	Logger.Error(err).Log(logger.ErrorLevel, args)
}
func WithErrorf(err error, format string, args ...interface{}) {
	Logger.Error(err).Logf(logger.ErrorLevel, format, args...)
}

func Error(args ...interface{}) {
	Logger.Log(logger.ErrorLevel, args)
}
func Errorf(format string, args ...interface{}) {
	Logger.Logf(logger.ErrorLevel, format, args...)
}

func Panic(args ...interface{}) {
	Logger.Log(logger.PanicLevel, args)
}
func Panicf(format string, args ...interface{}) {
	Logger.Logf(logger.PanicLevel, format, args...)
}

func Fatal(args ...interface{}) {
	Logger.Log(logger.FatalLevel, args)
}
func Fatalf(format string, args ...interface{}) {
	Logger.Logf(logger.FatalLevel, format, args...)
}

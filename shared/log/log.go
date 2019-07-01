package log

// TODO https://github.com/uknth/go-base/blob/master/base/log/log.go
// TODO https://github.com/pmker/vash/blob/master/common/log/log.go

import (
	"runtime"
	"strings"

	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/util/log"
	"github.com/sirupsen/logrus"
)

// Log is shared
var Log Logger

func init() {
	Log = NewLogger()
	log.SetLogger(Log)
}

func parentCaller() string {
	pc, _, _, ok := runtime.Caller(4)
	fn := runtime.FuncForPC(pc)
	if ok && fn != nil {
		return fn.Name()
	}
	return ""
}

// Logger interface for micro
type Logger interface {
	Log(v ...interface{})
	Logf(format string, v ...interface{})

	WithField(key string, value interface{}) *logrus.Entry
	WithFields(fields logrus.Fields) *logrus.Entry
	WithError(err error) *logrus.Entry

	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Printf(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Warningf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Panicf(format string, args ...interface{})

	Debug(args ...interface{})
	Info(args ...interface{})
	Print(args ...interface{})
	Warn(args ...interface{})
	Warning(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	Panic(args ...interface{})
}

type logger struct {
	logrus.Logger
}

func (l *logger) Log(v ...interface{}) {
	pc := parentCaller()
	if strings.HasSuffix(pc, "Fatal") {
		l.Fatal(v...)
	} else {
		l.Info(v...)
	}
}

func (l *logger) Logf(format string, v ...interface{}) {
	pc := parentCaller()
	if strings.HasSuffix(pc, "Fatalf") {
		l.Fatalf(format, v...)
	} else {
		l.Infof(format, v...)
	}
}

// NewLogger create new logrus logger from config
// log level: panic, fatal, error, warn, info, debug, trace
// log format: json, text
func NewLogger() Logger {
	level := config.Get("log", "level").String("info")
	format := config.Get("log", "format").String("text")
	lslog := logrus.New()

	logLevel, err := logrus.ParseLevel(level)
	if err != nil {
		logLevel = logrus.InfoLevel
	}
	lslog.SetLevel(logLevel)

	if format == "json" {
		lslog.SetFormatter(&logrus.JSONFormatter{})
	} else {
		lslog.SetFormatter(&logrus.TextFormatter{
			ForceColors:   true,
			FullTimestamp: true,
		})
	}

	lslog.WithFields(logrus.Fields{
		"level":  logLevel,
		"format": format,
	}).Debug("Logger set to Logrus with:")

	return &logger{*lslog}
}

package grpc

import (
    "fmt"

    "github.com/rs/zerolog"
)

type Logger struct {
    log zerolog.Logger
}

func New(log zerolog.Logger) Logger {
    return Logger{log: log}
}

func (l Logger) Fatal(args ...interface{}) {
    l.log.Fatal().Msg(fmt.Sprint(args...))
}

func (l Logger) Fatalf(format string, args ...interface{}) {
    l.log.Fatal().Msg(fmt.Sprintf(format, args...))
}

func (l Logger) Fatalln(args ...interface{}) {
    l.Fatal(args...)
}

func (l Logger) Error(args ...interface{}) {
    l.log.Error().Msg(fmt.Sprint(args...))
}

func (l Logger) Errorf(format string, args ...interface{}) {
    l.log.Error().Msg(fmt.Sprintf(format, args...))
}

func (l Logger) Errorln(args ...interface{}) {
    l.Error(args...)
}

func (l Logger) Info(args ...interface{}) {
    l.log.Info().Msg(fmt.Sprint(args...))
}

func (l Logger) Infof(format string, args ...interface{}) {
    l.log.Info().Msg(fmt.Sprintf(format, args...))
}

func (l Logger) Infoln(args ...interface{}) {
    l.Info(args...)
}

func (l Logger) Warning(args ...interface{}) {
    l.log.Warn().Msg(fmt.Sprint(args...))
}

func (l Logger) Warningf(format string, args ...interface{}) {
    l.log.Warn().Msg(fmt.Sprintf(format, args...))
}

func (l Logger) Warningln(args ...interface{}) {
    l.Warning(args...)
}

func (l Logger) Print(args ...interface{}) {
    l.log.Info().Msg(fmt.Sprint(args...))
}

func (l Logger) Printf(format string, args ...interface{}) {
    l.log.Info().Msg(fmt.Sprintf(format, args...))
}

func (l Logger) Println(args ...interface{}) {
    l.Print(args...)
}

func (l Logger) V(level int) bool {
    return true
}

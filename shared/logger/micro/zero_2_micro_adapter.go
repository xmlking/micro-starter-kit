package micro

import (
    "fmt"

    "github.com/micro/go-micro/v2/logger"
    "github.com/rs/zerolog"
)

type zeroLogger struct {
    zLog zerolog.Logger
    opts logger.Options
}

func (l *zeroLogger) Init(opts ...logger.Option) error {
    return nil
}

func (l *zeroLogger) Fields(fields map[string]interface{}) logger.Logger {
    l.zLog = l.zLog.With().Fields(fields).Logger()
    return l
}

func (l *zeroLogger) Error(err error) logger.Logger {
    l.zLog = l.zLog.With().Fields(map[string]interface{}{zerolog.ErrorFieldName: err}).Logger()
    return l
}

func (l *zeroLogger) Log(level logger.Level, args ...interface{}) {
    msg := fmt.Sprint(args...)
    l.zLog.WithLevel(loggerToZerologLevel(level)).Msg(msg)

}

func (l *zeroLogger) Logf(level logger.Level, format string, args ...interface{}) {
    l.zLog.WithLevel(loggerToZerologLevel(level)).Msgf(format, args...)
}

func (l *zeroLogger) String() string {
    return "zerolog"
}

func (l *zeroLogger) Options() logger.Options {
    return l.opts
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
    case logger.FatalLevel:
        return zerolog.FatalLevel
    default:
        return zerolog.InfoLevel
    }
}

func Convert(zLogger zerolog.Logger) logger.Logger {
    return  &zeroLogger{zLog: zLogger , opts:  logger.Options{}}
}

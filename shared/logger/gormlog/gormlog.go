// Package gormlog provides gorm logger implementation using Go-Micro's meta logger.
//
// Example usage:
//  orm, _ := gorm.Open("postgres", dsn)
//  orm.LogMode(true)
//  orm.SetLogger(gormlog.NewGormLogger(gormlog.WithLevel(logger.DebugLevel)))
package gormlog

import (
    "database/sql/driver"
    "fmt"
    "reflect"
    "strings"
    "time"
    "unicode"

    "github.com/rs/zerolog"
)

// Logger is a gorm logger implementation using zap.
type GormLogger struct {
    logger      zerolog.Logger
    level       zerolog.Level
    encoderFunc RecordToFields
}

// LoggerOption is an option for Logger.
type GormLoggerOption func(*GormLogger)

// WithLevel returns Logger option that sets level for gorm logs.
// It affects only general logs, e.g. those that contain SQL queries.
// Errors will be logged with error level independently of this option.
func WithLevel(level zerolog.Level) GormLoggerOption {
    return func(l *GormLogger) {
        l.level = level
    }
}

// WithRecordToFields returns Logger option that sets RecordToFields func which
// encodes log Record to a slice of micro logger fields.
//
// This can be used to control field names or field values types.
func WithRecordToFields(f RecordToFields) GormLoggerOption {
    return func(l *GormLogger) {
        l.encoderFunc = f
    }
}

// New returns a new gorm logger implemented using zap.
// By default it logs with debug level.
func NewGormLogger(logger zerolog.Logger, opts ...GormLoggerOption) *GormLogger {
    l := &GormLogger{
        logger:      logger,
        level:       zerolog.DebugLevel,
        encoderFunc: DefaultRecordToFields,
    }

    for _, o := range opts {
        o(l)
    }

    return l
}

// Print implements gorm's logger interface.
func (l *GormLogger) Print(values ...interface{}) {
    rec := l.newRecord(values...)
    l.logger.WithLevel(rec.Level).Fields(l.encoderFunc(rec)).Msg(rec.Message)
}

func (l *GormLogger) newRecord(values ...interface{}) Record {
    // See https://github.com/jinzhu/gorm/blob/master/main.go#L774
    // for info how gorm logs messages.

    if len(values) < 2 {
        // Should this ever happen?
        return Record{
            Message: fmt.Sprint(values...),
            Level:   l.level,
        }
    }

    // Handle https://github.com/jinzhu/gorm/blob/32455088f24d6b1e9a502fb8e40fdc16139dbea8/main.go#L716
    if len(values) == 2 {
        return Record{
            Message: fmt.Sprintf("%v", values[1]),
            Source:  fmt.Sprintf("%v", values[0]),
            Level:   zerolog.ErrorLevel,
        }
    }

    level := values[0]

    // Handle https://github.com/jinzhu/gorm/blob/32455088f24d6b1e9a502fb8e40fdc16139dbea8/main.go#L778
    if level == "log" {
        // By default, assume this is a user log.
        // See: https://github.com/jinzhu/gorm/blob/32455088f24d6b1e9a502fb8e40fdc16139dbea8/scope.go#L96
        // If this is an error log, we set level to error.
        // See: https://github.com/jinzhu/gorm/blob/32455088f24d6b1e9a502fb8e40fdc16139dbea8/main.go#L718
        logLevel := l.level
        if _, ok := values[2].(error); ok {
            logLevel = zerolog.ErrorLevel
        }

        return Record{
            Message: fmt.Sprint(values[2:]...),
            Source:  fmt.Sprintf("%v", values[1]),
            Level:   logLevel,
        }
    }

    // Handle https://github.com/jinzhu/gorm/blob/32455088f24d6b1e9a502fb8e40fdc16139dbea8/main.go#L786
    if level == "sql" {
        return Record{
            Message:      "gorm query",
            Source:       fmt.Sprintf("%v", values[1]),
            Duration:     values[2].(time.Duration),
            SQL:          formatSQL(values[3].(string), values[4].([]interface{})),
            RowsAffected: values[5].(int64),
            Level:        l.level,
        }
    }

    // Should this ever happen?
    return Record{
        Message: fmt.Sprint(values[2:]...),
        Source:  fmt.Sprintf("%v", values[1]),
        Level:   l.level,
    }
}

func formatSQL(sql string, values []interface{}) string {
    size := len(values)

    replacements := make([]string, size*2)

    var indexFunc func(int) string
    if strings.Contains(sql, "$1") {
        indexFunc = formatNumbered
    } else {
        indexFunc = formatQuestioned
    }

    for i := size - 1; i >= 0; i-- {
        replacements[(size-i-1)*2] = indexFunc(i)
        replacements[(size-i-1)*2+1] = formatValue(values[i])
    }

    r := strings.NewReplacer(replacements...)
    return r.Replace(sql)
}

func formatNumbered(index int) string {
    return fmt.Sprintf("$%d", index+1)
}

func formatQuestioned(index int) string {
    return "?"
}

func formatValue(value interface{}) string {
    indirectValue := reflect.Indirect(reflect.ValueOf(value))
    if !indirectValue.IsValid() {
        return "NULL"
    }

    value = indirectValue.Interface()

    switch v := value.(type) {
    case time.Time:
        return fmt.Sprintf("'%v'", v.Format("2006-01-02 15:04:05"))
    case []byte:
        s := string(v)
        if isPrintable(s) {
            return redactLong(fmt.Sprintf("'%s'", s))
        }
        return "'<binary>'"
    case int, int8, int16, int32, int64,
        uint, uint8, uint16, uint32, uint64:
        return fmt.Sprintf("%d", v)
    case driver.Valuer:
        if dv, err := v.Value(); err == nil && dv != nil {
            return formatValue(dv)
        }
        return "NULL"
    default:
        return redactLong(fmt.Sprintf("'%v'", value))
    }
}

func isPrintable(s string) bool {
    for _, r := range s {
        if !unicode.IsPrint(r) {
            return false
        }
    }
    return true
}

func redactLong(s string) string {
    if len(s) > maxLen {
        return "'<redacted>'"
    }
    return s
}

const maxLen = 255

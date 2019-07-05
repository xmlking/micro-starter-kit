package log

import (
	"database/sql/driver"
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"time"
	"unicode"

	"github.com/sirupsen/logrus"
)

var (
	sqlRegexp                = regexp.MustCompile(`\?`)
	numericPlaceHolderRegexp = regexp.MustCompile(`\$\d+`)
)

// GormLogger struct
type GormLogger struct {
	logrus logrus.FieldLogger
}

// NewGormLogger Logger
func NewGormLogger(l logrus.FieldLogger) *GormLogger {
	return &GormLogger{l}
}

// Print for Logger
func (l *GormLogger) Print(v ...interface{}) {
	message, fields := Formatter(v...)
	if v[0] == "sql" {
		l.logrus.WithFields(fields).WithFields(logrus.Fields{"type": "sql"}).Debug(message)
	}
	if v[0] == "log" {
		l.logrus.WithFields(fields).WithFields(logrus.Fields{"type": "log"}).Debug(message)
	}
}

// Formatter for Logger
var Formatter = func(values ...interface{}) (message interface{}, fields logrus.Fields) {
	if len(values) < 1 {
		return
	}
	fields = logrus.Fields{
		"_source": formatSource(values[1].(string)),
	}

	if values[0] == "sql" {
		formattedValues := []string{}

		for _, value := range values[4].([]interface{}) {
			indirectValue := reflect.Indirect(reflect.ValueOf(value))
			if indirectValue.IsValid() {
				value = indirectValue.Interface()
				if t, ok := value.(time.Time); ok {
					formattedValues = append(formattedValues, fmt.Sprintf("'%v'", t.Format("2006-01-02 15:04:05")))
				} else if b, ok := value.([]byte); ok {
					if str := string(b); isPrintable(str) {
						formattedValues = append(formattedValues, fmt.Sprintf("'%v'", str))
					} else {
						formattedValues = append(formattedValues, "'<binary>'")
					}
				} else if r, ok := value.(driver.Valuer); ok {
					if value, err := r.Value(); err == nil && value != nil {
						formattedValues = append(formattedValues, fmt.Sprintf("'%v'", value))
					} else {
						formattedValues = append(formattedValues, "NULL")
					}
				} else {
					formattedValues = append(formattedValues, fmt.Sprintf("'%v'", value))
				}
			} else {
				formattedValues = append(formattedValues, "NULL")
			}
		}

		sql := ""
		s := values[3].(string)
		// differentiate between $n placeholders or else treat like ?
		if numericPlaceHolderRegexp.MatchString(s) {
			sql = s
			for index, value := range formattedValues {
				placeholder := fmt.Sprintf(`\$%d([^\d]|$)`, index+1)
				sql = regexp.MustCompile(placeholder).ReplaceAllString(sql, value+"$1")
			}
		} else {
			formattedValuesLength := len(formattedValues)
			for index, value := range sqlRegexp.Split(s, -1) {
				sql += value
				if index < formattedValuesLength {
					sql += formattedValues[index]
				}
			}
		}

		message = sql
	} else {
		message = values[2]
	}

	return
}

func formatSource(source string) string {
	i := strings.LastIndex(source, "/")
	i = strings.LastIndex(source[:i], "/")

	return source[i+1:]
}

func isPrintable(s string) bool {
	for _, r := range s {
		if !unicode.IsPrint(r) {
			return false
		}
	}
	return true
}

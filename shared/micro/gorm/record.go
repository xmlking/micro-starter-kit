package gormlog

import (
	"time"

	ml "github.com/micro/go-micro/v2/logger"
)

// Record is gormlog log record.
type Record struct {
	Message string
	Source  string
	Level   ml.Level

	Duration     time.Duration
	SQL          string
	RowsAffected int64
}

// RecordToFields func can encode gormlog Record into a slice of zap fields.
type RecordToFields func(r Record) []ml.Field

// DefaultRecordToFields is default encoder func for gormzap log records.
func DefaultRecordToFields(r Record) []ml.Field {
	// Note that Level field is ignored here, because it is handled outside
	// by micro logger itself.

	if r.SQL != "" {
		return []ml.Field{
			{Key: "sql.source", Type: ml.StringType, Value: r.Source},
			{Key: "sql.duration", Type: ml.DurationType, Value: r.Duration},
			{Key: "sql.query", Type: ml.StringType, Value: r.SQL},
			{Key: "sql.rows_affected", Type: ml.Int64Type, Value: r.RowsAffected},
		}
	}

	return []ml.Field{{Key: "sql.source", Type: ml.StringType, Value: r.Source}}
}

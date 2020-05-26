package gormlog

import (
    "time"

    "github.com/rs/zerolog"
)

// Record is gormlog log record.
type Record struct {
    Message string
    Source  string
    Level   zerolog.Level

    Duration     time.Duration
    SQL          string
    RowsAffected int64
}

// RecordToFields func can encode gormlog Record into a slice of zap fields.
type RecordToFields func(r Record) map[string]interface{}

// DefaultRecordToFields is default encoder func for gormzap log records.
func DefaultRecordToFields(r Record) map[string]interface{} {
    // Note that Level field is ignored here, because it is handled outside
    // by micro logger itself.

    if r.SQL != "" {
        return map[string]interface{}{
            "source":        r.Source,
            "duration":      r.Duration,
            "query":         r.SQL,
            "rows_affected": r.RowsAffected,
        }
    }

    return map[string]interface{}{"source": r.Source}
}

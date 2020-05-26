package gormlog_test

import (
    "testing"
    "time"

    "github.com/rs/zerolog"
    "github.com/rs/zerolog/log"

    "github.com/xmlking/micro-starter-kit/shared/logger"
    "github.com/xmlking/micro-starter-kit/shared/logger/gormlog"
)

func BenchmarkLogger_Print(b *testing.B) {
  logger.Init(logger.WithLevel(zerolog.DebugLevel))
    l := gormlog.NewGormLogger(log.Logger)
    b.ResetTimer()

    for i := 0; i < b.N; i++ {
        l.Print(
            "sql",
            "/some/file.go:34",
            time.Millisecond*5,
            "SELECT * FROM test WHERE id = $1",
            []interface{}{42},
            int64(1),
        )
    }
}

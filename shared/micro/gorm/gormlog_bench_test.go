package gormlog_test

import (
	"os"
	"testing"
	"time"

	ml "github.com/micro/go-micro/v2/logger"
	gormlog "github.com/xmlking/micro-starter-kit/shared/micro/gorm"
	zero "github.com/xmlking/micro-starter-kit/shared/micro/logger/zerolog"
)

func BenchmarkLogger_Print(b *testing.B) {
	mLogger := zero.NewLogger(zero.WithOut(os.Stdout), zero.WithLevel(ml.DebugLevel))
	l := gormlog.NewGormLogger(mLogger)

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

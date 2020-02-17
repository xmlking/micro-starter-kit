package gormlog_test

import (
	"os"
	"testing"
	"time"

	"github.com/micro/go-micro/v2/logger"
	glog "github.com/xmlking/micro-starter-kit/shared/micro/logger/gormlog"
    zlog "github.com/micro/go-plugins/logger/zerolog/v2"
)

func BenchmarkLogger_Print(b *testing.B) {
	mLogger := zlog.NewLogger(zlog.WithOut(os.Stdout), zlog.WithLevel(logger.DebugLevel))
	l := glog.NewGormLogger(mLogger)
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

package gormlog_test

import (
	"os"
	"time"

	// "github.com/micro/go-micro/v2/logger"
	"github.com/xmlking/micro-starter-kit/shared/micro/logger"
	glog "github.com/xmlking/micro-starter-kit/shared/micro/logger/gormlog"
	zlog "github.com/xmlking/micro-starter-kit/shared/micro/logger/zerolog"
)

func ExampleLogger() {
	mLogger := zlog.NewLogger(zlog.WithOut(os.Stdout), zlog.WithLevel(logger.DebugLevel))

	l := glog.NewGormLogger(mLogger)

	l.Print(
		"sql",
		"/foo/bar.go",
		time.Second*2,
		"SELECT * FROM foo WHERE id = ?",
		[]interface{}{123},
		int64(2),
	)

	// Output:
	// {"level":"debug","sql.duration":2000,"sql.query":"SELECT * FROM foo WHERE id = 123","sql.rows_affected":2,"sql.source":"/foo/bar.go","time":"2020-02-14T21:16:52-08:00","message":"gorm query"}
}

func ExampleWithRecordToFields() {
	mLogger := zlog.NewLogger(zlog.WithOut(os.Stdout), zlog.WithLevel(logger.DebugLevel))

	l := glog.NewGormLogger(
		mLogger,
		glog.WithLevel(logger.DebugLevel),

		glog.WithRecordToFields(func(r glog.Record) map[string]interface{} {
			return map[string]interface{}{
				"caller":        r.Source,
				"duration_ms":   float32(r.Duration.Nanoseconds()/1000) / 1000,
				"query":         r.SQL,
				"rows_affected": r.RowsAffected,
			}
		}),
	)

	l.Print(
		"sql",
		"/foo/bar.go",
		time.Millisecond*200,
		"SELECT * FROM foo WHERE id = ?",
		[]interface{}{123},
		int64(2),
	)

	// Output:
	// {"level":"debug","caller":"/foo/bar.go","duration_ms":200,"query":"SELECT * FROM foo WHERE id = 123","rows_affected":2,"time":"2020-02-14T21:18:28-08:00","message":"gorm query"}
}

/**
func TestLogger_Print(t *testing.T) {
	t.Run("log with values < 2", func(t *testing.T) {
		l, buf := logger()

		l.Print("idunno")
		expected := `{"level":"debug","msg":"idunno","sql.source":""}`

		actual := buf.Lines()[0]
		if actual != expected {
			t.Fatalf("Expected %s but got %s", expected, actual)
		}
	})

	t.Run("log with values = 2 (error)", func(t *testing.T) {
		l, buf := logger()

		l.Print("/some/file.go:32", errors.New("some serious error!"))
		expected := `{"level":"error","msg":"some serious error!","sql.source":"/some/file.go:32"}`

		actual := buf.Lines()[0]
		if actual != expected {
			t.Fatalf("Expected %s but got %s", expected, actual)
		}
	})

	t.Run("log with level = log (error)", func(t *testing.T) {
		l, buf := logger()

		l.Print(
			"log",
			"/some/file.go:33",
			errors.New("some serious error!"),
		)
		expected := `{"level":"error","msg":"some serious error!","sql.source":"/some/file.go:33"}`

		actual := buf.Lines()[0]
		if actual != expected {
			t.Fatalf("Expected %s but got %s", expected, actual)
		}
	})

	t.Run("log with level = log (user log)", func(t *testing.T) {
		l, buf := logger()

		l.Print(
			"log",
			"/some/file.go:33",
			"foo",
			"bar",
		)
		expected := `{"level":"debug","msg":"foobar","sql.source":"/some/file.go:33"}`

		actual := buf.Lines()[0]
		if actual != expected {
			t.Fatalf("Expected %s but got %s", expected, actual)
		}
	})

	t.Run("log with level = sql", func(t *testing.T) {
		l, buf := logger()

		l.Print(
			"sql",
			"/some/file.go:34",
			time.Millisecond*5,
			"SELECT * FROM test WHERE id = $1",
			[]interface{}{42},
			int64(1),
		)
		expected := `{"level":"debug","msg":"gorm query","sql.source":"/some/file.go:34","sql.duration":"5ms","sql.query":"SELECT * FROM test WHERE id = 42","sql.rows_affected":1}`

		actual := buf.Lines()[0]
		if actual != expected {
			t.Fatalf("Expected %s but got %s", expected, actual)
		}
	})
}

func logger() (*gormlog.GormLogger, *eroztest.Buffer) {

	mLogger := zlog.NewLogger(zlog.WithLevel(ml.DebugLevel))

	return gormlog.NewGormLogger(mLogger), buf
}
**/

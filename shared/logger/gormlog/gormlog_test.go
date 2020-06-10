package gormlog_test

import (
    "os"
    "time"

    "github.com/rs/zerolog"
    "github.com/rs/zerolog/log"

    "github.com/xmlking/micro-starter-kit/shared/logger"
    "github.com/xmlking/micro-starter-kit/shared/logger/gormlog"
)

func ExampleLogger() {
   logger.Init(
        logger.WithOutput(os.Stdout),
        logger.WithFormat(logger.JSON),
        logger.WithTimeFormat("ddd"),
        logger.WithLevel(zerolog.DebugLevel),
    )

    l := gormlog.NewGormLogger(log.Logger, gormlog.WithLevel(zerolog.DebugLevel))

    l.Print(
        "sql",
        "/foo/bar.go",
        time.Second*2,
        "SELECT * FROM foo WHERE id = ?",
        []interface{}{123},
        int64(2),
    )

    // Output:
    //{"level":"info","LogLevel":"debug","LogFormat":"json","time":"ddd","message":"Logger set to Zerolog with:"}
    //{"level":"debug","duration":2000,"query":"SELECT * FROM foo WHERE id = 123","rows_affected":2,"source":"/foo/bar.go","time":"ddd","message":"gorm query"}

}

func ExampleWithRecordToFields() {
    logger.Init(
        logger.WithOutput(os.Stdout),
        logger.WithFormat(logger.JSON),
        logger.WithTimeFormat("ddd"),
        logger.WithLevel(zerolog.DebugLevel),
    )

    l := gormlog.NewGormLogger(
        log.Logger,
        gormlog.WithLevel(zerolog.DebugLevel),

        gormlog.WithRecordToFields(func(r gormlog.Record) map[string]interface{} {
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
    //{"level":"info","LogLevel":"debug","LogFormat":"json","time":"ddd","message":"Logger set to Zerolog with:"}
    //{"level":"debug","caller":"/foo/bar.go","duration_ms":200,"query":"SELECT * FROM foo WHERE id = 123","rows_affected":2,"time":"ddd","message":"gorm query"}
}

/**
func TestLogger_Print(t *testing.T) {
	t.Run("log with values < 2", func(t *testing.T) {
		l, buf := logger()
		l.Print("idunno")
		expected := `{"level":"debug","msg":"idunno","source":""}`
		actual := buf.Lines()[0]
		if actual != expected {
			t.Fatalf("Expected %s but got %s", expected, actual)
		}
	})
	t.Run("log with values = 2 (error)", func(t *testing.T) {
		l, buf := logger()
		l.Print("/some/file.go:32", errors.New("some serious error!"))
		expected := `{"level":"error","msg":"some serious error!","source":"/some/file.go:32"}`
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
		expected := `{"level":"error","msg":"some serious error!","source":"/some/file.go:33"}`
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
		expected := `{"level":"debug","msg":"foobar","source":"/some/file.go:33"}`
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
		expected := `{"level":"debug","msg":"gorm query","source":"/some/file.go:34","duration":"5ms","query":"SELECT * FROM test WHERE id = 42","rows_affected":1}`
		actual := buf.Lines()[0]
		if actual != expected {
			t.Fatalf("Expected %s but got %s", expected, actual)
		}
	})
}
func logger() (*gormlog.GormLogger, *eroztest.Buffer) {
	mLogger := logger.NewLogger(logger.WithLevel(ml.DebugLevel))
	return gormlog.NewGormLogger(mLogger), buf
}
**/

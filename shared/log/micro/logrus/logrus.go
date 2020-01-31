package micrologrus

import (
	"runtime"
	"strings"

	microlog "github.com/micro/go-micro/v2/debug/log"
	"github.com/sirupsen/logrus"
)

type MicroLogrus struct {
	l logrus.FieldLogger
}

func parentCaller() string {
	pc, _, _, ok := runtime.Caller(4)
	fn := runtime.FuncForPC(pc)
	if ok && fn != nil {
		return fn.Name()
	}

	return ""
}

func NewMicroLogrus(logger logrus.FieldLogger) MicroLogrus {
	ml := MicroLogrus{}
	ml.l = logger

	return ml
}

func (ml MicroLogrus) Read(ops ...microlog.ReadOption) ([]microlog.Record, error) {
	return nil, nil
}
func (ml MicroLogrus) Write(rec microlog.Record) error {
	pc := parentCaller()
	if strings.HasSuffix(pc, "Fatal") {
		ml.l.Fatal(rec.Message)
	} else {
		ml.l.Info(rec.Message)
	}
	return nil
}

func (ml MicroLogrus) Stream() (microlog.Stream, error) {
	return nil, nil
}

func (ml MicroLogrus) Log(v ...interface{}) {
	pc := parentCaller()
	if strings.HasSuffix(pc, "Fatal") {
		ml.l.Fatal(v...)
	} else {
		ml.l.Info(v...)
	}
}

func (ml MicroLogrus) Logf(format string, v ...interface{}) {
	pc := parentCaller()
	if strings.HasSuffix(pc, "Fatalf") {
		ml.l.Fatalf(format, v...)
	} else {
		ml.l.Infof(format, v...)
	}
}

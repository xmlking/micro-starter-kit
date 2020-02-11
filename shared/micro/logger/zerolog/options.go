package zero

import (
	"context"
	"io"

	"github.com/micro/go-micro/v2/logger"
	"github.com/rs/zerolog"
)

type formatterKey struct{}
type levelKey struct{}
type levelFieldKey struct{}
type outKey struct{}
type hooksKey struct{}
type reportCallerKey struct{}
type exitKey struct{}
type prettyKey struct{}
type useColorKey struct{}

type Options struct {
	logger.Options
}

func WithTimeFieldFormat(formatter zerolog.Formatter) logger.Option {
	return setOption(formatterKey{}, formatter)
}

func WithLevelFieldName(levelFieldName string) logger.Option {
	return setOption(levelFieldKey{}, levelFieldName)
}

func WithPretty(pretty bool) logger.Option {
	return setOption(prettyKey{}, pretty)
}
func WithColor(useColor bool) logger.Option {
	return setOption(useColorKey{}, useColor)
}

func WithLevel(lvl logger.Level) logger.Option {
	return setOption(levelKey{}, lvl)
}

func WithOut(out io.Writer) logger.Option {
	return setOption(outKey{}, out)
}

func WithHooks(hooks []zerolog.Hook) logger.Option {
	return setOption(hooksKey{}, hooks)
}

func WithReportCaller(reportCaller bool) logger.Option {
	return setOption(reportCallerKey{}, reportCaller)
}

func WithExitFunc(exit func(int)) logger.Option {
	return setOption(exitKey{}, exit)
}

func setOption(k, v interface{}) logger.Option {
	return func(o *logger.Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, k, v)
	}
}

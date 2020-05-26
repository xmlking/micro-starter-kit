package config

import "context"

type Option func(*Options)

type Options struct {
	ConfigDir  string
	ConfigFile string
	// Alternative options
	Context context.Context
}

func WithConfigDir(configDir string) Option {
	return func(args *Options) {
		args.ConfigDir = configDir
	}
}

func WithConfigFile(configFile string) Option {
	return func(args *Options) {
		args.ConfigFile = configFile
	}
}

func SetOption(k, v interface{}) Option {
	return func(o *Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, k, v)
	}
}

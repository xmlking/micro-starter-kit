package config

// ServiceConfiguration is the top level configuration struct which is loaded from the defined source(s)
type ServiceConfiguration struct {
	Log         LogConfiguration
	Environment string `json:"environment"`
	Database    DatabaseConfiguration
}

// LogConfiguration holds log config
type LogConfiguration struct {
	Level  string `json:"level"`
	Format string `json:"format"` // json or text
}

// DatabaseConfiguration holds db config
type DatabaseConfiguration struct {
	Dialect  string `json:"dialect" default:"postgres"`
	Driver   string `json:"driver"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"user"`
	Password string `json:"password"`
	Charset  string `json:"charset"`
}

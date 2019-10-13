package config

import "fmt"

const (
	// Development environment
	Development Environment = "development"
	// Test environment
	Test Environment = "test"
	// Production environment
	Production Environment = "production"
)

// Environment represents an application environment.
type Environment string

// TODO: use proto to define config.yaml schema
// ServiceConfiguration is the top level configuration struct which is loaded from the defined source(s)
type ServiceConfiguration struct {
	ServiceName string      `json:"name"`
	Version     string      `json:"version"`
	Environment Environment `json:"environment"`
	Log         LogConfiguration
	Database    DatabaseConfiguration
	Email       EmailConfiguration
}

// LogConfiguration holds log config
type LogConfiguration struct {
	Level  string `json:"level"`
	Format string `json:"format"` // json or text
}

const (
	// PostgreSQLDialect is the dialect name for PostgreSQL.
	PostgreSQLDialect DatabaseDialect = "postgres"
	// MySQLDialect is the dialect name for MySQL.
	MySQLDialect DatabaseDialect = "mysql"
	// SQLiteDialect is the dialect name for SQLite.
	SQLiteDialect DatabaseDialect = "sqlite3"
)

// DatabaseDialect represents a database dialect constant.
type DatabaseDialect string

// DatabaseConfiguration holds db config
type DatabaseConfiguration struct {
	Dialect     DatabaseDialect `json:"dialect" default:"postgres"`
	Host        string          `json:"host"`
	Port        int             `json:"port"`
	Username    string          `json:"username"`
	Password    string          `json:"password"`
	Database    string          `json:"database"`
	Charset     string          `json:"charset" default:"utf8"`
	UTC         bool            `default:"true"`
	Logging     bool            `default:"false"`
	Singularize bool            `default:"false"`
	Params      map[string]interface{}
}

// URL returns a connection string for the database.
func (d *DatabaseConfiguration) URL() (url string, err error) {

	switch d.Dialect {
	case SQLiteDialect:
		return d.Database, nil
	case PostgreSQLDialect:
		return fmt.Sprintf(
			"host=%s port=%v user=%s dbname=%s sslmode=disable password=%s",
			d.Host, d.Port, d.Username, d.Database, d.Password,
		), nil
	case MySQLDialect:
		return fmt.Sprintf(
			"%s:%s@(%s:%v)/%s?charset=%s&parseTime=True&loc=Local",
			d.Username, d.Password, d.Host, d.Port, d.Database, d.Charset,
		), nil
	default:
		return "", fmt.Errorf(" '%v' driver doesn't exist. ", d.Dialect)
	}

	// TODO https://github.com/coderosoio/cortito/blob/master/common/config/database.go

}

// EmailConfiguration holds email config
type EmailConfiguration struct {
	Username    string
	Password    string
	EmailServer string
	Port        int
	From        string
}

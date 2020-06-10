package database

import (
    "testing"
    "time"

    _ "github.com/jinzhu/gorm/dialects/sqlite"

    _ "github.com/xmlking/micro-starter-kit/shared/logger"
    configPB "github.com/xmlking/micro-starter-kit/shared/proto/config"
)

func TestDatabase(t *testing.T) {
    dur := time.Hour
	_, err := GetDatabaseConnection(configPB.DatabaseConfiguration{
		Dialect:  configPB.DatabaseDialect_SQLite3,
		Host:     "127.0.0.1",
		Port:     3306,
		Username: "root",
		Password: "123456",
		Database: ":memory:",
        MaxOpenConns: 1,
        MaxIdleConns: 1,
        ConnMaxLifetime: &dur,
	})
	if err != nil {
		t.Fatalf("Database connection failed, %v!", err)
	}
}

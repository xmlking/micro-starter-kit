package database

import (
    "testing"

    _ "github.com/jinzhu/gorm/dialects/sqlite"

    "github.com/xmlking/micro-starter-kit/shared/config"
    // bootstrap config and logger
    _ "github.com/xmlking/micro-starter-kit/shared/logger"
)

func TestDatabase(t *testing.T) {
	_, err := GetDatabaseConnection(config.DatabaseConfiguration{
		Dialect:  "sqlite3",
		Host:     "127.0.0.1",
		Port:     3306,
		Username: "root",
		Password: "123456",
		Database: ":memory:",
	})
	if err != nil {
		t.Fatalf("Database connection failed, %v!", err)
	}
}

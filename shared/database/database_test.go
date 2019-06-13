package database

import (
	"testing"

	"github.com/xmlking/micro-starter-kit/shared/config"
)

func TestDatabase(t *testing.T) {
	_, err := GetDatabaseConnection(&config.DatabaseConfiguration{
		Dialect:  "sqlite3",
		Host:     "127.0.0.1",
		Port:     3306,
		Username: "root",
		Password: "123456",
		Database: ":memory:",
	})
	if err != nil {
		t.Errorf("Database connection failed, %v!", err)
	}
}

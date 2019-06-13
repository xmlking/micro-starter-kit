package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/xmlking/micro-starter-kit/shared/config"
)

// GetDatabaseConnection return (gorm.DB or error)
func GetDatabaseConnection(dbConf *config.DatabaseConfiguration) (db *gorm.DB, err error) {
	var timezoneCommand string

	switch dbConf.Dialect {
	case config.SQLiteDialect:
		db, err = sqliteConnection(dbConf)
	case config.PostgreSQLDialect:
		timezoneCommand = "SET timezone = 'UTC'"
		db, err = postgresConnection(dbConf)
	case config.MySQLDialect:
		timezoneCommand = "SET time_zone = '+00:00'"
		db, err = mysqlConnection(dbConf)
	default:
		return nil, fmt.Errorf("database dialect %s not supported", dbConf.Dialect)
	}

	if err != nil {
		return
	}

	if dbConf.Logging {
		db.Debug()
	}

	db.LogMode(dbConf.Logging)
	db.SingularTable(dbConf.Singularize)
	// db.DB().SetMaxOpenConns(400)
	// db.DB().SetMaxIdleConns(0)
	// db.DB().SetConnMaxLifetime(100 * time.Second)

	if dbConf.UTC {
		if _, err = db.DB().Exec(timezoneCommand); err != nil {
			return nil, fmt.Errorf("error setting UTC timezone: %v", err)
		}
	}

	return
}

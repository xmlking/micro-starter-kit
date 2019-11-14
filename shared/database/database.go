package database

import (
	"fmt"

	"github.com/jinzhu/gorm"

	"github.com/sirupsen/logrus"
	"github.com/xmlking/micro-starter-kit/shared/config"

	gormlog "github.com/onrik/logrus/gorm"
)

// GetDatabaseConnection return (gorm.DB or error)
func GetDatabaseConnection(dbConf config.DatabaseConfiguration) (db *gorm.DB, err error) {
	var timezoneCommand string

	switch dbConf.Dialect {
	case config.SQLiteDialect:
		db, err = connection(dbConf)
	case config.PostgreSQLDialect:
		timezoneCommand = "SET timezone = 'UTC'"
		db, err = connection(dbConf)
	case config.MySQLDialect:
		timezoneCommand = "SET time_zone = '+00:00'"
		db, err = connection(dbConf)
	default:
		return nil, fmt.Errorf("database dialect %s not supported", dbConf.Dialect)
	}

	if err != nil {
		return
	}

	//db.SetLogger(logrus.WithFields(logrus.Fields{"module": "gorm"}))
	//db.SetLogger(logrus.StandardLogger())
	db.SetLogger(gormlog.New(logrus.WithFields(logrus.Fields{"module": "gorm"})))

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
			return nil, fmt.Errorf("error setting UTC timezone: %w", err)
		}
	}

	return
}

func connection(dbConf config.DatabaseConfiguration) (db *gorm.DB, err error) {
	url, err := dbConf.URL()
	if err != nil {
		return nil, err
	}
	db, err = gorm.Open(string(dbConf.Dialect), url)
	return
}

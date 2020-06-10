package database

import (
    "fmt"
    "strings"

    "github.com/jinzhu/gorm"
    "github.com/rs/zerolog/log"

    "github.com/xmlking/micro-starter-kit/shared/logger/gormlog"
    configPB "github.com/xmlking/micro-starter-kit/shared/proto/config"
)

// GetDatabaseConnection return (gorm.DB or error)
func GetDatabaseConnection(dbConf configPB.DatabaseConfiguration) (db *gorm.DB, err error) {
    var timezoneCommand string

    switch dbConf.Dialect {
    case configPB.DatabaseDialect_SQLite3:
        db, err = connection(dbConf)
    case configPB.DatabaseDialect_Postgre:
        timezoneCommand = "SET timezone = 'UTC'"
        db, err = connection(dbConf)
    case configPB.DatabaseDialect_MySQL:
        timezoneCommand = "SET time_zone = '+00:00'"
        db, err = connection(dbConf)
    default:
        return nil, fmt.Errorf("database dialect %s not supported", dbConf.Dialect)
    }

    if err != nil {
        return
    }
    gLogger := log.With().
        Str("module", "gorm").
        Logger()

    db.SetLogger(gormlog.NewGormLogger(gLogger))

    if dbConf.Logging {
        db.Debug()
    }

    db.LogMode(dbConf.Logging)
    db.SingularTable(dbConf.Singularize)
    db.DB().SetMaxOpenConns(int(dbConf.MaxOpenConns))
    db.DB().SetMaxIdleConns(int(dbConf.MaxIdleConns))
    db.DB().SetConnMaxLifetime(*dbConf.ConnMaxLifetime)

    if dbConf.Utc {
        if _, err = db.DB().Exec(timezoneCommand); err != nil {
            return nil, fmt.Errorf("error setting UTC timezone: %w", err)
        }
    }

    return
}

func connection(dbConf configPB.DatabaseConfiguration) (db *gorm.DB, err error) {
    url, err := dbConf.URL()
    if err != nil {
        return nil, err
    }
    db, err = gorm.Open(strings.ToLower(string(dbConf.Dialect.String())), url)
    return
}

package database

import (
	"github.com/jinzhu/gorm"
	// for postgres
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/xmlking/micro-starter-kit/shared/config"
)

// postgresConnection
func postgresConnection(dbConf *config.DatabaseConfiguration) (db *gorm.DB, err error) {
	url, err := dbConf.URL()
	if err != nil {
		return nil, err
	}
	db, err = gorm.Open(string(dbConf.Dialect), url)
	return
}

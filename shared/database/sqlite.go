package database

import (
	"github.com/jinzhu/gorm"
	// for sqlite
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/xmlking/micro-starter-kit/shared/config"
)

// sqliteConnection
func sqliteConnection(dbConf *config.DatabaseConfiguration) (db *gorm.DB, err error) {
	url, err := dbConf.URL()
	if err != nil {
		return nil, err
	}
	db, err = gorm.Open(string(dbConf.Dialect), url)
	return
}

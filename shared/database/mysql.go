package database

import (
	"github.com/jinzhu/gorm"
	// for mysql
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/xmlking/micro-starter-kit/shared/config"
)

// mysqlConnection
func mysqlConnection(dbConf *config.DatabaseConfiguration) (db *gorm.DB, err error) {
	url, err := dbConf.URL()
	if err != nil {
		return nil, err
	}
	db, err = gorm.Open(string(dbConf.Dialect), url)
	return
}

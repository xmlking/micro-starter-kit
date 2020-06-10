package config

import "fmt"

// URL returns a connection string for the database.
func (d *DatabaseConfiguration) URL() (url string, err error) {

    switch d.Dialect {
    case DatabaseDialect_SQLite3:
        return d.Database, nil
    case DatabaseDialect_Postgre:
        return fmt.Sprintf(
            "host=%s port=%v user=%s dbname=%s sslmode=disable password=%s",
            d.Host, d.Port, d.Username, d.Database, d.Password,
        ), nil
    case DatabaseDialect_MySQL:
        return fmt.Sprintf(
            "%s:%s@(%s:%v)/%s?charset=%s&parseTime=True&loc=Local",
            d.Username, d.Password, d.Host, d.Port, d.Database, d.Charset,
        ), nil
    default:
        return "", fmt.Errorf(" '%v' driver doesn't exist. ", d.Dialect)
    }

    // TODO https://github.com/coderosoio/cortito/blob/master/common/config/database.go
}

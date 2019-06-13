package database

import (
	"github.com/xmlking/micro-starter-kit/shared/config"
)

func bigtableConnection(dbConf *config.DatabaseConfiguration) (conn interface{}, err error) {
	// settings, found := dbConf.KeyValue[connectionName]
	// if !found {
	// 	return nil, fmt.Errorf("could not get connection settings for %s", connectionName)
	// }
	// switch settings.Adapter {
	// case cfg.RedisKeyValueAdapter:
	// 	fallthrough
	// default:
	// 	conn, err = getRedisConnection(settings)
	// }
	return nil, nil
}
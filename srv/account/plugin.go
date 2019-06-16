package main

import (
	_ "github.com/xmlking/micro-starter-kit/shared/config"
	_ "github.com/xmlking/micro-starter-kit/shared/log"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
	// _ "github.com/jinzhu/gorm/dialects/postgres"
	// _ "github.com/jinzhu/gorm/dialects/mysql"
	// _ "github.com/micro/go-plugins/registry/kubernetes"
	// _ "github.com/micro/go-plugins/transport/nats"
)

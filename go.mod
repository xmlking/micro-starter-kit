module github.com/xmlking/micro-starter-kit

go 1.14

// replace github.com/micro/go-micro/v2 => /Users/schintha/Developer/Work/go/3rd-party/go-micro
// FIXME : https://github.com/etcd-io/etcd/issues/11563
replace google.golang.org/grpc => google.golang.org/grpc v1.29.1

require (
	github.com/DATA-DOG/go-sqlmock v1.4.1 // indirect
	github.com/envoyproxy/protoc-gen-validate v0.1.0
	github.com/golang/protobuf v1.4.0
	github.com/google/uuid v1.1.1
	github.com/infobloxopen/atlas-app-toolkit v0.20.0
	github.com/infobloxopen/protoc-gen-gorm v0.20.0
	github.com/jinzhu/gorm v1.9.12
	github.com/markbates/pkger v0.15.1
	github.com/micro/cli/v2 v2.1.2
	github.com/micro/go-micro/v2 v2.5.0
	github.com/micro/go-plugins/config/source/pkger/v2 v2.5.0
	github.com/pkg/errors v0.9.1
	github.com/sarulabs/di/v2 v2.4.0
	github.com/satori/go.uuid v1.2.0
	github.com/stretchr/testify v1.5.1
	github.com/thoas/go-funk v0.6.0
	github.com/xmlking/logger v0.1.5
	github.com/xmlking/logger/gormlog v0.1.5
	github.com/xmlking/logger/zerolog v0.1.5
	golang.org/x/lint v0.0.0-20200302205851-738671d3881b // indirect
	golang.org/x/net v0.0.0-20200301022130-244492dfa37a // indirect
	golang.org/x/sys v0.0.0-20200302150141-5c8b2ff67527 // indirect
	golang.org/x/tools v0.0.0-20200312045724-11d5b4c81c7d // indirect
	google.golang.org/appengine v1.6.5 // indirect
	google.golang.org/genproto v0.0.0-20200420144010-e5e8543f8aeb
	google.golang.org/grpc v1.28.0 // indirect
	honnef.co/go/tools v0.0.1-2020.1.3 // indirect
)

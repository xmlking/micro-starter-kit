module github.com/xmlking/micro-starter-kit

go 1.14

// for local development, you can repoint go-micro to local development go-micro workspace
// replace github.com/micro/go-micro/v2 => /Users/schintha/Developer/Work/go/3rd-party/go-micro
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

// replace github.com/xmlking/configor => /Users/schintha/Developer/Work/go/configor

require (
	github.com/DATA-DOG/go-sqlmock v1.4.1 // indirect
	github.com/envoyproxy/protoc-gen-validate v0.1.0
	github.com/gogo/protobuf v1.3.1
	github.com/golang/protobuf v1.4.2
	github.com/infobloxopen/atlas-app-toolkit v0.21.1
	github.com/infobloxopen/protoc-gen-gorm v0.20.0
	github.com/jinzhu/gorm v1.9.13
	github.com/markbates/pkger v0.17.0
	github.com/micro/go-micro/v2 v2.9.1
	github.com/pkg/errors v0.9.1
	github.com/rs/zerolog v1.19.0
	github.com/sarulabs/di/v2 v2.4.0
	github.com/satori/go.uuid v1.2.0
	github.com/stretchr/testify v1.6.1
	github.com/thoas/go-funk v0.6.0
	github.com/xmlking/configor v0.2.2
	google.golang.org/genproto v0.0.0-20191216164720-4f79533eabd1
	google.golang.org/grpc v1.26.0
)

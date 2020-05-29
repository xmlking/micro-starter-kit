module github.com/xmlking/micro-starter-kit

go 1.14

// for local development, you can repoint go-micro to local development go-micro workspace
// replace github.com/micro/go-micro/v2 => /Users/schintha/Developer/Work/go/3rd-party/go-micro
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/envoyproxy/protoc-gen-validate v0.1.0
	github.com/golang/protobuf v1.4.2
	github.com/google/uuid v1.1.1
	github.com/infobloxopen/atlas-app-toolkit v0.21.1
	github.com/infobloxopen/protoc-gen-gorm v0.20.0
	github.com/jinzhu/gorm v1.9.12
	github.com/markbates/pkger v0.16.0
	github.com/micro/cli/v2 v2.1.2
	github.com/micro/go-micro/v2 v2.7.0
	github.com/micro/go-plugins/broker/googlepubsub/v2 v2.5.0
	github.com/micro/go-plugins/config/source/pkger/v2 v2.5.0
	github.com/pkg/errors v0.9.1
	github.com/rs/zerolog v1.18.0
	github.com/sarulabs/di/v2 v2.4.0
	github.com/satori/go.uuid v1.2.0
	github.com/stretchr/testify v1.6.0
	github.com/thoas/go-funk v0.6.0
	google.golang.org/genproto v0.0.0-20191216164720-4f79533eabd1
	google.golang.org/grpc v1.29.1
)

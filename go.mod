module github.com/xmlking/micro-starter-kit

go 1.13

// replace github.com/micro/go-micro/v2 => /Users/schintha/Developer/Work/go/3rd-party/go-micro
// FIXME : https://github.com/etcd-io/etcd/issues/11563
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

// replace github.com/coreos/go-systemd => github.com/coreos/go-systemd/v22 v22.0.0

require (
	github.com/DATA-DOG/go-sqlmock v1.4.1 // indirect
	github.com/envoyproxy/protoc-gen-validate v0.1.0
	github.com/golang/protobuf v1.3.3
	github.com/google/uuid v1.1.1
	github.com/infobloxopen/atlas-app-toolkit v0.19.2
	github.com/infobloxopen/protoc-gen-gorm v0.18.0
	github.com/jinzhu/gorm v1.9.12
	github.com/micro/cli/v2 v2.1.2-0.20200204093551-dfdc8f23b971
	github.com/micro/go-micro/v2 v2.1.0
	github.com/micro/go-plugins/config/source/pkger/v2 v2.0.1
	github.com/rs/zerolog v1.18.0
	github.com/sarulabs/di/v2 v2.4.0
	github.com/satori/go.uuid v1.2.0
	github.com/sirupsen/logrus v1.4.2
	github.com/stretchr/testify v1.4.0
	github.com/thoas/go-funk v0.5.0
	google.golang.org/genproto v0.0.0-20200207204624-4f3edf09f4f6
)

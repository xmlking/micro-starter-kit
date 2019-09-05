module github.com/xmlking/micro-starter-kit

go 1.13

replace (
	github.com/hashicorp/consul => github.com/hashicorp/consul v1.6.0
	k8s.io/api => k8s.io/api v0.0.0-20190813180838-e711354a0280
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20190813060636-0c17871ad6fd
)

require (
	github.com/envoyproxy/protoc-gen-validate v0.1.0
	github.com/golang/protobuf v1.3.2
	github.com/infobloxopen/atlas-app-toolkit v0.18.2
	github.com/infobloxopen/protoc-gen-gorm v0.17.0
	github.com/jinzhu/gorm v1.9.10
	github.com/micro/cli v0.2.0
	github.com/micro/go-micro v1.9.1
	github.com/micro/go-plugins v1.2.0
	github.com/micro/micro v1.9.1
	github.com/onrik/logrus v0.4.1
	github.com/sarulabs/di/v2 v2.2.0
	github.com/satori/go.uuid v1.2.0
	github.com/sirupsen/logrus v1.4.2
	github.com/stretchr/testify v1.4.0
	github.com/thoas/go-funk v0.4.0
	github.com/tudurom/micro-logrus v0.0.0-20171007082012-3704f28fa9d1
	google.golang.org/genproto v0.0.0-20190905072037-92dd089d5514
)

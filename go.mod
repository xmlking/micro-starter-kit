module github.com/xmlking/micro-starter-kit

go 1.13.1

replace (
	k8s.io/api => k8s.io/api v0.0.0-20191003035645-10e821c09743
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20191003115452-c31ffd88d5d2
)

require (
	github.com/DATA-DOG/go-sqlmock v1.3.3 // indirect
	github.com/chzyer/logex v1.1.10 // indirect
	github.com/chzyer/test v0.0.0-20180213035817-a1ea475d72b1 // indirect
	github.com/envoyproxy/protoc-gen-validate v0.1.0
	github.com/go-telegram-bot-api/telegram-bot-api v4.6.4+incompatible // indirect
	github.com/golang/protobuf v1.3.2
	github.com/infobloxopen/atlas-app-toolkit v0.19.0
	github.com/infobloxopen/protoc-gen-gorm v0.18.0
	github.com/jinzhu/gorm v1.9.11
	github.com/lusis/go-slackbot v0.0.0-20180109053408-401027ccfef5 // indirect
	github.com/lusis/slack-test v0.0.0-20190426140909-c40012f20018 // indirect
	github.com/micro/cli v0.2.0
	github.com/micro/go-micro v1.11.1
	github.com/micro/go-plugins v1.3.0
	github.com/micro/micro v1.11.1
	github.com/onrik/logrus v0.4.1
	github.com/sarulabs/di/v2 v2.2.0
	github.com/satori/go.uuid v1.2.0
	github.com/sirupsen/logrus v1.4.2
	github.com/stretchr/testify v1.4.0
	github.com/thoas/go-funk v0.4.0
	github.com/tudurom/micro-logrus v0.0.0-20171007082012-3704f28fa9d1
	google.golang.org/genproto v0.0.0-20191002211648-c459b9ce5143
)

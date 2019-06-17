# micro-starter-kit

> go-micro starter kit

## What you get

- [x] Monorepo
- [x] gRPC microservices with REST Gateway
- [ ] gRPC validation
- [x] config fallback
- [x] custom logging
- [x] CRUD via ORM
- [ ] Observability

## TODO

- [ ] [protoc-gen-gorm](https://github.com/infobloxopen/protoc-gen-gorm)
- [ ] [envoyproxy/protoc-gen-validate](https://github.com/envoyproxy/protoc-gen-validate)

## Prerequisite

> micro, go-micro versions are at `v.1.6.0`
> Global tools:

```bash
brew install protobuf
go get -u github.com/micro/micro
go get -u github.com/micro/go-micro

go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
go get -u github.com/micro/protoc-gen-micro
```

## Initial Setup

> (optional) setup your workspace from scratch

```bash
go mod init github.com/xmlking/micro-starter-kit
go get -u github.com/micro/go-micro
mkdir srv api fnc

# scaffold modules
micro new --namespace="go.micro" --type="srv" --gopath=false --alias="account" srv/account

micro new --namespace="go.micro" --type="srv" --gopath=false \
--alias="emailer"  --plugin=registry=mdns:broker=nats srv/emailer

micro new --namespace="go.micro" --type="api" --gopath=false --alias="account" api/account
```

## Build

```bash
make proto
# silence
make -s proto

# prod build. Build with plugins.go
go build -o echo srv/echo/main.go srv/echo/plugin.go
```

## Run

> Optionally start postgres and set it in `config.yaml`

```bash
postgres
docker-compose up postgres
```

```bash
# dev mode
# test account srv (plugin adds custom logger )
go run srv/account/main.go srv/account/plugin.go
go run srv/emailer/main.go srv/emailer/plugin.go

# prod mode
MICRO_BROKER=kafka \
MICRO_REGISTRY=kubernetes \
MICRO_TRANSPORT=nats \
./account-srv



# test config with CMD
go run cmd/demo/main.go --help
go run cmd/demo/main.go --database_host=1.1.1.1 --database_port=7777
APP_ENV=production go run cmd/demo/main.go
```

### Run Micro

```bash
micro list services
micro get service go.micro.srv.echo

# run API Gateway
micro api --namespace=go.micro.srv
# run Web UX for testing
micro web --namespace=go.micro.srv
```

## Test

> create new user from `Micro Web UI` and see if an email is send

```json
{
"username": "sumo",
"firstName": "sumoto",
"lastName": "demo",
"email": "sumo@demo.com"
}
```

## Reference

1. [examples](https://github.com/micro/examples) - example usage code for micro
2. [microhq](https://github.com/microhq) - a place for prebuilt microservices
3. [explorer](https://micro.mu/explore/) - which aggregates micro based open source projects
4. [micro-plugins](https://github.com/micro/go-plugins) extensible micro plugins

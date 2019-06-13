# micro-starter-kit

> go-micro starter kit

## What you get

- [x] Monorepo
- [x] gRPC microservices with REST Gateway
- [ ] gRPC validation
- [x] config fallback
- [ ] custom logging
- [x] CRUD via ORM
- [x] Observability

## TODO

- [ ] [protoc-gen-gorm](https://github.com/infobloxopen/protoc-gen-gorm)
- [ ] [envoyproxy/protoc-gen-validate](https://github.com/envoyproxy/protoc-gen-validate)

## Prerequisite

> micro, go-micro versions are at `v.1.6.0`
> Global tools:

```bash
brew install protobuf
go get -u github.com/micro/micro

go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
go get -u github.com/micro/protoc-gen-micro
```

## Setup

```bash
go mod init github.com/xmlking/micro-starter-kit
go get -u github.com/micro/go-micro
mkdir srv api client

# scaffold demo module
micro new --namespace="go.micro" --type="srv" \
--alias="echo" github.com/xmlking/micro-starter-kit/srv/echo

mv  /Users/schintha/go/src/github.com/xmlking/micro-starter-kit/srv/echo srv


micro new --namespace="go.micro" --type="srv" \
--alias="account" github.com/xmlking/micro-starter-kit/srv/account

mv  /Users/schintha/go/src/github.com/xmlking/micro-starter-kit/srv/account srv
```

## Build

```bash
make proto

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
go run srv/echo/main.go

# prod mode
MICRO_BROKER=kafka \
MICRO_REGISTRY=kubernetes \
MICRO_TRANSPORT=nats \
./echo

# test config with CMD
go run cmd/demo/main.go --help
go run cmd/demo/main.go --database_host=1.1.1.1 --database_port=7777
APP_ENV=production go run cmd/demo/main.go


# test account srv
go run srv/account/main.go
```

### Run Micro

```bash
micro list services
micro get service go.micro.srv.echo

# run API Gateway
micro api --namespace=go.micro.srv
micro web --namespace=go.micro.srv
```

## Test

## Reference

1. [examples](https://github.com/micro/examples) - example usage code for micro
2. [microhq](https://github.com/microhq) - a place for prebuilt microservices
3. [explorer](https://micro.mu/explore/) - which aggregates micro based open source projects
4. [micro-plugins](https://github.com/micro/go-plugins) extensible micro plugins

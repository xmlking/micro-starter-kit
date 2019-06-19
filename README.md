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
# fetche micro into $GOPATH
go get -u github.com/micro/micro
go get -u github.com/micro/go-micro

# for mac, use brew to install protobuf
brew install protobuf

# fetch protoc plugins into $GOPATH
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
go get -u github.com/micro/protoc-gen-micro
# fetch PGV protoc plugin
go get -d github.com/envoyproxy/protoc-gen-validate
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

> Node: `--server_address=<myVpnIp>:5501x --broker_address=<myVpnIp>:5502x` required only when you are behind VPN

```bash
# dev mode
# test account srv (plugin adds custom logger )
# myVpnIp=$(ifconfig | grep "inet " | grep -Fv 127.0.0.1 |  grep -Fv '192.168' | awk '{print $2}')
# go run srv/account/main.go srv/account/plugin.go --server_address=${myVpnIp}:55011 --broker_address=${myVpnIp}:55021
go run srv/account/main.go srv/account/plugin.go
# go run srv/emailer/main.go srv/emailer/plugin.go --server_address=${myVpnIp}:55012 --broker_address=${myVpnIp}:55022
go run srv/emailer/main.go srv/emailer/plugin.go

# prod mode
MICRO_BROKER=nats \
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
micro get service go.micro.srv.account

# run API Gateway
micro api --namespace=go.micro.srv
# (or) run Web UX for testing
micro web --namespace=go.micro.srv
```

## Test

> create new user from `Micro Web UI` and see if an email is send

```json
{
"username": "sumo",
"firstName": "sumo",
"lastName": "demo",
"email": "sumo@demo.com"
}
```

## Reference

1. [examples](https://github.com/micro/examples) - example usage code for micro
2. [microhq](https://github.com/microhq) - a place for prebuilt microservices
3. [explorer](https://micro.mu/explore/) - which aggregates micro based open source projects
4. [micro-plugins](https://github.com/micro/go-plugins) extensible micro plugins

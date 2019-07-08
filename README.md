# micro-starter-kit

> go-micro starter kit

![Image of Yaktocat](docs/deployment.png)

## What you get

- [x] Monorepo
- [x] gRPC microservices with REST Gateway
- [x] Proto Validation
- [x] Config fallback
- [x] Custom logging
- [x] CRUD via ORM
- [x] DI Container
- [ ] Observability

## TODO

- [ ] [protoc-gen-gorm](https://github.com/infobloxopen/protoc-gen-gorm)

## Prerequisite

> micro, go-micro versions are at `v.1.7.0`
> Global tools:

```bash
# fetch micro into $GOPATH
go get -u github.com/micro/micro
go get -u github.com/micro/go-micro
go get -u github.com/google/ko/cmd/ko

# for mac, use brew to install protobuf
brew install protobuf

# fetch protoc plugins into $GOPATH
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
go get -u github.com/micro/protoc-gen-micro
# fetch PGV protoc plugin
go get -u github.com/envoyproxy/protoc-gen-validate
```

## Initial Setup

> (optional) setup your workspace from scratch

```bash
go mod init github.com/xmlking/micro-starter-kit
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

> Node: `--server_address=<MY_VPM_IP>:5501x --broker_address=<MY_VPN_IP>:5502x` required only when you are behind VPN

```bash
# dev mode
# test account srv (plugin adds custom logger )
# MY_VPM_IP=$(ifconfig | grep 172 | awk '{print $2; exit}')
# go run srv/account/main.go srv/account/plugin.go --server_address=${MY_VPN_IP}:55011 --broker_address=${MY_VPN_IP}:55021
go run srv/account/main.go srv/account/plugin.go
# go run srv/emailer/main.go srv/emailer/plugin.go --server_address=${MY_VPN_IP}:55012 --broker_address=${MY_VPN_IP}:55022
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

## Test

> using `micro` CLI

```bash
micro list services
micro get service go.micro.srv.account

# Start API Gateway
micro api --namespace=go.micro.srv
# (or) Start Web UX for testing
micro web --namespace=go.micro.srv
```

### Test gRPC Directly

```bash
micro call go.micro.srv.account UserService.Create '{"username": "sumo", "firstName": "sumo", "lastName": "demo", "email": "sumo@demo.com"}'
micro call go.micro.srv.account UserService.List '{}'
micro call go.micro.srv.account UserService.Get '{"id": 1}'
```

### Test via Micro Web UI

```
open http://localhost:8082
```

> create new user from `Micro Web UI` and see if an email is send

```json
{
  "username": "sumo",
  "firstName": "sumo",
  "lastName": "demo",
  "email": "sumo@demo.com"
}
```

### Test via Micro API Gateway

> Start `API Gateway` and run **REST Client** [tests](test/test-rest-api.http)

## Deploy

Use `ko`. If you are new to `ko` check out the [ko-demo](https://github.com/xmlking/ko-demo)

Set a registry and make sure you can push to it:

```bash
export PROJECT_ID=ngx-starter-kit
export KO_DOCKER_REPO=gcr.io/${PROJECT_ID}
```

Then `apply` like this:

```bash
ko apply -f deploy/
```

To deploy in a different namespace:

```bash
ko -n nondefault apply -f deploy/
```

## Reference

1. [examples](https://github.com/micro/examples) - example usage code for micro
2. [microhq](https://github.com/microhq) - a place for prebuilt microservices
3. [explorer](https://micro.mu/explore/) - which aggregates micro based open source projects
4. [micro-plugins](https://github.com/micro/go-plugins) extensible micro plugins

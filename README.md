# micro-starter-kit

Microservices starter kit for **Golang**, aims to be developer friendly.

[![Build Status](https://img.shields.io/endpoint.svg?url=https%3A%2F%2Factions-badge.atrox.dev%2Fxmlking%2Fmicro-starter-kit%2Fbadge%3Fref%3Ddevelop&style=popout)](https://actions-badge.atrox.dev/xmlking/micro-starter-kit/goto?ref=develop)
[![codecov](https://codecov.io/gh/xmlking/micro-starter-kit/branch/develop/graph/badge.svg)](https://codecov.io/gh/xmlking/micro-starter-kit)
[![Go Report Card](https://goreportcard.com/badge/github.com/xmlking/micro-starter-kit)](https://goreportcard.com/report/github.com/xmlking/micro-starter-kit)
[![fuzzit](https://app.fuzzit.dev/badge?org_id=xmlkinggithub&branch=develop)](https://app.fuzzit.dev/orgs/xmlkinggithub/dashboard)
[![GoDoc](https://godoc.org/github.com/xmlking/micro-starter-kit?status.svg)](https://godoc.org/github.com/xmlking/micro-starter-kit)
[![Go 1.13](https://img.shields.io/badge/go-1.13-9cf.svg)](https://golang.org/dl/)
[![Renovate dependency Status](https://img.shields.io/badge/renovate-enabled-brightgreen.svg)](https://renovatebot.com/)
[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)

## Overview

![Image of Deployment](docs/images/deployment.png)

### What you get

- [x] Monorepo - Sharing Code Between Microservices
- [x] gRPC microservices with REST Gateway
- [x] Input Validation with [protoc-gen-validate (PGV)](https://github.com/envoyproxy/protoc-gen-validate)
- [x] Config - Pluggable Sources, Mergeable Config, Safe Recovery
- [x] Customizable Logging
- [x] CRUD Example using [GORM](https://gorm.io/)
- [x] GORM code gen via [protoc-gen-gorm](https://github.com/infobloxopen/protoc-gen-gorm)
- [x] Dependency injection [Container](https://github.com/sarulabs/di)
- [x] One Step _build/publish/deploy_ with [ko](https://github.com/google/ko)
- [x] BuildInfo with [govvv](https://github.com/ahmetb/govvv)
- [x] Linting with [GolangCI](https://github.com/golangci/golangci-lint) linters aggregator
- [x] CICD Pipelines with [GitHub Actions](https://github.com/features/actions)
- [x] Kubernetes _Matrix Deployment_ with [Kustomize](https://kustomize.io/)
- [ ] Observability
- [ ] Service Mesh with [Istio](https://istio.io/)
- [ ] GraphQL Gateway with [gqlgen](https://gqlgen.com/), [rejoiner](https://github.com/google/rejoiner)

## Getting Started

### Prerequisite

> run following `go get ...` commands outside **this project root** and `$GOPATH`<br/>
> if you get error, try setting `export GO111MODULE=on` befor running `go get ...`

```bash
# fetch micro into $GOPATH
go get github.com/micro/micro
go get github.com/micro/go-micro
# go lang  build/publish/deploy tool
go get github.com/google/ko/cmd/ko
# go better build tool
go get github.com/ahmetb/govvv
# for static check/linter
go get github.com/golangci/golangci-lint/cmd/golangci-lint
# for mac, use brew to install protobuf
brew install protobuf
# GUI Client for GRPC Services
brew cask install bloomrpc
# k8s tool similar to helm
brew install kustomize

# fetch protoc plugins into $GOPATH
go get github.com/golang/protobuf/{proto,protoc-gen-go}
go get github.com/micro/protoc-gen-micro
# go get -u github.com/envoyproxy/protoc-gen-validate
# go get -u github.com/infobloxopen/protoc-gen-gorm
```

> Installing PGV can currently only be done from source:

```bash
go get -d github.com/envoyproxy/protoc-gen-validate
cd ~/go/src/github.com/envoyproxy/protoc-gen-validate
make build
```

> Installing `protoc-gen-gorm` can currently only be done from source:

```
go get -d github.com/infobloxopen/protoc-gen-gorm
cd ~/go/src/github.com/infobloxopen/protoc-gen-gorm
make install
```

### Initial Setup

> (optional) setup your workspace from scratch

```bash
go mod init github.com/xmlking/micro-starter-kit
mkdir srv api fnc

# scaffold modules
micro new --fqdn="account-srv" --type="srv" --gopath=false \
--alias="account" --plugin=registry=kubernetes srv/account

micro new --fqdn="emailer-srv" --type="srv" --gopath=false \
--alias="emailer"  --plugin=registry=kubernetes:broker=nats srv/emailer

micro new --fqdn="greeter-srv" --type="srv" --gopath=false \
--alias="greeter"  --plugin=registry=kubernetes srv/greeter

micro new --fqdn="account-api" --type="api" --gopath=false \
--alias="account" --plugin=registry=kubernetes api/account
```

### Build

```bash
make proto
# silence
make -s proto

make build
```

### Run

#### Database

By default this project use embedded `sqlite3` database. if you want to use **postgreSQL**,

- start **postgres** via `docker-compose` command provided below
- uncommend `postgres` import statement and comment `sqlite` in `plugin.go`
- start micro server with `--configFile=config.dev.postgres.yaml` flag <br/>
  i.e., `go run srv/account/main.go srv/account/plugin.go --configFile=config.dev.postgres.yaml`

```bash
# to start postgres in foreground
docker-compose up postgres
# to stop postgres
docker-compose down
# if needed, remove `postgres_data` volume to recreate database next time, when you start.
docker system prune --volumes
```

#### Services

> Node: `--server_address=localhost:5501x --broker_address=localhost:5502x` required only when you are behind VPN a.k.a `Work From Home`

```bash
# dev mode
make run-account
# or
make run-account ARGS="--server_address=localhost:55012 --broker_address=localhost:55022"
# or
go run srv/account/main.go srv/account/plugin.go \
--configDir deploy/bases/account-srv/config \
--server_address=localhost:55011 --broker_address=localhost:55021

make run-emailer
# or
make run-emailer ARGS="--server_address=localhost:55011 --broker_address=localhost:55021"

# integration tests for config module via CMD
make run TARGET=demo TYPE=cmd
go run cmd/demo/main.go --help
go run cmd/demo/main.go --database_host=1.1.1.1 --database_port=7777

export APP_ENV=production
go run cmd/demo/main.go
```

#### Services with `gRPC` transport

> go-micro by default use `transport=http` and `protocol=mucp`. If you want to switch to `gRPC`, do this

```bash
# run account micro with gRPC transport, NOTE: we also have to add --server_name as it will remove server_name
go run srv/account/main.go srv/account/plugin.go --client=grpc --server=grpc --server_name=account-srv
go run srv/emailer/main.go srv/emailer/plugin.go --client=grpc --server=grpc --server_name=emailer-srv
# we also has to use grpc for gateway and `micro call` cli
go run cmd/micro/main.go cmd/micro/plugin.go --client=grpc --server=grpc api
# go run cmd/micro/main.go cmd/micro/plugin.go --client=grpc --server=grpc api --enable_rpc=true
go run cmd/micro/main.go cmd/micro/plugin.go  --api_handler=rpc  api  --enable_rpc=true

micro --client=grpc call account-srv UserService.List '{ "limit": 10, "page": 1}'
```

### Test

```bash
# Run only Unit tests:
make test-emailer
go test -v -short
go test -v -short ./srv/emailer/service
# Run only Integration Tests: Useful for smoke testing canaries in production.
make inte-emailer
go test -v -run Integration ./srv/emailer/service
```

### UAT Test

> using `micro` CLI

```bash
micro list services
micro get service account-srv
micro get service emailer-srv

# how to start proxy
micro proxy --protocol=grpc
```

#### Test gRPC Directly

> remember to use `micro --client=grpc` when microservices and gateway are using `grpc` transport

```bash
# micro --client=grpc call account-srv UserService.Create \
# '{"username": "sumo", "firstName": "sumo", "lastName": "demo", "email": "sumo@demo.com"}'
micro call  account-srv UserService.Create \
'{"username": "sumo", "firstName": "sumo", "lastName": "demo", "email": "sumo@demo.com"}'
micro call account-srv UserService.Create \
'{"username": "sumo", "firstName": "sumo", "lastName": "demo", "email": "sumo@demo.com"}'
micro call account-srv UserService.List '{}'
micro call account-srv UserService.List '{ "limit": 10, "page": 1}'
micro call account-srv UserService.Get '{"id": "UserIdFromList"}'
micro call account-srv UserService.Exist '{"username": "sumo", "email": "sumo@demo.com"}'
micro call account-srv UserService.Update \
'{"id": "UserIdFromGet", "firstName": "sumoto222","email": "sumo222@demo.com"}'
micro call account-srv UserService.Delete '{ "id": "UserIdFromGet" }'
```

#### Test via Micro Web UI

```bash
# Start Web UI for testing
micro web

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

#### Test via Micro API Gateway

# Start API Gateway

> Start `API Gateway` and run **REST Client** [tests](test/test-rest-api.http)

```bash
# start API Gateway to test via REST-Client
micro api --enable_rpc=true
```

## GitOps

### Deploy

Use `ko`. If you are new to `ko` check out the [ko-demo](https://github.com/xmlking/ko-demo)

Set a registry and make sure you can push to it:

```bash
export PROJECT_ID=ngx-starter-kit
export KO_DOCKER_REPO=gcr.io/${PROJECT_ID}
```

> to publish locally set: `export KO_DOCKER_REPO=ko.local`

Then `apply` like this:

```bash
ko apply -f deploy/
```

To deploy in a different namespace:

```bash
ko -n nondefault apply -f deploy/
```

### Release

> This will publish all of the binary components as container images and create a release.yaml

```bash
# publish to  docker repo ar KO_DOCKER_REPO
ko resolve -P -f deploy/ > release.yaml
# publish to local docker repo
ko resolve -P -L -f deploy/ > release.yaml
```

> run local image

```bash
docker run -it \
-e MICRO_SERVER_ADDRESS=0.0.0.0:8080 \
-e MICRO_BROKER_ADDRESS=0.0.0.0:10001 \
-e MICRO_REGISTRY=mdns \
-e CONFIG_DIR=/var/run/ko \
-e CONFIG_FILE=config.yaml \
-p 8080:8080 -p 10001:10001 ko.local/github.com/xmlking/micro-starter-kit/srv/account
```

### Docker

#### Docker Build

```bash
# build
TYPE=srv
TARGET=account
VERSION=0.0.5-SNAPSHOT
# DOCKER_REGISTRY=gcr.io
DOCKER_CONTEXT_PATH=xmlking
# docker build --force-rm=true --rm=true --no-cache \
docker build --rm \
--build-arg VERSION=$VERSION \
--build-arg TYPE=${TYPE} \
--build-arg TARGET=${TARGET} \
--build-arg DOCKER_REGISTRY=${DOCKER_REGISTRY} \
--build-arg DOCKER_CONTEXT_PATH=${DOCKER_CONTEXT_PATH} \
--build-arg VCS_REF=$(shell git rev-parse --short HEAD) \
--build-arg BUILD_DATE=$(date -u +'%Y-%m-%dT%H:%M:%SZ') \
-t ${DOCKER_REGISTRY:+${DOCKER_REGISTRY}/}${DOCKER_CONTEXT_PATH}/${TARGET}-${TYPE}:${VERSION} .

IMANGE_NAME=${DOCKER_REGISTRY:+${DOCKER_REGISTRY}/}${DOCKER_CONTEXT_PATH}/${TARGET}-${TYPE}:${VERSION}

# push
docker push $IMANGE_NAME

# check
docker inspect  $IMANGE_NAME
# remove temp images after build
docker image prune -f
# Remove dangling images
docker rmi $(docker images -f "dangling=true" -q)
# Remove images tagged with vendor=sumo
docker rmi $(docker images -f "label=org.label-schema.vendor=sumo"  -q)
# Remove all <none> layers
docker rmi $(docker images -a|grep "<none>"|awk '$1=="<none>" {print $3}')
```

#### Docker Run

> run just for testing image...

```bash
docker run -it \
-e MICRO_SERVER_ADDRESS=0.0.0.0:8080 \
-e MICRO_BROKER_ADDRESS=0.0.0.0:10001 \
-e MICRO_REGISTRY=mdns \
-p 8080:8080 -p 10001:10001 $IMANGE_NAME
```

#### Docker Compose Run

Run complete app suite with `docker-compose`

```bash
docker-compose up consul
docker-compose up account-srv
docker-compose up emailer-srv
docker-compose up gateway
docker-compose up account-api
docker-compose up gateway-api
curl "http://localhost:8081/account/AccountService/list?limit=10"
```

#### Kubernetes Run

> run just for testing image in k8s...

```bash
# account-srv
kubectl run --rm mytest --image=xmlking/account-srv:latest \
--env="MICRO_REGISTRY=kubernetes" \
--env="MICRO_SELECTOR=static" \
--env="MICRO_SERVER_ADDRESS=0.0.0.0:8080" \
--env="MICRO_BROKER_ADDRESS=0.0.0.0:10001" \
--restart=Never -it

# gateway
kubectl run --rm mygateway --image=microhq/micro:kubernetes \
--env="MICRO_REGISTRY=kubernetes" \
--env="MICRO_SELECTOR=static" \
--restart=Never -it \
--command ./micro api
```

### Make

using Makefile

> use `-n` flag for `dry-run`, `-s` or '--silent' flag to suppress echoing

```bash
# codegen from proto
make proto
make proto TARGET=account
make proto TARGET=account TYPE=api
make proto-account
make proto-account-api
## generate for protos in shared package
make proto TARGET=shared TYPE=.

# unit tests
make test-account
make test-emailer
make test-account-api
make test-config-shared
make test-demo-cmd

# integration tests
make inte-account
make inte-emailer

# run
make run-account
make run-emailer
make run-account-api
make run-micro-cmd ARGS="--api_address=0.0.0.0:8088 api"
make run-demo-cmd

# Lint
make lint-account-srv

# build
make build VERSION=v0.1.1
make build TARGET=account VERSION=v0.1.1
make build TARGET=account TYPE=srv VERSION=v0.1.1
make build TARGET=emailer TYPE=srv VERSION=v0.1.1
make build TARGET=account TYPE=api VERSION=v0.1.1
make build-account VERSION=v0.1.1
make build-account-api VERSION=v0.1.1

# push tag to git
make release VERSION=v0.1.1

# build docker image
make docker-account VERSION=v0.1.1
make docker-account-srv VERSION=v0.1.1
make docker TARGET=account VERSION=v0.1.1
make docker TARGET=account TYPE=srv VERSION=v0.1.1
make docker TARGET=account DOCKER_REGISTRY=gcr.io DOCKER_CONTEXT_PATH=micro-starter-kit

make docker-emailer-srv
make docker-account-api

# build all docker images
make docker
# publish all microservices images
make docker_push
# remove all previous microservices images and any dangling images
make docker_clean
```

## Reference

1. [examples](https://github.com/micro/examples) - example usage code for micro
2. [microhq](https://github.com/microhq) - a place for prebuilt microservices
3. [explorer](https://micro.mu/explore/) - which aggregates micro based open source projects
4. [micro-plugins](https://github.com/micro/go-plugins) extensible micro plugins

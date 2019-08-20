# Account Service

This is the Account API service

Showcase:

1. use of `*api.Request` `*api.Response` from `github.com/micro/go-micro/api/proto/api.proto`
2. calling external microservices
3. Error Handling with `micro/go-micro/errors`

Generated with

```bash
micro new --namespace=go.micro --type=srv --gopath=false --alias=account srv/account
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.api.account
- Type: api
- Alias: account

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend consul.

```bash
# install consul
brew install consul

# run consul
consul agent -dev
```

## Usage

A Makefile is included for convenience

Build the binary

```bash
make build TARGET=account TYPE=api VERSION=v0.1.1
```

Run the service

> make sure `account-srv` and `emailer-srv` services are running before you start `account-api`

```bash
# Run the go.micro.api.account API Service
go run api/account/main.go api/account/plugin.go

# Run the micro API
micro --client=grpc --server=grpc api --namespace=go.micro.api --handler=api
# or use custom built micro with `grpc`
go run cmd/micro/main.go cmd/micro/plugin.go api  --namespace=go.micro.api --handler=api
# (or) Run Micro Web to test via Web UI
micro web --namespace=go.micro.api

# see service definitions
micro get service go.micro.api.account

# Call go.micro.api.account via API
curl "http://localhost:8080/account/AccountService/list?limit=10"
```

Build a docker image

```bash
make docker TARGET=account TYPE=api VERSION=v0.1.1
```

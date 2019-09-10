# Account Service

This is the Account API service

Showcase:

1. use of `*api.Request` `*api.Response` from `github.com/micro/go-micro/api/proto/api.proto`
2. calling external microservices
3. Error Handling with `micro/go-micro/errors`

## Configuration

- FQDN: account-api
- Type: api
- Alias: account

## Usage

A Makefile is included for convenience

Build the binary

```bash
make build TARGET=account TYPE=api VERSION=v0.1.1
```

Run the service

> make sure `account-srv` and `emailer-srv` services are running before you start `account-api`

```bash
# make run-account-api
# FIXME: `micro api --handler=api` automatically add `go.micro.api` namespace
# so I cannot use short serviceName
make run-account-api ARGS=--server_name=go.micro.api.account
# or
go run api/account/main.go api/account/plugin.go --configDir deploy/bases/account-api/config --server_name=go.micro.api.account

# Run the micro API
micro api --handler=api
# (or) Run Micro Web to trst via Web UI
micro web

# see service definitions
# micro get service account-api
micro get service go.micro.api.account

# Call go.micro.api.account via API
curl "http://localhost:8080/account/AccountService/list?limit=10"
```

Build a docker image

```bash
make docker TARGET=account TYPE=api VERSION=v0.1.1
```

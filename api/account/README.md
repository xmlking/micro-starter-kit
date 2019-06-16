# Account Service

This is the Account service

Generated with

```
micro new github.com/xmlking/micro-starter-kit/api/account --namespace=go.micro --alias=account --type=api
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

```
# install consul
brew install consul

# run consul
consul agent -dev
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./account-api
```

Build a docker image
```
make docker
```

# Start the API

```bash
go run api/account/main.go

micro api --handler=api --namespace=go.micro.api
micro web --namespace=go.micro.api

micro get service go.micro.api.account
curl "http://localhost:8080/account/AccountService/list?limit=10"

curl "http://localhost:63572/account/AccountService/list?limit=10"
```

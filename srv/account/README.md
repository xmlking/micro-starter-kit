# Account Service

This is the Account service

showcase

1. Implements basic CRUD API
2. Multiple handlers, repositories, subscribers  
3. Publishing events
4. GORM data access
5. Config Managment
6. Custom Logging

Generated with

```
micro new --namespace=go.micro --type=srv --gopath=false --alias=account srv/account
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.srv.account
- Type: srv
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
./account-srv
```

Build a docker image
```
make docker
```

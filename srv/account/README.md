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

```bash
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
make build TARGET=account TYPE=srv VERSION=v0.1.1
```

Run the service

```bash
go run srv/account/main.go srv/account/plugin.go
```

Build a docker image

```bash
make docker TARGET=account TYPE=srv VERSION=v0.1.1
```

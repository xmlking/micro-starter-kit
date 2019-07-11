# Emailer Service

This is the Emailer service

Showcase

1. Async service

Generated with

```bash
micro new srv/emailer --namespace=go.micro --alias=emailer --type=srv --gopath=false --plugin=registry=mdns:broker=nats
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.srv.emailer
- Type: srv
- Alias: emailer

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
make build TARGET=emailer TYPE=srv VERSION=v0.1.1
```

Run the service

```bash
go run srv/emailer/main.go srv/emailer/plugin.go
```

Build a docker image

```bash
make docker TARGET=emailer TYPE=srv VERSION=v0.1.1
```

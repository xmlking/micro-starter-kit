# Echo Service

This is the Echo service

Generated with

```
micro new github.com/xmlking/micro-starter-kit/srv/echo --namespace=go.micro --alias=echo --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.srv.echo
- Type: srv
- Alias: echo

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
./echo-srv
```

Build a docker image
```
make docker
```
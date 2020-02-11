# Emailer Service

This is the Emailer service

Showcase

1. Async service

## Configuration

- FQDN: emailer_srv
- Type: srv
- Alias: emailer

## Usage

A Makefile is included for convenience

Build the binary

```bash
make build TARGET=emailer TYPE=srv VERSION=v0.1.1
```

Run the service

```bash
make run-emailer
# or
go run srv/emailer/main.go srv/emailer/plugin.go --configDir deploy/bases/emailer-srv/config
```

Build a docker image

```bash
make docker TARGET=emailer TYPE=srv VERSION=v0.1.1
```

Test the service

```bash
micro publish emailersrv  '{ "to" : "sumo@demo.com", "from": "demo@sumo.com", "subject": "sub", "body": "mybody"  }'
```

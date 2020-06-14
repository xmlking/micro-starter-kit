# Emailer Service

This is the Emailer service

Showcase

1. Async service

## Configuration

- FQDN: mkit.service.emailer
- Type: service
- Alias: emailer

## Usage

A Makefile is included for convenience

Build the binary

```bash
make build TARGET=emailer TYPE=service VERSION=v0.1.1
```

Run the service

```bash
make run-emailer
make run-emailer ARGS="--server_address=:8082"
# or
go run service/emailer/main.go service/emailer/plugin.go
```

Build a docker image

```bash
make docker TARGET=emailer TYPE=service VERSION=v0.1.1
```

Test the service

```bash
micro publish mkit.service.emailer  '{ "to" : "sumo@demo.com", "from": "demo@sumo.com", "subject": "sub", "body": "mybody"  }'
```

# Emailer Service

This is the Emailer service

Showcase

1. Async service

## Configuration

- FQDN: go.micro.srv.emailer
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
go run srv/emailer/main.go srv/emailer/plugin.go
```

Build a docker image

```bash
make docker TARGET=emailer TYPE=srv VERSION=v0.1.1
```

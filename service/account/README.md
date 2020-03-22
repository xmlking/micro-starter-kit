# Account Service

This is the Account service

showcase

1. Implements basic CRUD API
2. Multiple handlers, repositories, subscribers
3. Publishing events
4. GORM data access
5. Config Managment
6. Custom Logging

## Configuration

- FQDN: mkit.service.account
- Type: service
- Alias: account

## Usage

A Makefile is included for convenience

Build the binary

```bash
make build TARGET=account TYPE=service VERSION=v0.1.1
```

Run the service

```bash
make run-account
# or
go run service/account/main.go service/account/plugin.go --configDir deploy/bases/service/account/config
```

Build a docker image

```bash
make docker TARGET=account TYPE=service VERSION=v0.1.1
```

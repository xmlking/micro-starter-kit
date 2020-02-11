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

- FQDN: accountsrv
- Type: srv
- Alias: account

## Usage

A Makefile is included for convenience

Build the binary

```bash
make build TARGET=account TYPE=srv VERSION=v0.1.1
```

Run the service

```bash
make run-account
# or
go run srv/account/main.go srv/account/plugin.go --configDir deploy/bases/account-srv/config
```

Build a docker image

```bash
make docker TARGET=account TYPE=srv VERSION=v0.1.1
```

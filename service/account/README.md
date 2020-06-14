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
make run-account ARGS="--server_address=:8080"
# or
go run service/account/main.go service/account/plugin.go
```

Build a docker image

```bash
make docker TARGET=account TYPE=service VERSION=v0.1.1
```

Test the service

```bash
micro call  mkit.service.account UserService.Create \
'{"username": "sumo", "firstName": "sumo", "lastName": "demo", "email": "sumo@demo.com"}'
micro call mkit.service.account UserService.Create \
'{"username": "sumo1", "firstName": "sumo1", "lastName": "demo1", "email": "sumo1@demo.com"}'
micro call mkit.service.account UserService.List '{}'
```

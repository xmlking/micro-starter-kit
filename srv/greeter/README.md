# Greeter Service

This is the Greeter service

## Configuration

- FQDN: greeter-srv
- Type: srv
- Alias: greeter

## Usage

A Makefile is included for convenience

Build the binary

```bash
make build TARGET=greeter TYPE=srv VERSION=v0.1.1
```

Run the service

```bash
make run-greeter
# or
go run srv/greeter/main.go srv/greeter/plugin.go --configDir deploy/bases/greeter-srv/config
```

Build a docker image

```bash
make docker TARGET=greeter TYPE=srv VERSION=v0.1.1
```

### Test

```bash
# test with grpc cli
grpc_cli call localhost:8080 Greeter.Hello  'name: "sumo"'  --protofiles=srv/greeter/proto/greeter/greeter.proto

micro --client=grpc call greeter-srv Greeter.Hello  '{"name": "John"}'

# in k8s container
./micro --client=grpc call greeter-srv Greeter.Hello  '{"name": "John"}'

curl --request POST \
--url http://localhost:8080/rpc \
--header 'accept: application/json' \
--header 'content-type: application/json' \
--data '{"service": "greeter-srv", "method": "Greeter.Hello","request": {"name": "sumo"}}'
```

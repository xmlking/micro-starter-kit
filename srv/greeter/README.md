# Greeter Service

This is the Greeter service

## Configuration

- FQDN: greeter_srv
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
go run srv/greeter/main.go srv/greeter/plugin.go --configDir deploy/bases/greeter_srv/config
```

Build a docker image

```bash
make docker TARGET=greeter TYPE=srv VERSION=v0.1.1
```

### Test

```bash
# start the server on fixed port
make run-greeter ARGS="--server_address=localhost:8080 --broker_address=localhost:10001"
# make run-greeter ARGS="--server_address=greetersrv:8080 --broker_address=greetersrv:10001"

# test with grpc cli
grpcurl -plaintext -proto srv/greeter/proto/greeter/greeter.proto list
grpcurl -plaintext -proto srv/greeter/proto/greeter/greeter.proto describe
grpcurl -plaintext -proto srv/greeter/proto/greeter/greeter.proto -d '{"name": "sumo"}' localhost:8080  greetersrv.Greeter/Hello
# testing via micro-cli
micro --client=grpc call --metadata trans-id=1234 greetersrv Greeter.Hello  '{"name": "John"}'

# start REST gateway
micro api --enable_rpc=true

# testing via rest proxy
curl --request POST \
--url http://localhost:8080/rpc \
--header 'accept: application/json' \
--header 'content-type: application/json' \
--data '{"service": "greetersrv", "method": "Greeter.Hello","request": {"name": "sumo"}}'
```

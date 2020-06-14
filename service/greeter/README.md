# Greeter Service

This is the Greeter service

## Configuration

- FQDN: mkit.service.greeter
- Type: service
- Alias: greeter

## Usage

Build the binary

```bash
make build TARGET=greeter TYPE=service
# then run with custom env
CONFIGOR_ENV_PREFIX=APP APP_FEATURES_TLS_ENABLED=true ./build/greeter-service
```

Run the service

```bash
make run-greeter
make run-greeter ARGS="--server_address=:8081"
# or
go run service/greeter/main.go service/greeter/plugin.go
```

Build a docker image

```bash
make docker TARGET=greeter TYPE=service VERSION=v0.1.1
```

Test the service

```bash
# start greeter service first
make run-greeter

# test with grpc cli
grpcurl -plaintext -proto service/greeter/proto/greeter/greeter.proto list
grpcurl -plaintext -proto service/greeter/proto/greeter/greeter.proto describe
grpcurl -plaintext -proto service/greeter/proto/greeter/greeter.proto -d '{"name": "sumo"}' localhost:8081  mkit.service.greeter.v1.GreeterService/Hello
# testing via micro-cli
micro call mkit.service.greeter GreeterService.Hello  '{"name": "John"}'

# start REST gateway
micro api --enable_rpc=true

# testing via rest proxy
curl --request POST \
--url http://localhost:8080/rpc \
--header 'accept: application/json' \
--header 'content-type: application/json' \
--data '{"service": "mkit.service.greeter", "method": "GreeterService.Hello","request": {"name": "sumo"}}'
```

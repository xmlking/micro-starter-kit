# Testing

## Unit Test

```bash
make test-unit TARGET=emailer
go test -v -short
go test -v -short ./srv/emailer/service
```

## Integration Test

> Run only Integration Tests: Useful for smoke testing canaries in production.

```bash
make test-inte TARGET=emailer
make test-inte TARGET=emailer TIMEOUT=300ms
go test -v -run Integration ./srv/emailer/service
```

## UAT Test

> we can use one of the option below. They are various options for manual testing

### BloomRPC UI Client

1. add `~/go/src` to import paths, also add any other paths to shared proto files.
2. click (+) and import proto file you want to test.
3. add optional metadata in the JSON format in the `Metadata Section`. e.g., Authorization Headers etc

### gRPC CLI Client

```bash
grpc_cli call localhost:8080 Greeter.Hello  'name: "sumo"'  --protofiles=srv/greeter/proto/greeter/greeter.proto
```

### Micro CLI

> test with gRPC clients such as Micro CLI, BloomRPC or grpcurl

```bash
micro list services
micro get service accountsrv
micro get service emailersrv
```

```bash
micro call  accountsrv UserService.Create \
'{"username": "sumo", "firstName": "sumo", "lastName": "demo", "email": "sumo@demo.com"}'
micro call accountsrv UserService.Create \
'{"username": "sumo1", "firstName": "sumo1", "lastName": "demo1", "email": "sumo1@demo.com"}'
micro call accountsrv UserService.List '{}'
micro call accountsrv UserService.List '{ "limit": 10, "page": 1}'
micro call accountsrv UserService.Get '{"id": "UserIdFromList"}'
micro call accountsrv UserService.Exist '{"username": "sumo", "email": "sumo@demo.com"}'
micro call accountsrv UserService.Update \
'{"id": "UserIdFromGet", "firstName": "sumoto222","email": "sumo222@demo.com"}'
micro call accountsrv UserService.Delete '{ "id": "UserIdFromGet" }'
```

> For k8s: SSH to gateway container and run micro cli....

```bash
kubectl exec -it -c srv gateway-srv-c86cb8667-g2rmc -- busybox sh
micro call accountsrv UserService.List '{}'
```

### Micro Web UI

```bash
# Start Web UI for testing
micro web

open http://localhost:8082
```

> create new user from `Micro Web UI` and see if an email is send

```json
{
  "username": "sumo",
  "firstName": "sumo",
  "lastName": "demo",
  "email": "sumo@demo.com"
}
```

### Micro API Gateway

> Start API Gateway

Start `API Gateway` and run **REST Client** [tests](../e2e/test-rest-api.http)

```bash
micro  api --enable_rpc=true
```

## E2E Testing

> Assume, you are running all microservices on [local k8s cluster](../e2e/README.md) with one of the profiles(`e2e`, `production`)

### E2E tests with tools

```bash
# with `grpc_cli`
# micro proxy --protocol=grpc
grpc_cli call localhost:8081 Greeter.Hello  'name: "sumo"'  --protofiles=srv/greeter/proto/greeter/greeter.proto
# with Micro CLI
MICRO_PROXY_ADDRESS=localhost:8081 micro list services
MICRO_PROXY_ADDRESS=localhost:8081 micro --client=grpc call greetersrv Greeter.Hello  '{"name": "John"}'
MICRO_PROXY_ADDRESS=localhost:8081 micro call accountsrv UserService.List '{}'
MICRO_PROXY_ADDRESS=localhost:8081  micro health greetersrv
```

### E2E tests via code

```bash
MICRO_PROXY_ADDRESS="localhost:8081" \
make test-e2e
```

## Fuzzing

TODO

> fuzz testing with [fuzzit](https://fuzzit.dev/2019/10/02/how-to-fuzz-go-code-with-go-fuzz-continuously/)

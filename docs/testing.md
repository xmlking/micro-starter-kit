# Testing

## Unit Test

```bash
make test-emailer
go test -v -short
go test -v -short ./srv/emailer/service
```

## Integration Test

> Run only Integration Tests: Useful for smoke testing canaries in production.

```bash
make inte-emailer
go test -v -run Integration ./srv/emailer/service
```

## UAT Test

> we can use one of the option below.

### Micro CLI

> test with gRPC clients such as Micro CLI, BloomRPC or grpcurl

```bash
micro list services
micro get service account-srv
micro get service emailer-srv

# how to start proxy
micro proxy --protocol=grpc
```

```bash
## local build has gRPC by default.
# ./build/micro call  account-srv UserService.Create \
#   '{"username": "sumo", "firstName": "sumo", "lastName": "demo", "email": "sumo@demo.com"}'
micro call  account-srv UserService.Create \
'{"username": "sumo", "firstName": "sumo", "lastName": "demo", "email": "sumo@demo.com"}'
micro call account-srv UserService.Create \
'{"username": "sumo", "firstName": "sumo", "lastName": "demo", "email": "sumo@demo.com"}'
micro call account-srv UserService.List '{}'
micro call account-srv UserService.List '{ "limit": 10, "page": 1}'
micro call account-srv UserService.Get '{"id": "UserIdFromList"}'
micro  call account-srv UserService.Exist '{"username": "sumo", "email": "sumo@demo.com"}'
micro call account-srv UserService.Update \
'{"id": "UserIdFromGet", "firstName": "sumoto222","email": "sumo222@demo.com"}'
micro call account-srv UserService.Delete '{ "id": "UserIdFromGet" }'
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

#### Micro API Gateway

> Start API Gateway

Start `API Gateway` and run **REST Client** [tests](test/test-rest-api.http)

```bash
# start local micro (grpc pre-loaded micro)
make run-micro-cmd ARGS="api --enable_rpc=true"
# (or)
go run cmd/micro/main.go  api --enable_rpc=true

# (or) start global micro
micro  api --enable_rpc=true
```

## Fuzzing

TODO

> fuzz testing with [fuzzit](https://fuzzit.dev/2019/10/02/how-to-fuzz-go-code-with-go-fuzz-continuously/)

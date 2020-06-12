# grpc-zerolog

grpc-zerolog is a simple implementation of grpclog.LoggerV2 interface using zerolog. Use this to log the internal actions of a gRPC server or client.

## Usage

Add the following before you `grpc.Dial` either a client or server.

```go
logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
logger = logger.With().Str("module", "grpc").Logger()

grpclog.SetLoggerV2(grpc.New(logger))
```

Start your server/client with the following environment variable.

`GRPC_GO_LOG_VERBOSITY_LEVEL=info GRPC_GO_LOG_VERBOSITY_LEVEL=99 go run cmd/account/main.go`
 

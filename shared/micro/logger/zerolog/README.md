# Zerolog

[Zerolog](https://github.com/rs/zerolog) logger implementation for __go-micro__ [meta logger](https://github.com/micro/go-micro/tree/master/logger).

## Usage

```go
func ExampleWithOut() {
  l := zero.NewLogger(zero.WithOut(os.Stdout), zero.WithLevel(logger.DebugLevel))

  l.Logf(logger.InfoLevel, "testing: %s", "logf")

  // Output:
  // {"level":"info","message":"testing: logf"}
}
```

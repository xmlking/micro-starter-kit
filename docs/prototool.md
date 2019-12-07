# prototool

Protos should follow [Uber Protobuf Style Guide V2](https://github.com/uber/prototool/blob/dev/style/README.md)

> Generating Go gRPC from protobuf definitions

```bash
prototool generate
```

> Lint

```bash
prototool lint  ./srv/greeter
```

> generate-ignores

```bash
prototool lint e2e  --generate-ignores
prototool lint srv/greeter  --generate-ignores
```

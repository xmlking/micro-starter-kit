# HOWTO

- How to update 3rd party dependencies?

```bash
go get -u # to use the latest minor or patch releases
go get -u=patch # to use the latest patch releases
go mod tidy
```

- How to Prepare for a Release?

```bash
go mod tidy
go test all
```

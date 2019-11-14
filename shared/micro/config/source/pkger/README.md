# pkger

pkger plugin for `go-config`

## Prerequisites

> Install `pkger` cli

```bash
go get github.com/markbates/pkger/cmd/pkger
pkger -h
```

> generating `pkged.go` with all files in `/config` for production envelopment

```bash
pkger -o srv/greeter -include /config
```

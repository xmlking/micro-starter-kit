# pkger

pkger plugin for `go-config`

### Prerequisites

> Install `pkger` cli

```bash
go install github.com/markbates/pkger/cmd/pkger
pkger -h
```

### Packager

> generating `pkged.go` with all files in `/config` for production build

```bash
pkger -o srv/greeter -include /config
# (or)
make make pkger-greeter
# (or)
make pkger
```

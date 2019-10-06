# Prerequisites

You should have:

1. working **cntlm** (NTLM local Proxy)
2. **golang** installed via **brew**

> run following `go get ...` commands outside **this project root** and outside `$GOPATH` i.e, `~/go`<br/>
> if you get error, try setting `export GO111MODULE=on` befor running `go get ...`

> Lets build and install `grpc` pre-loaded **Micro CLI** from [here](../cmd/micro/README.md#Build) instead of official **Micro CLI**

```bash
# fetch micro into $GOPATH
# build and install your own grpc pre-loaded micro-cli
# go get github.com/micro/micro

# go lang  build/publish/deploy tool
go get github.com/google/ko/cmd/ko
# go better build tool
go get github.com/ahmetb/govvv
# for static check/linter
go get github.com/golangci/golangci-lint/cmd/golangci-lint
# for mac, use brew to install protobuf
brew install protobuf
# GUI Client for GRPC Services
brew cask install bloomrpc
# k8s tool similar to helm
brew install kustomize

# fetch protoc plugins into $GOPATH
go get github.com/golang/protobuf/{proto,protoc-gen-go}
go get github.com/micro/protoc-gen-micro
# go get -u github.com/envoyproxy/protoc-gen-validate
# go get -u github.com/infobloxopen/protoc-gen-gorm
```

> Installing PGV can currently only be done from source:

```bash
go get -d github.com/envoyproxy/protoc-gen-validate
cd ~/go/src/github.com/envoyproxy/protoc-gen-validate
make build
```

> Installing `protoc-gen-gorm` can currently only be done from source:

```bash
go get -d github.com/infobloxopen/protoc-gen-gorm
cd ~/go/src/github.com/infobloxopen/protoc-gen-gorm
make install
```

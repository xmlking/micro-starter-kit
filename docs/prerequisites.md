# Prerequisites

You should have:

**golang** installed via **brew**

### third-party tools

```bash
# for mac, use brew to install protobuf
brew install protobuf
# k8s tool similar to helm  (optional)
brew install kustomize
# kubeval - validate one or more Kubernetes config files(optional)
brew tap instrumenta/instrumenta
brew install kubeval
# Manage Your lk8s In Style!
brew install derailed/k9s/k9s
# grpc cli client (optional)
brew install grpcurl
# bloomrpc is a UI client for gRPC (optional)
# install `bloomrpc` via `brew` into ~/Applications)
brew cask install --appdir=~/Applications bloomrpc
```

### third-party golang tools

```bash
# go better build tool
GO111MODULE=off go get github.com/ahmetb/govvv
# for static check/linter
GO111MODULE=off go get github.com/golangci/golangci-lint/cmd/golangci-lint
# kind - kubernetes in docker (optional)
GO111MODULE=on go get sigs.k8s.io/kind
# go lang  build/publish/deploy tool (optional)
GO111MODULE=off go get github.com/google/ko/cmd/ko
# other way to get latest kustomize
GO111MODULE=on go get sigs.k8s.io/kustomize/kustomize/v3@v3.3.0
# pkger cli
go install github.com/markbates/pkger/cmd/pkger

# fetch protoc plugins into $GOPATH
GO111MODULE=off go get github.com/golang/protobuf/{proto,protoc-gen-go}
GO111MODULE=off go get github.com/micro/protoc-gen-micro
# GO111MODULE=off go get -u github.com/envoyproxy/protoc-gen-validate
# GO111MODULE=off go get -u github.com/infobloxopen/protoc-gen-gorm
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

> Installing `micro-cli`<br/>
> instead of using default micro-cli, pull custom build `micro-cli` with `gRPC` and `CORS` enabled

```bash
# getting official micro cli
## GO111MODULE=on go get github.com/micro/micro
# lets use custom made micro with gRPC enabled by default...
# GO111MODULE=on go get github.com/xmlking/micro@master
# Temp workaround...
GO111MODULE=off go get -d github.com/xmlking/micro
cd ~/go/src/github.com/xmlking/micro
git pull --all
go install ./...
```

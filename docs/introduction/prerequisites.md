# Prerequisites

You should have:

**golang** installed via **brew**

## third-party

### third-party tools

```bash
# github CLI
brew install hub
# for mac, use brew to install protobuf
brew install protobuf
# VS Code plugin `vscode-proto3` need clang-format
brew install clang-format
# k8s tool similar to helm  (optional)
# generate fill k8s yaml files from overlays
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
# gRPC mock server for testing
yarn global add bloomrpc-mock
# for etcdctl
brew install etcd
# CHANGELOG generator
brew tap git-chglog/git-chglog
brew install git-chglog
# buf: proto tool https://buf.build/docs/tour-1
brew tap bufbuild/buf
brew install buf
```

### third-party golang tools

```bash
# go better build tool
GO111MODULE=off go get github.com/ahmetb/govvv
# for static check/linter
GO111MODULE=off go get github.com/golangci/golangci-lint/cmd/golangci-lint
# linter and tool for proto files
# (if you use brew to install buf, skip next line)
GO111MODULE=on go get github.com/bufbuild/buf/cmd/buf
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
GO111MODULE=on go get github.com/micro/micro/v2/cmd/protoc-gen-micro@master
GO111MODULE=on go get github.com/gogo/protobuf/protoc-gen-gofast
# GO111MODULE=off go get -u github.com/envoyproxy/protoc-gen-validate
# GO111MODULE=off go get -u github.com/infobloxopen/protoc-gen-gorm
# getting official micro cli
GO111MODULE=on go get github.com/micro/micro/v2
# goup checks if there are any updates for imports in your module.
# the main purpose is using it as a linter in continuous integration or in development process.
# Usage: goup -v -m ./...
GO111MODULE=on go get github.com/rvflash/goup
```

> Installing PGV can currently only be done from source:

```bash
go get -d github.com/envoyproxy/protoc-gen-validate
cd ~/go/src/github.com/envoyproxy/protoc-gen-validate
git pull
make build
```

> Installing `protoc-gen-gorm` can currently only be done from source:

```bash
go get -d github.com/infobloxopen/protoc-gen-gorm
cd ~/go/src/github.com/infobloxopen/protoc-gen-gorm
git pull
make install
```


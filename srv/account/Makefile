VERSION:=$(shell git describe --tags || echo "HEAD")
GOPATH:=$(shell go env GOPATH)

.PHONY: proto
proto:
	protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. proto/account/account.proto

.PHONY: build
build: proto

	go build -o account-srv main.go plugin.go

.PHONY: test
test:
	APP_ENV=test go test -v ./... -cover

.PHONY: fmt
fmt:
	go fmt -v ./...

.PHONY: docker
docker:
	docker build . -t account-srv:$(VERSION)

.PHONY: clean
clean:
	rm -f  ./build/account-srv

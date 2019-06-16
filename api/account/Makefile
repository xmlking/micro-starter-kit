
VERSION:=$(shell cat ../../VERSION)
GOPATH:=$(shell go env GOPATH)


.PHONY: proto
proto:
	protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. proto/account/account.proto

.PHONY: build
build: proto

	go build -o account-api main.go plugin.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t account-api:$(VERSION)

.PHONY: clean
clean:
	rm -f  ./account-api
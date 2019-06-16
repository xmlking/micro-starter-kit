VERSION:=$(shell cat ../../VERSION)
GOPATH:=$(shell go env GOPATH)


.PHONY: proto
proto:
	protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. proto/emailer/emailer.proto

.PHONY: build
build: proto

	go build -o emailer-srv main.go plugin.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t emailer-srv:$(VERSION)

.PHONY: clean
clean:
	rm -f  ./account-srv
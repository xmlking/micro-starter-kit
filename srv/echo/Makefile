
VERSION:=$(shell cat ../../VERSION)
GOPATH:=$(shell go env GOPATH)


.PHONY: proto
proto:
	protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. proto/echo/echo.proto

.PHONY: build
build: proto

	go build -o echo-srv main.go plugin.go

.PHONY: test
test:
	APP_ENV=test go test -v ./... -cover

.PHONY: fmt
fmt:
	go fmt -v ./...

.PHONY: docker
docker:
	docker build . -t echo-srv:$(VERSION)

.PHONY: clean
clean:
	rm -f  ./account-srv
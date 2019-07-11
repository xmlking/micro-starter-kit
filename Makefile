VERSION		:= $(shell git describe --tags || echo "HEAD")
GOPATH		:= $(shell go env GOPATH)
HAS_GOVVV	:= $(shell command -v govvv 2> /dev/null)
HAS_KO		:= $(shell command -v ko 2> /dev/null)
# Type of service e.g api, fnc, srv, web (default: "srv")
TYPE:= "srv"
override TYPES:= api srv
# Target for running the action
TARGET:=
override VERSION_PACKAGE=$(shell go list ./shared/config)
BUILD_FLAGS=$(shell govvv -flags -version $(VERSION) -pkg $(VERSION_PACKAGE))

.PHONY: proto data build

tools:
	@echo "==> Installing dev tools"
	# go install github.com/ahmetb/govvv
	# go install github.com/google/ko/cmd/ko

proto:
ifndef TARGET
	@for d in $(TYPES); do \
		for f in $$d/**/proto/**/*.proto; do \
			protoc --proto_path=.:${GOPATH}/src --micro_out=. --go_out=. --validate_out=lang=go:. $$f; \
			echo compiled: $$f; \
		done \
	done
else
	@for f in ${TYPE}/${TARGET}/proto/**/*.proto; do \
		protoc --proto_path=.:${GOPATH}/src --micro_out=. --go_out=. --validate_out=lang=go:. $$f; \
		echo compiled: $$f; \
	done
endif

lint:
	./scripts/lint.sh

build:
ifndef HAS_GOVVV
	$(error "No govvv in PATH". Please install via 'go install github.com/ahmetb/govvv'")
endif
ifndef TARGET
	# @for i in $(shell ls -1  srv); do echo  $${i}; done
	@for type in $(TYPES); do \
		echo "Building $${type}..."; \
		for _target in $${type}/*/; do \
			temp=$${_target%%/}; target=$${temp#*/}; \
			echo "\tBuilding $${target}-$${type}"; \
			CGO_ENABLED=0 GOOS=linux go build -o build/$${target}-$${type} -a -ldflags "-w -s ${BUILD_FLAGS}" ./$${type}/$${target}; \
		done \
	done
else
	@echo "Building ${TARGET}-${TYPE}"; \
	go build -o  build/${TARGET}-${TYPE} -a -ldflags "-w -s ${BUILD_FLAGS}" ./${TYPE}/${TARGET};
	# CGO_ENABLED=0 GOOS=linux go build -o  build/${TARGET}-${TYPE} -a -ldflags "-w -s ${BUILD_FLAGS}" ./${TYPE}/${TARGET};
endif

test:
	go test -v ./... -cover

data:
	go-bindata -o data/bindata.go -pkg data data/*.json

run:
	docker-compose build
	docker-compose up

release:
	git tag -a $(VERSION) -m "Release" || true
	git push origin $(VERSION)
	# goreleaser --rm-dist

# delegate to sub projects
clean:
	@for d in $(TYPES); do \
		for sd in $$d/*; do \
			$(MAKE) -C $$sd $(MAKECMDGOALS); \
			echo cleaned: $$sd; \
		done \
	done

update_deps:
	go mod verify
	go mod tidy


docker:
ifndef TARGET
	$(error "no  TARGET. example usage: make docker TARGET=account")
endif
	@docker build --rm \
	--build-arg VERSION=$(VERSION) \
	--build-arg BUILD_PKG=./${TYPE}/${TARGET} \
	--build-arg BUILD_DATE=$(shell date +%FT%T%Z) \
	-t xmlking/${TARGET}-${TYPE} .

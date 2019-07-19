VERSION		:= $(shell git describe --tags || echo "HEAD")
GOPATH		:= $(shell go env GOPATH)
HAS_GOVVV	:= $(shell command -v govvv 2> /dev/null)
HAS_KO		:= $(shell command -v ko 2> /dev/null)
# Type of service e.g api, fnc, srv, web (default: "srv")
TYPE = $(or $(word 2,$(subst -, ,$*)), srv)
override TYPES:= api srv
# Target for running the action
TARGET = $(word 1,$(subst -, ,$*))

override VERSION_PACKAGE = $(shell go list ./shared/config)
BUILD_FLAGS = $(shell govvv -flags -version $(VERSION) -pkg $(VERSION_PACKAGE))

# $(warning TYPES = $(TYPE), TARGET = $(TARGET))
# $(warning VERSION = $(VERSION), HAS_GOVVV = $(HAS_GOVVV), HAS_KO = $(HAS_KO))
# $(warning VERSION_PACKAGE = $(VERSION_PACKAGE), BUILD_FLAGS = $(BUILD_FLAGS))

.PHONY: proto proto-% lint build build-% test run release clean update_deps docker

tools:
	@echo "==> Installing dev tools"
	# go install github.com/ahmetb/govvv
	# go install github.com/google/ko/cmd/ko

proto proto-%:
	@if [ -z $(TARGET) ]; then \
		for d in $(TYPES); do \
			for f in $$d/**/proto/**/*.proto; do \
				protoc --proto_path=.:${GOPATH}/src \
				--go_out=paths=source_relative:. \
				--micro_out=paths=source_relative:. \
				--validate_out=lang=go,paths=source_relative:. $$f; \
				echo compiled: $$f; \
			done \
		done \
	else \
		for f in ${TYPE}/${TARGET}/proto/**/*.proto; do \
			protoc --proto_path=.:${GOPATH}/src \
			--go_out=paths=source_relative:. \
			--micro_out=paths=source_relative:. \
			--validate_out=lang=go,paths=source_relative:. $$f; \
			echo compiled: $$f; \
		done \
	fi

lint:
	./scripts/lint.sh

build build-%:
ifndef HAS_GOVVV
	$(error "No govvv in PATH". Please install via 'go install github.com/ahmetb/govvv'")
endif
	@if [ -z $(TARGET) ]; then \
		for type in $(TYPES); do \
			echo "Building $${type}..."; \
			for _target in $${type}/*/; do \
				temp=$${_target%%/}; target=$${temp#*/}; \
				echo "\tBuilding $${target}-$${type}"; \
				CGO_ENABLED=0 GOOS=linux go build -o build/$${target}-$${type} -a -ldflags "-w -s ${BUILD_FLAGS}" ./$${type}/$${target}; \
			done \
		done \
	else \
		echo "Building ${TARGET}-${TYPE}"; \
		go build -o  build/${TARGET}-${TYPE} -a -ldflags "-w -s ${BUILD_FLAGS}" ./${TYPE}/${TARGET}; \
	fi

test:
	go test -v ./... -cover

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

docker docker-%:
	@if [ -z $(TARGET) ]; then \
		echo "no  TARGET. example usage: make docker TARGET=account"; \
	else \
		echo "in else"; \
		docker build --rm \
		--build-arg VERSION=$(VERSION) \
		--build-arg BUILD_PKG=./${TYPE}/${TARGET} \
		--build-arg IMANGE_NAME=xmlking/${TARGET}-${TYPE} \
		--build-arg BUILD_DATE=$(shell date +%FT%T%Z) \
		-t xmlking/${TARGET}-${TYPE} .; \
	fi

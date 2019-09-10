VERSION					:= $(shell git describe --tags || echo "HEAD")
GOPATH					:= $(shell go env GOPATH)
HAS_GOVVV				:= $(shell command -v govvv 2> /dev/null)
HAS_KO					:= $(shell command -v ko 2> /dev/null)
CODECOV_FILE 		:= build/coverage.txt
# DOCKER_CONTEXT_PATH 			:= my_project_id/micro-starter-kit
DOCKER_CONTEXT_PATH 			:= xmlking

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

.PHONY: proto proto-% lint lint-% build build-% test test-% inte inte-% run run-% release clean update_deps docker docker-%

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
				--gorm_out=paths=source_relative:. \
				--validate_out=lang=go,paths=source_relative:. $$f; \
				echo compiled: $$f; \
			done \
		done \
	else \
		for f in ${TYPE}/${TARGET}/proto/**/*.proto; do \
			protoc --proto_path=.:${GOPATH}/src \
			--go_out=paths=source_relative:. \
			--micro_out=paths=source_relative:. \
			--gorm_out=paths=source_relative:. \
			--validate_out=lang=go,paths=source_relative:. $$f; \
			echo compiled: $$f; \
		done \
	fi

lint lint-%:
	@if [ -z $(TARGET) ]; then \
		echo "Linting all"; \
		${GOPATH}/bin/golangci-lint run ./... ; \
	else \
		echo "Linting ${TARGET}-${TYPE}..."; \
		${GOPATH}/bin/golangci-lint run ./${TYPE}/${TARGET}/... ; \
	fi

build build-%:
ifndef HAS_GOVVV
	$(error "No govvv in PATH". Please install via 'go install github.com/ahmetb/govvv'")
endif
	@if [ -z $(TARGET) ]; then \
		for type in $(TYPES); do \
			echo "Building Type: $${type}..."; \
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

test test-%:
	@if [ -z $(TARGET) ]; then \
		echo "Testing all"; \
		go test -short -race ./... -coverprofile=${CODECOV_FILE} -covermode=atomic; \
	else \
		echo "Testing ${TARGET}-${TYPE}..."; \
		go test -v -short  -race ./${TYPE}/${TARGET}/... -coverprofile=${CODECOV_FILE} -covermode=atomic; \
	fi

inte inte-%:
	@if [ -z $(TARGET) ]; then \
		echo "Integration testing all"; \
		go test -run Integration ./... ; \
	else \
		echo "Integration testing ${TARGET}-${TYPE}..."; \
		go test -v -run Integration ./${TYPE}/${TARGET}/... ; \
	fi

run run-%:
	@if [ -z $(TARGET) ]; then \
		echo "no  TARGET. example usage: make test TARGET=account"; \
	else \
		go run  ./${TYPE}/${TARGET} ${ARGS}; \
	fi

release:
	git tag -a $(VERSION) -m "Release" || true
	git push origin $(VERSION)
	# goreleaser --rm-dist

clean:
	@for d in ./build/*-{srv,api}; do \
		echo "Deleting $$d;"; \
		rm -f $$d; \
	done

update_deps:
	go mod verify
	go mod tidy

docker docker-%:
	@if [ -z $(TARGET) ]; then \
		echo "Building images for all services..."; \
		echo "no  TARGET. example usage: make docker TARGET=account"; \
		for type in $(TYPES); do \
			echo "Building Type: $${type}..."; \
			for _target in $${type}/*/; do \
				temp=$${_target%%/}; target=$${temp#*/}; \
				echo "Building Image $${target}-$${type}..."; \
				docker build --rm \
				--build-arg VERSION=$(VERSION) \
				--build-arg TYPE=$${type} \
				--build-arg TARGET=$${target} \
				--build-arg DOCKER_REGISTRY=${DOCKER_REGISTRY} \
				--build-arg DOCKER_CONTEXT_PATH=${DOCKER_CONTEXT_PATH} \
				--build-arg BUILD_DATE=$(shell date +%FT%T%Z) \
				-t $${DOCKER_REGISTRY:+${DOCKER_REGISTRY}/}${DOCKER_CONTEXT_PATH}/$${target}-$${type}:$(VERSION) .; \
			done \
		done \
	else \
		echo "Building image for ${TARGET}-${TYPE}..."; \
		docker build --rm \
		--build-arg VERSION=$(VERSION) \
		--build-arg TYPE=${TYPE} \
		--build-arg TARGET=${TARGET} \
		--build-arg DOCKER_REGISTRY=${DOCKER_REGISTRY} \
		--build-arg DOCKER_CONTEXT_PATH=${DOCKER_CONTEXT_PATH} \
		--build-arg BUILD_DATE=$(shell date +%FT%T%Z) \
		-t $${DOCKER_REGISTRY:+${DOCKER_REGISTRY}/}${DOCKER_CONTEXT_PATH}/${TARGET}-${TYPE}:$(VERSION) .; \
	fi

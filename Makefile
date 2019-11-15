# Usage:
# make        	# compile all binary
# make clean  	# remove ALL binaries and objects
# make release  # add git TAG and push
GITHUB_REPO_OWNER 				:= xmlking
GITHUB_REPO_NAME 					:= micro-starter-kit
GITHUB_RELEASES_UI_URL 		:= https://github.com/$(GITHUB_REPO_OWNER)/$(GITHUB_REPO_NAME)/releases
GITHUB_RELEASES_API_URL 	:= https://api.github.com/repos/$(GITHUB_REPO_OWNER)/$(GITHUB_REPO_NAME)/releases
GITHUB_RELEASE_ASSET_URL	:= https://uploads.github.com/repos/$(GITHUB_REPO_OWNER)/$(GITHUB_REPO_NAME)/releases
GITHUB_DEPLOY_API_URL			:= https://api.github.com/repos/$(GITHUB_REPO_OWNER)/$(GITHUB_REPO_NAME)/deployments
DOCKER_REGISTRY 					:= docker.pkg.github.com
DOCKER_CONTEXT_PATH 			:= $(GITHUB_REPO_OWNER)/$(GITHUB_REPO_NAME)
# DOCKER_REGISTRY 					:= us.gcr.io
# DOCKER_CONTEXT_PATH 			:= xmlking

VERSION					:= $(shell git describe --tags || echo "HEAD")
CURRENT_BRANCH  := $(shell git rev-parse --abbrev-ref HEAD)
GOPATH					:= $(shell go env GOPATH)
HAS_GOVVV				:= $(shell command -v govvv 2> /dev/null)
HAS_PKGER				:= $(shell command -v pkger 2> /dev/null)
GIT_DIRTY 			:= $(shell git status --porcelain 2> /dev/null)
HAS_KO					:= $(shell command -v ko 2> /dev/null)
CODECOV_FILE 		:= build/coverage.txt
TIMEOUT  				:= 60s

# Type of service e.g api, fnc, srv, web (default: "srv")
TYPE = $(or $(word 2,$(subst -, ,$*)), srv)
override TYPES:= srv
# Target for running the action
TARGET = $(word 1,$(subst -, ,$*))

override VERSION_PACKAGE = $(shell go list ./shared/config)
BUILD_FLAGS = $(shell govvv -flags -version $(VERSION) -pkg $(VERSION_PACKAGE))

# $(warning TYPES = $(TYPE), TARGET = $(TARGET))
# $(warning VERSION = $(VERSION), HAS_GOVVV = $(HAS_GOVVV), HAS_KO = $(HAS_KO))
# $(warning VERSION_PACKAGE = $(VERSION_PACKAGE), BUILD_FLAGS = $(BUILD_FLAGS))

.PHONY: all tools, check_dirty, clean, update_deps
.PHONY: proto proto-%
.PHONY: lint lint-%
.PHONY: pkger pkger-%
.PHONY: build build-%
.PHONY: run run-%
.PHONY: docker_clean docker docker-% docker_push
.PHONY: kustomize build/kustomize
.PHONY: release/draft release/publish
.PHONY: deploy/e2e deploy/prod

all: build

tools:
	@echo "==> Installing dev tools"
	# go install github.com/ahmetb/govvv
	# go install github.com/markbates/pkger/cmd/pkger

check_dirty:
ifdef GIT_DIRTY
	$(error "Won't run on a dirty working copy. Commit or stash and try again.")
endif

clean:
	@for d in ./build/*-srv; do \
		echo "Deleting $$d;"; \
		rm -f $$d; \
	done
	@for f in */*/pkged.go ; do \
		echo "Deleting $$f;"; \
		rm -f $$f; \
	done

update_deps:
	go mod verify
	go mod tidy

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
		${GOPATH}/bin/golangci-lint run ./... --deadline=5m; \
	else \
		echo "Linting ${TARGET}-${TYPE}..."; \
		${GOPATH}/bin/golangci-lint run ./${TYPE}/${TARGET}/... ; \
	fi

pkger pkger-%:
ifndef HAS_PKGER
	$(error "No pkger in PATH". Please install via 'go install github.com/markbates/pkger/cmd/pkger'")
endif
	@if [ -z $(TARGET) ]; then \
		for type in $(TYPES); do \
			echo "Packaging config for Type: $${type}..."; \
			for _target in $${type}/*/; do \
				temp=$${_target%%/}; target=$${temp#*/}; \
				echo "\tPackaging config for $${target}-$${type}"; \
				${GOPATH}/bin/pkger -o $${type}/$${target} -include /config; \
			done \
		done \
	else \
		echo "Packaging config for ${TARGET}-${TYPE}..."; \
		${GOPATH}/bin/pkger -o ${TYPE}/${TARGET} -include /config ; \
	fi

build build-%: pkger-%
ifndef HAS_GOVVV
	$(error "No govvv in PATH". Please install via 'go install github.com/ahmetb/govvv'")
endif
	@if [ -z $(TARGET) ]; then \
		for type in $(TYPES); do \
			echo "Building Type: $${type}..."; \
			for _target in $${type}/*/; do \
				temp=$${_target%%/}; target=$${temp#*/}; \
				echo "\tBuilding $${target}-$${type}"; \
				CGO_ENABLED=0 GOOS=linux go build -o build/$${target}-$${type} -a -trimpath -ldflags "-w -s ${BUILD_FLAGS}" ./$${type}/$${target}; \
			done \
		done \
	else \
		echo "Building ${TARGET}-${TYPE}"; \
		go build -o  build/${TARGET}-${TYPE} -a -trimpath -ldflags "-w -s ${BUILD_FLAGS}" ./${TYPE}/${TARGET}; \
	fi

TEST_TARGETS := test-default test-bench test-unit test-inte test-e2e test-race test-cover
.PHONY: $(TEST_TARGETS) check test tests
test-bench:   	ARGS=-run=__absolutelynothing__ -bench=. ## Run benchmarks
test-unit:   		ARGS=-short        					## Run only unit tests
test-inte:   		ARGS=-run Integration       ## Run only integration tests
test-e2e:   		ARGS=-run E2E       				## Run only E2E tests
test-race:    	ARGS=-race         					## Run tests with race detector
test-cover:   	ARGS=-cover -short -coverprofile=${CODECOV_FILE} -covermode=atomic ## Run tests in verbose mode with coverage reporting
$(TEST_TARGETS): NAME=$(MAKECMDGOALS:test-%=%)
$(TEST_TARGETS): test
check test tests:
	@if [ -z $(TARGET) ]; then \
		echo "Running $(NAME:%=% )tests for all"; \
		go test -timeout $(TIMEOUT) $(ARGS) ./... ; \
	else \
		echo "Running $(NAME:%=% )tests for ${TARGET}-${TYPE}"; \
		go test -timeout $(TIMEOUT) -v $(ARGS) ./${TYPE}/${TARGET}/... ; \
	fi

run run-%:
	@if [ -z $(TARGET) ]; then \
		echo "no  TARGET. example usage: make test TARGET=account"; \
	else \
		go run  ./${TYPE}/${TARGET} ${ARGS}; \
	fi

release/draft: check_dirty
	@echo Publishing Draft: $(VERSION)
	@git tag -a $(VERSION) -m "[skip ci] Release: $(VERSION)" || true
	@git push origin $(VERSION)
	@echo "\n\nPlease inspect the release and run `make release/publish` if it looks good"
	@open "$(GITHUB_RELEASES_UI_URL)/$(VERSION)"

release/publish:
	@echo Publishing Release: $(VERSION)

deploy/e2e:
	@curl -H "Content-Type: application/json" \
		-H "Accept: application/vnd.github.ant-man-preview+json"  \
		-H "Authorization: token $(GITHUB_TOKEN)" \
    -XPOST $(GITHUB_DEPLOY_API_URL) \
    -d '{"ref": "develop", "environment": "e2e", "payload": { "what": "deployment for e2e testing"}}'

deploy/prod:
	@curl -H "Content-Type: application/json" \
		-H "Accept: application/vnd.github.ant-man-preview+json"  \
		-H "Authorization: token $(GITHUB_TOKEN)" \
    -XPOST $(GITHUB_DEPLOY_API_URL) \
    -d '{"ref": "develop", "environment": "production", "payload": { "what": "production deployment to GKE"}}'

# TODO: DOCKER_BUILDKIT=1 docker build --rm
docker docker-%:
	@if [ -z $(TARGET) ]; then \
		echo "Building images for all services..."; \
		for type in $(TYPES); do \
			echo "Building Type: $${type}..."; \
			for _target in $${type}/*/; do \
				temp=$${_target%%/}; target=$${temp#*/}; \
				echo "Building Image $${target}-$${type}..."; \
				docker build --rm \
				--build-arg BUILDKIT_INLINE_CACHE=1 \
				--build-arg VERSION=$(VERSION) \
				--build-arg TYPE=$${type} \
				--build-arg TARGET=$${target} \
				--build-arg DOCKER_REGISTRY=${DOCKER_REGISTRY} \
				--build-arg DOCKER_CONTEXT_PATH=${DOCKER_CONTEXT_PATH} \
				--build-arg VCS_REF=$(shell git rev-parse --short HEAD) \
				--build-arg BUILD_DATE=$(shell date +%FT%T%Z) \
				-t $${DOCKER_REGISTRY:+${DOCKER_REGISTRY}/}${DOCKER_CONTEXT_PATH}/$${target}-$${type}:$(VERSION) .; \
			done \
		done \
	else \
		echo "Building image for ${TARGET}-${TYPE}..."; \
		docker build --rm \
		--build-arg BUILDKIT_INLINE_CACHE=1 \
		--build-arg VERSION=$(VERSION) \
		--build-arg TYPE=${TYPE} \
		--build-arg TARGET=${TARGET} \
		--build-arg DOCKER_REGISTRY=${DOCKER_REGISTRY} \
		--build-arg DOCKER_CONTEXT_PATH=${DOCKER_CONTEXT_PATH} \
		--build-arg VCS_REF=$(shell git rev-parse --short HEAD) \
		--build-arg BUILD_DATE=$(shell date +%FT%T%Z) \
		-t $${DOCKER_REGISTRY:+${DOCKER_REGISTRY}/}${DOCKER_CONTEXT_PATH}/${TARGET}-${TYPE}:$(VERSION) .; \
	fi

docker_clean:
	@echo "Cleaning dangling images..."
	@docker images -f "dangling=true" -q  | xargs docker rmi
	@echo "Removing microservice images..."
	@docker images -f "label=org.label-schema.vendor=sumo" -q | xargs docker rmi
	@echo "Pruneing images..."
	@docker image prune -f

docker_push:
	@echo "Piblishing images with VCS_REF=$(shell git rev-parse --short HEAD)"
	@docker images -f "label=org.label-schema.vcs-ref=$(shell git rev-parse --short HEAD)" --format {{.Repository}}:{{.Tag}} | \
	while read -r image; do \
		echo Now pushing $$image; \
		docker push $$image; \
	done;

kustomize: OVERLAY 	:= e2e
kustomize: NS 			:= default
kustomize:
	# @kustomize build deploy/overlays/${OVERLAY}/ | sed -e "s|\$$(NS)|${NS}|g" -e "s|\$$(IMAGE_VERSION)|${VERSION}|g" | kubectl apply -f -
	@kustomize build deploy/overlays/${OVERLAY}/ | sed -e "s|\$$(NS)|${NS}|g" 	-e "s|\$$(IMAGE_VERSION)|${VERSION}|g" > build/deploy.yaml

build/kustomize: check_dirty
	@kustomize build deploy/overlays/e2e/ 			 | sed -e "s|\$$(NS)|default|g" -e "s|\$$(IMAGE_VERSION)|${VERSION}|g" > build/deploy.e2e.yaml
	@kustomize build deploy/overlays/production/ | sed -e "s|\$$(NS)|default|g" -e "s|\$$(IMAGE_VERSION)|${VERSION}|g" > build/deploy.production.yaml
	@kustomize build deploy/overlays/production/ | sed -e "s|\$$(NS)|mynamespace|g" 	-e "s|\$$(IMAGE_VERSION)|${VERSION}|g" > build/deploy.production.mynamespace.yaml

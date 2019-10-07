# Make

using Makefile

> use `-n` flag for `dry-run`, `-s` or '--silent' flag to suppress echoing

## Targets

- default **VERSION** is `git tag`
- default **TYPE** is `srv`

### proto

> codegen from proto

```bash
make proto
make proto TARGET=account
make proto TARGET=account TYPE=api
make proto-account
make proto-account-api
## generate for protos in shared package
make proto TARGET=shared TYPE=.
```

### test

```bash
# unit tests
make test-account
make test-emailer
make test-account-api
make test-config-shared
make test-demo-cmd

# integration tests
make inte-account
make inte-emailer

# end-to-end tests
make e2e-account
make e2e-emailer
make e2e
```

### e2e tests on CI

> trigger e2e tests on GitHub Actions

```bash
make start_e2e GITHUB_TOKEN=123...
```

### run

```bash
make run-account
make run-emailer
make run-account-api
make run-micro-cmd ARGS="--api_address=0.0.0.0:8088 api"
make run-demo-cmd
```

### lint

```bash
# lint all
make lint
# lint account srv
make lint-account
make lint-account-srv
```

### build

```bash
# use git tag as VERSION
make build VERSION=v0.1.1
make build TARGET=account VERSION=v0.1.1
make build TARGET=account TYPE=srv VERSION=v0.1.1
make build TARGET=emailer TYPE=srv VERSION=v0.1.1
make build TARGET=account TYPE=api VERSION=v0.1.1
make build-account VERSION=v0.1.1
make build-account-api VERSION=v0.1.1
```

### release

> push tag to git

```bash
make release VERSION=v0.1.1 GITHUB_TOKEN=123...
```

### docker

```bash
make docker-account VERSION=v0.1.1
make docker-account-srv VERSION=v0.1.1
make docker TARGET=account VERSION=v0.1.1
make docker TARGET=account TYPE=srv VERSION=v0.1.1
make docker TARGET=account DOCKER_REGISTRY=us.gcr.io DOCKER_CONTEXT_PATH=<MY_PROJECT_ID>/micro-starter-kit

# short hand for TARGET and TYPE args
make docker-emailer-srv

# build all docker images for docker-compose
make docker
make docker DOCKER_REGISTRY=us.gcr.io
make docker DOCKER_REGISTRY=docker.pkg.github.com DOCKER_CONTEXT_PATH=xmlking/micro-starter-kit


# publish all microservices images
make docker_push

# remove all previous microservices images and any dangling images
make docker_clean
```

### kustomize

> generate `deploy/deploy.yaml` for given `overlay` and `namespace` using **kustomize**

```bash
make kustomize OVERLAY=production NS=default VERSION=v1.0.1
make kustomize OVERLAY=production NS=default
make kustomize OVERLAY=production
make kustomize NS=default
# default ENV=e2e,  NS=default VERSION=git tag
make kustomize
```

# Micro

> Custom build for `microhq/micro:kubernetes` to use with k8s.

> To use as a `REST Gateway` for gRPC microservices. CORS enabled.

## Run

```bash
make run-micro-cmd ARGS=api
make run-micro-cmd ARGS="--api_address=0.0.0.0:8088 api"
# with plugins
go run cmd/micro/main.go cmd/micro/plugin.go --api_address=0.0.0.0:8088  api
# without plugins (same as official micro cli)
go run cmd/micro/main.go  --api_address=0.0.0.0:8088  api
```

## Docker

> from project root directory, run following commands.

### Docker Build

```bash
# build
VERSION=0.0.9-SNAPSHOT
# DOCKER_REGISTRY=gcr.io
DOCKER_CONTEXT_PATH=xmlking
docker build --rm \
--build-arg VERSION=$VERSION \
--build-arg DOCKER_REGISTRY=${DOCKER_REGISTRY} \
--build-arg DOCKER_CONTEXT_PATH=${DOCKER_CONTEXT_PATH} \
--build-arg VCS_REF=$(shell git rev-parse --short HEAD) \
--build-arg BUILD_DATE=$(date -u +'%Y-%m-%dT%H:%M:%SZ') \
-t ${DOCKER_REGISTRY:+${DOCKER_REGISTRY}/}${DOCKER_CONTEXT_PATH}/micro:${VERSION} -f cmd/micro/Dockerfile .

IMANGE_NAME=${DOCKER_REGISTRY:+${DOCKER_REGISTRY}/}${DOCKER_CONTEXT_PATH}/micro:${VERSION}

# push
docker push $IMANGE_NAME

# check
docker inspect  $IMANGE_NAME
# remove temp images after build
docker image prune -f
# Remove all untagged images
docker rmi $(docker images | grep "^<none>" | awk "{print $3}")
```

### Docker Run

> run just for testing image...

```bash
docker run -it \
-e MICRO_API_ADDRESS=0.0.0.0:8080 \
-e MICRO_BROKER_ADDRESS=0.0.0.0:10001 \
-e MICRO_REGISTRY=mdns \
-p 8080:8080 -p 10001:10001 $IMANGE_NAME api
```

## Environment variables

```bash
MICRO_REGISTRY="kubernetes"
MICRO_SELECTOR="static"
```

```bash
CORS_ALLOWED_HEADERS="Authorization,Content-Type"
# CORS_ALLOWED_ORIGINS="*"
# important - don't  put a / at the end of the ORIGINS
CORS_ALLOWED_ORIGINS="http://localhost:4200,https://api.kashmora.com"
CORS_ALLOWED_METHODS="POST,GET"
```

### Ref

<https://micro.mu/docs/go-grpc.html>

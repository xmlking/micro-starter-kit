# Micro

> Custom build for `microhq/micro:kubernetes` to use with k8s.

> To use as a `REST Gateway` for gRPC microservices. CORS enabled.

## Run

```bash
# with plugins
go run cmd/micro/main.go cmd/micro/plugin.go --api_address=0.0.0.0:8088  api
# without plugins (same as official micro cli)
go run cmd/micro/main.go  --api_address=0.0.0.0:8088  api
```

## Docker

### Docker Build

```bash
# build
VERSION=0.0.9-SNAPSHOT
BUILD_PKG=./cmd/micro
IMANGE_NAME=xmlking/micro
docker build --rm \
--build-arg VERSION=$VERSION \
--build-arg BUILD_PKG=$BUILD_PKG \
--build-arg IMANGE_NAME=$IMANGE_NAME \
--build-arg BUILD_DATE=$(date -u +'%Y-%m-%dT%H:%M:%SZ') \
-t $IMANGE_NAME -f micro.dockerfile .

# tag
docker tag $IMANGE_NAME $IMANGE_NAME:$VERSION

# push
docker push $IMANGE_NAME:$VERSION
docker push $IMANGE_NAME:"kubernetes"


# check
docker inspect  $IMANGE_NAME:$VERSION
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

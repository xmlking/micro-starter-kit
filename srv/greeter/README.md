# Greeter Service

This is the Greeter service

### Test

```bash
micro call go.micro.srv.greeter Greeter.Hello  '{"name": "John"}'

# in k8s container
./micro call greeter Greeter.Hello  '{"name": "John"}'

curl --request POST \
--url http://localhost:8080/rpc \
--header 'accept: application/json' \
--header 'content-type: application/json' \
--data '{"service": "greeter", "method": "Greeter.Hello","request": {"name": "sumo"}}'
```

### Docker

#### Docker Build

```bash
# build
VERSION=0.0.4-SNAPSHOT
BUILD_PKG=./srv/greeter
IMANGE_NAME=xmlking/greeter-srv
docker build --rm \
--build-arg VERSION=$VERSION \
--build-arg BUILD_PKG=$BUILD_PKG \
--build-arg IMANGE_NAME=$IMANGE_NAME \
--build-arg BUILD_DATE=$(date -u +'%Y-%m-%dT%H:%M:%SZ') \
-t $IMANGE_NAME .

# tag
docker tag $IMANGE_NAME $IMANGE_NAME:$VERSION

# push
docker push $IMANGE_NAME:$VERSION
docker push $IMANGE_NAME:"latest"

# check
docker inspect  $IMANGE_NAME:$VERSION
# remove temp images after build
docker image prune -f
# Remove all untagged images
docker rmi $(docker images | grep "^<none>" | awk "{print $3}")
```

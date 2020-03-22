# envoy

### Prerequisite

```bash
wget -O ~/Downloads/protoc-gen-grpc-web https://github.com/grpc/grpc-web/releases/download/1.0.7/protoc-gen-grpc-web-1.0.7-darwin-x86_64
chmod +x ~/Downloads/protoc-gen-grpc-web
mv  ~/Downloads/protoc-gen-grpc-web /usr/local/bin/protoc-gen-grpc-web

yarn global add grpc-tools
```

### Reference

```bash
# minikube mount /Users/schintha/Developer/Work:/Work
docker-compose up envoy


docker run -it --rm --name envoy \
-p 9090:9090 -p 9901:9901  \
-v "$(pwd)/deploy/bases/envoy/envoy.yaml:/etc/envoy/envoy.yaml:ro"  \
envoyproxy/envoy:latest

docker exec -it envoy /bin/bash

# admin http://localhost:9901/

 curl 'http://localhost:9090/mkit.service.greeter/Greeter.Hello' \
 -H 'Content-Type: application/grpc-web+proto' \
 -H 'X-Grpc-Web: 1' \
 -H 'custom-header-1: value1' \
 -H 'Accept: */*' \
 -H 'Connection: keep-alive' \
 --data-binary $'\x00\x00\x00\x00\x05\n\x03abc' --compressed

curl 'http://localhost:9090/yeti.EchoService/Echo' \
-H 'Accept: application/grpc-web-text' \
-H 'Content-Type: application/grpc-web-text' \
-H 'X-Grpc-Web: 1' \
-H 'Connection: keep-alive' \
-H 'Accept-Encoding: gzip, deflate, br' \
--data-binary 'AAAAAAYKBHN1bW8=' --compressed


```

1. https://github.com/jrockway/jrock.us/blob/master/ingress/envoy.yaml

# End 2 End Tests

**Integration** tests are e2e tests that invoke `Handler` methods directly and ignore networking completely.<br/>
True **e2e** tests are Black-box tests that invoke network endpoint.

## Start Cluster

> start minimal e2e test cluster locally

```bash
# start
kubectl apply -f deploy/deploy.e2e.yaml
# stop
kubectl delete -f deploy/deploy.e2e.yaml
```

> (Or) start production like e2e test cluster locally

```bash
# start
kubectl apply -f deploy/deploy.production.yaml
# stop
kubectl delete -f  deploy/deploy.production.yaml
```

## E2E Test via REST Gateway

Use **REST Client** [tests](./test-rest-api.http) for manual testing

## E2E Test via gRPC Gateway

- set envelopment variables for CI e2e tests via `micro` proxy.
- You can also run this test against local standalone service(`go run ./srv/greeter`), without any extra settings

```bash
# e2e tests in CI envelopment with micro gRPC Gateway
MICRO_PROXY_ADDRESS="localhost:8081" \
make test-e2e
# e2e tests against local standalone services
make test-e2e
```

### Reference

- Simple gRPC benchmarking and load testing tool <https://ghz.sh>

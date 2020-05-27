# etcd Operator

Default etcd operator behavior is to only manage etcd clusters created in the same namespace.

## Installing etcd operator

### Manual

```bash
# Deploy etcd operator
kubectl apply -f deploy/bases/etcd-operator/manual/deployment.yaml

```

## Uninstalling etcd operator

```bash
# Undeploy etcd operator
kubectl delete -f deploy/bases/etcd-operator/manual/deployment.yaml
```

## Reference

- https://github.com/micro/micro/tree/master/network/config/kubernetes
- https://github.com/micro/micro/blob/master/network/config/kubernetes/gcloud.md

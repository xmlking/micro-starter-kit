# NATS

Default NATS operator behavior is to only manage NATS clusters created in the same namespace.

## Installing NATS Operator

```bash
# install NATS operator
kubectl apply -f https://github.com/nats-io/nats-operator/releases/latest/download/00-prereqs.yaml
kubectl apply -f https://github.com/nats-io/nats-operator/releases/latest/download/10-deployment.yaml
# verify NATS operator
kubectl get crd

# deploy NATS cluster
kubectl apply -f deploy/bases/nats/nats.yaml
```

## Uninstalling NATS Operator

```bash
# uninstall NATS operator
kubectl delete -f https://github.com/nats-io/nats-operator/releases/latest/download/00-prereqs.yaml
kubectl delete -f https://github.com/nats-io/nats-operator/releases/latest/download/10-deployment.yaml
# uninstall NATS cluster
kubectl delete -f deploy/bases/nats/nats.yaml
```

## Reference

- https://github.com/nats-io/nats-operator
- https://github.com/micro/micro/blob/master/network/config/kubernetes/gcloud.md

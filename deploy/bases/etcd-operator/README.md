# etcd Operator

Default etcd operator behavior is to only manage etcd clusters created in the same namespace.

## Installing etcd operator

### Manual

**Don't use Helm, recommended this manual deployment for etcd-operator**

```bash
# Deploy etcd operator
kubectl apply -f deploy/bases/etcd-operator/manual/deployment.yaml
# Undeploy etcd operator
kubectl delete -f deploy/bases/etcd-operator/manual/deployment.yaml
```

### Helm

```bash
# helm v3.0.0 or higher needed
helm repo add stable https://kubernetes-charts.storage.googleapis.com
helm repo update
```

```bash
# install etcd operator
helm install sumo --set deployments.backupOperator=false  --set deployments.restoreOperator=false stable/etcd-operator

 # Check the etcd-operator logs
export POD=$(kubectl get pods -l app=sumo-etcd-operator-etcd-operator --namespace default --output name)
kubectl logs -f $POD --namespace=default
```

## Uninstalling etcd operator

```bash
# uninstall ETCD operator
helm delete sumo
helm del --purge sumo
# uninstall ETCD cluster
kubectl delete -f deploy/bases/etcd/deployment.yaml
```

## Updating ETCD Operator

```bash
# update ETCD operator
helm upgrade sumo stable/etcd-operato
```

## Reference

- https://github.com/micro/micro/tree/master/network/config/kubernetes
- https://github.com/micro/micro/blob/master/network/config/kubernetes/gcloud.md

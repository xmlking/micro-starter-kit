# etcd

Default etcd operator behavior is to only manage etcd clusters created in the same namespace.

## Installing ETCD Operator

```bash
helm repo update
# install ETCD operator
helm install --name sumo --set deployments.backupOperator=false  --set deployments.restoreOperator=false stable/etcd-operator

 # Check the etcd-operator logs
export POD=$(kubectl get pods -l app=sumo-etcd-operator-etcd-operator --namespace default --output name)
kubectl logs $POD --namespace=default

# deploy ETCD cluster
kubectl create -f deploy/bases/etcd/deployment.yaml
# Optionally create load balancer  service (to access from laptop and test etcd is responding)
kubectl create -f deploy/bases/etcd/service.json
```

## Uninstalling ETCD Operator

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
helm upgrade sumo stable/etcd-operator

# Resize an etcd cluster
 kubectl apply -f deploy/bases/etcd/deployment.yaml
```

### ETCD

Install etcd in your local environment and query remote etcd cluster

```bash
kubectl port-forward service/etcd-cluster-client -n default 2379
ETCDCTL_API=3 etcdctl version
ETCDCTL_API=3 etcdctl -w table member list
```

## Reference

- https://github.com/micro/micro/tree/master/network/config/kubernetes
- https://github.com/micro/micro/blob/master/network/config/kubernetes/gcloud.md

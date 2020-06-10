# etcd-operator

## Setup

> install etcd operator

```bash
# install etcd-operator
kubectl create -f deploy/bases/etcd-operator/manual/deployment.yaml
# or install clusterwide etcd-operator
kubectl create -f deploy/bases/etcd-operator/manual/clusterwide-deployment.yaml
```

> create etcd cluster

```bash
# install etcd-cluster
kubectl create -f deploy/bases/etcd/deployment.yaml
```

## Verify

```bash
# list all with label name=etcd-operator
kubectl get all,configmap,secret,ingress,replicasets,crd,roles,rolebindings,serviceAccounts --show-labels \
-lapp.kubernetes.io/name=etcd-operator --namespace key-dev
# list all with label name=etcd-cluster
kubectl get all --show-labels -lapp.kubernetes.io/name=etcd-cluster --namespace key-dev
# describe etcd-cluster
kubectl describe etcdclusters.etcd.database.coreos.com etcd-cluster
# list etcd-cluster services
kubectl get services
NAME                  TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)             AGE
etcd-cluster          ClusterIP   None            <none>        2379/TCP,2380/TCP   20m
etcd-cluster-client   ClusterIP   10.106.12.199   <none>        2379/TCP            20m
# connect test for  etcd-cluster
kubectl run --rm -i --tty fun --image quay.io/coreos/etcd --restart=Never -- /bin/sh
/ # ETCDCTL_API=3 etcdctl --endpoints http://etcd-cluster-client:2379 put foo bar
OK
(ctrl-D to exit)
```

## Cleanup

```bash
# delete etcd-operator
kubectl delete -f deploy/bases/etcd-operator/manual/deployment.yaml
# delete etcd-cluster
kubectl delete -f deploy/bases/etcd/deployment.yaml
```

## Reference

- https://github.com/coreos/etcd-operator/blob/master/doc/user/spec_examples.md

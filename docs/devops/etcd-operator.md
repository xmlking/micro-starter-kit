# etcd operator

## Install operator

Reference installation [instructions](https://github.com/coreos/etcd-operator/blob/master/doc/user/install_guide.md)

We will be using forked image from [cbws/etcd-operator](https://github.com/cbws/etcd-operator/releases)

About __cbws__ fork: <https://blog.cloudbear.nl/reviving-etcd-operator/>

`docker.pkg.github.com/cbws/etcd-operator/operator:v0.10.0`

```yaml
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  name: etcd-operator
spec:
  replicas: 1
  template:
    metadata:
      labels:
        name: etcd-operator
    spec:
      containers:
      - name: etcd-operator
        image: docker.pkg.github.com/cbws/etcd-operator/operator:v0.10.0
        command:
        - etcd-operator
        # Uncomment to act for resources in all namespaces. More information in doc/user/clusterwide.md
        #- -cluster-wide
        env:
        - name: MY_POD_NAMESPACE
          valueFrom:
            fieldRef:
              fieldPath: metadata.namespace
        - name: MY_POD_NAME
          valueFrom:
            fieldRef:
              fieldPath: metadata.name
```

## Install etcd

Refer for rest of instructions [here](../../deploy/bases/etcd/README.md)
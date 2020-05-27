# etcd operator

## Install operator

Reference installation [instructions](https://github.com/coreos/etcd-operator/blob/master/doc/user/install_guide.md)

We will be using forked image from [xmlking/etcd-operator](https://github.com/xmlking/etcd-operator)

About __cbws__ fork: <https://blog.cloudbear.nl/reviving-etcd-operator/>

`docker.pkg.github.com/cbws/etcd-operator/operator:v0.10.0`

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: etcd-operator
  namespace: default
  labels:
    app.kubernetes.io/name: etcd-operator
spec:
  selector:
    matchLabels:
      app.kubernetes.io/name: etcd-operator
  replicas: 1
  template:
    metadata:
      labels:
        app.kubernetes.io/name: etcd-operator
    spec:
      serviceAccountName: etcd-operator
      securityContext:
        runAsUser: 1000
        fsGroup: 1000
      containers:
        - name: etcd-operator
          image: "xmlking/etcd-operator:v0.10.1"
          # image: 'quay.io/coreos/etcd-operator:latest'
          command:
            - etcd-operator
            # Uncomment to act for resources in all namespaces. More information in doc/user/clusterwide.md
            # - -cluster-wide
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

Refer for rest of instructions [here](../../config/base/etcd/README.md)

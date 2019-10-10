# Istio

**NOTE: If you have Teller, then recommended to install Istio via Helm as documented [here](istio-helm.md)**

> assume you already have `helm` cli installed and activated on your k8s with `helm init`

1. Enter the following commands to download Istio:

```bash
# Download and unpack Istio
cd ~/Developer/Apps/
export ISTIO_VERSION=1.3.2
curl -L https://git.io/getLatestIstio | ISTIO_VERSION=1.3.2 sh -
cd istio-${ISTIO_VERSION}
```

2. Enter the following command to install the Istio CRDs:

```bash
# Create a namespace for the istio-system components:
kubectl create namespace istio-system
# Install all the Istio Custom Resource Definitions (CRDs)
helm template install/kubernetes/helm/istio-init --name istio-init --namespace istio-system | kubectl apply -f -
# Verify that all 23 Istio CRDs
kubectl get crds | grep 'istio.io' | wc -l
```

3. Install the Istio CNI components

```bash
 helm template install/kubernetes/helm/istio-cni --name=istio-cni --namespace=kube-system | kubectl apply -f -
```

3. Install one of the following variants of the demo profile: `default`, `demo`, `demo-auth`, `cni`

> Enable CNI in Istio by setting `--set istio_cni.enabled=true` in addition to the settings for your chosen profile.<br/> For example, to configure the `cni` profile:

```bash
helm template install/kubernetes/helm/istio --name istio --namespace istio-system \
    --values install/kubernetes/helm/istio/values-istio-demo-auth.yaml \
    --set istio_cni.enabled=true | kubectl apply -f -
```

4. Verifying the installation

```bash
kubectl get svc -n istio-system
ubectl get pods -n istio-system
```

5. Clean up Istio

```bash
helm template install/kubernetes/helm/istio --name istio --namespace istio-system \
    --values install/kubernetes/helm/istio/values-istio-demo-auth.yaml \
    --set istio_cni.enabled=true | kubectl delete -f -
# delete cni
helm template install/kubernetes/helm/istio-cni --name=istio-cni --namespace=kube-system | kubectl apply -f -
# delete CRDs
kubectl delete -f install/kubernetes/helm/istio-init/files
# delete namespace
kubectl delete namespace istio-system
```

## Access

[kiali](https://istio.io/docs/tasks/telemetry/kiali/)

```bash
# open kiali
kubectl -n istio-system port-forward $(kubectl -n istio-system get pod -l app=kiali -o jsonpath='{.items[0].metadata.name}') 20001:20001
open http://localhost:20001/kiali/console
admin : admin
```

[Jaeger](https://istio.io/docs/tasks/telemetry/distributed-tracing/jaeger/)

```bash
kubectl -n istio-system port-forward $(kubectl -n istio-system get pod -l app=jaeger -o jsonpath='{.items[0].metadata.name}') 15032:16686
open  http://localhost:15032
```

## Deploy your application

```bash
# Create istio enabled namespace for the go-micro components:
kubectl create namespace micro
kubectl label namespace micro istio-injection=enabled
# verify
kubectl get namespace -L istio-injection
```

## Reference

- [ISTIO WORKSHOP](https://polarsquad.github.io/istio-workshop/install-istio/)
- https://dzone.com/articles/setup-of-a-local-kubernetes-and-istio-dev-environm-1
- https://istio.io/docs/setup/install/helm/

# Istio

> assume you already have `helm` cli installed and activated on your k8s with `helm init`

## Enter the following commands to download Istio

```bash
# Download latest Istio and unpack Istio
cd ~/Developer/Work/tools/
export ISTIO_VERSION=1.3.2
curl -L https://git.io/getLatestIstio | ISTIO_VERSION=1.3.2 sh -
cd istio-${ISTIO_VERSION}
```

## Installing **Istio** via Helm (recommended)

```bash
cd  ~/Developer/Work/tools/istio-1.3.2/
kubectl apply -f install/kubernetes/helm/helm-service-account.yaml
helm init --service-account tiller
helm install install/kubernetes/helm/istio-init --name istio-init --namespace istio-system
kubectl get crds | grep 'istio.io' | wc -l
helm install install/kubernetes/helm/istio --name istio --namespace istio-system \
    --values install/kubernetes/helm/istio/values-istio-demo.yaml
```

## Verify

```bash
helm ls -a

kubectl get svc -n istio-system
kubectl get pods -n istio-system
```

## Uninstall

```bash
helm delete --purge istio
helm delete --purge istio-init
helm delete --purge istio-cni
kubectl delete namespace istio-system

# Deleting CRDs and Istio Configuration
kubectl delete -f install/kubernetes/helm/istio-init/files
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

> switch on/off istio for `default` namespace

```bash
kubectl label namespace default istio-injection=enabled
kubectl get namespace -L istio-injection
# Disabling injection for the `default` namespace
kubectl label namespace default istio-injection-
```

> if you want to exclude a specific pod from getting istio sidecar injected, add this to `Deployment` kind

```yaml
metadata:
  annotations:
  sidecar.istio.io/inject: "false"
```

## Reference

- [ISTIO WORKSHOP](https://polarsquad.github.io/istio-workshop/install-istio/)
- https://dzone.com/articles/setup-of-a-local-kubernetes-and-istio-dev-environm-1
- https://istio.io/docs/setup/install/helm/

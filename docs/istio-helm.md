# Istio

> Installing **Istio** via Helm (recommended)

> assume you already have `helm` cli installed and activated on your k8s with `helm init`

```bash
cd  ~/Developer/Apps/istio-1.3.2/
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

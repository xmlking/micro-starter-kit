# Istio

1. Enter the following commands to download Istio:

```bash
# Download and unpack Istio
export ISTIO_VERSION=1.2.5
curl -L https://git.io/getLatestIstio | ISTIO_VERSION=1.2.5 sh -
cd istio-${ISTIO_VERSION}
```

2. Enter the following command to install the Istio CRDs first:

```bash
for i in install/kubernetes/helm/istio-init/files/crd*yaml; do kubectl apply -f $i; done
```

3. Install one of the following variants of the demo profile:

```bash
kubectl apply -f install/kubernetes/istio-demo.yaml
```

4. Verifying the installation

```bash
kubectl get svc -n istio-system
ubectl get pods -n istio-system
```

5. Clean up Istio

```bash
kubectl delete -f install/kubernetes/istio-demo.yaml
for i in install/kubernetes/helm/istio-init/files/crd*yaml; do kubectl delete -f $i; done
```

## Ref

- https://dzone.com/articles/setup-of-a-local-kubernetes-and-istio-dev-environm-1
- https://istio.io/docs/setup/kubernetes/install/kubernetes/

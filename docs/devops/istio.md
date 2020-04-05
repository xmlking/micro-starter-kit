# Istio

 
## Install

```bash
brew install istioctl
Optionally souce below file to enable auto-completion
. /Users/schintha/Developer/Apps/istio-1.5.1/tools/_istioctl
```

```bash
$ istioctl proxy-<TAB>
proxy-config proxy-status
```



### Profiles

To see list of available profiles
```bash
istioctl profile list
# Display the configuration of a profile
istioctl profile dump demo
istioctl profile dump --config-path components.pilot demo
# Show differences in profiles
istioctl profile diff default demo
# Generateing a manifest before installation
istioctl manifest generate > generated-manifest.yaml
```

### Setup

> For this installation, we use the demo [configuration profile](https://istio.io/docs/setup/additional-setup/config-profiles/)

```bash
# setup Istio into your kubernetes cluster
$ istioctl manifest apply --set profile=demo
# To enable the Grafana dashboard on top of the default profile
$ istioctl manifest apply --set addonComponents.grafana.enabled=true

Detected that your cluster does not support third party JWT authentication. Falling back to less secure first party JWT
- Applying manifest for component Base...
✔ Finished applying manifest for component Base.
- Applying manifest for component Pilot...
✔ Finished applying manifest for component Pilot.
Waiting for resources to become ready...
- Applying manifest for component EgressGateways...
- Applying manifest for component IngressGateways...
- Applying manifest for component AddonComponents...
✔ Finished applying manifest for component EgressGateways.
✔ Finished applying manifest for component IngressGateways.
✔ Finished applying manifest for component AddonComponents.

✔ Installation complete
```

### Verify

```bash
kubectl get svc -n istio-system
kubectl get pods -n istio-system
# verify generated
istioctl manifest generate <your original installation options> > $HOME/generated-manifest.yaml
istioctl verify-install -f $HOME/generated-manifest.yaml
```

### Enable

Add a namespace label to instruct Istio to automatically inject Envoy sidecar proxies when you deploy your application later:

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

## Access

[kiali](https://istio.io/docs/tasks/telemetry/kiali/)

```bash
istioctl dashboard kiali
open http://localhost:20001/kiali/console
# admin : admin
```

[Jaeger](https://istio.io/docs/tasks/telemetry/distributed-tracing/jaeger/)

```bash
istioctl dashboard Jaeger
open  http://localhost:15032
```

Get an overview of your mesh

```bash
istioctl proxy-status
```

## Uninstall

```bash
istioctl manifest generate --set profile=demo | kubectl delete -f -
```

## Reference

- https://istio.io/docs/setup/getting-started/

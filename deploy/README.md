# Deployment

deploying with **Kustomize**

Why Kustomize?

- Natively built into kubectl
- Purely declarative approach to configuration customization
- Kustomize encourages a fork/modify/rebase workflow
- Manage an arbitrary number of distinctly customized Kubernetes configurations
- Setting cross-cutting fields - e.g. namespace, labels, annotations, name-prefixes, etc

## Prerequisite

```bash
brew install kubernetes-cli
# make sure you v3.2.0 or above
brew install kustomize
# optional
brew install skaffold
brew install kubernetes-helm
```

## Workflows

A _workflow_ is the sequence of steps one takes to use and maintain a configuration.

![bespoke config workflow image](../docs/images/workflow.jpg)

## Matrix Deployment

Typical release process includes deploying multiple related components to multiple environments/profiles <br/>
_kustomize_ matrix layout helps organizing kubernetes manifest files, reducing duplication <br/>
and apply consistence labels, environment specific overlays

|      | account-µs (r,sr,k) | emailer-µs (r,sr,k) | account-api (r,sr,k) | gateway (r,sr,k) |
| ---- | :-----------------: | :-----------------: | :------------------: | :--------------: |
| Dev  |       (p,c,k)       |       (p,c,k)       |        (p,k)         |       (k)        |
| Test |      (p,c,s,k)      |      (p,c,s,k)      |      (p,c,s,k)       |      (p,k)       |
| Prod |     (r,p,c,s,k)     |      (p,c,s,k)      |      (p,c,s,k)       |     (p,i,k)      |

### Legend

| symbol | description   |
| ------ | ------------- |
| r      | Resources     |
| sr     | Service       |
| i      | Ingress       |
| c      | ConfigMap     |
| s      | Secret        |
| p      | Patches       |
| k      | Kustomization |

### Layout for Matrix Deployment

> single component layout

```
├── base
│   ├── deployment.yaml
│   ├── kustomization.yaml
│   └── service.yaml
└── overlays
    ├── dev
    │   ├── kustomization.yaml
    │   └── patch.yaml
    ├── prod
    │   ├── kustomization.yaml
    │   └── patch.yaml
    └── staging
        ├── kustomization.yaml
      └── patch.yaml
```

> multi component layout

```
deploy
├── bases <-- components
│   ├── account-api
│   │   ├── deployment.yaml
│   │   ├── kustomization.yaml
│   │   └── service.yaml
│   ├── account-srv
│   │   ├── deployment.yaml
│   │   ├── kustomization.yaml
│   │   └── service.yaml
│   ├── config
│   │   └── config.yaml
│   ├── emailer-srv
│   │   ├── deployment.yaml
│   │   ├── kustomization.yaml
│   │   └── service.yaml
│   ├── gateway
│   │   ├── deployment.yaml
│   │   ├── kustomization.yaml
│   │   └── service.yaml
│   ├── kustomization.yaml
│   └──  micro-service-account.yaml
├── kustomization.yaml
└── overlays <-- environments
    ├── dev
    │   ├── config
    │   │   └── config.yaml
    │   └── kustomization.yaml
    ├── production
    │   ├── config
    │   │   └── config.yaml
    │   ├── kustomization.yaml
    │   ├── patches
    │   │   ├── replica_count.yaml
    │   │   └── resource_limit.yaml
    │   └── resources
    │       ├── hpa.yaml
    │       └── namespace.yaml
    └── staging
        ├── config
        │   └── config.yaml
        ├── config.env
        ├── deployment.yaml
        └── kustomization.yaml
```

## Kustomize

```bash
# Kustomize command the modified manifests can be generated and printed to the terminal with: --load_restrictions none
kubectl kustomize ./deploy
# only production env
kubectl kustomize ./deploy/overlays/production

# The manifests can be applied
kubectl apply -k ./deploy
# only production env
kubectl kustomize ./deploy/overlays/production

# update image version
IMAGE_VERSION=v0.1.0-118-g21f8a30
cd deploy && kustomize edit set image xmlking/account-srv:$IMAGE_VERSION && cd ..

kustomize build someapp/overlays/staging | kubectl apply -f -
kustomize build someapp/overlays/production | kubectl apply -f -

# Fix the missing and deprecated fields in kustomization file
kustomize edit fix

kustomize build ./deploy

kubectl get -k ./deploy
kubectl describe -k ./deploy
kubectl delete -k ./deploy
```

## verify

```bash
# highlight `microhq/micro:latest`
kustomize build ./deploy | grep -C 3 microhq/micro:latest

# compare the output directly to see how consul and production differ:
diff \
  <(kustomize build ./deploy/overlays/consul) \
  <(kustomize build ./deploy/overlays/production) |\
  more

kustomize build ./deploy > release.yaml
k apply -f release.yaml
k get all -l app.kubernetes.io/managed-by=kustomize
open http://localhost:8500/ui/#/dc1/services

POD_NAME=$(kubectl get pods  -lapp.kubernetes.io/name=account-srv -o jsonpath='{.items[0].metadata.name}')
kubectl logs $POD_NAME -f
kubectl logs $POD_NAME -c srv -f
kubectl logs $POD_NAME -c health -f

kubectl exec -it $POD_NAME -- /bin/sh

k get svc

k delete -f release.yaml
```

## kustomize-sopssecret-plugin

> Installing [kustomize-sopssecret-plugin](https://github.com/goabout/kustomize-sopssecret-plugin)

```bash
# VERSION=1.0.0 PLATFORM=linux ARCH=amd64
VERSION=1.0.0 PLATFORM=darwin ARCH=amd64
curl -Lo SopsSecret https://github.com/goabout/kustomize-sopssecret-plugin/releases/download/v${VERSION}/SopsSecret_${VERSION}_${PLATFORM}_${ARCH}
chmod +x SopsSecret
mkdir -p "${XDG_CONFIG_HOME:-$HOME/.config}/kustomize/plugin/goabout.com/v1beta1/sopssecret"
mv SopsSecret "${XDG_CONFIG_HOME:-$HOME/.config}/kustomize/plugin/goabout.com/v1beta1/sopssecret"
```

### Usage

> Create some encrypted values using sops:

```bash
echo FOO=secret >secret-vars.env
sops -e -i secret-vars.env

echo secret >secret-file.txt
sops -e -i secret-file.txt
```

> Add a generator to your kustomization:

```bash
cat <<. >kustomization.yaml
generators:
  - generator.yaml
.

cat <<. >generator.yaml
apiVersion: goabout.com/v1beta1
kind: SopsSecret
metadata:
  name: my-secret
envs:
  - secret-vars.env
files:
  - secret-file.txt
.
```

> Run kustomize build with the --enable_alpha_plugins flag:

`kustomize build --enable_alpha_plugins`

## Reference

1. <https://github.com/kubernetes-sigs/kustomize/blob/master/docs/glossary.md>
2. <https://blog.jetstack.io/blog/kustomize-cert-manager/>
3. <https://kustomize.io/>
4. with sops <https://teuto.net/deploying-jupyterhub-to-kubernetes-via-kustomize-using-sops-secret-management/?lang=en>
5. <https://github.com/pwittrock-me/petclinic-config/tree/master/config>
6. <https://github.com/venilnoronha/grpc-web-istio-demo>
7. patch example, keycloak traefik <https://github.com/piotrjanik/opa-warsaw-cloud-native-conf/tree/master/manifests>

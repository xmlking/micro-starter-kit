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

kustomize edit set image busybox=alpine:3.6

kustomize build someapp/overlays/staging | kubectl apply -f -
kustomize build someapp/overlays/production | kubectl apply -f -

# Fix the missing and deprecated fields in kustomization file
kustomize edit fix

kustomize build ./deploy

kubectl get -k ./deploy
kubectl describe -k ./deploy
kubectl delete -k ./deploy
```

## Reference

1. <https://github.com/kubernetes-sigs/kustomize/blob/master/docs/glossary.md>
2. <https://blog.jetstack.io/blog/kustomize-cert-manager/>
3. <https://kustomize.io/>
4. with sops <https://teuto.net/deploying-jupyterhub-to-kubernetes-via-kustomize-using-sops-secret-management/?lang=en>
5. <https://github.com/pwittrock-me/petclinic-config/tree/master/config>
6. <https://github.com/venilnoronha/grpc-web-istio-demo>
7. patch example, keycloak traefik <https://github.com/piotrjanik/opa-warsaw-cloud-native-conf/tree/master/manifests>

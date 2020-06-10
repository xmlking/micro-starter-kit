# Deployment

deploying with **appctl** and **Kustomize**

### Structure

```
├── config
│   ├── base
│   │   ├── kustomization.yaml
│   │   └── myapp.yaml
│   └── envs
│       ├── prod
│       │   └── kustomization.yaml
│       └── staging
│           ├── kustomization.yaml
│           └── patch-replicas.yaml
└── delivery
    └── envs
        ├── prod.yaml
        └── staging.yaml
```

we can even have multiple levels of nested overlays

e.g., base, envs, <regions/zones>

#### config/base

  Configuration in the directory `config/base` applies to all environments. Additional configuration in `config/envs` can modify this configuration.

#### config/envs

  The repository contains information on two environments, defined in the directory `config/envs`: `prod` and `staging`.

  The `prod` environment refers only to the configuration in `config/base/myapp.yaml`.

  The `staging` environment has an additional customization in `config/envs/staging/patch-replicas.yaml`. This customization is referenced in `config/envs/staging/kustomization.yaml`.

#### delivery/envs

  Subdirectories in `delivery/envs` contain information on the GKE clusters that host each environment. These files are automatically generated and don't need to be modified directly.

### Install

```bash
gcloud components install appctl
```

### Setup

> One time setup

```bash
# Initialize existing repository
# make sure the `git remote -v` show `git@github.com:xmlking/micro-starter-kit.git`
cd ..
appctl init micro-starter-kit --app-config-repo=github.com/xmlking/micro-starter-kit
cd micro-starter-kit
# Create the configuration for your Kubernetes workload. i.e., add/update `config/base`, then test:
# kubectl apply -k config/base/ --dry-run=client -o yaml
kustomize build config/base
# if works, add changes to git and commit.
git add .
git commit -m "chore(deploy): bootstraping config"
git push
# 2. add new envs and connect to cluster
appctl env add development --cluster=sumo --namespace=development --review-required=false
appctl env add staging --cluster=sumo --namespace=staging --review-required=false
appctl env add production --cluster=sumo --namespace=production --review-required=true
# To see appctl changes, run `git log -p *`.
# push auto-generated configurations
git push
# Create the configuration for your enveronments. i.e., add/update `config/envs`, then test, push code.
# dry run to see what you will create
# kubectl apply -k config/envs/development  --dry-run=client -o yaml
mkdir -p {./build/kubernetes/development,./build/kubernetes/production,./build/kubernetes/staging}
kustomize build config/envs/development --output ./build/kubernetes/development --load_restrictor none 
kustomize build config/envs/production --output ./build/kubernetes/production --load_restrictor none 
kustomize build config/envs/staging --output ./build/kubernetes/staging --load_restrictor none
# tag changes
git tag v0.1.3
git push origin  v0.1.3
# prepare env PR (response with created PR in seymour-env)
appctl prepare development
appctl prepare development --from-tag v0.1.3
# run apply without merge the PR -> deny
appctl apply staging
# merge PR in seymour-env and see created dev branch
# rerun apply
# open GCP and see GKE/Applications
# to promote a release candidate from one environment to another, run the following command:
appctl prepare prod --from-env staging
# to deploy the release candidate to the target environment, run the following command:
appctl apply prod
# rollback
appctl apply development --from-tag v0.1.0
```

### Reference

1. [kustomize](https://kubectl.docs.kubernetes.io/pages/examples/kustomize.html)
1. <https://github.com/kubernetes-sigs/kustomize/blob/master/docs/glossary.md>
1. <https://blog.jetstack.io/blog/kustomize-cert-manager/>
1. <https://kustomize.io/>
1. with sops <https://teuto.net/deploying-jupyterhub-to-kubernetes-via-kustomize-using-sops-secret-management/?lang=en>
1. <https://github.com/pwittrock-me/petclinic-config/tree/master/config>
1. [TODO: gRPC-Web Istio Demo](https://github.com/venilnoronha/grpc-web-istio-demo)
1. patch example, keycloak traefik <https://github.com/piotrjanik/opa-warsaw-cloud-native-conf/tree/master/manifests>
1. [Application Delivery](https://cloud.google.com/kubernetes-engine/docs/concepts/add-on/application-delivery)
1. [Application Manager for GKE](https://cloud.google.com/blog/products/containers-kubernetes/announcing-application-manager-for-google-kubernetes-engine)
1. [Application Manager brings GitOps to GKE](https://www.youtube.com/watch?v=r5_xYtbZPfc)

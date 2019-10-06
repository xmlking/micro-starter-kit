# GitOps

> TODO CI/CD [Pipeline](https://github.com/tektoncd/pipeline/tree/master/tekton)

- Push Pipeline
- Pull Request(PR) Pipeline
- Release Pipeline

## Feature Branch Push Pipeline

- lint
- build
- tests

## PR Pipeline

TODO: lets automate _PR Pipeline_ that is triggered when PR is created for `develop` branch

- lint
- build
- tests

## Release Pipeline

- build docker images
- sign images
- push images to GCR
- generate release.yaml for k8s with Helm or Kustomize
- deploy to GKE

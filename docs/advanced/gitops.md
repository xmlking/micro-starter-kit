# GitOps

> TODO CI/CD [Pipeline](https://github.com/tektoncd/pipeline/tree/master/tekton)

- Push Pipeline
- Pull Request(PR) Pipeline
- Release Pipeline
- Deployment Pipeline

## Feature Branch Push Pipeline

- lint
- unit tests

## PR Pipeline

> triggered when PR is created for `develop` branch

- lint
- unit tests
- integration tests

## Release Pipeline

- build docker images
- sign images
- push images to GCR
- generate build/deploy.yaml for k8s with Helm or Kustomize
- generate release on GitHub

## Deployment Pipeline

- Deploy to KinD on CI
- E2E Test

- deploy to GKE

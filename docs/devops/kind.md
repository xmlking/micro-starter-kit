# Kubernetes

Ephemeral Kubernetes Clusters with KinD

## Prerequisite

```bash
# install `docker for mac`
brew cask install docker
# then go to the gui launcher and start up docker, and follow the prompts.
# uninstalling `docker for mac`
brew cask zap docker

# brew cask install minikube
```

## Install

```bash
GO111MODULE=on go get sigs.k8s.io/kind
```

## Usage

### Create a cluster

> for help `kind [command] --help`

```bash
kind create cluster # Default cluster context name is `kind`.
# with name
kind create cluster --name blog --wait 5m

# list clusters
kind get clusters

# Deleting a Cluster
kind delete cluster
```

> verify `kindest/node` container running with `docker ps`

### Interacting With Your Cluster

```bash
export KUBECONFIG="$(kind get kubeconfig-path)"
# export KUBECONFIG="$(kind get kubeconfig-path --name blog)"
kubectl cluster-info
```

### Loading

> Loading an Image Into Your Cluster

```bash
kind load docker-image my-custom-image:unique-tag
kind load docker-image ko.local/ko-demo-7a5550aba07ee9abc7c6c2992dc2c243:0f9dea87eb5c56703dc806e05d70276ca14014c9dc49ca8c8cb88507f8997a72
```

## Reference

1. <https://kind.sigs.k8s.io/>
2. <https://blog.alexellis.io/be-kind-to-yourself/>
3. <https://garethr.dev/2019/05/ephemeral-kubernetes-clusters-with-kind-and-make/>
4. [Running end-to-end tests on your Kubernetes cluster with Kind and Brigade](https://radu-matei.com/blog/kubernetes-e2e-kind-brigade/)

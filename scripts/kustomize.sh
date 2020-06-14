#!/usr/bin/env bash

env

set -x
NS=test

echo "Deploying micro-starter-kit..."

# Set the image tag to the sha hash that we just built in the previous stage.
# K8S will do a rolling deployment
(   cd config/envs/production;  \
    kustomize edit set docker.pkg.github.com/xmlking/micro-starter-kit/greeter-service:"$SHORT_SHA" && \
    kustomize edit set namespace $NS  && \
    kustomize build | \
    kubectl apply -f -  )

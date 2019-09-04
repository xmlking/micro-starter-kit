# Testing k8s consul

## consul deployment via Helm

> check for [values.yaml](https://www.consul.io/docs/platform/k8s/helm.html)

```bash
# Clone the chart repo
$ git clone https://github.com/hashicorp/consul-helm.git
$ cd consul-helm

# Checkout a tagged version
$ git checkout v0.1.0

# Run Helm
$ helm install --name consul ./
```

```bash
# create service account
kubectl create -f deploy/consul/micro-service-account.yaml
kubectl get sa
kubectl get clusterroles | grep micro-role
kubectl get ClusterRoleBinding | grep micro-role-binding
kubectl delete -f deploy/consul/micro-service-account.yaml

# account consul
kubectl create -f deploy/consul/consul.yaml
kubectl create -f deploy/consul/consul-svc.yaml

# access consul UI at http://localhost:8500/ui/#/dc1/services

POD_NAME=$(kubectl get pods  -lapp=consul -o jsonpath='{.items[0].metadata.name}')
kubectl logs $POD_NAME -f
kubectl exec -it $POD_NAME -- /bin/busybox sh

kubectl delete -f deploy/consul/consul.yaml
kubectl delete -f deploy/consul/consul-svc.yaml

# account service
kubectl create -f deploy/consul/account.yaml
kubectl create -f deploy/consul/account-svc.yaml

POD_NAME=$(kubectl get pods  -lapp=account -o jsonpath='{.items[0].metadata.name}')
kubectl logs $POD_NAME -f
kubectl exec -it $POD_NAME -- /bin/busybox sh

kubectl delete -f deploy/consul/account.yaml
kubectl delete -f deploy/consul/account-svc.yaml

# emailer service
kubectl create -f deploy/consul/emailer.yaml
kubectl create -f deploy/consul/emailer-svc.yaml

POD_NAME=$(kubectl get pods  -lapp=emailer -o jsonpath='{.items[0].metadata.name}')
kubectl logs $POD_NAME -f
kubectl exec -it $POD_NAME -- /bin/busybox sh

kubectl delete -f deploy/consul/emailer.yaml
kubectl delete -f deploy/consul/emailer-svc.yaml

# greeter service
kubectl create -f deploy/consul/greeter.yaml
kubectl create -f deploy/consul/greeter-svc.yaml

POD_NAME=$(kubectl get pods  -lapp=greeter -o jsonpath='{.items[0].metadata.name}')
kubectl logs $POD_NAME -f
kubectl exec -it $POD_NAME -- /bin/busybox sh

kubectl delete -f deploy/consul/greeter.yaml
kubectl delete -f deploy/consul/greeter-svc.yaml

# Gateway service
kubectl create -f deploy/consul/gateway.yaml
kubectl create -f deploy/consul/gateway-svc.yaml

POD_NAME=$(kubectl get pods  -lapp=gateway -o jsonpath='{.items[0].metadata.name}')
kubectl logs $POD_NAME -f
kubectl exec -it $POD_NAME -- /bin/busybox sh

kubectl delete -f deploy/consul/gateway.yaml
kubectl delete -f deploy/consul/gateway-svc.yaml
```

> Test from inside `micro:kubernetes` container

```bash
POD_NAME=$(kubectl get pods  -lapp=gateway -o jsonpath='{.items[0].metadata.name}')
kubectl exec -it $POD_NAME -- /bin/busybox sh

# list services
./micro list services
# describe `gateway` service
./micro get service go.micro.api
# describe `account` service
./micro get service  go.micro.srv.account | /bin/busybox less

# list users
./micro call go.micro.srv.account UserService.List '{}'

# call Greeter
./micro call go.micro.srv.greeter Greeter.Hello  '{"name": "John"}'

# list users
./micro call go.micro.srv.account UserService.List '{"limit": 10, "page": 1}'

# create new user  - not working
./micro call go.micro.srv.account UserService.Create \
'{"username": "sumo", "firstName": "sumo", "lastName": "demo", "email": "sumo@demo.com"}'

./micro call go.micro.srv.account ProfileService.Create  \
'{"userId": "467c715c-3f6e-40c4-9256-4c3e32d6870a", "tz" : "PST", "avatar": "sumo1.jpg", "gender": "F", "birthday": "2017-01-15T01:30:15.01Z"}'

# lets try with REST Gateway..
curl --request POST \
--url http://localhost:8080/rpc \
--header 'accept: application/json' \
--header 'content-type: application/json' \
--data '{"service": "go.micro.srv.account","method": "UserService.List","request": {}}'

curl --request POST \
--url http://localhost:8080/rpc \
--header 'accept: application/json' \
--header 'content-type: application/json' \
--data '{"service": "go.micro.srv.account", "method": "UserService.List","request": {"limit": 10, "page": 1}}'


```

## Reference

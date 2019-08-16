# Testing k8s deployment

```bash
# create service account
kubectl create -f deploy/basic/micro-service-account.yaml
kubectl get sa
kubectl get clusterroles | grep micro-role
kubectl get ClusterRoleBinding | grep micro-role-binding
kubectl delete -f deploy/basic/micro-service-account.yaml

# account service
kubectl create -f deploy/basic/account.yaml
kubectl create -f deploy/basic/account-svc.yaml

POD_NAME=$(kubectl get pods  -lapp=account -o jsonpath='{.items[0].metadata.name}')
kubectl logs $POD_NAME -f
kubectl exec -it $POD_NAME -- /bin/busybox sh

kubectl delete -f deploy/basic/account.yaml
kubectl delete -f deploy/basic/account-svc.yaml

# emailer service
kubectl create -f deploy/basic/emailer.yaml
kubectl create -f deploy/basic/emailer-svc.yaml

POD_NAME=$(kubectl get pods  -lapp=emailer -o jsonpath='{.items[0].metadata.name}')
kubectl logs $POD_NAME -f
kubectl exec -it $POD_NAME -- /bin/busybox sh

kubectl delete -f deploy/basic/emailer.yaml
kubectl delete -f deploy/basic/emailer-svc.yaml

# Gateway service
kubectl create -f deploy/basic/gateway.yaml
kubectl create -f deploy/basic/gateway-svc.yaml

POD_NAME=$(kubectl get pods  -lapp=gateway -o jsonpath='{.items[0].metadata.name}')
kubectl logs $POD_NAME -f
kubectl exec -it $POD_NAME -- /bin/busybox sh

kubectl delete -f deploy/basic/gateway.yaml
kubectl delete -f deploy/basic/gateway-svc.yaml
```

> Test from inside `micro:kubernetes` container

```bash
POD_NAME=$(kubectl get pods  -lapp=gateway -o jsonpath='{.items[0].metadata.name}')
kubectl exec -it $POD_NAME -- /bin/busybox sh

# list services
./micro --registry=kubernetes --selector=static list services
# describe `account` service
./micro  --registry=kubernetes --selector=static  get service account

# create new user
./micro --registry=kubernetes --selector=static \
call account UserService.Create \
'{"username": "sumo", "firstName": "sumo", "lastName": "demo", "email": "sumo@demo.com"}'

# list users - working
./micro --registry=kubernetes --selector=static \
call account UserService.List '{}'

# list users - not working
./micro --registry=kubernetes --selector=static \
call account UserService.List '{"limit": 10, "page": 1}'

./micro --registry=kubernetes --selector=static \
call account ProfileService.Create  \
'{"userId": "05cf2a6a-c063-11e9-9cb5-2a2ae2dbcce4", "tz" : "PST", "avatar": "sumo1.jpg", "gender": "F", "birthday": "2017-01-15T01:30:15.01Z"}'

# lets try with REST Gateway..
curl --request POST \
--url http://localhost:8080/rpc \
--header 'accept: application/json' \
--header 'content-type: application/json' \
--data '{"service": "account","method": "UserService.List","request": {}}'

curl --request POST \
--url http://localhost:8080/rpc \
--header 'accept: application/json' \
--header 'content-type: application/json' \
--data '{"service": "account","method": "UserService.List","request": {"limit": 10, "page": 1}}'

```

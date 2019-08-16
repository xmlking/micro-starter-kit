# Testing k8s deployment

```bash
# service account
kubectl create -f deploy/basic/micro-service-account.yaml
kubectl get sa
kubectl get clusterroles | grep micro-role
kubectl get ClusterRoleBinding | grep micro-role-binding
kubectl delete -f deploy/basic/micro-service-account.yaml

# account-srv
kubectl create -f deploy/basic/account.yaml
kubectl create -f deploy/basic/account-svc.yaml

POD_NAME=$(kubectl get pods  -lapp=account-srv -o jsonpath='{.items[0].metadata.name}')
kubectl logs $POD_NAME -f
kubectl exec -it $POD_NAME -- /bin/busybox sh

kubectl delete -f deploy/basic/account.yaml
kubectl delete -f deploy/basic/account-svc.yaml

# emailer-srv
kubectl create -f deploy/basic/emailer.yaml
kubectl create -f deploy/basic/emailer-svc.yaml

POD_NAME=$(kubectl get pods  -lapp=emailer-srv -o jsonpath='{.items[0].metadata.name}')
kubectl logs $POD_NAME -f
kubectl exec -it $POD_NAME -- /bin/busybox sh

kubectl delete -f deploy/basic/emailer.yaml
kubectl delete -f deploy/basic/emailer-svc.yaml

# Gateway
kubectl create -f deploy/basic/gateway.yaml
kubectl create -f deploy/basic/gateway-svc.yaml

POD_NAME=$(kubectl get pods  -lapp=gateway -o jsonpath='{.items[0].metadata.name}')
kubectl logs $POD_NAME -f
kubectl exec -it $POD_NAME -- /bin/busybox sh

kubectl delete -f deploy/basic/gateway.yaml
kubectl delete -f deploy/basic/gateway-svc.yaml
```

> Test inside `micro:kubernetes` container

```bash
POD_NAME=$(kubectl get pods  -lapp=gateway -o jsonpath='{.items[0].metadata.name}')
kubectl exec -it $POD_NAME -- /bin/busybox sh
./micro \
  --registry=kubernetes \
  --selector=static \
  --api_namespace=go.micro.srv \
  call go.micro.srv.account UserService.Create \
  '{"username": "sumo", "firstName": "sumo", "lastName": "demo", "email": "sumo@demo.com"}'
```

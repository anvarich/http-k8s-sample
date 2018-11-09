# Containerized microservice challenge

1 
`http_srv`
Implement a piece of software exposing a JSON document:
```json
{
"id": "1",
"message": "Hello world"
}
```
build docker container with running `./build.sh` script

---
2 
`http_proxy`
second application, that utilizes the first and displays reversed message text
build docker container with running `./build.sh` script

---
3 You can deploy it to  to Minikube Kubernetes cluster or to particular GCE/Azure AKS k8s cluster 
with `./deploy.sh` script.

---
4 Yu can check how it works with `./check.sh` script

---

####Code review example questions

Q: Explain, how to ensure running multiple instances of the application 

A: You can check it usually with `kubectl get deployment hello -n demo`, current row should show you 
how many copies of pods now running

Q: Explain, how you would organize regular application upgrades

A: For version update you should up version of docker container in `./build.sh` script and set new version of image 
in `deploy.yaml` manifests or you can update it in runtime set image version by running
`kubectl set image deployments/hello hello=tshafeev/hello:0.2` for example
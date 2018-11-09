#!/bin/sh
set -e
external_ip_hello=$(kubectl get svc hello --template="{{range .status.loadBalancer.ingress}}{{.ip}}{{end}}")
external_ip_frontend=$(kubectl get svc frontend --template="{{range .status.loadBalancer.ingress}}{{.ip}}{{end}}")
echo "hello world Check from $external_ip_hello port 8080\n"
curl http://$external_ip_hello:8080
echo "--------------\n"
echo "hello world reverse string Check from $external_ip_frontend port 80\n"
curl http://$external_ip_frontend:80
echo "---------------\n"
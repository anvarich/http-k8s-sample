#!/bin/sh
kubectl create ns demo
kubectl apply -f http_srv/deploy.yaml -n demo
kubectl apply -f http_proxy/deploy.yaml -n demo
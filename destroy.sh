#!/bin/sh
kubectl delete -f http_srv/deploy.yaml -n demo
kubectl delete -f http_proxy/deploy.yaml -n demo
kubectl delete ns demo

#!/bin/sh
docker build -t shafeev/http_proxy:0.1 .
docker push shafeev/http_proxy:0.1
kubectl apply -f deploy.yaml
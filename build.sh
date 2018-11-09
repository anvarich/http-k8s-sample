#!/bin/sh
cd http_srv/
docker build -t shafeev/http_srv:0.1 .
docker push shafeev/http_srv:0.1
cd -
cd http_proxy/
docker build -t shafeev/http_proxy:0.1 .
docker push shafeev/http_proxy:0.1
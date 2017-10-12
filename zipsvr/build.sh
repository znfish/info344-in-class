#!/usr/bin/env bash
set -e
GOOS=linux go build #linux executable
docker build -t znfish/zipsvr .
docker push znfish/zipsvr
go clean
#docker run -d -p 80:80 znfish/zipsvr

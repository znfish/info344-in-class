#!/usr/bin/env bash
set -e #if any of the commands get error, stop, don't try to keep going
GOOS=linux go build
docker build -t znfish/testserver .
docker push znfish/testserver

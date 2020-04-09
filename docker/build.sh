#!/bin/bash

set -eo pipefail

GOOS=linux GOARCH=amd64 go build -o main ./..
docker build -t 200313983415.dkr.ecr.us-east-2.amazonaws.com/liquidata/vault:v1.3.2-aws-patch .
docker push 200313983415.dkr.ecr.us-east-2.amazonaws.com/liquidata/vault:v1.3.2-aws-patch

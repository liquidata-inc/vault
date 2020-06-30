#!/bin/bash

set -eo pipefail

(cd ..; make bootstrap; make static-dist; XC_OSARCH=linux/amd64 make bin)
cp ../bin/vault .
docker build -t 200313983415.dkr.ecr.us-east-2.amazonaws.com/liquidata/vault:v1.4.2-aws-patch .
docker push 200313983415.dkr.ecr.us-east-2.amazonaws.com/liquidata/vault:v1.4.2-aws-patch

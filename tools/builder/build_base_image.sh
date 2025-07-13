#!/bin/bash

set -e

PROXY=$1
image="ghcr.io/lukaproject/linda/linda-buildbase-image:latest"
dockerfilePath="tools/dockerimages/buildbase/Dockerfile.buildbase"

echo "build image $image"
echo "dockerfile $dockerfilePath"

docker buildx build -f $dockerfilePath -t $image . --build-arg PROXY=$PROXY

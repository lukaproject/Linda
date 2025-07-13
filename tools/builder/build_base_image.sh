#!/bin/bash

set -e

PUSH=$1
PROXY=$2
image="ghcr.io/lukaproject/linda/linda-buildbase-image:latest"
dockerfilePath="tools/dockerimages/buildbase/Dockerfile.buildbase"

echo "build image $image, push $PUSH"
echo "dockerfile $dockerfilePath"

docker buildx build -f $dockerfilePath -t $image . --build-arg PROXY=$PROXY

if [[ $PUSH = true ]]; then
    echo "docker push $image"
    docker push $image
fi
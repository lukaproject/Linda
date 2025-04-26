#!/bin/bash

if [[ ( $@ == "--help") ||  $@ == "-h" ]]
then 
	echo "Usage: "
    echo "$0 imagesList buildTool[buildx|docker]"
    echo "  imagesList: Images you want to build."
    echo "  buildTool: docker or buildx, you can choose what is your prefer one."
	exit 0
fi 
subImagesPath=$1
buildTool=$2

cat $subImagesPath | while read line
do
    IFS=',' read -ra arr <<< $line
    dockerfilePath=${arr[0]}
    imageName=${arr[1]}
    echo "dockerfile: $dockerfilePath image name: $imageName"
    docker rmi $imageName:latest
    set -e
    $buildTool build -f $dockerfilePath -t $imageName .
    echo "build success: $imageName"
    set +e
done

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
    echo "dockerfile: ${arr[0]} image tag: ${arr[1]}"
    docker rmi ${arr[1]}:latest
    $buildTool build -f ${arr[0]} -t ${arr[1]} .
done

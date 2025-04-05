# /bin/bash
# pls run this command in root dir

NUMBER_OF_AGENTS=$1

docker compose -f tools/dockerimages/agent/docker-compose.yml up --scale agents-cluster=$NUMBER_OF_AGENTS -d
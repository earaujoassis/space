#!/usr/bin/env bash

set -e

function create_containers() {
    docker-compose up -d --build
}

function update_containers() {
    docker-compose build space
    docker-compose up --no-deps -d space
    # docker rmi $(docker images -a --filter=dangling=true -q)
    # docker rm $(docker ps --filter=status=exited --filter=status=created -q)
    docker system prune -a -f
}

if [ "$1" == "-u" ]; then
    echo "`date`: space/deploy: Updating containers"
    update_containers
    echo "`date`: space/deploy: Containers updated"
    exit 0
else
    echo "`date`: space/deploy: Building and creating containers"
    create_containers
    echo "`date`: space/deploy: Containers created"
    exit 0
fi

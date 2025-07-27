#!/bin/bash

# sudo systemctl start docker
docker stop firebird

docker-compose up -d

docker ps --format "{{.Names}}"

#!/bin/bash
docker run --rm -v $PWD/data:/data --name mongodb -d -p 27017-27019:27017-27019 mongo:latest
sleep 1
#進入monog shell
docker exec -ti mongodb bash
docker stop mongodb
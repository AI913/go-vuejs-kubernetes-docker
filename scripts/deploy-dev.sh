#!/bin/bash

if  [ -z ${SHA+x}]; then SHA=$(git rev-parse HEAD); fi

docker build -t dennischiu/echo1:lasetest-dev -f ./echoservers/echo1/Dockerfile.dev ./echoservers/echo1
docker build -t dennischiu/echo1:$SHA -f ./echoservers/echo1/Dockerfile.dev ./echoservers/echo1

docker push dennischiu/echo1:lasetest-dev
docker push dennischiu/echo1:$SHA
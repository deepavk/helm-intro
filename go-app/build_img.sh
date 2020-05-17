#!/bin/bash

#docker login -u dvk16

docker build -t go-app:latest .

docker tag go-app:latest dvk16/go-app:latest

docker push dvk16/go-app:latest

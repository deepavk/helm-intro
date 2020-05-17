#!/bin/bash

docker login -u dvk16

docker build -t example-app:latest app

docker tag example-app:latest dvk16/helm-intro:latest

docker push dvk16/helm-intro:latest

# docker run -d -p 5000:5000 example-app

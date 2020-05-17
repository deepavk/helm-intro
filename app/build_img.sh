#!/bin/bash

docker login -u dvk16

docker build -t flask-app:latest .

docker tag flask-app:latest dvk16/helm-intro:latest

docker push dvk16/helm-intro:latest

# docker run -d -p 5000:5000 flask-app

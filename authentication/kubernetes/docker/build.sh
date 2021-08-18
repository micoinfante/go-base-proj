#!/bin/bash

cp ../../authentication/authsvc .
cp ../../api/apisvc .

docker build -t authentication:v1 .
docker inspect authentication:v1
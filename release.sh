#!/bin/bash

set -e

gf build main.go -a amd64 -s linux -p ./temp
gf docker main.go -p -t epicmo/soruxgpt-saas-oauth:latest
now=$(date +"%Y%m%d%H%M%S")
# 以当前时间为版本号
docker tag epicmo/soruxgpt-saas-oauth:latest epicmo/soruxgpt-saas-oauth:$now
docker push epicmo/soruxgpt-saas-oauth:$now
echo "release success" $now

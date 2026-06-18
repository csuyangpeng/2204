# Script used to start lite5gc-dev containter
#!/usr/bin/env bash
set -e

cd "$(dirname "${BASH_SOURCE[0]}")"

# REPOSITORY=10.18.1.2:5000
ALIREPOSITORY=registry.cn-hangzhou.aliyuncs.com/10_18_1_2_5000
TAGS=latest

# registry.cn-hangzhou.aliyuncs.com/10_18_1_2_5000/kamailio:latest
docker build --progress=plain -t ${ALIREPOSITORY}/kamailio:${TAGS} . -f Dockerfile
docker push ${ALIREPOSITORY}/kamailio:${TAGS}
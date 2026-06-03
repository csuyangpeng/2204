#!/usr/bin/env bash
set -e

cd "$(dirname "${BASH_SOURCE[0]}")"

REPOSITORY=10.18.1.2:5000
# ALIREPOSITORY=registry.cn-hangzhou.aliyuncs.com/10_18_1_2_5000
TAGS=dev-x86_64-22.04 #-`date +%Y%m%d`
docker build --progress=plain -t sems:${TAGS} . -f Dockerfile
docker tag sems:${TAGS} ${REPOSITORY}/sems:${TAGS}
docker push ${REPOSITORY}/sems:${TAGS}
# docker tag sems:${TAGS} ${ALIREPOSITORY}/sems:${TAGS}
# docker push ${ALIREPOSITORY}/sems:${TAGS}
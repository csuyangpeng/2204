# Script used to start epc-dev containter
#!/usr/bin/env bash
set -e

cd "$(dirname "${BASH_SOURCE[0]}")"

REPOSITORY=10.18.1.2:5000
# TAGS=dev-x86_64-18.04-`date +%Y%m%d_%H%M`-${VER}
ALIREPOSITORY=registry.cn-hangzhou.aliyuncs.com/10_18_1_2_5000
TAGS=dev-x86_64-22.04-dev-2.5.13
docker build --progress=plain  -t epc:${TAGS} . -f Dockerfile
docker tag epc:${TAGS} ${REPOSITORY}/epc:${TAGS}
docker push ${REPOSITORY}/epc:${TAGS}
docker tag epc:${TAGS} ${ALIREPOSITORY}/epc:${TAGS}
docker push ${ALIREPOSITORY}/epc:${TAGS}

# Script used to start lite5gc-dev containter
#!/usr/bin/env bash
set -e

cd "$(dirname "${BASH_SOURCE[0]}")"

REPOSITORY=10.18.1.2:5000
ALIREPOSITORY=registry.cn-hangzhou.aliyuncs.com/10_18_1_2_5000
TAGS=dev-x86_64-18.04-`date +%Y%m%d_%H%M`
docker build --progress=plain -t lite5gc:${TAGS} . -f Dockerfile
docker tag lite5gc:${TAGS} ${REPOSITORY}/lite5gc:${TAGS}
docker push ${REPOSITORY}/lite5gc:${TAGS}
docker tag lite5gc:${TAGS} ${ALIREPOSITORY}/lite5gc:${TAGS}
docker push ${ALIREPOSITORY}/lite5gc:${TAGS}

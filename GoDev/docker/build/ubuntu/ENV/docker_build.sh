# Script used to start lite5gc-dev containter
#!/usr/bin/env bash
set -e

cd "$(dirname "${BASH_SOURCE[0]}")"

REPOSITORY=10.18.1.2:5000
ALIREPOSITORY=registry.cn-hangzhou.aliyuncs.com/10_18_1_2_5000
TAGS=dev-22.04-20250219

# registry.cn-hangzhou.aliyuncs.com/10_18_1_2_5000/lite5gc:dev-22.04-20250219
docker build --progress=plain -t ${ALIREPOSITORY}/lite5gc:${TAGS} . -f Dockerfile
docker push ${ALIREPOSITORY}/lite5gc:${TAGS}
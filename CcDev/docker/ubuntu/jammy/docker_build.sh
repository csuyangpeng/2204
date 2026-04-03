# Script used to start epc4g-dev containter
#!/usr/bin/env bash
set -e

cd "$(dirname "${BASH_SOURCE[0]}")"

REPOSITORY=10.18.1.2:5000
ALIREPOSITORY=registry.cn-hangzhou.aliyuncs.com/10_18_1_2_5000
# TAGS=dev-x86_64-22.04 #-`date +%Y%m%d`
TAGS=dev-ARM_64-22.04-`date +%Y%m%d`
docker build --build-arg https_proxy=http://10.18.11.52:7890 --build-arg http_proxy=http://10.18.11.52:7890 \
  --network host --progress=plain -t epc4g:${TAGS} . -f Dockerfile
docker tag epc4g:${TAGS} ${REPOSITORY}/epc4g:${TAGS}
docker push ${REPOSITORY}/epc4g:${TAGS}
docker tag epc4g:${TAGS} ${ALIREPOSITORY}/epc4g:${TAGS}
docker push ${ALIREPOSITORY}/epc4g:${TAGS}

TAGS=dev-22.04-`date +%Y%m%d`
# docker buildx build --progress=plain --platform linux/amd64,linux/arm64 -f Dockerfile -t ${ALIREPOSITORY}/epc4g:${TAGS} --push .
docker buildx build --build-arg https_proxy=http://10.18.11.52:7890 --build-arg http_proxy=http://10.18.11.52:7890 \
  --progress=plain --platform linux/amd64,linux/arm64 -f Dockerfile -t ${ALIREPOSITORY}/epc4g:${TAGS} --push .

# Script used to start cc-dev containter
#!/usr/bin/env bash
set -e

cd "$(dirname "${BASH_SOURCE[0]}")"

# REPOSITORY=10.18.1.2:5000
ALIREPOSITORY=registry.cn-hangzhou.aliyuncs.com/10_18_1_2_5000
TAGS=dev-22.04 #-`date +%Y%m%d`
# docker build --progress=plain -t cc:${TAGS} . -f Dockerfile
# docker tag epc4g:${TAGS} ${REPOSITORY}/epc4g:${TAGS}
# docker push ${REPOSITORY}/epc4g:${TAGS}
docker buildx build --progress=plain --platform linux/amd64,linux/arm64 \
    -f Dockerfile -t ${ALIREPOSITORY}/cc:${TAGS} --push .
# docker tag cc:${TAGS} ${ALIREPOSITORY}/cc:${TAGS}
# docker push ${ALIREPOSITORY}/cc:${TAGS}
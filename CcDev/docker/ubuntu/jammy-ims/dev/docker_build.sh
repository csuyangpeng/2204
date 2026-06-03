#!/usr/bin/env bash
set -e

cd "$(dirname "${BASH_SOURCE[0]}")"

REPOSITORY=10.18.1.2:5000
ALIREPOSITORY=registry.cn-hangzhou.aliyuncs.com/10_18_1_2_5000
TAGS=latest
# docker build --network=host --progress=plain -t kamailio:${TAGS} . -f Dockerfile
# docker tag kamailio:${TAGS} ${REPOSITORY}/kamailio:${TAGS}
# docker push ${REPOSITORY}/kamailio:${TAGS}
# docker tag kamailio:${TAGS} ${ALIREPOSITORY}/kamailio:${TAGS}
# docker push ${ALIREPOSITORY}/kamailio:${TAGS}

docker buildx build --progress=plain --platform linux/amd64,linux/arm64 \
    -f Dockerfile -t ${ALIREPOSITORY}/kamailio:${TAGS} --push .

# # Build docker images for kamailio IMS components
# cd ../ims_base
# docker build --no-cache --force-rm -t docker_kamailio .

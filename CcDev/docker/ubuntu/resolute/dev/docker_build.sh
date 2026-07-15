#!/usr/bin/env bash
set -e

cd "$(dirname "${BASH_SOURCE[0]}")"

REPOSITORY=10.18.1.2:5000
ALIREPOSITORY=registry.cn-hangzhou.aliyuncs.com/10_18_1_2_5000
# TAGS=latest
# TAGS=dev-26.04-`date +%Y%m%d`
TAGS=dev-26.04-20260624

GITPASS_FILE="${GITPASS_FILE:-$(dirname "$0")/.gitpass}"
if [ ! -s "${GITPASS_FILE}" ]; then
    # fallback: 与 base 共享 .gitpass
    GITPASS_FILE="$(dirname "$0")/../base/.gitpass"
fi
if [ ! -s "${GITPASS_FILE}" ]; then
    echo "ERROR: ${GITPASS_FILE} not found or empty." >&2
    echo "Copy .gitpass.example to .gitpass and paste your GitLab token." >&2
    exit 1
fi

# docker build --network=host --progress=plain -t kamailio:${TAGS} . -f Dockerfile
# docker tag kamailio:${TAGS} ${REPOSITORY}/kamailio:${TAGS}
# docker push ${REPOSITORY}/kamailio:${TAGS}
# docker tag kamailio:${TAGS} ${ALIREPOSITORY}/kamailio:${TAGS}
# docker push ${ALIREPOSITORY}/kamailio:${TAGS}

docker buildx build --progress=plain --platform linux/amd64,linux/arm64 \
    --secret "id=gitpass,src=${GITPASS_FILE}" \
    -f Dockerfile -t ${ALIREPOSITORY}/kamailio:${TAGS} --push .

# # Build docker images for kamailio IMS components
# cd ../ims_base
# docker build --no-cache --force-rm -t docker_kamailio .

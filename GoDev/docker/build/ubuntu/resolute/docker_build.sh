# Script used to start lite5gc-dev containter
#!/usr/bin/env bash
set -e

cd "$(dirname "${BASH_SOURCE[0]}")"

REPOSITORY=10.18.1.2:5000
ALIREPOSITORY=registry.cn-hangzhou.aliyuncs.com/10_18_1_2_5000
# TAGS=dev-26.04-`date +%Y%m%d`
TAGS=dev-26.04-20260622

GITPASS_FILE="${GITPASS_FILE:-$(dirname "$0")/.gitpass}"
if [ ! -s "${GITPASS_FILE}" ]; then
    echo "ERROR: ${GITPASS_FILE} not found or empty." >&2
    echo "Copy .gitpass.example to .gitpass and paste your GitLab token." >&2
    exit 1
fi

export DOCKER_BUILDKIT=1
# 多 stage + BuildKit secret：token 不写入 Dockerfile / 镜像层
# docker build --progress=plain \
#     --secret "id=gitpass,src=${GITPASS_FILE}" \
#     -t lite5gc:${TAGS} . -f Dockerfile
# # docker tag lite5gc:${TAGS} ${REPOSITORY}/lite5gc:${TAGS}
# # docker push ${REPOSITORY}/lite5gc:${TAGS}
# docker tag lite5gc:${TAGS} ${ALIREPOSITORY}/lite5gc:${TAGS}
# docker push ${ALIREPOSITORY}/lite5gc:${TAGS}

docker buildx build --progress=plain --secret id=gitpass,src=${GITPASS_FILE} \
    --platform linux/amd64,linux/arm64 -f Dockerfile -t ${ALIREPOSITORY}/lite5gc:${TAGS} --push .

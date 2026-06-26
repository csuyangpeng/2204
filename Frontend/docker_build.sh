#!/usr/bin/env bash
# 构建 omc-frontend 开发环境镜像（对应 docker/script/start.sh 所用镜像）
set -e

cd "$(dirname "${BASH_SOURCE[0]}")"

REPOSITORY=10.18.1.2:5000
ALIREPOSITORY=registry.cn-hangzhou.aliyuncs.com/10_18_1_2_5000
IMAGE=ubuntu-node-builder
TAGS=26.04-node20-pnpm8-fast

docker build --progress=plain --network host \
    -f Dockerfile -t ${IMAGE}:${TAGS} .

docker tag ${IMAGE}:${TAGS} ${REPOSITORY}/${IMAGE}:${TAGS}
docker push ${REPOSITORY}/${IMAGE}:${TAGS}
docker tag ${IMAGE}:${TAGS} ${ALIREPOSITORY}/${IMAGE}:${TAGS}
docker push ${ALIREPOSITORY}/${IMAGE}:${TAGS}

echo "Image: ${REPOSITORY}/${IMAGE}:${TAGS}"
echo "Image: ${ALIREPOSITORY}/${IMAGE}:${TAGS}"

# Script used to start lite5gc-dev containter
#!/usr/bin/env bash
set -e

cd "$(dirname "${BASH_SOURCE[0]}")"

REPOSITORY=10.18.1.2:5000
ALIREPOSITORY=registry.cn-hangzhou.aliyuncs.com/10_18_1_2_5000
# TAGS=dev-22.04-`date +%Y%m%d`
# TAGS=dev-26.04-`date +%Y%m%d`
TAGS=dev-26.04-20260622
# 多 stage 构建：builder 编译层不进入最终镜像 history
docker build --progress=plain -t kims:${TAGS} . -f Dockerfile
docker tag kims:${TAGS} ${ALIREPOSITORY}/kims:${TAGS}
docker push ${ALIREPOSITORY}/kims:${TAGS}

# docker buildx build --progress=plain --platform linux/amd64,linux/arm64 -f Dockerfile -t ${ALIREPOSITORY}/lite5gc:${TAGS} --push .
# docker buildx build --progress=plain --platform linux/amd64 -f Dockerfile -t ${ALIREPOSITORY}/lite5gc:${TAGS} --push .
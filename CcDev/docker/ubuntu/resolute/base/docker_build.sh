#!/usr/bin/env bash
set -e

cd "$(dirname "${BASH_SOURCE[0]}")"

# REPOSITORY=10.18.1.2:5000
ALIREPOSITORY=registry.cn-hangzhou.aliyuncs.com/10_18_1_2_5000
TAGS=dev-26.04 #-`date +%Y%m%d`
CACHE_DIR="${HOME}/.cache/docker-buildx/cc-resolute-base"
mkdir -p "${CACHE_DIR}"

CACHE_FROM=()
if [ -f "${CACHE_DIR}/index.json" ]; then
  CACHE_FROM=(--cache-from "type=local,src=${CACHE_DIR}")
fi

GITPASS_FILE="${GITPASS_FILE:-$(dirname "$0")/.gitpass}"
SECRET_FLAGS=()
if [ -s "${GITPASS_FILE}" ]; then
    SECRET_FLAGS=(--secret "id=gitpass,src=${GITPASS_FILE}")
else
    echo "WARNING: ${GITPASS_FILE} not found. Private GitLab clones will fail." >&2
    echo "Copy .gitpass.example to .gitpass and paste your GitLab token if needed." >&2
fi

# 快速本地验证（仅 amd64、不 push、复用 build cache）:
#   FAST=1 bash docker_build.sh
# 改 install_cpprestsdk.sh 等脚本时，preinstall/wfrest 等已成功的层会直接命中缓存。
if [ "${FAST:-}" = "1" ]; then
  docker buildx build --progress=plain --platform linux/amd64 \
    "${CACHE_FROM[@]}" "${SECRET_FLAGS[@]}" \
    --cache-to "type=local,dest=${CACHE_DIR},mode=max" \
    -f Dockerfile -t "cc:${TAGS}" --load .
  exit 0
fi

# 正式多架构构建并推送
docker buildx build --progress=plain --platform linux/amd64,linux/arm64 \
  "${CACHE_FROM[@]}" "${SECRET_FLAGS[@]}" \
  --cache-to "type=local,dest=${CACHE_DIR},mode=max" \
  -f Dockerfile -t "${ALIREPOSITORY}/cc:${TAGS}" --push .

#!/usr/bin/env bash
# 从 BuildKit secret 读取 GitLab token，避免写入 Dockerfile/镜像层
export GITUSER="${GITUSER:-peng.yang}"

if [ -r /run/secrets/gitpass ]; then
    GITPASS="$(grep -v '^[[:space:]]*#' /run/secrets/gitpass | grep -v '^[[:space:]]*$' | tail -1 | tr -d '[:space:]')"
    export GITPASS
elif [ -n "${GITPASS:-}" ]; then
    export GITPASS
else
    echo "ERROR: Git token missing. Build with:" >&2
    echo "  docker build --secret id=gitpass,src=.gitpass ..." >&2
    exit 1
fi

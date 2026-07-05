#!/usr/bin/env bash

set -e

cd "$(dirname "${BASH_SOURCE[0]}")"

# Ubuntu 26.04 基础镜像自带 ubuntu(1000)，adduser sder 会得到 1001。
# 宿主 bind-mount 源码多为 uid=1000，需与 lite5gc:dev-22.04 对齐。
if getent passwd ubuntu >/dev/null; then
    userdel -r ubuntu 2>/dev/null || userdel ubuntu
fi

adduser --disabled-password --gecos '' --uid 1000 sder
usermod -aG sudo sder
echo '%sudo ALL=(ALL) NOPASSWD:ALL' >> /etc/sudoers

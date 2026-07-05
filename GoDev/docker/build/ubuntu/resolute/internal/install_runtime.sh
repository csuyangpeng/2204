#!/usr/bin/env bash
# 最终 stage：仅安装运行时/开发 apt 包与用户，不暴露 builder 编译过程
set -e

cd "$(dirname "${BASH_SOURCE[0]}")"

bash preinstall.sh
bash install_sder.sh

mkdir -p "${GOPATH}/bin" "${GOPATH}/src"
chmod 777 "${GOPATH}" -R
echo "Asia/Shanghai" > /etc/timezone

#!/usr/bin/env bash
# builder stage：编译安装 golang / nDPI / tcpdump，不进入最终镜像 history
set -e

cd "$(dirname "${BASH_SOURCE[0]}")"

bash preinstall.sh
bash install_golang.sh
bash install_3rd.sh
bash install_tcpdump.sh

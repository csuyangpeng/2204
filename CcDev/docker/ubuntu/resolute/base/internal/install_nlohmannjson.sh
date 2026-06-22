#!/usr/bin/env bash

set -e

cd "$(dirname "${BASH_SOURCE[0]}")"
source "$(dirname "${BASH_SOURCE[0]}")/git_env.sh"

# 全历史极大，浅克隆可显著降低经代理/QEMU 时的断流概率
git clone --depth 1 --single-branch https://github.com/nlohmann/json.git
pushd json
mkdir build && cd build && cmake .. && make -j8 && make install
popd
rm -rf json

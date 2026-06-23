#!/usr/bin/env bash

set -euo pipefail

cd "$(dirname "${BASH_SOURCE[0]}")"
source "$(dirname "${BASH_SOURCE[0]}")/git_env.sh"
git config --global http.sslVerify false

git clone https://github.com/etcd-cpp-apiv3/etcd-cpp-apiv3.git
pushd etcd-cpp-apiv3

# CMake 4.x 不再兼容 cmake_minimum_required < 3.5
sed -i 's/cmake_minimum_required(VERSION [0-9.]*)/cmake_minimum_required(VERSION 3.5)/' CMakeLists.txt

mkdir -p build && cd build
cmake .. -DCMAKE_POLICY_VERSION_MINIMUM=3.5
make -j"$(nproc)"
make install

popd
rm -rf etcd-cpp-apiv3

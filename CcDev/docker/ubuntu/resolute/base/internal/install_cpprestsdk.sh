#!/usr/bin/env bash

set -euo pipefail

cd "$(dirname "${BASH_SOURCE[0]}")"
source "$(dirname "${BASH_SOURCE[0]}")/git_env.sh"
git config --global http.sslVerify false

PATCH_DIR="$(dirname "${BASH_SOURCE[0]}")/patches"

# v2.10.19 + vcpkg asio 补丁（Boost 1.89+ io_service 已移除）
git clone --depth 1 --branch v2.10.19 https://github.com/microsoft/cpprestsdk.git
pushd cpprestsdk
git apply "$PATCH_DIR/cpprestsdk-fix-asio-error.patch"

# Boost 1.89+ 移除 boost_system 编译库 stub（1.69 起已是 header-only）
# https://github.com/microsoft/cpprestsdk/pull/1838
BOOST_CMAKE=Release/cmake/cpprest_find_boost.cmake
sed -i \
  -e 's/ random system thread/ random thread/g' \
  -e 's/REQUIRED COMPONENTS system date_time/REQUIRED COMPONENTS date_time/g' \
  -e '/Boost::system$/d' \
  "$BOOST_CMAKE"

# GCC -Werror=format-truncation: buffer[9] 装不下 64 位 %8zX
# https://github.com/microsoft/cpprestsdk/issues/1796
HELPERS_CPP=Release/src/http/common/http_helpers.cpp
sed -i 's/buffer\[9\]/buffer[17]/' "$HELPERS_CPP"
grep -q 'buffer\[17\]' "$HELPERS_CPP" || { echo "patch $HELPERS_CPP failed"; exit 1; }

mkdir -p build && cd build
cmake .. \
  -DCPPREST_EXCLUDE_WEBSOCKETS=ON \
  -DBUILD_TESTS=OFF \
  -DCMAKE_CXX_STANDARD=14 \
  -DCMAKE_CXX_STANDARD_REQUIRED=ON \
  -Wno-dev \
  -DCMAKE_CXX_FLAGS="-Wno-error=format-truncation -Wno-format-truncation"
make -j"$(nproc)"
make install

popd
rm -rf cpprestsdk

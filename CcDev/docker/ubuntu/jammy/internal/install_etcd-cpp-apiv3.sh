#!/usr/bin/env bash

set -e

cd "$(dirname "${BASH_SOURCE[0]}")"
git config --global http.sslVerify false

git clone https://github.com/etcd-cpp-apiv3/etcd-cpp-apiv3.git
pushd etcd-cpp-apiv3 && mkdir build && cd build && cmake .. && make -j8 && make install && popd
rm -rf etcd-cpp-apiv3

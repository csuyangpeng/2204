#!/usr/bin/env bash

set -e

cd "$(dirname "${BASH_SOURCE[0]}")"

version=4.8
wget https://github.com/ntop/nDPI/archive/refs/tags/${version}.tar.gz
tar -xf ${version}.tar.gz
pushd nDPI-${version}
./autogen.sh
unset MAKEFLAGS
if [ "$(uname -m)" = "aarch64" ]; then
    export CFLAGS="-O0 -g"
    export CXXFLAGS="-O0 -g"
    jobs=1
else
    jobs="$(nproc)"
fi
make -j"${jobs}"
make install
popd

rm ${version}.tar.gz nDPI-${version} -rf

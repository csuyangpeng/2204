#!/usr/bin/env bash

set -e

cd "$(dirname "${BASH_SOURCE[0]}")"
source "$(dirname "${BASH_SOURCE[0]}")/git_env.sh"

git clone http://${GITUSER}:${GITPASS}@10.18.1.2:9999/plt/tcpdump.git
pushd tcpdump
mkdir build && cd build
unset MAKEFLAGS
if [ "$(uname -m)" = "aarch64" ]; then
    cmake -DCMAKE_POLICY_VERSION_MINIMUM=3.5 -DCMAKE_C_FLAGS="-O0 -g" ..
    make -j1
else
    cmake -DCMAKE_POLICY_VERSION_MINIMUM=3.5 ..
    make -j"$(nproc)"
fi
make install
cd -
popd

rm tcpdump -rf

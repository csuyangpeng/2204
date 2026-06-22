#!/usr/bin/env bash

set -e

cd "$(dirname "${BASH_SOURCE[0]}")"

git clone http://${GITUSER}:${GITPASS}@10.18.1.2:9999/plt/tcpdump.git
pushd tcpdump
mkdir build && cd build
cmake -DCMAKE_POLICY_VERSION_MINIMUM=3.5 ..
make -j2
make install
cd -
popd

rm tcpdump -rf

#!/usr/bin/env bash

set -e

cd "$(dirname "${BASH_SOURCE[0]}")"

pushd cJSON && mkdir -p build && cd build && cmake .. && make -j && make install && cd .. && popd
pushd hiredis && mkdir -p build && cd build && cmake .. && make -j && make install && cd .. && popd

cd epc
bash epc.sh buildrelease && bash epc.sh install
cd && rm /tmp/internal/ -rf

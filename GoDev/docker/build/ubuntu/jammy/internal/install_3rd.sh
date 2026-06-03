#!/usr/bin/env bash

set -e

cd "$(dirname "${BASH_SOURCE[0]}")"

wget https://github.com/ntop/nDPI/archive/refs/tags/4.8.tar.gz
tar -xf 4.8.tar.gz
pushd nDPI-4.8 && ./autogen.sh && make && make install && popd

rm 4.8.tar.gz nDPI-4.8 -rf

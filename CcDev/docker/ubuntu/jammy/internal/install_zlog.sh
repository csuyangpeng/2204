#!/usr/bin/env bash

set -e

cd "$(dirname "${BASH_SOURCE[0]}")"

git clone https://github.com/HardySimpson/zlog.git
pushd zlog && make -j8 && make install && popd 
rm -rf zlog

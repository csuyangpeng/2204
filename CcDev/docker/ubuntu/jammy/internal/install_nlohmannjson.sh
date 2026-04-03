#!/usr/bin/env bash

set -e

cd "$(dirname "${BASH_SOURCE[0]}")"

git clone https://github.com/nlohmann/json.git
pushd json
mkdir build && cd build && cmake .. && make -j8 && make install
popd
rm -rf json

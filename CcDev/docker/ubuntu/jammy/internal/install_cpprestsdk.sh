#!/usr/bin/env bash

set -e

cd "$(dirname "${BASH_SOURCE[0]}")"

git clone https://github.com/microsoft/cpprestsdk.git
pushd cpprestsdk
mkdir build && cd build && cmake .. -DCPPREST_EXCLUDE_WEBSOCKETS=ON && make -j8 && make install
popd
rm -rf cpprestsdk

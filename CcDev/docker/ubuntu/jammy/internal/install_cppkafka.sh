#!/usr/bin/env bash

set -e

cd "$(dirname "${BASH_SOURCE[0]}")"

git clone https://github.com/confluentinc/librdkafka.git
pushd librdkafka
./configure && make -j8 && make install
popd
rm -rf librdkafka

git clone https://github.com/mfontanini/cppkafka.git
pushd cppkafka
mkdir build && cd build && cmake .. && make -j8 && make install
popd
rm -rf cppkafka

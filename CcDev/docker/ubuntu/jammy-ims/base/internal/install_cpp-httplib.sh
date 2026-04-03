#!/usr/bin/env bash

set -e

cd "$(dirname "${BASH_SOURCE[0]}")"
source "$(dirname "${BASH_SOURCE[0]}")/git_env.sh"

git clone https://github.com/yhirose/cpp-httplib.git
pushd cpp-httplib
sed -i 's/<milliseconds>/<std::chrono::milliseconds>/g' httplib.h
mkdir build && cd build && cmake .. && make -j8 && make install
popd
rm -rf cpp-httplib

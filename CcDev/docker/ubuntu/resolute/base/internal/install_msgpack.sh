#!/usr/bin/env bash

set -e

cd "$(dirname "${BASH_SOURCE[0]}")"

tar -xf msgpack-c-master.tar
pushd msgpack-c-master
# sed -i '1i add_compile_options(-Wno-error=class-memaccess)' CMakeLists.txt
sed -i 's/-Wno-mismatched-tags//g' CMakeLists.txt
mkdir build
pushd build
cmake -DCMAKE_CXX_FLAGS=-Wno-implicit-fallthrough -DMSGPACK_BUILD_EXAMPLES=OFF -DMSGPACK_BUILD_TESTS=OFF .. && make -j8 install
ln -sf /usr/local/include/msgpack.hpp /usr/local/include/msgpack/msgpack.hpp
popd
popd
rm -rf msgpack-c-master

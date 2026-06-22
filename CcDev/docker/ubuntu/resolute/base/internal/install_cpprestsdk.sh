#!/usr/bin/env bash

set -euo pipefail

cd "$(dirname "${BASH_SOURCE[0]}")"
git config --global http.sslVerify false

git clone https://github.com/microsoft/cpprestsdk.git
pushd cpprestsdk

# Boost 1.89+ 移除 boost_system 编译库 stub（1.69 起已是 header-only）
# https://github.com/microsoft/cpprestsdk/pull/1838
BOOST_CMAKE=Release/cmake/cpprest_find_boost.cmake
sed -i \
  -e 's/ random system thread/ random thread/g' \
  -e 's/REQUIRED COMPONENTS system date_time/REQUIRED COMPONENTS date_time/g' \
  -e '/Boost::system$/d' \
  "$BOOST_CMAKE"

ASIO_CPP=Release/src/http/client/http_client_asio.cpp
if [ -f "$ASIO_CPP" ] && ! grep -q 'boost/asio/deadline_timer.hpp' "$ASIO_CPP"; then
  sed -i '1s/^/#include <boost\/asio\/deadline_timer.hpp>\n/' "$ASIO_CPP"
fi

mkdir build && cd build
cmake .. -DCPPREST_EXCLUDE_WEBSOCKETS=ON -Wno-dev
make -j8
make install

popd
rm -rf cpprestsdk

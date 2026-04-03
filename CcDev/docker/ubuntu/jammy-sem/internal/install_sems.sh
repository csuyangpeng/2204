#!/usr/bin/env bash

set -e

cd "$(dirname "${BASH_SOURCE[0]}")"

git clone https://github.com/sems-server/sems.git
pushd sems
# git checkout tags/1.6.0
sed -i 's/SEMS_VERSION/"1.6.0"/g' core/plug-in/stats/StatsUDPServer.cpp
sed -i 's/event_pthreads/event_pthreads dl/g' core/CMakeLists.txt
sed -i '321 a\#ifdef PROPAGATE_COREDUMP_SETTINGS' core/sems.cpp
sed -i '337 a\#endif' core/sems.cpp

mkdir build && cd build && cmake .. && make -j8 && make install && cd ..
popd

# cp sems.conf /usr/local/etc/sems/sems.conf

rm sems -rf

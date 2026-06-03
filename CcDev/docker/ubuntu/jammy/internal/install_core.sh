#!/usr/bin/env bash

set -e

cd "$(dirname "${BASH_SOURCE[0]}")"

git clone http://${GITUSER}:${GITPASS}@10.18.1.2:9999/develop/epc4g.git
# install epc4g lib
pushd epc4g
pushd hss/src/omc/src/ && make -j8 && make install && popd
pushd hss/src/omc/ && cp -a localexp/* /usr/local/include/ && popd
mv /usr/local/include/libs/* /usr/local/lib/
popd

rm epc4g -rf

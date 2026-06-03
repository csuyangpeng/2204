#!/usr/bin/env bash

set -e

cd "$(dirname "${BASH_SOURCE[0]}")"

cp ipsecstart /usr/local/bin/
tar -zxf ipsec.tar.gz && chmod +x ipsec -R

pushd ipsec
sed -i '/#include <bcon.h>/d' ./src/libcharon/attributes/mem_pool.c
sed -i '/#include <bson.h>/s/.*/#include <bson\/bson.h>/' ./src/libcharon/attributes/mem_pool.c
autoreconf -ivf
./configure --prefix=/usr/local/ipsec/install/ --sysconfdir=/usr/local/ipsec/install/etc
make -j
make install

popd && rm ipsec.tar.gz ipsec -rf

#!/usr/bin/env bash

set -e

cd "$(dirname "${BASH_SOURCE[0]}")"

cp ipsecstart /usr/local/bin/
tar -zxvf ipsec.tar.gz && chmod +x ipsec -R

pushd ipsec
./configure --prefix=/usr/local/ipsec/install/ --sysconfdir=/usr/local/ipsec/install/etc
make -j
make install

popd && rm ipsec.tar.gz ipsec -rf

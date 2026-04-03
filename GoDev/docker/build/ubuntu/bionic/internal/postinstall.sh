#!/usr/bin/env bash

set -e

cd "$(dirname "${BASH_SOURCE[0]}")"

echo "/usr/lib64" > /etc/ld.so.conf.d/usrlib64.conf
# echo "Asia/Shanghai" > /etc/timezone

ldconfig
mkdir -p ${GOPATH}
chmod 777 ${GOPATH} -R

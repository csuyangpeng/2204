#!/usr/bin/env bash

set -e

cd "$(dirname "${BASH_SOURCE[0]}")"

git clone http://${GITUSER}:${GITPASS}@10.18.1.2:9999/plt/traceserver.git
pushd traceserver
git checkout tcp
make && mv trace ${GOPATH}/bin
popd

rm traceserver /root/.cache ${GOPATH}/pkg -rf

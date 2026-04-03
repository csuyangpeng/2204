#!/usr/bin/env bash

set -e

cd "$(dirname "${BASH_SOURCE[0]}")"

git clone http://${GITUSER}:${GITPASS}@10.18.1.2:9999/plt/omc-backend.git
pushd omc-backend
git checkout master
go mod tidy && go build backend.go && mv backend ${GOPATH}/bin/

popd

rm omc-backend /root/.cache ${GOPATH}/pkg -rf

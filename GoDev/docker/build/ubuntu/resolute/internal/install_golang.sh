#!/usr/bin/env bash

set -e

cd "$(dirname "${BASH_SOURCE[0]}")"

mkdir -p ${GOPATH} && mkdir -p ${GOPATH}/bin && mkdir -p ${GOPATH}/src

version="1.26.4"
echo "https://go.dev/dl/go${version}.linux-${TARGETARCH}.tar.gz"
cd /opt && curl -L -s https://go.dev/dl/go${version}.linux-${TARGETARCH}.tar.gz | tar zx

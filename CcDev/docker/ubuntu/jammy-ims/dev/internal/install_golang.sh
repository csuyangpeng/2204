#!/usr/bin/env bash

set -e

export http_proxy="${http_proxy:-http://10.18.11.52:7890}"
export https_proxy="${https_proxy:-http://10.18.11.52:7890}"
git config --global http.version HTTP/1.1
git config --global http.postBuffer 524288000
git config --global http.lowSpeedLimit 0
git config --global http.lowSpeedTime 999999

cd "$(dirname "${BASH_SOURCE[0]}")"

mkdir -p ${GOPATH} && mkdir -p ${GOPATH}/bin && mkdir -p ${GOPATH}/src

VER=1.22.4
echo "https://go.dev/dl/go${VER}.linux-${TARGETARCH}.tar.gz"

cd /opt && curl -L -s https://go.dev/dl/go${VER}.linux-${TARGETARCH}.tar.gz | tar zx

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

VER="1.26.4"
if [ -n "${TARGETARCH:-}" ]; then
  go_arch="${TARGETARCH}"
else
  case "$(uname -m)" in
    x86_64) go_arch=amd64 ;;
    aarch64|arm64) go_arch=arm64 ;;
    armv7l) go_arch=armv6l ;;
    *)
      echo "unsupported architecture: $(uname -m)" >&2
      exit 1
      ;;
  esac
fi

go_url="https://go.dev/dl/go${VER}.linux-${go_arch}.tar.gz"
echo "${go_url}"
cd /opt && curl -fL "${go_url}" | tar zx

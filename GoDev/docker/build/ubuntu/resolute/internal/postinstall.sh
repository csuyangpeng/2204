#!/usr/bin/env bash

set -e

cd "$(dirname "${BASH_SOURCE[0]}")"

#go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
# echo "/usr/lib64" > /etc/ld.so.conf.d/usrlib64.conf
echo "Asia/Shanghai" > /etc/timezone
#cli completion bash > /etc/bash_completion.d/cli
# ldconfig
chmod 777 ${GOPATH} -R

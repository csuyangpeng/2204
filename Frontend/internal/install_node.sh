#!/usr/bin/env bash

set -e

cd "$(dirname "${BASH_SOURCE[0]}")"

node_version=20.18.0
case "$(uname -m)" in
    x86_64) node_arch=linux-x64 ;;
    aarch64) node_arch=linux-arm64 ;;
    *)
        echo "unsupported architecture: $(uname -m)" >&2
        exit 1
    ;;
esac

tarball="node-v${node_version}-${node_arch}.tar.xz"
curl -fsSL "https://npmmirror.com/mirrors/node/v${node_version}/${tarball}" -o /tmp/node.tar.xz
tar -xJf /tmp/node.tar.xz -C /usr/local --strip-components=1
rm -f /tmp/node.tar.xz

npm config set registry https://registry.npmmirror.com
npm install -g pnpm@8.15.0
pnpm config set registry https://registry.npmmirror.com

git config --global --add safe.directory '*'

node --version
npm --version
pnpm --version

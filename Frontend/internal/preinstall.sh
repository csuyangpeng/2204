#!/usr/bin/env bash

set -e

cd "$(dirname "${BASH_SOURCE[0]}")"

unset http_proxy https_proxy HTTP_PROXY HTTPS_PROXY
rm -f /etc/apt/sources.list.d/ubuntu.sources
cp /tmp/internal/sources.list /etc/apt/sources.list
echo 'Acquire::Retries "5";' > /etc/apt/apt.conf.d/80-retries
apt-get -q update && apt-get -q -y install \
    curl \
    wget \
    git \
    ca-certificates \
    xz-utils \
    vim \
    sudo \
    bash-completion

apt-get clean && rm -rf /var/lib/apt/lists/*

#!/usr/bin/env bash

set -e

cd "$(dirname "${BASH_SOURCE[0]}")"

apt-get -q update && apt-get -q --no-install-recommends -y install \
    wget && \
    rm -rf /var/lib/apt/lists/*

apt-get clean

cd /opt && wget --no-check-certificate https://dl.google.com/go/go1.15.6.linux-amd64.tar.gz && tar xzf go1.15.6.linux-amd64.tar.gz

rm go1.15.6.linux-amd64.tar.gz

apt remove -y wget

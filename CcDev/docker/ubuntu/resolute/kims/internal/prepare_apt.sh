#!/usr/bin/env bash
set -e

unset http_proxy https_proxy HTTP_PROXY HTTPS_PROXY
rm -f /etc/apt/sources.list.d/ubuntu.sources
cp /tmp/internal/sources.list /etc/apt/sources.list
echo 'Acquire::Retries "5";' > /etc/apt/apt.conf.d/80-retries
apt-get -q update

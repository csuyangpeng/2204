#!/usr/bin/env bash

set -e

cd "$(dirname "${BASH_SOURCE[0]}")"

adduser --disabled-password --gecos '' sder
usermod -aG sudo sder
echo '%sudo ALL=(ALL) NOPASSWD:ALL' >> /etc/sudoers

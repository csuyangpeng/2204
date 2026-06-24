#!/usr/bin/env bash
# Ubuntu 26.04 已移除 libpcre3，rtpengine 等仍依赖 libpcre3-dev
set -euo pipefail

ARCH="$(dpkg --print-architecture)"
PCRE_VER="8.39-13ubuntu0.22.04.1"

case "$ARCH" in
  amd64|i386)
    BASE="http://archive.ubuntu.com/ubuntu/pool/main/p/pcre3"
    ;;
  arm64|armhf|ppc64el|s390x|riscv64)
    BASE="http://ports.ubuntu.com/ubuntu-ports/pool/main/p/pcre3"
    ;;
  *)
    echo "unsupported arch: $ARCH" >&2
    exit 1
    ;;
esac

pkgs=(
  "libpcre3_${PCRE_VER}_${ARCH}.deb"
  "libpcre16-3_${PCRE_VER}_${ARCH}.deb"
  "libpcre32-3_${PCRE_VER}_${ARCH}.deb"
  "libpcrecpp0v5_${PCRE_VER}_${ARCH}.deb"
  "libpcre3-dev_${PCRE_VER}_${ARCH}.deb"
)

tmpdir="$(mktemp -d)"
trap 'rm -rf "$tmpdir"' EXIT
cd "$tmpdir"

for pkg in "${pkgs[@]}"; do
  wget -O "$pkg" "${BASE}/${pkg}"
done

dpkg -i ./*.deb

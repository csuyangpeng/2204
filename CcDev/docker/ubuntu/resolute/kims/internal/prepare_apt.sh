#!/usr/bin/env bash
set -e

unset http_proxy https_proxy HTTP_PROXY HTTPS_PROXY
rm -f /etc/apt/sources.list.d/ubuntu.sources

# arm64 需要 ubuntu-ports，amd64 用主 archive
ARCH="${TARGETARCH:-$(dpkg --print-architecture)}"
case "$ARCH" in
    arm64|aarch64|armhf|armv7l|riscv64|ppc64el|s390x)
        MIRROR="http://mirrors.aliyun.com/ubuntu-ports"
        ;;
    *)
        MIRROR="http://mirrors.aliyun.com/ubuntu"
        ;;
esac

cat > /etc/apt/sources.list <<EOF
deb ${MIRROR} resolute main restricted universe multiverse
deb-src ${MIRROR} resolute main restricted universe multiverse
deb ${MIRROR} resolute-security main restricted universe multiverse
deb-src ${MIRROR} resolute-security main restricted universe multiverse
deb ${MIRROR} resolute-updates main restricted universe multiverse
deb-src ${MIRROR} resolute-updates main restricted universe multiverse
deb ${MIRROR} resolute-proposed main restricted universe multiverse
deb-src ${MIRROR} resolute-proposed main restricted universe multiverse
deb ${MIRROR} resolute-backports main restricted universe multiverse
deb-src ${MIRROR} resolute-backports main restricted universe multiverse
EOF

echo 'Acquire::Retries "5";' > /etc/apt/apt.conf.d/80-retries
apt-get -q update

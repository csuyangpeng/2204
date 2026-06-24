#!/usr/bin/env bash
# 为 rtpengine 9.4.1 提供与 Ubuntu 22.04 一致的 FFmpeg 4.4.2 库
set -euo pipefail
export http_proxy="${http_proxy:-http://10.18.11.52:7890}"
export https_proxy="${https_proxy:-http://10.18.11.52:7890}"
FFMPEG_PREFIX="${FFMPEG_PREFIX:-/usr/local/ffmpeg-4.4.2}"
FFMPEG_VERSION=4.4.2

tmpdir="$(mktemp -d)"
trap 'rm -rf "$tmpdir"' EXIT
cd "$tmpdir"

wget "https://ffmpeg.org/releases/ffmpeg-${FFMPEG_VERSION}.tar.bz2"
tar -xjf "ffmpeg-${FFMPEG_VERSION}.tar.bz2"
cd "ffmpeg-${FFMPEG_VERSION}"

# 26.04 的 binutils/GCC 与 4.4.2 内联汇编不兼容（mathops.h shr）；禁用全部 asm
./configure \
  --prefix="${FFMPEG_PREFIX}" \
  --enable-shared \
  --disable-static \
  --enable-gpl \
  --enable-version3 \
  --enable-nonfree \
  --disable-doc \
  --disable-ffmpeg \
  --disable-ffplay \
  --disable-ffprobe \
  --enable-avfilter \
  --enable-avformat \
  --enable-avcodec \
  --enable-swresample \
  --enable-swscale \
  --enable-postproc \
  --enable-pthreads \
  --disable-asm

make -j"$(nproc)"
make install

echo "${FFMPEG_PREFIX}/lib" > /etc/ld.so.conf.d/ffmpeg-4.4.2.conf
ldconfig

# 验证 pkg-config 指向 4.4.x（libavcodec 58）
export PKG_CONFIG_PATH="${FFMPEG_PREFIX}/lib/pkgconfig"
av_ver="$(pkg-config --modversion libavcodec)"
case "$av_ver" in 58.*|4.4.*) ;; *)
  echo "unexpected libavcodec version: $av_ver" >&2
  exit 1
esac
echo "FFmpeg ${FFMPEG_VERSION} installed at ${FFMPEG_PREFIX} (libavcodec ${av_ver})"

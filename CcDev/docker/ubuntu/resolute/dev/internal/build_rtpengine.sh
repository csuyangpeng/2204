#!/usr/bin/env bash
set -euo pipefail

cd "$(dirname "${BASH_SOURCE[0]}")/.."

PATCH_DIR="$(dirname "${BASH_SOURCE[0]}")/patches"

git clone "http://${GITUSER}:${GITPASS}@10.18.1.2:9999/develop/ims.git"
cd ims && git checkout cc_yp && cd rtpengine

# FFmpeg 7 (Ubuntu 26.04): AVCodecParameters.channels 已移除，需 ch_layout API
if [ ! -f lib/fix_frame_channel_layout-01.h ]; then
  cp "$PATCH_DIR/fix_frame_channel_layout-01.h" lib/fix_frame_channel_layout-01.h
fi
{
  echo '/******** GENERATED FILE ********/'
  cat lib/fix_frame_channel_layout-01.h
} > lib/fix_frame_channel_layout.compat

find daemon recording-daemon lib -name '*.c' 2>/dev/null | while read -r f; do
  sed -i \
    -e 's/\([[:alnum:]_>]*\)CODECPAR->channels/GET_CHANNELS(\1CODECPAR)/g' \
    -e 's/\([[:alnum:]_>]*\)codecpar->channels/GET_CHANNELS(\1codecpar)/g' \
    "$f"
done

# dpkg-buildflags 在 daemon/ 等子目录编译时只找本目录 debian/changelog
mkdir -p debian
if [ ! -f debian/changelog ]; then
  cat > debian/changelog <<'EOF'
rtpengine (0.0.0) unstable; urgency=medium

  * docker build

 -- docker <docker@local>  Mon, 01 Jan 2024 00:00:00 +0000
EOF
fi
for sub in */; do
  [ "$sub" = "debian/" ] && continue
  mkdir -p "${sub}debian"
  ln -sfn "../debian/changelog" "${sub}debian/changelog"
done

export DEB_BUILD_PROFILES="${DEB_BUILD_PROFILES:-pkg.ngcp-rtpengine.nobcg729}"

dpkg-checkbuilddeps

build_log=/tmp/rtpengine-build.log
set -o pipefail
if ! dpkg-buildpackage -b -uc -us 2>&1 | tee "$build_log" | grep -v --line-buffered 'dpkg-buildflags: warning: debian/changelog not found'; then
  echo "========== rtpengine build failed; last errors ==========" >&2
  grep -E 'error:|undefined reference|fatal error:|Error [0-9]|recipe for target' "$build_log" | tail -40 >&2 || tail -80 "$build_log" >&2
  exit 1
fi

cd ..
dpkg -i ./*.deb
ldconfig
cd ..
rm -rf ims

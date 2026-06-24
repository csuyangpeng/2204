#!/usr/bin/env bash
set -euo pipefail

cd "$(dirname "${BASH_SOURCE[0]}")/.."

FFMPEG_PREFIX="${FFMPEG_PREFIX:-/usr/local/ffmpeg-4.4.2}"
export PKG_CONFIG_PATH="${FFMPEG_PREFIX}/lib/pkgconfig${PKG_CONFIG_PATH:+:${PKG_CONFIG_PATH}}"
export LD_LIBRARY_PATH="${FFMPEG_PREFIX}/lib${LD_LIBRARY_PATH:+:${LD_LIBRARY_PATH}}"
export CPPFLAGS="-I${FFMPEG_PREFIX}/include ${CPPFLAGS:-}"

git clone "http://${GITUSER}:${GITPASS}@10.18.1.2:9999/develop/ims.git"
cd ims && git checkout cc_yp && cd rtpengine

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

for pkg in libavcodec-dev libavfilter-dev libavformat-dev libavutil-dev libswresample-dev; do
  sed -i "/${pkg}/d" debian/control
done
for pc in libavcodec libavfilter libavformat libavutil libswresample; do
  pkg-config --exists "$pc" || {
    echo "missing pkg-config module: $pc (PKG_CONFIG_PATH=${PKG_CONFIG_PATH})" >&2
    exit 1
  }
done

dpkg-checkbuilddeps

build_jobs="${RTPENGINE_BUILD_JOBS:-20}"
export DEB_BUILD_OPTIONS="parallel=${build_jobs}"
# 关键：optimize=-lto 才能去掉 -flto=auto -ffat-lto-objects（仅 -fno-lto 不够，fat-lto 仍极慢）
export DEB_BUILD_MAINT_OPTIONS="optimize=-lto"
export DEB_CFLAGS_MAINT_APPEND="-O1 -Wno-deprecated-declarations ${DEB_CFLAGS_MAINT_APPEND:-}"
export DEB_CXXFLAGS_MAINT_APPEND="-O1 -Wno-deprecated-declarations ${DEB_CXXFLAGS_MAINT_APPEND:-}"

# 顶层 Makefile 串行 daemon→recording-daemon→iptables，改为子目录并行
# cc_yp 分支带 -j20，上游无 -j；用正则兼容两种格式
python3 - Makefile <<'PY'
import re
import sys
from pathlib import Path

path = Path(sys.argv[1])
text = path.read_text()

make_j = r"(?:-j\d+ )?"
pat_all = re.compile(
    rf"all:\n\t\$\(MAKE\) {make_j}-C daemon\n"
    rf"ifeq \(\$\(with_transcoding\),yes\)\n"
    rf"\t\$\(MAKE\) {make_j}-C recording-daemon\n"
    rf"endif\n"
    rf"\t\$\(MAKE\) {make_j}-C iptables-extension",
    re.MULTILINE,
)

new_all = """SUBDIRS = daemon iptables-extension
ifeq ($(with_transcoding),yes)
SUBDIRS += recording-daemon
endif

CLEAN_DIRS = $(SUBDIRS) kernel-module t

.PHONY: all distclean clean coverity $(SUBDIRS) $(CLEAN_DIRS:%=clean-%)

all: $(SUBDIRS)

$(SUBDIRS):
\t+$(MAKE) -C $@"""

pat_clean = re.compile(
    r"distclean clean:\n"
    r"\t\$\(MAKE\) -C daemon clean\n"
    r"\t\$\(MAKE\) -C recording-daemon clean\n"
    r"\t\$\(MAKE\) -C iptables-extension clean\n"
    r"\t\$\(MAKE\) -C kernel-module clean\n"
    r"\t\$\(MAKE\) -C t clean",
    re.MULTILINE,
)

new_clean = """distclean clean: $(CLEAN_DIRS:%=clean-%)

clean-%:
\t+$(MAKE) -C $* clean"""

m = pat_all.search(text)
if not m:
    sys.exit("Makefile parallel patch: all: target not found")
text = pat_all.sub(new_all, text, count=1)

if not pat_clean.search(text):
    sys.exit("Makefile parallel patch: distclean clean target not found")
text = pat_clean.sub(new_clean, text, count=1)
text = text.replace(".PHONY:\tall distclean clean coverity\n\n", "", 1)
path.write_text(text)
print("Makefile: parallel subdir builds enabled")
PY

# iptables-extension：GCC 15 下 xtables _init() 与 crti.o 冲突，需隔离 flags + 分步链接
python3 - iptables-extension/Makefile <<'PY'
import re
import sys
from pathlib import Path

path = Path(sys.argv[1])
text = path.read_text()

pat = re.compile(
    r"^CC\?=gcc\n"
    r"CFLAGS\s*\?= -O2 -Wall -Wstrict-prototypes\n"
    r"CFLAGS\s*\+= -shared -fPIC\n",
    re.MULTILINE,
)
repl = (
    "CC?=gcc\n"
    "# xtables .so 隔离 dpkg-buildflags；GCC 15 需 -nostartfiles 避免 _init 冲突\n"
    "override PICFLAGS := -O2 -Wall -Wstrict-prototypes -fPIC -fcf-protection=none\n"
)
if not pat.search(text):
    sys.exit("iptables-extension/Makefile patch: CFLAGS block not found")
text = pat.sub(repl, text, count=1)
text = re.sub(r"^(\s*)CFLAGS(\s+)\+=", r"\1override PICFLAGS\2+=", text, flags=re.MULTILINE)
text = re.sub(
    r"^override PICFLAGS(\s+)\+= \$\(shell pkg-config --cflags --libs xtables\)",
    "XTABLES_LIBS := $(shell pkg-config --libs xtables 2>/dev/null || echo -lxtables)\n"
    "override PICFLAGS += $(shell pkg-config --cflags xtables 2>/dev/null)",
    text,
    flags=re.MULTILINE,
)
text = text.replace(
    "\t$(CC) $(CFLAGS) -o libxt_RTPENGINE.so libxt_RTPENGINE.c\n",
    "\t$(CC) $(PICFLAGS) -c -o libxt_RTPENGINE.o libxt_RTPENGINE.c\n"
    "\t$(CC) -shared -nostartfiles -fPIC -o libxt_RTPENGINE.so libxt_RTPENGINE.o $(XTABLES_LIBS)\n",
)
text = text.replace(
    "\t$(CC) $(CFLAGS) -D__ipt -o libipt_RTPENGINE.so libxt_RTPENGINE.c\n",
    "\t$(CC) $(PICFLAGS) -D__ipt -c -o libipt_RTPENGINE.o libxt_RTPENGINE.c\n"
    "\t$(CC) -shared -nostartfiles -fPIC -o libipt_RTPENGINE.so libipt_RTPENGINE.o $(XTABLES_LIBS)\n",
)
text = text.replace(
    "\t$(CC) $(CFLAGS) -D__ip6t -o libip6t_RTPENGINE.so libxt_RTPENGINE.c\n",
    "\t$(CC) $(PICFLAGS) -D__ip6t -c -o libip6t_RTPENGINE.o libxt_RTPENGINE.c\n"
    "\t$(CC) -shared -nostartfiles -fPIC -o libip6t_RTPENGINE.so libip6t_RTPENGINE.o $(XTABLES_LIBS)\n",
)
text = text.replace(
    "\trm -f libxt_RTPENGINE.so libipt_RTPENGINE.so libip6t_RTPENGINE.so\n",
    "\trm -f libxt_RTPENGINE.so libxt_RTPENGINE.o libipt_RTPENGINE.so libipt_RTPENGINE.o "
    "libip6t_RTPENGINE.so libip6t_RTPENGINE.o\n",
)
path.write_text(text)
print("iptables-extension/Makefile: PIC compile + -nostartfiles shared link")
PY

export RTPENGINE_VERSION="$(
  dpkg-parsechangelog -l debian/changelog 2>/dev/null | awk '/^Version: / {print $2; exit}'
)"
export RTPENGINE_VERSION="${RTPENGINE_VERSION:-9.4.1.6+0~mr9.4.1.6}"

eval "$(dpkg-buildflags --export=sh)"
echo "dpkg-buildflags CFLAGS: ${CFLAGS:-}"
case "${CFLAGS:-}" in
  *flto*|*ffat-lto-objects*)
    echo "WARNING: LTO flags still present in CFLAGS, stripping" >&2
    for var in CFLAGS CXXFLAGS LDFLAGS; do
      val="${!var:-}"
      val="${val//-flto=auto/}"
      val="${val//-flto /}"
      val="${val//-ffat-lto-objects/}"
      export "$var=$val"
    done
    echo "stripped CFLAGS: ${CFLAGS:-}" >&2
    ;;
esac
export CFLAGS="-I${FFMPEG_PREFIX}/include ${CFLAGS:-}"
export LDFLAGS="-L${FFMPEG_PREFIX}/lib ${LDFLAGS:-}"

# 自编译 FFmpeg、jammy libpcre、/usr/local 库无 Debian shlibs 元数据
sed -i "s|\tdh_shlibdeps -l/usr/local/lib|\tdh_shlibdeps -l/usr/local/lib -l${FFMPEG_PREFIX}/lib -- --ignore-missing-info|" debian/rules
grep -F 'ignore-missing-info' debian/rules

echo "rtpengine build: ${build_jobs} jobs, LTO off, -O1, subdir parallel"

build_log=/tmp/rtpengine-build.log
set -o pipefail
if ! dpkg-buildpackage --jobs-force="${build_jobs}" -b -uc -us 2>&1 \
  | tee "$build_log" \
  | grep -v --line-buffered \
    -e 'dpkg-buildflags: warning: debian/changelog not found' \
    -e 'is deprecated' \
    -e 'note: declared here'; then
  echo "========== rtpengine build failed; last errors ==========" >&2
  grep -E 'error:|undefined reference|fatal error:|Error [0-9]|recipe for target' "$build_log" | tail -40 >&2 || tail -80 "$build_log" >&2
  exit 1
fi

cd ..
# ngcp-rtpengine-kernel-dkms 依赖 lsb-release；invoke-rc.d 在 Docker 里拒绝 start 属正常
if ! dpkg -s lsb-release >/dev/null 2>&1; then
  apt-get update && apt-get install -y --no-install-recommends lsb-release
fi
dpkg -i ./*.deb || apt-get install -f -y
ldconfig
cd ..
rm -rf ims

if command -v rtpengine >/dev/null; then
  ldd "$(command -v rtpengine)" | grep -F "${FFMPEG_PREFIX}/lib" || {
    echo "rtpengine not linked against ${FFMPEG_PREFIX}" >&2
    ldd "$(command -v rtpengine)" | grep -E 'libav(codec|format|util)' >&2 || true
    exit 1
  }
fi

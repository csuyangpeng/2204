#!/usr/bin/env bash

set -e

cd "$(dirname "${BASH_SOURCE[0]}")"

version=4.8
wget https://github.com/ntop/nDPI/archive/refs/tags/${version}.tar.gz
tar -xf ${version}.tar.gz
pushd nDPI-${version}
unset MAKEFLAGS
# TARGETARCH 优先，uname -m 兜底（QEMU 模拟时 uname -m 也返回 aarch64）
ARCH="${TARGETARCH:-$(uname -m)}"
if [ "$ARCH" = "arm64" ] || [ "$ARCH" = "aarch64" ]; then
    # ARM64 QEMU 模拟下，gcc 内存占用大，大文件编译容易 OOM 导致 segfault：
    #   -O0: 禁用优化（-O2 更耗内存，容易触发 gcc 内部错误）
    #   -g0: 明确禁用调试信息（-g 生成 DWARF 很耗内存）
    #   -fno-var-tracking*: 禁用变量跟踪（GCC 最耗内存的 pass 之一）
    #   --param ggc-min-expand=0: 堆一扩展就触发 GC，压低峰值内存
    #   --param ggc-min-heapsize=: 调大 GC 堆阈值减少 GC 频率
    export CFLAGS="-O0 -g0 -fno-var-tracking -fno-var-tracking-assignments --param ggc-min-expand=0 --param ggc-min-heapsize=32768"
    export CXXFLAGS="-O0 -g0 -fno-var-tracking -fno-var-tracking-assignments --param ggc-min-expand=0 --param ggc-min-heapsize=32768"
    jobs=1
else
    jobs="$(nproc)"
fi
./autogen.sh
# nDPI 的 configure 会在 CFLAGS 末尾强制追加 -O2，覆盖我们的 -O0
# ARM64 QEMU 模拟下 -O2 容易触发 gcc 内部段错误，强制替换为 -O0
# Makefile 里还可能注入 -g/-g3/-ggdb，一并清除
if [ "$ARCH" = "arm64" ] || [ "$ARCH" = "aarch64" ]; then
    find . -name 'Makefile' -exec sed -i 's/-O2/-O0/g; s/ -g[0-9a-z]* / /g; s/ -g$/ /g' {} \;
fi
make -j"${jobs}"
make install
popd

rm ${version}.tar.gz nDPI-${version} -rf

#!/usr/bin/env bash
# 经 HTTP 代理访问 GitHub：HTTP/2 易 framing 错误；慢链路/ QEMU arm64 易 TLS 断流、early EOF
export http_proxy="${http_proxy:-http://10.18.11.52:7890}"
export https_proxy="${https_proxy:-http://10.18.11.52:7890}"
git config --global http.version HTTP/1.1
git config --global http.postBuffer 524288000
git config --global http.lowSpeedLimit 0
git config --global http.lowSpeedTime 999999

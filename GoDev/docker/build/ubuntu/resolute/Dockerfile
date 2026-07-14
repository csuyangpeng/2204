# ---------- builder：编译安装，层不进入最终镜像 ----------
FROM registry.cn-hangzhou.aliyuncs.com/10_18_1_2_5000/ubuntu:26.04 AS builder

ARG TARGETARCH
ENV TARGETARCH=${TARGETARCH}
ENV GOROOT=/opt/go
ENV GOPATH=/home/sder/go
ENV http_proxy="http://10.18.11.52:7890" https_proxy="http://10.18.11.52:7890" \
    no_proxy="mirrors.aliyun.com,10.18.1.2,registry.cn-hangzhou.aliyuncs.com,goproxy.cn"

COPY internal /tmp/internal

# 拆成多个 RUN，失败时只重跑失败层及之后（apt/golang/nDPI 可命中缓存）
RUN --mount=type=cache,target=/var/cache/apt,sharing=locked \
    --mount=type=cache,target=/var/lib/apt,sharing=locked \
    bash /tmp/internal/preinstall.sh && bash /tmp/internal/install_golang.sh

RUN bash /tmp/internal/install_3rd.sh

RUN --mount=type=secret,id=gitpass \
    bash /tmp/internal/install_tcpdump.sh

RUN rm -rf /tmp/internal

# ---------- 最终镜像：只保留运行时与 COPY 产物 ----------
FROM registry.cn-hangzhou.aliyuncs.com/10_18_1_2_5000/ubuntu:26.04

ARG TARGETARCH
ENV GOROOT=/opt/go
ENV GOPATH=/home/sder/go
ENV GO111MODULE=on
ENV GOPROXY="http://10.18.1.2:8081,https://goproxy.cn,direct" \
    GOSUMDB="sum.golang.org http://10.18.1.2:8081/sumdb/sum.golang.org" \
    GO111MODULE="on"
ENV PATH=${GOROOT}/bin:${GOPATH}/bin:${PATH}

COPY internal /tmp/internal
RUN --mount=type=cache,target=/var/cache/apt,sharing=locked \
    --mount=type=cache,target=/var/lib/apt,sharing=locked \
    bash /tmp/internal/install_runtime.sh && rm -rf /tmp/internal

COPY --from=builder /opt/go /opt/go
COPY --from=builder /usr/local /usr/local
RUN ldconfig 2>/dev/null || true

USER sder
WORKDIR /home/sder

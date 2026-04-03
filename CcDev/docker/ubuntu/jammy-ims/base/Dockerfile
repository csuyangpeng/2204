FROM registry.cn-hangzhou.aliyuncs.com/10_18_1_2_5000/ubuntu:22.04

ARG MAKEFLAGS=-j
ENV GITUSER=peng.yang GITPASS=11asdASD
# ENV http_proxy="http://10.18.11.52:7890" https_proxy="http://10.18.11.52:7890"
ENV LD_LIBRARY_PATH=$LD_LIBRARY_PATH:/usr/local/lib

COPY internal /tmp/internal
RUN bash /tmp/internal/preinstall.sh
# RUN bash /tmp/internal/install_msgpack.sh
RUN bash /tmp/internal/install_wfrest.sh
RUN bash /tmp/internal/install_cpp-httplib.sh
RUN bash /tmp/internal/install_cpprestsdk.sh
RUN bash /tmp/internal/install_etcd-cpp-apiv3.sh
# RUN bash /tmp/internal/install_cppkafka.sh
RUN bash /tmp/internal/install_nlohmannjson.sh
# RUN bash /tmp/internal/install_zlog.sh
# RUN bash /tmp/internal/install_core.sh
RUN bash /tmp/internal/postinstall.sh
RUN rm /tmp/internal/ -rf

ENV GITUSER= GITPASS= http_proxy= https_proxy=
USER root
WORKDIR /home/root/epc4g

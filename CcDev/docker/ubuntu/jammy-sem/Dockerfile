FROM ubuntu:16.04

ARG MAKEFLAGS=-j
ENV http_proxy="http://10.18.11.52:7890" https_proxy="http://10.18.11.52:7890"

COPY internal /tmp/internal
RUN bash /tmp/internal/preinstall.sh
RUN bash /tmp/internal/install_sems.sh
RUN bash /tmp/internal/postinstall.sh
RUN rm /tmp/internal/ -rf

ENV http_proxy= https_proxy=
RUN groupadd -r -g 996 sems && useradd -M -r -u 996 -g 996 sems
CMD ["/usr/local/sbin/sems", "-E", "-u", "sems", "-g", "sems",  "-f", "/etc/sems/sems.conf"]
STOPSIGNAL SIGTERM

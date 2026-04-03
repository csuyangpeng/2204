#!/usr/bin/env bash


set -e

cd "$(dirname "${BASH_SOURCE[0]}")"

pushd cJSON && mkdir -p build && cd build && cmake .. && make -j && make install && cd .. && popd
pushd hiredis && mkdir -p build && cd build && cmake .. && make -j && make install && cd .. && popd

git clone http://${GITUSER}:${GITPASS}@10.18.1.2:9999/develop/epc.git
# install epc lib
pushd epc
# sed -i 's/#if (!defined(SCTP_CONNECTX_4_ARGS) || (!defined(SCTP_RECVRCVINFO)) || (!defined(SCTP_SNDINFO)))/#if ((!defined(SCTP_CONNECTX_4_ARGS)) || (!defined(SCTP_RECVRCVINFO)) || (!defined(SCTP_SNDINFO)) || (!defined(SCTP_SEND_FAILED_EVENT)) || (!defined(SCTP_NOTIFICATIONS_STOPPED_EVENT)))/g' ./lib/freeDiameter-1.2.1/libfdcore/sctp.c
# sed -i 's/#include <sys\/sysctl.h>/\/* #include <sys\/sysctl.h>  deprecated *\//g' ./lib/ipfw/ipfw2.c ./lib/ipfw/tables.c ./lib/ipfw/dummynet.c
bash epc.sh buildrelease && bash epc.sh install  
popd

rm epc -rf

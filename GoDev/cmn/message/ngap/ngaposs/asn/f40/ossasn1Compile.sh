#!/bin/bash

asn1codegen="/usr/local/ossasn1/linux-glibc2.3-amd64/10.1.1/bin/asn1"
asn1dflt="/usr/local/ossasn1/linux-glibc2.3-amd64/10.1.1/asn1dflt.linux-amd64"

asn1out="ngapToed_x86"
asn1opts="-codefile ngapToed_x86.cc -per -C++ -2002 -constraints -autoencdec -output"
asn1srcs="
  ngap-common.asn
  ngap-constants.asn
  ngap-containers.asn
  ngap-ies.asn
  ngap-pdus.asn
  ngap-procedures.asn
  "
$asn1codegen $asn1dflt $asn1srcs $asn1opts $asn1out
mv ngapToed_x86.* ../../ngapToed 


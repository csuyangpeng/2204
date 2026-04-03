#include <string.h>
#include <iostream>

#include "snssaiPdu.h"

SnssaiPdu::SnssaiPdu()
{
    reset_v();
}

SnssaiPdu::~SnssaiPdu()
{
}

void SnssaiPdu::reset_v()
{
    snssai_mp = NULL;
    snssai_m.bit_mask = 0;
    snssai_m.iE_Extensions = NULL;
    snssai_m.sST.length = 0;
    memset(snssai_m.sST.value, 0xff,sizeof(snssai_m.sST.value));
    snssai_m.sD.length = 0;
    memset(snssai_m.sD.value, 0xff,sizeof(snssai_m.sD.value));
}

unsigned int 
SnssaiPdu::decodePerToOss(NgapOssCtxt* ctxt)
{
   if(NULL == ctxt)
   {
       return -1;
   }

   OssBuf inputBuf;
   OssBuf outputBuf;
   int pduNum = S_NSSAI_PDU;

   inputBuf = perBufRef_m;

   ctxt->reset_v();
   outputBuf = ctxt->getPrimaryCodecBuf();

   // default action -- automatically fully decode the message
   unsigned int result = ctxt->decodePerToOss(&pduNum, &inputBuf, &outputBuf);

   if (result == 0)
   {
      ossBufRef_m = outputBuf;
      getDecodedOssDataPtrs_v();
   }

   return result;
}

unsigned int
SnssaiPdu::encodeOssToPer(NgapOssCtxt *ctxt)
{
    if (NULL == ctxt){
        return -1;
    }

    S_NSSAI ngApPdu;
    ngApPdu.bit_mask = 0;
    ngApPdu.iE_Extensions = NULL;

    ngApPdu.sST.length = 1;
    memcpy(ngApPdu.sST.value,snssai_m.sST.value,sizeof(snssai_m.sST.value));
    std::cout<< "the 2 ceng shen du value:" << snssai_m.sST.length <<std::endl;
    std::cout<< "fu zhi zhi hou de chang du:" << int(ngApPdu.sST.value[0]) <<std::endl;
    if (snssai_m.bit_mask & sD_present) {
        ngApPdu.bit_mask |= sD_present;
        ngApPdu.sD.length = snssai_m.sD.length;
        memcpy(ngApPdu.sD.value,snssai_m.sD.value,snssai_m.sD.length);
    }
    
    //Convert Structures to PER
    OssBuf outputBuf;
    unsigned int result;

    ctxt->reset_v();
    outputBuf = ctxt->getPrimaryCodecBuf();

    std::cout <<"bian ma zi jie chang du:" << sizeof(ngApPdu)<<std::endl;

    result = ctxt->encodeOssToPer(S_NSSAI_PDU, &ngApPdu, &outputBuf);

    if (result == encodeSuccessCode_c){
        perBufRef_m = outputBuf;
    }

    return result;
}

void SnssaiPdu::getDecodedOssDataPtrs_v()
{
    reset_v();

    S_NSSAI* ptr = reinterpret_cast<S_NSSAI *>(ossBufRef_m.value);

    if (ptr == NULL){
        return;
    }
    snssai_mp = ptr;
}

S_NSSAI * 
SnssaiPdu::getSnssaiEncodePtr()
{
    return &snssai_m;
}

S_NSSAI * 
SnssaiPdu::getSnssaiDecodePtr()
{
    return snssai_mp;
}

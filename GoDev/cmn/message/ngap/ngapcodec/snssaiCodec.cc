#include <iostream>
#include <sstream>
#include <string>

#include "snssaiCodec.h"

using namespace std;

SnssaiCodec::SnssaiCodec() {
  reset();
}

SnssaiCodec::~SnssaiCodec() {}

void 
SnssaiCodec::setSnssai(SNssai& snssai)
{
    snssaiValue_m = snssai;
}

SNssai& 
SnssaiCodec::getSnssai()
{
    return snssaiValue_m;
}

bool SnssaiCodec::reset()
{
    msgPdu_m.reset_v();
    memset(&snssaiValue_m,0,sizeof(snssaiValue_m));
}

MsgBuffer SnssaiCodec::encode(NgapOssCtxt *ctxt)
{
    MsgBuffer msgBuffer;
    msgBuffer.length = 0;
    msgBuffer.value = NULL;

    //snssai
    S_NSSAI* snssaiPtr = msgPdu_m.getSnssaiEncodePtr();
    
    snssaiPtr->sST.length = sizeof(snssaiValue_m.sst);
    std::cout <<"the sst length:" << snssaiPtr->sST.length <<std::endl;
    memcpy(snssaiPtr->sST.value,
           snssaiValue_m.sst,
           snssaiPtr->sST.length);

    if (snssaiValue_m.sdPresent) {
      snssaiPtr->bit_mask |= sD_present;
      snssaiPtr->sD.length = sizeof(snssaiValue_m.sd);
      memcpy(snssaiPtr->sD.value,
             snssaiValue_m.sd,
             snssaiPtr->sD.length);
    }

    msgPdu_m.encodeOssToPer(ctxt);

    OssBuf *perBuf = msgPdu_m.getPerBufRef();

    msgBuffer.length = perBuf->length;
    msgBuffer.value = perBuf->value;

    return msgBuffer;

}
bool SnssaiCodec::decode(NgapOssCtxt *ctxt, MsgBuffer msgBuf)
{
    reset();

    OssBuf ossbuf;
    ossbuf.length = msgBuf.length;
    ossbuf.value = msgBuf.value;

    msgPdu_m.setPerBufRef_v(ossbuf);

    if (0 != msgPdu_m.decodePerToOss(ctxt)) {
       return false;
    }

    S_NSSAI* snssaiPtr = msgPdu_m.getSnssaiDecodePtr();
    memcpy(snssaiValue_m.sst,
           snssaiPtr->sST.value,
           sizeof(snssaiPtr->sST.value));
    if (snssaiPtr->bit_mask & sD_present) {

        snssaiValue_m.sdPresent = true;
        memcpy(snssaiValue_m.sd,
             snssaiPtr->sD.value,
             snssaiPtr->sD.length);
    }
}

std::string SnssaiCodec::toString()
{

  ostringstream oss;

  oss << "snssai: sst(" << int(snssaiValue_m.sst[0]) << ")" << endl
      << "sd:prst(" << snssaiValue_m.sdPresent << ")" << endl
      << "sd(" <<int(snssaiValue_m.sd[0]) <<snssaiValue_m.sd[1] <<snssaiValue_m.sd[2] << ")" << endl;
      
  return oss.str();
}

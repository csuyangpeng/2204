#ifndef __snssaiCodec_H__
#define __snssaiCodec_H__

#include <string>
#include <ngapMsgTypes.h>
#include <ngapOssCtxt.h>
#include <snssaiPdu.h>

class SnssaiCodec
{
public:
  SnssaiCodec();
  ~SnssaiCodec();

  void setSnssai(SNssai& snssai);
  SNssai& getSnssai();

  MsgBuffer encode(NgapOssCtxt *ctxt);
  bool decode(NgapOssCtxt *ctxt, MsgBuffer msgBuf);

  bool reset();

  std::string toString();

private:

  SnssaiPdu msgPdu_m;

  SNssai snssaiValue_m;
};
#endif

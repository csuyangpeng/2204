#ifndef __pduSessResMdfyUnSuccTransferCodec_H__
#define __pduSessResMdfyUnSuccTransferCodec_H__

#include <string>
#include <ngapMsgTypes.h>
#include <ngapOssCtxt.h>
#include <pduSessionResourceModifyUnSuccTransferPdu.h>

class PduSessResMdfyUnSuccTransferCodec
{
public:
  PduSessResMdfyUnSuccTransferCodec();
  ~PduSessResMdfyUnSuccTransferCodec();

  void setCause(Uint8 cType, Uint8 cValue);
  Uint8 getCauseType();
  Uint8 getCauseValue();

  MsgBuffer encode(NgapOssCtxt *ctxt);
  bool decode(NgapOssCtxt *ctxt, MsgBuffer msgBuf);

  bool reset();

  std::string toString();

private:

  PduSessResMdfyUnsuccTransferPdu msgPdu_m;

  Uint8 causeType_m;
  Uint8 causeValue_m;
};
#endif


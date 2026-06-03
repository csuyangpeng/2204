#ifndef __pduSessionResourceReleaseCommandTransferCodec_H__
#define __pduSessionResourceReleaseCommandTransferCodec_H__

#include <string>
#include <ngapMsgTypes.h>
#include <ngapOssCtxt.h>
#include <pduSessionResourceReleaseCommandTransferPdu.h>

class PduSessResRelCmdTransferCodec
{
public:
  PduSessResRelCmdTransferCodec();
  ~PduSessResRelCmdTransferCodec();

  void setCause(Uint8 cType, Uint8 cValue);
  Uint8 getCauseType();
  Uint8 getCauseValue();

  MsgBuffer encode(NgapOssCtxt *ctxt);
  bool decode(NgapOssCtxt *ctxt, MsgBuffer msgBuf);

  bool reset();

  std::string toString();

private:
	
  PduSessResRelCmdTransferPdu msgPdu_m;

  Uint8 causeType_m;
  Uint8 causeValue_m;
};
#endif
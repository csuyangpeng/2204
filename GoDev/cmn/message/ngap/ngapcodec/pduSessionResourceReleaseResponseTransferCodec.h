#ifndef __pduSessionResourceReleaseResponseTransferCodec_H__
#define __pduSessionResourceReleaseResponseTransferCodec_H__

#include <string>
#include <ngapMsgTypes.h>
#include <ngapOssCtxt.h>
#include <pduSessionResourceReleaseResponseTransferPdu.h>

class PduSessResRelRespTransferCodec
{
public:
  PduSessResRelRespTransferCodec();
  ~PduSessResRelRespTransferCodec();

  bool isSecRatUsageInfoPrst();
  void setSecRatUsageInfo(SecRatUsageInformation& secRatUsage);
  SecRatUsageInformation& getSecRatUsageInfo();

  MsgBuffer encode(NgapOssCtxt *ctxt);
  bool decode(NgapOssCtxt *ctxt, MsgBuffer msgBuf);

  bool reset();

  std::string toString();

private:

   PduSesResRelRespTransferPdu msgPdu_m;

   bool isSecRatUsageInfoPrst_m;
   SecRatUsageInformation secRatUsageInfo_m;

};
#endif

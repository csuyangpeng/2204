#ifndef __pduSessionResourceModifyResponseCodec_H__
#define __pduSessionResourceModifyResponseCodec_H__

#include <string>
#include <ngapMsgTypes.h>
#include <ngapOssCtxt.h>
#include <pduSessionResourceModifyResponsePdu.h>

class PduSessResModifyRespCodec
{
public:
  PduSessResModifyRespCodec();
  ~PduSessResModifyRespCodec();

  void setRanUeNgapId(Uint32 id);
  Uint32 getRanUeNgapId();

  void setAmfUeNgapId(Uint64 id);
  Uint64 getAmfUeNgapId();

  void addPduSessResModifyRespList(PduSessResModifyRespItem &item);
  PduSessResModifyRespList& getPduSessResModifyRespList();
  bool isPduSessResMdfyRespListPrst();

  void addPduSessResFailedModifyList(PduSessResFailedMdfyRespItem &item);
  PduSessResFailedModifyList& getPduSessResFailedModifyList();
  bool isPduSessResFailMdfyRespListPrst();

  void setUserLocationInfo(UserLocationInfo& info);
  UserLocationInfo& getUserLocationInfo();
  bool isUserLocInfoPrst();

  MsgBuffer encode(NgapOssCtxt *ctxt);
  bool decode(NgapOssCtxt *ctxt, MsgBuffer msgBuf);

  bool reset();

  std::string toString();

private:
  PduSessResMdfyRespPdu msgPdu_m;

  Uint64 amfUeNgapId_m;
  Uint32 ranUeNgapId_m;

  PduSessResModifyRespList pduSessResModRespList_m;
  bool pduSessResModRespListPrst_m;
  PduSessResFailedModifyList pduSessResFailedModList_m;
  bool pduSessResFailModRespListPrst_m;

  UserLocationInfo userLocationInfo_m;
  bool userLocInfoPrst_m;
};

#endif


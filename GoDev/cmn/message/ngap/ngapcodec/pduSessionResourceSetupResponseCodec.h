#ifndef __pduSessionResourceSetupResponseCodec_H__
#define __pduSessionResourceSetupResponseCodec_H__

#include <string>
#include <ngapMsgTypes.h>
#include <ngapOssCtxt.h>
#include <pduSessionResourceSetupResponsePdu.h>

class PduSessResSetupRespCodec
{
public:
  PduSessResSetupRespCodec();
  ~PduSessResSetupRespCodec();

  void setRanUeNgapId(Uint32 id);
  Uint32 getRanUeNgapId();

  void setAmfUeNgapId(Uint64 id);
  Uint64 getAmfUeNgapId();

  void addPduSessResSetupRespList(PduSessResSetupRespItem &item);
  PduSessResSetupRespList& getPduSessResSetupRespList();

  void addPduSessResFailedSetupList(PduSessResFailedSetupItem &item);
  PduSessResFailedSetupList& getPduSessResFailedToSetupList();

  //bool encode(NgapOssCtxt* ctxt, MsgBuffer msgBuffer);
  MsgBuffer encode(NgapOssCtxt *ctxt);
  bool decode(NgapOssCtxt *ctxt, MsgBuffer msgBuf);

  bool reset();

  std::string toString();

private:
  PDUSessionResourceSetupResponsePdu msgPdu_m;

  Uint64 amfUeNgapId_m;
  Uint32 ranUeNgapId_m;

  PduSessResSetupRespList pduSessResSetupRespList_m;
  PduSessResFailedSetupList pduSessResFailedSetupList_m;
};

#endif


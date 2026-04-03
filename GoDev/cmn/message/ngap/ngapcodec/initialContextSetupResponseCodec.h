#ifndef __initialContextSetupResponseCodec_H__
#define __initialContextSetupResponseCodec_H__

#include <string>
#include <ngapMsgTypes.h>
#include <ngapOssCtxt.h>
#include <initialContextSetupResponsePdu.h>

class InitialContextSetupRespCodec
{
public:
  InitialContextSetupRespCodec();
  ~InitialContextSetupRespCodec();

  void setRanUeNgapId(Uint32 id);
  Uint32 getRanUeNgapId();

  void setAmfUeNgapId(Uint64 id);
  Uint64 getAmfUeNgapId();

  void addPduSessResSetupRespList(PduSessResSetupRespItem &item);
  PduSessResSetupRespList& getPduSessResSetupRespList();

  void addPduSessResFailedSetupList(PduSessResFailedSetupItem &item);
  PduSessResFailedSetupList& getPduSessResFailedToSetupList();

  MsgBuffer encode(NgapOssCtxt *ctxt);
  //bool encode(NgapOssCtxt* ctxt, MsgBuffer msgBuffer);
  bool decode(NgapOssCtxt *ctxt, MsgBuffer msgBuf);

  bool reset();

  std::string toString();

private:
  InitialContextSetupResponsePdu msgPdu_m;

  Uint64 amfUeNgapId_m;
  Uint32 ranUeNgapId_m;

  PduSessResSetupRespList pduSessResSetupRespList_m;
  PduSessResFailedSetupList pduSessResFailedSetupList_m;
};

#endif
#ifndef __pduSessResSetupReqCodec_H__
#define __pduSessResSetupReqCodec_H__

#include <string>
#include <ngapMsgTypes.h>
#include <ngapOssCtxt.h>
#include <pduSessionResourceSetupRequestPdu.h>

class PduSessResSetupReqCodec
{
public:
  PduSessResSetupReqCodec();
  ~PduSessResSetupReqCodec();

  void setRanUeNgapId(Uint32 id);
  Uint32 getRanUeNgapId();

  void setAmfUeNgapId(Uint64 id);
  Uint64 getAmfUeNgapId();

  void setNasPdu(std::string naspdu);
  std::string getNasPdu();
  bool isNasPduPrst();

  void setUeAmbr(UeAmbr& ambr);
  UeAmbr& getUeAmbr();
  bool isUeAMBRPrst();

  void setRanPagingPriority(Uint16 val);
  Uint16 getRanPagingPriority();
  bool isRanPagPriPrst();

  void addPduSessResSetupReqList(PduSessResSetupReqItem& item);
  PduSessResSetupReqList &getPduSessResSetupReqList();

  MsgBuffer encode(NgapOssCtxt* ctxt);
  bool decode(NgapOssCtxt *ctxt, MsgBuffer msgBuf);

  bool reset();

  std::string toString();

private:
	
  PDUSessionResourceSetupRequestPdu msgPdu_m;

  Uint64 amfUeNgapId_m;
  Uint32 ranUeNgapId_m;

  std::string nasPdu_m;
  bool nasPduPrst_m;

  UeAmbr ueAmbr_m;
  bool ueAMBRPrst_m;

  Uint16 ranPagingPriority_m;
  bool ranPagingPriPrst_m;

  PduSessResSetupReqList pduSessResSetupReqList_m;
};

#endif
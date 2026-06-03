#ifndef __pduSessionResourceSetupResponseTransferCodec_H__
#define __pduSessionResourceSetupResponseTransferCodec_H__

#include <string>
#include <ngapMsgTypes.h>
#include <ngapOssCtxt.h>
#include <pduSessionResourceSetupResponseTransferPdu.h>

class PduSesResSetupRespTransferCodec
{
public:
  PduSesResSetupRespTransferCodec();
  ~PduSesResSetupRespTransferCodec();

  void setGtpTunnel(GtpTunnelInfo& tunnel);
  GtpTunnelInfo& getGtpTunnel();
  void addQosFlowIndList(AssQosFlowItem& val);
  QosFlowIndList& getQosFlowIndList();

  bool isQosFlowTNLInfoListPrst();
  void addQosFlowTNLInfoList(AddQosFlowPerTNLInfo& item);
  AddQosFlowPerTNLInfoList& getQosFlowTNLInfoList();

  bool isSecResultPrst();
  void setSecResult(Uint8 intInd, Uint8 confdInd);
  Uint8 getIndPrctResult();
  Uint8 getConfdPrctResult();

  bool isFailedSetupListPrst();
  void addQosFlowFailedSetupList(QosFlowCodecItem& item);
  QosFlowCodecList& getQosFlowFailedSetupList();

  //bool encode(NgapOssCtxt* ctxt, MsgBuffer msgBuffer);
  MsgBuffer encode(NgapOssCtxt *ctxt);
  bool decode(NgapOssCtxt *ctxt, MsgBuffer msgBuf);

  bool reset();

  std::string toString();

private:
	
  PduSesResSetupRespTransferPdu msgPdu_m;

  GtpTunnelInfo gtpTunnel_m;
  QosFlowIndList qfiList_m;

  bool AddQosFlowTNLListPrst_m;
  AddQosFlowPerTNLInfoList addQosFlowTNLInfoList_m;

  bool isSecResult_m;
  Uint8 intProtectRst_m;
  Uint8 confdProtecRst_m;

  bool isFailedSetupListPrst_m;  
  QosFlowCodecList qosFlowFailedSetupList_m;  
};
#endif

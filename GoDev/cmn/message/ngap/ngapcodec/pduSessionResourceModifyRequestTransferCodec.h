#ifndef __pduSessionResourceModifyRequestTransferCodec_H__
#define __pduSessionResourceModifyRequestTransferCodec_H__

#include <string>
#include <ngapMsgTypes.h>
#include <ngapOssCtxt.h>
#include <pduSessionResourceModifyRequestTransferPdu.h>

class PduSesResModReqTransferCodec
{
public:
  PduSesResModReqTransferCodec();
  ~PduSesResModReqTransferCodec();

  bool isSessAMBRPrst();
  void setSessAmbr(long ulBitRate, long dlBitRate);
  long getSessAmbrUl();
  long getSessAmbrDl();

  bool isNtwkInstancePrst();
  void setNtwkInstance(Uint64 ntwkInstanceVal);
  Uint64 getNtwkInstance();

  void addUlNguUpTnlModifyList(UlNguUpTnlModifyItem& item);
  UlNguUpTnlModifyList& getUlNguUpTnlModifyList();

  void addAdditionalGtpTunnelList(GtpTunnel& item);
  AddUpTransportLayerInfoList& getAdditionalGtpTunnelList();
  bool isAddGtpTunnelPrst();
  
  void addQosFlowAddOrModReqList(QosFlowAddOrModReqItem& item);
  QosFlowAddOrModReqList& getQosFlowAddOrModReqList();
  
  void addQosFlowToReleaseList(QosFlowCodecItem& item);
  QosFlowCodecList& getQosFlowToReleaseList();

  void setCmmnNtwrkInstance(std::string networkInstance);
  std::string getCmmnNtwrkInstance();
  bool isCmmnNtwrkInstacePrst();

  MsgBuffer encode(NgapOssCtxt *ctxt);
  bool decode(NgapOssCtxt *ctxt, MsgBuffer msgBuf);

  bool reset(); 

  std::string toString();

private:

  PduSesResModifyReqTransferPdu msgPdu_m;

  bool isSessAMBRPrst_m;
  long sessAmbrULBitRate_m;
  long sessAmbrDLBitRate_m;
  
  Uint64 ntwkInstance_m;
  bool isNtwkInstancePrst_m;

  UlNguUpTnlModifyList ulNguUpTnlModifyList_m;

  AddUpTransportLayerInfoList addGtpTunnelList_m;
  bool addGtpTunnelListPrst_m;

  QosFlowAddOrModReqList qosFlowAddOrModReqList_m;

  QosFlowCodecList qosFlowToReleaseList_m;

  std::string cmmnNtwkInstance_m;
  bool cmmnNtwkInstancePrst_m;
  
};

#endif


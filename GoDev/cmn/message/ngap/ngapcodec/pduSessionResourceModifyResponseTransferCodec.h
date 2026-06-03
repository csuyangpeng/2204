#ifndef __pduSessionResponseModifyResponseTransferCodec_H__
#define __pduSessionResponseModifyResponseTransferCodec_H__

#include <string>
#include <ngapMsgTypes.h>
#include <ngapOssCtxt.h>
#include <pduSessionResourceModifyResponseTransferPdu.h>

class PduSessResMdfyRespTransferCodec
{
public:
  PduSessResMdfyRespTransferCodec();
  ~PduSessResMdfyRespTransferCodec();

  void setDlGtpTunnel(GtpTunnelInfo& tunnel);
  GtpTunnelInfo& getDlGtpTunnel();
  bool isDlGtpTunnelPrst();

  void setUlGtpTunnel(GtpTunnelInfo& tunnel);
  GtpTunnelInfo& getUlGtpTunnel();
  bool isUlGtpTunnelPrst();

  bool isQosFlowTNLInfoListPrst();
  void addQosFlowTNLInfoList(AddQosFlowPerTNLInfo& item);
  AddQosFlowPerTNLInfoList& getQosFlowTNLInfoList();

  bool isQosFlowAddOrMdfyRespListPrst();
  void addQosFlowAddOrMdfyRespList(Uint32 item);
  std::vector<unsigned int>& getQosFlowAddOrMdfyRespList();
  
  bool isFailedAddOrMdfyListPrst();
  void addQosFlowFailedAddOrMdfyList(QosFlowCodecItem& item);
  QosFlowCodecList& getQosFlowFailedAddOrMdfyList();
  
  MsgBuffer encode(NgapOssCtxt *ctxt);
  bool decode(NgapOssCtxt *ctxt, MsgBuffer msgBuf);

  bool reset();

  std::string toString();

private:
    
  PduSesResModifyRespTransferPdu msgPdu_m;

  bool dlGtpTunnelPrst_m;
  GtpTunnelInfo dlGtpTunnel_m;

  bool ulGtpTunnelPrst_m;
  GtpTunnelInfo ulGtpTunnel_m;

  bool AddQosFlowTNLListPrst_m;
  AddQosFlowPerTNLInfoList addQosFlowTNLInfoList_m;

  bool isQosFlowAddOrMdfyRespListPrst_m;
  std::vector<unsigned int> qosFlowAddOrMdfyRespList_m;

  bool isFailedAddOrMdfyListPrst_m;
  QosFlowCodecList qosFlowFailedAddOrMdfyList_m;
};
#endif

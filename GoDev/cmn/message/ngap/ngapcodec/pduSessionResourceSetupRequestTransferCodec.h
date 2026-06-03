#ifndef __pduSessionResourceSetupRequestTransferCodec_H__
#define __pduSessionResourceSetupRequestTransferCodec_H__

#include <string>
#include <ngapMsgTypes.h>
#include <ngapOssCtxt.h>
#include <pduSessionResourceSetupRequestTransferPdu.h>

class PduSesResSetupReqTransferCodec
{
public:
  PduSesResSetupReqTransferCodec();
  ~PduSesResSetupReqTransferCodec();

  bool isSessAMBRPrst();
  void setSessAmbr(long ulBitRate, long dlBitRate);
  long getSessAmbrUl();
  long getSessAmbrDl();

  bool isNtwkInstancePrst();
  void setNtwkInstance(Uint64 ntwkInstanceVal);
  Uint64 getNtwkInstance();

  void setUpGtpTunnel(GtpTunnelInfo& tunnel);
  GtpTunnelInfo& getUpGtpTunnel();

  void addUpTransLayerInfoList(GtpTunnel& item);
  AddUpTransportLayerInfoList& getUpTransLayerInfoList();
  bool isAddUpTransLayerInfoListPrst();

  bool isDataFwNotPossiblePrst();
  void setDataFwNotPssble(Uint8 value);
  Uint8 getDataFwNotPssble();

  void setPduSessType(Uint8 value);
  Uint8 getPduSessType();

  bool isSecIndPrst();
  void setSecInd(Uint8 intInd, Uint8 confdInd);
  Uint8 getSecIndProtectInd();
  Uint8 getSecConfdProtectInd();
  bool isMaxPrtDataRatePrst();
  void setMaxPrtDataRate(Uint8 maxPrtData);
  Uint8 getMaxPrtDataRate();
  bool isMaxPrtDataRateDlPrst();
  void setMaxPrtDataRateDl(Uint8 maxPrtDataDl);
  Uint8 getMaxPrtDataRateDl();
    
  void addQosFlowSetupReqList(QosFlowSetupReqItem& item);
  QosFlowSetupReqList& getQosFlowSetupReqList();

  void setCommonNetworkInstance(std::string instance);
  std::string getCommonNetworkInstance();
  bool isCommonNetworkInstancePrst();

  MsgBuffer encode(NgapOssCtxt *ctxt);
  //bool encode(NgapOssCtxt* ctxt, MsgBuffer msgBuffer);
  bool decode(NgapOssCtxt *ctxt, MsgBuffer msgBuf);

  bool reset(); 

  std::string toString();

private:
    
  PduSesResSetupReqTransferPdu msgPdu_m;

  bool isSessAMBRPrst_m;
  long sessAmbrULBitRate_m;
  long sessAmbrDLBitRate_m;
  
  Uint64 ntwkInstance_m;
  bool isNtwkInstancePrst_m;

  GtpTunnelInfo upGtpTunnel_m;
  bool isAddTransLayerInfoListPrst_m;
  AddUpTransportLayerInfoList addUpTransportLayerInfoList_m;

  bool isDataFwNotPossible_m;
  Uint8 dataFwNotPssble_m;
  
  Uint8 pduSessType_m;

  bool isSecInd_m;
  Uint8 intProtectInd_m;
  Uint8 confdProtecInd_m;
  bool isMaxPrtDataRate_m;
  Uint8 maxPrtDataRate_m;
  bool maxIntProDataRateDownlinkPrst_m;
  Uint8 maxIntProDataRateDownlink_m;

  QosFlowSetupReqList qosFlowSetupReqList_m;

  std::string cmmnNtwrkInstance_m;
  bool cmmnNtwrkInstancePrst_m;
  
};

#endif


#ifndef __pduSessionResourceSetupRequestTransferPdu_h__INCLUDED__
#define __pduSessionResourceSetupRequestTransferPdu_h__INCLUDED__

#include "ngapOssPdu.h"
#include <ngapMsgTypes.h>
///////////////////////////////////////////////////////////////////////////////
// @class PDUSessionResourceSetupRequestTransferPdu
//-- **************************************************************
//--
//-- PDU Session Resource Setup Request Transfer 
//--
//-- **************************************************************
//
// PDUSessionResourceSetupRequestTransfer ::= SEQUENCE {
//  protocolIEs     ProtocolIE-Container    { {PDUSessionResourceSetupRequestTransfer-IEs} },
//  ...
//}
// PDUSessionResourceSetupRequestTransfer-IEs NGAP-PROTOCOL-IES ::= {
//  { ID id-PDUSessionAggregateMaximumBitRate       CRITICALITY reject  TYPE PDUSessionAggregateMaximumBitRate      PRESENCE optional   }|
//  { ID id-UPTransportLayerInformation             CRITICALITY reject  TYPE UPTransportLayerInformation            PRESENCE mandatory  }|
//  { ID id-AdditionalUL-NGU-UP-TNLInformation      CRITICALITY reject  TYPE UPTransportLayerInformationList        PRESENCE optional   }|
//  { ID id-DataForwardingNotPossible               CRITICALITY reject  TYPE DataForwardingNotPossible              PRESENCE optional   }|
//  { ID id-PDUSessionType                          CRITICALITY reject  TYPE PDUSessionType                         PRESENCE mandatory  }|
//  { ID id-SecurityIndication                      CRITICALITY reject  TYPE SecurityIndication                     PRESENCE optional   }|
//  { ID id-NetworkInstance                         CRITICALITY reject  TYPE NetworkInstance                        PRESENCE optional   }|
//  { ID id-QosFlowSetupRequestList                 CRITICALITY reject  TYPE QosFlowSetupRequestList                PRESENCE mandatory  }|
//  { ID id-CommonNetworkInstance                   CRITICALITY reject  TYPE CommonNetworkInstance                  PRESENCE optional },
//  ...
//}
///////////////////////////////////////////////////////////////////////////////
class PduSesResSetupReqTransferPdu : public NgapOssPdu
{
  public:
    PduSesResSetupReqTransferPdu();
    //virtual
    ~PduSesResSetupReqTransferPdu();

    /////////////////////////////////////////////////////////////
    /// @brief reset_v
    ///  resets all data members to default value
    /////////////////////////////////////////////////////////////
    //virtual
    void reset_v();

    /////////////////////////////////////////////////////////////
    /// @brief encode Oss to PER using ASN.1
    ///  builds and processes the OSS PDU to encode
    ///  it to PER using ASN.1
    /////////////////////////////////////////////////////////////
    //virtual
    unsigned int encodeOssToPer(NgapOssCtxt *ctxt);

    //overwrite the decode methord
    unsigned int decodePerToOss(NgapOssCtxt* ctxt);

    /////////////////////////////////////////////////////////////
    /// @brief get Decoded Oss Data Pointers
    ///  processes the decoded OSS buffer and populates
    ///  the data access pointers in the class for the IEs present in the
    ///  decoded OSS buffer
    /////////////////////////////////////////////////////////////
    //virtual
    void getDecodedOssDataPtrs_v();

    /////////////////////////////////////////////////////////////
    /// @brief get PDUSessionResourceSetupRequestTransfer pointer
    /// returns the pointer to the
    /// PDUSessionResourceSetupRequestTransfer in the decoded message
    /////////////////////////////////////////////////////////////
    PDUSessionResourceSetupRequestTransfer *getPduSessResSetupReqTransferDecodedPtr() const;

    /////////////////////////////////////////////////////////////
    /// PDUSessionAggregateMaximumBitRate
    /////////////////////////////////////////////////////////////
    PDUSessionAggregateMaximumBitRate *getSessAmbrDecodedPtr() const;
    PDUSessionAggregateMaximumBitRate *getSessAmbrEncodePtr();
    bool isSessAmbrPrst();

    /////////////////////////////////////////////////////////////
    /// UPTransportLayerInformation
    /////////////////////////////////////////////////////////////
    UPTransportLayerInformation *getUpTransLayerInfoDecodedPtr() const;
    UPTransportLayerInformation *getUpTransLayerInfoEncodePtr();

    /////////////////////////////////////////////////////////////
    /// Addtional UPTransportLayerInformation
    /////////////////////////////////////////////////////////////
    UPTransportLayerInformationList* getUpTransLayerInfoListDecodedPtr() const;
    UPTransportLayerInformationList* getUpTransLayerInfoListEncodePtr();

    UPTransportLayerInformationItem* getUpTransLayerInfoItemDecodedPtr(unsigned int index) const;
    UPTransportLayerInformationItem* getUpTransLayerInfoItemEncodePtr(unsigned int index);

    void setNumofUpTransLayerInfoListItem_m_v(unsigned int num);
    unsigned int getNumofUpTransLayerInfoListItem() const;
    bool isUpTransLayerInfoListPrst();

    /////////////////////////////////////////////////////////////
    /// DataForwardingNotPossible
    /////////////////////////////////////////////////////////////
    DataForwardingNotPossible *getDataFwdNotPssbleDecodedPtr() const;
    DataForwardingNotPossible *getDataFwdNotPssbleEncodePtr();
    bool isDataFwdNotPssblePrst();

    /////////////////////////////////////////////////////////////
    /// PDUSessionType
    /////////////////////////////////////////////////////////////
    PDUSessionType *getPduSessTypeDecodedPtr() const;
    PDUSessionType *getPduSessTypeEncodePtr();

    /////////////////////////////////////////////////////////////
    /// SecurityIndication
    /////////////////////////////////////////////////////////////
    SecurityIndication *getSecIndDecodedPtr() const;
    SecurityIndication *getSecIndEncodePtr();
    bool isSecIndPrst();

    MaximumIntegrityProtectedDataRate* getMaxIntProDataRateEncodePtr();
    SecurityIndicationIeExtension* getSecIndIeExtenDecodedPtr();

    /////////////////////////////////////////////////////////////
    /// NetworkInstance
    /////////////////////////////////////////////////////////////
    NetworkInstance *getNetworkInstanceDecodedPtr() const;
    NetworkInstance *getNetworkInstanceEncodePtr();
    bool isNetworkInstancePrst();

    /////////////////////////////////////////////////////////////
    /// QosFlowSetupRequestList
    /////////////////////////////////////////////////////////////
    QosFlowSetupRequestList* getQosFlowSetupReqListDecodedPtr() const;
    QosFlowSetupRequestList* getQosFlowSetupReqListEncodePtr();
  
    QosFlowSetupRequestItem* getQosFlowSetupReqItemDecodedPtr(unsigned int index) const;
    QosFlowSetupRequestItem* getQosFlowSetupReqItemEncodePtr(unsigned int index);
  
    void setNumofQosFlowSetupReqListItem_m_v(unsigned int num);
    unsigned int getNumofQosFlowSetupReqListItem() const;

    /////////////////////////////////////////////////////////////
    /// CommonNetworkInstance
    /////////////////////////////////////////////////////////////
    CommonNetworkInstance * getCommonNetworkInstanceDecodedPtr() const;
    CommonNetworkInstance * getCommonNetworkInstanceEncodePtr();
    bool isCommonNetworkInstancePrst();

  private:
    typedef PduSessResSetupReqTransferIe ProtocolIeType;
    typedef SecurityIndicationIeExtension IeExtenIeType;

    //OSS structures
    PDUSessionResourceSetupRequestTransfer* pduSessResSetupReqTransfer_mp;

    //Session ambr
    PDUSessionAggregateMaximumBitRate *sessAMBR_mp; //for decode
    PDUSessionAggregateMaximumBitRate sessAMBR_m; //for encode
    bool isSessAmbrPrst_m;

    //up transport layer information
    UPTransportLayerInformation *upTransportLayerInfo_mp; //for decode
    UPTransportLayerInformation upTransportLayerInfo_m;   //for encode
    TransLayerAddr transLayerAddress_m;

    //addtional up transport information
    UPTransportLayerInformationList upTransportLayerInfoList_m;
    UPTransportLayerInformationList* upTransportLayerInfoList_mp;
    struct UPTransportLayerInformationList_ upTransportLayerInfoList_ma[maxnoofMultiConnectivityMinusOne_c];
    int numofUpTransportLayerInfo_m;
    bool isUpTransportLayerInfoPrst_m;

    TransLayerAddr addTransLayerAddress_m[maxnoofMultiConnectivityMinusOne_c];

    //data forwarding not possible 
    DataForwardingNotPossible *dataFwdNotPssble_mp; //for decode
    DataForwardingNotPossible dataFwdNotPssble_m; //for encode
    bool isDataFwdNotPssblePrst_m;

    //pdu session type
    PDUSessionType *pduSessType_mp; //for decode
    PDUSessionType pduSessType_m; //for encode

    //security indication
    SecurityIndication *secInd_mp;
    SecurityIndication secInd_m;
    bool isSecIndPrst_m;
    IeExtenIeType MaxIntProDataRateDlIe;
    MaximumIntegrityProtectedDataRate maxIntProDataRate_m;

    //network instance
    NetworkInstance *networkInstance_mp;
    NetworkInstance networkInstance_m;
    bool isnetworkInstancePrst_m;

    //qos flow setup request list
    QosFlowSetupRequestList qosFlowSetupReqList_m;
    QosFlowSetupRequestList* qosFlowSetupReqList_mp;
    struct QosFlowSetupRequestList_ qosFlowSetupReqList_ma[maxnoofQoSFlows_c];
    int numofQosFlowSetupReq_m;

    //Common Network Instance
    CommonNetworkInstance commonNetworkInstance_m;
    CmmnNtwrkInstance cmmnNtwrkInstance_ma;
    CommonNetworkInstance* commonNetworkInstance_mp;
    bool commonntkInstancePresent_m;
};

#endif
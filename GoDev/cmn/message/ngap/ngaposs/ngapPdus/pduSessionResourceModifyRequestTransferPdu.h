#ifndef __pduSessionResourceModifyRequestTransferPdu_h__INCLUDED__
#define __pduSessionResourceModifyRequestTransferPdu_h__INCLUDED__

#include "ngapOssPdu.h"
#include <ngapMsgTypes.h>
///////////////////////////////////////////////////////////////////////////////
// @class PDUSessionResourceModifyRequestTransferPdu
//-- **************************************************************
//--
//-- PDU Session Resource Modify Request Transfer 
//--
//-- **************************************************************
//

// PDUSessionResourceModifyRequestTransfer ::= SEQUENCE {
//        protocolIEs             ProtocolIE-Container            { {PDUSessionResourceModifyRequestTransferIEs} },
//        ...
// }
//
// PDUSessionResourceModifyRequestTransferIEs NGAP-PROTOCOL-IES ::= {
//        { ID id-PDUSessionAggregateMaximumBitRate       CRITICALITY reject      TYPE PDUSessionAggregateMaximumBitRate          PRESENCE optional               }|
//        { ID id-UL-NGU-UP-TNLModifyList                 CRITICALITY reject      TYPE UL-NGU-UP-TNLModifyList                    PRESENCE optional               }|
//        { ID id-NetworkInstance                         CRITICALITY reject      TYPE NetworkInstance                            PRESENCE optional               }|
//        { ID id-QosFlowAddOrModifyRequestList           CRITICALITY reject      TYPE QosFlowAddOrModifyRequestList              PRESENCE optional               }|
//        { ID id-QosFlowToReleaseList                    CRITICALITY reject      TYPE QosFlowListWithCause                       PRESENCE optional               }|
//        { ID id-AdditionalUL-NGU-UP-TNLInformation      CRITICALITY reject      TYPE UPTransportLayerInformationList            PRESENCE optional               }|
//        { ID id-CommonNetworkInstance                   CRITICALITY ignore      TYPE CommonNetworkInstance                      PRESENCE optional               },
//        ...
//}
///////////////////////////////////////////////////////////////////////////////
class PduSesResModifyReqTransferPdu : public NgapOssPdu
{
  public:
    PduSesResModifyReqTransferPdu();
    //virtual
    ~PduSesResModifyReqTransferPdu();

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
    /// @brief get PDUSessionResourceModifyRequestTransfer pointer
    /// returns the pointer to the
    /// PDUSessionResourceModifyRequestTransfer in the decoded message
    /////////////////////////////////////////////////////////////
    PDUSessionResourceModifyRequestTransfer *getPduSessResModifyReqTransferDecodedPtr() const;

    /////////////////////////////////////////////////////////////
    /// PDUSessionAggregateMaximumBitRate
    /////////////////////////////////////////////////////////////
    PDUSessionAggregateMaximumBitRate *getSessAmbrDecodedPtr() const;
    PDUSessionAggregateMaximumBitRate *getSessAmbrEncodePtr();
    bool isSessAmbrPrst();

    /////////////////////////////////////////////////////////////
    ///  ul ngu up tnl modify list
    /////////////////////////////////////////////////////////////
    UL_NGU_UP_TNLModifyList* getUlNguUpTnlModifyListDecodedPtr() const;
    UL_NGU_UP_TNLModifyList* getUlNguUpTnlModifyListEncodePtr();
    UL_NGU_UP_TNLModifyItem* getUlNguUpTnlModifyItemDecodedPtr(unsigned int index) const;
    UL_NGU_UP_TNLModifyItem* getUlNguUpTnlModifyItemEncodePtr(unsigned int index);
    void setNumofUlNguUpTnlModifyListItem(unsigned int num);
    unsigned int getNumofUlNguUpTnlModifyListItem() const;

    /////////////////////////////////////////////////////////////
    /// NetworkInstance
    /////////////////////////////////////////////////////////////
    NetworkInstance *getNetworkInstanceDecodedPtr() const;
    NetworkInstance *getNetworkInstanceEncodePtr();
    bool isNetworkInstancePrst();

    /////////////////////////////////////////////////////////////
    /// Additional Up Transport Layer Information List
    /////////////////////////////////////////////////////////////
    UPTransportLayerInformationList *getAddUpTransportLayerInfoListDecodedPtr() const;
    UPTransportLayerInformationList *getAddUpTransportLayerInfoListEncodePtr();

    UPTransportLayerInformationItem *getAddUpTransportLayerInfoItemDecodedPtr(unsigned int index) const;
    UPTransportLayerInformationItem *getAddUpTransportLayerInfoItemEncodePtr(unsigned int index);
    void setNumofAddUpTransportLayerInfoListItem(unsigned int num);
    unsigned int getNumOfAddUpTransportLayerInfoListItem() const;
    bool isAddUpTransportLayerInfoListPrst();

    /////////////////////////////////////////////////////////////
    /// QosFlowAddOrModifyRequestList
    /////////////////////////////////////////////////////////////
    QosFlowAddOrModifyRequestList* getQosFlowAddOrModReqListDecodedPtr() const;
    QosFlowAddOrModifyRequestList* getQosFlowAddOrModReqListEncodePtr();
  
    QosFlowAddOrModifyRequestItem* getQosFlowAddOrModReqItemDecodedPtr(unsigned int index) const;
    QosFlowAddOrModifyRequestItem* getQosFlowAddOrModReqItemEncodePtr(unsigned int index);
  
    void setNumofQosFlowAddOrModReqListItem(unsigned int num);
    unsigned int getNumofQosFlowAddOrModReqListItem() const;   

    /////////////////////////////////////////////////////////////
    /// QosFlowToReleaseList
    /////////////////////////////////////////////////////////////
    QosFlowListWithCause* getQosFlowToRelListDecodedPtr() const;
    QosFlowListWithCause* getQosFlowToRelListEncodePtr();
  
    QosFlowWithCauseItem* getQosFlowToRelItemDecodedPtr(unsigned int index) const;
    QosFlowWithCauseItem* getQosFlowToRelItemEncodePtr(unsigned int index);
  
    void setNumofQosFlowToRelListItem(unsigned int num);
    unsigned int getNumofQosFlowToRelListItem() const;

    /////////////////////////////////////////////////////////////
    /// CommonNetworkInstance
    /////////////////////////////////////////////////////////////
    CommonNetworkInstance * getCommonNetworkInstanceDecodedPtr() const;
    CommonNetworkInstance * getCommonNetworkInstanceEncodePtr();
    bool isCommonNetworkInstancePrst(); 

  private:
    typedef PduSessResModifyReqTransferIe ProtocolIeType;

    //OSS structures
    PDUSessionResourceModifyRequestTransfer* pduSessResModifyReqTransfer_mp;

    //Session ambr
    PDUSessionAggregateMaximumBitRate *sessAMBR_mp; //for decode
    PDUSessionAggregateMaximumBitRate sessAMBR_m; //for encode
    bool isSessAmbrPrst_m;

    // ul ngu up tnl modify list
    UL_NGU_UP_TNLModifyList  ulNguUpTnlModifyList_m;
    UL_NGU_UP_TNLModifyList* ulNguUpTnlModifyList_mp;
    struct UL_NGU_UP_TNLModifyList_ ulNguUpTnlModifyList_ma[maxnoofMultiConnectivity];
    TransLayerAddr ulTransLayerAddress_m[maxnoofMultiConnectivity];
    TransLayerAddr dlTransLayerAddress_m[maxnoofMultiConnectivity];
    int numofUlNguUpTnlModifyList_m;

    //addtional up transport information list
    UPTransportLayerInformationList  addUpTransportLayerInfoList_m;
    UPTransportLayerInformationList* addUpTransportLayerInfoList_mp;
    bool isAddUpTransportLayerInfoListPrst_m;
    struct UPTransportLayerInformationList_ addTransLayerAddrList_ma[maxnoofMultiConnectivityMinusOne_c];
    TransLayerAddr addTransLayerAddress_m[maxnoofMultiConnectivity];
    int numofAddUpTransportLayerInfo_m;

    //network instance
    NetworkInstance *networkInstance_mp;
    NetworkInstance networkInstance_m;
    bool isNetworkInstancePrst_m;

    //qos flow add or modify request list  QosFlowAddOrModifyRequestList
    QosFlowAddOrModifyRequestList  qosFlowAddOrModifyReqList_m;
    QosFlowAddOrModifyRequestList* qosFlowAddOrModifyReqList_mp;
    struct QosFlowAddOrModifyRequestList_ qosFlowAddOrModifyReqList_ma[maxnoofQoSFlows_c];
    int numofQosFlowAddOrModifyReq_m;

    //qos flow to release list
    QosFlowListWithCause qosFlowToReleaseList_m;
    QosFlowListWithCause* qosFlowToReleaseList_mp;
    struct QosFlowListWithCause_ qosFlowToReleaseList_ma[maxnoofQoSFlows_c];
    int numofQosFlowToRelease_m;

    //Common Network Instance
    CommonNetworkInstance commonNetworkInstance_m;
    CmmnNtwrkInstance cmmnNtwrkInstance_ma;
    CommonNetworkInstance* commonNetworkInstance_mp;
    bool commonntkInstancePresent_m;
};

#endif

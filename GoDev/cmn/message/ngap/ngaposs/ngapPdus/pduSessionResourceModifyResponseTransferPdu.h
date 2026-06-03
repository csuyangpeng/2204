#ifndef __pduSessionResourceModifyResponseTransferPdu_h__INCLUDED__
#define __pduSessionResourceModifyResponseTransferPdu_h__INCLUDED__

#include "ngapOssPdu.h"
#include <ngapMsgTypes.h>
///////////////////////////////////////////////////////////////////////////////
// @class PDUSessionResourceModifyResponseTransferPdu
//-- **************************************************************
//--
//-- PDU Session Resource Modify Response Transfer 
//--
//-- **************************************************************
//

// PDUSessionResourceModifyResponseTransfer ::= SEQUENCE {
//        protocolIEs             ProtocolIE-Container            { {PDUSessionResourceModifyResponseTransferIEs} },
//        ...
// }
//
// PDUSessionResourceModifyResponseTransfer ::= SEQUENCE {
//   dL-NGU-UP-TNLInformation               UPTransportLayerInformation                                         OPTIONAL,
//   uL-NGU-UP-TNLInformation               UPTransportLayerInformation                                         OPTIONAL,
//   qosFlowAddOrModifyResponseList         QosFlowAddOrModifyResponseList                                      OPTIONAL,
//   additionalDLQosFlowPerTNLInformation   QosFlowPerTNLInformationList                                        OPTIONAL,
//   qosFlowFailedToAddOrModifyList         QosFlowListWithCause                                                OPTIONAL,
//   iE-Extensions      ProtocolExtensionContainer { {PDUSessionResourceModifyResponseTransfer-ExtIEs} }    OPTIONAL,
//   ...
// }
//
///////////////////////////////////////////////////////////////////////////////
class PduSesResModifyRespTransferPdu : public NgapOssPdu
{
  public:
    PduSesResModifyRespTransferPdu();
    //virtual
    ~PduSesResModifyRespTransferPdu();

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
    /// @brief get PduSesResModifyResponseTransferPdu pointer
    /// returns the pointer to the
    /// PduSesResModifyResponseTransferPdu in the decoded message
    /////////////////////////////////////////////////////////////
    PDUSessionResourceModifyResponseTransfer *getPduSessResModifyRespTransferDecodedPtr() const;

    /////////////////////////////////////////////////////////////
    /// DL UPTransportLayerInformation
    /////////////////////////////////////////////////////////////
    UPTransportLayerInformation *getDlUpTNLInformationDecodedPtr() const;
    UPTransportLayerInformation *getDlUpTNLInformationEncodePtr();
    bool isDlUpTNLInformationPrst();

    /////////////////////////////////////////////////////////////
    /// UL UPTransportLayerInformation
    /////////////////////////////////////////////////////////////
    UPTransportLayerInformation *getUlUpTNLInformationDecodedPtr() const;
    UPTransportLayerInformation *getUlUpTNLInformationEncodePtr();
    bool isUlUpTNLInformationPrst();

    /////////////////////////////////////////////////////////////
    /// QosFlowAddOrModifyResponseList
    /////////////////////////////////////////////////////////////
    QosFlowAddOrModifyResponseList* getQosFlowAddOrModRespListDecodedPtr() const;
    QosFlowAddOrModifyResponseList* getQosFlowAddOrModRespListEncodePtr();
  
    QosFlowAddOrModifyResponseItem* getQosFlowAddOrModRespItemDecodedPtr(unsigned int index) const;
    QosFlowAddOrModifyResponseItem* getQosFlowAddOrModRespItemEncodePtr(unsigned int index);
  
    void setNumofQosFlowAddOrModRespListItem(unsigned int num);
    unsigned int getNumofQosFlowAddOrModRespListItem() const; 
    bool isQosFlowAddOrModifyRespListPrst();

    /////////////////////////////////////////////////////////////
    /// Qos Flow per TNL Information List
    /////////////////////////////////////////////////////////////
    QosFlowPerTNLInformationList *getQosFlowPerTNLInformationListDecodedPtr() const;
    QosFlowPerTNLInformationList *getQosFlowPerTNLInformationListEncodePtr();

    QosFlowPerTNLInformationItem* getQosFlowPerTNLInformationItemDecodedPtr(unsigned int index) const;
    QosFlowPerTNLInformationItem* getQosFlowPerTNLInformationItemEncodePtr(unsigned int index);

    void setNumofQosFlowPerTNLInformationListItem(unsigned int num);
    unsigned int getNumofQosFlowPerTNLInformationListItem() const;

    bool isQosFlowPerTNLInformationListPrst();

    AssociatedQosFlowItem* getAddAssQosFlowItemDecodedPtr(QosFlowPerTNLInformationItem * InfoPtr,
                                                                         unsigned int indAss) const;
    AssociatedQosFlowItem* getAddAssQosFlowItemEncodePtr(unsigned int indInfo,  
                                                                        unsigned int index);
    void setNumofAddAssQosFlowListItem(unsigned int num,unsigned int index);
    unsigned int getNumofAddAssQosFlowListItem(unsigned int index) const;

    /////////////////////////////////////////////////////////////
    /// QosFlowToReleaseList
    /////////////////////////////////////////////////////////////
    QosFlowListWithCause* getQosFlowToRelListDecodedPtr() const;
    QosFlowListWithCause* getQosFlowToRelListEncodePtr();
  
    QosFlowWithCauseItem* getQosFlowToRelItemDecodedPtr(unsigned int index) const;
    QosFlowWithCauseItem* getQosFlowToRelItemEncodePtr(unsigned int index);
  
    void setNumofQosFlowToRelListItem(unsigned int num);
    unsigned int getNumofQosFlowToRelListItem() const;
    bool isQosFlowToRelListPrst();

  private:

    //OSS structures
    PDUSessionResourceModifyResponseTransfer* pduSessResModifyRespTransfer_mp;

    //dl up transport layer information
    UPTransportLayerInformation *dlUpTransportLayerInfo_mp; //for decode
    UPTransportLayerInformation dlUpTransportLayerInfo_m;   //for encode
    bool isDlUpTransportLayerInfoPrst_m;

    //ul up transport layer information
    UPTransportLayerInformation *ulUpTransportLayerInfo_mp; //for decode
    UPTransportLayerInformation ulUpTransportLayerInfo_m;   //for encode
    bool isUlUpTransportLayerInfoPrst_m;

    EnhanceTransLayerAddr transLayerAddress_m;

    //additional dl qos flow per tul information
    QosFlowPerTNLInformationList *qosFlowPerTNLInformationList_mp;
    QosFlowPerTNLInformationList qosFlowPerTNLInformationList_m;
    struct QosFlowPerTNLInformationList_ qosFlowPerTNLInformationList_ma[maxnoofMultiConnectivityMinusOne_c]; 
    unsigned int numOfQosFlowPerTNLInformationListItem_m;
    bool isAddqosFlowPerTnlInfoPrst_m;

    //AssQosFlowList
    unsigned int numofAddQosFlowInd_m[maxnoofMultiConnectivityMinusOne_c];
    struct AssociatedQosFlowList_ addAssQosFlowList_ma[maxnoofMultiConnectivityMinusOne_c][maxnoofQoSFlows_c];

    EnhanceTransLayerAddr assTransLayerAddress_m[maxnoofMultiConnectivityMinusOne_c];

    //qos flow add or modify response list  QosFlowAddOrModifyResponseList
    QosFlowAddOrModifyResponseList  qosFlowAddOrModifyRespList_m;
    QosFlowAddOrModifyResponseList* qosFlowAddOrModifyRespList_mp;
    struct QosFlowAddOrModifyResponseList_ qosFlowAddOrModifyRespList_ma[maxnoofQoSFlows_c];
    int numofQosFlowAddOrModifyResp_m;
    bool isQosFlowAddOrModifyResponsePrst_m;

    //qos flow to modify list   QosFlowFailedToAddOrModifyList
    QosFlowListWithCause qosFlowToReleaseList_m;
    QosFlowListWithCause* qosFlowToReleaseList_mp;
    struct QosFlowListWithCause_ qosFlowToReleaseList_ma[maxnoofQoSFlows_c];
    int numofQosFlowToRelease_m;
    bool isQosFlowListPrst_m;
    
};

#endif

#ifndef __pduSessionResourceSetupResponseTransferPdu_h__INCLUDED__
#define __pduSessionResourceSetupResponseTransferPdu_h__INCLUDED__

#include "ngapOssPdu.h"
#include <ngapMsgTypes.h>
///////////////////////////////////////////////////////////////////////////////
// @class PDUSessionResourceSetupResponseTransferPdu
//-- **************************************************************
//--
//-- PDU Session Resource Setup Response Transfer 
//--
//-- **************************************************************
//
//
//PDUSessionResourceSetupResponseTransfer ::= SEQUENCE {
//  qosFlowPerTNLInformation                QosFlowPerTNLInformation,
//  additionalQosFlowPerTNLInformation      QosFlowPerTNLInformation                                            OPTIONAL,
//  securityResult                          SecurityResult                                                      OPTIONAL,
//  qosFlowFailedToSetupList                QosFlowList                                                         OPTIONAL,
//  iE-Extensions       ProtocolExtensionContainer { {PDUSessionResourceSetupResponseTransfer-ExtIEs} }     OPTIONAL,
//  ...
//}
//}
///////////////////////////////////////////////////////////////////////////////

class PduSesResSetupRespTransferPdu : public NgapOssPdu
{
  public:
    PduSesResSetupRespTransferPdu();
    //virtual
    ~PduSesResSetupRespTransferPdu();

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
    /// @brief get PDUSessionResourceSetupResponseTransfer pointer
    /// returns the pointer to the
    /// PDUSessionResourceSetupResponseTransfer in the decoded message
    /////////////////////////////////////////////////////////////
    PDUSessionResourceSetupResponseTransfer *
    getPduSessResSetupRespTransferDecodedPtr() const;

    /////////////////////////////////////////////////////////////
    /// QosFlowPerTNLInformation
    /////////////////////////////////////////////////////////////
    QosFlowPerTNLInformation *getQosFlowInfoDecodedPtr() const;
    QosFlowPerTNLInformation *getQosFlowInfoEncodePtr();
    AssociatedQosFlowItem* getAssQosFlowItemDecodedPtr(unsigned int index) const;
    AssociatedQosFlowItem* getAssQosFlowItemEncodePtr(unsigned int index);
    void setNumofAssQosFlowListItem_m_v(unsigned int num);
    unsigned int getNumofAssQosFlowListItem() const;    

    /////////////////////////////////////////////////////////////
    /// Additional DL QosFlowPerTNLInformation List
    /////////////////////////////////////////////////////////////
    QosFlowPerTNLInformationList* getQosFlowPerTNLInformationListDecodedPtr() const;
    QosFlowPerTNLInformationList* getQosFlowPerTNLInformationListEncodePtr();

    QosFlowPerTNLInformationItem* getQosFlowPerTNLInformationListItemDecodedPtr(unsigned int index) const;
    QosFlowPerTNLInformationItem* getQosFlowPerTNLInformationListItemEncodePtr(unsigned int index);

    void setNumOfQosFlowPerTNLInformationListItem_v(unsigned int num);
    unsigned int getNumOfQosFlowPerTNLInformationListItem() const;

    bool isQosFlowPerTNLInformationListPrst();

    AssociatedQosFlowItem* getAddAssQosFlowItemDecodedPtr(QosFlowPerTNLInformationItem *addQosFlowInfoPtr ,
                                                                         unsigned int indAssQosFlow) const;
    AssociatedQosFlowItem* getAddAssQosFlowItemEncodePtr(unsigned int index,
                                                                        unsigned int indAssQosFlow);
    void setNumofAddAssQosFlowListItem_m_v(unsigned int num,unsigned int index);
    unsigned int getNumofAddAssQosFlowListItem(unsigned int index) const;
    
    /////////////////////////////////////////////////////////////
    /// SecurityResult
    /////////////////////////////////////////////////////////////
    SecurityResult *getSecurityResultDecodedPtr() const;
    SecurityResult *getSecurityResultEncodePtr();
    bool isSecResultPrst();

    /////////////////////////////////////////////////////////////
    /// qos flow failed to setup list
    /////////////////////////////////////////////////////////////
    QosFlowListWithCause* getQosFlowFailedSetupListDecodedPtr() const;
    QosFlowListWithCause* getQosFlowFailedSetupListEncodePtr();
  
    QosFlowWithCauseItem* getQosFlowItemDecodedPtr(unsigned int index) const;
    QosFlowWithCauseItem* getQosFlowItemEncodePtr(unsigned int index);
  
    void setNumofQosFlowListItem_m_v(unsigned int num);
    unsigned int getNumofQosFlowListItem() const;
    bool isQosFlowFailedSetupListPrst();

  private:

    //OSS structures
    PDUSessionResourceSetupResponseTransfer* pduSessResSetupRespTransfer_mp;

    //qos flow per tul information
    QosFlowPerTNLInformation *qosFlowPerTnlInfo_mp; //for decode
    QosFlowPerTNLInformation  qosFlowPerTnlInfo_m;   //for encode
    struct AssociatedQosFlowList_ assQosFlowList_ma[maxnoofQoSFlows_c];
    int numofQosFlowInd_m;
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

    //security result 
    SecurityResult *secResult_mp; //for decode
    SecurityResult secResult_m; //for encode
    bool isSecResultPrst_m;

    //qos flow failed to setup list
    QosFlowListWithCause qosFlowFailedSetupList_m;
    QosFlowListWithCause* qosFlowFailedSetupList_mp;
    struct QosFlowListWithCause_ qosFlowList_ma[maxnoofQoSFlows_c];
    int numofQosFlowlist_m;
    bool isQosFlowFailedListPrst_m;
    
};

#endif


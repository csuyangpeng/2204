#ifndef __pduSessionResourceModifyRequestPdu_h__INCLUDED__
#define __pduSessionResourceModifyRequestPdu_h__INCLUDED__

#include "ngapOssPdu.h"
#include <ngapMsgTypes.h>
///////////////////////////////////////////////////////////////////////////////
//-- **************************************************************
//--
//-- PDU Session Resource Modify Elementary Procedure
//--
//-- **************************************************************
//
//-- **************************************************************
//--
//-- PDU SESSION RESOURCE MODIFY REQUEST
//--
//-- **************************************************************
//
//PDUSessionResourceModifyRequest ::= SEQUENCE {
//      protocolIEs             ProtocolIE-Container            { {PDUSessionResourceModifyRequestIEs} },
//      ...
//}
//
//PDUSessionResourceModifyRequestIEs NGAP-PROTOCOL-IES ::= {
//      { ID id-AMF-UE-NGAP-ID                                                  CRITICALITY reject      TYPE AMF-UE-NGAP-ID                                                             PRESENCE mandatory      }|
//      { ID id-RAN-UE-NGAP-ID                                                  CRITICALITY reject      TYPE RAN-UE-NGAP-ID                                                             PRESENCE mandatory      }|
//      { ID id-RANPagingPriority                                               CRITICALITY ignore      TYPE RANPagingPriority                                                  PRESENCE optional               }|
//      { ID id-PDUSessionResourceModifyListModReq              CRITICALITY reject      TYPE PDUSessionResourceModifyListModReq PRESENCE mandatory      },
//      ...
//}
//PDUSessionResourceModifyListModReq ::= SEQUENCE (SIZE(1..maxnoofPDUSessions)) OF PDUSessionResourceModifyItemModReq
//PDUSessionResourceModifyItemModReq ::= SEQUENCE {
//        pDUSessionID                                                            PDUSessionID,
//        nAS-PDU                                                                         NAS-PDU                                                                                                 OPTIONAL,
//        pDUSessionResourceModifyRequestTransfer         OCTET STRING (CONTAINING PDUSessionResourceModifyRequestTransfer),
//        iE-Extensions           ProtocolExtensionContainer { {PDUSessionResourceModifyItemModReq-ExtIEs} }      OPTIONAL,
//        ...
//}
///////////////////////////////////////////////////////////////////////////////

class PduSessResModReqPdu:public NgapOssPdu
{
  public:
    PduSessResModReqPdu();
    //virtual
    ~PduSessResModReqPdu();

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
    unsigned int encodeOssToPer(NgapOssCtxt* ctxt);

    /////////////////////////////////////////////////////////////
    /// @brief get Decoded Oss Data Pointers
    ///  processes the decoded OSS buffer and populates
    ///  the data access pointers in the class for the IEs present in the
    ///  decoded OSS buffer
    /////////////////////////////////////////////////////////////
    //virtual
    void getDecodedOssDataPtrs_v();

    /////////////////////////////////////////////////////////////
    /// @brief get PDUSessionResourceModifyRequestPdu pointer
    /// returns the pointer to the
    /// PDUSessionResourceModifyRequestPdu in the decoded message
    /////////////////////////////////////////////////////////////
    PDUSessionResourceModifyRequest * getPduSessResModReqDecodedPtr() const;

    /////////////////////////////////////////////////////////////
    /// RAN_UE_NGAP_ID
    /////////////////////////////////////////////////////////////
    RAN_UE_NGAP_ID* getRanUeNgapIdDecodedPtr() const;
    RAN_UE_NGAP_ID* getRanUeNgapIdEncodePtr();

    /////////////////////////////////////////////////////////////
    /// AMF_UE_NGAP_ID
    /////////////////////////////////////////////////////////////
    AMF_UE_NGAP_ID* getAmfUeNgapIdDecodedPtr() const;
    AMF_UE_NGAP_ID* getAmfUeNgapIdEncodePtr();

    /////////////////////////////////////////////////////////////
    /// RANPagingPriority
    /////////////////////////////////////////////////////////////
    RANPagingPriority *getRANPagingPriorityDecodedPtr() const;
    RANPagingPriority *getRANPagingPriorityEncodePtr();
    bool isRANPagingPriorityPresent();

    /////////////////////////////////////////////////////////////
    /// PDUSessionResourceModifyRequest
    /////////////////////////////////////////////////////////////
    PDUSessionResourceModifyListModReq* getPduSessResModifyListModReqDecodedPtr() const;
    PDUSessionResourceModifyListModReq* getPduSessResModifyListModReqEncodePtr();

    PDUSessionResourceModifyItemModReq* getPduSessResModifyItemModReqDecodedPtr(unsigned int index) const;
    PDUSessionResourceModifyItemModReq* getPduSessResModifyItemModReqEncodePtr(unsigned int index);

    void setNumofPduSessResModifyItemModReq_v(unsigned int num);
    unsigned int getNumofPduSessResModifyItemModReq() const;

    /////////////////////////////////////////////////////////////
    /// Ie_Extension
    /////////////////////////////////////////////////////////////
    SnssaiIeValue* getIeExtenIeTypeDecodedPtr();
    S_NSSAI* getSnssaiEncodePtr();
    unsigned char* getSnssaiStr();

  private:
    typedef PduSessResMdfyReqIe ProtocolIeType;
    typedef SnssaiIeValue IeExtenIeType;

    //OSS structures
    PDUSessionResourceModifyRequest *pduSessResMdfyReq_mp; //for encode

    //AMF-UE-NGAP-ID
    AMF_UE_NGAP_ID amfUeNgapId_m;
    AMF_UE_NGAP_ID* amfUeNgapId_mp;

    //RAN-UE-NGAP-ID
    RAN_UE_NGAP_ID ranUeNgapId_m;
    RAN_UE_NGAP_ID* ranUeNgapId_mp;

    //RANPagingPriority
    RANPagingPriority ranPagingPriority_m;
    RANPagingPriority *ranPagingPriority_mp;
    bool ranPagingPriorityPresent_m;

    //PDUSessionResourceModifyListModReq
    PDUSessionResourceModifyListModReq pduSessResModifyListModReq_m;
    struct PDUSessionResourceModifyListModReq_ pduSessResModifyListModReq_ma[maxnoofPDUSessions_c];
    unsigned int numofPduSessResModifyItemModReq_m;
    PDUSessionResourceModifyListModReq* pduSessResModifyListModReq_mp;

    NasPdu pduNasPdus_ma[maxnoofPDUSessions_c];
    SmfTransferPdu transfers_ma[maxnoofPDUSessions_c];
    SnssaiStr snssaiStr_ma[maxnoofPDUSessions_c];
    S_NSSAI snssai_m;
    //Ie_Extension
    IeExtenIeType snssaiIe;
	//IeExtenIeType snssaiIes[maxnoofPDUSessions_c];	
};

#endif

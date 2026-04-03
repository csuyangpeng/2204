#ifndef __pduSessionResourceReleaseResponsePdu_h__INCLUDED__
#define __pduSessionResourceReleaseResponsePdu_h__INCLUDED__

#include "ngapOssPdu.h"
#include <ngapMsgTypes.h>
///////////////////////////////////////////////////////////////////////////////
// @class PDUSessionResourceReleaseResponsePdu
// -- **************************************************************
// --
// -- PDU SESSION RESOURCE RELEASE RESPONSE
// --
// -- **************************************************************

// PDUSessionResourceReleaseResponse ::= SEQUENCE {
//  protocolIEs     ProtocolIE-Container        { {PDUSessionResourceReleaseResponseIEs} },
//  ...
// }

// PDUSessionResourceReleaseResponseIEs NGAP-PROTOCOL-IES ::= {
//  { ID id-AMF-UE-NGAP-ID                              CRITICALITY ignore  TYPE AMF-UE-NGAP-ID                                     PRESENCE mandatory  }|
//  { ID id-RAN-UE-NGAP-ID                              CRITICALITY ignore  TYPE RAN-UE-NGAP-ID                                     PRESENCE mandatory  }|
//  { ID id-PDUSessionResourceReleasedListRelRes            CRITICALITY ignore  TYPE PDUSessionResourceReleasedListRelRes       PRESENCE mandatory  }|
//  { ID id-UserLocationInformation                     CRITICALITY ignore  TYPE UserLocationInformation                        PRESENCE optional       }|
//  { ID id-CriticalityDiagnostics                      CRITICALITY ignore  TYPE CriticalityDiagnostics                         PRESENCE optional       },
//  ...
// }
///////////////////////////////////////////////////////////////////////////////

class PduSessResRelRespPdu:public NgapOssPdu
{
  public:
    PduSessResRelRespPdu();
    //virtual
    ~PduSessResRelRespPdu();
    
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
    /// @brief get PDUSessionResourceReleaseResponse pointer
    /// returns the pointer to the
    /// PDUSessionResourceReleaseResponse in the decoded message
    /////////////////////////////////////////////////////////////
    PDUSessionResourceReleaseResponse * getPduSessResRelRespDecodedPtr() const;

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
    /// UserLocationInformation
    /////////////////////////////////////////////////////////////
    UserLocationInformation* getUseLocationInformationDecodedPtr() const;
    UserLocationInformation* getUseLocationInformationEncodePtr();
    bool isUseLocationInformationPresent();

    NGRAN_CGI* getNgRanCgiEncodePtr();
    UserLocationInfomationNRIeExtension* getIeExtenIeTypeDecodedPtr();

    /////////////////////////////////////////////////////////////
    /// PDUSessionResourceReleasedListRelRes
    /////////////////////////////////////////////////////////////
    PDUSessionResourceReleasedListRelRes* getPduSessResListRelRespDecodedPtr() const;
    PDUSessionResourceReleasedListRelRes* getPduSessResListRelRespEncodePtr();
    
    PDUSessionResourceReleasedItemRelRes* getPduSessResItemRelRespDecodedPtr(unsigned int index) const;
    PDUSessionResourceReleasedItemRelRes* getPduSessResItemRelRespEncodePtr(unsigned int index);

    void setNumofPduSessResListItemRelResp_v(unsigned int num);
    unsigned int getNumofPduSessResListItemRelResp() const;

  private:
    typedef PduSessResRelRespIe ProtocolIeType;
    typedef UserLocationInfomationNRIeExtension IeExtenIeType;
  
    //OSS structures
    PDUSessionResourceReleaseResponse *pduSessResRelResp_mp; //for encode

    //AMF-UE-NGAP-ID
    AMF_UE_NGAP_ID amfUeNgapId_m;
    AMF_UE_NGAP_ID* amfUeNgapId_mp;

    //RAN-UE-NGAP-ID
    RAN_UE_NGAP_ID ranUeNgapId_m;
    RAN_UE_NGAP_ID* ranUeNgapId_mp;

    //User Location Information
    UserLocationInformation useLocationInformation_m;
    UserLocationInformation* useLocationInformation_mp;
    unsigned char nrCellId_ma[maxNrCellIdLength_c];
    bool userLocInfoPresent_m;
    //Ie_Extension
    IeExtenIeType userLocalInfoIe;
    NGRAN_CGI ngranCgi_m;
    unsigned char nrCellId_mb[maxNrCellIdLength_c];

    //PDUSessionResourceReleasedListRelRes
    PDUSessionResourceReleasedListRelRes pduSessResListRelRes_m;
    struct PDUSessionResourceReleasedListRelRes_ pduSessResListRelResp_ma[maxnoofPDUSessions_c];
    unsigned int numofPduSessResListItemCxtReq_m;
    PDUSessionResourceReleasedListRelRes* pduSessResListRelRes_mp;

    SmfTransferPdu transfers_ma[maxnoofPDUSessions_c];
};

#endif
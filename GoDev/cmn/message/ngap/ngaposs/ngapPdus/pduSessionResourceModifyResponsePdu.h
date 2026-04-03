#ifndef __pduSessionResourceModifyResponsePdu_h__INCLUDED__
#define __pduSessionResourceModifyResponsePdu_h__INCLUDED__

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
//-- PDU SESSION RESOURCE MODIFY RESPONSE
//--
//-- **************************************************************
//
//PDUSessionResourceModifyResponse ::= SEQUENCE {
//  protocolIEs     ProtocolIE-Container        { {PDUSessionResourceModifyResponseIEs} },
//  ...
//}
//
//PDUSessionResourceModifyResponseIEs NGAP-PROTOCOL-IES ::= {
//  { ID id-AMF-UE-NGAP-ID                                  CRITICALITY ignore  TYPE AMF-UE-NGAP-ID                                     PRESENCE mandatory  }|
//  { ID id-RAN-UE-NGAP-ID                                  CRITICALITY ignore  TYPE RAN-UE-NGAP-ID                                     PRESENCE mandatory  }|
//  { ID id-PDUSessionResourceModifyListModRes              CRITICALITY ignore  TYPE PDUSessionResourceModifyListModRes         PRESENCE optional       }|
//  { ID id-PDUSessionResourceFailedToModifyListModRes      CRITICALITY ignore  TYPE PDUSessionResourceFailedToModifyListModRes     PRESENCE optional   }|
//  { ID id-UserLocationInformation                         CRITICALITY ignore  TYPE UserLocationInformation                            PRESENCE optional       }|
//  { ID id-CriticalityDiagnostics                          CRITICALITY ignore  TYPE CriticalityDiagnostics                         PRESENCE optional       },
//  ...
//}
//
///////////////////////////////////////////////////////////////////////////////

class PduSessResMdfyRespPdu:public NgapOssPdu
{
  public:
    PduSessResMdfyRespPdu();
    //virtual
    ~PduSessResMdfyRespPdu();

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
    /// @brief get PDUSessionResourceModifyResponsePdu pointer
    /// returns the pointer to the
    /// PDUSessionResourceModifyResponsePdu in the decoded message
    /////////////////////////////////////////////////////////////
    PDUSessionResourceModifyResponse * getPduSessResMdfyRespDecodedPtr() const;

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
    /// PDUSessionResourceModifyListModRes
    /////////////////////////////////////////////////////////////
    PDUSessionResourceModifyListModRes* getPduSessResModifyListModResDecodedPtr() const;
    PDUSessionResourceModifyListModRes* getPduSessResModifyListModResEncodePtr();

    PDUSessionResourceModifyItemModRes* getPduSessResModifyItemModResDecodedPtr(unsigned int index) const;
    PDUSessionResourceModifyItemModRes* getPduSessResModifyItemModResEncodePtr(unsigned int index);

    void setNumofPduSessResModifyItemModRes_v(unsigned int num);
    unsigned int getNumofPduSessResModifyItemModRes() const;
    bool isPduSessionResourceModifyListPresent();
    
    /////////////////////////////////////////////////////////////
    /// PDUSessionResourceFailedToModifyListModRes
    /////////////////////////////////////////////////////////////
    PDUSessionResourceFailedToModifyListModRes* getPduSessResFailedModifyListModResDecodedPtr() const;
    PDUSessionResourceFailedToModifyListModRes* getPduSessResFailedModifyListModResEncodePtr();
    
    PDUSessionResourceFailedToModifyItemModRes* getPduSessResFailedModifyItemModResDecodedPtr(unsigned int index) const;
    PDUSessionResourceFailedToModifyItemModRes* getPduSessResFailedModifyItemModResEncodePtr(unsigned int index);
    
    void setNumofFailedToModifyItem_v(unsigned int num);
    unsigned int getNumofFailedToModifyItem() const;
    bool isPduSessionResourceFailModifyListPresent();

    /////////////////////////////////////////////////////////////
    /// User Location Information
    /////////////////////////////////////////////////////////////
    UserLocationInformation* getUseLocationInformationDecodedPtr() const;
    UserLocationInformation* getUseLocationInformationEncodePtr();
    bool isUseLocationInformationPresent();

    NGRAN_CGI* getNgRanCgiEncodePtr();
    UserLocationInfomationNRIeExtension* getIeExtenIeTypeDecodedPtr();

  private:
    typedef PduSessResMdfyRespIe ProtocolIeType;
    typedef UserLocationInfomationNRIeExtension IeExtenIeType;

    //OSS structures
    PDUSessionResourceModifyResponse *PduSessResMdfyResp_mp; //for encode

    //AMF-UE-NGAP-ID
    AMF_UE_NGAP_ID amfUeNgapId_m;
    AMF_UE_NGAP_ID* amfUeNgapId_mp;

    //RAN-UE-NGAP-ID
    RAN_UE_NGAP_ID ranUeNgapId_m;
    RAN_UE_NGAP_ID* ranUeNgapId_mp;

    //PDUSessionResourceModifyListModResp
    PDUSessionResourceModifyListModRes pduSessResModifyListModResp_m;
    struct PDUSessionResourceModifyListModRes_ pduSessResModifyListModResp_ma[maxnoofPDUSessions_c];
    unsigned int numofPduSessResModifyItemModResp_m;
    PDUSessionResourceModifyListModRes* pduSessResModifyListModResp_mp;
    SmfTransferPdu respTransfers_ma[maxnoofPDUSessions_c];
    bool pduSessResModifyRespListPresent_m;

    //Failed list
    PDUSessionResourceFailedToModifyListModRes pduSessResFailedModifyList_m;
    struct PDUSessionResourceFailedToModifyListModRes_ pduSessResFailedModifyList_ma[maxnoofPDUSessions_c];
    unsigned int numofpduSessResFailedModifyList_m;
    PDUSessionResourceFailedToModifyListModRes* pduSessResFailedModifyList_mp;
    SmfTransferPdu failedTransfers_ma[maxnoofPDUSessions_c];
    bool pduSessResFailModifyRespListPresent_m;

    //User Location Information
    UserLocationInformation useLocationInformation_m; 
    UserLocationInformation* useLocationInformation_mp; 
    unsigned char nrCellId_ma[maxNrCellIdLength_c];
    bool userLocInfoPresent_m;
    //Ie_Extension
    IeExtenIeType userLocalInfoIe;
    NGRAN_CGI ngranCgi_m;
    unsigned char nrCellId_mb[maxNrCellIdLength_c];
};

#endif

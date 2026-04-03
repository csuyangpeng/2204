#ifndef __ueContextReleaseCompletePdu_h__INCLUDED__
#define __ueContextReleaseCompletePdu_h__INCLUDED__

#include "ngapOssPdu.h"
#include <ngapMsgTypes.h>
///////////////////////////////////////////////////////////////////////////////
//-- **************************************************************
//--
//-- UE CONTEXT RELEASE COMPLETE
//--
//-- **************************************************************
//
//UEContextReleaseComplete ::= SEQUENCE {
//  protocolIEs     ProtocolIE-Container        { {UEContextReleaseComplete-IEs} },
//  ...
//}
//
//UEContextReleaseComplete-IEs NGAP-PROTOCOL-IES ::= {
//  { ID id-AMF-UE-NGAP-ID                              CRITICALITY ignore  TYPE AMF-UE-NGAP-ID                                         PRESENCE mandatory  }|
//  { ID id-RAN-UE-NGAP-ID                              CRITICALITY ignore  TYPE RAN-UE-NGAP-ID                                         PRESENCE mandatory  }|
//  { ID id-UserLocationInformation                     CRITICALITY ignore  TYPE UserLocationInformation                            PRESENCE optional       }|
//  { ID id-InfoOnRecommendedCellsAndRANNodesForPaging  CRITICALITY ignore  TYPE InfoOnRecommendedCellsAndRANNodesForPaging PRESENCE optional       }|
//  { ID id-PDUSessionResourceListCxtRelCpl             CRITICALITY reject  TYPE PDUSessionResourceListCxtRelCpl                    PRESENCE optional       }|
//  { ID id-CriticalityDiagnostics                      CRITICALITY ignore  TYPE CriticalityDiagnostics                             PRESENCE optional       },
//  ...
//}
///////////////////////////////////////////////////////////////////////////////

class UeContextReleaseCompletePdu:public NgapOssPdu
{
  public:
    UeContextReleaseCompletePdu();
    //virtual
    ~UeContextReleaseCompletePdu();
    
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
    /// @brief get UeContextReleaseComplete pointer
    /// returns the pointer to the
    /// UeContextReleaseComplete in the decoded message
    /////////////////////////////////////////////////////////////
    UEContextReleaseComplete * getUeContextRelCplDecodedPtr() const;

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
    bool isUserLocInfoPrst();
    
    NGRAN_CGI* getNgRanCgiEncodePtr();
    UserLocationInfomationNRIeExtension* getIeExtenIeTypeDecodedPtr();
    
    /////////////////////////////////////////////////////////////
    /// InfoOnRecommendedCellsAndRANNodesForPaging
    /////////////////////////////////////////////////////////////
    InfoOnRecommendedCellsAndRANNodesForPaging* getInfoOnRecommendedCellsAndRANNodesForPagingDecodedPtr() const;
    InfoOnRecommendedCellsAndRANNodesForPaging* getInfoOnRecommendedCellsAndRANNodesForPagingEncodePtr();

    RecommendedCellItem * getRecommendedCellItemDecodedPtr(unsigned int index) const;
    RecommendedCellItem * getRecommendedCellItemEncodePtr(unsigned int index);

    void setNumOfRecommendedCellItem_v(unsigned int num);
    unsigned int getNumOfRecommendedCellItem() const;

    RecommendedRANNodeItem * getRecommendedRANNodeItemDecodedPtr(unsigned int index) const;
    RecommendedRANNodeItem * getRecommendedRANNodeItemEncodePtr(unsigned int index);

    void setNumOfRecommendedRANNodeItem_v(unsigned int num);
    unsigned int getNumOfRecommendedRANNodeItem() const;

    bool isRecommendInfoPrst();

    /////////////////////////////////////////////////////////////
    /// PDUSessionResourceListCxtRelCpl
    /////////////////////////////////////////////////////////////
    PDUSessionResourceListCxtRelCpl* getPduSessResListCxtRelCplDecodedPtr() const;
    PDUSessionResourceListCxtRelCpl* getPduSessResListCxtRelCplEncodePtr();
    
    PDUSessionResourceItemCxtRelCpl* getPduSessResItemCxtRelCplDecodedPtr(unsigned int index) const;
    PDUSessionResourceItemCxtRelCpl* getPduSessResItemCxtRelCplEncodePtr(unsigned int index);

    void setNumofPduSessResListItemCxtRelCpl_v(unsigned int num);
    unsigned int getNumofPduSessResListItemCxtRelCpl() const;

//  PDUSessionResourceItemCxtRelCpl_ExtIEs_EXTENSION* getPduSessResItemCxtRelCplExtIeEncodePtr();
//  PduSessResItemRelCplIeExtension* getIeExtenIeTypeDecodedPtr();

  private:
    typedef UEContextReleaseCompleteIe ProtocolIeType;
    typedef UserLocationInfomationNRIeExtension IeExtenIeType;
  
    //OSS structures
    UEContextReleaseComplete *UeContextReleaseComplete_mp; //for encode
//    typedef PduSessResItemRelCplIeExtension IeExtenIeType;

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
    bool isUliPrst_m;
    //Ie_Extension
    IeExtenIeType userLocalInfoIe;
    NGRAN_CGI ngranCgi_m;
    unsigned char nrCellId_mb[maxNrCellIdLength_c];

    //InfoOnRecommendedCellsAndRANNodesForPaging
    InfoOnRecommendedCellsAndRANNodesForPaging infoOnRecommendedCellsAndRANNodesForPaging_m;
    InfoOnRecommendedCellsAndRANNodesForPaging* infoOnRecommendedCellsAndRANNodesForPaging_mp;

    struct RecommendedCellList_ recommendedCellList_ma[maxnoofRecommendedCells_c];
    unsigned char nrCellmaId[maxnoofRecommendedCells_c];
    unsigned int numOfRecommendedCellItem_m;

    struct RecommendedRANNodeList_ recommendedRANNodeList_ma[maxnoofRecommendedRANNodes_c];
    unsigned char gnbIdValue_ma[maxGnbIdValue_c];
    unsigned int numOfRecommendedRANNodeItem_m;
    bool isRecommendInfoPrst_m;

    //PDUSessionResourceListCxtRelCpl
    PDUSessionResourceListCxtRelCpl pduSessResListCxtRelCpl_m;
    struct PDUSessionResourceListCxtRelCpl_ pduSessResListCxtRelCpl_ma[maxnoofPDUSessions_c];
    unsigned int numofPduSessResListItemCxtRelCpl_m;
    PDUSessionResourceListCxtRelCpl* pduSessResListCxtRelCpl_mp;
    //Ie_Extension
    IeExtenIeType pduSessResListCxtRelCplIe;
    PDUSessionResourceItemCxtRelCpl_ExtIEs_EXTENSION pduSessResListCxtRelCplExtIe_m;
};

#endif
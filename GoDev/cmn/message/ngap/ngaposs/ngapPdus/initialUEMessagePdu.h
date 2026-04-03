#ifndef __ngApOssInitialUEMessagePdu_h__INCLUDED__
#define __ngApOssInitialUEMessagePdu_h__INCLUDED__

#include "ngapOssPdu.h"
#include <ngapMsgTypes.h>
///////////////////////////////////////////////////////////////////////////////
// @class InitialUEMessagePdu
//-- **************************************************************
//--
//-- INITIAL UE MESSAGE
//--
//-- **************************************************************
//
//InitialUEMessage ::= SEQUENCE {
//  protocolIEs     ProtocolIE-Container        { {InitialUEMessage-IEs} },
//  ...
//}
//
//InitialUEMessage-IEs NGAP-PROTOCOL-IES ::= {
//  { ID id-RAN-UE-NGAP-ID              CRITICALITY reject  TYPE RAN-UE-NGAP-ID             PRESENCE mandatory  }|
//  { ID id-NAS-PDU                     CRITICALITY reject  TYPE NAS-PDU                    PRESENCE mandatory  }|
//  { ID id-UserLocationInformation     CRITICALITY reject  TYPE UserLocationInformation    PRESENCE mandatory  }|
//  { ID id-RRCEstablishmentCause       CRITICALITY ignore  TYPE RRCEstablishmentCause      PRESENCE mandatory      }|
//  { ID id-FiveG-S-TMSI                CRITICALITY reject  TYPE FiveG-S-TMSI               PRESENCE optional       }|
//  { ID id-AMFSetID                    CRITICALITY ignore  TYPE AMFSetID                   PRESENCE optional       }|
//  { ID id-UEContextRequest            CRITICALITY ignore  TYPE UEContextRequest           PRESENCE optional       }|
//  { ID id-AllowedNSSAI                CRITICALITY reject  TYPE AllowedNSSAI               PRESENCE optional       },
//  ...
//}
///////////////////////////////////////////////////////////////////////////////

class InitialUEMessagePdu:public NgapOssPdu
{
  public:
    InitialUEMessagePdu();
    //virtual
    ~InitialUEMessagePdu();
    
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
    /// @brief get InitialUEMessage pointer
    /// returns the pointer to the
    /// InitialUEMessage in the decoded message
    /////////////////////////////////////////////////////////////
    InitialUEMessage * getInitialUEMessageDecodedPtr() const;

    /////////////////////////////////////////////////////////////
    /// RAN_UE_NGAP_ID
    /////////////////////////////////////////////////////////////
    RAN_UE_NGAP_ID* getRanUeNgapIdDecodedPtr() const;
    RAN_UE_NGAP_ID* getRanUeNgapIdEncodePtr();

    /////////////////////////////////////////////////////////////
    /// NAS_PDU
    /////////////////////////////////////////////////////////////
    NAS_PDU* getNasPduDecodedPtr() const;
    NAS_PDU* getNasPduEncodePtr();

    /////////////////////////////////////////////////////////////
    /// UserLocationInformation
    /////////////////////////////////////////////////////////////
    UserLocationInformation* getUseLocationInformationDecodedPtr() const;
    UserLocationInformation* getUseLocationInformationEncodePtr();

    NGRAN_CGI* getNgRanCgiEncodePtr();
    UserLocationInfomationNRIeExtension* getIeExtenIeTypeDecodedPtr();

    /////////////////////////////////////////////////////////////
    /// FiveG_S_TMSI
    /////////////////////////////////////////////////////////////
    FiveG_S_TMSI* get5gStmsiDecodedPtr() const;
    FiveG_S_TMSI* get5gStmsiEncodePtr();
    bool is5gStmsiPresent();

    /////////////////////////////////////////////////////////////
    /// RRCEstablishmentCause
    /////////////////////////////////////////////////////////////
    RRCEstablishmentCause* getRrcEstablishmentCauseDecodedPtr() const;
    RRCEstablishmentCause* getRrcEstablishmentCauseEncodePtr();

    /////////////////////////////////////////////////////////////
    /// AMFSetID
    /////////////////////////////////////////////////////////////
    AMFSetID* getAmfSetIdDecodedPtr() const;
    AMFSetID* getAmfSetIdEncodePtr();
    bool isAmfSetIdPresent();

    /////////////////////////////////////////////////////////////
    /// UEContextRequest
    /////////////////////////////////////////////////////////////
    UEContextRequest* getUeContextRequestDecodedPtr() const;
    UEContextRequest* getUeContextRequestEncodePtr();
    bool isUeContextRequestPresent();
    
    /////////////////////////////////////////////////////////////
    /// Allowed NSSAI
    /////////////////////////////////////////////////////////////
    AllowedNSSAI* getAllowedNssaiDecodedPtr() const;
    AllowedNSSAI* getAllowedNssaiEncodePtr();
    bool isAllowedNssaiPresent();

    AllowedNSSAI_Item* getAllowedNssaiItemDecodedPtr(unsigned int index) const;
    AllowedNSSAI_Item* getAllowedNssaiItemEncodePtr(unsigned int index);

    void setNumofAllowedNssaiItem_v(unsigned int num);
    unsigned int getNumofAllowedNssaiItem() const;

  private:
    typedef InitialUEMessageIe ProtocolIeType;
    typedef UserLocationInfomationNRIeExtension IeExtenIeType;
  
    //OSS structures
    InitialUEMessage *InitialUEMessage_mp; //for encode

    //RAN UE NGAP ID
    RAN_UE_NGAP_ID  ranUeNgapId_m;
    RAN_UE_NGAP_ID* ranUeNgapId_mp;
    
    //NAS PDU
    NAS_PDU nasPdu_m;
    unsigned char nasPdu_ma[maxNasPduValue_c];
    NAS_PDU* nasPdu_mp;

    //User Location Information
    UserLocationInformation useLocationInformation_m; 
    UserLocationInformation* useLocationInformation_mp; 
    unsigned char nrCellId_ma[maxNrCellIdLength_c];
    //Ie_Extension
    IeExtenIeType userLocalInfoIe;
    NGRAN_CGI ngranCgi_m;
    unsigned char nrCellId_mb[maxNrCellIdLength_c];

    //FiveG S TMSI
    FiveG_S_TMSI stmsi_m;
    Stmsi5G stmsi5g_ma; 
    FiveG_S_TMSI* stmsi_mp;
    bool stmsiPresent_m;

    //RRCE Stablishment Cause
    RRCEstablishmentCause rrcEstablishmentCause_m;
    RRCEstablishmentCause* rrcEstablishmentCause_mp;

    //AMF Set ID
    AMFSetID amfSetId_m;
    unsigned char amfSetIdValue_ma[maxAmfSetIdValue_c];
    AMFSetID* amfSetId_mp;
    bool amfSetIdPresent_m;

    //UE Context Request
    UEContextRequest ueContextRequest_m;
    UEContextRequest* ueContextRequest_mp;
    bool ueContextRequestPresent_m;

    //Allowed NSSAI
    AllowedNSSAI  allowedNssai_m;
    struct AllowedNSSAI_ allowedNssai_ma[maxnoofAllowedSnssaiValue_c];
    unsigned int numOfAllowedNassiItem_m;
    AllowedNSSAI* allowedNssai_mp;
    bool allowedNssaiPresent_m;
};

#endif // __ngApOssInitialUEMessagePdu_h__INCLUDED__

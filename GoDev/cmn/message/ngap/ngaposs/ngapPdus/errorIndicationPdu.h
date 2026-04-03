#ifndef __errorIndicationPdu_h__INCLUDED__
#define __errorIndicationPdu_h__INCLUDED__

#include "ngapOssPdu.h"
#include <ngapMsgTypes.h>
///////////////////////////////////////////////////////////////////////////////
// @class ErrorIndicationPdu
// -- **************************************************************
// --
// -- Error Indication Elementary Procedure
// --
// -- **************************************************************

// -- **************************************************************
// --
// -- ERROR INDICATION
// --
// -- **************************************************************

// ErrorIndication ::= SEQUENCE {
//  protocolIEs     ProtocolIE-Container        { {ErrorIndicationIEs} },
//  ...
// }

// ErrorIndicationIEs NGAP-PROTOCOL-IES ::= {
//  { ID id-AMF-UE-NGAP-ID              CRITICALITY ignore  TYPE AMF-UE-NGAP-ID             PRESENCE optional       }|
//  { ID id-RAN-UE-NGAP-ID              CRITICALITY ignore  TYPE RAN-UE-NGAP-ID             PRESENCE optional       }|
//  { ID id-Cause                       CRITICALITY ignore  TYPE Cause                      PRESENCE optional       }|
//  { ID id-CriticalityDiagnostics      CRITICALITY ignore  TYPE CriticalityDiagnostics     PRESENCE optional       },
//  ...
// }
///////////////////////////////////////////////////////////////////////////////

class ErrorIndicationPdu : public NgapOssPdu
{
  public:
    ErrorIndicationPdu();
    //virtual
    ~ErrorIndicationPdu();
    
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
    /// @brief get ErrorIndication pointer
    /// returns the pointer to the
    /// ErrorIndication in the decoded message
    /////////////////////////////////////////////////////////////
    ErrorIndication * getErrorIndicationDecodedPtr() const;

    /////////////////////////////////////////////////////////////
    /// RAN_UE_NGAP_ID
    /////////////////////////////////////////////////////////////
    RAN_UE_NGAP_ID* getRanUeNgapIdDecodedPtr() const;
    RAN_UE_NGAP_ID* getRanUeNgapIdEncodePtr();
    bool isRanUeNgapIdPresent();

    /////////////////////////////////////////////////////////////
    /// AMF_UE_NGAP_ID
    /////////////////////////////////////////////////////////////
    AMF_UE_NGAP_ID* getAmfUeNgapIdDecodedPtr() const;
    AMF_UE_NGAP_ID* getAmfUeNgapIdEncodePtr();
    bool isAmfUeNgapIdPresent();

    /////////////////////////////////////////////////////////////
    /// Cause
    /////////////////////////////////////////////////////////////
    Cause* getCauseDecodedPtr() const;
    Cause* getCauseEncodePtr();
    bool isCausePresent();

  private:
    typedef ErrorIndicationIe ProtocolIeType;
  
    //OSS structures
    ErrorIndication *errorIndication_mp; //for encode

    //AMF-UE-NGAP-ID
    AMF_UE_NGAP_ID amfUeNgapId_m;
    AMF_UE_NGAP_ID* amfUeNgapId_mp;
    bool amfUeNgapIdPresent_m;

    //RAN-UE-NGAP-ID
    RAN_UE_NGAP_ID ranUeNgapId_m;
    RAN_UE_NGAP_ID* ranUeNgapId_mp;
    bool ranUeNgapIdPresent_m;

    //Cause
    Cause cause_m;
    Cause* cause_mp;
    bool causePresent_m;
};

#endif 
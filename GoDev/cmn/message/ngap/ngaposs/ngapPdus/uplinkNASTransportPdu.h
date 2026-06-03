#ifndef __uplinkNasTransportPdu_h__INCLUDED__
#define __uplinkNasTransportPdu_h__INCLUDED__

#include "ngapOssPdu.h"
#include <ngapMsgTypes.h>
///////////////////////////////////////////////////////////////////////////////
// @class UplinkNasTransportPdu
//-- **************************************************************
//--
//-- UPLINK NAS TRANSPORT
//--
//-- **************************************************************
//
//UplinkNASTransport ::= SEQUENCE {
//  protocolIEs     ProtocolIE-Container    { {UplinkNASTransport-IEs} },
//  ...
//}
//
//UplinkNASTransport-IEs NGAP-PROTOCOL-IES ::= {
//  { ID id-AMF-UE-NGAP-ID              CRITICALITY reject  TYPE AMF-UE-NGAP-ID                 PRESENCE mandatory  }|
//  { ID id-RAN-UE-NGAP-ID              CRITICALITY reject  TYPE RAN-UE-NGAP-ID                 PRESENCE mandatory  }|
//  { ID id-NAS-PDU                     CRITICALITY reject  TYPE NAS-PDU                        PRESENCE mandatory  }|
//  { ID id-UserLocationInformation     CRITICALITY ignore  TYPE UserLocationInformation        PRESENCE mandatory  },
//  ...
//}
///////////////////////////////////////////////////////////////////////////////

class UplinkNASTransportPdu:public NgapOssPdu
{
  public:
    UplinkNASTransportPdu();
    //virtual
    ~UplinkNASTransportPdu();
    
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
    /// @brief get UplinkNASTransport pointer
    /// returns the pointer to the
    /// UplinkNASTransport in the decoded message
    /////////////////////////////////////////////////////////////
    UplinkNASTransport * getUplinkNASTransportDecodedPtr() const;

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

  private:
    typedef UplinkNASTransportIe ProtocolIeType;
    typedef UserLocationInfomationNRIeExtension IeExtenIeType;

    //OSS structures
    UplinkNASTransport *UplinkNASTransport_mp; //for encode

    //AMF-UE-NGAP-ID
    AMF_UE_NGAP_ID amfUeNgapId_m;
    AMF_UE_NGAP_ID* amfUeNgapId_mp;

    //RAN-UE-NGAP-ID
    RAN_UE_NGAP_ID ranUeNgapId_m;
    RAN_UE_NGAP_ID* ranUeNgapId_mp;

    //NAS-PDU
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
};

#endif // __ngApOssUplinkNasTransportPdu_h__INCLUDED__
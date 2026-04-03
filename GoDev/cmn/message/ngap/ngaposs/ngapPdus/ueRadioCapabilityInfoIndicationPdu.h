#ifndef __ueRadioCapabilityInfoIndicationPdu_h__INCLUDED__
#define __ueRadioCapabilityInfoIndicationPdu_h__INCLUDED__

#include "ngapOssPdu.h"
#include <ngapMsgTypes.h>
///////////////////////////////////////////////////////////////////////////////
// @class UeRadioCapaInfoIndicationPdu
// -- **************************************************************
// --
// -- UE RADIO CAPABILITY INFO INDICATION
// --
// -- **************************************************************

// UERadioCapabilityInfoIndication ::= SEQUENCE {
// 	protocolIEs			ProtocolIE-Container		{ {UERadioCapabilityInfoIndicationIEs} },
// 	...
// }

// UERadioCapabilityInfoIndicationIEs NGAP-PROTOCOL-IES ::= {
// 	{ ID id-AMF-UE-NGAP-ID					CRITICALITY reject	TYPE AMF-UE-NGAP-ID					PRESENCE mandatory	}|
// 	{ ID id-RAN-UE-NGAP-ID					CRITICALITY reject	TYPE RAN-UE-NGAP-ID					PRESENCE mandatory	}|
// 	{ ID id-UERadioCapability				CRITICALITY ignore	TYPE UERadioCapability				PRESENCE mandatory	}|
// 	{ ID id-UERadioCapabilityForPaging		CRITICALITY ignore	TYPE UERadioCapabilityForPaging		PRESENCE optional		},
// 	...
// }
///////////////////////////////////////////////////////////////////////////////

class UeRadioCapaInfoIndicationPdu:public NgapOssPdu
{
  public:
    UeRadioCapaInfoIndicationPdu();
    //virtual
    ~UeRadioCapaInfoIndicationPdu();
    
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
    /// @brief get UeRadioCapabilityInfoIndication pointer
    /// returns the pointer to the
    /// UeRadioCapabilityInfoIndication in the decoded message
    /////////////////////////////////////////////////////////////
    UERadioCapabilityInfoIndication * getUeRadioCapaInfoIndicationDecodedPtr() const;

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
    /// UERadioCapability
    /////////////////////////////////////////////////////////////
    UERadioCapability* getUeRadioCapabilityDecodedPtr() const;
    UERadioCapability* getUeRadioCapabilityEncodePtr();

    /////////////////////////////////////////////////////////////
    /// UERadioCapabilityForPaging
    /////////////////////////////////////////////////////////////
    UERadioCapabilityForPaging* getUeRadioCapabilityForPagingDecodedPtr() const;
    UERadioCapabilityForPaging* getUeRadioCapabilityForPagingEncodePtr();
    bool isUeRadioCapabilityForPagingPresent();

  private:
    typedef UeRadioCapaInfoIndicationIe ProtocolIeType;
  
    //OSS structures
    UERadioCapabilityInfoIndication *ueRadioCapaInfoIndication_mp; //for encode

    //AMF-UE-NGAP-ID
    AMF_UE_NGAP_ID amfUeNgapId_m;
    AMF_UE_NGAP_ID* amfUeNgapId_mp;

    //RAN-UE-NGAP-ID
    RAN_UE_NGAP_ID ranUeNgapId_m;
    RAN_UE_NGAP_ID* ranUeNgapId_mp;

    //UERadioCapability
    UERadioCapability ueRadioCapability_m;
    unsigned char ueRadioCapability_ma[maxUeRadioCapabilityValue_c];
    UERadioCapability* ueRadioCapability_mp;

    //UERadioCapabilityForPaging
    UERadioCapabilityForPaging ueRadioCapabilityForPaging_m;
    unsigned char ueRadioCapabilityForPagingNr_ma[maxUeRadioCapabilityValue_c];
	unsigned char ueRadioCapabilityForPagingEutra_ma[maxUeRadioCapabilityValue_c];
    UERadioCapabilityForPaging* ueRadioCapabilityForPaging_mp;
    bool  ueRadioCapabilityForPagingPresent_m;
};

#endif

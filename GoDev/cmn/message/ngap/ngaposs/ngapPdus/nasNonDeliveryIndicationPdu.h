#ifndef __ngApOssNasNonDeliveryIndicationPdu_h__INCLUDED__
#define __ngApOssNasNonDeliveryIndicationPdu_h__INCLUDED__

#include "ngapOssPdu.h"
#include <ngapMsgTypes.h>
///////////////////////////////////////////////////////////////////////////////
// @class NasNonDeliveryIndicationPdu
//-- **************************************************************
//--
//-- NAS NON DELIVERY INDICATION
//--
//-- **************************************************************
//
//NASNonDeliveryIndication ::= SEQUENCE {
//	protocolIEs		ProtocolIE-Container		{ {NASNonDeliveryIndication-IEs} },
//	...
//}
//
//NASNonDeliveryIndication-IEs NGAP-PROTOCOL-IES ::= {
//	{ ID id-AMF-UE-NGAP-ID		CRITICALITY reject	TYPE AMF-UE-NGAP-ID		PRESENCE mandatory	}|
//	{ ID id-RAN-UE-NGAP-ID		CRITICALITY reject	TYPE RAN-UE-NGAP-ID		PRESENCE mandatory	}|
//	{ ID id-NAS-PDU				CRITICALITY ignore	TYPE NAS-PDU			PRESENCE mandatory	}|
//	{ ID id-Cause				CRITICALITY ignore	TYPE Cause				PRESENCE mandatory	},
//	...
//}
//
///////////////////////////////////////////////////////////////////////////////

class NasNonDeliveryIndicationPdu:public NgapOssPdu
{
  public:
    NasNonDeliveryIndicationPdu();
    //virtual
    ~NasNonDeliveryIndicationPdu();
    
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
    /// @brief get Nas Non Delivery Indication pointer
    /// returns the pointer to the
    /// NasNonDeliveryIndication in the decoded message
    /////////////////////////////////////////////////////////////
    NASNonDeliveryIndication *getNasNonDeliveryIndicationDecodedPtr() const;

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
    /// Cause
    /////////////////////////////////////////////////////////////
    Cause* getCauseDecodedPtr() const;
    Cause* getCauseEncodePtr();

  private:
    typedef NasNonDeliveryIndicationIe ProtocolIeType;

    //OSS structures
    NASNonDeliveryIndication *NasNonDeliveryIndication_mp; //for encode

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
    
    //Cause
	Cause cause_m;
	Cause* cause_mp;

};

#endif // __ngApOssUplinkNasTransportPdu_h__INCLUDED__

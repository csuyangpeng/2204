#ifndef __downlinkNasTransportPdu_h__INCLUDED__
#define __downlinkNasTransportPdu_h__INCLUDED__

#include "ngapOssPdu.h"
#include <ngapMsgTypes.h>
///////////////////////////////////////////////////////////////////////////////
// @class DownlinkNasTransportPdu
//-- **************************************************************
//--
//-- DOWNLINK NAS TRANSPORT
//--
//-- **************************************************************
//
//DownlinkNASTransport ::= SEQUENCE {
//	protocolIEs		ProtocolIE-Container		{ {DownlinkNASTransport-IEs} },
//	...
//}
//
//DownlinkNASTransport-IEs NGAP-PROTOCOL-IES ::= {
//	{ ID id-AMF-UE-NGAP-ID					CRITICALITY reject	TYPE AMF-UE-NGAP-ID					PRESENCE mandatory	}|
//	{ ID id-RAN-UE-NGAP-ID					CRITICALITY reject	TYPE RAN-UE-NGAP-ID					PRESENCE mandatory	}|
//	{ ID id-OldAMF							CRITICALITY reject	TYPE AMFName						PRESENCE optional		}|
//	{ ID id-RANPagingPriority				CRITICALITY ignore	TYPE RANPagingPriority				PRESENCE optional		}|
//	{ ID id-NAS-PDU							CRITICALITY reject	TYPE NAS-PDU						PRESENCE mandatory	}|
//	{ ID id-MobilityRestrictionList			CRITICALITY ignore	TYPE MobilityRestrictionList		PRESENCE optional		}|
//	{ ID id-IndexToRFSP						CRITICALITY ignore	TYPE IndexToRFSP					PRESENCE optional		}|
//	{ ID id-UEAggregateMaximumBitRate		CRITICALITY ignore	TYPE UEAggregateMaximumBitRate		PRESENCE optional		}|
//	{ ID id-AllowedNSSAI					CRITICALITY reject	TYPE AllowedNSSAI					PRESENCE optional		},
//	...
//}
///////////////////////////////////////////////////////////////////////////////

class DownlinkNASTransportPdu : public NgapOssPdu
{
  public:
    DownlinkNASTransportPdu();
    //virtual
    ~DownlinkNASTransportPdu();
    
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
    /// @brief get DownlinkNasTransport pointer
    /// returns the pointer to the
    /// DownlinkNasTransport in the decoded message
    /////////////////////////////////////////////////////////////
    DownlinkNASTransport * getDownlinkNasTransportDecodedPtr() const;

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
    /// UE Aggregate Maximum Bit Rate
    /////////////////////////////////////////////////////////////
    UEAggregateMaximumBitRate* getUeAmbrDecodedPtr() const;
    UEAggregateMaximumBitRate* getUeAmbrEncodePtr();
	bool isUeAmbrPresent();

    /////////////////////////////////////////////////////////////
    /// Index To RFSP
    /////////////////////////////////////////////////////////////
    IndexToRFSP* getIndexToRfspDecodedPtr() const;
    IndexToRFSP* getIndexToRfspEncodePtr();
	bool isIndexToRfspPresent();

    /////////////////////////////////////////////////////////////
    /// RANPagingPriority
    /////////////////////////////////////////////////////////////
    RANPagingPriority* getRanPagingPriorityDecodedPtr() const;
    RANPagingPriority* getRanPagingPriorityEncodePtr();
	bool isRanPagingPriorityPresent();

    /////////////////////////////////////////////////////////////
    /// AmfName 
    /////////////////////////////////////////////////////////////
    AMFName* getOldAmfNameDecodedPtr() const;
    AMFName* getOldAmfNameEncodePtr();
	bool isOldAmfNamePresent();
    
    /////////////////////////////////////////////////////////////
    /// Allowed NSSAI
    /////////////////////////////////////////////////////////////
    AllowedNSSAI* getAllowedNssaiDecodedPtr() const;
    AllowedNSSAI* getAllowedNssaiEncodePtr();

    AllowedNSSAI_Item* getAllowedNssaiItemDecodedPtr(unsigned int index) const;
    AllowedNSSAI_Item* getAllowedNssaiItemEncodePtr(unsigned int index);

    void setNumofAllowedNssaiItem_v(unsigned int num);
    unsigned int getNumofAllowedNssaiItem() const;
	bool isAllowedNssaiPresent();

  private:
  	
    typedef DownlinkNASTransportIe ProtocolIeType;
  
    //OSS structures
    DownlinkNASTransport *downlinkNasTransport_mp; //for encode

    //AMF-UE-NGAP-ID
    AMF_UE_NGAP_ID amfUeNgapId_m;
    AMF_UE_NGAP_ID* amfUeNgapId_mp;

    //RAN-UE-NGAP-ID
    RAN_UE_NGAP_ID ranUeNgapId_m;
    RAN_UE_NGAP_ID* ranUeNgapId_mp;

    //OldAMF
    AMFName oldAmfName_m;  //for encode
    char oldAmfNameValue_ma[maxAMFNameValue_c];
    AMFName *oldAmfName_mp;  //for decode
    bool oldAmfNamePresent_m;

    //RANPagingPriority
    RANPagingPriority   ranPagingPriority_m;
    RANPagingPriority*  ranPagingPriority_mp;
	bool ranPagingPriorityPresent_m;
	
    //NAS-PDU
    NAS_PDU nasPdu_m;
	unsigned char nasPdu_ma[maxNasPduValue_c];
    NAS_PDU* nasPdu_mp;

    //Index To RFSP
    IndexToRFSP  indexToRfsp_m; 
    IndexToRFSP* indexToRfsp_mp;
	bool indexToRfspPresent_m;

    //UE Aggregate Maximum Bit Rate
    UEAggregateMaximumBitRate     ueAmbr_m; 
    UEAggregateMaximumBitRate*    ueAmbr_mp;
	bool ueAmbrPresent_m;

    //Allowed NSSAI
    AllowedNSSAI  allowedNssai_m;
    struct AllowedNSSAI_ allowedNssai_ma[maxnoofAllowedSnssaiValue_c];
    unsigned int numOfAllowedNassiItem_m;
    AllowedNSSAI* allowedNssai_mp;
	bool allowedNssaiPresent_m;
};

#endif // __ngApOssDownlinkNasTransportPdu_h__INCLUDED__

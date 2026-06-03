#ifndef __pduSessionResourceSetupRequestPdu_h__INCLUDED__
#define __pduSessionResourceSetupRequestPdu_h__INCLUDED__

#include "ngapOssPdu.h"
#include <ngapMsgTypes.h>
///////////////////////////////////////////////////////////////////////////////
// @class PDUSessionResourceSetupRequestPdu
//-- **************************************************************
//--
//-- PDU Session Resource Setup Request
//--
//-- **************************************************************
//
//
//PDUSessionResourceSetupRequest ::= SEQUENCE {
//	protocolIEs		ProtocolIE-Container		{ {PDUSessionResourceSetupRequestIEs} },
//	...
//}
//
//PDUSessionResourceSetupRequestIEs NGAP-PROTOCOL-IES ::= {
//	{ ID id-AMF-UE-NGAP-ID							CRITICALITY reject	TYPE AMF-UE-NGAP-ID								PRESENCE mandatory	}|
//	{ ID id-RAN-UE-NGAP-ID							CRITICALITY reject	TYPE RAN-UE-NGAP-ID								PRESENCE mandatory	}|
//	{ ID id-RANPagingPriority						CRITICALITY ignore	TYPE RANPagingPriority							PRESENCE optional		}|
//	{ ID id-NAS-PDU									CRITICALITY reject	TYPE NAS-PDU									PRESENCE optional		}|
//	{ ID id-PDUSessionResourceSetupListSUReq		CRITICALITY reject	TYPE PDUSessionResourceSetupListSUReq		PRESENCE mandatory	},
//	...
//}
///////////////////////////////////////////////////////////////////////////////

class PDUSessionResourceSetupRequestPdu : public NgapOssPdu
{
  public:
    PDUSessionResourceSetupRequestPdu();
    //virtual
    ~PDUSessionResourceSetupRequestPdu();

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
    unsigned int encodeOssToPer(NgapOssCtxt *ctxt);

    //virtual
	//unsigned int decodePerToOss(NgapOssCtxt * ctxt);

    /////////////////////////////////////////////////////////////
    /// @brief get Decoded Oss Data Pointers
    ///  processes the decoded OSS buffer and populates
    ///  the data access pointers in the class for the IEs present in the
    ///  decoded OSS buffer
    /////////////////////////////////////////////////////////////
    //virtual
    void getDecodedOssDataPtrs_v();

    /////////////////////////////////////////////////////////////
    /// @brief get PDUSessionResourceSetupRequest pointer
    /// returns the pointer to the
    /// PDUSessionResourceSetupRequest in the decoded message
    /////////////////////////////////////////////////////////////
    PDUSessionResourceSetupRequest *getPduSessResSetupReqDecodedPtr() const;

    /////////////////////////////////////////////////////////////
    /// RAN_UE_NGAP_ID
    /////////////////////////////////////////////////////////////
    RAN_UE_NGAP_ID *getRanUeNgapIdDecodedPtr() const;
    RAN_UE_NGAP_ID *getRanUeNgapIdEncodePtr();

    /////////////////////////////////////////////////////////////
    /// AMF_UE_NGAP_ID
    /////////////////////////////////////////////////////////////
    AMF_UE_NGAP_ID *getAmfUeNgapIdDecodedPtr() const;
    AMF_UE_NGAP_ID *getAmfUeNgapIdEncodePtr();

    /////////////////////////////////////////////////////////////
    /// NAS_PDU
    /////////////////////////////////////////////////////////////
    NAS_PDU *getNasPduDecodedPtr() const;
    NAS_PDU *getNasPduEncodePtr();
    bool isNasPduPresent();

    /////////////////////////////////////////////////////////////
    /// UE Aggregate Maximum Bit Rate
    /////////////////////////////////////////////////////////////
    UEAggregateMaximumBitRate* getUeAggregateMaximumBitRateDecodedPtr() const;
    UEAggregateMaximumBitRate* getUeAggregateMaximumBitRateEncodePtr();
    bool isUeAMBRPresent();

    /////////////////////////////////////////////////////////////
    /// RANPagingPriority
    /////////////////////////////////////////////////////////////
    RANPagingPriority *getRanPagePriorityDecodedPtr() const;
    RANPagingPriority *getRanPagePriorityEncodePtr();
    bool isRanPagePriorityPrst();

    /////////////////////////////////////////////////////////////
    /// PDUSessionResourceSetupListSUReq
    /////////////////////////////////////////////////////////////
    PDUSessionResourceSetupListSUReq *getPduSessResSetupListSUReqDecodedPtr() const;
    PDUSessionResourceSetupListSUReq *getPduSessResSetupListSUReqEncodePtr();

    PDUSessionResourceSetupItemSUReq *getPduSessResSetupSUReqItemDecodedPtr(unsigned int index) const;
    PDUSessionResourceSetupItemSUReq *getPduSessResSetupSUReqItemEncodePtr(unsigned int index);

    void setNumofPduSessResSetupListSUReqItem_v(unsigned int num);
    unsigned int getNumofPduSessResSetupListSUReqItem() const;

  private:
  	
    typedef PDUSessResSetupReqIe ProtocolIeType;

    //OSS structures
    PDUSessionResourceSetupRequest *pduSessResSetupRequest_mp; //for encode

    //AMF-UE-NGAP-ID
    AMF_UE_NGAP_ID amfUeNgapId_m;
    AMF_UE_NGAP_ID *amfUeNgapId_mp;

    //RAN-UE-NGAP-ID
    RAN_UE_NGAP_ID ranUeNgapId_m;
    RAN_UE_NGAP_ID *ranUeNgapId_mp;

    //NAS-PDU
    NAS_PDU nasPdu_m;
    NasPdu nasPdu_ma;
    NAS_PDU *nasPdu_mp;
    bool nasPduPrst_m;

    //UE Aggregate Maximum Bit Rate
    UEAggregateMaximumBitRate     ueAggregateMaximumBitRate_m;
    UEAggregateMaximumBitRate*    ueAggregateMaximumBitRate_mp;
    bool ueAMBRPresent_m;

    //RANPagingPriority
    RANPagingPriority ranPagePriority_m;
    RANPagingPriority *ranPagePriority_mp;
    bool ranPagePriorityPrst_m;

    //PDUSessionResourceSetupListSUReq
    PDUSessionResourceSetupListSUReq pduSessResSetupListSuReq_m;
    struct PDUSessionResourceSetupListSUReq_ pduSessResSetupListSuReq_ma[maxnoofPDUSessions_c];
    unsigned int numofPduSessResSetupListItem_m;
    PDUSessionResourceSetupListSUReq *pduSessResSetupListSuReq_mp;

    NasPdu pduNasPdus_ma[maxnoofPDUSessions_c];
    SmfTransferPdu transfers_ma[maxnoofPDUSessions_c];

};
	
#endif
#ifndef __ueContextReleaseRequestPdu_h__INCLUDED__
#define __ueContextReleaseRequestPdu_h__INCLUDED__

#include "ngapOssPdu.h"
#include <ngapMsgTypes.h>
///////////////////////////////////////////////////////////////////////////////
//-- **************************************************************
//--
//-- UE Context Release Request Elementary Procedure
//--
//-- **************************************************************
//
//-- **************************************************************
//--
//-- UE CONTEXT RELEASE REQUEST
//--
//-- **************************************************************
//
//UEContextReleaseRequest ::= SEQUENCE {
//	protocolIEs		ProtocolIE-Container		{ {UEContextReleaseRequest-IEs} },
//	...
//}
//
//UEContextReleaseRequest-IEs NGAP-PROTOCOL-IES ::= {
//	{ ID id-AMF-UE-NGAP-ID		                CRITICALITY reject	TYPE AMF-UE-NGAP-ID		                    PRESENCE mandatory	}|
//	{ ID id-RAN-UE-NGAP-ID		                CRITICALITY reject	TYPE RAN-UE-NGAP-ID		                    PRESENCE mandatory	}|
//  { ID id-PDUSessionResourceListCxtRelReq		CRITICALITY reject	TYPE PDUSessionResourceListCxtRelReq		PRESENCE optional		}|
//	{ ID id-Cause				                CRITICALITY ignore	TYPE Cause				                    PRESENCE mandatory	},
//	...
//}
///////////////////////////////////////////////////////////////////////////////

class UeContextReleaseRequestPdu:public NgapOssPdu
{
  public:
    UeContextReleaseRequestPdu();
    //virtual
    ~UeContextReleaseRequestPdu();
    
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
    /// @brief get UeContextReleaseRequest pointer
    /// returns the pointer to the
    /// UeContextReleaseRequest in the decoded message
    /////////////////////////////////////////////////////////////
    UEContextReleaseRequest * getUeContextReleaseRequestDecodedPtr() const;

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
    /// Cause
    /////////////////////////////////////////////////////////////
    Cause* getCauseDecodedPtr() const;
    Cause* getCauseEncodePtr();

    /////////////////////////////////////////////////////////////
    /// PDUSessionResourceListCxtRelReq
    /////////////////////////////////////////////////////////////
    PDUSessionResourceListCxtRelReq* getPduSessResListCxtRelReqDecodedPtr() const;
    PDUSessionResourceListCxtRelReq* getPduSessResListCxtRelReqEncodePtr();
    bool isPduSessResListCxtRelReqPresent();	
    
    PDUSessionResourceItemCxtRelReq* getPduSessResItemCxtRelReqDecodedPtr(unsigned int index) const;
    PDUSessionResourceItemCxtRelReq* getPduSessResItemCxtRelReqEncodePtr(unsigned int index);

    void setNumofPduSessResListItemCxtRelReq_v(unsigned int num);
    unsigned int getNumofPduSessResListItemCxtRelReq() const;


  private:
    typedef UEContextReleaseRequestIe ProtocolIeType;
  
    //OSS structures
    UEContextReleaseRequest *UeContextReleaseRequest_mp; //for encode

    //AMF-UE-NGAP-ID
    AMF_UE_NGAP_ID amfUeNgapId_m;
    AMF_UE_NGAP_ID* amfUeNgapId_mp;

    //RAN-UE-NGAP-ID
    RAN_UE_NGAP_ID ranUeNgapId_m;
    RAN_UE_NGAP_ID* ranUeNgapId_mp;

    //Cause
	Cause cause_m;
	Cause* cause_mp;

	//PDUSessionResourceListCxtRelReq
    PDUSessionResourceListCxtRelReq pduSessResListCxtRelReq_m;
    struct PDUSessionResourceListCxtRelReq_ pduSessResListCxtRelReq_ma[maxnoofPDUSessions_c];
    unsigned int numofPduSessResListItemCxtRelReq_m;
    PDUSessionResourceListCxtRelReq* pduSessResListCxtRelReq_mp;
	bool pduSessResListCxtRelReqPresent_m;

};

#endif
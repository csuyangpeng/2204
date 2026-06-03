#ifndef __initialContextSetupFailurePdu_h__INCLUDED__
#define __initialContextSetupFailurePdu_h__INCLUDED__

#include "ngapOssPdu.h"
#include <ngapMsgTypes.h>
///////////////////////////////////////////////////////////////////////////////
// @class InitialContextSetupFailurePdu
// -- **************************************************************
// --
// -- INITIAL CONTEXT SETUP FAILURE
// --
// -- **************************************************************
// 
// InitialContextSetupFailure ::= SEQUENCE {
// 	protocolIEs		ProtocolIE-Container		{ {InitialContextSetupFailureIEs} },
// 	...
// }
// 
// InitialContextSetupFailureIEs NGAP-PROTOCOL-IES ::= {
// 	{ ID id-AMF-UE-NGAP-ID								CRITICALITY ignore	TYPE AMF-UE-NGAP-ID												PRESENCE mandatory	}|
// 	{ ID id-RAN-UE-NGAP-ID								CRITICALITY ignore	TYPE RAN-UE-NGAP-ID												PRESENCE mandatory	}|
// 	{ ID id-PDUSessionResourceFailedToSetupListCxtFail	CRITICALITY ignore	TYPE PDUSessionResourceFailedToSetupListCxtFail		PRESENCE optional		}|
// 	{ ID id-Cause										CRITICALITY ignore	TYPE Cause														PRESENCE mandatory	}|
// 	{ ID id-CriticalityDiagnostics						CRITICALITY ignore	TYPE CriticalityDiagnostics									PRESENCE optional		},
// 	...
// }
///////////////////////////////////////////////////////////////////////////////

class InitialContextSetupFailurePdu:public NgapOssPdu
{
  public:
    InitialContextSetupFailurePdu();
    //virtual
    ~InitialContextSetupFailurePdu();
    
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
    /// @brief get InitialContextSetupRequest pointer
    /// returns the pointer to the
    /// InitialContextSetupRequest in the decoded message
    /////////////////////////////////////////////////////////////
    InitialContextSetupFailure * getInitialContextSetupFailureDecodedPtr() const;

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
    /// PDUSessionResourceFailedToSetupListCxtFail
    /////////////////////////////////////////////////////////////
    PDUSessionResourceFailedToSetupListCxtFail* getPDUSessionResourceSetupListCxtFailDecodedPtr() const;
    PDUSessionResourceFailedToSetupListCxtFail* getPDUSessionResourceSetupListCxtFailEncodePtr();
    
    PDUSessionResourceFailedToSetupItemCxtFail* getPDUSessionResourceSetupItemCxtFailDecodedPtr(unsigned int index) const;
    PDUSessionResourceFailedToSetupItemCxtFail* getPDUSessionResourceSetupItemCxtFailEncodePtr(unsigned int index);

    void setNumofPDUSessionResourceSetupListItemCxtFail_v(unsigned int num);
    unsigned int getNumofPDUSessionResourceSetupListItemCxtFail() const;

  private:
    typedef InitialContextSetupFailureIe ProtocolIeType;
  
    //OSS structures
    InitialContextSetupFailure *InitialContextSetupFailure_mp; //for encode

    //AMF-UE-NGAP-ID
    AMF_UE_NGAP_ID amfUeNgapId_m;
    AMF_UE_NGAP_ID* amfUeNgapId_mp;

    //RAN-UE-NGAP-ID
    RAN_UE_NGAP_ID ranUeNgapId_m;
    RAN_UE_NGAP_ID* ranUeNgapId_mp;

    //Cause
    Cause cause_m;
    Cause* cause_mp;

    //PDUSessionResourceFailedToSetupListCxtFail
    PDUSessionResourceFailedToSetupListCxtFail pDUSessionResourceSetupListCxtFail_m;
    struct PDUSessionResourceFailedToSetupListCxtFail_ PDUSessionResSetupListCxtFail_ma[maxnoofPDUSessions_c];
    unsigned int numofPDUSessionResourceSetupListItemCxtFail_m;
    PDUSessionResourceFailedToSetupListCxtFail* pDUSessionResourceSetupListCxtFail_mp;

    SmfTransferPdu transfers_ma[maxnoofPDUSessions_c];
};

#endif
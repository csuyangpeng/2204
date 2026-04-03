#ifndef __ngSetupFailurePdu_h__INCLUDED__
#define __ngSetupFailurePdu_h__INCLUDED__

#include "ngapOssPdu.h"
#include <ngapMsgTypes.h>
///////////////////////////////////////////////////////////////////////////////
// @class NGSetupFailurePdu
//-- **************************************************************
//--
//-- NG SETUP FAILURE
//--
//-- **************************************************************
//
//NGSetupFailure ::= SEQUENCE {
//	protocolIEs		ProtocolIE-Container		{ {NGSetupFailureIEs} },
//	...
//}
//
//NGSetupFailureIEs NGAP-PROTOCOL-IES ::= {
//	{ ID id-Cause						CRITICALITY ignore	TYPE Cause					    PRESENCE mandatory	}|
//	{ ID id-TimeToWait			        CRITICALITY ignore	TYPE TimeToWait			        PRESENCE optional	}|
//	{ ID id-CriticalityDiagnostics		CRITICALITY ignore	TYPE CriticalityDiagnostics		PRESENCE optional		},
//	...
//}
///////////////////////////////////////////////////////////////////////////////

class NGSetupFailurePdu:public NgapOssPdu
{
  public:
    NGSetupFailurePdu();
    //virtual
    ~NGSetupFailurePdu();
    
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
    /// @brief get NGSetupFailure pointer
    /// returns the pointer to the
    /// NGSetupFailure in the decoded message
    /////////////////////////////////////////////////////////////
    NGSetupFailure * getNGSetupFailureDecodedPtr() const;

    /////////////////////////////////////////////////////////////
    /// Cause
    /////////////////////////////////////////////////////////////
    Cause* getCauseDecodedPtr() const;
    Cause* getCauseEncodePtr();

    /////////////////////////////////////////////////////////////
    /// TimeToWait
    /////////////////////////////////////////////////////////////
    TimeToWait* getTimeToWaitDecodedPtr() const;
    TimeToWait* getTimeToWaitEncodePtr();
	bool isTimeToWaitPresent();

  private:
  
    typedef NGSetupFailureIe ProtocolIeType;

    //OSS structures
    NGSetupFailure *NGSetupFailure_mp; //for encode

    //Cause
	Cause cause_m;
	Cause* cause_mp;

    //Time To Wait
	TimeToWait timeToWait_m;
	TimeToWait* timeToWait_mp;
	bool timeToWaitPresent_m;

};

#endif // __ngSetupFailurePdu_h__INCLUDED__
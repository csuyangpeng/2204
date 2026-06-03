#ifndef __ngapResetPdu_h__INCLUDED__
#define __ngapResetPdu_h__INCLUDED__

#include "ngapOssPdu.h"
#include <ngapMsgTypes.h>
///////////////////////////////////////////////////////////////////////////////
//-- **************************************************************
//--
//-- NG Reset Elementary Procedure
//--
//-- **************************************************************
//
//-- **************************************************************
//--
//-- NG RESET
//--
//-- **************************************************************
//
// NGReset ::= SEQUENCE {
//	protocolIEs		ProtocolIE-Container		{ {NGResetIEs} },
//	...
// }
//
// NGResetIEs NGAP-PROTOCOL-IES ::= {
// 	{ ID id-Cause						CRITICALITY ignore	TYPE Cause						PRESENCE mandatory	}|
//	{ ID id-ResetType					CRITICALITY reject	TYPE ResetType					PRESENCE mandatory	},
//	...
// }
///////////////////////////////////////////////////////////////////////////////

class NGResetPdu:public NgapOssPdu
{
  public:
    NGResetPdu();
    //virtual
    ~NGResetPdu();
    
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
    NGReset * getNGResetDecodedPtr() const;

    /////////////////////////////////////////////////////////////
    /// Cause
    /////////////////////////////////////////////////////////////
    Cause* getCauseDecodedPtr() const;
    Cause* getCauseEncodePtr();

    /////////////////////////////////////////////////////////////
    /// ResetType
    /////////////////////////////////////////////////////////////
    ResetType* getResetTypeDecodedPtr() const;
    ResetType* getResetTypeEncodePtr();

  private:
    typedef NGResetIe ProtocolIeType;
  
    //OSS structures
    NGReset *NGReset_mp; //for encode

    //Cause
    Cause cause_m;
    Cause* cause_mp;
	
    //ResetType
    ResetType resetType_m; 
    ResetType* resetType_mp; 

};

#endif // __ngapResetPdu_h__INCLUDED__
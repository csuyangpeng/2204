#ifndef __ueContextReleaseCommandPdu_h__INCLUDED__
#define __ueContextReleaseCommandPdu_h__INCLUDED__

#include "ngapOssPdu.h"
#include <ngapMsgTypes.h>
///////////////////////////////////////////////////////////////////////////////
//-- **************************************************************
//--
//-- UE Context Release Elementary Procedure
//--
//-- **************************************************************
//
//-- **************************************************************
//--
//-- UE CONTEXT RELEASE COMMAND
//--
//-- **************************************************************
//
//UEContextReleaseCommand ::= SEQUENCE {
//	protocolIEs		ProtocolIE-Container		{ {UEContextReleaseCommand-IEs} },
//	...
//}
//
//UEContextReleaseCommand-IEs NGAP-PROTOCOL-IES ::= {
//	{ ID id-UE-NGAP-IDs				CRITICALITY reject	TYPE UE-NGAP-IDs			PRESENCE mandatory	}|
//	{ ID id-Cause					CRITICALITY ignore	TYPE Cause					PRESENCE mandatory	},
//	...
//}
///////////////////////////////////////////////////////////////////////////////

class UeContextReleaseCommandPdu:public NgapOssPdu
{
  public:
    UeContextReleaseCommandPdu();
    //virtual
    ~UeContextReleaseCommandPdu();
    
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
    /// @brief get UeContextReleaseCommand pointer
    /// returns the pointer to the
    /// UeContextReleaseCommand in the decoded message
    /////////////////////////////////////////////////////////////
    UEContextReleaseCommand * getUeContextReleaseCommandDecodedPtr() const;

    /////////////////////////////////////////////////////////////
    /// UE_NGAP_IDs
    /////////////////////////////////////////////////////////////
    UE_NGAP_IDs* getUeNgapIdsDecodedPtr() const;
    UE_NGAP_IDs* getUeNgapIdsEncodePtr();

    /////////////////////////////////////////////////////////////
    /// Cause
    /////////////////////////////////////////////////////////////
    Cause* getCauseDecodedPtr() const;
    Cause* getCauseEncodePtr();

  private:
    typedef UEContextReleaseCommandIe ProtocolIeType;
  
    //OSS structures
    UEContextReleaseCommand *ueContextReleaseCommand_mp; //for encode

    //UE-NGAP-IDs
    UE_NGAP_IDs ueNgapIds_m;
    UE_NGAP_IDs* ueNgapIds_mp;

    //Cause
	Cause cause_m;
	Cause* cause_mp;
};

#endif
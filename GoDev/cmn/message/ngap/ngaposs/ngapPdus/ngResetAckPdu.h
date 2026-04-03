#ifndef __ngapResetAckPdu_h__INCLUDED__
#define __ngapResetAckPdu_h__INCLUDED__

#include "ngapOssPdu.h"
#include <ngapMsgTypes.h>
///////////////////////////////////////////////////////////////////////////////
// -- **************************************************************
// --
// -- NG RESET ACKNOWLEDGE
// --
// -- **************************************************************
//
// NGResetAcknowledge ::= SEQUENCE {
// 	protocolIEs		ProtocolIE-Container		{ {NGResetAcknowledgeIEs} },
// 	...
// }
//
// NGResetAcknowledgeIEs NGAP-PROTOCOL-IES ::= {
// 	{ ID id-UE-associatedLogicalNG-connectionList			CRITICALITY ignore	TYPE UE-associatedLogicalNG-connectionList			PRESENCE optional	}|
// 	{ ID id-CriticalityDiagnostics							CRITICALITY ignore	TYPE CriticalityDiagnostics								PRESENCE optional	},
// 	...
// }
///////////////////////////////////////////////////////////////////////////////

class NGResetAckPdu:public NgapOssPdu
{
  public:
    NGResetAckPdu();
    //virtual
    ~NGResetAckPdu();
    
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
    /// @brief get NGResetAcknowledge pointer
    /// returns the pointer to the
    /// NGResetAcknowledge in the decoded message
    /////////////////////////////////////////////////////////////
    NGResetAcknowledge * getNGResetAckDecodedPtr() const;

    /////////////////////////////////////////////////////////////
    /// UE Associated Logical NG Connection List
    /////////////////////////////////////////////////////////////
    UE_associatedLogicalNG_connectionList * getUEAssNGConnListDecodedPtr() const;
    UE_associatedLogicalNG_connectionList * getUEAssNGConnListEncodePtr();
    bool isUEAssNGConnPresent();

    UE_associatedLogicalNG_connectionItem * getUEAssNGConnListItemDecodedPtr(unsigned int index) const;
    UE_associatedLogicalNG_connectionItem * getUEAssNGConnListItemEncodePtr(unsigned int index);

    void setNumOfUEAssNGConnListItem_v(unsigned int num);
    unsigned int getNumOfUEAssNGConnListItem() const;

  private:
    typedef NGResetAckIe ProtocolIeType;
  
    //OSS structures
    NGResetAcknowledge *NGResetAck_mp; //for encode
	
    //UE Associated Logical NG Connection List
    UE_associatedLogicalNG_connectionList ueAssNGConnList_m;
    struct UE_associatedLogicalNG_connectionList_ ueAssNGConnList_ma[maxnoofNGConnectionsToReset_c];
	unsigned int numOfUEAssNGConnListItem_m;
    UE_associatedLogicalNG_connectionList *ueAssNGConnList_mp;
    bool ueAssNGConnPresent_m;

};

#endif // __ngapResetAckPdu_h__INCLUDED__
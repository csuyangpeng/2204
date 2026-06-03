#ifndef __ngApOssInitialContextSetupResponsePdu_h__INCLUDED__
#define __ngApOssInitialContextSetupResponsePdu_h__INCLUDED__

#include "ngapOssPdu.h"
#include <ngapMsgTypes.h>
///////////////////////////////////////////////////////////////////////////////
// @class InitialContextSetupResponsePdu
//-- **************************************************************
//--
//-- INITIAL CONTEXT SETUP RESPONSE
//--
//-- **************************************************************
//
//InitialContextSetupResponse ::= SEQUENCE {
//	protocolIEs		ProtocolIE-Container		{ {InitialContextSetupResponseIEs} },
//	...
//}
//
//InitialContextSetupResponseIEs NGAP-PROTOCOL-IES ::= {
//	{ ID id-AMF-UE-NGAP-ID								CRITICALITY ignore	TYPE AMF-UE-NGAP-ID												PRESENCE mandatory	}|
//	{ ID id-RAN-UE-NGAP-ID								CRITICALITY ignore	TYPE RAN-UE-NGAP-ID												PRESENCE mandatory	}|
//	{ ID id-PDUSessionResourceSetupListCxtRes			CRITICALITY ignore	TYPE PDUSessionResourceSetupListCxtRes					PRESENCE optional		}|
//	{ ID id-PDUSessionResourceFailedToSetupListCxtRes	CRITICALITY ignore	TYPE PDUSessionResourceFailedToSetupListCxtRes		PRESENCE optional		}|
//	{ ID id-CriticalityDiagnostics						CRITICALITY ignore	TYPE CriticalityDiagnostics									PRESENCE optional		},
//	...
//}
///////////////////////////////////////////////////////////////////////////////

class InitialContextSetupResponsePdu:public NgapOssPdu
{
  public:
    InitialContextSetupResponsePdu();
    //virtual
    ~InitialContextSetupResponsePdu();
    
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
    /// @brief get InitialContextSetupResponse pointer
    /// returns the pointer to the
    /// InitialContextSetupResponse in the decoded message
    /////////////////////////////////////////////////////////////
    InitialContextSetupResponse * getInitialContextSetupResponseDecodedPtr() const;

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
    /// PDUSessionResourceSetupListCxtRes
    /////////////////////////////////////////////////////////////
    PDUSessionResourceSetupListCxtRes* getPDUSessionResourceSetupListCxtResDecodedPtr() const;
    PDUSessionResourceSetupListCxtRes* getPDUSessionResourceSetupListCxtResEncodePtr();
    
    PDUSessionResourceSetupItemCxtRes* getPDUSessionResourceSetupItemCxtResDecodedPtr(unsigned int index) const;
    PDUSessionResourceSetupItemCxtRes* getPDUSessionResourceSetupItemCxtResEncodePtr(unsigned int index);

    void setNumofPDUSessionResourceSetupListItemCxtRes_v(unsigned int num);
    unsigned int getNumofPDUSessionResourceSetupListItemCxtRes() const;

    /////////////////////////////////////////////////////////////
    /// PDUSessionResourceFailedToSetupListCxtRes
    /////////////////////////////////////////////////////////////
    PDUSessionResourceFailedToSetupListCxtRes* getPDUSessionResourceFailedToSetupListCxtResDecodedPtr() const;
    PDUSessionResourceFailedToSetupListCxtRes* getPDUSessionResourceFailedToSetupListCxtResEncodePtr();
    
    PDUSessionResourceFailedToSetupItemCxtRes* getPDUSessionResourceFailedToSetupItemCxtResDecodedPtr(unsigned int index) const;
    PDUSessionResourceFailedToSetupItemCxtRes* getPDUSessionResourceFailedToSetupItemCxtResEncodePtr(unsigned int index);

    void setNumofPDUSessionResourceFailedToSetupListItemCxtRes_v(unsigned int num);
    unsigned int getNumofPDUSessionResourceFailedToSetupListItemCxtRes() const;


  private:
    typedef InitialContextSetupResponseIe ProtocolIeType;
  
    //OSS structures
    InitialContextSetupResponse *InitialContextSetupResponse_mp; //for encode

    //AMF-UE-NGAP-ID
    AMF_UE_NGAP_ID amfUeNgapId_m;
    AMF_UE_NGAP_ID* amfUeNgapId_mp;

    //RAN-UE-NGAP-ID
    RAN_UE_NGAP_ID ranUeNgapId_m;
    RAN_UE_NGAP_ID* ranUeNgapId_mp;

  	//PDUSessionResourceSetupListCxtRes
    PDUSessionResourceSetupListCxtRes pDUSessionResourceSetupListCxtRes_m;
    struct PDUSessionResourceSetupListCxtRes_ pDUSessionResSetupListCxtRes_ma[maxnoofPDUSessions_c];
    unsigned int numofPDUSessionResourceSetupListItemCxtRes_m;
    PDUSessionResourceSetupListCxtRes* pDUSessionResourceSetupListCxtRes_mp;
    SmfTransferPdu setupTransfers_ma[maxnoofPDUSessions_c];

	  //Failed list
  	PDUSessionResourceFailedToSetupListCxtRes pDUSessionResourceFailedToSetupListCxtRes_m;
  	struct PDUSessionResourceFailedToSetupListCxtRes_ pDUSessionResourceFailedToSetupListCxtRes_ma[maxnoofPDUSessions_c];
	  unsigned int numofPDUSessionResourceFailedToSetupItemCxtRes_m;
  	PDUSessionResourceFailedToSetupListCxtRes* pDUSessionResourceFailedToSetupListCxtRes_mp;
    SmfTransferPdu failedTransfers_ma[maxnoofPDUSessions_c];

};

#endif
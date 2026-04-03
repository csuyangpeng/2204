#ifndef __pduSessionResourceSetupResponsePdu_h__INCLUDED__
#define __pduSessionResourceSetupResponsePdu_h__INCLUDED__

#include "ngapOssPdu.h"
#include <ngapMsgTypes.h>
///////////////////////////////////////////////////////////////////////////////
// @class NgApOssPDUSessionResourceSetupResponsePdu
//-- **************************************************************
//--
//-- PDU Session Resource Setup Response
//--
//-- **************************************************************
//
//
//PDUSessionResourceSetupResponse ::= SEQUENCE {
//	protocolIEs		ProtocolIE-Container		{ {PDUSessionResourceSetupResponseIEs} },
//	...
//}
//
//PDUSessionResourceSetupResponseIEs NGAP-PROTOCOL-IES ::= {
//	{ ID id-AMF-UE-NGAP-ID								CRITICALITY ignore	TYPE AMF-UE-NGAP-ID						        	PRESENCE mandatory	}|
//	{ ID id-RAN-UE-NGAP-ID								CRITICALITY ignore	TYPE RAN-UE-NGAP-ID									PRESENCE mandatory	}|
//	{ ID id-PDUSessionResourceSetupListSURes			CRITICALITY ignore	TYPE PDUSessionResourceSetupListSURes				PRESENCE optional		}|
//	{ ID id-PDUSessionResourceFailedToSetupListSURes	CRITICALITY ignore	TYPE PDUSessionResourceFailedToSetupListSURes		PRESENCE optional		},
//	...	
//}
///////////////////////////////////////////////////////////////////////////////

class PDUSessionResourceSetupResponsePdu:public NgapOssPdu
{
  public:
    PDUSessionResourceSetupResponsePdu();
    //virtual
    ~PDUSessionResourceSetupResponsePdu();
    
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
    /// @brief get PDUSessionResourceSetupResponse pointer
    /// returns the pointer to the
    /// PDUSessionResourceSetupResponse in the decoded message
    /////////////////////////////////////////////////////////////
    PDUSessionResourceSetupResponse * getPduSessResSetupRespDecodedPtr() const;

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
    /// PDUSessionResourceSetupListSURes
    /////////////////////////////////////////////////////////////
    PDUSessionResourceSetupListSURes* getPduSessResSetupListDecodedPtr() const;
    PDUSessionResourceSetupListSURes* getPduSessResSetupListEncodePtr();
    
    PDUSessionResourceSetupItemSURes* getPduSessResSetupItemDecodedPtr(unsigned int index) const;
    PDUSessionResourceSetupItemSURes* getPduSessResSetupItemEncodePtr(unsigned int index);
    
    void setNumofSetupListItem_v(unsigned int num);
    unsigned int getNumofSetupListItem() const;

    /////////////////////////////////////////////////////////////
    /// PDUSessionResourceFailedToSetupListSURes
    /////////////////////////////////////////////////////////////
    PDUSessionResourceFailedToSetupListSURes* getPduSessResFailedSetupListpDecodedPtr() const;
    PDUSessionResourceFailedToSetupListSURes* getPduSessResFailedSetupListEncodePtr();
    
    PDUSessionResourceFailedToSetupItemSURes* getPduSessResFailedSetupItemDecodedPtr(unsigned int index) const;
    PDUSessionResourceFailedToSetupItemSURes* getPduSessResFailedSetupItemEncodePtr(unsigned int index);
    
    void setNumofFailedSetupItem_v(unsigned int num);
    unsigned int getNumofFailedSetupItem() const;

  private:
    typedef PDUSessResSetupRespIe ProtocolIeType;
  
    //OSS structures
    PDUSessionResourceSetupResponse *pduSessResSetupResp_mp; 

    //AMF-UE-NGAP-ID
    AMF_UE_NGAP_ID amfUeNgapId_m;
    AMF_UE_NGAP_ID* amfUeNgapId_mp;

    //RAN-UE-NGAP-ID
    RAN_UE_NGAP_ID ranUeNgapId_m;
    RAN_UE_NGAP_ID* ranUeNgapId_mp;

    //PDUSessionResourceSetupListSUResp
    PDUSessionResourceSetupListSURes pduSessResSetupList_m;
    struct PDUSessionResourceSetupListSURes_ pduSessResSetupList_ma[maxnoofPDUSessions_c];
    unsigned int numofPduSessResSetupList_m;
    PDUSessionResourceSetupListSURes* pduSessResSetupList_mp;
    SmfTransferPdu setupTransfers_ma[maxnoofPDUSessions_c];

    //Failed list
    PDUSessionResourceFailedToSetupListSURes pduSessResFailedSetupList_m;
    struct PDUSessionResourceFailedToSetupListSURes_ pduSessResFailedSetupList_ma[maxnoofPDUSessions_c];
    unsigned int numofpduSessResFailedSetupList_m;
    PDUSessionResourceFailedToSetupListSURes* pduSessResFailedSetupList_mp;
    SmfTransferPdu failedTransfers_ma[maxnoofPDUSessions_c];
};

#endif
#ifndef __pduSessResRelCmdTransferPdu_h__INCLUDED__
#define __pduSessResRelCmdTransferPdu_h__INCLUDED__

#include "ngapOssPdu.h"
#include <ngapMsgTypes.h>
///////////////////////////////////////////////////////////////////////////////
// @class PduSessResRelCmdTransferPdu
//-- **************************************************************
//--
//-- PDU Session Resource Release Command Transfer 
//--
//-- **************************************************************
//
//
//Cause ::= CHOICE {
//	radioNetwork		CauseRadioNetwork,
//	transport			CauseTransport,
//	nas					CauseNas,
//	protocol			CauseProtocol,
//	misc				CauseMisc,
//	choice-Extensions		ProtocolIE-SingleContainer { {Cause-ExtIEs} }
//}
///////////////////////////////////////////////////////////////////////////////

class PduSessResRelCmdTransferPdu : public NgapOssPdu
{
  public:
    PduSessResRelCmdTransferPdu();
    //virtual
    ~PduSessResRelCmdTransferPdu();

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

    //overwrite the decode methord
	  unsigned int decodePerToOss(NgapOssCtxt* ctxt);
	  void getDecodedOssDataPtrs_v();

    PDUSessionResourceReleaseCommandTransfer *getPduSessResRelCmdTransferDecodedPtr() const;

    Cause * getCauseEncodePtr();
    Cause * getCauseDecodePtr();

  private:
    //OSS structures
	  PDUSessionResourceReleaseCommandTransfer* pduSessResRelCmdTransfer_mp;

	  Cause cause_m;
	  Cause* cause_mp;
};

#endif

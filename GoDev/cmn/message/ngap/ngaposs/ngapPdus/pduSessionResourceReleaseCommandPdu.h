#ifndef __pduSessionResourceReleaseCommandPdu_h__INCLUDED__
#define __pduSessionResourceReleaseCommandPdu_h__INCLUDED__

#include "ngapOssPdu.h"
#include <ngapMsgTypes.h>
///////////////////////////////////////////////////////////////////////////////
// @class PDUSessionResourceReleaseCommandPdu
//-- **************************************************************
//--
//-- PDU SESSION RESOURCE RELEASE COMMAND
//--
//-- **************************************************************
//
//
// PDUSessionResourceReleaseCommandIEs NGAP-PROTOCOL-IES ::= {
// { ID id-AMF-UE-NGAP-ID								CRITICALITY reject	TYPE AMF-UE-NGAP-ID										PRESENCE mandatory	}|
// { ID id-RAN-UE-NGAP-ID								CRITICALITY reject	TYPE RAN-UE-NGAP-ID										PRESENCE mandatory	}|
// { ID id-RANPagingPriority							CRITICALITY ignore	TYPE RANPagingPriority								PRESENCE optional		}|
// { ID id-NAS-PDU										CRITICALITY ignore	TYPE NAS-PDU											PRESENCE optional		}|
// { ID id-PDUSessionResourceToReleaseListRelCmd		CRITICALITY reject	TYPE PDUSessionResourceToReleaseListRelCmd		PRESENCE mandatory	},
// ...
// }
// 
///////////////////////////////////////////////////////////////////////////////

class PduSessResRelCmdPdu:public NgapOssPdu
{
  public:
    PduSessResRelCmdPdu();
    //virtual
    ~PduSessResRelCmdPdu();
    
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
    /// @brief get PDUSessionResourceReleaseCommandPdu pointer
    /// returns the pointer to the
    /// PDUSessionResourceReleaseCommandPdu in the decoded message
    /////////////////////////////////////////////////////////////
    PDUSessionResourceReleaseCommand * getPduSessResRelCmdDecodedPtr() const;

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
    /// NAS_PDU
    /////////////////////////////////////////////////////////////
    NAS_PDU* getNasPduDecodedPtr() const;
    NAS_PDU* getNasPduEncodePtr();
    bool isNasPduPresent();

    /////////////////////////////////////////////////////////////
    /// RANPagingPriority
    /////////////////////////////////////////////////////////////
    RANPagingPriority *getRANPagingPriorityDecodedPtr() const;
    RANPagingPriority *getRANPagingPriorityEncodePtr();
    bool isRANPagingPriorityPresent();

    /////////////////////////////////////////////////////////////
    /// PDUSessionResourceReleaseCommand
    /////////////////////////////////////////////////////////////
    PDUSessionResourceToReleaseListRelCmd* getPduSessResToRelListRelCmdDecodedPtr() const;
    PDUSessionResourceToReleaseListRelCmd* getPduSessResToRelListRelCmdEncodePtr();
    
    PDUSessionResourceToReleaseItemRelCmd* getPduSessResToRelItemRelCmdDecodedPtr(unsigned int index) const;
    PDUSessionResourceToReleaseItemRelCmd* getPduSessResToRelItemRelCmdEncodePtr(unsigned int index);

    void setNumofPduSessResToRelItemRelCmd_v(unsigned int num);
    unsigned int getNumofPduSessResToRelItemRelCmd() const;

  private:
    typedef PduSessResRelCmdIe ProtocolIeType;
  
    //OSS structures
    PDUSessionResourceReleaseCommand *PduSessResRelCmd_mp; //for encode

    //AMF-UE-NGAP-ID
    AMF_UE_NGAP_ID amfUeNgapId_m;
    AMF_UE_NGAP_ID* amfUeNgapId_mp;

    //RAN-UE-NGAP-ID
    RAN_UE_NGAP_ID ranUeNgapId_m;
    RAN_UE_NGAP_ID* ranUeNgapId_mp;

    //NAS-PDU
    NAS_PDU nasPdu_m;
    unsigned char nasPdu_ma[maxNasPduValue_c];
    NAS_PDU* nasPdu_mp;
    bool nasPduPresent_m;

    //RANPagingPriority
    RANPagingPriority ranPagingPriority_m;
    RANPagingPriority *ranPagingPriority_mp;
    bool ranPagingPriorityPresent_m;

    //PDUSessionResourceToReleaseListRelCmd
    PDUSessionResourceToReleaseListRelCmd pduSessResToRelListRelCmd_m;
    struct PDUSessionResourceToReleaseListRelCmd_ pduSessResToRelListRelCmd_ma[maxnoofPDUSessions_c];
    unsigned int numofPduSessResToRelListItemRelCmd_m;
    PDUSessionResourceToReleaseListRelCmd* pduSessResToRelListRelCmd_mp;

    SmfTransferPdu transfers_ma[maxnoofPDUSessions_c];
};

#endif
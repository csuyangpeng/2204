#include <string.h>
#include <iostream>

#include "nasNonDeliveryIndicationPdu.h"

NasNonDeliveryIndicationPdu::NasNonDeliveryIndicationPdu()
{
    reset_v();
}

NasNonDeliveryIndicationPdu::~NasNonDeliveryIndicationPdu()
{
}

void
NasNonDeliveryIndicationPdu::reset_v()
{
    NasNonDeliveryIndication_mp = NULL;
    
    //RAN UE NGAP ID
    ranUeNgapId_m=0;
    ranUeNgapId_mp=NULL;
    
    //AMF UE NGAP ID
    amfUeNgapId_m=0;
    amfUeNgapId_mp=NULL;

    //NAS PDU
    nasPdu_m.length=0;
    memset(nasPdu_ma, 0xff, sizeof(nasPdu_ma));
    nasPdu_m.value=nasPdu_ma;
    nasPdu_mp=NULL;

    //Cause
    cause_mp=NULL;
	memset(&cause_m,0,sizeof(cause_m));    
}

unsigned int
NasNonDeliveryIndicationPdu::encodeOssToPer(NgapOssCtxt* ctxt)                   
{
    if(NULL == ctxt)
    {
        std::cout <<"Invalid NgapOssCtxt!" << std::endl;                  
        return -1;
    }
    
    ProtocolIeType *nextIe_p = NULL;

    //Cause IE encode
    ProtocolIeType causeIe;
    causeIe.next=nextIe_p;
    causeIe.value.id=id_Cause;
    causeIe.value.criticality=ignore;
    causeIe.value.value.pduNum=PDU_NASNonDeliveryIndication_IEs_Value_Cause;
    causeIe.value.value.decoded.pdu_Cause=&cause_m;
    causeIe.value.value.encoded.length=0;
    causeIe.value.value.encoded.value=NULL;
    nextIe_p=&causeIe;

    //NAS PDU
    ProtocolIeType nasPduIe;
    nasPduIe.next =nextIe_p;
    nasPduIe.value.id = id_NAS_PDU;
    nasPduIe.value.criticality = ignore;
    nasPduIe.value.value.pduNum =PDU_NASNonDeliveryIndication_IEs_Value_NAS_PDU;
    nasPduIe.value.value.decoded.pdu_NAS_PDU = &nasPdu_m;
    nasPduIe.value.value.encoded.length = 0;
    nasPduIe.value.value.encoded.value = NULL;
    nextIe_p = &nasPduIe;

    //RAN_UE_NGAP_ID
    ProtocolIeType ranUeNgapIdIe;
    ranUeNgapIdIe.next =nextIe_p;
    ranUeNgapIdIe.value.id = id_RAN_UE_NGAP_ID;
    ranUeNgapIdIe.value.criticality = reject;
    ranUeNgapIdIe.value.value.pduNum =PDU_NASNonDeliveryIndication_IEs_Value_RAN_UE_NGAP_ID;
    ranUeNgapIdIe.value.value.decoded.pdu_RAN_UE_NGAP_ID = &ranUeNgapId_m;
    ranUeNgapIdIe.value.value.encoded.length = 0;
    ranUeNgapIdIe.value.value.encoded.value = NULL;
    nextIe_p = &ranUeNgapIdIe;

    //AMF_UE_NGAP_ID
    ProtocolIeType amfUeNgapIdIe;
    amfUeNgapIdIe.next =nextIe_p;
    amfUeNgapIdIe.value.id = id_AMF_UE_NGAP_ID;
    amfUeNgapIdIe.value.criticality = reject;
    amfUeNgapIdIe.value.value.pduNum =PDU_NASNonDeliveryIndication_IEs_Value_AMF_UE_NGAP_ID;
    amfUeNgapIdIe.value.value.decoded.pdu_AMF_UE_NGAP_ID = &amfUeNgapId_m;
    amfUeNgapIdIe.value.value.encoded.length = 0;
    amfUeNgapIdIe.value.value.encoded.value = NULL;
    nextIe_p = &amfUeNgapIdIe;

    //message
    NASNonDeliveryIndication msg;
    msg.protocolIEs = nextIe_p;

    NGAP_PDU ngApPdu;

    ngApPdu.choice = initiatingMessage_chosen;
    ngApPdu.u.initiatingMessage.procedureCode = id_NASNonDeliveryIndication;
    ngApPdu.u.initiatingMessage.criticality = ignore;
    ngApPdu.u.initiatingMessage.value.pduNum = PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_NASNonDeliveryIndication;
    ngApPdu.u.initiatingMessage.value.decoded.pdu_NASNonDeliveryIndication = &msg;
    ngApPdu.u.initiatingMessage.value.encoded.length = 0;
    ngApPdu.u.initiatingMessage.value.encoded.value = NULL; 

    //Convert Structures to PER
    OssBuf outputBuf;
    unsigned int result;

    ctxt->reset_v();
    outputBuf = ctxt->getPrimaryCodecBuf();

    result = ctxt->encodeOssToPer(NGAP_PDU_PDU, &ngApPdu, &outputBuf);

    if (result == encodeSuccessCode_c)
    {
        perBufRef_m = outputBuf;
    }

    return (result);
}

void
NasNonDeliveryIndicationPdu::getDecodedOssDataPtrs_v()
{
    ProtocolIeType *currIe_p;
    reset_v();

    ngapPdu_mp = reinterpret_cast<NGAP_PDU *> (ossBufRef_m.value);
    if (ngapPdu_mp == NULL)
    {
        return;
    }
    
    NasNonDeliveryIndication_mp =
            ngapPdu_mp->u.initiatingMessage.value.decoded.pdu_NASNonDeliveryIndication;
    
    if (NasNonDeliveryIndication_mp == NULL)
    {
        return;
    }

    currIe_p = NasNonDeliveryIndication_mp->protocolIEs;

    while (currIe_p != NULL)
    {
        if (currIe_p->value.id == id_RAN_UE_NGAP_ID
            && currIe_p->value.value.decoded.pdu_RAN_UE_NGAP_ID != NULL)
        {
            ranUeNgapId_mp = 
                      currIe_p->value.value.decoded.pdu_RAN_UE_NGAP_ID;
        }
        else if (currIe_p->value.id == id_NAS_PDU
                 && currIe_p->value.value.decoded.pdu_NAS_PDU != NULL)
        {
            nasPdu_mp = currIe_p->value.value.decoded.pdu_NAS_PDU;
        }
        else if (currIe_p->value.id == id_AMF_UE_NGAP_ID
            && currIe_p->value.value.decoded.pdu_AMF_UE_NGAP_ID != NULL)
        {
            amfUeNgapId_mp = 
                      currIe_p->value.value.decoded.pdu_AMF_UE_NGAP_ID;
        }
		else if (currIe_p->value.id == id_Cause
			&& currIe_p->value.value.decoded.pdu_Cause != NULL)
		{
			cause_mp = 
					  currIe_p->value.value.decoded.pdu_Cause;
		}
        currIe_p = currIe_p->next;
    }
}

NASNonDeliveryIndication *
NasNonDeliveryIndicationPdu::getNasNonDeliveryIndicationDecodedPtr() const
{
    return NasNonDeliveryIndication_mp;
}

RAN_UE_NGAP_ID * 
NasNonDeliveryIndicationPdu::getRanUeNgapIdDecodedPtr() const
{
    return ranUeNgapId_mp;
}

RAN_UE_NGAP_ID * 
NasNonDeliveryIndicationPdu::getRanUeNgapIdEncodePtr()
{
    return &ranUeNgapId_m;
}

AMF_UE_NGAP_ID * 
NasNonDeliveryIndicationPdu::getAmfUeNgapIdDecodedPtr() const
{
    return amfUeNgapId_mp;
}

AMF_UE_NGAP_ID * 
NasNonDeliveryIndicationPdu::getAmfUeNgapIdEncodePtr()
{
    return &amfUeNgapId_m;
}

NAS_PDU * 
NasNonDeliveryIndicationPdu::getNasPduDecodedPtr() const
{
    return nasPdu_mp;
}

NAS_PDU * 
NasNonDeliveryIndicationPdu::getNasPduEncodePtr()
{
    return &nasPdu_m;
}

Cause* 
NasNonDeliveryIndicationPdu::getCauseDecodedPtr() const
{
	return cause_mp;
}

Cause* 
NasNonDeliveryIndicationPdu::getCauseEncodePtr()
{
	return &cause_m;
}

#include <string.h>
#include <iostream>

#include "ngResetPdu.h"
NGResetPdu::NGResetPdu()
{
    reset_v();
}

NGResetPdu::~NGResetPdu()
{
}

void
NGResetPdu::reset_v()
{
    NGReset_mp = NULL;
    
    //Cause
    cause_mp=NULL;
	memset(&cause_m,0,sizeof(cause_m));

    //Reset Type
    resetType_m.choice=nG_Interface_chosen;
    resetType_m.u.nG_Interface = reset_all;
           
    resetType_mp=NULL;
}

unsigned int
NGResetPdu::encodeOssToPer(NgapOssCtxt* ctxt)                   
{
    if(NULL == ctxt)
	{
	    return -1;
    }
	
    ProtocolIeType *nextIe_p = NULL;

    //Reset Type
    ProtocolIeType resetTypeIe;
    resetTypeIe.next =nextIe_p;
    resetTypeIe.value.id = id_ResetType;
    resetTypeIe.value.criticality = reject;
    resetTypeIe.value.value.pduNum =PDU_NGResetIEs_Value_ResetType;
    resetTypeIe.value.value.decoded.pdu_ResetType = &resetType_m;
    resetTypeIe.value.value.encoded.length = 0;
    resetTypeIe.value.value.encoded.value = NULL;
    nextIe_p = &resetTypeIe;

    //Cause IE encode
    ProtocolIeType causeIe;
    causeIe.next=nextIe_p;
    causeIe.value.id=id_Cause;
    causeIe.value.criticality=ignore;
    causeIe.value.value.pduNum=PDU_NGResetIEs_Value_Cause;
    causeIe.value.value.decoded.pdu_Cause=&cause_m;
    causeIe.value.value.encoded.length=0;
    causeIe.value.value.encoded.value=NULL;
    nextIe_p=&causeIe;

    //message
    NGReset msg;
    msg.protocolIEs = nextIe_p;

    NGAP_PDU ngApPdu;

	ngApPdu.choice = initiatingMessage_chosen;
    ngApPdu.u.initiatingMessage.procedureCode = id_NGReset;
    ngApPdu.u.initiatingMessage.criticality = reject;
    ngApPdu.u.initiatingMessage.value.pduNum = PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_NGReset;
    ngApPdu.u.initiatingMessage.value.decoded.pdu_NGReset = &msg;
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
NGResetPdu::getDecodedOssDataPtrs_v()
{
    ProtocolIeType *currIe_p;
    reset_v();

    ngapPdu_mp = reinterpret_cast<NGAP_PDU *> (ossBufRef_m.value);
    if (ngapPdu_mp == NULL)
    {
        return;
    }
    
    NGReset_mp = ngapPdu_mp->u.initiatingMessage.value.decoded.pdu_NGReset;
    
    if (NGReset_mp == NULL)
    {
        return;
    }

    currIe_p = NGReset_mp->protocolIEs;

    while (currIe_p != NULL)
    {
        if (currIe_p->value.id == id_Cause
            && currIe_p->value.value.decoded.pdu_Cause != NULL)
        {
            cause_mp = currIe_p->value.value.decoded.pdu_Cause;
        }
        else if (currIe_p->value.id == id_ResetType
                 && currIe_p->value.value.decoded.pdu_ResetType != NULL)
        {
        	resetType_mp = currIe_p->value.value.decoded.pdu_ResetType;
        }

        currIe_p = currIe_p->next;
    }
}

NGReset *
NGResetPdu::getNGResetDecodedPtr() const
{
    return NGReset_mp;
}

Cause* 
NGResetPdu::getCauseDecodedPtr() const
{
	return cause_mp;
}

Cause* 
NGResetPdu::getCauseEncodePtr()
{
	return &cause_m;
}

ResetType * 
NGResetPdu::getResetTypeDecodedPtr() const
{
    return resetType_mp;
}

ResetType * 
NGResetPdu::getResetTypeEncodePtr()
{
    return &resetType_m;
}
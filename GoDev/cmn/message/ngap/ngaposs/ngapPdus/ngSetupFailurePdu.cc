#include <string.h>
#include <iostream>

#include "ngSetupFailurePdu.h"

NGSetupFailurePdu::NGSetupFailurePdu()
{
    reset_v();
}

NGSetupFailurePdu::~NGSetupFailurePdu()
{
}

void
NGSetupFailurePdu::reset_v()
{
    NGSetupFailure_mp=NULL;

    //Cause
    cause_mp=NULL;
	memset(&cause_m,0,sizeof(cause_m));

    //Time To Wait
    timeToWait_mp=NULL;
    timeToWait_m=v1s;
	timeToWaitPresent_m = false;
}

unsigned int
NGSetupFailurePdu::encodeOssToPer(NgapOssCtxt* ctxt)                           
{
    if(NULL == ctxt)
	{
	    return -1;
    }

    ProtocolIeType *nextIe_p = NULL;

    //Time To Wait
    ProtocolIeType timeToWaitIe;
	if(timeToWaitPresent_m)
	{
        timeToWaitIe.next=nextIe_p;
        timeToWaitIe.value.id=id_TimeToWait;
        timeToWaitIe.value.criticality=ignore;
        timeToWaitIe.value.value.pduNum=PDU_NGSetupFailureIEs_Value_TimeToWait;
        timeToWaitIe.value.value.decoded.pdu_TimeToWait=&timeToWait_m;
        timeToWaitIe.value.value.encoded.length=0;
        timeToWaitIe.value.value.encoded.value=NULL;
        nextIe_p=&timeToWaitIe;
	}

    //Cause IE encode
    ProtocolIeType causeIe;
    causeIe.next=nextIe_p;
    causeIe.value.id=id_Cause;
    causeIe.value.criticality=ignore;
    causeIe.value.value.pduNum=PDU_NGSetupFailureIEs_Value_Cause;
    causeIe.value.value.decoded.pdu_Cause=&cause_m;
    causeIe.value.value.encoded.length=0;
    causeIe.value.value.encoded.value=NULL;
    nextIe_p=&causeIe;

    NGSetupFailure msg;
    msg.protocolIEs = nextIe_p;
	
    NGAP_PDU ngApPdu;

    ngApPdu.choice = unsuccessfulOutcome_chosen;
    ngApPdu.u.unsuccessfulOutcome.procedureCode = id_NGSetup;
    ngApPdu.u.unsuccessfulOutcome.criticality = reject;
    ngApPdu.u.unsuccessfulOutcome.value.pduNum = PDU_NGAP_ELEMENTARY_PROCEDURES_UnsuccessfulOutcome_NGSetupFailure;
    ngApPdu.u.unsuccessfulOutcome.value.decoded.pdu_NGSetupFailure = &msg;
    ngApPdu.u.unsuccessfulOutcome.value.encoded.length = 0;
    ngApPdu.u.unsuccessfulOutcome.value.encoded.value = NULL;
	

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
NGSetupFailurePdu::getDecodedOssDataPtrs_v()
{
    ProtocolIeType *currIe_p;
    reset_v();

    ngapPdu_mp = reinterpret_cast<NGAP_PDU *> (ossBufRef_m.value);
    if (ngapPdu_mp == NULL)
    {
        return;
    }

    NGSetupFailure_mp = 
             ngapPdu_mp->u.unsuccessfulOutcome.value.decoded.pdu_NGSetupFailure;             
    if (NGSetupFailure_mp == NULL)
    {
        return;
    }

    currIe_p = NGSetupFailure_mp->protocolIEs;

    while (currIe_p != NULL)
    {
        if (currIe_p->value.id == id_Cause
            && currIe_p->value.value.decoded.pdu_Cause != NULL)
        {
            cause_mp = 
                      currIe_p->value.value.decoded.pdu_Cause;
        }
        else if (currIe_p->value.id == id_TimeToWait
                 && currIe_p->value.value.decoded.pdu_TimeToWait != NULL)
        {
        	timeToWait_mp = currIe_p->value.value.decoded.pdu_TimeToWait;
			timeToWaitPresent_m = true;
        }

        currIe_p = currIe_p->next;
    }
}

NGSetupFailure *
NGSetupFailurePdu::getNGSetupFailureDecodedPtr() const
{
    return NGSetupFailure_mp;
}

Cause* 
NGSetupFailurePdu::getCauseDecodedPtr() const
{
	return cause_mp;
}

Cause* 
NGSetupFailurePdu::getCauseEncodePtr()
{
	return &cause_m;
}

TimeToWait* 
NGSetupFailurePdu::getTimeToWaitDecodedPtr() const
{
	return timeToWait_mp;
}

TimeToWait* 
NGSetupFailurePdu::getTimeToWaitEncodePtr()
{
	timeToWaitPresent_m = true;
	return &timeToWait_m;
}

bool 
NGSetupFailurePdu::isTimeToWaitPresent()
{
	return timeToWaitPresent_m;
}

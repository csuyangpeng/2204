#include <iostream>
#include <string.h>

#include "ueContextReleaseCommandPdu.h"

UeContextReleaseCommandPdu::UeContextReleaseCommandPdu()
{
    reset_v();
}

UeContextReleaseCommandPdu::~UeContextReleaseCommandPdu()
{
}

void UeContextReleaseCommandPdu::reset_v()
{
    ueContextReleaseCommand_mp = NULL;

    //UE NGAP IDs
    memset(&ueNgapIds_m,0,sizeof(ueNgapIds_m));
    ueNgapIds_mp = NULL;

    //Cause
    cause_mp=NULL;
	memset(&cause_m,0,sizeof(cause_m));
}

unsigned int
UeContextReleaseCommandPdu::encodeOssToPer(NgapOssCtxt* ctxt)                   
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
    causeIe.value.value.pduNum=PDU_UEContextReleaseCommand_IEs_Value_Cause;
    causeIe.value.value.decoded.pdu_Cause=&cause_m;
    causeIe.value.value.encoded.length=0;
    causeIe.value.value.encoded.value=NULL;
    nextIe_p=&causeIe;
    
	//RAN_UE_NGAP_ID
    ProtocolIeType ueNgapIdsIe;
    ueNgapIdsIe.next =nextIe_p;
    ueNgapIdsIe.value.id = id_UE_NGAP_IDs;
    ueNgapIdsIe.value.criticality = reject;
    ueNgapIdsIe.value.value.pduNum =PDU_UEContextReleaseCommand_IEs_Value_UE_NGAP_IDs;
    ueNgapIdsIe.value.value.decoded.pdu_UE_NGAP_IDs = &ueNgapIds_m;
    ueNgapIdsIe.value.value.encoded.length = 0;
    ueNgapIdsIe.value.value.encoded.value = NULL;
    nextIe_p = &ueNgapIdsIe;
	
    UEContextReleaseCommand msg;
    msg.protocolIEs = nextIe_p;
	
    NGAP_PDU ngApPdu;

    ngApPdu.choice = initiatingMessage_chosen;
    ngApPdu.u.initiatingMessage.procedureCode = id_UEContextRelease;
    ngApPdu.u.initiatingMessage.criticality = reject;
    ngApPdu.u.initiatingMessage.value.pduNum = PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_UEContextReleaseCommand;
    ngApPdu.u.initiatingMessage.value.decoded.pdu_UEContextReleaseCommand = &msg;
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
UeContextReleaseCommandPdu::getDecodedOssDataPtrs_v()
{
    ProtocolIeType *currIe_p;
    reset_v();

    ngapPdu_mp = reinterpret_cast<NGAP_PDU *> (ossBufRef_m.value);
    if (ngapPdu_mp == NULL)
    {
        return;
    }

    ueContextReleaseCommand_mp = 
             ngapPdu_mp->u.initiatingMessage.value.decoded.pdu_UEContextReleaseCommand;             
    if (ueContextReleaseCommand_mp == NULL)
    {
        return;
    }

    currIe_p = ueContextReleaseCommand_mp->protocolIEs;

    while (currIe_p != NULL)
    {
        if (currIe_p->value.id == id_Cause
            && currIe_p->value.value.decoded.pdu_Cause != NULL)
        {
            cause_mp = 
                      currIe_p->value.value.decoded.pdu_Cause;
        }
        else if (currIe_p->value.id == id_UE_NGAP_IDs
            && currIe_p->value.value.decoded.pdu_UE_NGAP_IDs != NULL)
        {
            ueNgapIds_mp = 
                      currIe_p->value.value.decoded.pdu_UE_NGAP_IDs;
        }
        currIe_p = currIe_p->next;
    }
	
}

UEContextReleaseCommand *
UeContextReleaseCommandPdu::getUeContextReleaseCommandDecodedPtr() const
{
    return ueContextReleaseCommand_mp;
}

UE_NGAP_IDs * 
UeContextReleaseCommandPdu::getUeNgapIdsDecodedPtr() const
{
    return ueNgapIds_mp;
}

UE_NGAP_IDs * 
UeContextReleaseCommandPdu::getUeNgapIdsEncodePtr()
{
    return &ueNgapIds_m;
}

Cause* 
UeContextReleaseCommandPdu::getCauseDecodedPtr() const
{
	return cause_mp;
}

Cause* 
UeContextReleaseCommandPdu::getCauseEncodePtr()
{
	return &cause_m;
}

#include <string.h>
#include <iostream>

#include "initialContextSetupResponsePdu.h"

InitialContextSetupResponsePdu::InitialContextSetupResponsePdu()
{
    reset_v();
}

InitialContextSetupResponsePdu::~InitialContextSetupResponsePdu()
{
}

void
InitialContextSetupResponsePdu::reset_v()
{
    InitialContextSetupResponse_mp = NULL;

    //RAN UE NGAP ID
    ranUeNgapId_m=0;
    ranUeNgapId_mp=NULL;

    //AMF UE NGAP ID
    amfUeNgapId_m=0;
    amfUeNgapId_mp=NULL;

    //PDUSessionResourceSetupListCxtRes
    pDUSessionResourceSetupListCxtRes_mp=NULL;
    pDUSessionResourceSetupListCxtRes_m=NULL;
    numofPDUSessionResourceSetupListItemCxtRes_m=0;
    memset(setupTransfers_ma,0xff,sizeof(setupTransfers_ma));

    for(unsigned int i = 0 ; i < maxnoofPDUSessions_c ; i++)
    {
        pDUSessionResSetupListCxtRes_ma[i].next=NULL;
        pDUSessionResSetupListCxtRes_ma[i].value.bit_mask=0;
        pDUSessionResSetupListCxtRes_ma[i].value.iE_Extensions = NULL;
        pDUSessionResSetupListCxtRes_ma[i].value.pDUSessionID=0;
        pDUSessionResSetupListCxtRes_ma[i].value.pDUSessionResourceSetupResponseTransfer
            .encoded.length=0;
        pDUSessionResSetupListCxtRes_ma[i].value.pDUSessionResourceSetupResponseTransfer
            .encoded.value=setupTransfers_ma[i].data;
        pDUSessionResSetupListCxtRes_ma[i].value.pDUSessionResourceSetupResponseTransfer
            .decoded = NULL;
    }

    //PDUSessionResourceFailedToSetupListCxtRes
    pDUSessionResourceFailedToSetupListCxtRes_mp=NULL;
    pDUSessionResourceFailedToSetupListCxtRes_m=NULL;
    numofPDUSessionResourceFailedToSetupItemCxtRes_m=0;
    memset(failedTransfers_ma,0xff,sizeof(failedTransfers_ma));

    for(unsigned int i = 0 ; i < maxnoofPDUSessions_c ; i++)
    {
        pDUSessionResourceFailedToSetupListCxtRes_ma[i].next=NULL;
        pDUSessionResourceFailedToSetupListCxtRes_ma[i].value.bit_mask=0;
        pDUSessionResourceFailedToSetupListCxtRes_ma[i].value.iE_Extensions = NULL;
        pDUSessionResourceFailedToSetupListCxtRes_ma[i].value.pDUSessionID=0;
        pDUSessionResourceFailedToSetupListCxtRes_ma[i].value.pDUSessionResourceSetupUnsuccessfulTransfer
            .encoded.length=0;
        pDUSessionResourceFailedToSetupListCxtRes_ma[i].value.pDUSessionResourceSetupUnsuccessfulTransfer
            .encoded.value=failedTransfers_ma[i].data;
        pDUSessionResourceFailedToSetupListCxtRes_ma[i].value.pDUSessionResourceSetupUnsuccessfulTransfer
            .decoded = NULL;
    }
}

unsigned int
InitialContextSetupResponsePdu::encodeOssToPer(NgapOssCtxt* ctxt)                   
{
    if(NULL == ctxt)
	{
	    std::cout <<"Invalid NgapOssCtxt!" << std::endl;                  
	    return -1;
    }
	
    ProtocolIeType *nextIe_p = NULL;

    //PDUSessionResourceFailedToSetupListCxtRes
    ProtocolIeType pduSessionResourceFailedToSetupListCxtResIe;
    if(numofPDUSessionResourceFailedToSetupItemCxtRes_m> maxnoofPDUSessions ||
          numofPDUSessionResourceFailedToSetupItemCxtRes_m < 0)
    {
        return defaultFailureCode_c;
    }
    else
    {
        struct PDUSessionResourceFailedToSetupListCxtRes_ *nextPduSessFailedSetupRes_p=NULL;
        for(unsigned int i=0;i<numofPDUSessionResourceFailedToSetupItemCxtRes_m;i++)
        {
            pDUSessionResourceFailedToSetupListCxtRes_ma[i].next=nextPduSessFailedSetupRes_p;
            nextPduSessFailedSetupRes_p=&pDUSessionResourceFailedToSetupListCxtRes_ma[i];
        
	        pDUSessionResourceFailedToSetupListCxtRes_m=nextPduSessFailedSetupRes_p;
	        pduSessionResourceFailedToSetupListCxtResIe.value.value.pduNum=
	            PDU_InitialContextSetupResponseIEs_Value_PDUSessionResourceFailedToSetupListCxtRes;
	        pduSessionResourceFailedToSetupListCxtResIe.value.value.decoded
	            .pdu_PDUSessionResourceFailedToSetupListCxtRes=&pDUSessionResourceFailedToSetupListCxtRes_m;
	        pduSessionResourceFailedToSetupListCxtResIe.value.value.encoded.length = 0;
	        pduSessionResourceFailedToSetupListCxtResIe.value.value.encoded.value = NULL;
	        pduSessionResourceFailedToSetupListCxtResIe.next=nextIe_p;
	        pduSessionResourceFailedToSetupListCxtResIe.value.id=id_PDUSessionResourceFailedToSetupListCxtRes;
	        pduSessionResourceFailedToSetupListCxtResIe.value.criticality=ignore;
	        nextIe_p=&pduSessionResourceFailedToSetupListCxtResIe;
	    }
    }

    //PDUSessionResourceSetupListCxtRes
    ProtocolIeType pduSessionResourceSetupListCxtResIe;
    if(numofPDUSessionResourceSetupListItemCxtRes_m > maxnoofPDUSessions ||
          numofPDUSessionResourceSetupListItemCxtRes_m < 0)
    {
        return defaultFailureCode_c;
    }
    else
    {
        struct PDUSessionResourceSetupListCxtRes_ *nextPDUSessionResourceSetupListCxtRes_p=NULL;
        for(unsigned int i=0;i<numofPDUSessionResourceSetupListItemCxtRes_m;i++)
        {
            pDUSessionResSetupListCxtRes_ma[i].next=nextPDUSessionResourceSetupListCxtRes_p;
            nextPDUSessionResourceSetupListCxtRes_p=&pDUSessionResSetupListCxtRes_ma[i];

    	    pDUSessionResourceSetupListCxtRes_m=nextPDUSessionResourceSetupListCxtRes_p;
	        pduSessionResourceSetupListCxtResIe.value.value.pduNum=
	            PDU_InitialContextSetupResponseIEs_Value_PDUSessionResourceSetupListCxtRes;
	        pduSessionResourceSetupListCxtResIe.value.value.decoded.pdu_PDUSessionResourceSetupListCxtRes
				=&pDUSessionResourceSetupListCxtRes_m;
	        pduSessionResourceSetupListCxtResIe.value.value.encoded.length = 0;
	        pduSessionResourceSetupListCxtResIe.value.value.encoded.value = NULL;
	        pduSessionResourceSetupListCxtResIe.next=nextIe_p;
	        pduSessionResourceSetupListCxtResIe.value.id=id_PDUSessionResourceSetupListCxtRes;
	        pduSessionResourceSetupListCxtResIe.value.criticality=ignore;
	        nextIe_p=&pduSessionResourceSetupListCxtResIe;
		}
    }

    //RAN_UE_NGAP_ID
    ProtocolIeType ranUeNgapIdIe;
    ranUeNgapIdIe.next =nextIe_p;
    ranUeNgapIdIe.value.id = id_RAN_UE_NGAP_ID;
    ranUeNgapIdIe.value.criticality = ignore;
    ranUeNgapIdIe.value.value.pduNum =
        PDU_InitialContextSetupResponseIEs_Value_RAN_UE_NGAP_ID;
    ranUeNgapIdIe.value.value.decoded.pdu_RAN_UE_NGAP_ID = &ranUeNgapId_m;
    ranUeNgapIdIe.value.value.encoded.length = 0;
    ranUeNgapIdIe.value.value.encoded.value = NULL;
    nextIe_p = &ranUeNgapIdIe;

    //AMF_UE_NGAP_ID
    ProtocolIeType amfUeNgapIdIe;
    amfUeNgapIdIe.next =nextIe_p;
    amfUeNgapIdIe.value.id = id_AMF_UE_NGAP_ID;
    amfUeNgapIdIe.value.criticality = ignore;
    amfUeNgapIdIe.value.value.pduNum =
        PDU_InitialContextSetupResponseIEs_Value_AMF_UE_NGAP_ID;
    amfUeNgapIdIe.value.value.decoded.pdu_AMF_UE_NGAP_ID = &amfUeNgapId_m;
    amfUeNgapIdIe.value.value.encoded.length = 0;
    amfUeNgapIdIe.value.value.encoded.value = NULL;
    nextIe_p = &amfUeNgapIdIe;

    //message
    InitialContextSetupResponse msg;
    msg.protocolIEs = nextIe_p;
    
    NGAP_PDU ngApPdu;

    ngApPdu.choice = successfulOutcome_chosen;
    ngApPdu.u.successfulOutcome.procedureCode = id_InitialContextSetup;
    ngApPdu.u.successfulOutcome.criticality = reject;
    ngApPdu.u.successfulOutcome.value.pduNum = 
        PDU_NGAP_ELEMENTARY_PROCEDURES_SuccessfulOutcome_InitialContextSetupResponse;
    ngApPdu.u.successfulOutcome.value.decoded.pdu_InitialContextSetupResponse = &msg;
    ngApPdu.u.successfulOutcome.value.encoded.length = 0;
    ngApPdu.u.successfulOutcome.value.encoded.value = NULL;

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
InitialContextSetupResponsePdu::getDecodedOssDataPtrs_v()
{
    ProtocolIeType *currIe_p;
    reset_v();

    ngapPdu_mp = reinterpret_cast<NGAP_PDU *> (ossBufRef_m.value);
    if (ngapPdu_mp == NULL)
    {
        return;
    }
    
    InitialContextSetupResponse_mp =
            ngapPdu_mp->u.successfulOutcome.value.decoded.pdu_InitialContextSetupResponse;
    
    if (InitialContextSetupResponse_mp == NULL)
    {
        return;
    }

    currIe_p = InitialContextSetupResponse_mp->protocolIEs;

    while (currIe_p != NULL)
    {
        if (currIe_p->value.id == id_RAN_UE_NGAP_ID
            && currIe_p->value.value.decoded.pdu_RAN_UE_NGAP_ID != NULL)
        {
            ranUeNgapId_mp = 
                      currIe_p->value.value.decoded.pdu_RAN_UE_NGAP_ID;
        }
        else if (currIe_p->value.id == id_AMF_UE_NGAP_ID
            && currIe_p->value.value.decoded.pdu_AMF_UE_NGAP_ID != NULL)
        {
            amfUeNgapId_mp = 
                      currIe_p->value.value.decoded.pdu_AMF_UE_NGAP_ID;
        }
        else if (currIe_p->value.id == id_PDUSessionResourceSetupListCxtRes
            && currIe_p->value.value.decoded.pdu_PDUSessionResourceSetupListCxtRes != NULL)
        {
            pDUSessionResourceSetupListCxtRes_mp = 
                      currIe_p->value.value.decoded.pdu_PDUSessionResourceSetupListCxtRes;
            PDUSessionResourceSetupListCxtRes setupListCxtRes_p=NULL;
            // AssociatedQosFlowList associatedQosFlowList_p=NULL;
            // QosFlowList qosFlowList_p=NULL;
            setupListCxtRes_p=*pDUSessionResourceSetupListCxtRes_mp;
            while(setupListCxtRes_p!=NULL)
            {
                numofPDUSessionResourceSetupListItemCxtRes_m++;
                setupListCxtRes_p=setupListCxtRes_p->next;
            }
        }
        else if (currIe_p->value.id == id_PDUSessionResourceFailedToSetupListCxtRes
            && currIe_p->value.value.decoded.pdu_PDUSessionResourceFailedToSetupListCxtRes != NULL)
        {
            pDUSessionResourceFailedToSetupListCxtRes_mp = 
                      currIe_p->value.value.decoded.pdu_PDUSessionResourceFailedToSetupListCxtRes;
            PDUSessionResourceFailedToSetupListCxtRes failedToSetup_p = NULL;
            failedToSetup_p = *pDUSessionResourceFailedToSetupListCxtRes_mp;
            while(failedToSetup_p!=NULL)
            {
                numofPDUSessionResourceFailedToSetupItemCxtRes_m++;
                failedToSetup_p=failedToSetup_p->next;
            }
        }
        currIe_p = currIe_p->next;
    }
}

PDUSessionResourceFailedToSetupListCxtRes * 
InitialContextSetupResponsePdu::getPDUSessionResourceFailedToSetupListCxtResDecodedPtr() const
{
    return pDUSessionResourceFailedToSetupListCxtRes_mp;
}

PDUSessionResourceFailedToSetupListCxtRes * 
InitialContextSetupResponsePdu::getPDUSessionResourceFailedToSetupListCxtResEncodePtr()
{
    return &pDUSessionResourceFailedToSetupListCxtRes_m;
}

PDUSessionResourceSetupItemCxtRes* 
InitialContextSetupResponsePdu::getPDUSessionResourceSetupItemCxtResDecodedPtr(unsigned int index) const
{
    PDUSessionResourceSetupItemCxtRes *item_p = NULL;

    PDUSessionResourceSetupListCxtRes pduSessResSetupList_p = NULL;

    if (pDUSessionResourceSetupListCxtRes_mp != NULL)
    {
        pduSessResSetupList_p = *(pDUSessionResourceSetupListCxtRes_mp);
    }

    if (pduSessResSetupList_p != NULL && (index >= 0 && index < numofPDUSessionResourceSetupListItemCxtRes_m))
    {
        unsigned int i = 0;
        for (; i < index && pduSessResSetupList_p != NULL; i++)
        {
            pduSessResSetupList_p = pduSessResSetupList_p->next;
        }
        if (pduSessResSetupList_p != NULL && i == index)
        {
            item_p = &(pduSessResSetupList_p->value);
        }
    }

    return item_p;
}

PDUSessionResourceSetupItemCxtRes* 
InitialContextSetupResponsePdu::getPDUSessionResourceSetupItemCxtResEncodePtr(unsigned int index)
{
    if (index >= 0 && index < maxnoofPDUSessions)
    {
        return &(pDUSessionResSetupListCxtRes_ma[index].value);
    }
    else
    {
        return NULL;
    }
}

void 
InitialContextSetupResponsePdu::setNumofPDUSessionResourceSetupListItemCxtRes_v(unsigned int num)
{
    numofPDUSessionResourceSetupListItemCxtRes_m=num;
}

unsigned int 
InitialContextSetupResponsePdu::getNumofPDUSessionResourceSetupListItemCxtRes() const
{
    return numofPDUSessionResourceSetupListItemCxtRes_m;
}

PDUSessionResourceSetupListCxtRes * 
InitialContextSetupResponsePdu::getPDUSessionResourceSetupListCxtResDecodedPtr() const
{
    return pDUSessionResourceSetupListCxtRes_mp;
}

PDUSessionResourceSetupListCxtRes * 
InitialContextSetupResponsePdu::getPDUSessionResourceSetupListCxtResEncodePtr()
{
    return &pDUSessionResourceSetupListCxtRes_m;
}

PDUSessionResourceFailedToSetupItemCxtRes* 
InitialContextSetupResponsePdu::getPDUSessionResourceFailedToSetupItemCxtResDecodedPtr(unsigned int index) const
{
    PDUSessionResourceFailedToSetupItemCxtRes *item_p = NULL;

    PDUSessionResourceFailedToSetupListCxtRes pduSessResSetupList_p = NULL;

    if (pDUSessionResourceFailedToSetupListCxtRes_mp != NULL)
    {
        pduSessResSetupList_p = *(pDUSessionResourceFailedToSetupListCxtRes_mp);
    }

    if (pduSessResSetupList_p != NULL && (index >= 0 && index < numofPDUSessionResourceFailedToSetupItemCxtRes_m))
    {
        unsigned int i = 0;
        for (; i < index && pduSessResSetupList_p != NULL; i++)
        {
            pduSessResSetupList_p = pduSessResSetupList_p->next;
        }
        if (pduSessResSetupList_p != NULL && i == index)
        {
            item_p = &(pduSessResSetupList_p->value);
        }
    }

    return item_p;
}

PDUSessionResourceFailedToSetupItemCxtRes* 
InitialContextSetupResponsePdu::getPDUSessionResourceFailedToSetupItemCxtResEncodePtr(unsigned int index)
{
    if (index >= 0 && index < maxnoofPDUSessions)
    {
        return &(pDUSessionResourceFailedToSetupListCxtRes_ma[index].value);
    }
    else
    {
        return NULL;
    }
}

void 
InitialContextSetupResponsePdu::setNumofPDUSessionResourceFailedToSetupListItemCxtRes_v(unsigned int num)
{
    numofPDUSessionResourceFailedToSetupItemCxtRes_m = num;
}

unsigned int 
InitialContextSetupResponsePdu::getNumofPDUSessionResourceFailedToSetupListItemCxtRes() const
{
    return numofPDUSessionResourceFailedToSetupItemCxtRes_m;
}

InitialContextSetupResponse *
InitialContextSetupResponsePdu::getInitialContextSetupResponseDecodedPtr() const
{
    return InitialContextSetupResponse_mp;
}

//----------------------------------------------------------------------------
// This is a public API. Refer to InitialContextSetupResponsePdu.h
//----------------------------------------------------------------------------
RAN_UE_NGAP_ID * 
InitialContextSetupResponsePdu::getRanUeNgapIdDecodedPtr() const
{
    return ranUeNgapId_mp;
}

RAN_UE_NGAP_ID * 
InitialContextSetupResponsePdu::getRanUeNgapIdEncodePtr()
{
    return &ranUeNgapId_m;
}

AMF_UE_NGAP_ID * 
InitialContextSetupResponsePdu::getAmfUeNgapIdDecodedPtr() const
{
    return amfUeNgapId_mp;
}

AMF_UE_NGAP_ID * 
InitialContextSetupResponsePdu::getAmfUeNgapIdEncodePtr()
{
    return &amfUeNgapId_m;
}

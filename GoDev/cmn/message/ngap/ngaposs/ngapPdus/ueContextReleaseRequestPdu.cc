#include <iostream>
#include <string.h>

#include "ueContextReleaseRequestPdu.h"

UeContextReleaseRequestPdu::UeContextReleaseRequestPdu()
{
    reset_v();
}

UeContextReleaseRequestPdu::~UeContextReleaseRequestPdu()
{
}

void UeContextReleaseRequestPdu::reset_v()
{
    UeContextReleaseRequest_mp = NULL;

    //RAN UE NGAP ID
    ranUeNgapId_m = 0;
    ranUeNgapId_mp = NULL;

    //AMF UE NGAP ID
    amfUeNgapId_m = 0;
    amfUeNgapId_mp = NULL;

    //Cause
    cause_mp=NULL;
	memset(&cause_m,0,sizeof(cause_m));

    //PDUSessionResourceListCxtRelReq
    pduSessResListCxtRelReq_mp = NULL;
    pduSessResListCxtRelReq_m = NULL;
    numofPduSessResListItemCxtRelReq_m = 0;
	pduSessResListCxtRelReqPresent_m = false;

    for (unsigned int i = 0; i < maxnoofPDUSessions_c; i++) {
        pduSessResListCxtRelReq_ma[i].next = NULL;
        pduSessResListCxtRelReq_ma[i].value.bit_mask = 0;
        pduSessResListCxtRelReq_ma[i].value.iE_Extensions=NULL;
        pduSessResListCxtRelReq_ma[i].value.pDUSessionID = 0;
    }
}

unsigned int
UeContextReleaseRequestPdu::encodeOssToPer(NgapOssCtxt* ctxt)
{
    if (NULL == ctxt) {
        return -1;
    }

    ProtocolIeType* nextIe_p = NULL;

    //Cause IE encode
    ProtocolIeType causeIe;
    causeIe.next=nextIe_p;
    causeIe.value.id=id_Cause;
    causeIe.value.criticality=ignore;
    causeIe.value.value.pduNum=PDU_UEContextReleaseRequest_IEs_Value_Cause;
    causeIe.value.value.decoded.pdu_Cause=&cause_m;
    causeIe.value.value.encoded.length=0;
    causeIe.value.value.encoded.value=NULL;
    nextIe_p=&causeIe;

    //PDUSessionResourceListCxtRelReq
	ProtocolIeType PDUSessionResourceListCxtRelReqIe;
    if(pduSessResListCxtRelReqPresent_m)
	{
        if(numofPduSessResListItemCxtRelReq_m > maxnoofPDUSessions ||
              numofPduSessResListItemCxtRelReq_m < 0)
        {
            return defaultFailureCode_c;
        }
        else
        {
        	if (numofPduSessResListItemCxtRelReq_m > 0 )
        	{
    
	            struct PDUSessionResourceListCxtRelReq_ *nextPduSessResListCxtRelReq_p = NULL;
	            for(unsigned int i = 0;i<numofPduSessResListItemCxtRelReq_m;i++)
	            {
    
	                pduSessResListCxtRelReq_ma[i].next = nextPduSessResListCxtRelReq_p;
	                nextPduSessResListCxtRelReq_p = &(pduSessResListCxtRelReq_ma[i]);
    
	            }
	            pduSessResListCxtRelReq_m = nextPduSessResListCxtRelReq_p;
	            PDUSessionResourceListCxtRelReqIe.value.value.pduNum=
	                PDU_UEContextReleaseRequest_IEs_Value_PDUSessionResourceListCxtRelReq;
	            PDUSessionResourceListCxtRelReqIe.value.value.decoded.pdu_PDUSessionResourceListCxtRelReq
	    			= &pduSessResListCxtRelReq_m;
	            PDUSessionResourceListCxtRelReqIe.value.value.encoded.length = 0;
	            PDUSessionResourceListCxtRelReqIe.value.value.encoded.value = NULL;
	    		
	            PDUSessionResourceListCxtRelReqIe.next = nextIe_p;
	            PDUSessionResourceListCxtRelReqIe.value.id = id_PDUSessionResourceListCxtRelReq;
	            PDUSessionResourceListCxtRelReqIe.value.criticality = reject;
	            nextIe_p = &PDUSessionResourceListCxtRelReqIe;
        	}
        }
    }

    //RAN_UE_NGAP_ID
    ProtocolIeType ranUeNgapIdIe;
    ranUeNgapIdIe.next = nextIe_p;
    ranUeNgapIdIe.value.id = id_RAN_UE_NGAP_ID;
    ranUeNgapIdIe.value.criticality = reject;
    ranUeNgapIdIe.value.value.pduNum =
        PDU_UEContextReleaseRequest_IEs_Value_RAN_UE_NGAP_ID;
    ranUeNgapIdIe.value.value.decoded.pdu_RAN_UE_NGAP_ID = &ranUeNgapId_m;
    ranUeNgapIdIe.value.value.encoded.length = 0;
    ranUeNgapIdIe.value.value.encoded.value = NULL;
    nextIe_p = &ranUeNgapIdIe;

    //AMF_UE_NGAP_ID
    ProtocolIeType amfUeNgapIdIe;
    amfUeNgapIdIe.next = nextIe_p;
    amfUeNgapIdIe.value.id = id_AMF_UE_NGAP_ID;
    amfUeNgapIdIe.value.criticality = reject;
    amfUeNgapIdIe.value.value.pduNum =
        PDU_UEContextReleaseRequest_IEs_Value_AMF_UE_NGAP_ID;
    amfUeNgapIdIe.value.value.decoded.pdu_AMF_UE_NGAP_ID = &amfUeNgapId_m;
    amfUeNgapIdIe.value.value.encoded.length = 0;
    amfUeNgapIdIe.value.value.encoded.value = NULL;
    nextIe_p = &amfUeNgapIdIe;

    //message
    UEContextReleaseRequest msg;
    msg.protocolIEs = nextIe_p;

    NGAP_PDU ngApPdu;

    ngApPdu.choice = initiatingMessage_chosen;
    ngApPdu.u.initiatingMessage.procedureCode = id_UEContextReleaseRequest;
    ngApPdu.u.initiatingMessage.criticality = ignore;
    ngApPdu.u.initiatingMessage.value.pduNum = 
        PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_UEContextReleaseRequest;
    ngApPdu.u.initiatingMessage.value.decoded.pdu_UEContextReleaseRequest = &msg;
    ngApPdu.u.initiatingMessage.value.encoded.length = 0;
    ngApPdu.u.initiatingMessage.value.encoded.value = NULL;

    //Convert Structures to PER
    OssBuf outputBuf;
    unsigned int result;

    ctxt->reset_v();
    outputBuf = ctxt->getPrimaryCodecBuf();

    result = ctxt->encodeOssToPer(NGAP_PDU_PDU, &ngApPdu, &outputBuf);

    if (result == encodeSuccessCode_c) {
        perBufRef_m = outputBuf;
    }

    return (result);
}

void UeContextReleaseRequestPdu::getDecodedOssDataPtrs_v()
{
    ProtocolIeType* currIe_p;
    reset_v();

    ngapPdu_mp = reinterpret_cast<NGAP_PDU*>(ossBufRef_m.value);
    if (ngapPdu_mp == NULL) {
        return;
    }

    UeContextReleaseRequest_mp = 
        ngapPdu_mp->u.initiatingMessage.value.decoded.pdu_UEContextReleaseRequest;

    if (UeContextReleaseRequest_mp == NULL) {
        return;
    }

    currIe_p = UeContextReleaseRequest_mp->protocolIEs;

    while (currIe_p != NULL) {
        if (currIe_p->value.id == id_RAN_UE_NGAP_ID
            && currIe_p->value.value.decoded.pdu_RAN_UE_NGAP_ID != NULL) {
            ranUeNgapId_mp = currIe_p->value.value.decoded.pdu_RAN_UE_NGAP_ID;
        } else if (currIe_p->value.id == id_AMF_UE_NGAP_ID
            && currIe_p->value.value.decoded.pdu_AMF_UE_NGAP_ID != NULL) {
            amfUeNgapId_mp = currIe_p->value.value.decoded.pdu_AMF_UE_NGAP_ID;
        } else if (currIe_p->value.id == id_Cause
            && currIe_p->value.value.decoded.pdu_Cause != NULL)
        {
            cause_mp = currIe_p->value.value.decoded.pdu_Cause;
        } else if (currIe_p->value.id == id_PDUSessionResourceListCxtRelReq
            && currIe_p->value.value.decoded.pdu_PDUSessionResourceListCxtRelReq != NULL)
        {
            pduSessResListCxtRelReq_mp = 
                      currIe_p->value.value.decoded.pdu_PDUSessionResourceListCxtRelReq;

            PDUSessionResourceListCxtRelReq pduSessResSetup_p=NULL;

            pduSessResSetup_p=*pduSessResListCxtRelReq_mp;
            while(pduSessResSetup_p!=NULL)
            {
            	numofPduSessResListItemCxtRelReq_m++;
                pduSessResSetup_p=pduSessResSetup_p->next;
            }
            pduSessResListCxtRelReqPresent_m = true;
        }
        currIe_p = currIe_p->next;
    }
}

UEContextReleaseRequest*
UeContextReleaseRequestPdu::getUeContextReleaseRequestDecodedPtr() const
{
    return UeContextReleaseRequest_mp;
}

RAN_UE_NGAP_ID*
UeContextReleaseRequestPdu::getRanUeNgapIdDecodedPtr() const
{
    return ranUeNgapId_mp;
}

RAN_UE_NGAP_ID*
UeContextReleaseRequestPdu::getRanUeNgapIdEncodePtr()
{
    return &ranUeNgapId_m;
}

AMF_UE_NGAP_ID*
UeContextReleaseRequestPdu::getAmfUeNgapIdDecodedPtr() const
{
    return amfUeNgapId_mp;
}

AMF_UE_NGAP_ID*
UeContextReleaseRequestPdu::getAmfUeNgapIdEncodePtr()
{
    return &amfUeNgapId_m;
}

Cause* 
UeContextReleaseRequestPdu::getCauseDecodedPtr() const
{
	return cause_mp;
}

Cause* 
UeContextReleaseRequestPdu::getCauseEncodePtr()
{
	return &cause_m;
}

PDUSessionResourceListCxtRelReq*
UeContextReleaseRequestPdu::getPduSessResListCxtRelReqDecodedPtr() const
{
    return pduSessResListCxtRelReq_mp;
}

PDUSessionResourceListCxtRelReq*
UeContextReleaseRequestPdu::getPduSessResListCxtRelReqEncodePtr()
{
    return &pduSessResListCxtRelReq_m;
}

bool 
UeContextReleaseRequestPdu::isPduSessResListCxtRelReqPresent()
{
    return pduSessResListCxtRelReqPresent_m;
}


PDUSessionResourceItemCxtRelReq*
UeContextReleaseRequestPdu::getPduSessResItemCxtRelReqDecodedPtr(unsigned int index) const
{
    PDUSessionResourceItemCxtRelReq* item_p = NULL;

    PDUSessionResourceListCxtRelReq pduSessResListRelReq_p = NULL;

    if (pduSessResListCxtRelReq_mp != NULL) {
        pduSessResListRelReq_p = *(pduSessResListCxtRelReq_mp);
    }

    if (pduSessResListRelReq_p != NULL
        && (index >= 0 && index < numofPduSessResListItemCxtRelReq_m)) {
        unsigned int i = 0;
        for (; i < index && pduSessResListRelReq_p != NULL; i++) {
            pduSessResListRelReq_p = pduSessResListRelReq_p->next;
        }
        if (pduSessResListRelReq_p != NULL && i == index) {
            item_p = &(pduSessResListRelReq_p->value);
        }
    }

    return item_p;
}

PDUSessionResourceItemCxtRelReq*
UeContextReleaseRequestPdu::getPduSessResItemCxtRelReqEncodePtr(unsigned int index)
{
    if (index >= 0 && index < maxnoofPDUSessions) {
		pduSessResListCxtRelReqPresent_m = true;
        return &(pduSessResListCxtRelReq_ma[index].value);
    } else {
        return NULL;
    }
}
void UeContextReleaseRequestPdu::setNumofPduSessResListItemCxtRelReq_v(unsigned int num)
{
    numofPduSessResListItemCxtRelReq_m = num;
}

unsigned int
UeContextReleaseRequestPdu::getNumofPduSessResListItemCxtRelReq() const
{
    return numofPduSessResListItemCxtRelReq_m;
}


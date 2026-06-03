#include <string.h>
#include <iostream>

#include "pduSessionResourceSetupResponsePdu.h"

PDUSessionResourceSetupResponsePdu::PDUSessionResourceSetupResponsePdu()
{
    reset_v();
}

PDUSessionResourceSetupResponsePdu::~PDUSessionResourceSetupResponsePdu()
{
}

void
PDUSessionResourceSetupResponsePdu::reset_v()
{
    pduSessResSetupResp_mp = NULL;

    //RAN UE NGAP ID
    ranUeNgapId_m=0;
    ranUeNgapId_mp=NULL;

    //AMF UE NGAP ID
    amfUeNgapId_m=0;
    amfUeNgapId_mp=NULL;

    //PDUSessionResourceSetupListSUResp
    pduSessResSetupList_mp=NULL;
	pduSessResSetupList_m=NULL;
    numofPduSessResSetupList_m=0;
    memset(setupTransfers_ma, 0xff, sizeof(setupTransfers_ma));
	
    for(unsigned int i = 0 ; i < maxnoofPDUSessions_c ; i++){
        pduSessResSetupList_ma[i].value.bit_mask = 0;
		pduSessResSetupList_ma[i].value.iE_Extensions = NULL;
		pduSessResSetupList_ma[i].value.pDUSessionID = 0;
		pduSessResSetupList_ma[i].value.pDUSessionResourceSetupResponseTransfer.decoded = NULL;
		pduSessResSetupList_ma[i].value.pDUSessionResourceSetupResponseTransfer.encoded.length = 0;
		pduSessResSetupList_ma[i].value.pDUSessionResourceSetupResponseTransfer.encoded.value = setupTransfers_ma[i].data;
    }

	//PDUSessionResourceFailedToSetupListSUResp
	pduSessResFailedSetupList_mp = NULL;
	pduSessResFailedSetupList_m = NULL;
	numofpduSessResFailedSetupList_m = 0;
	memset(failedTransfers_ma, 0xff, sizeof(failedTransfers_ma));
	for(int i=0;i<maxnoofPDUSessions_c;i++){
		pduSessResFailedSetupList_ma[i].value.bit_mask = 0;
		pduSessResFailedSetupList_ma[i].value.iE_Extensions = NULL;
		pduSessResFailedSetupList_ma[i].value.pDUSessionID = 0;
		pduSessResFailedSetupList_ma[i].value.pDUSessionResourceSetupUnsuccessfulTransfer.decoded = NULL;
		pduSessResFailedSetupList_ma[i].value.pDUSessionResourceSetupUnsuccessfulTransfer.encoded.length = 0;
		pduSessResFailedSetupList_ma[i].value.pDUSessionResourceSetupUnsuccessfulTransfer.encoded.value = failedTransfers_ma[i].data;
	}
}

unsigned int
PDUSessionResourceSetupResponsePdu::encodeOssToPer(NgapOssCtxt* ctxt)                   
{
    if(NULL == ctxt)
	{         
	    return -1;
    }
	
    ProtocolIeType *nextIe_p = NULL;
	
    //PDUSessionResourceFailedToModifyListSURes
    ProtocolIeType pduSessResFailedSetupListRespIe;
    if(numofpduSessResFailedSetupList_m > maxnoofPDUSessions)
    {
        return defaultFailureCode_c;
    }
    else if (numofpduSessResFailedSetupList_m >= 1)
    {
		struct PDUSessionResourceFailedToSetupListSURes_ * nextPtr = NULL;
		for (unsigned int i = 0; i < numofpduSessResFailedSetupList_m;i++){
			pduSessResFailedSetupList_ma[i].next = nextPtr;
			nextPtr = &pduSessResFailedSetupList_ma[i];
		}
	    pduSessResFailedSetupList_m = nextPtr;
		
        pduSessResFailedSetupListRespIe.value.value.pduNum =
            PDU_PDUSessionResourceSetupResponseIEs_Value_PDUSessionResourceFailedToSetupListSURes;
        pduSessResFailedSetupListRespIe.value.value.decoded.pdu_PDUSessionResourceFailedToSetupListSURes = 
			&pduSessResFailedSetupList_m;
        pduSessResFailedSetupListRespIe.value.value.encoded.length = 0;
        pduSessResFailedSetupListRespIe.value.value.encoded.value = NULL;
		
		pduSessResFailedSetupListRespIe.next=nextIe_p;
		pduSessResFailedSetupListRespIe.value.id=id_PDUSessionResourceFailedToSetupListSURes;
		pduSessResFailedSetupListRespIe.value.criticality=ignore;
		nextIe_p=&pduSessResFailedSetupListRespIe;
    }

    //PDUSessionResourceSetupListSURes
    ProtocolIeType pduSessResSetupListRespIe;
    if(numofPduSessResSetupList_m > maxnoofPDUSessions)
    {
        return defaultFailureCode_c;
    }
    else if(numofPduSessResSetupList_m >= 1 )
    {
        struct PDUSessionResourceSetupListSURes_ * nextPtr = NULL;
		for (unsigned int i = 0; i< numofPduSessResSetupList_m; i++){
			pduSessResSetupList_ma[i].next = nextPtr;
		    nextPtr = &pduSessResSetupList_ma[i];
		}
		pduSessResSetupList_m = nextPtr;

        pduSessResSetupListRespIe.value.value.pduNum =
            PDU_PDUSessionResourceSetupResponseIEs_Value_PDUSessionResourceSetupListSURes;
        pduSessResSetupListRespIe.value.value.decoded.pdu_PDUSessionResourceSetupListSURes =
            &pduSessResSetupList_m;
        pduSessResSetupListRespIe.value.value.encoded.length = 0;
        pduSessResSetupListRespIe.value.value.encoded.value = NULL;

		pduSessResSetupListRespIe.next=nextIe_p;
	    pduSessResSetupListRespIe.value.id=id_PDUSessionResourceSetupListSURes;
	    pduSessResSetupListRespIe.value.criticality=ignore;
	    nextIe_p=&pduSessResSetupListRespIe;
    }

    //RAN_UE_NGAP_ID
    ProtocolIeType ranUeNgapIdIe;
    ranUeNgapIdIe.next =nextIe_p;
    ranUeNgapIdIe.value.id = id_RAN_UE_NGAP_ID;
    ranUeNgapIdIe.value.criticality = ignore;
    ranUeNgapIdIe.value.value.pduNum =PDU_PDUSessionResourceSetupResponseIEs_Value_RAN_UE_NGAP_ID;
    ranUeNgapIdIe.value.value.decoded.pdu_RAN_UE_NGAP_ID = &ranUeNgapId_m;
    ranUeNgapIdIe.value.value.encoded.length = 0;
    ranUeNgapIdIe.value.value.encoded.value = NULL;
    nextIe_p = &ranUeNgapIdIe;

    //AMF_UE_NGAP_ID
    ProtocolIeType amfUeNgapIdIe;
    amfUeNgapIdIe.next =nextIe_p;
    amfUeNgapIdIe.value.id = id_AMF_UE_NGAP_ID;
    amfUeNgapIdIe.value.criticality = ignore;
    amfUeNgapIdIe.value.value.pduNum =PDU_PDUSessionResourceSetupResponseIEs_Value_AMF_UE_NGAP_ID;
    amfUeNgapIdIe.value.value.decoded.pdu_AMF_UE_NGAP_ID = &amfUeNgapId_m;
    amfUeNgapIdIe.value.value.encoded.length = 0;
    amfUeNgapIdIe.value.value.encoded.value = NULL;
    nextIe_p = &amfUeNgapIdIe;

    //message
    PDUSessionResourceSetupResponse msg;
    msg.protocolIEs = nextIe_p;
    
    NGAP_PDU ngApPdu;

    ngApPdu.choice = successfulOutcome_chosen;
    ngApPdu.u.successfulOutcome.procedureCode = id_PDUSessionResourceSetup;
    ngApPdu.u.successfulOutcome.criticality = reject;
    ngApPdu.u.successfulOutcome.value.pduNum = 
        PDU_NGAP_ELEMENTARY_PROCEDURES_SuccessfulOutcome_PDUSessionResourceSetupResponse;
    ngApPdu.u.successfulOutcome.value.decoded.pdu_PDUSessionResourceSetupResponse = &msg;
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
PDUSessionResourceSetupResponsePdu::getDecodedOssDataPtrs_v()
{
    ProtocolIeType *currIe_p;
    reset_v();

    ngapPdu_mp = reinterpret_cast<NGAP_PDU *> (ossBufRef_m.value);
    if (ngapPdu_mp == NULL)
    {
        return;
    }
    
    pduSessResSetupResp_mp =
            ngapPdu_mp->u.successfulOutcome.value.decoded.pdu_PDUSessionResourceSetupResponse;
    
    if (pduSessResSetupResp_mp == NULL)
    {
        return;
    }

    currIe_p = pduSessResSetupResp_mp->protocolIEs;

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
        else if (currIe_p->value.id == id_PDUSessionResourceSetupListSURes
            && currIe_p->value.value.decoded.pdu_PDUSessionResourceSetupListSURes != NULL)
        {
            pduSessResSetupList_mp =
                currIe_p->value.value.decoded.pdu_PDUSessionResourceSetupListSURes;

			PDUSessionResourceSetupListSURes nextPtr = NULL;
            nextPtr = *pduSessResSetupList_mp;
            while(nextPtr != NULL)
            {
                numofPduSessResSetupList_m++;
                nextPtr = nextPtr->next;
            }
        }
        else if (currIe_p->value.id == id_PDUSessionResourceFailedToSetupListSURes
            && currIe_p->value.value.decoded.pdu_PDUSessionResourceFailedToSetupListSURes != NULL)
        {
            pduSessResFailedSetupList_mp=
                currIe_p->value.value.decoded.pdu_PDUSessionResourceFailedToSetupListSURes;

            PDUSessionResourceFailedToSetupListSURes nextPtr = NULL;
            nextPtr = *pduSessResFailedSetupList_mp;
            while(nextPtr != NULL)
            {
                numofpduSessResFailedSetupList_m++;
                nextPtr=nextPtr->next;
            }
        }
        currIe_p = currIe_p->next;
    }
}

PDUSessionResourceSetupResponse *
PDUSessionResourceSetupResponsePdu::getPduSessResSetupRespDecodedPtr() const
{
    return pduSessResSetupResp_mp;
}

RAN_UE_NGAP_ID * 
PDUSessionResourceSetupResponsePdu::getRanUeNgapIdDecodedPtr() const
{
    return ranUeNgapId_mp;
}

RAN_UE_NGAP_ID * 
PDUSessionResourceSetupResponsePdu::getRanUeNgapIdEncodePtr()
{
    return &ranUeNgapId_m;
}

AMF_UE_NGAP_ID * 
PDUSessionResourceSetupResponsePdu::getAmfUeNgapIdDecodedPtr() const
{
    return amfUeNgapId_mp;
}

AMF_UE_NGAP_ID * 
PDUSessionResourceSetupResponsePdu::getAmfUeNgapIdEncodePtr()
{
    return &amfUeNgapId_m;
}

PDUSessionResourceSetupListSURes * 
PDUSessionResourceSetupResponsePdu::getPduSessResSetupListDecodedPtr() const
{
    return pduSessResSetupList_mp;
}

PDUSessionResourceSetupListSURes * 
PDUSessionResourceSetupResponsePdu::getPduSessResSetupListEncodePtr()
{
    return &pduSessResSetupList_m;
}

PDUSessionResourceSetupItemSURes*
PDUSessionResourceSetupResponsePdu::getPduSessResSetupItemDecodedPtr(unsigned int index) const
{
    PDUSessionResourceSetupItemSURes *item_p = NULL;

    PDUSessionResourceSetupListSURes pduSessResSetupList_p = NULL;

    if (pduSessResSetupList_mp != NULL)
    {
        pduSessResSetupList_p = *(pduSessResSetupList_mp);
    }

    if (pduSessResSetupList_p != NULL && (index >= 0 && index < numofPduSessResSetupList_m))
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

PDUSessionResourceSetupItemSURes*
PDUSessionResourceSetupResponsePdu::getPduSessResSetupItemEncodePtr(unsigned int index)
{
    if (index >= 0 && index < maxnoofPDUSessions)
    {
        return &(pduSessResSetupList_ma[index].value);
    }
    else
    {
        return NULL;
    }
}

void 
PDUSessionResourceSetupResponsePdu::setNumofSetupListItem_v(unsigned int num)
{
	numofPduSessResSetupList_m = num;
}

unsigned int 
PDUSessionResourceSetupResponsePdu::getNumofSetupListItem() const
{
   return numofPduSessResSetupList_m;
}

PDUSessionResourceFailedToSetupListSURes * 
PDUSessionResourceSetupResponsePdu::getPduSessResFailedSetupListpDecodedPtr() const
{
    return pduSessResFailedSetupList_mp;
}

PDUSessionResourceFailedToSetupListSURes * 
PDUSessionResourceSetupResponsePdu::getPduSessResFailedSetupListEncodePtr()
{
    return &pduSessResFailedSetupList_m;
}

PDUSessionResourceFailedToSetupItemSURes*
PDUSessionResourceSetupResponsePdu::getPduSessResFailedSetupItemDecodedPtr(unsigned int index) const
{
    PDUSessionResourceFailedToSetupItemSURes *item_p = NULL;

    PDUSessionResourceFailedToSetupListSURes pduSessResFailedToSetupList_p = NULL;

    if (pduSessResFailedSetupList_mp != NULL)
    {
        pduSessResFailedToSetupList_p = *(pduSessResFailedSetupList_mp);
    }

    if (pduSessResFailedToSetupList_p != NULL && 
        (index >= 0 && index < numofpduSessResFailedSetupList_m))
    {
        unsigned int i = 0;
        for (; i < index && pduSessResFailedToSetupList_p != NULL; i++)
        {
            pduSessResFailedToSetupList_p = pduSessResFailedToSetupList_p->next;
        }
        if (pduSessResFailedToSetupList_p != NULL && i == index)
        {
            item_p = &(pduSessResFailedToSetupList_p->value);
        }
    }

    return item_p;
}

PDUSessionResourceFailedToSetupItemSURes*
PDUSessionResourceSetupResponsePdu::getPduSessResFailedSetupItemEncodePtr(unsigned int index)
{
    if (index >= 0 && index < maxnoofPDUSessions)
    {
        return &(pduSessResFailedSetupList_ma[index].value);
    }
    else
    {
        return NULL;
    }
}

void 
PDUSessionResourceSetupResponsePdu::setNumofFailedSetupItem_v(unsigned int num)
{
	numofpduSessResFailedSetupList_m = num;
}

unsigned int 
PDUSessionResourceSetupResponsePdu::getNumofFailedSetupItem() const
{
   return numofpduSessResFailedSetupList_m;
}


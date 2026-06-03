#include <iostream>
#include <string.h>

#include "initialContextSetupFailurePdu.h"

InitialContextSetupFailurePdu::InitialContextSetupFailurePdu()
{
    reset_v();
}

InitialContextSetupFailurePdu::~InitialContextSetupFailurePdu()
{
}

void InitialContextSetupFailurePdu::reset_v()
{
    InitialContextSetupFailure_mp = NULL;

    //RAN UE NGAP ID
    ranUeNgapId_m = 0;
    ranUeNgapId_mp = NULL;

    //AMF UE NGAP ID
    amfUeNgapId_m = 0;
    amfUeNgapId_mp = NULL;

    //Cause
    cause_mp=NULL;
	memset(&cause_m,0,sizeof(cause_m));

    //PDUSessionResourceFailedToSetupListCxtFail
    pDUSessionResourceSetupListCxtFail_mp = NULL;
    pDUSessionResourceSetupListCxtFail_m = NULL;
    numofPDUSessionResourceSetupListItemCxtFail_m = 0;

    memset(transfers_ma, 0xff, sizeof(transfers_ma));

    for (unsigned int i = 0; i < maxnoofPDUSessions_c; i++) {
        PDUSessionResSetupListCxtFail_ma[i].next = NULL;
        PDUSessionResSetupListCxtFail_ma[i].value.bit_mask = 0;
        PDUSessionResSetupListCxtFail_ma[i].value.iE_Extensions=NULL;
        PDUSessionResSetupListCxtFail_ma[i].value.pDUSessionID = 0;
        PDUSessionResSetupListCxtFail_ma[i].value.pDUSessionResourceSetupUnsuccessfulTransfer.encoded.length = 0;
        PDUSessionResSetupListCxtFail_ma[i].value.pDUSessionResourceSetupUnsuccessfulTransfer.encoded.value = transfers_ma[i].data;
        PDUSessionResSetupListCxtFail_ma[i].value.pDUSessionResourceSetupUnsuccessfulTransfer.decoded = NULL;
    }
}

unsigned int
InitialContextSetupFailurePdu::encodeOssToPer(NgapOssCtxt* ctxt)
{
    if (NULL == ctxt) {
        std::cout << "Invalid NgapOssCtxt!" << std::endl;
        return -1;
    }

    ProtocolIeType* nextIe_p = NULL;

    //Cause IE encode
    ProtocolIeType causeIe;
    causeIe.next=nextIe_p;
    causeIe.value.id=id_Cause;
    causeIe.value.criticality=ignore;
    causeIe.value.value.pduNum=PDU_InitialContextSetupFailureIEs_Value_Cause;
    causeIe.value.value.decoded.pdu_Cause=&cause_m;
    causeIe.value.value.encoded.length=0;
    causeIe.value.value.encoded.value=NULL;
    nextIe_p=&causeIe;
	
	//PDUSessionResourceFailedToSetupListCxtFail
    ProtocolIeType PDUSessionResourceSetupListCxtFailIe;
    if(numofPDUSessionResourceSetupListItemCxtFail_m > maxnoofPDUSessions ||
          numofPDUSessionResourceSetupListItemCxtFail_m < 0)
    {
        return defaultFailureCode_c;
    }
    else
    {
        struct PDUSessionResourceFailedToSetupListCxtFail_ *nextPDUSessionResourceSetupListCxtFail_p=NULL;
        for(unsigned int i=0;i<numofPDUSessionResourceSetupListItemCxtFail_m;i++)
        {
            PDUSessionResSetupListCxtFail_ma[i].next=nextPDUSessionResourceSetupListCxtFail_p;
            nextPDUSessionResourceSetupListCxtFail_p=&(PDUSessionResSetupListCxtFail_ma[i]);
        }

        if(nextPDUSessionResourceSetupListCxtFail_p != NULL)
        {
            pDUSessionResourceSetupListCxtFail_m=nextPDUSessionResourceSetupListCxtFail_p;

            PDUSessionResourceSetupListCxtFailIe.value.value.pduNum=
                PDU_InitialContextSetupFailureIEs_Value_PDUSessionResourceFailedToSetupListCxtFail;
            PDUSessionResourceSetupListCxtFailIe.value.value.decoded.pdu_PDUSessionResourceFailedToSetupListCxtFail
                =&pDUSessionResourceSetupListCxtFail_m;
            PDUSessionResourceSetupListCxtFailIe.value.value.encoded.length = 0;
            PDUSessionResourceSetupListCxtFailIe.value.value.encoded.value = NULL;
            PDUSessionResourceSetupListCxtFailIe.next=nextIe_p;
            PDUSessionResourceSetupListCxtFailIe.value.id=id_PDUSessionResourceFailedToSetupListCxtFail;
            PDUSessionResourceSetupListCxtFailIe.value.criticality=ignore;
            nextIe_p=&PDUSessionResourceSetupListCxtFailIe;
        }
    }
	
    //RAN_UE_NGAP_ID
    ProtocolIeType ranUeNgapIdIe;
    ranUeNgapIdIe.next = nextIe_p;
    ranUeNgapIdIe.value.id = id_RAN_UE_NGAP_ID;
    ranUeNgapIdIe.value.criticality = ignore;
    ranUeNgapIdIe.value.value.pduNum = 
        PDU_InitialContextSetupFailureIEs_Value_RAN_UE_NGAP_ID;
    ranUeNgapIdIe.value.value.decoded.pdu_RAN_UE_NGAP_ID = &ranUeNgapId_m;
    ranUeNgapIdIe.value.value.encoded.length = 0;
    ranUeNgapIdIe.value.value.encoded.value = NULL;
    nextIe_p = &ranUeNgapIdIe;

    //AMF_UE_NGAP_ID
    ProtocolIeType amfUeNgapIdIe;
    amfUeNgapIdIe.next = nextIe_p;
    amfUeNgapIdIe.value.id = id_AMF_UE_NGAP_ID;
    amfUeNgapIdIe.value.criticality = ignore;
    amfUeNgapIdIe.value.value.pduNum = 
        PDU_InitialContextSetupFailureIEs_Value_AMF_UE_NGAP_ID;
    amfUeNgapIdIe.value.value.decoded.pdu_AMF_UE_NGAP_ID = &amfUeNgapId_m;
    amfUeNgapIdIe.value.value.encoded.length = 0;
    amfUeNgapIdIe.value.value.encoded.value = NULL;
    nextIe_p = &amfUeNgapIdIe;

    //message
    InitialContextSetupFailure msg;
    msg.protocolIEs = nextIe_p;

    NGAP_PDU ngApPdu;

    ngApPdu.choice = unsuccessfulOutcome_chosen;
    ngApPdu.u.unsuccessfulOutcome.procedureCode = id_InitialContextSetup;
    ngApPdu.u.unsuccessfulOutcome.criticality = reject;
    ngApPdu.u.unsuccessfulOutcome.value.pduNum = 
        PDU_NGAP_ELEMENTARY_PROCEDURES_UnsuccessfulOutcome_InitialContextSetupFailure;
    ngApPdu.u.unsuccessfulOutcome.value.decoded.pdu_InitialContextSetupFailure = &msg;
    ngApPdu.u.unsuccessfulOutcome.value.encoded.length = 0;
    ngApPdu.u.unsuccessfulOutcome.value.encoded.value = NULL;

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

void InitialContextSetupFailurePdu::getDecodedOssDataPtrs_v()
{
    ProtocolIeType* currIe_p;
    reset_v();

    ngapPdu_mp = reinterpret_cast<NGAP_PDU*>(ossBufRef_m.value);
    if (ngapPdu_mp == NULL) {
        return;
    }

    InitialContextSetupFailure_mp = 
        ngapPdu_mp->u.unsuccessfulOutcome.value.decoded.pdu_InitialContextSetupFailure;

    if (InitialContextSetupFailure_mp == NULL) {
        return;
    }

    currIe_p = InitialContextSetupFailure_mp->protocolIEs;

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
            cause_mp = 
                      currIe_p->value.value.decoded.pdu_Cause;
        } else if (currIe_p->value.id == id_PDUSessionResourceFailedToSetupListCxtFail
            && currIe_p->value.value.decoded.pdu_PDUSessionResourceFailedToSetupListCxtFail != NULL)
        {
            pDUSessionResourceSetupListCxtFail_mp = 
                      currIe_p->value.value.decoded.pdu_PDUSessionResourceFailedToSetupListCxtFail;

            PDUSessionResourceFailedToSetupListCxtFail pduSessResSetupFail_p=NULL;

            pduSessResSetupFail_p=*pDUSessionResourceSetupListCxtFail_mp;
            while(pduSessResSetupFail_p!=NULL)
            {
                numofPDUSessionResourceSetupListItemCxtFail_m ++;
                pduSessResSetupFail_p=pduSessResSetupFail_p->next;
            }
        }
        currIe_p = currIe_p->next;
    }
}

//----------------------------------------------------------------------------
// This is a public API. Refer to ngApOssInitialContextSetupFailurePdu.h
//----------------------------------------------------------------------------
InitialContextSetupFailure*
InitialContextSetupFailurePdu::getInitialContextSetupFailureDecodedPtr() const
{
    return InitialContextSetupFailure_mp;
}

//----------------------------------------------------------------------------
// This is a public API. Refer to ngApOssInitialContextSetupFailurePdu.h
//----------------------------------------------------------------------------
RAN_UE_NGAP_ID*
InitialContextSetupFailurePdu::getRanUeNgapIdDecodedPtr() const
{
    return ranUeNgapId_mp;
}

RAN_UE_NGAP_ID*
InitialContextSetupFailurePdu::getRanUeNgapIdEncodePtr()
{
    return &ranUeNgapId_m;
}

//----------------------------------------------------------------------------
// This is a public API. Refer to ngApOssInitialContextSetupFailurePdu.h
//----------------------------------------------------------------------------
AMF_UE_NGAP_ID*
InitialContextSetupFailurePdu::getAmfUeNgapIdDecodedPtr() const
{
    return amfUeNgapId_mp;
}

AMF_UE_NGAP_ID*
InitialContextSetupFailurePdu::getAmfUeNgapIdEncodePtr()
{
    return &amfUeNgapId_m;
}

//----------------------------------------------------------------------------
// This is a public API. Refer to ngApOssInitialContextSetupFailurePdu.h
//----------------------------------------------------------------------------
//PDUSessionResourceFailedToSetupListCxtFail
PDUSessionResourceFailedToSetupListCxtFail*
InitialContextSetupFailurePdu::getPDUSessionResourceSetupListCxtFailDecodedPtr() const
{
    return pDUSessionResourceSetupListCxtFail_mp;
}

PDUSessionResourceFailedToSetupListCxtFail*
InitialContextSetupFailurePdu::getPDUSessionResourceSetupListCxtFailEncodePtr()
{
    return &pDUSessionResourceSetupListCxtFail_m;
}

PDUSessionResourceFailedToSetupItemCxtFail*
InitialContextSetupFailurePdu::getPDUSessionResourceSetupItemCxtFailDecodedPtr(unsigned int index) const
{
    PDUSessionResourceFailedToSetupItemCxtFail* item_p = NULL;

    PDUSessionResourceFailedToSetupListCxtFail pduSessResSetupListFail_p = NULL;

    if (pDUSessionResourceSetupListCxtFail_mp != NULL) {
        pduSessResSetupListFail_p = *(pDUSessionResourceSetupListCxtFail_mp);
    }

    if (pduSessResSetupListFail_p != NULL
        && (index >= 0 && index < numofPDUSessionResourceSetupListItemCxtFail_m)) {
        unsigned int i = 0;
        for (; i < index && pduSessResSetupListFail_p != NULL; i++) {
            pduSessResSetupListFail_p = pduSessResSetupListFail_p->next;
        }
        if (pduSessResSetupListFail_p != NULL && i == index) {
            item_p = &(pduSessResSetupListFail_p->value);
        }
    }

    return item_p;
}

PDUSessionResourceFailedToSetupItemCxtFail*
InitialContextSetupFailurePdu::getPDUSessionResourceSetupItemCxtFailEncodePtr(unsigned int index)
{
    if (index >= 0 && index < maxnoofPDUSessions) {
        return &(PDUSessionResSetupListCxtFail_ma[index].value);
    } else {
        return NULL;
    }
}
void InitialContextSetupFailurePdu::setNumofPDUSessionResourceSetupListItemCxtFail_v(unsigned int num)
{
    numofPDUSessionResourceSetupListItemCxtFail_m = num;
}

unsigned int
InitialContextSetupFailurePdu::getNumofPDUSessionResourceSetupListItemCxtFail() const
{
    return numofPDUSessionResourceSetupListItemCxtFail_m;
}

//----------------------------------------------------------------------------
// This is a public API. Refer to ngApOssInitialContextSetupFailurePdu.h
//----------------------------------------------------------------------------
Cause* 
InitialContextSetupFailurePdu::getCauseDecodedPtr() const
{
	return cause_mp;
}

Cause* 
InitialContextSetupFailurePdu::getCauseEncodePtr()
{
	return &cause_m;
}
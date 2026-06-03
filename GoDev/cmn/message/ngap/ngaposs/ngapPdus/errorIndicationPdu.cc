#include <iostream>
#include <string.h>

#include "errorIndicationPdu.h"

ErrorIndicationPdu::ErrorIndicationPdu()
{
    reset_v();
}

ErrorIndicationPdu::~ErrorIndicationPdu()
{
}

void ErrorIndicationPdu::reset_v()
{
    errorIndication_mp = NULL;

    //RAN UE NGAP ID
    ranUeNgapId_m = 0;
    ranUeNgapId_mp = NULL;
    ranUeNgapIdPresent_m = false;

    //AMF UE NGAP ID
    amfUeNgapId_m = 0;
    amfUeNgapId_mp = NULL;
    amfUeNgapIdPresent_m = false;

    //Cause
    cause_mp=NULL;
    memset(&cause_m,0,sizeof(cause_m));
    causePresent_m = false;
}

unsigned int
ErrorIndicationPdu::encodeOssToPer(NgapOssCtxt* ctxt)
{
    if (NULL == ctxt) {
        std::cout << "Invalid NgapOssCtxt!" << std::endl;
        return -1;
    }

    ProtocolIeType* nextIe_p = NULL;

    //Cause IE encode
    ProtocolIeType causeIe;
    if(causePresent_m)
    {
        causeIe.next=nextIe_p;
        causeIe.value.id=id_Cause;
        causeIe.value.criticality = ignore;
        causeIe.value.value.pduNum = PDU_UEContextModificationFailureIEs_Value_Cause;
        causeIe.value.value.decoded.pdu_Cause=&cause_m;
        causeIe.value.value.encoded.length=0;
        causeIe.value.value.encoded.value=NULL;
        nextIe_p=&causeIe;
    }

    //RAN_UE_NGAP_ID
    ProtocolIeType ranUeNgapIdIe;
    if(ranUeNgapIdPresent_m)
    {
        ranUeNgapIdIe.next = nextIe_p;
        ranUeNgapIdIe.value.id = id_RAN_UE_NGAP_ID;
        ranUeNgapIdIe.value.criticality = ignore;
        ranUeNgapIdIe.value.value.pduNum =
            PDU_UEContextModificationFailureIEs_Value_RAN_UE_NGAP_ID;
        ranUeNgapIdIe.value.value.decoded.pdu_RAN_UE_NGAP_ID = &ranUeNgapId_m;
        ranUeNgapIdIe.value.value.encoded.length = 0;
        ranUeNgapIdIe.value.value.encoded.value = NULL;
        nextIe_p = &ranUeNgapIdIe;
    }

    //AMF_UE_NGAP_ID
    ProtocolIeType amfUeNgapIdIe;
    if(amfUeNgapIdPresent_m)
    {
        amfUeNgapIdIe.next = nextIe_p;
        amfUeNgapIdIe.value.id = id_AMF_UE_NGAP_ID;
        amfUeNgapIdIe.value.criticality = ignore;
        amfUeNgapIdIe.value.value.pduNum =
            PDU_UEContextModificationFailureIEs_Value_AMF_UE_NGAP_ID;
        amfUeNgapIdIe.value.value.decoded.pdu_AMF_UE_NGAP_ID = &amfUeNgapId_m;
        amfUeNgapIdIe.value.value.encoded.length = 0;
        amfUeNgapIdIe.value.value.encoded.value = NULL;
        nextIe_p = &amfUeNgapIdIe;
    }

    //message
    ErrorIndication msg;
    msg.protocolIEs = nextIe_p;

    NGAP_PDU ngApPdu;

    ngApPdu.choice = initiatingMessage_chosen;
    ngApPdu.u.initiatingMessage.procedureCode = id_ErrorIndication;
    ngApPdu.u.initiatingMessage.criticality = ignore;
    ngApPdu.u.initiatingMessage.value.pduNum = 
        PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_ErrorIndication;
    ngApPdu.u.initiatingMessage.value.decoded.pdu_ErrorIndication = &msg;
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

void ErrorIndicationPdu::getDecodedOssDataPtrs_v()
{
    ProtocolIeType* currIe_p;
    reset_v();

    ngapPdu_mp = reinterpret_cast<NGAP_PDU*>(ossBufRef_m.value);
    if (ngapPdu_mp == NULL) {
        return;
    }

    errorIndication_mp = ngapPdu_mp->u.initiatingMessage.value.decoded.pdu_ErrorIndication;

    if (errorIndication_mp == NULL) {
        return;
    }

    currIe_p = errorIndication_mp->protocolIEs;

    while (currIe_p != NULL) {
        if (currIe_p->value.id == id_RAN_UE_NGAP_ID
            && currIe_p->value.value.decoded.pdu_RAN_UE_NGAP_ID != NULL) {
            ranUeNgapId_mp = currIe_p->value.value.decoded.pdu_RAN_UE_NGAP_ID;
            ranUeNgapIdPresent_m = true;
        } else if (currIe_p->value.id == id_AMF_UE_NGAP_ID
            && currIe_p->value.value.decoded.pdu_AMF_UE_NGAP_ID != NULL) {
            amfUeNgapId_mp = currIe_p->value.value.decoded.pdu_AMF_UE_NGAP_ID;
            amfUeNgapIdPresent_m = true;
        } else if (currIe_p->value.id == id_Cause
            && currIe_p->value.value.decoded.pdu_Cause != NULL) {
            cause_mp = currIe_p->value.value.decoded.pdu_Cause;
            causePresent_m = true;
        }
        currIe_p = currIe_p->next;
    }
}

ErrorIndication*
ErrorIndicationPdu::getErrorIndicationDecodedPtr() const
{
    return errorIndication_mp;
}

RAN_UE_NGAP_ID*
ErrorIndicationPdu::getRanUeNgapIdDecodedPtr() const
{
    return ranUeNgapId_mp;
}

RAN_UE_NGAP_ID*
ErrorIndicationPdu::getRanUeNgapIdEncodePtr()
{
    ranUeNgapIdPresent_m = true;
    return &ranUeNgapId_m;
}

bool 
ErrorIndicationPdu::isRanUeNgapIdPresent()
{
    return ranUeNgapIdPresent_m;
}

AMF_UE_NGAP_ID*
ErrorIndicationPdu::getAmfUeNgapIdDecodedPtr() const
{
    return amfUeNgapId_mp;
}

AMF_UE_NGAP_ID*
ErrorIndicationPdu::getAmfUeNgapIdEncodePtr()
{
    amfUeNgapIdPresent_m = true;
    return &amfUeNgapId_m;
}

bool 
ErrorIndicationPdu::isAmfUeNgapIdPresent()
{
    return amfUeNgapIdPresent_m;
}

Cause* 
ErrorIndicationPdu::getCauseDecodedPtr() const
{
    return cause_mp;
}

Cause* 
ErrorIndicationPdu::getCauseEncodePtr()
{
    causePresent_m = true;
    return &cause_m;
}

bool 
ErrorIndicationPdu::isCausePresent()
{
    return causePresent_m;
}

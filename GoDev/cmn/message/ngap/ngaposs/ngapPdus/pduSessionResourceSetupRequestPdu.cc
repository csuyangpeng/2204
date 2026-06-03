#include <string.h>
#include <iostream>

#include "pduSessionResourceSetupRequestPdu.h"

PDUSessionResourceSetupRequestPdu::PDUSessionResourceSetupRequestPdu()
{
    reset_v();
}

PDUSessionResourceSetupRequestPdu::~PDUSessionResourceSetupRequestPdu()
{
}

void PDUSessionResourceSetupRequestPdu::reset_v()
{
    pduSessResSetupRequest_mp = NULL;

    //RAN UE NGAP ID
    ranUeNgapId_m = 0;
    ranUeNgapId_mp = NULL;

    //AMF UE NGAP ID
    amfUeNgapId_m = 0;
    amfUeNgapId_mp = NULL;

    //NAS PDU
    nasPdu_m.length = 0;
    memset(nasPdu_ma.data, 0xff, sizeof(nasPdu_ma.data));
    nasPdu_m.value = nasPdu_ma.data;
    nasPdu_mp = NULL;
    nasPduPrst_m = false;

    //UE Aggregate Maximum Bit Rate
    ueAggregateMaximumBitRate_m.bit_mask = 0;
    ueAggregateMaximumBitRate_m.uEAggregateMaximumBitRateDL = 0;
    ueAggregateMaximumBitRate_m.uEAggregateMaximumBitRateUL = 0;
    ueAggregateMaximumBitRate_m.iE_Extensions=NULL;
    ueAggregateMaximumBitRate_mp = NULL;
    ueAMBRPresent_m=false;

    //RAN Paging Priority
    ranPagePriority_m = 0;
    ranPagePriority_mp = NULL;
    ranPagePriorityPrst_m = false;

    //PDUSessionResourceSetupListSUReq
    pduSessResSetupListSuReq_mp = NULL;
    pduSessResSetupListSuReq_m = NULL;
    numofPduSessResSetupListItem_m = 0;
    memset(pduNasPdus_ma, 0xff, sizeof(pduNasPdus_ma));
    memset(transfers_ma, 0xff, sizeof(transfers_ma));

    for (int i = 0; i < maxnoofPDUSessions_c; i++)
    {
        pduSessResSetupListSuReq_ma[i].next = NULL;
        pduSessResSetupListSuReq_ma[i].value.bit_mask = 0;
        pduSessResSetupListSuReq_ma[i].value.pDUSessionID = 0;
        pduSessResSetupListSuReq_ma[i].value.pDUSessionNAS_PDU.length = 0;
        pduSessResSetupListSuReq_ma[i].value.pDUSessionNAS_PDU.value = pduNasPdus_ma[i].data;
        pduSessResSetupListSuReq_ma[i].value.s_NSSAI.bit_mask = 0;
        pduSessResSetupListSuReq_ma[i].value.s_NSSAI.sST.length = 0;
        memset(pduSessResSetupListSuReq_ma[i].value.s_NSSAI.sST.value, 0xff,
            sizeof(pduSessResSetupListSuReq_ma[i].value.s_NSSAI.sST.value));
        pduSessResSetupListSuReq_ma[i].value.s_NSSAI.sD.length = 0;
        memset(pduSessResSetupListSuReq_ma[i].value.s_NSSAI.sD.value, 0xff, 
            sizeof(pduSessResSetupListSuReq_ma[i].value.s_NSSAI.sD.value));
        pduSessResSetupListSuReq_ma[i].value.pDUSessionResourceSetupRequestTransfer.encoded.length = 0;
        pduSessResSetupListSuReq_ma[i].value.pDUSessionResourceSetupRequestTransfer.encoded.value = transfers_ma[i].data;
        pduSessResSetupListSuReq_ma[i].value.pDUSessionResourceSetupRequestTransfer.decoded = NULL;
    }
}

unsigned int
PDUSessionResourceSetupRequestPdu::encodeOssToPer(NgapOssCtxt *ctxt)
{
    if (NULL == ctxt)
    {
        return -1;
    }

    ProtocolIeType *nextIe_p = NULL;

    //PDUSessionResourceSetupListSUReq
    ProtocolIeType PDUSessionResourceSetupListSUReqIe;
    if (numofPduSessResSetupListItem_m > maxnoofPDUSessions 
        || numofPduSessResSetupListItem_m < 1)
    {
        return defaultFailureCode_c;
    }
    else
    {
        struct PDUSessionResourceSetupListSUReq_ *nextPDUSessionResourceSetupListSUReq_p = NULL;
        for (unsigned int i = 0; i < numofPduSessResSetupListItem_m; i++)
        {
            pduSessResSetupListSuReq_ma[i].next = nextPDUSessionResourceSetupListSUReq_p;
            nextPDUSessionResourceSetupListSUReq_p = &(pduSessResSetupListSuReq_ma[i]);
        }
        pduSessResSetupListSuReq_m = nextPDUSessionResourceSetupListSUReq_p;
        PDUSessionResourceSetupListSUReqIe.value.value.pduNum = 
            PDU_PDUSessionResourceSetupRequestIEs_Value_PDUSessionResourceSetupListSUReq;
        PDUSessionResourceSetupListSUReqIe.value.value.decoded.pdu_PDUSessionResourceSetupListSUReq = &pduSessResSetupListSuReq_m;
        PDUSessionResourceSetupListSUReqIe.value.value.encoded.length = 0;
        PDUSessionResourceSetupListSUReqIe.value.value.encoded.value = NULL;
        PDUSessionResourceSetupListSUReqIe.next = nextIe_p;
        PDUSessionResourceSetupListSUReqIe.value.id = id_PDUSessionResourceSetupListSUReq;
        PDUSessionResourceSetupListSUReqIe.value.criticality = reject;
        nextIe_p = &PDUSessionResourceSetupListSUReqIe;
    }

    //NAS PDU
    ProtocolIeType nasPduIe;
    if (nasPduPrst_m)
    {
        nasPduIe.next = nextIe_p;
        nasPduIe.value.id = id_NAS_PDU;
        nasPduIe.value.criticality = reject;
        nasPduIe.value.value.pduNum = 
            PDU_PDUSessionResourceSetupRequestIEs_Value_NAS_PDU;
        nasPduIe.value.value.decoded.pdu_NAS_PDU = &nasPdu_m;
        nasPduIe.value.value.encoded.length = 0;
        nasPduIe.value.value.encoded.value = NULL;
        nextIe_p = &nasPduIe;
    }

    //UE Aggregate Maximum Bit Rate
    ProtocolIeType ueAggregateMaximumBitRateIe;
    if(ueAMBRPresent_m)
    {
        ueAggregateMaximumBitRateIe.next = nextIe_p;
        ueAggregateMaximumBitRateIe.value.id = id_UEAggregateMaximumBitRate;
        ueAggregateMaximumBitRateIe.value.criticality = ignore;
        ueAggregateMaximumBitRateIe.value.value.pduNum =
            PDU_PDUSessionResourceSetupRequestIEs_Value_UEAggregateMaximumBitRate;
        ueAggregateMaximumBitRateIe.value.value.decoded.pdu_UEAggregateMaximumBitRate =
            &ueAggregateMaximumBitRate_m;
        ueAggregateMaximumBitRateIe.value.value.encoded.length = 0;
        ueAggregateMaximumBitRateIe.value.value.encoded.value = NULL;
        nextIe_p = &ueAggregateMaximumBitRateIe;
    }

    //RAN Paging Priority
    ProtocolIeType ranPagingPriorityIe;
    if (ranPagePriorityPrst_m)
    {
        ranPagingPriorityIe.next = nextIe_p;
        ranPagingPriorityIe.value.id = id_RANPagingPriority;
        ranPagingPriorityIe.value.criticality = ignore;
        ranPagingPriorityIe.value.value.pduNum = 
            PDU_PDUSessionResourceSetupRequestIEs_Value_RANPagingPriority;
        ranPagingPriorityIe.value.value.decoded.pdu_RANPagingPriority = &ranPagePriority_m;
        ranPagingPriorityIe.value.value.encoded.length = 0;
        ranPagingPriorityIe.value.value.encoded.value = NULL;
        nextIe_p = &ranPagingPriorityIe;
    }

    //RAN_UE_NGAP_ID
    ProtocolIeType ranUeNgapIdIe;
    ranUeNgapIdIe.next = nextIe_p;
    ranUeNgapIdIe.value.id = id_RAN_UE_NGAP_ID;
    ranUeNgapIdIe.value.criticality = reject;
    ranUeNgapIdIe.value.value.pduNum = 
        PDU_PDUSessionResourceSetupRequestIEs_Value_RAN_UE_NGAP_ID;
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
        PDU_PDUSessionResourceSetupRequestIEs_Value_AMF_UE_NGAP_ID;
    amfUeNgapIdIe.value.value.decoded.pdu_AMF_UE_NGAP_ID = &amfUeNgapId_m;
    amfUeNgapIdIe.value.value.encoded.length = 0;
    amfUeNgapIdIe.value.value.encoded.value = NULL;
    nextIe_p = &amfUeNgapIdIe;

    //message
    PDUSessionResourceSetupRequest msg;
    msg.protocolIEs = nextIe_p;

    NGAP_PDU ngApPdu;

    ngApPdu.choice = initiatingMessage_chosen;
    ngApPdu.u.initiatingMessage.procedureCode = id_PDUSessionResourceSetup;
    ngApPdu.u.initiatingMessage.criticality = reject;
    ngApPdu.u.initiatingMessage.value.pduNum = 
        PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_PDUSessionResourceSetupRequest;
    ngApPdu.u.initiatingMessage.value.decoded.pdu_PDUSessionResourceSetupRequest = &msg;
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

void PDUSessionResourceSetupRequestPdu::getDecodedOssDataPtrs_v()
{
    ProtocolIeType *currIe_p;
    reset_v();

    ngapPdu_mp = reinterpret_cast<NGAP_PDU *>(ossBufRef_m.value);
    if (ngapPdu_mp == NULL)
    {
        return;
    }

    pduSessResSetupRequest_mp =
        ngapPdu_mp->u.initiatingMessage.value.decoded.pdu_PDUSessionResourceSetupRequest;

    if (pduSessResSetupRequest_mp == NULL)
    {
        return;
    }

    currIe_p = pduSessResSetupRequest_mp->protocolIEs;

    while (currIe_p != NULL)
    {
        if (currIe_p->value.id == id_RAN_UE_NGAP_ID && 
            currIe_p->value.value.decoded.pdu_RAN_UE_NGAP_ID != NULL)
        {
            ranUeNgapId_mp = currIe_p->value.value.decoded.pdu_RAN_UE_NGAP_ID;
        }
        else if (currIe_p->value.id == id_AMF_UE_NGAP_ID && 
            currIe_p->value.value.decoded.pdu_AMF_UE_NGAP_ID != NULL)
        {
            amfUeNgapId_mp = currIe_p->value.value.decoded.pdu_AMF_UE_NGAP_ID;
        }
        else if (currIe_p->value.id == id_NAS_PDU && 
            currIe_p->value.value.decoded.pdu_NAS_PDU != NULL)
        {
            nasPdu_mp = currIe_p->value.value.decoded.pdu_NAS_PDU;
            nasPduPrst_m = true;
        }
        else if (currIe_p->value.id == id_UEAggregateMaximumBitRate
            && currIe_p->value.value.decoded.pdu_UEAggregateMaximumBitRate != NULL) {
            ueAggregateMaximumBitRate_mp =
                currIe_p->value.value.decoded.pdu_UEAggregateMaximumBitRate;
            ueAMBRPresent_m=true;
        }
        else if (currIe_p->value.id == id_RANPagingPriority && 
            currIe_p->value.value.decoded.pdu_RANPagingPriority != NULL)
        {
            ranPagePriority_mp =
                currIe_p->value.value.decoded.pdu_RANPagingPriority;
            ranPagePriorityPrst_m = true;
        }
        else if (currIe_p->value.id == id_PDUSessionResourceSetupListSUReq && 
            currIe_p->value.value.decoded.pdu_PDUSessionResourceSetupListSUReq != NULL)
        {
            pduSessResSetupListSuReq_mp =
                currIe_p->value.value.decoded.pdu_PDUSessionResourceSetupListSUReq;
            PDUSessionResourceSetupListSUReq pduSessSetupReq_p = *pduSessResSetupListSuReq_mp;
            while (pduSessSetupReq_p != NULL)
            {
                numofPduSessResSetupListItem_m++;
                pduSessSetupReq_p = pduSessSetupReq_p->next;
            }
        }
        currIe_p = currIe_p->next;
    }
}

PDUSessionResourceSetupRequest *
PDUSessionResourceSetupRequestPdu::getPduSessResSetupReqDecodedPtr() const
{
    return pduSessResSetupRequest_mp;
}

RAN_UE_NGAP_ID *
PDUSessionResourceSetupRequestPdu::getRanUeNgapIdDecodedPtr() const
{
    return ranUeNgapId_mp;
}

RAN_UE_NGAP_ID *
PDUSessionResourceSetupRequestPdu::getRanUeNgapIdEncodePtr()
{
    return &ranUeNgapId_m;
}

AMF_UE_NGAP_ID *
PDUSessionResourceSetupRequestPdu::getAmfUeNgapIdDecodedPtr() const
{
    return amfUeNgapId_mp;
}

AMF_UE_NGAP_ID *
PDUSessionResourceSetupRequestPdu::getAmfUeNgapIdEncodePtr()
{
    return &amfUeNgapId_m;
}

NAS_PDU *
PDUSessionResourceSetupRequestPdu::getNasPduDecodedPtr() const
{
    return nasPdu_mp;
}

NAS_PDU *
PDUSessionResourceSetupRequestPdu::getNasPduEncodePtr()
{
    nasPduPrst_m = true;
    return &nasPdu_m;
}

bool PDUSessionResourceSetupRequestPdu::isNasPduPresent()
{
    return nasPduPrst_m;
}

RANPagingPriority *
PDUSessionResourceSetupRequestPdu::getRanPagePriorityDecodedPtr() const
{
    return ranPagePriority_mp;
}

RANPagingPriority *
PDUSessionResourceSetupRequestPdu::getRanPagePriorityEncodePtr()
{
    ranPagePriorityPrst_m = true;
    return &ranPagePriority_m;
}

bool 
PDUSessionResourceSetupRequestPdu::isRanPagePriorityPrst()
{
    return ranPagePriorityPrst_m;
}

PDUSessionResourceSetupListSUReq *
PDUSessionResourceSetupRequestPdu::getPduSessResSetupListSUReqDecodedPtr() const
{
    return pduSessResSetupListSuReq_mp;
}

PDUSessionResourceSetupListSUReq *
PDUSessionResourceSetupRequestPdu::getPduSessResSetupListSUReqEncodePtr()
{
    return &pduSessResSetupListSuReq_m;
}

PDUSessionResourceSetupItemSUReq *
PDUSessionResourceSetupRequestPdu::getPduSessResSetupSUReqItemDecodedPtr(unsigned int index) const
{
    PDUSessionResourceSetupItemSUReq *item_p = NULL;

    PDUSessionResourceSetupListSUReq pduSessResSetupList_p = NULL;

    if (pduSessResSetupListSuReq_mp != NULL)
    {
        pduSessResSetupList_p = *(pduSessResSetupListSuReq_mp);
    }

    if (pduSessResSetupList_p != NULL && (index >= 0 && index < numofPduSessResSetupListItem_m))
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

PDUSessionResourceSetupItemSUReq *
PDUSessionResourceSetupRequestPdu::getPduSessResSetupSUReqItemEncodePtr(unsigned int index)
{
    if (index >= 0 && index < maxnoofPDUSessions)
    {
        return &(pduSessResSetupListSuReq_ma[index].value);
    }
    else
    {
        return NULL;
    }
}
void PDUSessionResourceSetupRequestPdu::setNumofPduSessResSetupListSUReqItem_v(unsigned int num)
{
    numofPduSessResSetupListItem_m = num;
}

unsigned int
PDUSessionResourceSetupRequestPdu::getNumofPduSessResSetupListSUReqItem() const
{
    return numofPduSessResSetupListItem_m;
}

UEAggregateMaximumBitRate*
PDUSessionResourceSetupRequestPdu::getUeAggregateMaximumBitRateDecodedPtr() const
{
    return ueAggregateMaximumBitRate_mp;
}

UEAggregateMaximumBitRate*
PDUSessionResourceSetupRequestPdu::getUeAggregateMaximumBitRateEncodePtr()
{
    ueAMBRPresent_m=true;
    return &ueAggregateMaximumBitRate_m;
}

bool
PDUSessionResourceSetupRequestPdu::isUeAMBRPresent()
{
    return ueAMBRPresent_m;
}

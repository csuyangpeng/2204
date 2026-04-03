#include <iostream>
#include <string.h>

#include "pduSessionResourceModifyRequestPdu.h"

PduSessResModReqPdu::PduSessResModReqPdu()
{
    reset_v();
}

PduSessResModReqPdu::~PduSessResModReqPdu()
{
}

void PduSessResModReqPdu::reset_v()
{
    pduSessResModifyListModReq_mp = NULL;

    //RAN UE NGAP ID
    ranUeNgapId_m = 0;
    ranUeNgapId_mp = NULL;

    //AMF UE NGAP ID
    amfUeNgapId_m = 0;
    amfUeNgapId_mp = NULL;

    //RAN Paging Priority
    ranPagingPriority_m = 0;
    ranPagingPriority_mp = NULL;
    ranPagingPriorityPresent_m = false;

    //PDUSessionResourceModifyListModReq
    pduSessResModifyListModReq_mp = NULL;
    pduSessResModifyListModReq_m = NULL;
    numofPduSessResModifyItemModReq_m = 0;

    memset(pduNasPdus_ma, 0xff, sizeof(pduNasPdus_ma));
    memset(transfers_ma, 0xff, sizeof(transfers_ma));

    for (unsigned int i = 0; i < maxnoofPDUSessions_c; i++) {
        pduSessResModifyListModReq_ma[i].next = NULL;
        pduSessResModifyListModReq_ma[i].value.bit_mask = 0;
        pduSessResModifyListModReq_ma[i].value.iE_Extensions = NULL;
        pduSessResModifyListModReq_ma[i].value.pDUSessionID = 0;
                pduSessResModifyListModReq_ma[i].value.nAS_PDU.length = 0;
        pduSessResModifyListModReq_ma[i].value.nAS_PDU.value = pduNasPdus_ma[i].data;
        pduSessResModifyListModReq_ma[i].value.pDUSessionResourceModifyRequestTransfer.encoded.length = 0;
        pduSessResModifyListModReq_ma[i].value.pDUSessionResourceModifyRequestTransfer.encoded.value = transfers_ma[i].data;
        pduSessResModifyListModReq_ma[i].value.pDUSessionResourceModifyRequestTransfer.decoded = NULL;
    }
}

unsigned int
PduSessResModReqPdu::encodeOssToPer(NgapOssCtxt* ctxt)
{
    if (NULL == ctxt) {
        return -1;
    }

    ProtocolIeType* nextIe_p = NULL;

    //PDUSessionResourceModifyListModReq
    ProtocolIeType pduSessResModifyListModReqIe;
    if(numofPduSessResModifyItemModReq_m > maxnoofPDUSessions ||
          numofPduSessResModifyItemModReq_m < 0)
    {
        return defaultFailureCode_c;
    }
    else
    {
        struct PDUSessionResourceModifyListModReq_ *nextPduSessResModifyListModReq_p=NULL;
        for(unsigned int i=0;i<numofPduSessResModifyItemModReq_m;i++)
        {
            pduSessResModifyListModReq_ma[i].next=nextPduSessResModifyListModReq_p;
            nextPduSessResModifyListModReq_p=&(pduSessResModifyListModReq_ma[i]);

                pduSessResModifyListModReq_m=nextPduSessResModifyListModReq_p;

            pduSessResModifyListModReqIe.value.value.pduNum=
                    PDU_PDUSessionResourceModifyRequestIEs_Value_PDUSessionResourceModifyListModReq;
                pduSessResModifyListModReqIe.value.value.decoded.pdu_PDUSessionResourceModifyListModReq
                                =&pduSessResModifyListModReq_m;
                pduSessResModifyListModReqIe.value.value.encoded.length = 0;
                pduSessResModifyListModReqIe.value.value.encoded.value = NULL;
            pduSessResModifyListModReqIe.next=nextIe_p;
                pduSessResModifyListModReqIe.value.id=id_PDUSessionResourceModifyListModReq;
            pduSessResModifyListModReqIe.value.criticality=reject;
                nextIe_p=&pduSessResModifyListModReqIe;
        }
    }

    //RAN Paging Priority
    ProtocolIeType ranPagingPriorityIe;
    if (ranPagingPriorityPresent_m)
    {
        ranPagingPriorityIe.next = nextIe_p;
        ranPagingPriorityIe.value.id = id_RANPagingPriority;
        ranPagingPriorityIe.value.criticality = ignore;
        ranPagingPriorityIe.value.value.pduNum =
            PDU_PDUSessionResourceModifyRequestIEs_Value_RANPagingPriority;
        ranPagingPriorityIe.value.value.decoded.pdu_RANPagingPriority = &ranPagingPriority_m;
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
        PDU_PDUSessionResourceModifyRequestIEs_Value_RAN_UE_NGAP_ID;
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
        PDU_PDUSessionResourceModifyRequestIEs_Value_AMF_UE_NGAP_ID;
    amfUeNgapIdIe.value.value.decoded.pdu_AMF_UE_NGAP_ID = &amfUeNgapId_m;
    amfUeNgapIdIe.value.value.encoded.length = 0;
    amfUeNgapIdIe.value.value.encoded.value = NULL;
    nextIe_p = &amfUeNgapIdIe;

    //message
    PDUSessionResourceModifyRequest msg;
    msg.protocolIEs = nextIe_p;

    NGAP_PDU ngApPdu;

    ngApPdu.choice = initiatingMessage_chosen;
    ngApPdu.u.initiatingMessage.procedureCode = id_PDUSessionResourceModify;
    ngApPdu.u.initiatingMessage.criticality = reject;
    ngApPdu.u.initiatingMessage.value.pduNum =
		PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_PDUSessionResourceModifyRequest;
    ngApPdu.u.initiatingMessage.value.decoded.pdu_PDUSessionResourceModifyRequest = &msg;
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

void PduSessResModReqPdu::getDecodedOssDataPtrs_v()
{
    ProtocolIeType* currIe_p;
    reset_v();

    ngapPdu_mp = reinterpret_cast<NGAP_PDU*>(ossBufRef_m.value);
    if (ngapPdu_mp == NULL) {
        return;
    }

    pduSessResMdfyReq_mp =
        ngapPdu_mp->u.initiatingMessage.value.decoded.pdu_PDUSessionResourceModifyRequest;

    if (pduSessResMdfyReq_mp == NULL) {
        return;
    }

    currIe_p = pduSessResMdfyReq_mp->protocolIEs;

    while (currIe_p != NULL) {
        if (currIe_p->value.id == id_RAN_UE_NGAP_ID
            && currIe_p->value.value.decoded.pdu_RAN_UE_NGAP_ID != NULL) {
            ranUeNgapId_mp = currIe_p->value.value.decoded.pdu_RAN_UE_NGAP_ID;
        } else if (currIe_p->value.id == id_AMF_UE_NGAP_ID
            && currIe_p->value.value.decoded.pdu_AMF_UE_NGAP_ID != NULL) {
            amfUeNgapId_mp = currIe_p->value.value.decoded.pdu_AMF_UE_NGAP_ID;
        } else if (currIe_p->value.id == id_RANPagingPriority &&
            currIe_p->value.value.decoded.pdu_RANPagingPriority != NULL)
        {
            ranPagingPriority_mp =
                currIe_p->value.value.decoded.pdu_RANPagingPriority;
            ranPagingPriorityPresent_m = true;
        } else if (currIe_p->value.id == id_PDUSessionResourceModifyListModReq
            && currIe_p->value.value.decoded.pdu_PDUSessionResourceModifyListModReq != NULL)
        {
            pduSessResModifyListModReq_mp =
                      currIe_p->value.value.decoded.pdu_PDUSessionResourceModifyListModReq;

            PDUSessionResourceModifyListModReq pduSessResModReq_p=NULL;

            pduSessResModReq_p=*pduSessResModifyListModReq_mp;
            while(pduSessResModReq_p!=NULL)
            {
                numofPduSessResModifyItemModReq_m++;
                pduSessResModReq_p=pduSessResModReq_p->next;
            }
        }
        currIe_p = currIe_p->next;
    }
}

PDUSessionResourceModifyRequest*
PduSessResModReqPdu::getPduSessResModReqDecodedPtr() const
{
    return pduSessResMdfyReq_mp;
}

RAN_UE_NGAP_ID*
PduSessResModReqPdu::getRanUeNgapIdDecodedPtr() const
{
    return ranUeNgapId_mp;
}

RAN_UE_NGAP_ID*
PduSessResModReqPdu::getRanUeNgapIdEncodePtr()
{
    return &ranUeNgapId_m;
}

AMF_UE_NGAP_ID*
PduSessResModReqPdu::getAmfUeNgapIdDecodedPtr() const
{
    return amfUeNgapId_mp;
}

AMF_UE_NGAP_ID*
PduSessResModReqPdu::getAmfUeNgapIdEncodePtr()
{
    return &amfUeNgapId_m;
}

RANPagingPriority *
PduSessResModReqPdu::getRANPagingPriorityDecodedPtr() const
{
    return ranPagingPriority_mp;
}

RANPagingPriority *
PduSessResModReqPdu::getRANPagingPriorityEncodePtr()
{
    ranPagingPriorityPresent_m = true;
    return &ranPagingPriority_m;
}

bool PduSessResModReqPdu::isRANPagingPriorityPresent()
{
    return ranPagingPriorityPresent_m;
}

PDUSessionResourceModifyListModReq*
PduSessResModReqPdu::getPduSessResModifyListModReqDecodedPtr() const
{
    return pduSessResModifyListModReq_mp;
}

PDUSessionResourceModifyListModReq*
PduSessResModReqPdu::getPduSessResModifyListModReqEncodePtr()
{
    return &pduSessResModifyListModReq_m;
}

PDUSessionResourceModifyItemModReq*
PduSessResModReqPdu::getPduSessResModifyItemModReqDecodedPtr(unsigned int index) const
{
    PDUSessionResourceModifyItemModReq* item_p = NULL;

    PDUSessionResourceModifyListModReq pduSessResModList_p = NULL;

    if (pduSessResModifyListModReq_mp != NULL) {
        pduSessResModList_p = *(pduSessResModifyListModReq_mp);
    }

    if (pduSessResModList_p != NULL
        && (index >= 0 && index < numofPduSessResModifyItemModReq_m)) {
        unsigned int i = 0;
        for (; i < index && pduSessResModList_p != NULL; i++) {
            pduSessResModList_p = pduSessResModList_p->next;
        }
        if (pduSessResModList_p != NULL && i == index) {
            item_p = &(pduSessResModList_p->value);
        }
    }

    return item_p;
}

PDUSessionResourceModifyItemModReq*
PduSessResModReqPdu::getPduSessResModifyItemModReqEncodePtr(unsigned int index)
{
    if (index >= 0 && index < maxnoofPDUSessions) {
        return &(pduSessResModifyListModReq_ma[index].value);
    } else {
        return NULL;
    }
}

void PduSessResModReqPdu::setNumofPduSessResModifyItemModReq_v(unsigned int num)
{
    numofPduSessResModifyItemModReq_m = num;
}

unsigned int
PduSessResModReqPdu::getNumofPduSessResModifyItemModReq() const
{
    return numofPduSessResModifyItemModReq_m;
}

S_NSSAI* 
PduSessResModReqPdu::getSnssaiEncodePtr()
{
    return &snssai_m;
}

SnssaiIeValue* 
PduSessResModReqPdu::getIeExtenIeTypeDecodedPtr()
{
    return &snssaiIe;
}

unsigned char*
PduSessResModReqPdu::getSnssaiStr()
{
    return snssaiStr_ma[0].data;
}


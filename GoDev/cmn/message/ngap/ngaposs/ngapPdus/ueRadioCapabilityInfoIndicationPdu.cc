#include <iostream>
#include <string.h>

#include "ueRadioCapabilityInfoIndicationPdu.h"

UeRadioCapaInfoIndicationPdu::UeRadioCapaInfoIndicationPdu()
{
    reset_v();
}

UeRadioCapaInfoIndicationPdu::~UeRadioCapaInfoIndicationPdu()
{
}

void UeRadioCapaInfoIndicationPdu::reset_v()
{
    ueRadioCapaInfoIndication_mp = NULL;

    //RAN UE NGAP ID
    ranUeNgapId_m = 0;
    ranUeNgapId_mp = NULL;

    //AMF UE NGAP ID
    amfUeNgapId_m = 0;
    amfUeNgapId_mp = NULL;

    //UERadioCapability
    ueRadioCapability_mp = NULL;
    ueRadioCapability_m.length = 0;
    memset(ueRadioCapability_ma, 0xff, sizeof(ueRadioCapability_ma));
    ueRadioCapability_m.value = ueRadioCapability_ma;


    //UERadioCapabilityForPaging
    ueRadioCapabilityForPaging_mp = NULL;
    ueRadioCapabilityForPaging_m.uERadioCapabilityForPagingOfNR.length = 0;
    memset(ueRadioCapabilityForPagingNr_ma, 0xff, sizeof(ueRadioCapabilityForPagingNr_ma));
    ueRadioCapabilityForPaging_m.uERadioCapabilityForPagingOfNR.value = ueRadioCapabilityForPagingNr_ma;
	ueRadioCapabilityForPaging_m.uERadioCapabilityForPagingOfEUTRA.length = 0;
	memset(ueRadioCapabilityForPagingEutra_ma, 0xff, sizeof(ueRadioCapabilityForPagingEutra_ma));
	ueRadioCapabilityForPaging_m.uERadioCapabilityForPagingOfEUTRA.value = ueRadioCapabilityForPagingEutra_ma;
    ueRadioCapabilityForPagingPresent_m=false;

}

unsigned int
UeRadioCapaInfoIndicationPdu::encodeOssToPer(NgapOssCtxt* ctxt)
{
    if (NULL == ctxt) {
        return -1;
    }

    ProtocolIeType* nextIe_p = NULL;

    //UERadioCapabilityForPaging
    ProtocolIeType ueRadioCapabilityForPagingIe;
    if(ueRadioCapabilityForPagingPresent_m)
    {
        ueRadioCapabilityForPagingIe.next = nextIe_p;
        ueRadioCapabilityForPagingIe.value.id = id_UERadioCapabilityForPaging;
        ueRadioCapabilityForPagingIe.value.criticality = ignore;
        ueRadioCapabilityForPagingIe.value.value.pduNum = 
            PDU_UERadioCapabilityInfoIndicationIEs_Value_UERadioCapabilityForPaging;
        ueRadioCapabilityForPagingIe.value.value.decoded.pdu_UERadioCapabilityForPaging = 
            &ueRadioCapabilityForPaging_m;
        ueRadioCapabilityForPagingIe.value.value.encoded.length = 0;
        ueRadioCapabilityForPagingIe.value.value.encoded.value = NULL;
        nextIe_p = &ueRadioCapabilityForPagingIe;
    }

    //UERadioCapability
    ProtocolIeType ueRadioCapabilityIe;
    ueRadioCapabilityIe.next = nextIe_p;
    ueRadioCapabilityIe.value.id = id_UERadioCapability;
    ueRadioCapabilityIe.value.criticality = ignore;
    ueRadioCapabilityIe.value.value.pduNum =
        PDU_UERadioCapabilityInfoIndicationIEs_Value_UERadioCapability;
    ueRadioCapabilityIe.value.value.decoded.pdu_UERadioCapability = &ueRadioCapability_m;
    ueRadioCapabilityIe.value.value.encoded.length = 0;
    ueRadioCapabilityIe.value.value.encoded.value = NULL;
    nextIe_p = &ueRadioCapabilityIe;

    //RAN_UE_NGAP_ID
    ProtocolIeType ranUeNgapIdIe;
    ranUeNgapIdIe.next = nextIe_p;
    ranUeNgapIdIe.value.id = id_RAN_UE_NGAP_ID;
    ranUeNgapIdIe.value.criticality = reject;
    ranUeNgapIdIe.value.value.pduNum =
        PDU_UERadioCapabilityInfoIndicationIEs_Value_RAN_UE_NGAP_ID;
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
        PDU_UERadioCapabilityInfoIndicationIEs_Value_AMF_UE_NGAP_ID;
    amfUeNgapIdIe.value.value.decoded.pdu_AMF_UE_NGAP_ID = &amfUeNgapId_m;
    amfUeNgapIdIe.value.value.encoded.length = 0;
    amfUeNgapIdIe.value.value.encoded.value = NULL;
    nextIe_p = &amfUeNgapIdIe;

    //message
    UERadioCapabilityInfoIndication msg;
    msg.protocolIEs = nextIe_p;

    NGAP_PDU ngApPdu;

    ngApPdu.choice = initiatingMessage_chosen;
    ngApPdu.u.initiatingMessage.procedureCode = id_UERadioCapabilityInfoIndication;
    ngApPdu.u.initiatingMessage.criticality = ignore;
    ngApPdu.u.initiatingMessage.value.pduNum = 
        PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_UERadioCapabilityInfoIndication;
    ngApPdu.u.initiatingMessage.value.decoded.pdu_UERadioCapabilityInfoIndication = &msg;
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

void UeRadioCapaInfoIndicationPdu::getDecodedOssDataPtrs_v()
{
    ProtocolIeType* currIe_p;
    reset_v();

    ngapPdu_mp = reinterpret_cast<NGAP_PDU*>(ossBufRef_m.value);
    if (ngapPdu_mp == NULL) {
        return;
    }

    ueRadioCapaInfoIndication_mp = 
        ngapPdu_mp->u.initiatingMessage.value.decoded.pdu_UERadioCapabilityInfoIndication;

    if (ueRadioCapaInfoIndication_mp == NULL) {
        return;
    }

    currIe_p = ueRadioCapaInfoIndication_mp->protocolIEs;

    while (currIe_p != NULL) {
        if (currIe_p->value.id == id_RAN_UE_NGAP_ID
            && currIe_p->value.value.decoded.pdu_RAN_UE_NGAP_ID != NULL) {
            ranUeNgapId_mp = currIe_p->value.value.decoded.pdu_RAN_UE_NGAP_ID;
        } else if (currIe_p->value.id == id_AMF_UE_NGAP_ID
            && currIe_p->value.value.decoded.pdu_AMF_UE_NGAP_ID != NULL) {
            amfUeNgapId_mp = currIe_p->value.value.decoded.pdu_AMF_UE_NGAP_ID;
        } else if (currIe_p->value.id == id_UERadioCapability
            && currIe_p->value.value.decoded.pdu_UERadioCapability != NULL) {
            ueRadioCapability_mp = currIe_p->value.value.decoded.pdu_UERadioCapability;
        } else if (currIe_p->value.id == id_UERadioCapabilityForPaging
            && currIe_p->value.value.decoded.pdu_UERadioCapabilityForPaging != NULL) {
            ueRadioCapabilityForPaging_mp = currIe_p->value.value.decoded.pdu_UERadioCapabilityForPaging;
            ueRadioCapabilityForPagingPresent_m=true;
        }
        currIe_p = currIe_p->next;
    }
}

UERadioCapabilityInfoIndication*
UeRadioCapaInfoIndicationPdu::getUeRadioCapaInfoIndicationDecodedPtr() const
{
    return ueRadioCapaInfoIndication_mp;
}

RAN_UE_NGAP_ID*
UeRadioCapaInfoIndicationPdu::getRanUeNgapIdDecodedPtr() const
{
    return ranUeNgapId_mp;
}

RAN_UE_NGAP_ID*
UeRadioCapaInfoIndicationPdu::getRanUeNgapIdEncodePtr()
{
    return &ranUeNgapId_m;
}

AMF_UE_NGAP_ID*
UeRadioCapaInfoIndicationPdu::getAmfUeNgapIdDecodedPtr() const
{
    return amfUeNgapId_mp;
}

AMF_UE_NGAP_ID*
UeRadioCapaInfoIndicationPdu::getAmfUeNgapIdEncodePtr()
{
    return &amfUeNgapId_m;
}

UERadioCapability*
UeRadioCapaInfoIndicationPdu::getUeRadioCapabilityDecodedPtr() const
{
    return ueRadioCapability_mp;
}

UERadioCapability*
UeRadioCapaInfoIndicationPdu::getUeRadioCapabilityEncodePtr()
{
    return &ueRadioCapability_m;
}

UERadioCapabilityForPaging*
UeRadioCapaInfoIndicationPdu::getUeRadioCapabilityForPagingDecodedPtr() const
{
    return ueRadioCapabilityForPaging_mp;
}

UERadioCapabilityForPaging*
UeRadioCapaInfoIndicationPdu::getUeRadioCapabilityForPagingEncodePtr()
{
    ueRadioCapabilityForPagingPresent_m=true;
    return &ueRadioCapabilityForPaging_m;
}

bool
UeRadioCapaInfoIndicationPdu::isUeRadioCapabilityForPagingPresent()
{
    return ueRadioCapabilityForPagingPresent_m;
}

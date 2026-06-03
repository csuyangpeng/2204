#include <iostream>
#include <string.h>

#include "pduSessionResourceReleaseCommandPdu.h"

PduSessResRelCmdPdu::PduSessResRelCmdPdu()
{
    reset_v();
}

PduSessResRelCmdPdu::~PduSessResRelCmdPdu()
{
}

void PduSessResRelCmdPdu::reset_v()
{
    PduSessResRelCmd_mp = NULL;

    //RAN UE NGAP ID
    ranUeNgapId_m = 0;
    ranUeNgapId_mp = NULL;

    //AMF UE NGAP ID
    amfUeNgapId_m = 0;
    amfUeNgapId_mp = NULL;

    //NAS PDU
    nasPdu_m.length = 0;
    memset(nasPdu_ma, 0xff, sizeof(nasPdu_ma));
    nasPdu_m.value = nasPdu_ma;
    nasPdu_mp = NULL;
    nasPduPresent_m=false;

    //RAN Paging Priority
    ranPagingPriority_m = 0;
    ranPagingPriority_mp = NULL;
    ranPagingPriorityPresent_m = false;

    //PDUSessionResourceSetupListCxtReq
    pduSessResToRelListRelCmd_mp = NULL;
    pduSessResToRelListRelCmd_m = NULL;
    numofPduSessResToRelListItemRelCmd_m = 0;

    memset(transfers_ma, 0xff, sizeof(transfers_ma));

    for (unsigned int i = 0; i < maxnoofPDUSessions_c; i++) {
        pduSessResToRelListRelCmd_ma[i].next = NULL;
        pduSessResToRelListRelCmd_ma[i].value.bit_mask = 0;
        pduSessResToRelListRelCmd_ma[i].value.iE_Extensions=NULL;
        pduSessResToRelListRelCmd_ma[i].value.pDUSessionID = 0;
        pduSessResToRelListRelCmd_ma[i].value.pDUSessionResourceReleaseCommandTransfer
            .encoded.length = 0;
        pduSessResToRelListRelCmd_ma[i].value.pDUSessionResourceReleaseCommandTransfer
            .encoded.value = transfers_ma[i].data;
        pduSessResToRelListRelCmd_ma[i].value.pDUSessionResourceReleaseCommandTransfer
            .decoded = NULL;
    }
}

unsigned int
PduSessResRelCmdPdu::encodeOssToPer(NgapOssCtxt* ctxt)
{
    if (NULL == ctxt) {
        return -1;
    }

    ProtocolIeType* nextIe_p = NULL;

    //PDUSessionResourceToReleaseListRelCmd
    ProtocolIeType pduSessResToRelListRelCmdIe;
    if(numofPduSessResToRelListItemRelCmd_m > maxnoofPDUSessions ||
          numofPduSessResToRelListItemRelCmd_m < 0)
    {
        return defaultFailureCode_c;
    }
    else
    {
        struct PDUSessionResourceToReleaseListRelCmd_ *nextPduSessResToRelListRelCmd_p=NULL;
        for(unsigned int i=0;i<numofPduSessResToRelListItemRelCmd_m;i++)
        {
            pduSessResToRelListRelCmd_ma[i].next=nextPduSessResToRelListRelCmd_p;
            nextPduSessResToRelListRelCmd_p=&(pduSessResToRelListRelCmd_ma[i]);

	        pduSessResToRelListRelCmd_m=nextPduSessResToRelListRelCmd_p;
	
    	    pduSessResToRelListRelCmdIe.value.value.pduNum=
        	    PDU_PDUSessionResourceReleaseCommandIEs_Value_PDUSessionResourceToReleaseListRelCmd;
	        pduSessResToRelListRelCmdIe.value.value.decoded.pdu_PDUSessionResourceToReleaseListRelCmd
				=&pduSessResToRelListRelCmd_m;
        	pduSessResToRelListRelCmdIe.value.value.encoded.length = 0;
	        pduSessResToRelListRelCmdIe.value.value.encoded.value = NULL;
    	    pduSessResToRelListRelCmdIe.next=nextIe_p;
        	pduSessResToRelListRelCmdIe.value.id=id_PDUSessionResourceToReleaseListRelCmd;
    	    pduSessResToRelListRelCmdIe.value.criticality=reject;
	        nextIe_p=&pduSessResToRelListRelCmdIe;
        }
    }

    //NAS PDU
    ProtocolIeType nasPduIe;
    if(nasPduPresent_m)
    {
        nasPduIe.next = nextIe_p;
        nasPduIe.value.id = id_NAS_PDU;
        nasPduIe.value.criticality = ignore;
        nasPduIe.value.value.pduNum =
            PDU_PDUSessionResourceReleaseCommandIEs_Value_NAS_PDU;
        nasPduIe.value.value.decoded.pdu_NAS_PDU = &nasPdu_m;
        nasPduIe.value.value.encoded.length = 0;
        nasPduIe.value.value.encoded.value = NULL;
        nextIe_p = &nasPduIe;
    }

    //RAN Paging Priority
    ProtocolIeType ranPagingPriorityIe;
    if (ranPagingPriorityPresent_m)
    {
        ranPagingPriorityIe.next = nextIe_p;
        ranPagingPriorityIe.value.id = id_RANPagingPriority;
        ranPagingPriorityIe.value.criticality = ignore;
        ranPagingPriorityIe.value.value.pduNum =
            PDU_PDUSessionResourceReleaseCommandIEs_Value_RANPagingPriority;
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
        PDU_PDUSessionResourceReleaseCommandIEs_Value_RAN_UE_NGAP_ID;
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
        PDU_PDUSessionResourceReleaseCommandIEs_Value_AMF_UE_NGAP_ID;
    amfUeNgapIdIe.value.value.decoded.pdu_AMF_UE_NGAP_ID = &amfUeNgapId_m;
    amfUeNgapIdIe.value.value.encoded.length = 0;
    amfUeNgapIdIe.value.value.encoded.value = NULL;
    nextIe_p = &amfUeNgapIdIe;

    //message
    PDUSessionResourceReleaseCommand msg;
    msg.protocolIEs = nextIe_p;

    NGAP_PDU ngApPdu;

    ngApPdu.choice = initiatingMessage_chosen;
    ngApPdu.u.initiatingMessage.procedureCode = id_PDUSessionResourceRelease;
    ngApPdu.u.initiatingMessage.criticality = reject;
    ngApPdu.u.initiatingMessage.value.pduNum = 
        PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_PDUSessionResourceReleaseCommand;
    ngApPdu.u.initiatingMessage.value.decoded.pdu_PDUSessionResourceReleaseCommand = &msg;
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

void PduSessResRelCmdPdu::getDecodedOssDataPtrs_v()
{
    ProtocolIeType* currIe_p;
    reset_v();

    ngapPdu_mp = reinterpret_cast<NGAP_PDU*>(ossBufRef_m.value);
    if (ngapPdu_mp == NULL) {
        return;
    }

    PduSessResRelCmd_mp = 
        ngapPdu_mp->u.initiatingMessage.value.decoded.pdu_PDUSessionResourceReleaseCommand;

    if (PduSessResRelCmd_mp == NULL) {
        return;
    }

    currIe_p = PduSessResRelCmd_mp->protocolIEs;

    while (currIe_p != NULL) {
        if (currIe_p->value.id == id_RAN_UE_NGAP_ID
            && currIe_p->value.value.decoded.pdu_RAN_UE_NGAP_ID != NULL) {
            ranUeNgapId_mp = currIe_p->value.value.decoded.pdu_RAN_UE_NGAP_ID;
        } else if (currIe_p->value.id == id_AMF_UE_NGAP_ID
            && currIe_p->value.value.decoded.pdu_AMF_UE_NGAP_ID != NULL) {
            amfUeNgapId_mp = currIe_p->value.value.decoded.pdu_AMF_UE_NGAP_ID;
        } else if (currIe_p->value.id == id_NAS_PDU
            && currIe_p->value.value.decoded.pdu_NAS_PDU != NULL) {
            nasPdu_mp = currIe_p->value.value.decoded.pdu_NAS_PDU;
            nasPduPresent_m=true;
        } else if (currIe_p->value.id == id_RANPagingPriority && 
            currIe_p->value.value.decoded.pdu_RANPagingPriority != NULL)
        {
            ranPagingPriority_mp =
                currIe_p->value.value.decoded.pdu_RANPagingPriority;
            ranPagingPriorityPresent_m = true;
        } else if (currIe_p->value.id == id_PDUSessionResourceSetupListCxtReq
            && currIe_p->value.value.decoded.pdu_PDUSessionResourceToReleaseListRelCmd != NULL)
        {
            pduSessResToRelListRelCmd_mp = 
                      currIe_p->value.value.decoded.pdu_PDUSessionResourceToReleaseListRelCmd;

            PDUSessionResourceToReleaseListRelCmd pduSessResRelCmd_p=NULL;

            pduSessResRelCmd_p=*pduSessResToRelListRelCmd_mp;
            while(pduSessResRelCmd_p!=NULL)
            {
                pduSessResRelCmd_p=pduSessResRelCmd_p->next;
            }
        }
        currIe_p = currIe_p->next;
    }
}

PDUSessionResourceReleaseCommand*
PduSessResRelCmdPdu::getPduSessResRelCmdDecodedPtr() const
{
    return PduSessResRelCmd_mp;
}

RAN_UE_NGAP_ID*
PduSessResRelCmdPdu::getRanUeNgapIdDecodedPtr() const
{
    return ranUeNgapId_mp;
}

RAN_UE_NGAP_ID*
PduSessResRelCmdPdu::getRanUeNgapIdEncodePtr()
{
    return &ranUeNgapId_m;
}

AMF_UE_NGAP_ID*
PduSessResRelCmdPdu::getAmfUeNgapIdDecodedPtr() const
{
    return amfUeNgapId_mp;
}

AMF_UE_NGAP_ID*
PduSessResRelCmdPdu::getAmfUeNgapIdEncodePtr()
{
    return &amfUeNgapId_m;
}

NAS_PDU*
PduSessResRelCmdPdu::getNasPduDecodedPtr() const
{
    return nasPdu_mp;
}

NAS_PDU*
PduSessResRelCmdPdu::getNasPduEncodePtr()
{
    nasPduPresent_m=true;
    return &nasPdu_m;
}

bool 
PduSessResRelCmdPdu::isNasPduPresent()
{
    return nasPduPresent_m;
}

RANPagingPriority *
PduSessResRelCmdPdu::getRANPagingPriorityDecodedPtr() const
{
    return ranPagingPriority_mp;
}

RANPagingPriority *
PduSessResRelCmdPdu::getRANPagingPriorityEncodePtr()
{
    ranPagingPriorityPresent_m = true;
    return &ranPagingPriority_m;
}

bool PduSessResRelCmdPdu::isRANPagingPriorityPresent()
{
    return ranPagingPriorityPresent_m;
}

PDUSessionResourceToReleaseListRelCmd*
PduSessResRelCmdPdu::getPduSessResToRelListRelCmdDecodedPtr() const
{
    return pduSessResToRelListRelCmd_mp;
}

PDUSessionResourceToReleaseListRelCmd*
PduSessResRelCmdPdu::getPduSessResToRelListRelCmdEncodePtr()
{
    return &pduSessResToRelListRelCmd_m;
}

PDUSessionResourceToReleaseItemRelCmd*
PduSessResRelCmdPdu::getPduSessResToRelItemRelCmdDecodedPtr(unsigned int index) const
{
    PDUSessionResourceToReleaseItemRelCmd* item_p = NULL;

    PDUSessionResourceToReleaseListRelCmd pduSessResRelList_p = NULL;

    if (pduSessResToRelListRelCmd_mp != NULL) {
        pduSessResRelList_p = *(pduSessResToRelListRelCmd_mp);
    }

    if (pduSessResRelList_p != NULL
        && (index >= 0 && index < numofPduSessResToRelListItemRelCmd_m)) {
        unsigned int i = 0;
        for (; i < index && pduSessResRelList_p != NULL; i++) {
            pduSessResRelList_p = pduSessResRelList_p->next;
        }
        if (pduSessResRelList_p != NULL && i == index) {
            item_p = &(pduSessResRelList_p->value);
        }
    }

    return item_p;
}

PDUSessionResourceToReleaseItemRelCmd*
PduSessResRelCmdPdu::getPduSessResToRelItemRelCmdEncodePtr(unsigned int index)
{
    if (index >= 0 && index < maxnoofPDUSessions) {
        return &(pduSessResToRelListRelCmd_ma[index].value);
    } else {
        return NULL;
    }
}
void PduSessResRelCmdPdu::setNumofPduSessResToRelItemRelCmd_v(unsigned int num)
{
    numofPduSessResToRelListItemRelCmd_m = num;
}

unsigned int
PduSessResRelCmdPdu::getNumofPduSessResToRelItemRelCmd() const
{
    return numofPduSessResToRelListItemRelCmd_m;
}
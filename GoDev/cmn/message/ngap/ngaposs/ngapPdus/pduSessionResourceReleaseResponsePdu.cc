#include <iostream>
#include <string.h>

#include "pduSessionResourceReleaseResponsePdu.h"

PduSessResRelRespPdu::PduSessResRelRespPdu()
{
    reset_v();
}

PduSessResRelRespPdu::~PduSessResRelRespPdu()
{
}

void PduSessResRelRespPdu::reset_v()
{
    pduSessResRelResp_mp = NULL;

    //RAN UE NGAP ID
    ranUeNgapId_m = 0;
    ranUeNgapId_mp = NULL;

    //AMF UE NGAP ID
    amfUeNgapId_m = 0;
    amfUeNgapId_mp = NULL;

    //NgRanCgi
    ngranCgi_m.choice = nR_CGI_chosen;
    ngranCgi_m.u.nR_CGI.bit_mask=0;
    ngranCgi_m.u.nR_CGI.iE_Extensions = NULL;
    ngranCgi_m.u.nR_CGI.pLMNIdentity.length=0;
    memset(ngranCgi_m.u.nR_CGI.pLMNIdentity.value, 0xff,
           sizeof(ngranCgi_m.u.nR_CGI.pLMNIdentity.value));
    ngranCgi_m.u.nR_CGI.nRCellIdentity.length=0;
    memset(nrCellId_mb,0xff,sizeof(nrCellId_mb));
    ngranCgi_m.u.nR_CGI.nRCellIdentity.value = nrCellId_mb;
    ngranCgi_m.u.nR_CGI.iE_Extensions=NULL;

    //Ie Extension
    userLocalInfoIe.next = NULL;
    userLocalInfoIe.value.id = id_PSCellInformation;
    userLocalInfoIe.value.criticality = ignore;
    userLocalInfoIe.value.extensionValue.pduNum = 
        PDU_UserLocationInformationEUTRA_ExtIEs_Extension_NGRAN_CGI;
    userLocalInfoIe.value.extensionValue.decoded.pdu_NGRAN_CGI = &ngranCgi_m;
    userLocalInfoIe.value.extensionValue.encoded.length = 0;
    userLocalInfoIe.value.extensionValue.encoded.value = NULL;

    //User Location Information
    //only NR supported currently
    useLocationInformation_m.choice=userLocationInformationNR_chosen;
    useLocationInformation_m.u.userLocationInformationNR.bit_mask=0;
    useLocationInformation_m.u.userLocationInformationNR.iE_Extensions = NULL;
    useLocationInformation_m.u.userLocationInformationNR.nR_CGI.bit_mask=0;
    useLocationInformation_m.u.userLocationInformationNR.nR_CGI.pLMNIdentity.length=0;
    memset(useLocationInformation_m.u.userLocationInformationNR.nR_CGI.pLMNIdentity.value, 0xff,
           sizeof(useLocationInformation_m.u.userLocationInformationNR.nR_CGI.pLMNIdentity.value));
    useLocationInformation_m.u.userLocationInformationNR.nR_CGI.nRCellIdentity.length=0;
    memset(nrCellId_ma,0xff,sizeof(nrCellId_ma));
    useLocationInformation_m.u.userLocationInformationNR.nR_CGI.nRCellIdentity.value = nrCellId_ma;
    useLocationInformation_m.u.userLocationInformationNR.nR_CGI.iE_Extensions=NULL;
    
    useLocationInformation_m.u.userLocationInformationNR.tAI.bit_mask=0;
    useLocationInformation_m.u.userLocationInformationNR.tAI.pLMNIdentity.length=0;
    memset(useLocationInformation_m.u.userLocationInformationNR.tAI.pLMNIdentity.value, 0xff,
           sizeof(useLocationInformation_m.u.userLocationInformationNR.tAI.pLMNIdentity.value));
    useLocationInformation_m.u.userLocationInformationNR.tAI.tAC.length=0;
    memset(useLocationInformation_m.u.userLocationInformationNR.tAI.tAC.value, 0xff,
           sizeof(useLocationInformation_m.u.userLocationInformationNR.tAI.tAC.value));
    useLocationInformation_m.u.userLocationInformationNR.tAI.iE_Extensions=NULL;
    
    useLocationInformation_m.u.userLocationInformationNR.timeStamp.length=0;
    memset(useLocationInformation_m.u.userLocationInformationNR.timeStamp.value, 0xff,
           sizeof(useLocationInformation_m.u.userLocationInformationNR.timeStamp.value));
    
    useLocationInformation_m.u.userLocationInformationNR.iE_Extensions = &userLocalInfoIe;
    userLocInfoPresent_m = false;
           
    useLocationInformation_mp=NULL;

    //PDUSessionResourceReleasedListRelRes
    pduSessResListRelRes_mp = NULL;
    pduSessResListRelRes_m = NULL;
    numofPduSessResListItemCxtReq_m = 0;

    memset(transfers_ma, 0xff, sizeof(transfers_ma));

    for (unsigned int i = 0; i < maxnoofPDUSessions_c; i++) {
        pduSessResListRelResp_ma[i].next = NULL;
        pduSessResListRelResp_ma[i].value.bit_mask = 0;
        pduSessResListRelResp_ma[i].value.iE_Extensions=NULL;
        pduSessResListRelResp_ma[i].value.pDUSessionID = 0;
        pduSessResListRelResp_ma[i].value.pDUSessionResourceReleaseResponseTransfer
            .encoded.length = 0;
        pduSessResListRelResp_ma[i].value.pDUSessionResourceReleaseResponseTransfer
            .encoded.value = transfers_ma[i].data;
        pduSessResListRelResp_ma[i].value.pDUSessionResourceReleaseResponseTransfer
            .decoded = NULL;
    }
}

unsigned int
PduSessResRelRespPdu::encodeOssToPer(NgapOssCtxt* ctxt)
{
    if (NULL == ctxt) {
        return -1;
    }

    ProtocolIeType* nextIe_p = NULL;

    //User Location Information
    ProtocolIeType useLocationInformationIe;
    if(userLocInfoPresent_m){
        useLocationInformationIe.next =nextIe_p;
        useLocationInformationIe.value.id = id_UserLocationInformation;
        useLocationInformationIe.value.criticality = ignore;
        useLocationInformationIe.value.value.pduNum =
            PDU_PDUSessionResourceReleaseResponseIEs_Value_UserLocationInformation;
        useLocationInformationIe.value.value.decoded.pdu_UserLocationInformation = &useLocationInformation_m;
        useLocationInformationIe.value.value.encoded.length = 0;
        useLocationInformationIe.value.value.encoded.value = NULL;
        nextIe_p = &useLocationInformationIe;
    }

    //PDUSessionResourceReleasedListRelRes
    ProtocolIeType pduSessResListRelRespIe;
    if(numofPduSessResListItemCxtReq_m > maxnoofPDUSessions ||
          numofPduSessResListItemCxtReq_m < 0)
    {
        return defaultFailureCode_c;
    }
    else
    {
        struct PDUSessionResourceReleasedListRelRes_ *nextPduSessResListRelResp_p=NULL;
        for(unsigned int i=0;i<numofPduSessResListItemCxtReq_m;i++)
        {
            pduSessResListRelResp_ma[i].next=nextPduSessResListRelResp_p;
            nextPduSessResListRelResp_p=&(pduSessResListRelResp_ma[i]);

            pduSessResListRelRes_m=nextPduSessResListRelResp_p;
    
            pduSessResListRelRespIe.value.value.pduNum=
                PDU_PDUSessionResourceReleaseResponseIEs_Value_PDUSessionResourceReleasedListRelRes;
            pduSessResListRelRespIe.value.value.decoded.pdu_PDUSessionResourceReleasedListRelRes
                =&pduSessResListRelRes_m;
            pduSessResListRelRespIe.value.value.encoded.length = 0;
            pduSessResListRelRespIe.value.value.encoded.value = NULL;
            pduSessResListRelRespIe.next=nextIe_p;
            pduSessResListRelRespIe.value.id=id_PDUSessionResourceReleasedListRelRes;
            pduSessResListRelRespIe.value.criticality= ignore;
            nextIe_p=&pduSessResListRelRespIe;
        }
    }

    //RAN_UE_NGAP_ID
    ProtocolIeType ranUeNgapIdIe;
    ranUeNgapIdIe.next = nextIe_p;
    ranUeNgapIdIe.value.id = id_RAN_UE_NGAP_ID;
    ranUeNgapIdIe.value.criticality = ignore;
    ranUeNgapIdIe.value.value.pduNum =
        PDU_PDUSessionResourceReleaseResponseIEs_Value_RAN_UE_NGAP_ID;
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
        PDU_PDUSessionResourceReleaseResponseIEs_Value_AMF_UE_NGAP_ID;
    amfUeNgapIdIe.value.value.decoded.pdu_AMF_UE_NGAP_ID = &amfUeNgapId_m;
    amfUeNgapIdIe.value.value.encoded.length = 0;
    amfUeNgapIdIe.value.value.encoded.value = NULL;
    nextIe_p = &amfUeNgapIdIe;

    //message
    PDUSessionResourceReleaseResponse msg;
    msg.protocolIEs = nextIe_p;

    NGAP_PDU ngApPdu;

    ngApPdu.choice = successfulOutcome_chosen;
    ngApPdu.u.successfulOutcome.procedureCode = id_PDUSessionResourceRelease;
    ngApPdu.u.successfulOutcome.criticality = reject;
    ngApPdu.u.successfulOutcome.value.pduNum = 
        PDU_NGAP_ELEMENTARY_PROCEDURES_SuccessfulOutcome_PDUSessionResourceReleaseResponse;
    ngApPdu.u.successfulOutcome.value.decoded.pdu_PDUSessionResourceReleaseResponse = &msg;
    ngApPdu.u.successfulOutcome.value.encoded.length = 0;
    ngApPdu.u.successfulOutcome.value.encoded.value = NULL;

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

void PduSessResRelRespPdu::getDecodedOssDataPtrs_v()
{
    ProtocolIeType* currIe_p;
    reset_v();

    ngapPdu_mp = reinterpret_cast<NGAP_PDU*>(ossBufRef_m.value);
    if (ngapPdu_mp == NULL) {
        return;
    }

    pduSessResRelResp_mp = 
        ngapPdu_mp->u.successfulOutcome.value.decoded.pdu_PDUSessionResourceReleaseResponse;

    if (pduSessResRelResp_mp == NULL) {
        return;
    }

    currIe_p = pduSessResRelResp_mp->protocolIEs;

    while (currIe_p != NULL) {
        if (currIe_p->value.id == id_RAN_UE_NGAP_ID
            && currIe_p->value.value.decoded.pdu_RAN_UE_NGAP_ID != NULL) {
            ranUeNgapId_mp = currIe_p->value.value.decoded.pdu_RAN_UE_NGAP_ID;
        } else if (currIe_p->value.id == id_AMF_UE_NGAP_ID
            && currIe_p->value.value.decoded.pdu_AMF_UE_NGAP_ID != NULL) {
            amfUeNgapId_mp = currIe_p->value.value.decoded.pdu_AMF_UE_NGAP_ID;
        } else if (currIe_p->value.id == id_UserLocationInformation
                 && currIe_p->value.value.decoded.pdu_UserLocationInformation != NULL)
        {
            useLocationInformation_mp = currIe_p->value.value.decoded.pdu_UserLocationInformation;
            userLocInfoPresent_m = true;
        } else if (currIe_p->value.id == id_PDUSessionResourceReleasedListRelRes
            && currIe_p->value.value.decoded.pdu_PDUSessionResourceReleasedListRelRes != NULL)
        {
            pduSessResListRelRes_mp = 
                currIe_p->value.value.decoded.pdu_PDUSessionResourceReleasedListRelRes;

            PDUSessionResourceReleasedListRelRes pduSessResRelease_p=NULL;

            pduSessResRelease_p=*pduSessResListRelRes_mp;
            while(pduSessResRelease_p!=NULL)
            {
                numofPduSessResListItemCxtReq_m++;
                pduSessResRelease_p=pduSessResRelease_p->next;
            }
        }
        currIe_p = currIe_p->next;
    }
}

PDUSessionResourceReleaseResponse*
PduSessResRelRespPdu::getPduSessResRelRespDecodedPtr() const
{
    return pduSessResRelResp_mp;
}

RAN_UE_NGAP_ID*
PduSessResRelRespPdu::getRanUeNgapIdDecodedPtr() const
{
    return ranUeNgapId_mp;
}

RAN_UE_NGAP_ID*
PduSessResRelRespPdu::getRanUeNgapIdEncodePtr()
{
    return &ranUeNgapId_m;
}

AMF_UE_NGAP_ID*
PduSessResRelRespPdu::getAmfUeNgapIdDecodedPtr() const
{
    return amfUeNgapId_mp;
}

AMF_UE_NGAP_ID*
PduSessResRelRespPdu::getAmfUeNgapIdEncodePtr()
{
    return &amfUeNgapId_m;
}

UserLocationInformation * 
PduSessResRelRespPdu::getUseLocationInformationDecodedPtr() const
{
    return useLocationInformation_mp;
}

UserLocationInformation * 
PduSessResRelRespPdu::getUseLocationInformationEncodePtr()
{
    userLocInfoPresent_m = true;
    return &useLocationInformation_m;
}

bool
PduSessResRelRespPdu::isUseLocationInformationPresent()
{
    return userLocInfoPresent_m;
}

NGRAN_CGI* 
PduSessResRelRespPdu::getNgRanCgiEncodePtr()
{
    return &ngranCgi_m;
}

UserLocationInfomationNRIeExtension* 
PduSessResRelRespPdu::getIeExtenIeTypeDecodedPtr()
{
    return &userLocalInfoIe;
}

PDUSessionResourceReleasedListRelRes*
PduSessResRelRespPdu::getPduSessResListRelRespDecodedPtr() const
{
    return pduSessResListRelRes_mp;
}

PDUSessionResourceReleasedListRelRes*
PduSessResRelRespPdu::getPduSessResListRelRespEncodePtr()
{
    return &pduSessResListRelRes_m;
}

PDUSessionResourceReleasedItemRelRes*
PduSessResRelRespPdu::getPduSessResItemRelRespDecodedPtr(unsigned int index) const
{
    PDUSessionResourceReleasedItemRelRes* item_p = NULL;

    PDUSessionResourceReleasedListRelRes pduSessResReleaseList_p = NULL;

    if (pduSessResListRelRes_mp != NULL) {
        pduSessResReleaseList_p = *(pduSessResListRelRes_mp);
    }

    if (pduSessResReleaseList_p != NULL
        && (index >= 0 && index < numofPduSessResListItemCxtReq_m)) {
        unsigned int i = 0;
        for (; i < index && pduSessResReleaseList_p != NULL; i++) {
            pduSessResReleaseList_p = pduSessResReleaseList_p->next;
        }
        if (pduSessResReleaseList_p != NULL && i == index) {
            item_p = &(pduSessResReleaseList_p->value);
        }
    }

    return item_p;
}

PDUSessionResourceReleasedItemRelRes*
PduSessResRelRespPdu::getPduSessResItemRelRespEncodePtr(unsigned int index)
{
    if (index >= 0 && index < maxnoofPDUSessions) {
        return &(pduSessResListRelResp_ma[index].value);
    } else {
        return NULL;
    }
}
void PduSessResRelRespPdu::setNumofPduSessResListItemRelResp_v(unsigned int num)
{
    numofPduSessResListItemCxtReq_m = num;
}

unsigned int
PduSessResRelRespPdu::getNumofPduSessResListItemRelResp() const
{
    return numofPduSessResListItemCxtReq_m;
}
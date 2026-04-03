#include <string.h>
#include <iostream>

#include "pduSessionResourceModifyResponsePdu.h"

PduSessResMdfyRespPdu::PduSessResMdfyRespPdu()
{
    reset_v();
}

PduSessResMdfyRespPdu::~PduSessResMdfyRespPdu()
{
}

void
PduSessResMdfyRespPdu::reset_v()
{
    PduSessResMdfyResp_mp = NULL;

    //RAN UE NGAP ID
    ranUeNgapId_m=0;
    ranUeNgapId_mp=NULL;

    //AMF UE NGAP ID
    amfUeNgapId_m=0;
    amfUeNgapId_mp=NULL;

    //PDUSessionResourceModifyListModResp
    pduSessResModifyListModResp_mp=NULL;
    pduSessResModifyListModResp_m=NULL;
    numofPduSessResModifyItemModResp_m=0;
    pduSessResModifyRespListPresent_m = false;

    memset(respTransfers_ma, 0xff, sizeof(respTransfers_ma));
    for(unsigned int i = 0 ; i < maxnoofPDUSessions_c ; i++){
        pduSessResModifyListModResp_ma[i].value.bit_mask = 0;
        pduSessResModifyListModResp_ma[i].value.iE_Extensions = NULL;
        pduSessResModifyListModResp_ma[i].value.pDUSessionID = 0;
        pduSessResModifyListModResp_ma[i].value.pDUSessionResourceModifyResponseTransfer.decoded = NULL;
        pduSessResModifyListModResp_ma[i].value.pDUSessionResourceModifyResponseTransfer.encoded.length = 0;
        pduSessResModifyListModResp_ma[i].value.pDUSessionResourceModifyResponseTransfer.encoded.value = respTransfers_ma[i].data;
    }

    //PDUSessionResourceFailedToModifyListModRes
    pduSessResFailedModifyList_mp = NULL;
    pduSessResFailedModifyList_m = NULL;
    numofpduSessResFailedModifyList_m = 0;
    pduSessResFailModifyRespListPresent_m = false;

    memset(failedTransfers_ma, 0xff, sizeof(failedTransfers_ma));
    for(int i=0;i<maxnoofPDUSessions_c;i++){
        pduSessResFailedModifyList_ma[i].value.bit_mask = 0;
        pduSessResFailedModifyList_ma[i].value.iE_Extensions = NULL;
        pduSessResFailedModifyList_ma[i].value.pDUSessionID = 0;
        pduSessResFailedModifyList_ma[i].value.pDUSessionResourceModifyUnsuccessfulTransfer.decoded = NULL;
        pduSessResFailedModifyList_ma[i].value.pDUSessionResourceModifyUnsuccessfulTransfer.encoded.length = 0;
        pduSessResFailedModifyList_ma[i].value.pDUSessionResourceModifyUnsuccessfulTransfer.encoded.value = failedTransfers_ma[i].data;
    }

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
}

unsigned int
PduSessResMdfyRespPdu::encodeOssToPer(NgapOssCtxt* ctxt)                   
{
    if(NULL == ctxt)
    {         
        return -1;
    }
    
    ProtocolIeType *nextIe_p = NULL;

    //User Location Information
    ProtocolIeType useLocationInformationIe;
    if(userLocInfoPresent_m){
        useLocationInformationIe.next =nextIe_p;
        useLocationInformationIe.value.id = id_UserLocationInformation;
        useLocationInformationIe.value.criticality = ignore;
        useLocationInformationIe.value.value.pduNum =
            PDU_PDUSessionResourceModifyResponseIEs_Value_UserLocationInformation;
        useLocationInformationIe.value.value.decoded.pdu_UserLocationInformation = &useLocationInformation_m;
        useLocationInformationIe.value.value.encoded.length = 0;
        useLocationInformationIe.value.value.encoded.value = NULL;
        nextIe_p = &useLocationInformationIe;
    }

    //PDUSessionResourceFailedToModifyListModRes
    ProtocolIeType pduSessResFailedModifyListRespIe;
    if (pduSessResFailModifyRespListPresent_m)
    {
        if(numofpduSessResFailedModifyList_m > maxnoofPDUSessions)
        {
            return defaultFailureCode_c;
        }
        else if (numofpduSessResFailedModifyList_m >= 1)
        {
            struct PDUSessionResourceFailedToModifyListModRes_ * nextPtr = NULL;
            for (unsigned int i = 0; i < numofpduSessResFailedModifyList_m;i++){
                pduSessResFailedModifyList_ma[i].next = nextPtr;
                nextPtr = &pduSessResFailedModifyList_ma[i];
            }
            pduSessResFailedModifyList_m = nextPtr;
        
            pduSessResFailedModifyListRespIe.value.value.pduNum =
                PDU_PDUSessionResourceModifyResponseIEs_Value_PDUSessionResourceFailedToModifyListModRes;
            pduSessResFailedModifyListRespIe.value.value.decoded.pdu_PDUSessionResourceFailedToModifyListModRes = 
                &pduSessResFailedModifyList_m;
            pduSessResFailedModifyListRespIe.value.value.encoded.length = 0;
            pduSessResFailedModifyListRespIe.value.value.encoded.value = NULL;
        
            pduSessResFailedModifyListRespIe.next=nextIe_p;
            pduSessResFailedModifyListRespIe.value.id=id_PDUSessionResourceFailedToModifyListModRes;
            pduSessResFailedModifyListRespIe.value.criticality=ignore;
            nextIe_p=&pduSessResFailedModifyListRespIe;
        }
    }


    //PDUSessionResourceModifyListModRes
    ProtocolIeType pduSessResModListRespIe;
    if (pduSessResModifyRespListPresent_m) 
    {
        if(numofPduSessResModifyItemModResp_m > maxnoofPDUSessions)
        {
            return defaultFailureCode_c;
        }
        else if(numofPduSessResModifyItemModResp_m >= 1 )
        {
            struct PDUSessionResourceModifyListModRes_ * nextPtr = NULL;
            for (unsigned int i = 0; i< numofPduSessResModifyItemModResp_m; i++){
                pduSessResModifyListModResp_ma[i].next = nextPtr;
                nextPtr = &pduSessResModifyListModResp_ma[i];
            }
            pduSessResModifyListModResp_m = nextPtr;

            pduSessResModListRespIe.value.value.pduNum =
                PDU_PDUSessionResourceModifyResponseIEs_Value_PDUSessionResourceModifyListModRes;
            pduSessResModListRespIe.value.value.decoded.pdu_PDUSessionResourceModifyListModRes =
                &pduSessResModifyListModResp_m;
            pduSessResModListRespIe.value.value.encoded.length = 0;
            pduSessResModListRespIe.value.value.encoded.value = NULL;

            pduSessResModListRespIe.next=nextIe_p;
            pduSessResModListRespIe.value.id=id_PDUSessionResourceModifyListModRes;
            pduSessResModListRespIe.value.criticality=ignore;
            nextIe_p=&pduSessResModListRespIe;
        }
    }

    //RAN_UE_NGAP_ID
    ProtocolIeType ranUeNgapIdIe;
    ranUeNgapIdIe.next =nextIe_p;
    ranUeNgapIdIe.value.id = id_RAN_UE_NGAP_ID;
    ranUeNgapIdIe.value.criticality = ignore;
    ranUeNgapIdIe.value.value.pduNum =PDU_PDUSessionResourceModifyResponseIEs_Value_RAN_UE_NGAP_ID;
    ranUeNgapIdIe.value.value.decoded.pdu_RAN_UE_NGAP_ID = &ranUeNgapId_m;
    ranUeNgapIdIe.value.value.encoded.length = 0;
    ranUeNgapIdIe.value.value.encoded.value = NULL;
    nextIe_p = &ranUeNgapIdIe;

    //AMF_UE_NGAP_ID
    ProtocolIeType amfUeNgapIdIe;
    amfUeNgapIdIe.next =nextIe_p;
    amfUeNgapIdIe.value.id = id_AMF_UE_NGAP_ID;
    amfUeNgapIdIe.value.criticality = ignore;
    amfUeNgapIdIe.value.value.pduNum =PDU_PDUSessionResourceModifyResponseIEs_Value_AMF_UE_NGAP_ID;
    amfUeNgapIdIe.value.value.decoded.pdu_AMF_UE_NGAP_ID = &amfUeNgapId_m;
    amfUeNgapIdIe.value.value.encoded.length = 0;
    amfUeNgapIdIe.value.value.encoded.value = NULL;
    nextIe_p = &amfUeNgapIdIe;

    //message
    PDUSessionResourceModifyResponse msg;
    msg.protocolIEs = nextIe_p;
    
    NGAP_PDU ngApPdu;

    ngApPdu.choice = successfulOutcome_chosen;
    ngApPdu.u.successfulOutcome.procedureCode = id_PDUSessionResourceModify;
    ngApPdu.u.successfulOutcome.criticality = reject;
    ngApPdu.u.successfulOutcome.value.pduNum = 
        PDU_NGAP_ELEMENTARY_PROCEDURES_SuccessfulOutcome_PDUSessionResourceModifyResponse;
    ngApPdu.u.successfulOutcome.value.decoded.pdu_PDUSessionResourceModifyResponse = &msg;
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
PduSessResMdfyRespPdu::getDecodedOssDataPtrs_v()
{
    ProtocolIeType *currIe_p;
    reset_v();

    ngapPdu_mp = reinterpret_cast<NGAP_PDU *> (ossBufRef_m.value);
    if (ngapPdu_mp == NULL)
    {
        return;
    }
    
    PduSessResMdfyResp_mp =
            ngapPdu_mp->u.successfulOutcome.value.decoded.pdu_PDUSessionResourceModifyResponse;
    
    if (PduSessResMdfyResp_mp == NULL)
    {
        return;
    }

    currIe_p = PduSessResMdfyResp_mp->protocolIEs;

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
        else if (currIe_p->value.id == id_PDUSessionResourceModifyListModRes
            && currIe_p->value.value.decoded.pdu_PDUSessionResourceModifyListModRes != NULL)
        {
            pduSessResModifyListModResp_mp =
                currIe_p->value.value.decoded.pdu_PDUSessionResourceModifyListModRes;

            PDUSessionResourceModifyListModRes nextPtr = NULL;
            nextPtr = *pduSessResModifyListModResp_mp;
            while(nextPtr != NULL)
            {
                numofPduSessResModifyItemModResp_m++;
                nextPtr = nextPtr->next;
            }
            pduSessResModifyRespListPresent_m = true;
        }
        else if (currIe_p->value.id == id_PDUSessionResourceFailedToModifyListModRes
            && currIe_p->value.value.decoded.pdu_PDUSessionResourceFailedToModifyListModRes != NULL)
        {
            pduSessResFailedModifyList_mp=
                currIe_p->value.value.decoded.pdu_PDUSessionResourceFailedToModifyListModRes;

            PDUSessionResourceFailedToModifyListModRes nextPtr = NULL;
            nextPtr = *pduSessResFailedModifyList_mp;
            while(nextPtr != NULL)
            {
                numofpduSessResFailedModifyList_m++;
                nextPtr=nextPtr->next;
            }
            pduSessResFailModifyRespListPresent_m = true;
        }
        else if (currIe_p->value.id == id_UserLocationInformation
            && currIe_p->value.value.decoded.pdu_UserLocationInformation != NULL)
        {
            useLocationInformation_mp = currIe_p->value.value.decoded.pdu_UserLocationInformation;
            userLocInfoPresent_m = true;
        }
        currIe_p = currIe_p->next;
    }
}

PDUSessionResourceModifyResponse *
PduSessResMdfyRespPdu::getPduSessResMdfyRespDecodedPtr() const
{
    return PduSessResMdfyResp_mp;
}

RAN_UE_NGAP_ID * 
PduSessResMdfyRespPdu::getRanUeNgapIdDecodedPtr() const
{
    return ranUeNgapId_mp;
}

RAN_UE_NGAP_ID * 
PduSessResMdfyRespPdu::getRanUeNgapIdEncodePtr()
{
    return &ranUeNgapId_m;
}

AMF_UE_NGAP_ID * 
PduSessResMdfyRespPdu::getAmfUeNgapIdDecodedPtr() const
{
    return amfUeNgapId_mp;
}

AMF_UE_NGAP_ID * 
PduSessResMdfyRespPdu::getAmfUeNgapIdEncodePtr()
{
    return &amfUeNgapId_m;
}

PDUSessionResourceModifyListModRes * 
PduSessResMdfyRespPdu::getPduSessResModifyListModResDecodedPtr() const
{
    return pduSessResModifyListModResp_mp;
}

PDUSessionResourceModifyListModRes * 
PduSessResMdfyRespPdu::getPduSessResModifyListModResEncodePtr()
{
    return &pduSessResModifyListModResp_m;
}

PDUSessionResourceModifyItemModRes*
PduSessResMdfyRespPdu::getPduSessResModifyItemModResDecodedPtr(unsigned int index) const
{
    PDUSessionResourceModifyItemModRes *item_p = NULL;

    PDUSessionResourceModifyListModRes pduSessResModifyList_p = NULL;

    if (pduSessResModifyListModResp_mp != NULL)
    {
        pduSessResModifyList_p = *(pduSessResModifyListModResp_mp);
    }

    if (pduSessResModifyList_p != NULL && (index >= 0 && index < numofPduSessResModifyItemModResp_m))
    {
        unsigned int i = 0;
        for (; i < index && pduSessResModifyList_p != NULL; i++)
        {
            pduSessResModifyList_p = pduSessResModifyList_p->next;
        }
        if (pduSessResModifyList_p != NULL && i == index)
        {
            item_p = &(pduSessResModifyList_p->value);
        }
    }

    return item_p;
}

PDUSessionResourceModifyItemModRes*
PduSessResMdfyRespPdu::getPduSessResModifyItemModResEncodePtr(unsigned int index)
{
    if (index >= 0 && index < maxnoofPDUSessions)
    {
        pduSessResModifyRespListPresent_m = true;
        return &(pduSessResModifyListModResp_ma[index].value);
    }
    else
    {
        return NULL;
    }
}

bool 
PduSessResMdfyRespPdu::isPduSessionResourceModifyListPresent()
{
    return pduSessResModifyRespListPresent_m;
}


void 
PduSessResMdfyRespPdu::setNumofPduSessResModifyItemModRes_v(unsigned int num)
{
    numofPduSessResModifyItemModResp_m = num;
}

unsigned int 
PduSessResMdfyRespPdu::getNumofPduSessResModifyItemModRes() const
{
   return numofPduSessResModifyItemModResp_m;
}

PDUSessionResourceFailedToModifyListModRes * 
PduSessResMdfyRespPdu::getPduSessResFailedModifyListModResDecodedPtr() const
{
    return pduSessResFailedModifyList_mp;
}

PDUSessionResourceFailedToModifyListModRes * 
PduSessResMdfyRespPdu::getPduSessResFailedModifyListModResEncodePtr()
{
    return &pduSessResFailedModifyList_m;
}

PDUSessionResourceFailedToModifyItemModRes*
PduSessResMdfyRespPdu::getPduSessResFailedModifyItemModResDecodedPtr(unsigned int index) const
{
    PDUSessionResourceFailedToModifyItemModRes *item_p = NULL;

    PDUSessionResourceFailedToModifyListModRes pduSessResFailedToModifyList_p = NULL;

    if (pduSessResFailedModifyList_mp != NULL)
    {
        pduSessResFailedToModifyList_p = *(pduSessResFailedModifyList_mp);
    }

    if (pduSessResFailedToModifyList_p != NULL && 
        (index >= 0 && index < numofpduSessResFailedModifyList_m))
    {
        unsigned int i = 0;
        for (; i < index && pduSessResFailedToModifyList_p != NULL; i++)
        {
            pduSessResFailedToModifyList_p = pduSessResFailedToModifyList_p->next;
        }
        if (pduSessResFailedToModifyList_p != NULL && i == index)
        {
            item_p = &(pduSessResFailedToModifyList_p->value);
        }
    }

    return item_p;
}

PDUSessionResourceFailedToModifyItemModRes*
PduSessResMdfyRespPdu::getPduSessResFailedModifyItemModResEncodePtr(unsigned int index)
{
    if (index >= 0 && index < maxnoofPDUSessions)
    {
        pduSessResFailModifyRespListPresent_m = true;
        return &(pduSessResFailedModifyList_ma[index].value);
    }
    else
    {
        return NULL;
    }
}

bool
PduSessResMdfyRespPdu::isPduSessionResourceFailModifyListPresent()
{
    return pduSessResFailModifyRespListPresent_m;
}

void 
PduSessResMdfyRespPdu::setNumofFailedToModifyItem_v(unsigned int num)
{
    numofpduSessResFailedModifyList_m = num;
}

unsigned int 
PduSessResMdfyRespPdu::getNumofFailedToModifyItem() const
{
   return numofpduSessResFailedModifyList_m;
}

UserLocationInformation * 
PduSessResMdfyRespPdu::getUseLocationInformationDecodedPtr() const
{
    return useLocationInformation_mp;
}

UserLocationInformation * 
PduSessResMdfyRespPdu::getUseLocationInformationEncodePtr()
{
    userLocInfoPresent_m = true;
    return &useLocationInformation_m;
}

bool
PduSessResMdfyRespPdu::isUseLocationInformationPresent()
{
    return userLocInfoPresent_m;
}

NGRAN_CGI* 
PduSessResMdfyRespPdu::getNgRanCgiEncodePtr()
{
    return &ngranCgi_m;
}

UserLocationInfomationNRIeExtension* 
PduSessResMdfyRespPdu::getIeExtenIeTypeDecodedPtr()
{
    return &userLocalInfoIe;
}


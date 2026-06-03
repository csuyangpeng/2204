#include <iostream>
#include <string.h>

#include "ueContextReleaseCompletePdu.h"

UeContextReleaseCompletePdu::UeContextReleaseCompletePdu()
{
    reset_v();
}

UeContextReleaseCompletePdu::~UeContextReleaseCompletePdu()
{
}

void UeContextReleaseCompletePdu::reset_v()
{
    UeContextReleaseComplete_mp = NULL;

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
    
    //User Local Information Ie Extension
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
           
    useLocationInformation_mp=NULL;
    isUliPrst_m = false;

    //InfoOnRecommendedCellsAndRANNodesForPaging
    infoOnRecommendedCellsAndRANNodesForPaging_m.bit_mask=0;
    infoOnRecommendedCellsAndRANNodesForPaging_m.iE_Extensions=NULL;
    infoOnRecommendedCellsAndRANNodesForPaging_m.recommendedCellsForPaging.bit_mask=0;
    infoOnRecommendedCellsAndRANNodesForPaging_m.recommendedCellsForPaging.iE_Extensions=NULL;
    infoOnRecommendedCellsAndRANNodesForPaging_m.recommendedCellsForPaging.recommendedCellList=NULL;
    numOfRecommendedCellItem_m=0;
    for (unsigned int i = 0; i < maxnoofRecommendedCells_c; i++) 
    {
        recommendedCellList_ma[i].next = NULL;
        recommendedCellList_ma[i].value.bit_mask = 0;
        recommendedCellList_ma[i].value.iE_Extensions=NULL;
        recommendedCellList_ma[i].value.nGRAN_CGI.choice = nR_CGI_chosen;
        recommendedCellList_ma[i].value.nGRAN_CGI.u.nR_CGI.bit_mask = 0;
        recommendedCellList_ma[i].value.nGRAN_CGI.u.nR_CGI.iE_Extensions=NULL;
        recommendedCellList_ma[i].value.nGRAN_CGI.u.nR_CGI.pLMNIdentity
            .length = 0;
        memset(recommendedCellList_ma[i].value.nGRAN_CGI.u.nR_CGI.pLMNIdentity.value,
            0xff, sizeof(recommendedCellList_ma[i].value.nGRAN_CGI.u.nR_CGI
                .pLMNIdentity.value));
        recommendedCellList_ma[i].value.nGRAN_CGI.u.nR_CGI.nRCellIdentity
            .length = 0;
        recommendedCellList_ma[i].value.nGRAN_CGI.u.nR_CGI.nRCellIdentity
            .value = &(nrCellmaId[i]);
        recommendedCellList_ma[i].value.timeStayedInCell = 0;
    }
    infoOnRecommendedCellsAndRANNodesForPaging_m.recommendRANNodesForPaging.bit_mask=0;
    infoOnRecommendedCellsAndRANNodesForPaging_m.recommendRANNodesForPaging.iE_Extensions=NULL;
    infoOnRecommendedCellsAndRANNodesForPaging_m.recommendRANNodesForPaging.recommendedRANNodeList=NULL;
    numOfRecommendedRANNodeItem_m=0;
    for (unsigned int i = 0; i < maxnoofRecommendedRANNodes_c; i++) 
    {
        recommendedRANNodeList_ma[i].next = NULL;
        recommendedRANNodeList_ma[i].value.bit_mask = 0;
        recommendedRANNodeList_ma[i].value.iE_Extensions=NULL;
        recommendedRANNodeList_ma[i].value.aMFPagingTarget.choice = globalRANNodeID_chosen;
        recommendedRANNodeList_ma[i].value.aMFPagingTarget.u.globalRANNodeID.choice = globalGNB_ID_chosen;
        recommendedRANNodeList_ma[i].value.aMFPagingTarget.u.globalRANNodeID.u.globalGNB_ID.bit_mask=0;
        recommendedRANNodeList_ma[i].value.aMFPagingTarget.u.globalRANNodeID.u.globalGNB_ID.iE_Extensions=NULL;
        recommendedRANNodeList_ma[i].value.aMFPagingTarget.u.globalRANNodeID.u.globalGNB_ID.pLMNIdentity.length=0;
        memset(recommendedRANNodeList_ma[i].value.aMFPagingTarget.u.globalRANNodeID.u.globalGNB_ID.pLMNIdentity.value,
               0xff,
               sizeof(recommendedRANNodeList_ma[i].value.aMFPagingTarget.u.globalRANNodeID.u.globalGNB_ID.pLMNIdentity.value));
        recommendedRANNodeList_ma[i].value.aMFPagingTarget.u.globalRANNodeID.u.globalGNB_ID.gNB_ID.choice=gNB_ID_chosen;
        recommendedRANNodeList_ma[i].value.aMFPagingTarget.u.globalRANNodeID.u.globalGNB_ID.gNB_ID.u.gNB_ID.length=0;
        memset(gnbIdValue_ma, 0xff, sizeof(gnbIdValue_ma));
        recommendedRANNodeList_ma[i].value.aMFPagingTarget.u.globalRANNodeID.u.globalGNB_ID.gNB_ID.u.gNB_ID.value = gnbIdValue_ma;
    }

    infoOnRecommendedCellsAndRANNodesForPaging_mp=NULL;
    isRecommendInfoPrst_m = false;

    //PDUSessionResourceListCxtRelCpl
    pduSessResListCxtRelCpl_mp = NULL;
    pduSessResListCxtRelCpl_m = NULL;
    numofPduSessResListItemCxtRelCpl_m = 0;

    for (unsigned int i = 0; i < maxnoofPDUSessions_c; i++) {
        pduSessResListCxtRelCpl_ma[i].next = NULL;
        pduSessResListCxtRelCpl_ma[i].value.bit_mask = 0;
        pduSessResListCxtRelCpl_ma[i].value.iE_Extensions=NULL;
        pduSessResListCxtRelCpl_ma[i].value.pDUSessionID = 0;
    }
}

unsigned int
UeContextReleaseCompletePdu::encodeOssToPer(NgapOssCtxt* ctxt)
{
    if (NULL == ctxt) {
        return -1;
    }

    ProtocolIeType* nextIe_p = NULL;

    //PDUSessionResourceListCxtRelCpl
    ProtocolIeType PDUSessionResourceListCxtRelCplIe;

    if(numofPduSessResListItemCxtRelCpl_m > maxnoofPDUSessions ||
          numofPduSessResListItemCxtRelCpl_m < 0)
    {
        return defaultFailureCode_c;
    }
    else
    {
        if (numofPduSessResListItemCxtRelCpl_m > 0 )
        {

            struct PDUSessionResourceListCxtRelCpl_ *nextPduSessResListCxtRelCpl_p = NULL;
            for(unsigned int i=0;i<numofPduSessResListItemCxtRelCpl_m;i++)
            {

                pduSessResListCxtRelCpl_ma[i].next=nextPduSessResListCxtRelCpl_p;
                nextPduSessResListCxtRelCpl_p=&(pduSessResListCxtRelCpl_ma[i]);

            }
            pduSessResListCxtRelCpl_m=nextPduSessResListCxtRelCpl_p;
            PDUSessionResourceListCxtRelCplIe.value.value.pduNum=
                PDU_UEContextReleaseComplete_IEs_Value_PDUSessionResourceListCxtRelCpl;
            PDUSessionResourceListCxtRelCplIe.value.value.decoded.pdu_PDUSessionResourceListCxtRelCpl
                = &pduSessResListCxtRelCpl_m;
            PDUSessionResourceListCxtRelCplIe.value.value.encoded.length = 0;
            PDUSessionResourceListCxtRelCplIe.value.value.encoded.value = NULL;
            
            PDUSessionResourceListCxtRelCplIe.next=nextIe_p;
            PDUSessionResourceListCxtRelCplIe.value.id=id_PDUSessionResourceListCxtRelCpl;
            PDUSessionResourceListCxtRelCplIe.value.criticality=reject;
            nextIe_p=&PDUSessionResourceListCxtRelCplIe;
        }
    }

    //InfoOnRecommendedCellsAndRANNodesForPaging
    ProtocolIeType recommendInfoIe;
    if(numOfRecommendedCellItem_m > maxnoofRecommendedCells ||
          numOfRecommendedCellItem_m < 0)
    {
        return defaultFailureCode_c;
    }
    else if(numOfRecommendedRANNodeItem_m > maxnoofRecommendedRANNodes ||
          numOfRecommendedRANNodeItem_m < 0)
    {
        return defaultFailureCode_c;
    }
    else
    {
        if(isRecommendInfoPrst_m)
        {
            struct RecommendedCellList_ *nextRecommendedCellList_p=NULL;
            for(int i=0;i<numOfRecommendedCellItem_m;i++)
            {
                recommendedCellList_ma[i].next=nextRecommendedCellList_p;
                nextRecommendedCellList_p=&recommendedCellList_ma[i];
            }
            infoOnRecommendedCellsAndRANNodesForPaging_m.recommendedCellsForPaging.recommendedCellList
                = nextRecommendedCellList_p;

            struct RecommendedRANNodeList_ *nextRecommendedRANNodeList_p=NULL;
            for(int i=0;i<numOfRecommendedRANNodeItem_m;i++)
            {
                recommendedRANNodeList_ma[i].next=nextRecommendedRANNodeList_p;
                nextRecommendedRANNodeList_p=&recommendedRANNodeList_ma[i];
            }
            infoOnRecommendedCellsAndRANNodesForPaging_m.recommendRANNodesForPaging.recommendedRANNodeList
                = nextRecommendedRANNodeList_p;

            recommendInfoIe.value.value.pduNum =
                PDU_UEContextReleaseComplete_IEs_Value_InfoOnRecommendedCellsAndRANNodesForPaging;
            recommendInfoIe.value.value.decoded.pdu_InfoOnRecommendedCellsAndRANNodesForPaging
                = &infoOnRecommendedCellsAndRANNodesForPaging_m;
            recommendInfoIe.value.value.encoded.length = 0;
            recommendInfoIe.value.value.encoded.value = NULL;

            recommendInfoIe.next = nextIe_p;
            recommendInfoIe.value.id = id_InfoOnRecommendedCellsAndRANNodesForPaging;
            recommendInfoIe.value.criticality = ignore;
            nextIe_p = &recommendInfoIe;
        }
    }

    //User Location Information
    ProtocolIeType useLocationInformationIe;
    if (isUliPrst_m){

        useLocationInformationIe.next =nextIe_p;
        useLocationInformationIe.value.id = id_UserLocationInformation;
        useLocationInformationIe.value.criticality = ignore;
        useLocationInformationIe.value.value.pduNum =PDU_UEContextReleaseComplete_IEs_Value_UserLocationInformation;
        useLocationInformationIe.value.value.decoded.pdu_UserLocationInformation = &useLocationInformation_m;
        useLocationInformationIe.value.value.encoded.length = 0;
        useLocationInformationIe.value.value.encoded.value = NULL;
        nextIe_p = &useLocationInformationIe;
    }

    //RAN_UE_NGAP_ID
    ProtocolIeType ranUeNgapIdIe;
    ranUeNgapIdIe.next = nextIe_p;
    ranUeNgapIdIe.value.id = id_RAN_UE_NGAP_ID;
    ranUeNgapIdIe.value.criticality = ignore;
    ranUeNgapIdIe.value.value.pduNum =
        PDU_UEContextReleaseComplete_IEs_Value_RAN_UE_NGAP_ID;
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
        PDU_UEContextReleaseComplete_IEs_Value_AMF_UE_NGAP_ID;
    amfUeNgapIdIe.value.value.decoded.pdu_AMF_UE_NGAP_ID = &amfUeNgapId_m;
    amfUeNgapIdIe.value.value.encoded.length = 0;
    amfUeNgapIdIe.value.value.encoded.value = NULL;
    nextIe_p = &amfUeNgapIdIe;

    //message
    UEContextReleaseComplete msg;
    msg.protocolIEs = nextIe_p;

    NGAP_PDU ngApPdu;
    ngApPdu.choice = successfulOutcome_chosen;
    ngApPdu.u.successfulOutcome.procedureCode = id_UEContextRelease;
    ngApPdu.u.successfulOutcome.criticality = reject;
    ngApPdu.u.successfulOutcome.value.pduNum = 
        PDU_NGAP_ELEMENTARY_PROCEDURES_SuccessfulOutcome_UEContextReleaseComplete;
    ngApPdu.u.successfulOutcome.value.decoded.pdu_UEContextReleaseComplete = &msg;
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

void UeContextReleaseCompletePdu::getDecodedOssDataPtrs_v()
{
    ProtocolIeType* currIe_p;
    reset_v();

    ngapPdu_mp = reinterpret_cast<NGAP_PDU*>(ossBufRef_m.value);
    if (ngapPdu_mp == NULL) {
        return;
    }

    UeContextReleaseComplete_mp = 
        ngapPdu_mp->u.successfulOutcome.value.decoded.pdu_UEContextReleaseComplete;

    if (UeContextReleaseComplete_mp == NULL) {
        return;
    }

    currIe_p = UeContextReleaseComplete_mp->protocolIEs;

    while (currIe_p != NULL) {
        if (currIe_p->value.id == id_RAN_UE_NGAP_ID
            && currIe_p->value.value.decoded.pdu_RAN_UE_NGAP_ID != NULL) {
            ranUeNgapId_mp = currIe_p->value.value.decoded.pdu_RAN_UE_NGAP_ID;
        } else if (currIe_p->value.id == id_AMF_UE_NGAP_ID
            && currIe_p->value.value.decoded.pdu_AMF_UE_NGAP_ID != NULL) {
            amfUeNgapId_mp = currIe_p->value.value.decoded.pdu_AMF_UE_NGAP_ID;
        }
        else if (currIe_p->value.id == id_UserLocationInformation
                 && currIe_p->value.value.decoded.pdu_UserLocationInformation != NULL)
        {
            useLocationInformation_mp = currIe_p->value.value.decoded.pdu_UserLocationInformation;
            isUliPrst_m = true;
        }
        else if (currIe_p->value.id == id_InfoOnRecommendedCellsAndRANNodesForPaging
                 && currIe_p->value.value.decoded.pdu_InfoOnRecommendedCellsAndRANNodesForPaging != NULL)
        {
            infoOnRecommendedCellsAndRANNodesForPaging_mp=
                  currIe_p->value.value.decoded.pdu_InfoOnRecommendedCellsAndRANNodesForPaging;
            
            RecommendedCellList cellList_p = NULL;
            cellList_p 
                = infoOnRecommendedCellsAndRANNodesForPaging_mp->recommendedCellsForPaging.recommendedCellList;
            while(cellList_p!=NULL)
            {
                numOfRecommendedCellItem_m++;
                cellList_p=cellList_p->next;
            }

            RecommendedRANNodeList RanNodeList_p=NULL;
            RanNodeList_p
                = infoOnRecommendedCellsAndRANNodesForPaging_mp->recommendRANNodesForPaging.recommendedRANNodeList;
            while(RanNodeList_p!=NULL)
            {
                numOfRecommendedRANNodeItem_m++;
                RanNodeList_p=RanNodeList_p->next;
            }

            isRecommendInfoPrst_m = true;
        }
        else if (currIe_p->value.id == id_PDUSessionResourceListCxtRelCpl
            && currIe_p->value.value.decoded.pdu_PDUSessionResourceListCxtRelCpl != NULL)
        {
            pduSessResListCxtRelCpl_mp = 
                      currIe_p->value.value.decoded.pdu_PDUSessionResourceListCxtRelCpl;

            PDUSessionResourceListCxtRelCpl pduSessResSetup_p=NULL;

            pduSessResSetup_p=*pduSessResListCxtRelCpl_mp;
            while(pduSessResSetup_p!=NULL)
            {
                numofPduSessResListItemCxtRelCpl_m++;
                pduSessResSetup_p=pduSessResSetup_p->next;
            }
        }
        currIe_p = currIe_p->next;
    }
}

UEContextReleaseComplete*
UeContextReleaseCompletePdu::getUeContextRelCplDecodedPtr() const
{
    return UeContextReleaseComplete_mp;
}

RAN_UE_NGAP_ID*
UeContextReleaseCompletePdu::getRanUeNgapIdDecodedPtr() const
{
    return ranUeNgapId_mp;
}

RAN_UE_NGAP_ID*
UeContextReleaseCompletePdu::getRanUeNgapIdEncodePtr()
{
    return &ranUeNgapId_m;
}

AMF_UE_NGAP_ID*
UeContextReleaseCompletePdu::getAmfUeNgapIdDecodedPtr() const
{
    return amfUeNgapId_mp;
}

AMF_UE_NGAP_ID*
UeContextReleaseCompletePdu::getAmfUeNgapIdEncodePtr()
{
    return &amfUeNgapId_m;
}

UserLocationInformation * 
UeContextReleaseCompletePdu::getUseLocationInformationDecodedPtr() const
{
    return useLocationInformation_mp;
}
bool 
UeContextReleaseCompletePdu::isUserLocInfoPrst(){
    return isUliPrst_m;
}

UserLocationInformation * 
UeContextReleaseCompletePdu::getUseLocationInformationEncodePtr()
{
    isUliPrst_m = true;
    return &useLocationInformation_m;
}

NGRAN_CGI* 
UeContextReleaseCompletePdu::getNgRanCgiEncodePtr()
{
    return &ngranCgi_m;
}

UserLocationInfomationNRIeExtension* 
UeContextReleaseCompletePdu::getIeExtenIeTypeDecodedPtr()
{
    return &userLocalInfoIe;
}

InfoOnRecommendedCellsAndRANNodesForPaging* 
UeContextReleaseCompletePdu::getInfoOnRecommendedCellsAndRANNodesForPagingDecodedPtr() const
{
    return infoOnRecommendedCellsAndRANNodesForPaging_mp;
}

bool 
UeContextReleaseCompletePdu::isRecommendInfoPrst(){
    return isRecommendInfoPrst_m;
}

InfoOnRecommendedCellsAndRANNodesForPaging* 
UeContextReleaseCompletePdu::getInfoOnRecommendedCellsAndRANNodesForPagingEncodePtr()
{
    isRecommendInfoPrst_m = true;
    return &infoOnRecommendedCellsAndRANNodesForPaging_m;
}

RecommendedCellItem * 
UeContextReleaseCompletePdu::getRecommendedCellItemDecodedPtr(unsigned int index) const
{
    RecommendedCellItem* item_p = NULL;
    RecommendedCellList RecommendedCellList_a = NULL;
    if (infoOnRecommendedCellsAndRANNodesForPaging_mp->recommendedCellsForPaging.recommendedCellList
        != NULL) 
    {
        RecommendedCellList_a 
            = infoOnRecommendedCellsAndRANNodesForPaging_mp->recommendedCellsForPaging.recommendedCellList;
    }

    if (RecommendedCellList_a != NULL
        && (index >= 0 && index < numOfRecommendedCellItem_m)) {
        unsigned int i = 0;
        for (; i < index && RecommendedCellList_a != NULL; i++) {
            RecommendedCellList_a = RecommendedCellList_a->next;
        }
        if (RecommendedCellList_a != NULL && i == index) {
            item_p = &(RecommendedCellList_a->value);
        }
    }

    return item_p;
}

RecommendedCellItem * 
UeContextReleaseCompletePdu::getRecommendedCellItemEncodePtr(unsigned int index)
{
    isRecommendInfoPrst_m = true;
    
    if (index >= 0 && index < maxnoofRecommendedCells) {
        return &(recommendedCellList_ma[index].value);
    } else {
        return NULL;
    }
}

void 
UeContextReleaseCompletePdu::setNumOfRecommendedCellItem_v(unsigned int num)
{
    numOfRecommendedCellItem_m = num;
}

unsigned int 
UeContextReleaseCompletePdu::getNumOfRecommendedCellItem() const
{
    return numOfRecommendedCellItem_m;
}

RecommendedRANNodeItem * 
UeContextReleaseCompletePdu::getRecommendedRANNodeItemDecodedPtr(unsigned int index) const
{
    RecommendedRANNodeItem* item_p = NULL;
    RecommendedRANNodeList RecommendedRANNodeList_a = NULL;
    if (infoOnRecommendedCellsAndRANNodesForPaging_mp->recommendRANNodesForPaging.recommendedRANNodeList
        != NULL) 
    {
        RecommendedRANNodeList_a 
            = infoOnRecommendedCellsAndRANNodesForPaging_mp->recommendRANNodesForPaging.recommendedRANNodeList;
    }

    if (RecommendedRANNodeList_a != NULL
        && (index >= 0 && index < numOfRecommendedRANNodeItem_m)) {
        unsigned int i = 0;
        for (; i < index && RecommendedRANNodeList_a != NULL; i++) {
            RecommendedRANNodeList_a = RecommendedRANNodeList_a->next;
        }
        if (RecommendedRANNodeList_a != NULL && i == index) {
            item_p = &(RecommendedRANNodeList_a->value);
        }
    }

    return item_p;
}

RecommendedRANNodeItem * 
UeContextReleaseCompletePdu::getRecommendedRANNodeItemEncodePtr(unsigned int index)
{
    isRecommendInfoPrst_m = true;

    if (index >= 0 && index < maxnoofRecommendedRANNodes) {
        return &(recommendedRANNodeList_ma[index].value);
    } else {
        return NULL;
    }
}

void 
UeContextReleaseCompletePdu::setNumOfRecommendedRANNodeItem_v(unsigned int num)
{
    numOfRecommendedRANNodeItem_m = num;
}

unsigned int 
UeContextReleaseCompletePdu::getNumOfRecommendedRANNodeItem() const
{
    return numOfRecommendedRANNodeItem_m;
}

PDUSessionResourceListCxtRelCpl*
UeContextReleaseCompletePdu::getPduSessResListCxtRelCplDecodedPtr() const
{
    return pduSessResListCxtRelCpl_mp;
}

PDUSessionResourceListCxtRelCpl*
UeContextReleaseCompletePdu::getPduSessResListCxtRelCplEncodePtr()
{
    return &pduSessResListCxtRelCpl_m;
}

PDUSessionResourceItemCxtRelCpl*
UeContextReleaseCompletePdu::getPduSessResItemCxtRelCplDecodedPtr(unsigned int index) const
{
    PDUSessionResourceItemCxtRelCpl* item_p = NULL;

    PDUSessionResourceListCxtRelCpl pduSessResListRelCpl_p = NULL;

    if (pduSessResListCxtRelCpl_mp != NULL) {
        pduSessResListRelCpl_p = *(pduSessResListCxtRelCpl_mp);
    }

    if (pduSessResListRelCpl_p != NULL
        && (index >= 0 && index < numofPduSessResListItemCxtRelCpl_m)) {
        unsigned int i = 0;
        for (; i < index && pduSessResListRelCpl_p != NULL; i++) {
            pduSessResListRelCpl_p = pduSessResListRelCpl_p->next;
        }
        if (pduSessResListRelCpl_p != NULL && i == index) {
            item_p = &(pduSessResListRelCpl_p->value);
        }
    }

    return item_p;
}

PDUSessionResourceItemCxtRelCpl*
UeContextReleaseCompletePdu::getPduSessResItemCxtRelCplEncodePtr(unsigned int index)
{
    if (index >= 0 && index < maxnoofPDUSessions) {
        return &(pduSessResListCxtRelCpl_ma[index].value);
    } else {
        return NULL;
    }
}

void UeContextReleaseCompletePdu::setNumofPduSessResListItemCxtRelCpl_v(unsigned int num)
{
    numofPduSessResListItemCxtRelCpl_m = num;
}

unsigned int
UeContextReleaseCompletePdu::getNumofPduSessResListItemCxtRelCpl() const
{
    return numofPduSessResListItemCxtRelCpl_m;
}


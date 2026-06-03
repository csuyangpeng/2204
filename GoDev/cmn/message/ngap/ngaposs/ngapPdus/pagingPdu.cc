#include <string.h>
#include <iostream>

#include "pagingPdu.h"

PagingPdu::PagingPdu()
{
    reset_v();
}

PagingPdu::~PagingPdu()
{
}

void
PagingPdu::reset_v()
{
    paging_mp = NULL;
	
    //UEPagingIdentity
    uePagingIdentity_mp = NULL;
    uePagingIdentity_m.choice = fiveG_S_TMSI_chosen;
    memset(&uePagingIdent_ma,0xff,sizeof(uePagingIdent_ma));

    uePagingIdentity_m.u.fiveG_S_TMSI.bit_mask=0;
    uePagingIdentity_m.u.fiveG_S_TMSI.iE_Extensions=NULL;
    uePagingIdentity_m.u.fiveG_S_TMSI.aMFSetID.length=0;
    uePagingIdentity_m.u.fiveG_S_TMSI.aMFSetID.value=uePagingIdent_ma.setId;
    uePagingIdentity_m.u.fiveG_S_TMSI.aMFPointer.length=0;
    uePagingIdentity_m.u.fiveG_S_TMSI.aMFPointer.value=&(uePagingIdent_ma.pointer);
    uePagingIdentity_m.u.fiveG_S_TMSI.fiveG_TMSI.length=0;
    memset(uePagingIdentity_m.u.fiveG_S_TMSI.fiveG_TMSI.value, 0xff, 
           sizeof(uePagingIdentity_m.u.fiveG_S_TMSI.fiveG_TMSI.value));

    //Paging DRX
    pagingDrx_m = v32;
    pagingDrx_mp = NULL;
    pagingDrxPresent_m = false;

    //TAIListForPaging
    taiListForPaging_mp = NULL;
    taiListForPaging_m = NULL;
    numofTAIListForPagingItem_m = 0;
    for(unsigned int i = 0; i < maxnoofTAIforPaging; i++)
    {
        taiListForPaging_ma[i].next=NULL;
        taiListForPaging_ma[i].value.bit_mask=0; 
        taiListForPaging_ma[i].value.iE_Extensions=NULL;
        taiListForPaging_ma[i].value.tAI.bit_mask = 0; 
        taiListForPaging_ma[i].value.tAI.iE_Extensions=NULL;
        taiListForPaging_ma[i].value.tAI.pLMNIdentity.length = 0;
        memset(taiListForPaging_ma[i].value.tAI.pLMNIdentity.value, 0xff, 
            sizeof(taiListForPaging_ma[i].value.tAI.pLMNIdentity.value));

        taiListForPaging_ma[i].value.tAI.tAC.length = 0;
        memset(taiListForPaging_ma[i].value.tAI.tAC.value, 0xff, 
            sizeof(taiListForPaging_ma[i].value.tAI.tAC.value));    
    }

    //Paging Priority
    pagingPriority_m = priolevel1;
    pagingPriority_mp=NULL;
    pagingPriorityPresent_m = false;

    //UERadioCapabilityForPaging
    ueRadioCapabilityForPaging_mp = NULL;
    ueRadioCapabilityForPaging_m.uERadioCapabilityForPagingOfNR.length = 0;
    memset(ueRadioCapabilityForPagingNr_ma, 0xff, sizeof(ueRadioCapabilityForPagingNr_ma));
	memset(ueRadioCapabilityForPagingEutra_ma, 0xff, sizeof(ueRadioCapabilityForPagingEutra_ma));
    ueRadioCapabilityForPaging_m.uERadioCapabilityForPagingOfNR.value = ueRadioCapabilityForPagingNr_ma;
	ueRadioCapabilityForPaging_m.uERadioCapabilityForPagingOfEUTRA.length = 0;
	ueRadioCapabilityForPaging_m.uERadioCapabilityForPagingOfEUTRA.value = ueRadioCapabilityForPagingEutra_ma;
    ueRadioCapabilityForPagingPresent_m = false;

    //PagingOrigin
    pagingOrigin_mp = NULL;
    pagingOrigin_m = non_3gpp;
    pagingOriginPresent_m = false;

    //AssistanceDataForPaging
    assDataForPaging_mp = NULL;
    assDataForPaging_m.bit_mask = 0;
    assDataForPaging_m.iE_Extensions = NULL;
    assDataForPaging_m.assistanceDataForRecommendedCells.bit_mask = 0;
    assDataForPaging_m.assistanceDataForRecommendedCells.iE_Extensions = NULL;
    assDataForPaging_m.assistanceDataForRecommendedCells.recommendedCellsForPaging
        .bit_mask=0;
    assDataForPaging_m.assistanceDataForRecommendedCells.recommendedCellsForPaging
        .iE_Extensions=NULL;
    assDataForPaging_m.assistanceDataForRecommendedCells.recommendedCellsForPaging
        .recommendedCellList=NULL;
    numOfRecommendedCellItem_m = 0;
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
    assDataForPaging_m.pagingAttemptInformation.bit_mask = 0;
    assDataForPaging_m.pagingAttemptInformation.iE_Extensions = NULL;
    assDataForPaging_m.pagingAttemptInformation.pagingAttemptCount = 0;
    assDataForPaging_m.pagingAttemptInformation.intendedNumberOfPagingAttempts = 0;
    assDataForPaging_m.pagingAttemptInformation.nextPagingAreaScope = same;

    assDataForPagingPresent_m = false;
}

unsigned int
PagingPdu::encodeOssToPer(NgapOssCtxt* ctxt)
{
    if (NULL == ctxt) {
        return -1;
    }

    ProtocolIeType* nextIe_p = NULL;

    //PagingOrigin
    ProtocolIeType pagingOriginIe;
    if(pagingOriginPresent_m)
    {
        pagingOriginIe.next = nextIe_p;
        pagingOriginIe.value.id = id_PagingOrigin;
        pagingOriginIe.value.criticality = ignore;
        pagingOriginIe.value.value.pduNum = PDU_PagingIEs_Value_PagingOrigin;
        pagingOriginIe.value.value.decoded.pdu_PagingOrigin = &pagingOrigin_m;
        pagingOriginIe.value.value.encoded.length = 0;
        pagingOriginIe.value.value.encoded.value = NULL;
        nextIe_p = &pagingOriginIe;
    }

    //AssistanceDataForPaging
    ProtocolIeType assDataForPagingIe;
    if(assDataForPagingPresent_m)
    {
        if(numOfRecommendedCellItem_m > maxnoofRecommendedCells_c ||
              numOfRecommendedCellItem_m < 0)
        {
            return defaultFailureCode_c;
        }
        else
        {
            struct RecommendedCellList_ *nextRecommendedCellList_p=NULL;
            for(int i=0;i<numOfRecommendedCellItem_m;i++){
                recommendedCellList_ma[i].next = nextRecommendedCellList_p;
                nextRecommendedCellList_p = &recommendedCellList_ma[i];
            }
            assDataForPaging_m.assistanceDataForRecommendedCells.recommendedCellsForPaging
                .recommendedCellList = nextRecommendedCellList_p;

            assDataForPagingIe.value.value.pduNum = PDU_PagingIEs_Value_AssistanceDataForPaging;
            assDataForPagingIe.value.value.decoded.pdu_AssistanceDataForPaging = &assDataForPaging_m;
            assDataForPagingIe.value.value.encoded.length = 0;
            assDataForPagingIe.value.value.encoded.value = NULL;
        }
        assDataForPagingIe.next = nextIe_p;
        assDataForPagingIe.value.id = id_AssistanceDataForPaging;
        assDataForPagingIe.value.criticality = ignore;
        nextIe_p = &assDataForPagingIe;
    }

    //UERadioCapabilityForPaging
    ProtocolIeType ueRadioCapabilityForPagingIe;
    if(ueRadioCapabilityForPagingPresent_m)
    {
        ueRadioCapabilityForPagingIe.next = nextIe_p;
        ueRadioCapabilityForPagingIe.value.id = id_UERadioCapabilityForPaging;
        ueRadioCapabilityForPagingIe.value.criticality = ignore;
        ueRadioCapabilityForPagingIe.value.value.pduNum =
            PDU_PagingIEs_Value_UERadioCapabilityForPaging;
        ueRadioCapabilityForPagingIe.value.value.decoded.pdu_UERadioCapabilityForPaging =
            &ueRadioCapabilityForPaging_m;
        ueRadioCapabilityForPagingIe.value.value.encoded.length = 0;
        ueRadioCapabilityForPagingIe.value.value.encoded.value = NULL;
        nextIe_p = &ueRadioCapabilityForPagingIe;
    }

    //Paging Priority Encode
    ProtocolIeType pagingPriorityIe;
    if(pagingPriorityPresent_m)
    {
        pagingPriorityIe.next =nextIe_p;
        pagingPriorityIe.value.id = id_PagingPriority;
        pagingPriorityIe.value.criticality = ignore;
        pagingPriorityIe.value.value.pduNum = PDU_PagingIEs_Value_PagingPriority;
        pagingPriorityIe.value.value.decoded.pdu_PagingPriority = &pagingPriority_m;
        pagingPriorityIe.value.value.encoded.length = 0;
        pagingPriorityIe.value.value.encoded.value = NULL;
        nextIe_p = &pagingPriorityIe;
    }

    //TAI List For Paging encode
    ProtocolIeType taiListForPagingIe;
    if (numofTAIListForPagingItem_m > maxnoofTAIforPaging || numofTAIListForPagingItem_m <= 0)
    {
        return defaultFailureCode_c;
    }
    else
    {
        struct TAIListForPaging_* nextTaiListForPaging_p = NULL;
        for (int i = 0; i < numofTAIListForPagingItem_m; i++)
        {
            taiListForPaging_ma[i].next = nextTaiListForPaging_p;
            nextTaiListForPaging_p = &taiListForPaging_ma[i];
        }

        taiListForPaging_m = nextTaiListForPaging_p;
    }
    taiListForPagingIe.next = nextIe_p;
    taiListForPagingIe.value.id = id_TAIListForPaging;
    taiListForPagingIe.value.criticality = ignore;
    taiListForPagingIe.value.value.pduNum = PDU_PagingIEs_Value_TAIListForPaging;
    taiListForPagingIe.value.value.decoded.pdu_TAIListForPaging = &taiListForPaging_m;
    taiListForPagingIe.value.value.encoded.length = 0;
    taiListForPagingIe.value.value.encoded.value = NULL;
    nextIe_p = &taiListForPagingIe;

    //Paging DRX IE encode
    ProtocolIeType pagingDrxIe;
    if(pagingDrxPresent_m)
    {
        pagingDrxIe.next = nextIe_p;
        pagingDrxIe.value.id = id_PagingDRX;
        pagingDrxIe.value.criticality = ignore;
        pagingDrxIe.value.value.pduNum = PDU_PagingIEs_Value_PagingDRX;
        pagingDrxIe.value.value.decoded.pdu_PagingDRX = &pagingDrx_m;
        pagingDrxIe.value.value.encoded.length = 0;
        pagingDrxIe.value.value.encoded.value = NULL;
        nextIe_p = &pagingDrxIe;
    }

    //UEPagingIdentity IE encode
    ProtocolIeType uePagingIdentityIe;
    uePagingIdentityIe.next=nextIe_p;
    uePagingIdentityIe.value.id=id_UEPagingIdentity;
    uePagingIdentityIe.value.criticality=ignore;
    uePagingIdentityIe.value.value.pduNum=PDU_PagingIEs_Value_UEPagingIdentity;
    uePagingIdentityIe.value.value.decoded.pdu_UEPagingIdentity = &uePagingIdentity_m;
    uePagingIdentityIe.value.value.encoded.length=0;
    uePagingIdentityIe.value.value.encoded.value=NULL;
    nextIe_p=&uePagingIdentityIe;

    //message
    Paging msg;
    msg.protocolIEs = nextIe_p;

    NGAP_PDU ngApPdu;

    ngApPdu.choice = initiatingMessage_chosen;
    ngApPdu.u.initiatingMessage.procedureCode = id_Paging;
    ngApPdu.u.initiatingMessage.criticality = ignore;
    ngApPdu.u.initiatingMessage.value.pduNum = 
        PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_Paging;
    ngApPdu.u.initiatingMessage.value.decoded.pdu_Paging = &msg;
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

void PagingPdu::getDecodedOssDataPtrs_v()
{
    ProtocolIeType* currIe_p;
    reset_v();

    ngapPdu_mp = reinterpret_cast<NGAP_PDU*>(ossBufRef_m.value);
    if (ngapPdu_mp == NULL) {
        return;
    }

    paging_mp = 
        ngapPdu_mp->u.initiatingMessage.value.decoded.pdu_Paging;

    if (paging_mp == NULL) {
        return;
    }

    currIe_p = paging_mp->protocolIEs;

    while (currIe_p != NULL) {
        if (currIe_p->value.id == id_UEPagingIdentity
            && currIe_p->value.value.decoded.pdu_UEPagingIdentity != NULL) {
            uePagingIdentity_mp = currIe_p->value.value.decoded.pdu_UEPagingIdentity;
        } else if (currIe_p->value.id == id_PagingDRX
            && currIe_p->value.value.decoded.pdu_PagingDRX != NULL) {
            pagingDrx_mp = currIe_p->value.value.decoded.pdu_PagingDRX;
            pagingDrxPresent_m = true;
        } else if (currIe_p->value.id == id_TAIListForPaging
            && currIe_p->value.value.decoded.pdu_TAIListForPaging != NULL) {
            taiListForPaging_mp = currIe_p->value.value.decoded.pdu_TAIListForPaging;
            TAIListForPaging curr_p = *taiListForPaging_mp;
            while (curr_p != NULL) {
                numofTAIListForPagingItem_m++;
                curr_p = curr_p->next;
            }
        } else if (currIe_p->value.id == id_PagingPriority
            && currIe_p->value.value.decoded.pdu_PagingPriority != NULL) {
            pagingPriority_mp = currIe_p->value.value.decoded.pdu_PagingPriority;
            pagingPriorityPresent_m = true;
        } else if (currIe_p->value.id == id_UERadioCapabilityForPaging
            && currIe_p->value.value.decoded.pdu_UERadioCapabilityForPaging != NULL) {
            ueRadioCapabilityForPaging_mp = 
                currIe_p->value.value.decoded.pdu_UERadioCapabilityForPaging;
            ueRadioCapabilityForPagingPresent_m = true;
        } else if (currIe_p->value.id == id_PagingOrigin
            && currIe_p->value.value.decoded.pdu_PagingOrigin != NULL) {
            pagingOrigin_mp = currIe_p->value.value.decoded.pdu_PagingOrigin;
            pagingOriginPresent_m = true;
        } else if (currIe_p->value.id == id_AssistanceDataForPaging
            && currIe_p->value.value.decoded.pdu_AssistanceDataForPaging != NULL) {
            assDataForPaging_mp = 
                      currIe_p->value.value.decoded.pdu_AssistanceDataForPaging;
            
            RecommendedCellList cellList_p = assDataForPaging_mp->assistanceDataForRecommendedCells
                .recommendedCellsForPaging.recommendedCellList;
            while(cellList_p!=NULL)
            {
                numOfRecommendedCellItem_m++;  
                cellList_p = cellList_p->next;
            }
            assDataForPagingPresent_m = true;
        }
        currIe_p = currIe_p->next;
    }
}

Paging*
PagingPdu::getPagingDecodedPtr() const
{
    return paging_mp;
}

UEPagingIdentity*
PagingPdu::getUePagingIdentityDecodedPtr() const
{
    return uePagingIdentity_mp;
}

UEPagingIdentity*
PagingPdu::getUePagingIdentityEncodePtr()
{
    return &uePagingIdentity_m;
}

PagingDRX*
PagingPdu::getPagingDrxDecodedPtr() const
{
    return pagingDrx_mp;
}

PagingDRX*
PagingPdu::getPagingDrxEncodePtr()
{
    pagingDrxPresent_m = true;
    return &pagingDrx_m;
}

bool 
PagingPdu::isPagingDrxPresent()
{
    return pagingDrxPresent_m;
}

TAIListForPaging*
PagingPdu::getTAIListForPagingDecodedPtr() const
{
    return taiListForPaging_mp;
}

TAIListForPaging*
PagingPdu::getTAIListForPagingEncodePtr()
{
    return &taiListForPaging_m;
}

TAIListForPagingItem*
PagingPdu::getTAIListForPagingItemDecodedPtr(unsigned int index) const
{
    TAIListForPagingItem* item_p = NULL;

    TAIListForPaging taiListForPaging_p = NULL;

    if (taiListForPaging_mp != NULL) {
        taiListForPaging_p = *(taiListForPaging_mp);
    }

    if (taiListForPaging_p != NULL
        && (index >= 0 && index < numofTAIListForPagingItem_m)) {
        unsigned int i = 0;
        for (; i < index && taiListForPaging_p != NULL; i++) {
            taiListForPaging_p = taiListForPaging_p->next;
        }
        if (taiListForPaging_p != NULL && i == index) {
            item_p = &(taiListForPaging_p->value);
        }
    }

    return item_p;
}

TAIListForPagingItem*
PagingPdu::getTAIListForPagingItemEncodePtr(unsigned int index)
{
    if (index >= 0 && index < maxnoofTAIforPaging) {
        return &(taiListForPaging_ma[index].value);
    } else {
        return NULL;
    }
}

void PagingPdu::setNumofTAIListForPagingItem_v(unsigned int num)
{
    numofTAIListForPagingItem_m = num;
}

unsigned int
PagingPdu::getNumofTAIListForPagingItem() const
{
    return numofTAIListForPagingItem_m;
}

PagingPriority*
PagingPdu::getPagePriorityDecodedPtr() const
{
    return pagingPriority_mp;
}

PagingPriority*
PagingPdu::getPagePriorityEncodePtr()
{
    pagingPriorityPresent_m = true;
    return &pagingPriority_m;
}

bool 
PagingPdu::isPagePriorityPresent()
{
    return pagingPriorityPresent_m;
}

UERadioCapabilityForPaging*
PagingPdu::getUeRadioCapForPagingDecodedPtr() const
{
    return ueRadioCapabilityForPaging_mp;
}

UERadioCapabilityForPaging*
PagingPdu::getUeRadioCapForPagingEncodePtr()
{
    ueRadioCapabilityForPagingPresent_m = true;
    return &ueRadioCapabilityForPaging_m;
}

bool 
PagingPdu::isUeRadioCapabilityForPagingPresent()
{
    return ueRadioCapabilityForPagingPresent_m;
}

PagingOrigin*
PagingPdu::getPagingOriginDecodedPtr() const
{
    return pagingOrigin_mp;
}

PagingOrigin*
PagingPdu::getPagingOriginEncodePtr()
{
    pagingOriginPresent_m = true;
    return &pagingOrigin_m;
}

bool 
PagingPdu::isPagingOriginPresent()
{
    return pagingOriginPresent_m;
}

AssistanceDataForPaging*
PagingPdu::getAssDataForPagingDecodedPtr() const
{
    return assDataForPaging_mp;
}

AssistanceDataForPaging*
PagingPdu::getAssDataForPagingEncodePtr()
{
    assDataForPagingPresent_m = true;
    return &assDataForPaging_m;
}

bool 
PagingPdu::isAssDataForPagingPresent()
{
    return assDataForPagingPresent_m;
}

RecommendedCellItem*
PagingPdu::getRecommendedCellItemDecodedPtr(unsigned int index) const
{
    RecommendedCellItem* item_p = NULL;
    RecommendedCellList RecommendedCellList_a = NULL;
    if (assDataForPaging_mp->assistanceDataForRecommendedCells.recommendedCellsForPaging
        .recommendedCellList != NULL) {
        RecommendedCellList_a = assDataForPaging_mp->assistanceDataForRecommendedCells
            .recommendedCellsForPaging.recommendedCellList;
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

RecommendedCellItem*
PagingPdu::getRecommendedCellItemEncodePtr(unsigned int index)
{
    if (index >= 0 && index < maxnoofRecommendedCells) {
        return &(recommendedCellList_ma[index].value);
    } else {
        return NULL;
    }
}

void PagingPdu::setNumOfRecommendedCellItem_v(unsigned int num)
{
    numOfRecommendedCellItem_m = num;
}

unsigned int
PagingPdu::getNumOfRecommendedCellItem() const
{
    return numOfRecommendedCellItem_m;
}

#include <string.h>
#include <iostream>

#include "pduSessionResourceModifyResponseTransferPdu.h"

PduSesResModifyRespTransferPdu::PduSesResModifyRespTransferPdu()
{
    reset_v();
}

PduSesResModifyRespTransferPdu::~PduSesResModifyRespTransferPdu()
{
}

void PduSesResModifyRespTransferPdu::reset_v()
{
    pduSessResModifyRespTransfer_mp = NULL;

    //dl up transport layer information
    memset(&transLayerAddress_m, 0, sizeof(transLayerAddress_m));

    isDlUpTransportLayerInfoPrst_m = false;
    dlUpTransportLayerInfo_mp = NULL;
    dlUpTransportLayerInfo_m.choice = gTPTunnel_chosen;
    dlUpTransportLayerInfo_m.u.gTPTunnel.bit_mask = 0;
    dlUpTransportLayerInfo_m.u.gTPTunnel.iE_Extensions = NULL;
    dlUpTransportLayerInfo_m.u.gTPTunnel.transportLayerAddress.length = 0;
    dlUpTransportLayerInfo_m.u.gTPTunnel.transportLayerAddress.value = transLayerAddress_m.dlTransLayerAddr;
    dlUpTransportLayerInfo_m.u.gTPTunnel.gTP_TEID.length = 0;
    memset(dlUpTransportLayerInfo_m.u.gTPTunnel.gTP_TEID.value, 0, sizeof(dlUpTransportLayerInfo_m.u.gTPTunnel.gTP_TEID.value));

    //ul up transport layer information
    isUlUpTransportLayerInfoPrst_m = false;
    ulUpTransportLayerInfo_mp = NULL;
    ulUpTransportLayerInfo_m.choice = gTPTunnel_chosen;
    ulUpTransportLayerInfo_m.u.gTPTunnel.bit_mask = 0;
    ulUpTransportLayerInfo_m.u.gTPTunnel.iE_Extensions = NULL;
    ulUpTransportLayerInfo_m.u.gTPTunnel.transportLayerAddress.length = 0;
    ulUpTransportLayerInfo_m.u.gTPTunnel.transportLayerAddress.value = transLayerAddress_m.ulTransLayerAddr;
    ulUpTransportLayerInfo_m.u.gTPTunnel.gTP_TEID.length = 0;
    memset(ulUpTransportLayerInfo_m.u.gTPTunnel.gTP_TEID.value,0,sizeof(ulUpTransportLayerInfo_m.u.gTPTunnel.gTP_TEID.value));

    // additional qos flow per tunnel info list
    qosFlowPerTNLInformationList_mp = NULL;
    qosFlowPerTNLInformationList_m = NULL;
    isAddqosFlowPerTnlInfoPrst_m = false;
    numOfQosFlowPerTNLInformationListItem_m = 0;
    for (unsigned int i = 0; i < maxnoofMultiConnectivityMinusOne; i++) {
        qosFlowPerTNLInformationList_ma[i].next = NULL;
        qosFlowPerTNLInformationList_ma[i].value.bit_mask = 0;
        qosFlowPerTNLInformationList_ma[i].value.iE_Extensions = NULL;
        qosFlowPerTNLInformationList_ma[i].value.qosFlowPerTNLInformation.uPTransportLayerInformation.choice = gTPTunnel_chosen;
        qosFlowPerTNLInformationList_ma[i].value.qosFlowPerTNLInformation.uPTransportLayerInformation.u.gTPTunnel.bit_mask = 0;
        qosFlowPerTNLInformationList_ma[i].value.qosFlowPerTNLInformation.uPTransportLayerInformation.u.gTPTunnel.iE_Extensions = NULL;
        qosFlowPerTNLInformationList_ma[i].value.qosFlowPerTNLInformation.uPTransportLayerInformation.u.gTPTunnel.transportLayerAddress.length = 0;
        qosFlowPerTNLInformationList_ma[i].value.qosFlowPerTNLInformation.uPTransportLayerInformation.u.gTPTunnel.transportLayerAddress.value = assTransLayerAddress_m[i].addTransLayerAddr;
        qosFlowPerTNLInformationList_ma[i].value.qosFlowPerTNLInformation.uPTransportLayerInformation.u.gTPTunnel.gTP_TEID.length = 0;
        memset(qosFlowPerTNLInformationList_ma[i].value.qosFlowPerTNLInformation.uPTransportLayerInformation.u.gTPTunnel.gTP_TEID.value,
               0,
               sizeof(qosFlowPerTNLInformationList_ma[i].value.qosFlowPerTNLInformation.uPTransportLayerInformation.u.gTPTunnel.gTP_TEID.value));
        qosFlowPerTNLInformationList_ma[i].value.qosFlowPerTNLInformation.associatedQosFlowList = NULL;
        numofAddQosFlowInd_m[i] = 0;
        for (unsigned int j = 0; j < maxnoofQosFlows; j++) {
            addAssQosFlowList_ma[i][j].next = NULL;
            addAssQosFlowList_ma[i][j].value.bit_mask = 0;
            addAssQosFlowList_ma[i][j].value.iE_Extensions = NULL;
            addAssQosFlowList_ma[i][j].value.qosFlowIdentifier = 0;
            addAssQosFlowList_ma[i][j].value.qosFlowMappingIndication = ul;
        }
    }

    //qos flow add or mod response list
    qosFlowAddOrModifyRespList_mp = NULL;
    qosFlowAddOrModifyRespList_m = NULL;
    isQosFlowAddOrModifyResponsePrst_m = false;
    numofQosFlowAddOrModifyResp_m = 0;
    for (unsigned int i = 0; i < maxnoofQosFlows; i++)
    {
        qosFlowAddOrModifyRespList_ma[i].value.bit_mask = 0;
        qosFlowAddOrModifyRespList_ma[i].value.iE_Extensions = NULL;
        qosFlowAddOrModifyRespList_ma[i].value.qosFlowIdentifier = 0;
    }
    
    //qos flow failed to modify list
    qosFlowToReleaseList_mp = NULL;
    qosFlowToReleaseList_m = NULL;
    isQosFlowListPrst_m = false;
    numofQosFlowToRelease_m = 0;
    for (unsigned int i = 0; i < maxnoofQosFlows; i++)
    {
        qosFlowToReleaseList_ma[i].value.bit_mask = 0;
        qosFlowToReleaseList_ma[i].value.iE_Extensions = NULL;
        
        qosFlowToReleaseList_ma[i].value.qosFlowIdentifier = 0;
        qosFlowToReleaseList_ma[i].value.cause.choice = misc_chosen;    
        qosFlowToReleaseList_ma[i].value.cause.u.misc = CauseMisc_unspecified;
    }
}

unsigned int 
PduSesResModifyRespTransferPdu::decodePerToOss(NgapOssCtxt* ctxt)
{
   if(NULL == ctxt)
   {
       return -1;
   }
   
   OssBuf inputBuf;
   OssBuf outputBuf;
   int pduNum = PDUSessionResourceModifyResponseTransfer_PDU;

   inputBuf = perBufRef_m;

   ctxt->reset_v();
   outputBuf = ctxt->getPrimaryCodecBuf();

   // default action -- automatically fully decode the message
   unsigned int result = ctxt->decodePerToOss(&pduNum, &inputBuf, &outputBuf);

   if (result == 0)
   {
      ossBufRef_m = outputBuf;
      getDecodedOssDataPtrs_v();
   }

   return result;
}

unsigned int
PduSesResModifyRespTransferPdu::encodeOssToPer(NgapOssCtxt *ctxt)
{
    if (NULL == ctxt){
        return -1;
    }

    PDUSessionResourceModifyResponseTransfer ngApPdu;
    ngApPdu.iE_Extensions = NULL;
    ngApPdu.bit_mask = 0;

    //dl up transport layer information
    if(isDlUpTransportLayerInfoPrst_m) {
        ngApPdu.bit_mask |= dL_NGU_UP_TNLInformation_present;
        memcpy(&ngApPdu.dL_NGU_UP_TNLInformation,&dlUpTransportLayerInfo_m,sizeof(dlUpTransportLayerInfo_m));
    }

    //ul up transport layer information
    if(isUlUpTransportLayerInfoPrst_m) {
        ngApPdu.bit_mask |= PDUSessionResourceModifyResponseTransfer_uL_NGU_UP_TNLInformation_present;
        memcpy(&ngApPdu.uL_NGU_UP_TNLInformation,&ulUpTransportLayerInfo_m,sizeof(ulUpTransportLayerInfo_m));
    }

    // additional qos flow per tul information list
    if(isAddqosFlowPerTnlInfoPrst_m) {
        ngApPdu.bit_mask |= PDUSessionResourceModifyResponseTransfer_additionalDLQosFlowPerTNLInformation_present;

        if (numOfQosFlowPerTNLInformationListItem_m > maxnoofMultiConnectivityMinusOne_c ||
            numOfQosFlowPerTNLInformationListItem_m <= 0) {
            return defaultFailureCode_c;
        }
        struct QosFlowPerTNLInformationList_ *nextQosFlowTNLInfo_p = NULL;
        for (unsigned int i = 0; i < numOfQosFlowPerTNLInformationListItem_m; i++) {
            struct AssociatedQosFlowList_ *AssQosFlow_p = NULL;
            for(unsigned int j = 0; j <numofAddQosFlowInd_m[i]; j++) {
                addAssQosFlowList_ma[i][j].next = AssQosFlow_p;
                AssQosFlow_p = &addAssQosFlowList_ma[i][j];
            }
            qosFlowPerTNLInformationList_ma[i].value.qosFlowPerTNLInformation.associatedQosFlowList = AssQosFlow_p;
            qosFlowPerTNLInformationList_ma[i].next = nextQosFlowTNLInfo_p;
            nextQosFlowTNLInfo_p = &qosFlowPerTNLInformationList_ma[i];
        }
        ngApPdu.additionalDLQosFlowPerTNLInformation = nextQosFlowTNLInfo_p;
    }

    //qos flow add or modify response list
    if(isQosFlowAddOrModifyResponsePrst_m){
        ngApPdu.bit_mask |= qosFlowAddOrModifyResponseList_present;

        if(numofQosFlowAddOrModifyResp_m > maxnoofQosFlows || numofQosFlowAddOrModifyResp_m <=0){
            return defaultFailureCode_c;
        }

        struct QosFlowAddOrModifyResponseList_ *nextQosFlowList_p = NULL;
        for(int i=0; i<numofQosFlowAddOrModifyResp_m; i++){
           qosFlowAddOrModifyRespList_ma[i].next = nextQosFlowList_p;
           nextQosFlowList_p = &qosFlowAddOrModifyRespList_ma[i];
        }
        ngApPdu.qosFlowAddOrModifyResponseList = nextQosFlowList_p; 
    }

    //qos flow to modify list
    if(isQosFlowListPrst_m){
        ngApPdu.bit_mask |= qosFlowFailedToAddOrModifyList_present;

        if(numofQosFlowToRelease_m > maxnoofQosFlows || numofQosFlowToRelease_m <=0){
            return defaultFailureCode_c;
        }

        struct QosFlowListWithCause_ *nextQosFlowList_p = NULL;
        for(int i=0; i<numofQosFlowToRelease_m; i++){
           qosFlowToReleaseList_ma[i].next = nextQosFlowList_p;
           nextQosFlowList_p = &qosFlowToReleaseList_ma[i];
        }
        ngApPdu.qosFlowFailedToAddOrModifyList = nextQosFlowList_p; 
    }

    //Convert Structures to PER
    OssBuf outputBuf;
    unsigned int result;

    ctxt->reset_v();
    outputBuf = ctxt->getPrimaryCodecBuf();

    result = ctxt->encodeOssToPer(PDUSessionResourceModifyResponseTransfer_PDU, &ngApPdu, &outputBuf);

    if (result == encodeSuccessCode_c){
        perBufRef_m = outputBuf;
    }

    return result;
}

void PduSesResModifyRespTransferPdu::getDecodedOssDataPtrs_v()
{
    reset_v();

    pduSessResModifyRespTransfer_mp = reinterpret_cast<PDUSessionResourceModifyResponseTransfer *>(ossBufRef_m.value);
    if (pduSessResModifyRespTransfer_mp == NULL){
        return;
    }

    //dl up transport layer information
    if (pduSessResModifyRespTransfer_mp->bit_mask & dL_NGU_UP_TNLInformation_present) {
        isDlUpTransportLayerInfoPrst_m = true;
        dlUpTransportLayerInfo_mp = &pduSessResModifyRespTransfer_mp->dL_NGU_UP_TNLInformation;
    }

    //ul up transport layer information
    if (pduSessResModifyRespTransfer_mp->bit_mask & PDUSessionResourceModifyResponseTransfer_uL_NGU_UP_TNLInformation_present) {
        isUlUpTransportLayerInfoPrst_m = true;
        ulUpTransportLayerInfo_mp = &pduSessResModifyRespTransfer_mp->uL_NGU_UP_TNLInformation;
    }

    //additional qos flow per tul information list
    if(pduSessResModifyRespTransfer_mp ->bit_mask & PDUSessionResourceModifyResponseTransfer_additionalDLQosFlowPerTNLInformation_present) {
        isAddqosFlowPerTnlInfoPrst_m = true;
        qosFlowPerTNLInformationList_mp = &pduSessResModifyRespTransfer_mp->additionalDLQosFlowPerTNLInformation;
        struct QosFlowPerTNLInformationList_ *qosFlowPerTNLInfoPtr = 
                                                           pduSessResModifyRespTransfer_mp->additionalDLQosFlowPerTNLInformation;
        struct AssociatedQosFlowList_ * AssQosFlowPtr = NULL;
        while(qosFlowPerTNLInfoPtr != NULL) {
            AssQosFlowPtr = qosFlowPerTNLInfoPtr->value.qosFlowPerTNLInformation.associatedQosFlowList;
            while(AssQosFlowPtr != NULL) {
                numofAddQosFlowInd_m[numOfQosFlowPerTNLInformationListItem_m]++;
                AssQosFlowPtr = AssQosFlowPtr->next;
            }
            numOfQosFlowPerTNLInformationListItem_m++;
            qosFlowPerTNLInfoPtr = qosFlowPerTNLInfoPtr->next;
        }
    }

    //qos flow add or modify response list
    if(pduSessResModifyRespTransfer_mp->bit_mask & qosFlowAddOrModifyResponseList_present)
    {
        isQosFlowAddOrModifyResponsePrst_m = true;
        qosFlowAddOrModifyRespList_mp = &pduSessResModifyRespTransfer_mp->qosFlowAddOrModifyResponseList;
        struct QosFlowAddOrModifyResponseList_ * nextPtr = pduSessResModifyRespTransfer_mp->qosFlowAddOrModifyResponseList;
        while (nextPtr != NULL)
        {
           nextPtr = nextPtr->next;
           numofQosFlowAddOrModifyResp_m++;
        }
    }

    //qos flow to modify list
    if(pduSessResModifyRespTransfer_mp->bit_mask & qosFlowFailedToAddOrModifyList_present)
    {
        isQosFlowListPrst_m = true;
        qosFlowToReleaseList_mp = &pduSessResModifyRespTransfer_mp->qosFlowFailedToAddOrModifyList;
        struct QosFlowListWithCause_ * nextPtr = pduSessResModifyRespTransfer_mp->qosFlowFailedToAddOrModifyList;
        while (nextPtr != NULL)
        {
           nextPtr = nextPtr->next;
           numofQosFlowToRelease_m++;
        }
    }
}

//----------------------------------------------------------------------------
// This is a public API. Refer to PduSesResModifyRespTransferPdu.h
//----------------------------------------------------------------------------
PDUSessionResourceModifyResponseTransfer *
PduSesResModifyRespTransferPdu::getPduSessResModifyRespTransferDecodedPtr() const
{
    return pduSessResModifyRespTransfer_mp;
}

//----------------------------------------------------------------------------
// This is a public API. Refer to PduSesResModifyRespTransferPdu.h
//----------------------------------------------------------------------------
UPTransportLayerInformation *
PduSesResModifyRespTransferPdu::getDlUpTNLInformationDecodedPtr() const
{
    return dlUpTransportLayerInfo_mp;
}

UPTransportLayerInformation *
PduSesResModifyRespTransferPdu::getDlUpTNLInformationEncodePtr()
{
    isDlUpTransportLayerInfoPrst_m = true;
    return &dlUpTransportLayerInfo_m;
}

bool 
PduSesResModifyRespTransferPdu::isDlUpTNLInformationPrst()
{
    return isDlUpTransportLayerInfoPrst_m;
}

//----------------------------------------------------------------------------
// This is a public API. Refer to PduSesResModifyRespTransferPdu.h
//----------------------------------------------------------------------------
UPTransportLayerInformation *
PduSesResModifyRespTransferPdu::getUlUpTNLInformationDecodedPtr() const
{
    return ulUpTransportLayerInfo_mp;
}

UPTransportLayerInformation *
PduSesResModifyRespTransferPdu::getUlUpTNLInformationEncodePtr()
{
    isUlUpTransportLayerInfoPrst_m = true;
    return &ulUpTransportLayerInfo_m;
}

bool 
PduSesResModifyRespTransferPdu::isUlUpTNLInformationPrst()
{
    return isUlUpTransportLayerInfoPrst_m;
}

//----------------------------------------------------------------------------
// This is a public API. Refer to PduSesResModifyRespTransferPdu.h
//----------------------------------------------------------------------------
QosFlowPerTNLInformationList *
PduSesResModifyRespTransferPdu::getQosFlowPerTNLInformationListDecodedPtr() const
{
    return qosFlowPerTNLInformationList_mp;
}

QosFlowPerTNLInformationList *
PduSesResModifyRespTransferPdu::getQosFlowPerTNLInformationListEncodePtr()
{
    return &qosFlowPerTNLInformationList_m;
}

QosFlowPerTNLInformationItem* 
PduSesResModifyRespTransferPdu::getQosFlowPerTNLInformationItemDecodedPtr(unsigned int index) const
{
    if (index >= 0 && index < maxnoofMultiConnectivityMinusOne_c) {
        QosFlowPerTNLInformationItem *item_p = NULL;
        QosFlowPerTNLInformationList listPtr = *qosFlowPerTNLInformationList_mp;

        if (listPtr != NULL && index <= maxnoofMultiConnectivityMinusOne_c) {
            unsigned int i = 0;
            for (; i < index && listPtr != NULL; i++) {
                listPtr = listPtr->next;
            }
            if (listPtr != NULL && i == index) {
                item_p = &(listPtr->value);
            }
        }
        return item_p;
    } else {
        return NULL;
    }
}

QosFlowPerTNLInformationItem* 
PduSesResModifyRespTransferPdu::getQosFlowPerTNLInformationItemEncodePtr(unsigned int index)
{
    if (index < maxnoofMultiConnectivityMinusOne_c) {
        isAddqosFlowPerTnlInfoPrst_m = true;
        return &qosFlowPerTNLInformationList_ma[index].value;
    } else {
        return NULL;
    }
}

void 
PduSesResModifyRespTransferPdu::setNumofQosFlowPerTNLInformationListItem(unsigned int num)
{
    numOfQosFlowPerTNLInformationListItem_m = num;
}

unsigned int 
PduSesResModifyRespTransferPdu::getNumofQosFlowPerTNLInformationListItem() const
{
    return numOfQosFlowPerTNLInformationListItem_m;
}

bool 
PduSesResModifyRespTransferPdu::isQosFlowPerTNLInformationListPrst()
{
    return isAddqosFlowPerTnlInfoPrst_m;
}

//----------------------------------------------------------------------------
// This is a public API. Refer to PduSesResModifyRespTransferPdu.h
//----------------------------------------------------------------------------
AssociatedQosFlowItem* 
PduSesResModifyRespTransferPdu::getAddAssQosFlowItemDecodedPtr(QosFlowPerTNLInformationItem * InfoPtr,
                                                                     unsigned int indAss) const
{
    AssociatedQosFlowItem* item_p = NULL;
    AssociatedQosFlowList assQosFlowList_p =NULL;
    if(InfoPtr != NULL) {
        assQosFlowList_p = InfoPtr->qosFlowPerTNLInformation.associatedQosFlowList;
        if(assQosFlowList_p != NULL && (indAss >= 0 && indAss <maxnoofQosFlows)) {
            unsigned int i = 0;
            for (; i < indAss && assQosFlowList_p != NULL; i++) {
                assQosFlowList_p = assQosFlowList_p->next;
            }
            if (assQosFlowList_p != NULL && i == indAss)
            {
                item_p = &(assQosFlowList_p->value);
            }
        }
    }
    return item_p;
}

AssociatedQosFlowItem* 
PduSesResModifyRespTransferPdu::getAddAssQosFlowItemEncodePtr(unsigned int indInfo,  
                                                                    unsigned int index)
{
    if ((indInfo >= 0 && indInfo < maxnoofMultiConnectivityMinusOne) &&
         (index >= 0 && index < maxnoofQosFlows)) 
    {
        return &(addAssQosFlowList_ma[indInfo][index].value);
    } else {
        return NULL;
    }
}
                                                                    
void 
PduSesResModifyRespTransferPdu::setNumofAddAssQosFlowListItem(unsigned int num,unsigned int index)
{
    if (index >= 0 && index < maxnoofQosFlows) 
    {
        numofAddQosFlowInd_m[index] = num;
    }
}

unsigned int 
PduSesResModifyRespTransferPdu::getNumofAddAssQosFlowListItem(unsigned int index) const
{
    if (index >= 0 && index < maxnoofQosFlows)
    {
        return numofAddQosFlowInd_m[index];
    }
    else
    {
        return 0;
    }
}

//----------------------------------------------------------------------------
// This is a public API. Refer to PduSesResModifyRespTransferPdu.h
//----------------------------------------------------------------------------
QosFlowAddOrModifyResponseList* 
PduSesResModifyRespTransferPdu::getQosFlowAddOrModRespListDecodedPtr() const
{
    return qosFlowAddOrModifyRespList_mp;
}
QosFlowAddOrModifyResponseList* 
PduSesResModifyRespTransferPdu::getQosFlowAddOrModRespListEncodePtr()
{
    return &qosFlowAddOrModifyRespList_m;
}

QosFlowAddOrModifyResponseItem* 
PduSesResModifyRespTransferPdu::getQosFlowAddOrModRespItemDecodedPtr(unsigned int index) const
{
    if (index >= 0 && index < maxnoofQosFlows)
    {
        QosFlowAddOrModifyResponseItem *item_p = NULL;
        QosFlowAddOrModifyResponseList listPtr = *qosFlowAddOrModifyRespList_mp;
        
        if (listPtr != NULL && index <= maxnoofQosFlows)    {
            unsigned int i = 0;
            for (; i < index && listPtr != NULL; i++){
                listPtr = listPtr->next;
            }
            if (listPtr != NULL && i == index){
                item_p = &(listPtr->value);
            }
        }
        return item_p;
    }
    else{
        return NULL;
    }
}

QosFlowAddOrModifyResponseItem* 
PduSesResModifyRespTransferPdu::getQosFlowAddOrModRespItemEncodePtr(unsigned int index)
{
    if (index < maxnoofQosFlows){
        isQosFlowAddOrModifyResponsePrst_m = true;
        return &qosFlowAddOrModifyRespList_ma[index].value;
    }
    else{
        return NULL;
    }
}
void 
PduSesResModifyRespTransferPdu::setNumofQosFlowAddOrModRespListItem(unsigned int num)
{
    numofQosFlowAddOrModifyResp_m = num;
}
unsigned int 
PduSesResModifyRespTransferPdu::getNumofQosFlowAddOrModRespListItem() const
{
    return numofQosFlowAddOrModifyResp_m;
}
bool 
PduSesResModifyRespTransferPdu::isQosFlowAddOrModifyRespListPrst()
{
    return isQosFlowAddOrModifyResponsePrst_m;
}

//----------------------------------------------------------------------------
// This is a public API. Refer to PduSesResModifyRespTransferPdu.h
//----------------------------------------------------------------------------
QosFlowListWithCause* 
PduSesResModifyRespTransferPdu::getQosFlowToRelListDecodedPtr() const
{
    return qosFlowToReleaseList_mp;
}
QosFlowListWithCause* 
PduSesResModifyRespTransferPdu::getQosFlowToRelListEncodePtr()
{
    return &qosFlowToReleaseList_m;
}

QosFlowWithCauseItem* 
PduSesResModifyRespTransferPdu::getQosFlowToRelItemDecodedPtr(unsigned int index) const
{
    if (index >= 0 && index < maxnoofQosFlows)
    {
        QosFlowWithCauseItem *item_p = NULL;
        QosFlowListWithCause listPtr = *qosFlowToReleaseList_mp;
        
        if (listPtr != NULL && index <= maxnoofQosFlows)    {
            unsigned int i = 0;
            for (; i < index && listPtr != NULL; i++){
                listPtr = listPtr->next;
            }
            if (listPtr != NULL && i == index){
                item_p = &(listPtr->value);
            }
        }
        return item_p;
    }
    else{
        return NULL;
    }
}

QosFlowWithCauseItem* 
PduSesResModifyRespTransferPdu::getQosFlowToRelItemEncodePtr(unsigned int index)
{
    if (index < maxnoofQosFlows){
        isQosFlowListPrst_m = true;
        return &qosFlowToReleaseList_ma[index].value;
    }
    else{
        return NULL;
    }
}
void 
PduSesResModifyRespTransferPdu::setNumofQosFlowToRelListItem(unsigned int num)
{
    numofQosFlowToRelease_m = num;
}
unsigned int 
PduSesResModifyRespTransferPdu::getNumofQosFlowToRelListItem() const
{
    return numofQosFlowToRelease_m;
}
bool 
PduSesResModifyRespTransferPdu::isQosFlowToRelListPrst()
{
    return isQosFlowListPrst_m;
}


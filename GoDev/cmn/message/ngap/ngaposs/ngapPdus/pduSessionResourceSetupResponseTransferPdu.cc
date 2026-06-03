#include <string.h>
#include <iostream>

#include "pduSessionResourceSetupResponseTransferPdu.h"

PduSesResSetupRespTransferPdu::PduSesResSetupRespTransferPdu()
{
    reset_v();
}

PduSesResSetupRespTransferPdu::~PduSesResSetupRespTransferPdu()
{
}

void PduSesResSetupRespTransferPdu::reset_v()
{
    pduSessResSetupRespTransfer_mp = NULL;

    // qos flow per tunnel info
    qosFlowPerTnlInfo_mp = NULL;
    qosFlowPerTnlInfo_m.bit_mask = 0;
    qosFlowPerTnlInfo_m.iE_Extensions = NULL;
    qosFlowPerTnlInfo_m.uPTransportLayerInformation.choice = gTPTunnel_chosen;
    qosFlowPerTnlInfo_m.uPTransportLayerInformation.u.gTPTunnel.bit_mask = 0;
    qosFlowPerTnlInfo_m.uPTransportLayerInformation.u.gTPTunnel.iE_Extensions = NULL;
    qosFlowPerTnlInfo_m.uPTransportLayerInformation.u.gTPTunnel.transportLayerAddress.length = 0;
    qosFlowPerTnlInfo_m.uPTransportLayerInformation.u.gTPTunnel.transportLayerAddress.value = transLayerAddress_m.dlTransLayerAddr;
    qosFlowPerTnlInfo_m.uPTransportLayerInformation.u.gTPTunnel.gTP_TEID.length = 0;
    memset(qosFlowPerTnlInfo_m.uPTransportLayerInformation.u.gTPTunnel.gTP_TEID.value, 
        0, 
        sizeof(qosFlowPerTnlInfo_m.uPTransportLayerInformation.u.gTPTunnel.gTP_TEID.value));
    qosFlowPerTnlInfo_m.associatedQosFlowList = NULL;

    numofQosFlowInd_m = 0;
    for (unsigned int i = 0; i < maxnoofQosFlows; i++)
    {
        assQosFlowList_ma[i].value.bit_mask = 0;
        assQosFlowList_ma[i].value.iE_Extensions = NULL;
        assQosFlowList_ma[i].value.qosFlowIdentifier = 0;
        assQosFlowList_ma[i].value.qosFlowMappingIndication = ul;
    }

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


    // security result
    secResult_mp = NULL;
    secResult_m.bit_mask = 0;
    secResult_m.iE_Extensions = NULL;
    secResult_m.confidentialityProtectionResult = ConfidentialityProtectionResult_not_performed;
    secResult_m.integrityProtectionResult = IntegrityProtectionResult_not_performed;
    isSecResultPrst_m = false;
    
    //qos flow failed to setup list
    qosFlowFailedSetupList_mp = NULL;
    qosFlowFailedSetupList_m = NULL;
    isQosFlowFailedListPrst_m = false;
    numofQosFlowlist_m = 0;
    for (unsigned int i = 0; i < maxnoofQosFlows; i++)
    {
        qosFlowList_ma[i].value.bit_mask = 0;
        qosFlowList_ma[i].value.iE_Extensions = NULL;
        
        qosFlowList_ma[i].value.qosFlowIdentifier = 0;
        qosFlowList_ma[i].value.cause.choice = misc_chosen; 
        qosFlowList_ma[i].value.cause.u.misc = CauseMisc_unspecified;
    }
}

unsigned int 
PduSesResSetupRespTransferPdu::decodePerToOss(NgapOssCtxt* ctxt)
{
   if(NULL == ctxt)
   {
       return -1;
   }
   
   OssBuf inputBuf;
   OssBuf outputBuf;
   int pduNum = PDUSessionResourceSetupResponseTransfer_PDU;

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
PduSesResSetupRespTransferPdu::encodeOssToPer(NgapOssCtxt *ctxt)
{
    if (NULL == ctxt){
        return -1;
    }

    PDUSessionResourceSetupResponseTransfer ngApPdu;
    ngApPdu.iE_Extensions = NULL;
    ngApPdu.bit_mask = 0;

    //mandatory qos flow per tul information
    
    //initial ngApPdu.qosFlowPerTNLInformation
    ngApPdu.dLQosFlowPerTNLInformation.bit_mask = 0;
    ngApPdu.dLQosFlowPerTNLInformation.iE_Extensions = NULL;
    ngApPdu.dLQosFlowPerTNLInformation.associatedQosFlowList = NULL;
    
    memcpy(&ngApPdu.dLQosFlowPerTNLInformation.uPTransportLayerInformation,
        &qosFlowPerTnlInfo_m.uPTransportLayerInformation, 
        sizeof(qosFlowPerTnlInfo_m.uPTransportLayerInformation));

    struct AssociatedQosFlowList_ * nexQosFlow_p = NULL;
    for(int i=0; i< numofQosFlowInd_m; i++){
        assQosFlowList_ma[i].next = nexQosFlow_p;
        nexQosFlow_p = &assQosFlowList_ma[i];
    }
    ngApPdu.dLQosFlowPerTNLInformation.associatedQosFlowList = nexQosFlow_p;

    // additional qos flow per tul information list
    if(isAddqosFlowPerTnlInfoPrst_m){
        ngApPdu.bit_mask |= PDUSessionResourceSetupResponseTransfer_additionalDLQosFlowPerTNLInformation_present;
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
     
    if(isSecResultPrst_m){
        ngApPdu.bit_mask |= PDUSessionResourceSetupResponseTransfer_securityResult_present;
        memcpy(&ngApPdu.securityResult, &secResult_m, sizeof(secResult_m));
    }

    if(isQosFlowFailedListPrst_m){
        ngApPdu.bit_mask |= PDUSessionResourceSetupResponseTransfer_qosFlowFailedToSetupList_present;

        if(numofQosFlowlist_m > maxnoofQosFlows || numofQosFlowlist_m <=0){
            return defaultFailureCode_c;
        }

        struct QosFlowListWithCause_ *nextQosFlowList_p = NULL;
        for(int i=0; i<numofQosFlowlist_m; i++){
           qosFlowList_ma[i].next = nextQosFlowList_p;
           nextQosFlowList_p = &qosFlowList_ma[i];
        }
        ngApPdu.qosFlowFailedToSetupList = nextQosFlowList_p; 
    }

    //Convert Structures to PER
    OssBuf outputBuf;
    unsigned int result;

    ctxt->reset_v();
    outputBuf = ctxt->getPrimaryCodecBuf();

    result = ctxt->encodeOssToPer(PDUSessionResourceSetupResponseTransfer_PDU, &ngApPdu, &outputBuf);

    if (result == encodeSuccessCode_c){
        perBufRef_m = outputBuf;
    }

    return result;
}

void PduSesResSetupRespTransferPdu::getDecodedOssDataPtrs_v()
{
    reset_v();

    pduSessResSetupRespTransfer_mp = reinterpret_cast<PDUSessionResourceSetupResponseTransfer *>(ossBufRef_m.value);
    if (pduSessResSetupRespTransfer_mp == NULL){
        return;
    }
    qosFlowPerTnlInfo_mp = &pduSessResSetupRespTransfer_mp->dLQosFlowPerTNLInformation;
    struct AssociatedQosFlowList_ * nextPtr = qosFlowPerTnlInfo_mp->associatedQosFlowList;
    while( nextPtr != NULL ){
        nextPtr = nextPtr->next;
        numofQosFlowInd_m ++;           
    }

    //additional qos flow per tul information list
    if(pduSessResSetupRespTransfer_mp ->bit_mask & PDUSessionResourceSetupResponseTransfer_additionalDLQosFlowPerTNLInformation_present) {
        isAddqosFlowPerTnlInfoPrst_m = true;
        qosFlowPerTNLInformationList_mp = &pduSessResSetupRespTransfer_mp->additionalDLQosFlowPerTNLInformation;
        struct QosFlowPerTNLInformationList_ *qosFlowPerTNLInfoPtr = 
                                                           pduSessResSetupRespTransfer_mp->additionalDLQosFlowPerTNLInformation;
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

    if(pduSessResSetupRespTransfer_mp->bit_mask & PDUSessionResourceSetupResponseTransfer_securityResult_present)
    {
        isSecResultPrst_m = true;
        secResult_mp = &pduSessResSetupRespTransfer_mp->securityResult;
    }

    if(pduSessResSetupRespTransfer_mp->bit_mask & PDUSessionResourceSetupResponseTransfer_qosFlowFailedToSetupList_present)
    {
        isQosFlowFailedListPrst_m = true;
        qosFlowFailedSetupList_mp = &pduSessResSetupRespTransfer_mp->qosFlowFailedToSetupList;
        struct QosFlowListWithCause_ * nextPtr = pduSessResSetupRespTransfer_mp->qosFlowFailedToSetupList;
        while (nextPtr != NULL)
        {
           nextPtr = nextPtr->next;
           numofQosFlowlist_m++;
        }
    }
}

PDUSessionResourceSetupResponseTransfer *
PduSesResSetupRespTransferPdu::getPduSessResSetupRespTransferDecodedPtr() const
{
    return pduSessResSetupRespTransfer_mp;
}

QosFlowPerTNLInformation *
PduSesResSetupRespTransferPdu::getQosFlowInfoDecodedPtr() const
{
    return qosFlowPerTnlInfo_mp;
}

QosFlowPerTNLInformation *
PduSesResSetupRespTransferPdu::getQosFlowInfoEncodePtr()
{
    return &qosFlowPerTnlInfo_m;
}
AssociatedQosFlowItem* 
PduSesResSetupRespTransferPdu::getAssQosFlowItemDecodedPtr(unsigned int index) const
{
    if (index >= 0 && index < maxnoofPDUSessions)
    {
        AssociatedQosFlowItem *item_p = NULL;
        struct AssociatedQosFlowList_ *listPtr = 
            qosFlowPerTnlInfo_mp->associatedQosFlowList;
        
        if ((listPtr != NULL) && (index <= maxnoofPDUSessions)) {
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

AssociatedQosFlowItem* 
PduSesResSetupRespTransferPdu::getAssQosFlowItemEncodePtr(unsigned int index)
{
    if (index < maxnoofPDUSessions){
        return &assQosFlowList_ma[index].value;
    }
    else{
        return NULL;
    }
}

void 
PduSesResSetupRespTransferPdu::setNumofAssQosFlowListItem_m_v(unsigned int num)
{
    numofQosFlowInd_m = num;
}

unsigned int 
PduSesResSetupRespTransferPdu::getNumofAssQosFlowListItem() const
{
    return numofQosFlowInd_m;
}

QosFlowPerTNLInformationList *
PduSesResSetupRespTransferPdu::getQosFlowPerTNLInformationListDecodedPtr() const
{
    return qosFlowPerTNLInformationList_mp;
}

QosFlowPerTNLInformationList *
PduSesResSetupRespTransferPdu::getQosFlowPerTNLInformationListEncodePtr()
{
    return &qosFlowPerTNLInformationList_m;
}

bool 
PduSesResSetupRespTransferPdu::isQosFlowPerTNLInformationListPrst()
{
    return isAddqosFlowPerTnlInfoPrst_m;
}

QosFlowPerTNLInformationItem*
PduSesResSetupRespTransferPdu::getQosFlowPerTNLInformationListItemDecodedPtr(unsigned int index) const
{
    QosFlowPerTNLInformationItem *item_p = NULL;
    QosFlowPerTNLInformationList qosFlowPerTNLInformation_p = NULL;
    if(qosFlowPerTNLInformationList_mp != NULL)
    {
        qosFlowPerTNLInformation_p = *(qosFlowPerTNLInformationList_mp);
    }

    if(qosFlowPerTNLInformation_p != NULL
        && (index >= 0 && index < numOfQosFlowPerTNLInformationListItem_m))
    {
        unsigned int i = 0;
        for(; i < index && qosFlowPerTNLInformation_p != NULL; i++)
        {
            qosFlowPerTNLInformation_p = qosFlowPerTNLInformation_p->next;
        }
        if(qosFlowPerTNLInformation_p != NULL && i == index)
        {
            item_p = &(qosFlowPerTNLInformation_p->value);
        }
    }

    return item_p;
}


QosFlowPerTNLInformationItem*
PduSesResSetupRespTransferPdu::getQosFlowPerTNLInformationListItemEncodePtr(unsigned int index)
{
    if(index >= 0 && index < maxnoofQosFlows)
    {
        isAddqosFlowPerTnlInfoPrst_m = true;
        return &(qosFlowPerTNLInformationList_ma[index].value);
    }
    else
    {
        return NULL;
    }
}

void 
PduSesResSetupRespTransferPdu::setNumOfQosFlowPerTNLInformationListItem_v(unsigned int num)
{
    numOfQosFlowPerTNLInformationListItem_m = num;
}

unsigned int 
PduSesResSetupRespTransferPdu::getNumOfQosFlowPerTNLInformationListItem() const
{
    return numOfQosFlowPerTNLInformationListItem_m;
}

AssociatedQosFlowItem* 
PduSesResSetupRespTransferPdu::getAddAssQosFlowItemDecodedPtr(QosFlowPerTNLInformationItem *addQosFlowInfoPtr ,
                                                                     unsigned int indAssQosFlow) const
{
    AssociatedQosFlowItem* item_p = NULL;
    AssociatedQosFlowList assQosFlow_p = NULL;
    if(addQosFlowInfoPtr != NULL)
    {
        assQosFlow_p = addQosFlowInfoPtr->qosFlowPerTNLInformation.associatedQosFlowList;
        if(assQosFlow_p != NULL && (indAssQosFlow >= 0 && indAssQosFlow < maxnoofQosFlows))
        {
            unsigned int i = 0;
            for(;i < indAssQosFlow && assQosFlow_p != NULL; i++)
            {
                assQosFlow_p = assQosFlow_p->next;
            }
            if(assQosFlow_p != NULL && i == indAssQosFlow)
            {
                item_p = &(assQosFlow_p->value);
            }
        }
    }
    return item_p;
}

AssociatedQosFlowItem* 
PduSesResSetupRespTransferPdu::getAddAssQosFlowItemEncodePtr(unsigned int index,
                                                                    unsigned int indAssQosFlow)
{
    if((index >= 0 && index < maxnoofQosFlows) &&
        (indAssQosFlow >= 0 && indAssQosFlow < maxnoofQosFlows))
    {
        return &(addAssQosFlowList_ma[index][indAssQosFlow].value);
    }
    else
    {
        return NULL;
    }
}

void 
PduSesResSetupRespTransferPdu::setNumofAddAssQosFlowListItem_m_v(unsigned int num,unsigned int index)
{
    if(index >= 0 && index < maxnoofQosFlows)
    {
        numofAddQosFlowInd_m[index] = num;
    }
}

unsigned int 
PduSesResSetupRespTransferPdu::getNumofAddAssQosFlowListItem(unsigned int index) const
{
    if(index >= 0 && index < maxnoofQosFlows)
    {
        return numofAddQosFlowInd_m[index];
    }
    else
    {
        return 0;
    }
}

SecurityResult *
PduSesResSetupRespTransferPdu::getSecurityResultDecodedPtr() const
{
    return secResult_mp;
}

SecurityResult *
PduSesResSetupRespTransferPdu::getSecurityResultEncodePtr()
{
    isSecResultPrst_m = true;
    return &secResult_m;
}

bool 
PduSesResSetupRespTransferPdu::isSecResultPrst()
{
    return isSecResultPrst_m;
}

QosFlowListWithCause* 
PduSesResSetupRespTransferPdu::getQosFlowFailedSetupListDecodedPtr() const
{
    return qosFlowFailedSetupList_mp;
}
QosFlowListWithCause* 
PduSesResSetupRespTransferPdu::getQosFlowFailedSetupListEncodePtr()
{
    return &qosFlowFailedSetupList_m;
}

QosFlowWithCauseItem* 
PduSesResSetupRespTransferPdu::getQosFlowItemDecodedPtr(unsigned int index) const
{
    if (index >= 0 && index < maxnoofQosFlows)
    {
        QosFlowWithCauseItem *item_p = NULL;
        QosFlowListWithCause listPtr = *qosFlowFailedSetupList_mp;
        
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
PduSesResSetupRespTransferPdu::getQosFlowItemEncodePtr(unsigned int index)
{
    if (index < maxnoofQosFlows){
        isQosFlowFailedListPrst_m = true;
        return &qosFlowList_ma[index].value;
    }
    else{
        return NULL;
    }
}

void 
PduSesResSetupRespTransferPdu::setNumofQosFlowListItem_m_v(unsigned int num)
{
    numofQosFlowlist_m = num;
}

unsigned int 
PduSesResSetupRespTransferPdu::getNumofQosFlowListItem() const
{
    return numofQosFlowlist_m;
}

bool 
PduSesResSetupRespTransferPdu::isQosFlowFailedSetupListPrst()
{
    return isQosFlowFailedListPrst_m;
}

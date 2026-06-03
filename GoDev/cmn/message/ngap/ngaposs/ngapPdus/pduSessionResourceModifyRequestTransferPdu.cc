#include <string.h>
#include <iostream>

#include "pduSessionResourceModifyRequestTransferPdu.h"

PduSesResModifyReqTransferPdu::PduSesResModifyReqTransferPdu()
{
    reset_v();
}

PduSesResModifyReqTransferPdu::~PduSesResModifyReqTransferPdu()
{
}

void PduSesResModifyReqTransferPdu::reset_v()
{
    pduSessResModifyReqTransfer_mp = NULL;

    // session ambr
    sessAMBR_mp = NULL;
    sessAMBR_m.bit_mask = 0;
    sessAMBR_m.iE_Extensions = NULL;
    sessAMBR_m.pDUSessionAggregateMaximumBitRateDL = 0;
    sessAMBR_m.pDUSessionAggregateMaximumBitRateUL = 0;
    isSessAmbrPrst_m = false;

    // network Instance
    networkInstance_mp = NULL;
    networkInstance_m = 0;
    isNetworkInstancePrst_m = false;
    
    // ul ngu up tnl modify list
    ulNguUpTnlModifyList_mp = NULL;
    ulNguUpTnlModifyList_m = NULL;
    numofUlNguUpTnlModifyList_m = 0;
    for (unsigned int i = 0; i < maxnoofMultiConnectivity; i++)
    {
       ulNguUpTnlModifyList_ma[i].value.bit_mask = 0;
       ulNguUpTnlModifyList_ma[i].value.iE_Extensions = NULL;

       memset(&(ulTransLayerAddress_m[i]), 0, sizeof(TransLayerAddr));
       
       ulNguUpTnlModifyList_ma[i].value.uL_NGU_UP_TNLInformation.choice = gTPTunnel_chosen;
       ulNguUpTnlModifyList_ma[i].value.uL_NGU_UP_TNLInformation.u.gTPTunnel.bit_mask = 0;
       ulNguUpTnlModifyList_ma[i].value.uL_NGU_UP_TNLInformation.u.gTPTunnel.iE_Extensions = NULL;
       ulNguUpTnlModifyList_ma[i].value.uL_NGU_UP_TNLInformation.u.gTPTunnel.transportLayerAddress.length = 0;
       ulNguUpTnlModifyList_ma[i].value.uL_NGU_UP_TNLInformation.u.gTPTunnel.transportLayerAddress.value = ulTransLayerAddress_m[i].transLayerAddr;
       ulNguUpTnlModifyList_ma[i].value.uL_NGU_UP_TNLInformation.u.gTPTunnel.gTP_TEID.length = 0;
       memset(ulNguUpTnlModifyList_ma[i].value.uL_NGU_UP_TNLInformation.u.gTPTunnel.gTP_TEID.value, 
              0, 
              sizeof(ulNguUpTnlModifyList_ma[i].value.uL_NGU_UP_TNLInformation.u.gTPTunnel.gTP_TEID.value));

       memset(&dlTransLayerAddress_m[i], 0, sizeof(TransLayerAddr));
       ulNguUpTnlModifyList_ma[i].value.dL_NGU_UP_TNLInformation.choice = gTPTunnel_chosen;
       ulNguUpTnlModifyList_ma[i].value.dL_NGU_UP_TNLInformation.u.gTPTunnel.bit_mask = 0;
       ulNguUpTnlModifyList_ma[i].value.dL_NGU_UP_TNLInformation.u.gTPTunnel.iE_Extensions = NULL;
       ulNguUpTnlModifyList_ma[i].value.dL_NGU_UP_TNLInformation.u.gTPTunnel.transportLayerAddress.length = 0;
       ulNguUpTnlModifyList_ma[i].value.dL_NGU_UP_TNLInformation.u.gTPTunnel.transportLayerAddress.value = dlTransLayerAddress_m[i].transLayerAddr;
       ulNguUpTnlModifyList_ma[i].value.dL_NGU_UP_TNLInformation.u.gTPTunnel.gTP_TEID.length = 0;
       memset(ulNguUpTnlModifyList_ma[i].value.dL_NGU_UP_TNLInformation.u.gTPTunnel.gTP_TEID.value, 
              0, 
              sizeof(ulNguUpTnlModifyList_ma[i].value.dL_NGU_UP_TNLInformation.u.gTPTunnel.gTP_TEID.value));
    }

    //add up transport layer information list
    isAddUpTransportLayerInfoListPrst_m = false;
    addUpTransportLayerInfoList_mp = NULL;
    addUpTransportLayerInfoList_m = NULL;
    numofAddUpTransportLayerInfo_m = 0;
    for (unsigned int i = 0;i <maxnoofMultiConnectivityMinusOne; i++)
    {
        addTransLayerAddrList_ma[i].value.bit_mask = 0;
        addTransLayerAddrList_ma[i].value.iE_Extensions = NULL;

        memset(&(addTransLayerAddress_m[i]), 0, sizeof(TransLayerAddr));

        addTransLayerAddrList_ma[i].value.nGU_UP_TNLInformation.choice = gTPTunnel_chosen;
        addTransLayerAddrList_ma[i].value.nGU_UP_TNLInformation.u.gTPTunnel.bit_mask = 0;
        addTransLayerAddrList_ma[i].value.nGU_UP_TNLInformation.u.gTPTunnel.iE_Extensions = NULL;
        addTransLayerAddrList_ma[i].value.nGU_UP_TNLInformation.u.gTPTunnel.transportLayerAddress.length = 0;
        addTransLayerAddrList_ma[i].value.nGU_UP_TNLInformation.u.gTPTunnel.transportLayerAddress.value = addTransLayerAddress_m[i].transLayerAddr;
        addTransLayerAddrList_ma[i].value.nGU_UP_TNLInformation.u.gTPTunnel.gTP_TEID.length = 0;
        memset(addTransLayerAddrList_ma[i].value.nGU_UP_TNLInformation.u.gTPTunnel.gTP_TEID.value, 0, 
               sizeof(addTransLayerAddrList_ma[i].value.nGU_UP_TNLInformation.u.gTPTunnel.gTP_TEID.value));
    }

    //qos flow add or modify request list
    qosFlowAddOrModifyReqList_mp = NULL;
    qosFlowAddOrModifyReqList_m = NULL;
    numofQosFlowAddOrModifyReq_m = 0;
    for (unsigned int i = 0; i < maxnoofQoSFlows_c; i++)
    {
       qosFlowAddOrModifyReqList_ma[i].value.bit_mask = 0;
       qosFlowAddOrModifyReqList_ma[i].value.iE_Extensions = NULL;
                
       qosFlowAddOrModifyReqList_ma[i].value.qosFlowIdentifier = 0;
       
       qosFlowAddOrModifyReqList_ma[i].value.qosFlowLevelQosParameters.bit_mask = 0;
       qosFlowAddOrModifyReqList_ma[i].value.qosFlowLevelQosParameters.iE_Extensions = NULL;
       qosFlowAddOrModifyReqList_ma[i].value.qosFlowLevelQosParameters.qosCharacteristics.choice = nonDynamic5QI_chosen;
       qosFlowAddOrModifyReqList_ma[i].value.qosFlowLevelQosParameters.qosCharacteristics.u.nonDynamic5QI.bit_mask = 0;
       qosFlowAddOrModifyReqList_ma[i].value.qosFlowLevelQosParameters.qosCharacteristics.u.nonDynamic5QI.iE_Extensions = NULL;
       qosFlowAddOrModifyReqList_ma[i].value.qosFlowLevelQosParameters.qosCharacteristics.u.nonDynamic5QI.fiveQI = 0;
       qosFlowAddOrModifyReqList_ma[i].value.qosFlowLevelQosParameters.qosCharacteristics.u.nonDynamic5QI.priorityLevelQos = 0;
       qosFlowAddOrModifyReqList_ma[i].value.qosFlowLevelQosParameters.qosCharacteristics.u.nonDynamic5QI.averagingWindow = 0;
       qosFlowAddOrModifyReqList_ma[i].value.qosFlowLevelQosParameters.qosCharacteristics.u.nonDynamic5QI.maximumDataBurstVolume = 0;
                
       qosFlowAddOrModifyReqList_ma[i].value.qosFlowLevelQosParameters.allocationAndRetentionPriority.bit_mask = 0;
       qosFlowAddOrModifyReqList_ma[i].value.qosFlowLevelQosParameters.allocationAndRetentionPriority.iE_Extensions = NULL;
       qosFlowAddOrModifyReqList_ma[i].value.qosFlowLevelQosParameters.allocationAndRetentionPriority.priorityLevelARP = 0;
       qosFlowAddOrModifyReqList_ma[i].value.qosFlowLevelQosParameters.allocationAndRetentionPriority.pre_emptionCapability = shall_not_trigger_pre_emption;
       qosFlowAddOrModifyReqList_ma[i].value.qosFlowLevelQosParameters.allocationAndRetentionPriority.pre_emptionVulnerability = not_pre_emptable;

       qosFlowAddOrModifyReqList_ma[i].value.qosFlowLevelQosParameters.gBR_QosInformation.bit_mask = 0;
       qosFlowAddOrModifyReqList_ma[i].value.qosFlowLevelQosParameters.gBR_QosInformation.iE_Extensions = NULL;
       qosFlowAddOrModifyReqList_ma[i].value.qosFlowLevelQosParameters.gBR_QosInformation.maximumFlowBitRateDL = 0;
       qosFlowAddOrModifyReqList_ma[i].value.qosFlowLevelQosParameters.gBR_QosInformation.maximumFlowBitRateUL = 0;
       qosFlowAddOrModifyReqList_ma[i].value.qosFlowLevelQosParameters.gBR_QosInformation.guaranteedFlowBitRateDL = 0;
       qosFlowAddOrModifyReqList_ma[i].value.qosFlowLevelQosParameters.gBR_QosInformation.guaranteedFlowBitRateUL = 0;
       qosFlowAddOrModifyReqList_ma[i].value.qosFlowLevelQosParameters.gBR_QosInformation.notificationControl = notification_requested;
       qosFlowAddOrModifyReqList_ma[i].value.qosFlowLevelQosParameters.gBR_QosInformation.maximumPacketLossRateDL = 0;
       qosFlowAddOrModifyReqList_ma[i].value.qosFlowLevelQosParameters.gBR_QosInformation.maximumPacketLossRateUL = 0;

       qosFlowAddOrModifyReqList_ma[i].value.qosFlowLevelQosParameters.reflectiveQosAttribute = subject_to;
       qosFlowAddOrModifyReqList_ma[i].value.qosFlowLevelQosParameters.additionalQosFlowInformation = more_likely;
                
       qosFlowAddOrModifyReqList_ma[i].value.e_RAB_ID = 0;
       qosFlowAddOrModifyReqList_ma[i].next = NULL;
    }

    //qos flow to release list
    qosFlowToReleaseList_mp = NULL;
    qosFlowToReleaseList_m = NULL;
    numofQosFlowToRelease_m = 0;
    for (unsigned int i = 0; i < maxnoofQoSFlows_c; i++)
    {
        qosFlowToReleaseList_ma[i].value.bit_mask = 0;
        qosFlowToReleaseList_ma[i].value.iE_Extensions = NULL;
                
        qosFlowToReleaseList_ma[i].value.qosFlowIdentifier = 0;
        qosFlowToReleaseList_ma[i].value.cause.choice = radioNetwork_chosen;
        qosFlowToReleaseList_ma[i].value.cause.u.radioNetwork = CauseRadioNetwork_unspecified;

        qosFlowToReleaseList_ma[i].next = NULL;
    }

    //Common Network Instance
    commonNetworkInstance_m.length = 0;
    memset(cmmnNtwrkInstance_ma.data,0xff,sizeof(cmmnNtwrkInstance_ma.data));
    commonNetworkInstance_m.value = cmmnNtwrkInstance_ma.data;
    commonNetworkInstance_mp = NULL;
    commonntkInstancePresent_m = false;
}

unsigned int 
PduSesResModifyReqTransferPdu::decodePerToOss(NgapOssCtxt* ctxt)
{
   if(NULL == ctxt)
   {
       return -1;
   }
   
   OssBuf inputBuf;
   OssBuf outputBuf;
   int pduNum = PDUSessionResourceModifyRequestTransfer_PDU;

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
PduSesResModifyReqTransferPdu::encodeOssToPer(NgapOssCtxt *ctxt)
{
    if (NULL == ctxt)
    {
        return -1;
    }

    ProtocolIeType *nextIe_p = NULL;

    //Common Network Instance
    ProtocolIeType commonNetworkInstaceIe;
    if(commonntkInstancePresent_m){
        commonNetworkInstaceIe.next =nextIe_p;
        commonNetworkInstaceIe.value.id = id_CommonNetworkInstance;
        commonNetworkInstaceIe.value.criticality = ignore;
        commonNetworkInstaceIe.value.value.pduNum = PDU_PDUSessionResourceModifyRequestTransferIEs_Value_CommonNetworkInstance;
        commonNetworkInstaceIe.value.value.decoded.pdu_CommonNetworkInstance = &commonNetworkInstance_m;
        commonNetworkInstaceIe.value.value.encoded.length = 0;
        commonNetworkInstaceIe.value.value.encoded.value = NULL;
        nextIe_p = &commonNetworkInstaceIe;
    }

    //qos flow to release list
    ProtocolIeType qosFlowToRelListIe;
    if(numofQosFlowToRelease_m < maxnoofQoSFlows_c && numofQosFlowToRelease_m > 0)
    {
        struct QosFlowListWithCause_  *nextQosFlowToRelList_p = NULL;
        for(int i = 0; i < numofQosFlowToRelease_m; i++)
        {
            qosFlowToReleaseList_ma[i].next = nextQosFlowToRelList_p;
            nextQosFlowToRelList_p = &qosFlowToReleaseList_ma[i];       
        }
        qosFlowToReleaseList_m = nextQosFlowToRelList_p;
        qosFlowToRelListIe.next =nextIe_p;
        qosFlowToRelListIe.value.id = id_QosFlowToReleaseList;
        qosFlowToRelListIe.value.criticality = reject;
        qosFlowToRelListIe.value.value.pduNum = PDU_PDUSessionResourceModifyRequestTransferIEs_Value_QosFlowListWithCause;
        qosFlowToRelListIe.value.value.decoded.pdu_QosFlowListWithCause = &qosFlowToReleaseList_m;
        qosFlowToRelListIe.value.value.encoded.length = 0;
        qosFlowToRelListIe.value.value.encoded.value = NULL;
        nextIe_p = &qosFlowToRelListIe;
    }

    //qos flow add or mod request list
    ProtocolIeType qosFlowAddOrModReqListIe;
    if(numofQosFlowAddOrModifyReq_m < maxnoofQoSFlows_c && numofQosFlowAddOrModifyReq_m > 0)
    {
        struct QosFlowAddOrModifyRequestList_ *nextQosFlowAddOrModifyRequestList_p = NULL;
        for(int i = 0; i < numofQosFlowAddOrModifyReq_m; i++)
        {
            qosFlowAddOrModifyReqList_ma[i].next = nextQosFlowAddOrModifyRequestList_p;
            nextQosFlowAddOrModifyRequestList_p = &qosFlowAddOrModifyReqList_ma[i];       
        }
        qosFlowAddOrModifyReqList_m = nextQosFlowAddOrModifyRequestList_p;

        qosFlowAddOrModReqListIe.next = nextIe_p;
        qosFlowAddOrModReqListIe.value.id = id_QosFlowAddOrModifyRequestList;
        qosFlowAddOrModReqListIe.value.criticality = reject;
        qosFlowAddOrModReqListIe.value.value.pduNum = PDU_PDUSessionResourceModifyRequestTransferIEs_Value_QosFlowAddOrModifyRequestList;
        qosFlowAddOrModReqListIe.value.value.decoded.pdu_QosFlowAddOrModifyRequestList = &qosFlowAddOrModifyReqList_m;
        qosFlowAddOrModReqListIe.value.value.encoded.length = 0;
        qosFlowAddOrModReqListIe.value.value.encoded.value = NULL;
        nextIe_p = &qosFlowAddOrModReqListIe;
    }
    
    //network instance
    ProtocolIeType networkInstanceIe;
    if(isNetworkInstancePrst_m){
        networkInstanceIe.next =nextIe_p;
        networkInstanceIe.value.id = id_NetworkInstance;
        networkInstanceIe.value.criticality = reject;
        networkInstanceIe.value.value.pduNum = PDU_PDUSessionResourceModifyRequestTransferIEs_Value_NetworkInstance;
        networkInstanceIe.value.value.decoded.pdu_NetworkInstance = &networkInstance_m;
        networkInstanceIe.value.value.encoded.length = 0;
        networkInstanceIe.value.value.encoded.value = NULL;
        nextIe_p = &networkInstanceIe;
    }

    //Up Transport Layer Information List
    ProtocolIeType addUpTransportLayerInfoListIe;
    if(numofAddUpTransportLayerInfo_m < maxnoofMultiConnectivityMinusOne && numofAddUpTransportLayerInfo_m > 0) 
    {
        struct UPTransportLayerInformationList_ *nextAddUpTransportLayerInfoList_p = NULL;
        for (int i = 0; i <numofAddUpTransportLayerInfo_m; i++)
        {
            addTransLayerAddrList_ma[i].next = nextAddUpTransportLayerInfoList_p;
            nextAddUpTransportLayerInfoList_p = &addTransLayerAddrList_ma[i];
        }
        addUpTransportLayerInfoList_m = nextAddUpTransportLayerInfoList_p;

        addUpTransportLayerInfoListIe.next = nextIe_p;
        addUpTransportLayerInfoListIe.value.id = id_AdditionalUL_NGU_UP_TNLInformation;
        addUpTransportLayerInfoListIe.value.criticality = reject;
        addUpTransportLayerInfoListIe.value.value.pduNum = PDU_PDUSessionResourceModifyRequestTransferIEs_Value_UPTransportLayerInformationList;
        addUpTransportLayerInfoListIe.value.value.decoded.pdu_UPTransportLayerInformationList = &addUpTransportLayerInfoList_m;
        addUpTransportLayerInfoListIe.value.value.encoded.length = 0;
        addUpTransportLayerInfoListIe.value.value.encoded.value = NULL;
        nextIe_p = &addUpTransportLayerInfoListIe;

    }
        
    // UL_NGU_UP_TNLModifyList
    ProtocolIeType upNguUpTnlModifyListIe;
    if(numofUlNguUpTnlModifyList_m < maxnoofMultiConnectivity && numofUlNguUpTnlModifyList_m > 0){
        struct UL_NGU_UP_TNLModifyList_ *nextUpNguUpTnlModifiyList_p = NULL;
        for(int i = 0; i < numofUlNguUpTnlModifyList_m; i++)
        {
            ulNguUpTnlModifyList_ma[i].next = nextUpNguUpTnlModifiyList_p;
            nextUpNguUpTnlModifiyList_p = &ulNguUpTnlModifyList_ma[i];       
        }
        ulNguUpTnlModifyList_m = nextUpNguUpTnlModifiyList_p;
        upNguUpTnlModifyListIe.next =nextIe_p;
        upNguUpTnlModifyListIe.value.id = id_UL_NGU_UP_TNLModifyList;
        upNguUpTnlModifyListIe.value.criticality = reject;
        upNguUpTnlModifyListIe.value.value.pduNum = PDU_PDUSessionResourceModifyRequestTransferIEs_Value_UL_NGU_UP_TNLModifyList;
        upNguUpTnlModifyListIe.value.value.decoded.pdu_UL_NGU_UP_TNLModifyList = &ulNguUpTnlModifyList_m;
        upNguUpTnlModifyListIe.value.value.encoded.length = 0;
        upNguUpTnlModifyListIe.value.value.encoded.value = NULL;
        nextIe_p = &upNguUpTnlModifyListIe;
    }
        
    //Session ambr
    ProtocolIeType sessAmbrIe;
    if(isSessAmbrPrst_m){
        sessAmbrIe.next = nextIe_p;
        sessAmbrIe.value.id = id_PDUSessionAggregateMaximumBitRate;
        sessAmbrIe.value.criticality = reject;
        sessAmbrIe.value.value.pduNum = PDU_PDUSessionResourceModifyRequestTransferIEs_Value_PDUSessionAggregateMaximumBitRate;
        sessAmbrIe.value.value.decoded.pdu_PDUSessionAggregateMaximumBitRate = &sessAMBR_m;
        sessAmbrIe.value.value.encoded.length = 0;
        sessAmbrIe.value.value.encoded.value = NULL;
       nextIe_p = &sessAmbrIe;
    }
        
    //message
    PDUSessionResourceModifyRequestTransfer msg;
    msg.protocolIEs = nextIe_p;

    //Convert Structures to PER
    OssBuf outputBuf;
    unsigned int result;

    ctxt->reset_v();
    outputBuf = ctxt->getPrimaryCodecBuf();

    result = ctxt->encodeOssToPer(PDUSessionResourceModifyRequestTransfer_PDU, &msg, &outputBuf);

    if (result == encodeSuccessCode_c)
    {
        perBufRef_m = outputBuf;
    }

    return result;
}

void PduSesResModifyReqTransferPdu::getDecodedOssDataPtrs_v()
{
    ProtocolIeType *currIe_p;
    reset_v();
    
    pduSessResModifyReqTransfer_mp = reinterpret_cast<PDUSessionResourceModifyRequestTransfer *>(ossBufRef_m.value);

    if (pduSessResModifyReqTransfer_mp == NULL)
    {
        return;
    }

    currIe_p = pduSessResModifyReqTransfer_mp->protocolIEs;

    while (currIe_p != NULL)
    {
        if (currIe_p->value.id == id_PDUSessionAggregateMaximumBitRate
            && currIe_p->value.value.decoded.pdu_PDUSessionAggregateMaximumBitRate != NULL)
        {
            sessAMBR_mp = currIe_p->value.value.decoded.pdu_PDUSessionAggregateMaximumBitRate;
            isSessAmbrPrst_m = true;
        }
        else if (currIe_p->value.id == id_UL_NGU_UP_TNLModifyList
            && currIe_p->value.value.decoded.pdu_UL_NGU_UP_TNLModifyList != NULL)
        {
            ulNguUpTnlModifyList_mp = currIe_p->value.value.decoded.pdu_UL_NGU_UP_TNLModifyList;
            UL_NGU_UP_TNLModifyList curr_p = *ulNguUpTnlModifyList_mp;
            while(curr_p != NULL)
            {
                numofUlNguUpTnlModifyList_m ++;
                curr_p = curr_p->next;    
            }
        }
        else if (currIe_p->value.id == id_AdditionalUL_NGU_UP_TNLInformation
            && currIe_p->value.value.decoded.pdu_UPTransportLayerInformationList != NULL)
        {
            addUpTransportLayerInfoList_mp = currIe_p->value.value.decoded.pdu_UPTransportLayerInformationList;
            UPTransportLayerInformationList curr_p = *addUpTransportLayerInfoList_mp;
            while(curr_p != NULL)
            {
                numofAddUpTransportLayerInfo_m ++;
                curr_p = curr_p->next;
            }
        }
        else if (currIe_p->value.id == id_NetworkInstance
            && currIe_p->value.value.decoded.pdu_NetworkInstance != NULL)
        {
            networkInstance_mp = 
            currIe_p->value.value.decoded.pdu_NetworkInstance;
            isNetworkInstancePrst_m = true;
        }
        else if (currIe_p->value.id == id_QosFlowAddOrModifyRequestList
            && currIe_p->value.value.decoded.pdu_QosFlowAddOrModifyRequestList != NULL)
        {
            qosFlowAddOrModifyReqList_mp = currIe_p->value.value.decoded.pdu_QosFlowAddOrModifyRequestList;
            QosFlowAddOrModifyRequestList curr_p = *qosFlowAddOrModifyReqList_mp;
            while(curr_p != NULL)
            {
                numofQosFlowAddOrModifyReq_m++;
                curr_p = curr_p->next;
            }
        }
        else if (currIe_p->value.id == id_CommonNetworkInstance
            && currIe_p->value.value.decoded.pdu_CommonNetworkInstance != NULL)
        {
            commonNetworkInstance_mp = currIe_p->value.value.decoded.pdu_CommonNetworkInstance;
            commonntkInstancePresent_m = true;
        }
        currIe_p = currIe_p->next;
    }
}

PDUSessionResourceModifyRequestTransfer *
PduSesResModifyReqTransferPdu::getPduSessResModifyReqTransferDecodedPtr() const
{
    return pduSessResModifyReqTransfer_mp;
}

PDUSessionAggregateMaximumBitRate *
PduSesResModifyReqTransferPdu::getSessAmbrDecodedPtr() const
{
    return sessAMBR_mp;
}

PDUSessionAggregateMaximumBitRate *
PduSesResModifyReqTransferPdu::getSessAmbrEncodePtr()
{
    isSessAmbrPrst_m = true; 
    return &sessAMBR_m;
}

bool 
PduSesResModifyReqTransferPdu::isSessAmbrPrst()
{
   return isSessAmbrPrst_m;
}

UL_NGU_UP_TNLModifyList *
PduSesResModifyReqTransferPdu::getUlNguUpTnlModifyListDecodedPtr() const
{
    return ulNguUpTnlModifyList_mp;
}

UL_NGU_UP_TNLModifyList *
PduSesResModifyReqTransferPdu::getUlNguUpTnlModifyListEncodePtr()
{
    return &ulNguUpTnlModifyList_m;
}


UL_NGU_UP_TNLModifyItem* 
PduSesResModifyReqTransferPdu::getUlNguUpTnlModifyItemDecodedPtr(unsigned int index) const
{
    if (index >= 0 && index < maxnoofMultiConnectivity)
    {
       UL_NGU_UP_TNLModifyItem *item_p = NULL;
       UL_NGU_UP_TNLModifyList ulNguUpTnlModifyList_p = *ulNguUpTnlModifyList_mp;
                
       if (ulNguUpTnlModifyList_p != NULL 
           && (index >= 0 && index < maxnoofMultiConnectivity))
       {
           unsigned int i = 0;
           for (; i < index && ulNguUpTnlModifyList_p != NULL; i++)
           {
              ulNguUpTnlModifyList_p = ulNguUpTnlModifyList_p->next;
           }

           if (ulNguUpTnlModifyList_p != NULL && i == index)
           {
               item_p = &(ulNguUpTnlModifyList_p->value);
           }
       }
       return item_p;
    }
    else
    {
       return NULL;
    }
}

UL_NGU_UP_TNLModifyItem* 
PduSesResModifyReqTransferPdu::getUlNguUpTnlModifyItemEncodePtr(unsigned int index)
{
    if (index >= 0 && index < maxnoofMultiConnectivity)
    {
        return &ulNguUpTnlModifyList_ma[index].value;
    }
    else
    {
        return NULL;
    }
}

void 
PduSesResModifyReqTransferPdu::setNumofUlNguUpTnlModifyListItem(unsigned int num)
{
        numofUlNguUpTnlModifyList_m = num;
}

unsigned int 
PduSesResModifyReqTransferPdu::getNumofUlNguUpTnlModifyListItem() const
{
   return numofUlNguUpTnlModifyList_m;
}

NetworkInstance *
PduSesResModifyReqTransferPdu::getNetworkInstanceDecodedPtr() const
{
    return networkInstance_mp;
}

NetworkInstance *
PduSesResModifyReqTransferPdu::getNetworkInstanceEncodePtr()
{
    return &networkInstance_m;
}

bool 
PduSesResModifyReqTransferPdu::isNetworkInstancePrst()
{
    return isNetworkInstancePrst_m;
}

QosFlowAddOrModifyRequestList *
PduSesResModifyReqTransferPdu::getQosFlowAddOrModReqListDecodedPtr() const
{
    return qosFlowAddOrModifyReqList_mp;
}

QosFlowAddOrModifyRequestList *
PduSesResModifyReqTransferPdu::getQosFlowAddOrModReqListEncodePtr()
{
    return &qosFlowAddOrModifyReqList_m;
}

QosFlowAddOrModifyRequestItem* 
PduSesResModifyReqTransferPdu::getQosFlowAddOrModReqItemDecodedPtr(unsigned int index) const
{
    if (index >= 0 && index < maxnoofQoSFlows_c)
    {
       QosFlowAddOrModifyRequestItem *item_p = NULL;
       QosFlowAddOrModifyRequestList qosFlowAddOrModifyReqList_p = *qosFlowAddOrModifyReqList_mp;
                
       if (qosFlowAddOrModifyReqList_p != NULL 
           && (index >= 0 && index < maxnoofQoSFlows_c))
       {
          unsigned int i = 0;
          for (; i < index && qosFlowAddOrModifyReqList_p != NULL; i++)
          {
             qosFlowAddOrModifyReqList_p = qosFlowAddOrModifyReqList_p->next;
          }
          if (qosFlowAddOrModifyReqList_p != NULL && i == index)
          {
             item_p = &(qosFlowAddOrModifyReqList_p->value);
          }
      }
      return item_p;

    }
    else
    {
      return NULL;
    }
}

QosFlowAddOrModifyRequestItem* 
PduSesResModifyReqTransferPdu::getQosFlowAddOrModReqItemEncodePtr(unsigned int index)
{
    if (index >= 0 && index < maxnoofQoSFlows_c)
    {
       return &qosFlowAddOrModifyReqList_ma[index].value;
    }
    else
    {
       return NULL;
    }
}

void 
PduSesResModifyReqTransferPdu::setNumofQosFlowAddOrModReqListItem(unsigned int num)
{
        numofQosFlowAddOrModifyReq_m = num;
}

unsigned int 
PduSesResModifyReqTransferPdu::getNumofQosFlowAddOrModReqListItem() const
{
   return numofQosFlowAddOrModifyReq_m;
}

QosFlowListWithCause *
PduSesResModifyReqTransferPdu::getQosFlowToRelListDecodedPtr() const
{
    return qosFlowToReleaseList_mp;
}

QosFlowListWithCause *
PduSesResModifyReqTransferPdu::getQosFlowToRelListEncodePtr()
{
    return &qosFlowToReleaseList_m;
}

QosFlowWithCauseItem* 
PduSesResModifyReqTransferPdu::getQosFlowToRelItemDecodedPtr(unsigned int index) const
{
    if (index >= 0 && index < maxnoofQoSFlows_c)
    {
       QosFlowWithCauseItem *item_p = NULL;
       QosFlowListWithCause qosFlowToReleaseList_p = *qosFlowToReleaseList_mp;
                
       if (qosFlowToReleaseList_p != NULL 
           && (index >= 0 && index < maxnoofQoSFlows_c))
       {
          unsigned int i = 0;
          for (; i < index && qosFlowToReleaseList_p != NULL; i++)
          {
             qosFlowToReleaseList_p = qosFlowToReleaseList_p->next;
          }
          if (qosFlowToReleaseList_p != NULL && i == index)
          {
             item_p = &(qosFlowToReleaseList_p->value);
          }
      }
      return item_p;

    }
    else
    {
      return NULL;
    }
}

QosFlowWithCauseItem* 
PduSesResModifyReqTransferPdu::getQosFlowToRelItemEncodePtr(unsigned int index)
{
    if (index >= 0 && index < maxnoofQoSFlows_c)
    {
       return &qosFlowToReleaseList_ma[index].value;
    }
    else
    {
       return NULL;
    }
}

void 
PduSesResModifyReqTransferPdu::setNumofQosFlowToRelListItem(unsigned int num)
{
        numofQosFlowToRelease_m = num;
}

unsigned int 
PduSesResModifyReqTransferPdu::getNumofQosFlowToRelListItem() const
{
   return numofQosFlowToRelease_m;
}

CommonNetworkInstance *
PduSesResModifyReqTransferPdu::getCommonNetworkInstanceDecodedPtr() const
{
    return commonNetworkInstance_mp;
}

CommonNetworkInstance *
PduSesResModifyReqTransferPdu::getCommonNetworkInstanceEncodePtr()
{
    commonntkInstancePresent_m = true;
    return &commonNetworkInstance_m;
}

bool
PduSesResModifyReqTransferPdu::isCommonNetworkInstancePrst()
{
    return commonntkInstancePresent_m;
}

UPTransportLayerInformationList *
PduSesResModifyReqTransferPdu::getAddUpTransportLayerInfoListDecodedPtr() const
{
    return addUpTransportLayerInfoList_mp;
}

UPTransportLayerInformationList *
PduSesResModifyReqTransferPdu::getAddUpTransportLayerInfoListEncodePtr()
{
    return &addUpTransportLayerInfoList_m;
}

UPTransportLayerInformationItem *
PduSesResModifyReqTransferPdu::getAddUpTransportLayerInfoItemDecodedPtr(unsigned int index) const
{
    if (index >= 0 && index < maxnoofMultiConnectivityMinusOne)
    {
        UPTransportLayerInformationItem *item_p = NULL;
        UPTransportLayerInformationList addUpTransLayerInfoList_p = *addUpTransportLayerInfoList_mp;
        if(addUpTransLayerInfoList_p != NULL && (index >= 0 && index <maxnoofMultiConnectivityMinusOne)) 
        {
            unsigned int i = 0;
            for (; i <index && addUpTransLayerInfoList_p != NULL; i++)
            {
                addUpTransLayerInfoList_p = addUpTransLayerInfoList_p->next;
            }
            if (addUpTransLayerInfoList_p != NULL && i == index)
            {
                item_p = &(addUpTransLayerInfoList_p->value);
            }
        }
        return item_p;
    }
    else
    {
        return NULL;
    }
}

UPTransportLayerInformationItem *
PduSesResModifyReqTransferPdu::getAddUpTransportLayerInfoItemEncodePtr(unsigned int index)
{
    if (index >= 0 && index <maxnoofMultiConnectivityMinusOne)
    {
        isAddUpTransportLayerInfoListPrst_m = true;
        return &addTransLayerAddrList_ma[index].value;
    }
    else
    {
        return NULL;
    }
}

void 
PduSesResModifyReqTransferPdu::setNumofAddUpTransportLayerInfoListItem(unsigned int num)
{
    numofAddUpTransportLayerInfo_m = num;
}

unsigned int 
PduSesResModifyReqTransferPdu::getNumOfAddUpTransportLayerInfoListItem() const
{
    return numofAddUpTransportLayerInfo_m;
}

bool 
PduSesResModifyReqTransferPdu::isAddUpTransportLayerInfoListPrst()
{
    return isAddUpTransportLayerInfoListPrst_m;
}

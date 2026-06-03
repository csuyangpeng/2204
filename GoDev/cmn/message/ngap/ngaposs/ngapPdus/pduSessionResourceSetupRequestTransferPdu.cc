#include <string.h>
#include <iostream>

#include "pduSessionResourceSetupRequestTransferPdu.h"

PduSesResSetupReqTransferPdu::PduSesResSetupReqTransferPdu()
{
    reset_v();
}

PduSesResSetupReqTransferPdu::~PduSesResSetupReqTransferPdu()
{
}

void PduSesResSetupReqTransferPdu::reset_v()
{
    pduSessResSetupReqTransfer_mp = NULL;

    //Session ambr
    sessAMBR_mp = NULL;
    sessAMBR_m.bit_mask = 0;
    sessAMBR_m.iE_Extensions = NULL;
    sessAMBR_m.pDUSessionAggregateMaximumBitRateDL = 0;
    sessAMBR_m.pDUSessionAggregateMaximumBitRateUL = 0;
    isSessAmbrPrst_m = false;

    //up transport layer information
    memset(&transLayerAddress_m, 0, sizeof(transLayerAddress_m));

    upTransportLayerInfo_mp = NULL;
    upTransportLayerInfo_m.choice = gTPTunnel_chosen;
    upTransportLayerInfo_m.u.gTPTunnel.bit_mask = 0;
    upTransportLayerInfo_m.u.gTPTunnel.iE_Extensions = NULL;
    upTransportLayerInfo_m.u.gTPTunnel.transportLayerAddress.length = 0;
    upTransportLayerInfo_m.u.gTPTunnel.transportLayerAddress.value = transLayerAddress_m.transLayerAddr;
    upTransportLayerInfo_m.u.gTPTunnel.gTP_TEID.length = 0;
    memset(upTransportLayerInfo_m.u.gTPTunnel.gTP_TEID.value, 0, sizeof(upTransportLayerInfo_m.u.gTPTunnel.gTP_TEID.value));

    //Add Up Transport Layer Information
    for (unsigned int i = 0;i < maxnoofMultiConnectivityMinusOne_c; i++)
    {
        memset(&addTransLayerAddress_m[i],0,sizeof(addTransLayerAddress_m[i]));
    }

    upTransportLayerInfoList_mp = NULL;
    upTransportLayerInfoList_m = NULL;
    numofUpTransportLayerInfo_m = 0;
    for (unsigned int i = 0;i < maxnoofMultiConnectivityMinusOne_c; i++)
    {
        upTransportLayerInfoList_ma[i].value.bit_mask = 0;
        upTransportLayerInfoList_ma[i].value.iE_Extensions = NULL;
        upTransportLayerInfoList_ma[i].value.nGU_UP_TNLInformation.choice = gTPTunnel_chosen;
        upTransportLayerInfoList_ma[i].value.nGU_UP_TNLInformation.u.gTPTunnel.bit_mask = 0;
        upTransportLayerInfoList_ma[i].value.nGU_UP_TNLInformation.u.gTPTunnel.iE_Extensions = NULL;
        upTransportLayerInfoList_ma[i].value.nGU_UP_TNLInformation.u.gTPTunnel.transportLayerAddress.length = 0;
        upTransportLayerInfoList_ma[i].value.nGU_UP_TNLInformation.u.gTPTunnel.transportLayerAddress.value = addTransLayerAddress_m[i].addTransLayerAddr;
        upTransportLayerInfoList_ma[i].value.nGU_UP_TNLInformation.u.gTPTunnel.gTP_TEID.length = 0;
        memset(upTransportLayerInfoList_ma[i].value.nGU_UP_TNLInformation.u.gTPTunnel.gTP_TEID.value,
            0,sizeof(upTransportLayerInfoList_ma[i].value.nGU_UP_TNLInformation.u.gTPTunnel.gTP_TEID.value));
    }
    isUpTransportLayerInfoPrst_m = false;

    // data forwarding not possible
    dataFwdNotPssble_mp = NULL;
    dataFwdNotPssble_m = data_forwarding_not_possible;
    isDataFwdNotPssblePrst_m = false;

    //pdu session type
    pduSessType_mp = NULL;
    pduSessType_m = ipv4;

    //Maximum Integrity Protected Data Rate Downlink
    MaxIntProDataRateDlIe.next = NULL;
    MaxIntProDataRateDlIe.value.id = id_MaximumIntegrityProtectedDataRate_DL;
    MaxIntProDataRateDlIe.value.criticality = ignore;
    MaxIntProDataRateDlIe.value.extensionValue.pduNum = 
        PDU_SecurityIndication_ExtIEs_Extension_MaximumIntegrityProtectedDataRate;
    MaxIntProDataRateDlIe.value.extensionValue.decoded.pdu_MaximumIntegrityProtectedDataRate = &maxIntProDataRate_m;
    MaxIntProDataRateDlIe.value.extensionValue.encoded.length = 0;
    MaxIntProDataRateDlIe.value.extensionValue.encoded.value = NULL;

    //security indication
    secInd_mp = NULL;
    secInd_m.bit_mask = 0;
    secInd_m.iE_Extensions = NULL;
    secInd_m.integrityProtectionIndication = IntegrityProtectionIndication_not_needed;
    secInd_m.confidentialityProtectionIndication = ConfidentialityProtectionIndication_not_needed;
    secInd_m.maximumIntegrityProtectedDataRate_UL = bitrate64kbs;
    secInd_m.iE_Extensions = &MaxIntProDataRateDlIe;
    isSecIndPrst_m = false;

    //network Instance
    networkInstance_mp = NULL;
    networkInstance_m = 0;
    isnetworkInstancePrst_m = false;

    //qos flow setup request list
    qosFlowSetupReqList_mp = NULL;
    qosFlowSetupReqList_m = NULL;
    numofQosFlowSetupReq_m = 0;
    for (unsigned int i = 0; i < maxnoofQosFlows; i++)
    {

        qosFlowSetupReqList_ma[i].value.bit_mask = 0;
        qosFlowSetupReqList_ma[i].value.iE_Extensions = NULL;

        qosFlowSetupReqList_ma[i].value.qosFlowIdentifier = 0;
        qosFlowSetupReqList_ma[i].value.qosFlowLevelQosParameters.bit_mask = 0;
        qosFlowSetupReqList_ma[i].value.qosFlowLevelQosParameters.iE_Extensions = NULL;

        qosFlowSetupReqList_ma[i].value.qosFlowLevelQosParameters.qosCharacteristics.choice = nonDynamic5QI_chosen;
        qosFlowSetupReqList_ma[i].value.qosFlowLevelQosParameters.qosCharacteristics.u.nonDynamic5QI.bit_mask = 0;
        qosFlowSetupReqList_ma[i].value.qosFlowLevelQosParameters.qosCharacteristics.u.nonDynamic5QI.iE_Extensions = NULL;
        qosFlowSetupReqList_ma[i].value.qosFlowLevelQosParameters.qosCharacteristics.u.nonDynamic5QI.fiveQI = 0;
        qosFlowSetupReqList_ma[i].value.qosFlowLevelQosParameters.qosCharacteristics.u.nonDynamic5QI.priorityLevelQos = 0;
        qosFlowSetupReqList_ma[i].value.qosFlowLevelQosParameters.qosCharacteristics.u.nonDynamic5QI.averagingWindow = 0;
        qosFlowSetupReqList_ma[i].value.qosFlowLevelQosParameters.qosCharacteristics.u.nonDynamic5QI.maximumDataBurstVolume = 0;

        qosFlowSetupReqList_ma[i].value.qosFlowLevelQosParameters.allocationAndRetentionPriority.bit_mask = 0;
        qosFlowSetupReqList_ma[i].value.qosFlowLevelQosParameters.allocationAndRetentionPriority.iE_Extensions = NULL;
        qosFlowSetupReqList_ma[i].value.qosFlowLevelQosParameters.allocationAndRetentionPriority.priorityLevelARP = 0;
        qosFlowSetupReqList_ma[i].value.qosFlowLevelQosParameters.allocationAndRetentionPriority.pre_emptionCapability = shall_not_trigger_pre_emption;
        qosFlowSetupReqList_ma[i].value.qosFlowLevelQosParameters.allocationAndRetentionPriority.pre_emptionVulnerability = not_pre_emptable;

        qosFlowSetupReqList_ma[i].value.qosFlowLevelQosParameters.gBR_QosInformation.bit_mask = 0;
        qosFlowSetupReqList_ma[i].value.qosFlowLevelQosParameters.gBR_QosInformation.iE_Extensions = NULL;
        qosFlowSetupReqList_ma[i].value.qosFlowLevelQosParameters.gBR_QosInformation.maximumFlowBitRateDL = 0;
        qosFlowSetupReqList_ma[i].value.qosFlowLevelQosParameters.gBR_QosInformation.maximumFlowBitRateUL = 0;
        qosFlowSetupReqList_ma[i].value.qosFlowLevelQosParameters.gBR_QosInformation.guaranteedFlowBitRateDL = 0;
        qosFlowSetupReqList_ma[i].value.qosFlowLevelQosParameters.gBR_QosInformation.guaranteedFlowBitRateUL = 0;
        qosFlowSetupReqList_ma[i].value.qosFlowLevelQosParameters.gBR_QosInformation.notificationControl = notification_requested;
        qosFlowSetupReqList_ma[i].value.qosFlowLevelQosParameters.gBR_QosInformation.maximumPacketLossRateDL = 0;
        qosFlowSetupReqList_ma[i].value.qosFlowLevelQosParameters.gBR_QosInformation.maximumPacketLossRateUL = 0;

        qosFlowSetupReqList_ma[i].value.qosFlowLevelQosParameters.reflectiveQosAttribute = subject_to;
        qosFlowSetupReqList_ma[i].value.qosFlowLevelQosParameters.additionalQosFlowInformation = more_likely;

        qosFlowSetupReqList_ma[i].value.e_RAB_ID = 0;
        qosFlowSetupReqList_ma[i].next = NULL;
    }

    //Common Network Instance
    commonNetworkInstance_m.length = 0;
    memset(cmmnNtwrkInstance_ma.data,0xff,sizeof(cmmnNtwrkInstance_ma.data));
    commonNetworkInstance_m.value = cmmnNtwrkInstance_ma.data;
    commonNetworkInstance_mp = NULL;
    commonntkInstancePresent_m = false;
}

unsigned int 
PduSesResSetupReqTransferPdu::decodePerToOss(NgapOssCtxt* ctxt)
{
   if(NULL == ctxt)
   {
       return -1;
   }
   
   OssBuf inputBuf;
   OssBuf outputBuf;
   int pduNum = PDUSessionResourceSetupRequestTransfer_PDU;

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
PduSesResSetupReqTransferPdu::encodeOssToPer(NgapOssCtxt *ctxt)
{
    if (NULL == ctxt)
    {
        return -1;
    }

    ProtocolIeType *nextIe_p = NULL;

    //qos flow setup request list
    ProtocolIeType qosFlowSetupReqListIe;
    if(numofQosFlowSetupReq_m > maxnoofQosFlows || numofQosFlowSetupReq_m <= 0)
    {
        return defaultFailureCode_c;
    }
    else
    {
        struct QosFlowSetupRequestList_ *nextQosFlowSetupRequestList_p = NULL;
        for(int i = 0; i < numofQosFlowSetupReq_m; i++)
        {
            qosFlowSetupReqList_ma[i].next = nextQosFlowSetupRequestList_p;
            nextQosFlowSetupRequestList_p = &qosFlowSetupReqList_ma[i];       
        }
        qosFlowSetupReqList_m = nextQosFlowSetupRequestList_p;

        qosFlowSetupReqListIe.next =nextIe_p;
        qosFlowSetupReqListIe.value.id = id_QosFlowSetupRequestList;
        qosFlowSetupReqListIe.value.criticality = reject;
        qosFlowSetupReqListIe.value.value.pduNum = PDU_PDUSessionResourceSetupRequestTransferIEs_Value_QosFlowSetupRequestList;
        qosFlowSetupReqListIe.value.value.decoded.pdu_QosFlowSetupRequestList = &qosFlowSetupReqList_m;
        qosFlowSetupReqListIe.value.value.encoded.length = 0;
        qosFlowSetupReqListIe.value.value.encoded.value = NULL;
        nextIe_p = &qosFlowSetupReqListIe;
    }

    //network instance
    ProtocolIeType networkInstanceIe;
    if(isnetworkInstancePrst_m){
        networkInstanceIe.next =nextIe_p;
        networkInstanceIe.value.id = id_NetworkInstance;
        networkInstanceIe.value.criticality = reject;
        networkInstanceIe.value.value.pduNum = PDU_PDUSessionResourceSetupRequestTransferIEs_Value_NetworkInstance;
        networkInstanceIe.value.value.decoded.pdu_NetworkInstance = &networkInstance_m;
        networkInstanceIe.value.value.encoded.length = 0;
        networkInstanceIe.value.value.encoded.value = NULL;
        nextIe_p = &networkInstanceIe;
    }

    //security indication
    ProtocolIeType secIndIe;
    if(isSecIndPrst_m){
        secIndIe.next =nextIe_p;
        secIndIe.value.id = id_SecurityIndication;
        secIndIe.value.criticality = reject;
        secIndIe.value.value.pduNum = PDU_PDUSessionResourceSetupRequestTransferIEs_Value_SecurityIndication;
        secIndIe.value.value.decoded.pdu_SecurityIndication = &secInd_m;
        secIndIe.value.value.encoded.length = 0;
        secIndIe.value.value.encoded.value = NULL;
        nextIe_p = &secIndIe;
    }

    //pdu session type
    ProtocolIeType pduSessTypeIe;
    pduSessTypeIe.next =nextIe_p;
    pduSessTypeIe.value.id = id_PDUSessionType;
    pduSessTypeIe.value.criticality = reject;
    pduSessTypeIe.value.value.pduNum = PDU_PDUSessionResourceSetupRequestTransferIEs_Value_PDUSessionType;
    pduSessTypeIe.value.value.decoded.pdu_PDUSessionType = &pduSessType_m;
    pduSessTypeIe.value.value.encoded.length = 0;
    pduSessTypeIe.value.value.encoded.value = NULL;
    nextIe_p = &pduSessTypeIe;

    //data forwarding not possible
    ProtocolIeType dataFwdNotPssbleIe;
    if(isDataFwdNotPssblePrst_m){
        dataFwdNotPssbleIe.next =nextIe_p;
        dataFwdNotPssbleIe.value.id = id_DataForwardingNotPossible;
        dataFwdNotPssbleIe.value.criticality = reject;
        dataFwdNotPssbleIe.value.value.pduNum = PDU_PDUSessionResourceSetupRequestTransferIEs_Value_DataForwardingNotPossible;
        dataFwdNotPssbleIe.value.value.decoded.pdu_DataForwardingNotPossible = &dataFwdNotPssble_m;
        dataFwdNotPssbleIe.value.value.encoded.length = 0;
        dataFwdNotPssbleIe.value.value.encoded.value = NULL;
        nextIe_p = &dataFwdNotPssbleIe;
    }

    //additional upTransportLayerInfo
    ProtocolIeType upTransportLayerInfoListIe;
//    std::cout <<"test question:" <<isUpTransportLayerInfoPrst_m << std::endl;
    if (isUpTransportLayerInfoPrst_m) {
        if(numofUpTransportLayerInfo_m > maxnoofMultiConnectivityMinusOne || numofUpTransportLayerInfo_m < 0)
        {
            return defaultFailureCode_c;
        }
        else
        {
            struct UPTransportLayerInformationList_ *nextUpTransportLayerInformationList_p = NULL;
            for (int i = 0; i < numofUpTransportLayerInfo_m; i++)
            {
                upTransportLayerInfoList_ma[i].next = nextUpTransportLayerInformationList_p;
                nextUpTransportLayerInformationList_p = &upTransportLayerInfoList_ma[i];
            }
            upTransportLayerInfoList_m = nextUpTransportLayerInformationList_p;
        }
        upTransportLayerInfoListIe.next =nextIe_p;
        upTransportLayerInfoListIe.value.id = id_AdditionalUL_NGU_UP_TNLInformation;
        upTransportLayerInfoListIe.value.criticality = reject;
        upTransportLayerInfoListIe.value.value.pduNum = PDU_PDUSessionResourceSetupRequestTransferIEs_Value_UPTransportLayerInformationList;
        upTransportLayerInfoListIe.value.value.decoded.pdu_UPTransportLayerInformationList = &upTransportLayerInfoList_m;
        upTransportLayerInfoListIe.value.value.encoded.length = 0;
        upTransportLayerInfoListIe.value.value.encoded.value = NULL;
        nextIe_p = &upTransportLayerInfoListIe;
    }

    //upTransportLayerInfo
    ProtocolIeType upTransportLayerInfoIe;
    upTransportLayerInfoIe.next =nextIe_p;
    upTransportLayerInfoIe.value.id = id_UL_NGU_UP_TNLInformation;
    upTransportLayerInfoIe.value.criticality = reject;
    upTransportLayerInfoIe.value.value.pduNum = PDU_PDUSessionResourceSetupRequestTransferIEs_Value_UPTransportLayerInformation;
    upTransportLayerInfoIe.value.value.decoded.pdu_UPTransportLayerInformation = &upTransportLayerInfo_m;
    upTransportLayerInfoIe.value.value.encoded.length = 0;
    upTransportLayerInfoIe.value.value.encoded.value = NULL;
    nextIe_p = &upTransportLayerInfoIe;

    //Session ambr
    ProtocolIeType sessAmbrIe;
    if(isSessAmbrPrst_m){
        sessAmbrIe.next =nextIe_p;
        sessAmbrIe.value.id = id_PDUSessionAggregateMaximumBitRate;
        sessAmbrIe.value.criticality = reject;
        sessAmbrIe.value.value.pduNum = PDU_PDUSessionResourceSetupRequestTransferIEs_Value_PDUSessionAggregateMaximumBitRate;
        sessAmbrIe.value.value.decoded.pdu_PDUSessionAggregateMaximumBitRate = &sessAMBR_m;
        sessAmbrIe.value.value.encoded.length = 0;
        sessAmbrIe.value.value.encoded.value = NULL;
        nextIe_p = &sessAmbrIe;
    }

    //Common Network Instance
    ProtocolIeType commonNetworkInstaceIe;
    if(commonntkInstancePresent_m){
        commonNetworkInstaceIe.next =nextIe_p;
        commonNetworkInstaceIe.value.id = id_CommonNetworkInstance;
        commonNetworkInstaceIe.value.criticality = ignore;
        commonNetworkInstaceIe.value.value.pduNum = PDU_PDUSessionResourceSetupRequestTransferIEs_Value_CommonNetworkInstance;
        commonNetworkInstaceIe.value.value.decoded.pdu_CommonNetworkInstance = &commonNetworkInstance_m;
        commonNetworkInstaceIe.value.value.encoded.length = 0;
        commonNetworkInstaceIe.value.value.encoded.value = NULL;
        nextIe_p = &commonNetworkInstaceIe;
    }

    //message
    PDUSessionResourceSetupRequestTransfer msg;
    msg.protocolIEs = nextIe_p;

    //Convert Structures to PER
    OssBuf outputBuf;
    unsigned int result;

    ctxt->reset_v();
    outputBuf = ctxt->getPrimaryCodecBuf();

    result = ctxt->encodeOssToPer(PDUSessionResourceSetupRequestTransfer_PDU, &msg, &outputBuf);

    if (result == encodeSuccessCode_c)
    {
        perBufRef_m = outputBuf;
    }

    return result;
}

void PduSesResSetupReqTransferPdu::getDecodedOssDataPtrs_v()
{
    ProtocolIeType *currIe_p;
    reset_v();

    pduSessResSetupReqTransfer_mp = reinterpret_cast<PDUSessionResourceSetupRequestTransfer *>(ossBufRef_m.value);

    if (pduSessResSetupReqTransfer_mp == NULL)
    {
        return;
    }

    currIe_p = pduSessResSetupReqTransfer_mp->protocolIEs;

    while (currIe_p != NULL)
    {
        if (currIe_p->value.id == id_PDUSessionAggregateMaximumBitRate
            && currIe_p->value.value.decoded.pdu_PDUSessionAggregateMaximumBitRate != NULL)
        {
            sessAMBR_mp = 
                      currIe_p->value.value.decoded.pdu_PDUSessionAggregateMaximumBitRate;
            isSessAmbrPrst_m = true;
        }
        else if (currIe_p->value.id == id_UL_NGU_UP_TNLInformation
            && currIe_p->value.value.decoded.pdu_UPTransportLayerInformation != NULL)
        {
            upTransportLayerInfo_mp = 
                      currIe_p->value.value.decoded.pdu_UPTransportLayerInformation;
        }
        else if (currIe_p->value.id == id_AdditionalUL_NGU_UP_TNLInformation
                 && currIe_p->value.value.decoded.pdu_UPTransportLayerInformation != NULL)
        {
            upTransportLayerInfoList_mp = currIe_p->value.value.decoded.pdu_UPTransportLayerInformationList;
            UPTransportLayerInformationList curr_p = *upTransportLayerInfoList_mp;
            while(curr_p)
            {
                numofUpTransportLayerInfo_m++;
                curr_p = curr_p->next;
            }
            isUpTransportLayerInfoPrst_m = true;
        }
        else if (currIe_p->value.id == id_DataForwardingNotPossible
            && currIe_p->value.value.decoded.pdu_DataForwardingNotPossible != NULL)
        {
            dataFwdNotPssble_mp = 
                currIe_p->value.value.decoded.pdu_DataForwardingNotPossible;
            isDataFwdNotPssblePrst_m = true;
        }
        else if (currIe_p->value.id == id_PDUSessionType
            && currIe_p->value.value.decoded.pdu_PDUSessionType != NULL)
        {
            pduSessType_mp = 
               currIe_p->value.value.decoded.pdu_PDUSessionType;
        }
        else if (currIe_p->value.id == id_SecurityIndication
            && currIe_p->value.value.decoded.pdu_SecurityIndication != NULL)
        {
            secInd_mp = 
               currIe_p->value.value.decoded.pdu_SecurityIndication;
            isSecIndPrst_m = true;
        }
        else if (currIe_p->value.id == id_NetworkInstance
            && currIe_p->value.value.decoded.pdu_NetworkInstance != NULL)
        {
            networkInstance_mp = 
                currIe_p->value.value.decoded.pdu_NetworkInstance;
            isnetworkInstancePrst_m = true;
        }
        else if (currIe_p->value.id == id_QosFlowSetupRequestList
            && currIe_p->value.value.decoded.pdu_QosFlowSetupRequestList != NULL)
        {
            qosFlowSetupReqList_mp = currIe_p->value.value.decoded.pdu_QosFlowSetupRequestList;
            QosFlowSetupRequestList curr_p = *qosFlowSetupReqList_mp;
            while(curr_p != NULL)
            {
                numofQosFlowSetupReq_m++;
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

PDUSessionResourceSetupRequestTransfer *
PduSesResSetupReqTransferPdu::getPduSessResSetupReqTransferDecodedPtr() const
{
    return pduSessResSetupReqTransfer_mp;
}

PDUSessionAggregateMaximumBitRate *
PduSesResSetupReqTransferPdu::getSessAmbrDecodedPtr() const
{
    return sessAMBR_mp;
}

PDUSessionAggregateMaximumBitRate *
PduSesResSetupReqTransferPdu::getSessAmbrEncodePtr()
{
    isSessAmbrPrst_m = true; //should set the value outside
    return &sessAMBR_m;
}

bool 
PduSesResSetupReqTransferPdu::isSessAmbrPrst()
{
   return isSessAmbrPrst_m;
}

UPTransportLayerInformation *
PduSesResSetupReqTransferPdu::getUpTransLayerInfoDecodedPtr() const
{
    return upTransportLayerInfo_mp;
}

UPTransportLayerInformation *
PduSesResSetupReqTransferPdu::getUpTransLayerInfoEncodePtr()
{
    return &upTransportLayerInfo_m;
}

UPTransportLayerInformationList*
PduSesResSetupReqTransferPdu::getUpTransLayerInfoListDecodedPtr() const
{
    return upTransportLayerInfoList_mp;
}

UPTransportLayerInformationList*
PduSesResSetupReqTransferPdu::getUpTransLayerInfoListEncodePtr()
{
    isUpTransportLayerInfoPrst_m = true;
    return &upTransportLayerInfoList_m;
}

UPTransportLayerInformationItem*
PduSesResSetupReqTransferPdu::getUpTransLayerInfoItemDecodedPtr(unsigned int index) const
{
    if (index >=0 && index <maxnoofMultiConnectivityMinusOne)
    {
        UPTransportLayerInformationItem *item_p = NULL;
        UPTransportLayerInformationList upTransportLayerInfoList_p = *upTransportLayerInfoList_mp;

        if (upTransportLayerInfoList_p != NULL
            && (index >= 0 && index < maxnoofMultiConnectivityMinusOne_c))
        {
            unsigned int i = 0;
            for (;i < index && upTransportLayerInfoList_p != NULL; i++)
            {
                upTransportLayerInfoList_p = upTransportLayerInfoList_p->next;
            }
            if (upTransportLayerInfoList_p != NULL && i == index)
            {
                item_p = &(upTransportLayerInfoList_p->value);
            }
        }
        return item_p;
    }
    else
    {
        return NULL;
    }
}

UPTransportLayerInformationItem*
PduSesResSetupReqTransferPdu::getUpTransLayerInfoItemEncodePtr(unsigned int index)
{
    if (index >= 0 && index <maxnoofMultiConnectivityMinusOne)
    {
        isUpTransportLayerInfoPrst_m = true;
        return &upTransportLayerInfoList_ma[index].value;
    }
    else
    {
        return NULL;
    }
}

void
PduSesResSetupReqTransferPdu::setNumofUpTransLayerInfoListItem_m_v(unsigned int num)
{
    numofUpTransportLayerInfo_m = num;
}

unsigned int
PduSesResSetupReqTransferPdu::getNumofUpTransLayerInfoListItem() const
{
    return numofUpTransportLayerInfo_m;
}

bool 
PduSesResSetupReqTransferPdu::isUpTransLayerInfoListPrst()
{
    return isUpTransportLayerInfoPrst_m;
}

DataForwardingNotPossible *
PduSesResSetupReqTransferPdu::getDataFwdNotPssbleDecodedPtr() const
{
    return dataFwdNotPssble_mp;
}

DataForwardingNotPossible *
PduSesResSetupReqTransferPdu::getDataFwdNotPssbleEncodePtr()
{
    isDataFwdNotPssblePrst_m = true;
    return &dataFwdNotPssble_m;
}

bool 
PduSesResSetupReqTransferPdu::isDataFwdNotPssblePrst()
{
    return isDataFwdNotPssblePrst_m;
}

PDUSessionType *
PduSesResSetupReqTransferPdu::getPduSessTypeDecodedPtr() const
{
    return pduSessType_mp;
}

PDUSessionType *
PduSesResSetupReqTransferPdu::getPduSessTypeEncodePtr()
{
    return &pduSessType_m;
}

SecurityIndication *
PduSesResSetupReqTransferPdu::getSecIndDecodedPtr() const
{
    return secInd_mp;
}

SecurityIndication *
PduSesResSetupReqTransferPdu::getSecIndEncodePtr()
{
    isSecIndPrst_m = true;
    return &secInd_m;
}
bool 
PduSesResSetupReqTransferPdu::isSecIndPrst()
{
    return isSecIndPrst_m;
}

MaximumIntegrityProtectedDataRate* 
PduSesResSetupReqTransferPdu::getMaxIntProDataRateEncodePtr()
{
    return &maxIntProDataRate_m;
}

SecurityIndicationIeExtension* 
PduSesResSetupReqTransferPdu::getSecIndIeExtenDecodedPtr()
{
    return &MaxIntProDataRateDlIe;
}

NetworkInstance *
PduSesResSetupReqTransferPdu::getNetworkInstanceDecodedPtr() const
{
    return networkInstance_mp;
}

NetworkInstance *
PduSesResSetupReqTransferPdu::getNetworkInstanceEncodePtr()
{
    isnetworkInstancePrst_m = true;
    return &networkInstance_m;
}

bool 
PduSesResSetupReqTransferPdu::isNetworkInstancePrst()
{
    return isnetworkInstancePrst_m;
}

QosFlowSetupRequestList *
PduSesResSetupReqTransferPdu::getQosFlowSetupReqListDecodedPtr() const
{
    return qosFlowSetupReqList_mp;
}

QosFlowSetupRequestList *
PduSesResSetupReqTransferPdu::getQosFlowSetupReqListEncodePtr()
{
    return &qosFlowSetupReqList_m;
}


QosFlowSetupRequestItem*
PduSesResSetupReqTransferPdu::getQosFlowSetupReqItemDecodedPtr(unsigned int index) const
{
    if (index >= 0 && index < maxnoofPDUSessions)
    {
        QosFlowSetupRequestItem *item_p = NULL;
        QosFlowSetupRequestList qosFlowSetupRequestList_p = *qosFlowSetupReqList_mp;

        if (qosFlowSetupRequestList_p != NULL
            && (index >= 0 && index < maxnoofQosFlows))
        {
            unsigned int i = 0;
            for (; i < index && qosFlowSetupRequestList_p != NULL; i++)
            {
                qosFlowSetupRequestList_p = qosFlowSetupRequestList_p->next;
            }
            if (qosFlowSetupRequestList_p != NULL && i == index)
            {
                item_p = &(qosFlowSetupRequestList_p->value);
            }
        }
        return item_p;
    }
    else
    {
        return NULL;
    }
}

QosFlowSetupRequestItem*
PduSesResSetupReqTransferPdu::getQosFlowSetupReqItemEncodePtr(unsigned int index)
{
    if (index >= 0 && index < maxnoofPDUSessions)
    {
        return &qosFlowSetupReqList_ma[index].value;
    }
    else
    {
        return NULL;
    }
}

void
PduSesResSetupReqTransferPdu::setNumofQosFlowSetupReqListItem_m_v(unsigned int num)
{
    numofQosFlowSetupReq_m = num;
}

unsigned int
PduSesResSetupReqTransferPdu::getNumofQosFlowSetupReqListItem() const
{
   return numofQosFlowSetupReq_m;
}

CommonNetworkInstance *
PduSesResSetupReqTransferPdu::getCommonNetworkInstanceDecodedPtr() const
{
    return commonNetworkInstance_mp;
}

CommonNetworkInstance *
PduSesResSetupReqTransferPdu::getCommonNetworkInstanceEncodePtr()
{
    commonntkInstancePresent_m = true;
    return &commonNetworkInstance_m;
}

bool
PduSesResSetupReqTransferPdu::isCommonNetworkInstancePrst()
{
    return commonntkInstancePresent_m;
}

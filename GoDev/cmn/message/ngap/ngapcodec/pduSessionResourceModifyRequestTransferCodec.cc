#include <iostream>
#include <sstream>
#include <string>

#include "pduSessionResourceModifyRequestTransferCodec.h"

using namespace std;

PduSesResModReqTransferCodec::PduSesResModReqTransferCodec() {
    reset();
}

PduSesResModReqTransferCodec::~PduSesResModReqTransferCodec() {}

bool 
PduSesResModReqTransferCodec::isSessAMBRPrst(){
    return isSessAMBRPrst_m;
}

void 
PduSesResModReqTransferCodec::setSessAmbr(long ulBitRate, long dlBitRate){
    sessAmbrULBitRate_m = ulBitRate;
    sessAmbrDLBitRate_m = dlBitRate;
    isSessAMBRPrst_m = true;
}

long 
PduSesResModReqTransferCodec::getSessAmbrUl(){
    return sessAmbrULBitRate_m;
}

long 
PduSesResModReqTransferCodec::getSessAmbrDl(){
    return sessAmbrDLBitRate_m;
}

bool
PduSesResModReqTransferCodec::isNtwkInstancePrst(){
    return isNtwkInstancePrst_m;
}

void 
PduSesResModReqTransferCodec::setNtwkInstance(Uint64 ntwkInstanceVal){
    ntwkInstance_m = ntwkInstanceVal;
    isNtwkInstancePrst_m = true;
}

Uint64 
PduSesResModReqTransferCodec::getNtwkInstance(){
    return ntwkInstance_m;
}

void 
PduSesResModReqTransferCodec::addUlNguUpTnlModifyList(UlNguUpTnlModifyItem& item){
    ulNguUpTnlModifyList_m.push_back(item);
}

UlNguUpTnlModifyList& 
PduSesResModReqTransferCodec::getUlNguUpTnlModifyList(){
    return ulNguUpTnlModifyList_m;
}

void 
PduSesResModReqTransferCodec::addAdditionalGtpTunnelList(GtpTunnel& item) {
    addGtpTunnelListPrst_m = true;
    addGtpTunnelList_m.push_back(item);
}

AddUpTransportLayerInfoList& 
PduSesResModReqTransferCodec::getAdditionalGtpTunnelList(){
    return addGtpTunnelList_m;
}

bool 
PduSesResModReqTransferCodec::isAddGtpTunnelPrst(){
    return addGtpTunnelListPrst_m;
}

void 
PduSesResModReqTransferCodec::addQosFlowAddOrModReqList(QosFlowAddOrModReqItem& item){
    qosFlowAddOrModReqList_m.push_back(item);
}
QosFlowAddOrModReqList& 
PduSesResModReqTransferCodec::getQosFlowAddOrModReqList(){
    return qosFlowAddOrModReqList_m;
}

void 
PduSesResModReqTransferCodec::addQosFlowToReleaseList(QosFlowCodecItem& item){
    qosFlowToReleaseList_m.push_back(item);
}
QosFlowCodecList& 
PduSesResModReqTransferCodec::getQosFlowToReleaseList(){
    return qosFlowToReleaseList_m;
}

void 
PduSesResModReqTransferCodec::setCmmnNtwrkInstance(std::string networkInstance)
{
    cmmnNtwkInstancePrst_m = true;
    cmmnNtwkInstance_m = networkInstance;
}

std::string 
PduSesResModReqTransferCodec::getCmmnNtwrkInstance()
{
    return cmmnNtwkInstance_m;
}

bool 
PduSesResModReqTransferCodec::isCmmnNtwrkInstacePrst()
{
    return cmmnNtwkInstancePrst_m;
}


bool 
PduSesResModReqTransferCodec::reset(){
    msgPdu_m.reset_v();
    isSessAMBRPrst_m = false;
    sessAmbrULBitRate_m = 0;
    sessAmbrDLBitRate_m = 0;

    ulNguUpTnlModifyList_m.clear();

    addGtpTunnelList_m.clear();
    addGtpTunnelListPrst_m = false;

    cmmnNtwkInstance_m.clear();
    cmmnNtwkInstancePrst_m = false;

    ntwkInstance_m = 0;
    isNtwkInstancePrst_m = false;

    qosFlowAddOrModReqList_m.clear();

    qosFlowToReleaseList_m.clear();
}

//bool PduSesResModReqTransferCodec::encode(NgapOssCtxt* ctxt, MsgBuffer msgBuffer)
MsgBuffer PduSesResModReqTransferCodec::encode(NgapOssCtxt* ctxt)
{
    MsgBuffer msgBuffer;
    msgBuffer.length = 0;
    msgBuffer.value = NULL;

    // session ambr 
    if(isSessAMBRPrst_m){
       PDUSessionAggregateMaximumBitRate * ambr = msgPdu_m.getSessAmbrEncodePtr();
       ambr->pDUSessionAggregateMaximumBitRateDL = sessAmbrDLBitRate_m;
       ambr->pDUSessionAggregateMaximumBitRateUL = sessAmbrULBitRate_m;
    }

    //network instance
    if(isNtwkInstancePrst_m){
        *msgPdu_m.getNetworkInstanceEncodePtr() = ntwkInstance_m;
    }

    // ul gnu tnl modify list
    unsigned int num = ulNguUpTnlModifyList_m.size();
    msgPdu_m.setNumofUlNguUpTnlModifyListItem(num);
    for (int i=0; i < num; i++){
        UL_NGU_UP_TNLModifyItem* item = msgPdu_m.getUlNguUpTnlModifyItemEncodePtr(i);
        UlNguUpTnlModifyItem* ulNgTnlModInfo = &ulNguUpTnlModifyList_m[i];

        if (ulNgTnlModInfo->uplinkGtpTunnel.ipType == ipv4){
            item->uL_NGU_UP_TNLInformation.u.gTPTunnel.transportLayerAddress.length = 32; // 32 bit for ipv4
            memcpy(item->uL_NGU_UP_TNLInformation.u.gTPTunnel.transportLayerAddress.value, 
                ulNgTnlModInfo->uplinkGtpTunnel.transportLayerAddr, 4);
        }else if(ulNgTnlModInfo->uplinkGtpTunnel.ipType == ipv6){
            item->uL_NGU_UP_TNLInformation.u.gTPTunnel.transportLayerAddress.length = 128; // 16*8=128 bit for ipv6
            memcpy(item->uL_NGU_UP_TNLInformation.u.gTPTunnel.transportLayerAddress.value, 
                ulNgTnlModInfo->uplinkGtpTunnel.transportLayerAddr, 16);
        }
        item->uL_NGU_UP_TNLInformation.u.gTPTunnel.gTP_TEID.length = sizeof(ulNgTnlModInfo->uplinkGtpTunnel.gtpTeid);
        memcpy(item->uL_NGU_UP_TNLInformation.u.gTPTunnel.gTP_TEID.value,
              ulNgTnlModInfo->uplinkGtpTunnel.gtpTeid,sizeof( ulNgTnlModInfo->uplinkGtpTunnel.gtpTeid));

        if (ulNgTnlModInfo->downlinkGtpTunnel.ipType == ipv4){
            item->dL_NGU_UP_TNLInformation.u.gTPTunnel.transportLayerAddress.length = 32; // 32 bit for ipv4
            memcpy(item->dL_NGU_UP_TNLInformation.u.gTPTunnel.transportLayerAddress.value, 
                   ulNgTnlModInfo->downlinkGtpTunnel.transportLayerAddr, 4);
        }else if(ulNgTnlModInfo->downlinkGtpTunnel.ipType == ipv6){
            item->dL_NGU_UP_TNLInformation.u.gTPTunnel.transportLayerAddress.length = 128; // 16*8=128 bit for ipv6
            memcpy(item->dL_NGU_UP_TNLInformation.u.gTPTunnel.transportLayerAddress.value, 
                   ulNgTnlModInfo->downlinkGtpTunnel.transportLayerAddr, 16);
        }
        item->dL_NGU_UP_TNLInformation.u.gTPTunnel.gTP_TEID.length = sizeof(ulNgTnlModInfo->downlinkGtpTunnel.gtpTeid);
        memcpy(item->dL_NGU_UP_TNLInformation.u.gTPTunnel.gTP_TEID.value,
               ulNgTnlModInfo->downlinkGtpTunnel.gtpTeid,sizeof( ulNgTnlModInfo->downlinkGtpTunnel.gtpTeid));
    }

    //additional ul ng-u up tnl information list
    if(addGtpTunnelListPrst_m){
        unsigned int num = addGtpTunnelList_m.size();
        msgPdu_m.setNumofAddUpTransportLayerInfoListItem(num);
        for (int i=0; i < num; i++){
            UPTransportLayerInformationItem* item = msgPdu_m.getAddUpTransportLayerInfoItemEncodePtr(i);
            GtpTunnel* addGtpTunnelInfo = &addGtpTunnelList_m[i];
            if (addGtpTunnelInfo->ipType == ipv4) {
                item->nGU_UP_TNLInformation.u.gTPTunnel.transportLayerAddress.length = 32; // 32 bit for ipv4
                memcpy(item->nGU_UP_TNLInformation.u.gTPTunnel.transportLayerAddress.value, 
                    addGtpTunnelInfo->transportLayerAddr, 4);
            }else if(addGtpTunnelInfo->ipType == ipv6){
                item->nGU_UP_TNLInformation.u.gTPTunnel.transportLayerAddress.length = 128; // 16*8=128 bit for ipv6
                memcpy(item->nGU_UP_TNLInformation.u.gTPTunnel.transportLayerAddress.value, 
                    addGtpTunnelInfo->transportLayerAddr, 16);
            }
            item->nGU_UP_TNLInformation.u.gTPTunnel.gTP_TEID.length = sizeof(addGtpTunnelInfo->gtpTeid);
            memcpy(item->nGU_UP_TNLInformation.u.gTPTunnel.gTP_TEID.value,
                  addGtpTunnelInfo->gtpTeid,sizeof( addGtpTunnelInfo->gtpTeid));
        }
    }
    
    // qos flow setup request list
    num = qosFlowAddOrModReqList_m.size();
    msgPdu_m.setNumofQosFlowAddOrModReqListItem(num);
    for (int i=0; i < num; i++){
        QosFlowAddOrModifyRequestItem* item = msgPdu_m.getQosFlowAddOrModReqItemEncodePtr(i);
        QosFlowAddOrModReqItem* qosFlow = &qosFlowAddOrModReqList_m[i];

        // qfi
        item->qosFlowIdentifier = qosFlow->qosFlowInd;

        // qos flow parameter
        if(qosFlow->qosFlowParaPresent){
            if(qosFlow->qosFlowLevQosPara.qosCharacter.isDynamic){
                item->qosFlowLevelQosParameters.qosCharacteristics.choice = dynamic5QI_chosen;
                item->qosFlowLevelQosParameters.qosCharacteristics.u.dynamic5QI.priorityLevelQos
                    = qosFlow->qosFlowLevQosPara.qosCharacter.dynamic5qi.priorityLevelQos;
                item->qosFlowLevelQosParameters.qosCharacteristics.u.dynamic5QI.packetDelayBudget
                    = qosFlow->qosFlowLevQosPara.qosCharacter.dynamic5qi.packetDelayBudget;
                item->qosFlowLevelQosParameters.qosCharacteristics.u.dynamic5QI.packetErrorRate.pERExponent
                    = qosFlow->qosFlowLevQosPara.qosCharacter.dynamic5qi.packetErrorRate.pERExponent;
                item->qosFlowLevelQosParameters.qosCharacteristics.u.dynamic5QI.packetErrorRate.pERScalar
                    = qosFlow->qosFlowLevQosPara.qosCharacter.dynamic5qi.packetErrorRate.pERScalar;
                if(qosFlow->qosFlowLevQosPara.qosCharacter.dynamic5qi.is5QIPrst){
                    item->qosFlowLevelQosParameters.qosCharacteristics.u.dynamic5QI.fiveQI
                        = qosFlow->qosFlowLevQosPara.qosCharacter.dynamic5qi.fiveQI;
                    item->qosFlowLevelQosParameters.qosCharacteristics.u.dynamic5QI.bit_mask 
                        |= fiveQI_present;
                }
                if(qosFlow->qosFlowLevQosPara.qosCharacter.dynamic5qi.isAveWindowPrst){
                    item->qosFlowLevelQosParameters.qosCharacteristics.u.dynamic5QI.averagingWindow
                        = qosFlow->qosFlowLevQosPara.qosCharacter.dynamic5qi.averagingWindow;        
                    item->qosFlowLevelQosParameters.qosCharacteristics.u.dynamic5QI.bit_mask 
                        |= Dynamic5QIDescriptor_averagingWindow_present;
                }                      
                if(qosFlow->qosFlowLevQosPara.qosCharacter.dynamic5qi.isDelayCriticalPrst){
                    item->qosFlowLevelQosParameters.qosCharacteristics.u.dynamic5QI.delayCritical
                        = (DelayCritical)(qosFlow->qosFlowLevQosPara.qosCharacter.dynamic5qi.delayCritical); 
                    item->qosFlowLevelQosParameters.qosCharacteristics.u.dynamic5QI.bit_mask 
                        |= delayCritical_present;
                }                   
                if(qosFlow->qosFlowLevQosPara.qosCharacter.dynamic5qi.isMaxDataBusrtVolPrst){
                    item->qosFlowLevelQosParameters.qosCharacteristics.u.dynamic5QI.maximumDataBurstVolume
                        = qosFlow->qosFlowLevQosPara.qosCharacter.dynamic5qi.maximumDataBurstVolume; 
                    item->qosFlowLevelQosParameters.qosCharacteristics.u.dynamic5QI.bit_mask 
                        |= Dynamic5QIDescriptor_maximumDataBurstVolume_present;
                }                       
            }
            else{
                item->qosFlowLevelQosParameters.qosCharacteristics.choice = nonDynamic5QI_chosen;        
                item->qosFlowLevelQosParameters.qosCharacteristics.u.nonDynamic5QI.fiveQI
                    = qosFlow->qosFlowLevQosPara.qosCharacter.nonDynamic5qi.fiveQI;
                if(qosFlow->qosFlowLevQosPara.qosCharacter.nonDynamic5qi.isMaxDataBusrtVolPrst){
                    item->qosFlowLevelQosParameters.qosCharacteristics.u.nonDynamic5QI.maximumDataBurstVolume
                        = qosFlow->qosFlowLevQosPara.qosCharacter.nonDynamic5qi.maximumDataBurstVolume;      
                    item->qosFlowLevelQosParameters.qosCharacteristics.u.nonDynamic5QI.bit_mask 
                        |= NonDynamic5QIDescriptor_maximumDataBurstVolume_present;
                }
                if(qosFlow->qosFlowLevQosPara.qosCharacter.nonDynamic5qi.isPriorityLevelQosPrst){
                    item->qosFlowLevelQosParameters.qosCharacteristics.u.nonDynamic5QI.priorityLevelQos
                        = qosFlow->qosFlowLevQosPara.qosCharacter.nonDynamic5qi.priorityLevelQos;    
                    item->qosFlowLevelQosParameters.qosCharacteristics.u.nonDynamic5QI.bit_mask 
                        |= priorityLevelQos_present;
                }
                if(qosFlow->qosFlowLevQosPara.qosCharacter.nonDynamic5qi.isAveWindowPrst){
                    item->qosFlowLevelQosParameters.qosCharacteristics.u.nonDynamic5QI.averagingWindow
                        = qosFlow->qosFlowLevQosPara.qosCharacter.nonDynamic5qi.averagingWindow;     
                    item->qosFlowLevelQosParameters.qosCharacteristics.u.nonDynamic5QI.bit_mask 
                        |= NonDynamic5QIDescriptor_averagingWindow_present;
                }
            }
        
            // qos flow parameter - arp
            item->qosFlowLevelQosParameters.allocationAndRetentionPriority.priorityLevelARP 
                = qosFlow->qosFlowLevQosPara.arp.priorityLevel;
            item->qosFlowLevelQosParameters.allocationAndRetentionPriority.pre_emptionCapability
                = (Pre_emptionCapability)(qosFlow->qosFlowLevQosPara.arp.emptionCapability);
            item->qosFlowLevelQosParameters.allocationAndRetentionPriority.pre_emptionVulnerability 
                = (Pre_emptionVulnerability)(qosFlow->qosFlowLevQosPara.arp.emptionVulnerability);
        
            // qos flow parameter - gbr qos info
            if(qosFlow->qosFlowLevQosPara.gBRQosInfoPresent){
                item->qosFlowLevelQosParameters.bit_mask 
                    |= gBR_QosInformation_present;
                item->qosFlowLevelQosParameters.gBR_QosInformation.maximumFlowBitRateDL 
                    = qosFlow->qosFlowLevQosPara.gBRQosInfo.maxFlowBitRateDL;
                item->qosFlowLevelQosParameters.gBR_QosInformation.maximumFlowBitRateUL
                    = qosFlow->qosFlowLevQosPara.gBRQosInfo.maxFlowBitRateUL;
                item->qosFlowLevelQosParameters.gBR_QosInformation.guaranteedFlowBitRateDL
                    = qosFlow->qosFlowLevQosPara.gBRQosInfo.guaFlowBitRateDL;
                item->qosFlowLevelQosParameters.gBR_QosInformation.guaranteedFlowBitRateUL
                    = qosFlow->qosFlowLevQosPara.gBRQosInfo.guaFlowBitRateUL;
                if(qosFlow->qosFlowLevQosPara.gBRQosInfo.notiControlPresent){
                    item->qosFlowLevelQosParameters.gBR_QosInformation.bit_mask 
                        |= notificationControl_present;
                    item->qosFlowLevelQosParameters.gBR_QosInformation.notificationControl
                        = (NotificationControl)(qosFlow->qosFlowLevQosPara.gBRQosInfo.notiControl);
                }
                if(qosFlow->qosFlowLevQosPara.gBRQosInfo.maxPacketLossRateDLPresent){
                    item->qosFlowLevelQosParameters.gBR_QosInformation.bit_mask 
                        |= maximumPacketLossRateDL_present;
                    item->qosFlowLevelQosParameters.gBR_QosInformation.maximumPacketLossRateDL
                        = qosFlow->qosFlowLevQosPara.gBRQosInfo.maxPacketLossRateDL;
                }
                if(qosFlow->qosFlowLevQosPara.gBRQosInfo.maxPacketLossRateULPresent){
                    item->qosFlowLevelQosParameters.gBR_QosInformation.bit_mask 
                        |= maximumPacketLossRateUL_present;
                    item->qosFlowLevelQosParameters.gBR_QosInformation.maximumPacketLossRateUL
                        = qosFlow->qosFlowLevQosPara.gBRQosInfo.maxPacketLossRateUL;
                }
            }
        
            // qos flow parameter - refQosAttrPresent
            if(qosFlow->qosFlowLevQosPara.refQosAttrPresent){
                item->qosFlowLevelQosParameters.bit_mask 
                    |= reflectiveQosAttribute_present;                      
                item->qosFlowLevelQosParameters.reflectiveQosAttribute
                    = (ReflectiveQosAttribute)(qosFlow->qosFlowLevQosPara.refQosAttr);
             }
        
             // qos flow parameter - addQosFlowInfo
             if(qosFlow->qosFlowLevQosPara.addQosFlowInfoPresent){
                 item->qosFlowLevelQosParameters.bit_mask 
                     |= additionalQosFlowInformation_present;                        
                 item->qosFlowLevelQosParameters.additionalQosFlowInformation
                     = (AdditionalQosFlowInformation)(qosFlow->qosFlowLevQosPara.addQosFlowInfo);
             }

             item->bit_mask |= qosFlowLevelQosParameters_present;
        }

        // qos flow e-rab id
        if(qosFlow->eRABIdPresent){
            item->e_RAB_ID = (long)qosFlow->eRABId;
            item->bit_mask |= QosFlowAddOrModifyRequestItem_e_RAB_ID_present;
        }
   }

    // qos flow setup request list
    num = qosFlowToReleaseList_m.size();
    msgPdu_m.setNumofQosFlowToRelListItem(num);
    for (int i=0; i < num; i++){
        QosFlowWithCauseItem* item = msgPdu_m.getQosFlowToRelItemEncodePtr(i);
        QosFlowCodecItem* qosFlow = &qosFlowToReleaseList_m[i];

        // qfi
        item->qosFlowIdentifier = qosFlow->qosFlowInd;
        switch(qosFlow->causeType){
            case radioNetwork_chosen:
                item->cause.choice = radioNetwork_chosen;
                item->cause.u.radioNetwork = CauseRadioNetwork(qosFlow->causeValue);
                break;
            case transport_chosen:
                item->cause.choice = transport_chosen;
                item->cause.u.transport = CauseTransport(qosFlow->causeValue);
                break;                              
            case nas_chosen:
                item->cause.choice = nas_chosen;
                item->cause.u.nas = CauseNas(qosFlow->causeValue);
                break;                              
            case protocol_chosen:
                item->cause.choice = protocol_chosen;
                item->cause.u.protocol = CauseProtocol(qosFlow->causeValue);
                break;                              
            case misc_chosen:
                item->cause.choice = misc_chosen;
                item->cause.u.misc = CauseMisc(qosFlow->causeValue);
                break;
            default:
               break;//do nothing
        }
    }

    //common network instance
    if (cmmnNtwkInstancePrst_m) {
        CommonNetworkInstance *cmmnNtwrkInstancePtr = msgPdu_m.getCommonNetworkInstanceEncodePtr();
        cmmnNtwrkInstancePtr->length = cmmnNtwkInstance_m.length();
        memcpy(cmmnNtwrkInstancePtr->value,cmmnNtwkInstance_m.c_str(),cmmnNtwrkInstancePtr->length);
    }
    
    msgPdu_m.encodeOssToPer(ctxt);

    OssBuf *perBuf = msgPdu_m.getPerBufRef();

    msgBuffer.length = perBuf->length;
    msgBuffer.value = perBuf->value;

    return msgBuffer;
}

bool 
PduSesResModReqTransferCodec::decode(NgapOssCtxt *ctxt, MsgBuffer msgBuf)
{
    reset();
       
    OssBuf ossbuf;
    ossbuf.length = msgBuf.length;
    ossbuf.value = msgBuf.value;
        
    msgPdu_m.setPerBufRef_v(ossbuf);
        
    if (0 != msgPdu_m.decodePerToOss(ctxt)) {
          return false;
    }

    // session ambr
    if(msgPdu_m.isSessAmbrPrst()){
        isSessAMBRPrst_m = true;
        PDUSessionAggregateMaximumBitRate * ambrPtr = msgPdu_m.getSessAmbrDecodedPtr();
        sessAmbrULBitRate_m = ambrPtr->pDUSessionAggregateMaximumBitRateUL;
        sessAmbrDLBitRate_m = ambrPtr->pDUSessionAggregateMaximumBitRateDL;
    }
    
    // network instance
    if(msgPdu_m.isNetworkInstancePrst()){
        isNtwkInstancePrst_m = true;
        ntwkInstance_m = *(msgPdu_m.getNetworkInstanceDecodedPtr());
    }

    // qos flow add or modify request list
    for(int i=0; i< msgPdu_m.getNumofQosFlowAddOrModReqListItem(); i++) {
        QosFlowAddOrModifyRequestItem* item = msgPdu_m.getQosFlowAddOrModReqItemDecodedPtr(i);
        QosFlowAddOrModReqItem qosFlow;
    
        // qfi
        qosFlow.qosFlowInd = item->qosFlowIdentifier;
    
        // erab id
        if (item->bit_mask & QosFlowAddOrModifyRequestItem_e_RAB_ID_present){
            qosFlow.eRABId =item->e_RAB_ID;
            qosFlow.eRABIdPresent = true;
        }

        if (item->bit_mask & qosFlowLevelQosParameters_present){
            qosFlow.qosFlowParaPresent = true;

            // qos parameter - arp
            qosFlow.qosFlowLevQosPara.arp.priorityLevel 
                = item->qosFlowLevelQosParameters.allocationAndRetentionPriority.priorityLevelARP;
            qosFlow.qosFlowLevQosPara.arp.emptionCapability 
                = item->qosFlowLevelQosParameters.allocationAndRetentionPriority.pre_emptionCapability;
            qosFlow.qosFlowLevQosPara.arp.emptionVulnerability
                = item->qosFlowLevelQosParameters.allocationAndRetentionPriority.pre_emptionVulnerability;
    
            // qos parameter - qos characteristics
            if(item->qosFlowLevelQosParameters.qosCharacteristics.choice == nonDynamic5QI_chosen){
                qosFlow.qosFlowLevQosPara.qosCharacter.isDynamic = false;
                         
                qosFlow.qosFlowLevQosPara.qosCharacter.nonDynamic5qi.fiveQI
                    = item->qosFlowLevelQosParameters.qosCharacteristics.u.nonDynamic5QI.fiveQI;
    
                if(item->qosFlowLevelQosParameters.qosCharacteristics.u.nonDynamic5QI.bit_mask & priorityLevelQos_present){
                    qosFlow.qosFlowLevQosPara.qosCharacter.nonDynamic5qi.isPriorityLevelQosPrst = true;
                    qosFlow.qosFlowLevQosPara.qosCharacter.nonDynamic5qi.priorityLevelQos
                        = item->qosFlowLevelQosParameters.qosCharacteristics.u.nonDynamic5QI.priorityLevelQos;    
                }
    
                if(item->qosFlowLevelQosParameters.qosCharacteristics.u.nonDynamic5QI.bit_mask
                         & NonDynamic5QIDescriptor_averagingWindow_present){
                    qosFlow.qosFlowLevQosPara.qosCharacter.nonDynamic5qi.isAveWindowPrst = true;
                    qosFlow.qosFlowLevQosPara.qosCharacter.nonDynamic5qi.averagingWindow
                        = item->qosFlowLevelQosParameters.qosCharacteristics.u.nonDynamic5QI.averagingWindow;     
                }
    
                if(item->qosFlowLevelQosParameters.qosCharacteristics.u.nonDynamic5QI.bit_mask
                         & NonDynamic5QIDescriptor_maximumDataBurstVolume_present){
                    qosFlow.qosFlowLevQosPara.qosCharacter.nonDynamic5qi.isMaxDataBusrtVolPrst = true;
                    qosFlow.qosFlowLevQosPara.qosCharacter.nonDynamic5qi.maximumDataBurstVolume
                        = item->qosFlowLevelQosParameters.qosCharacteristics.u.nonDynamic5QI.maximumDataBurstVolume;      
                }
            } 
            else if(item->qosFlowLevelQosParameters.qosCharacteristics.choice == dynamic5QI_chosen){
                qosFlow.qosFlowLevQosPara.qosCharacter.isDynamic = true;
                qosFlow.qosFlowLevQosPara.qosCharacter.dynamic5qi.priorityLevelQos  
                    = item->qosFlowLevelQosParameters.qosCharacteristics.u.dynamic5QI.priorityLevelQos;
                qosFlow.qosFlowLevQosPara.qosCharacter.dynamic5qi.packetDelayBudget  
                    = item->qosFlowLevelQosParameters.qosCharacteristics.u.dynamic5QI.packetDelayBudget;
                qosFlow.qosFlowLevQosPara.qosCharacter.dynamic5qi.packetErrorRate.pERExponent  
                    = item->qosFlowLevelQosParameters.qosCharacteristics.u.dynamic5QI.packetErrorRate.pERExponent;                    
                qosFlow.qosFlowLevQosPara.qosCharacter.dynamic5qi.packetErrorRate.pERScalar  
                    = item->qosFlowLevelQosParameters.qosCharacteristics.u.dynamic5QI.packetErrorRate.pERScalar;
                if(item->qosFlowLevelQosParameters.qosCharacteristics.u.dynamic5QI.bit_mask & fiveQI_present){
                    qosFlow.qosFlowLevQosPara.qosCharacter.dynamic5qi.is5QIPrst = true;
                    qosFlow.qosFlowLevQosPara.qosCharacter.dynamic5qi.fiveQI 
                        = item->qosFlowLevelQosParameters.qosCharacteristics.u.dynamic5QI.fiveQI;
                }
                if(item->qosFlowLevelQosParameters.qosCharacteristics.u.dynamic5QI.bit_mask & delayCritical_present){
                    qosFlow.qosFlowLevQosPara.qosCharacter.dynamic5qi.isDelayCriticalPrst = true;
                    qosFlow.qosFlowLevQosPara.qosCharacter.dynamic5qi.delayCritical  
                        = item->qosFlowLevelQosParameters.qosCharacteristics.u.dynamic5QI.delayCritical;
                }
                if(item->qosFlowLevelQosParameters.qosCharacteristics.u.dynamic5QI.bit_mask
                    & Dynamic5QIDescriptor_averagingWindow_present){
                    qosFlow.qosFlowLevQosPara.qosCharacter.dynamic5qi.isAveWindowPrst = true;
                    qosFlow.qosFlowLevQosPara.qosCharacter.dynamic5qi.averagingWindow  
                        = item->qosFlowLevelQosParameters.qosCharacteristics.u.dynamic5QI.averagingWindow;
                }
                if(item->qosFlowLevelQosParameters.qosCharacteristics.u.dynamic5QI.bit_mask
                    & Dynamic5QIDescriptor_maximumDataBurstVolume_present){
                    qosFlow.qosFlowLevQosPara.qosCharacter.dynamic5qi.isMaxDataBusrtVolPrst = true;
                    qosFlow.qosFlowLevQosPara.qosCharacter.dynamic5qi.maximumDataBurstVolume  
                        = item->qosFlowLevelQosParameters.qosCharacteristics.u.dynamic5QI.maximumDataBurstVolume;
                }
            } 
            else{
                std::cout << "unsupported qos parameters." << endl;
            }
    
            // qos parameter - GBR_QosInformation
            if(item->qosFlowLevelQosParameters.bit_mask & gBR_QosInformation_present){
                qosFlow.qosFlowLevQosPara.gBRQosInfoPresent = true;
                qosFlow.qosFlowLevQosPara.gBRQosInfo.maxFlowBitRateUL  
                    = item->qosFlowLevelQosParameters.gBR_QosInformation.maximumFlowBitRateUL;
                qosFlow.qosFlowLevQosPara.gBRQosInfo.maxFlowBitRateDL  
                    = item->qosFlowLevelQosParameters.gBR_QosInformation.maximumFlowBitRateDL;
                qosFlow.qosFlowLevQosPara.gBRQosInfo.guaFlowBitRateUL  
                    = item->qosFlowLevelQosParameters.gBR_QosInformation.guaranteedFlowBitRateUL;
                qosFlow.qosFlowLevQosPara.gBRQosInfo.guaFlowBitRateDL  
                    = item->qosFlowLevelQosParameters.gBR_QosInformation.guaranteedFlowBitRateDL;
                if(item->qosFlowLevelQosParameters.gBR_QosInformation.bit_mask & notificationControl_present){
                    qosFlow.qosFlowLevQosPara.gBRQosInfo.notiControlPresent = true;
                    qosFlow.qosFlowLevQosPara.gBRQosInfo.notiControl  
                        = item->qosFlowLevelQosParameters.gBR_QosInformation.notificationControl;
                }
                if(item->qosFlowLevelQosParameters.gBR_QosInformation.bit_mask
                    & maximumPacketLossRateDL_present){
                    qosFlow.qosFlowLevQosPara.gBRQosInfo.maxPacketLossRateDLPresent= true;
                    qosFlow.qosFlowLevQosPara.gBRQosInfo.maxPacketLossRateDL  
                        = item->qosFlowLevelQosParameters.gBR_QosInformation.maximumPacketLossRateDL;
                }
                if(item->qosFlowLevelQosParameters.gBR_QosInformation.bit_mask
                    & maximumPacketLossRateUL_present){
                    qosFlow.qosFlowLevQosPara.gBRQosInfo.maxPacketLossRateULPresent= true;
                    qosFlow.qosFlowLevQosPara.gBRQosInfo.maxPacketLossRateUL  
                        = item->qosFlowLevelQosParameters.gBR_QosInformation.maximumPacketLossRateUL;
                }
            }
    
            // qos parameter - ReflectiveQosAttribute
            if(item->qosFlowLevelQosParameters.bit_mask & reflectiveQosAttribute_present){
                qosFlow.qosFlowLevQosPara.refQosAttrPresent = true;
                qosFlow.qosFlowLevQosPara.refQosAttr  
                    = item->qosFlowLevelQosParameters.reflectiveQosAttribute;
            }
    
            // qos parameter - addtional QosInformation
            if(item->qosFlowLevelQosParameters.bit_mask & additionalQosFlowInformation_present){
                qosFlow.qosFlowLevQosPara.addQosFlowInfoPresent = true;
                qosFlow.qosFlowLevQosPara.addQosFlowInfo
                    = item->qosFlowLevelQosParameters.additionalQosFlowInformation;
            }
        }

        //push back the qos flow 
        qosFlowAddOrModReqList_m.push_back(qosFlow);
     }


    // ul ngu up tnl modify list
    for(int i=0; i< msgPdu_m.getNumofUlNguUpTnlModifyListItem(); i++) {
        UL_NGU_UP_TNLModifyItem* item = msgPdu_m.getUlNguUpTnlModifyItemDecodedPtr(i);
        UlNguUpTnlModifyItem ulNguUpTnlModInfo;

        // uplink transportlayer information
        if (item->uL_NGU_UP_TNLInformation.choice != gTPTunnel_chosen){
            std::cout << "unsupported transport layer, only gtp tunnel supported."<< std::endl;
            continue;
        }
        unsigned short addrLen = item->uL_NGU_UP_TNLInformation.u.gTPTunnel.transportLayerAddress.length;
        if (addrLen == 32){
            ulNguUpTnlModInfo.uplinkGtpTunnel.ipType = ipv4;
            memcpy(ulNguUpTnlModInfo.uplinkGtpTunnel.transportLayerAddr, 
                   item->uL_NGU_UP_TNLInformation.u.gTPTunnel.transportLayerAddress.value, 4);
        } else if(addrLen == 128 ){
            ulNguUpTnlModInfo.uplinkGtpTunnel.ipType = ipv6;
            memcpy(ulNguUpTnlModInfo.uplinkGtpTunnel.transportLayerAddr, 
                   item->uL_NGU_UP_TNLInformation.u.gTPTunnel.transportLayerAddress.value, 16);
        }else{
            std::cout << "unsupported ip address" << std::endl;
            continue;
        }
        if (item->uL_NGU_UP_TNLInformation.u.gTPTunnel.gTP_TEID.length == 4){
           memcpy(ulNguUpTnlModInfo.uplinkGtpTunnel.gtpTeid, 
                  item->uL_NGU_UP_TNLInformation.u.gTPTunnel.gTP_TEID.value,sizeof(Uint32));
        }else{
            std::cout << "invalid gtp teid" << std::endl;
            continue;
        }

        //downlink
        if (item->dL_NGU_UP_TNLInformation.choice != gTPTunnel_chosen){
            std::cout << "unsupported transport layer, only gtp tunnel supported."<< std::endl;
            continue;
        }
        addrLen = item->dL_NGU_UP_TNLInformation.u.gTPTunnel.transportLayerAddress.length;
        if (addrLen == 32){
            ulNguUpTnlModInfo.downlinkGtpTunnel.ipType = ipv4;
            memcpy(ulNguUpTnlModInfo.downlinkGtpTunnel.transportLayerAddr, 
                   item->dL_NGU_UP_TNLInformation.u.gTPTunnel.transportLayerAddress.value, 4);
        } else if(addrLen == 128 ){
            ulNguUpTnlModInfo.downlinkGtpTunnel.ipType = ipv6;
            memcpy(ulNguUpTnlModInfo.downlinkGtpTunnel.transportLayerAddr, 
                   item->dL_NGU_UP_TNLInformation.u.gTPTunnel.transportLayerAddress.value, 16);
        }else{
            std::cout << "unsupported ip address" << std::endl;
            continue;
        }
        if (item->dL_NGU_UP_TNLInformation.u.gTPTunnel.gTP_TEID.length == 4){
           memcpy(ulNguUpTnlModInfo.downlinkGtpTunnel.gtpTeid, 
                  item->dL_NGU_UP_TNLInformation.u.gTPTunnel.gTP_TEID.value,sizeof(Uint32));
        }else{
            std::cout << "invalid gtp teid" << std::endl;
            continue;
        }

        ulNguUpTnlModifyList_m.push_back(ulNguUpTnlModInfo);
    }

    //additional ul ng-u up tnl information list
    for(int i=0; i< msgPdu_m.getNumOfAddUpTransportLayerInfoListItem(); i++) {
        UPTransportLayerInformationItem* item = msgPdu_m.getAddUpTransportLayerInfoItemDecodedPtr(i);
        GtpTunnel addGtpTunnelInfo;

        if (item->nGU_UP_TNLInformation.choice != gTPTunnel_chosen){
            std::cout << "unsupported transport layer, only gtp tunnel supported."<< std::endl;
            continue;
        }
        unsigned short addrLen = item->nGU_UP_TNLInformation.u.gTPTunnel.transportLayerAddress.length;
        if (addrLen == 32){
            addGtpTunnelInfo.ipType = ipv4;
            memcpy(addGtpTunnelInfo.transportLayerAddr, 
                   item->nGU_UP_TNLInformation.u.gTPTunnel.transportLayerAddress.value, 4);
        } else if(addrLen == 128 ){
            addGtpTunnelInfo.ipType = ipv6;
            memcpy(addGtpTunnelInfo.transportLayerAddr, 
                   item->nGU_UP_TNLInformation.u.gTPTunnel.transportLayerAddress.value, 16);
        }else{
            std::cout << "unsupported ip address" << std::endl;
            continue;
        }
        if (item->nGU_UP_TNLInformation.u.gTPTunnel.gTP_TEID.length == 4){
           memcpy(addGtpTunnelInfo.gtpTeid, 
                  item->nGU_UP_TNLInformation.u.gTPTunnel.gTP_TEID.value,sizeof(Uint32));
        }else{
            std::cout << "invalid gtp teid" << std::endl;
            continue;
        }
        addGtpTunnelList_m.push_back(addGtpTunnelInfo);
        addGtpTunnelListPrst_m = true;
    }


    // qos flow to release list
    for(int i=0; i< msgPdu_m.getNumofQosFlowToRelListItem(); i++) {
        QosFlowWithCauseItem* item = msgPdu_m.getQosFlowToRelItemDecodedPtr(i);
        QosFlowCodecItem qosFlow;
    
        // qfi
        qosFlow.qosFlowInd = item->qosFlowIdentifier;
        switch (item->cause.choice){
            case radioNetwork_chosen:
                qosFlow.causeType = radioNetwork_chosen;
                qosFlow.causeValue = item->cause.u.radioNetwork;
                break;
            case transport_chosen:
                qosFlow.causeType = transport_chosen;
                qosFlow.causeValue = item->cause.u.transport;
                break;                              
            case nas_chosen:
                qosFlow.causeType = nas_chosen;
                qosFlow.causeValue = item->cause.u.nas;
                break;                              
            case protocol_chosen:
                qosFlow.causeType = protocol_chosen;
                qosFlow.causeValue = item->cause.u.protocol;
                break;                              
            case misc_chosen:
                qosFlow.causeType = misc_chosen;
                qosFlow.causeValue = item->cause.u.misc;
                break;
            default:
                break;//do nothing
            }

       qosFlowToReleaseList_m.push_back(qosFlow);
    }

    //common network instance
    if (msgPdu_m.isCommonNetworkInstancePrst()) {
        CommonNetworkInstance *cmmnNtwrkInstancePtr = msgPdu_m.getCommonNetworkInstanceDecodedPtr();
        cmmnNtwkInstance_m.assign((char *)(cmmnNtwrkInstancePtr->value), cmmnNtwrkInstancePtr->length);
        cmmnNtwkInstancePrst_m = true;
    }
    
    return true;
}


std::string 
PduSesResModReqTransferCodec::toString(){
  ostringstream oss;
    oss << "CommonNetworkInstance:Prst(" << cmmnNtwkInstancePrst_m << ") value (" << cmmnNtwkInstance_m << ")" <<endl; 
//  oss << "PduSesResSetupReqTransfer: {"
//      << "SessAmbrPrst(" << isSessAMBRPrst_m << "),"
//      << "Uplink("<<sessAmbrULBitRate_m<<"),"
//      << "Downlink("<<sessAmbrDLBitRate_m << "),"
//      << "Up Tnl Information:{},"
//      << "TransportLayerAddress:"
//      << (int)upGtpTunnel_m.transportLayerAddr[0]
//      << (int)upGtpTunnel_m.transportLayerAddr[1]
//      << (int)upGtpTunnel_m.transportLayerAddr[2]
//      << (int)upGtpTunnel_m.transportLayerAddr[3]
//      << (int)upGtpTunnel_m.transportLayerAddr[4]
//      << (int)upGtpTunnel_m.transportLayerAddr[5]
//      << (int)upGtpTunnel_m.transportLayerAddr[6]
//      << (int)upGtpTunnel_m.transportLayerAddr[7]
//      << (int)upGtpTunnel_m.transportLayerAddr[8]
//      << (int)upGtpTunnel_m.transportLayerAddr[9]
//      << (int)upGtpTunnel_m.transportLayerAddr[10]
//      << (int)upGtpTunnel_m.transportLayerAddr[11]
//      << (int)upGtpTunnel_m.transportLayerAddr[12]
//      << (int)upGtpTunnel_m.transportLayerAddr[13]
//      << (int)upGtpTunnel_m.transportLayerAddr[14]
//      << (int)upGtpTunnel_m.transportLayerAddr[15]
//      << (int)upGtpTunnel_m.transportLayerAddr[16]
//      << (int)upGtpTunnel_m.transportLayerAddr[17]
//      << (int)upGtpTunnel_m.transportLayerAddr[18]
//      << (int)upGtpTunnel_m.transportLayerAddr[19]
//      << ", Teid:"
//      << (int)upGtpTunnel_m.gtpTeid[0]
//      << (int)upGtpTunnel_m.gtpTeid[1]
//      << (int)upGtpTunnel_m.gtpTeid[2]
//      << (int)upGtpTunnel_m.gtpTeid[3]
//      /*
//      << "Add Up Tnl Information:Prst " << addUpGtpTunnelPrst_m <<endl
//      << "addTransportLayerAddress:"
//      << (int)addUpGtpTunnel_m.transportLayerAddr[0]
//      << (int)addUpGtpTunnel_m.transportLayerAddr[1]
//      << (int)addUpGtpTunnel_m.transportLayerAddr[2]
//      << (int)addUpGtpTunnel_m.transportLayerAddr[3]
//      << (int)addUpGtpTunnel_m.transportLayerAddr[4]
//      << (int)addUpGtpTunnel_m.transportLayerAddr[5]
//      << (int)addUpGtpTunnel_m.transportLayerAddr[6]
//      << (int)addUpGtpTunnel_m.transportLayerAddr[7]
//      << (int)addUpGtpTunnel_m.transportLayerAddr[8]
//      << (int)addUpGtpTunnel_m.transportLayerAddr[9]
//      << (int)addUpGtpTunnel_m.transportLayerAddr[10]
//      << (int)addUpGtpTunnel_m.transportLayerAddr[11]
//      << (int)addUpGtpTunnel_m.transportLayerAddr[12]
//      << (int)addUpGtpTunnel_m.transportLayerAddr[13]
//      << (int)addUpGtpTunnel_m.transportLayerAddr[14]
//      << (int)addUpGtpTunnel_m.transportLayerAddr[15]
//      << (int)addUpGtpTunnel_m.transportLayerAddr[16]
//      << (int)addUpGtpTunnel_m.transportLayerAddr[17]
//      << (int)addUpGtpTunnel_m.transportLayerAddr[18]
//      << (int)addUpGtpTunnel_m.transportLayerAddr[19] <<endl
//      << "addTeid:"
//      << (int)addUpGtpTunnel_m.gtpTeid[0]
//      << (int)addUpGtpTunnel_m.gtpTeid[1]
//      << (int)addUpGtpTunnel_m.gtpTeid[2]
//      << (int)addUpGtpTunnel_m.gtpTeid[3] <<endl
//      */
//      << ", DatFwNotPssblePrst(" << isDataFwNotPossible_m << "),"
//      << "value("<<(int)dataFwNotPssble_m<<"),"
//      << "Pdu Session Type(" << (int)pduSessType_m << "),"
//      << "SecIndPrst(" << isSecInd_m << "),"
//      << "int("<<(int)intProtectInd_m<<"),"
//      << "confd("<<(int)confdProtecInd_m << "),"
//      << "maxPrtDataRate:Prst " <<isMaxPrtDataRate_m
//      << " value: ("<<(int)maxPrtDataRate_m << "),"
//      << "network Instance:Prst " << isNtwkInstancePrst_m
//      << " value: (" << ntwkInstance_m  << ")," << endl
//      << "QosFlowSetupReqList: length("<<qosFlowSetupReqList_m.size()<<")" << endl;
//  
//       vector<QosFlowSetupReqItem>::iterator qit = qosFlowSetupReqList_m.begin();
//           int index = 0;
//       for (; qit != qosFlowSetupReqList_m.end(); ++qit) {
//       oss << index<<":{ qosFlowInd("<< qit->qosFlowInd << ")," << "isDynamic:"
//           << qit->qosFlowLevQosPara.qosCharacter.isDynamic <<"; "
////           << "qosFlowLevelQosPara.QosCharacter.nonDynamic5QI.5QI("
////           << "pduSessResSetupReqTransfer:QosFlowSetupReqItem."
////              "qosFlowLevelQosPara.NonDynamic5QI.priLevelQos:Flag("
////           << qit->qosFlowLevQosPara.nonDyna5qi.plQosPrst << ")" << endl
////           << "value(" << qit->qosFlowLevQosPara.nonDyna5qi.priLevelQos << ")" << endl
////           << "pduSessResSetupReqTransfer:QosFlowSetupReqItem."
////              "qosFlowLevelQosPara.NonDynamic5QI.aveWin:Flag("
////           << qit->qosFlowLevQosPara.nonDyna5qi.aveWinPresent << ")" << endl
////           << "value(" << qit->qosFlowLevQosPara.nonDyna5qi.aveWin << ")" << endl
////           << "pduSessResSetupReqTransfer:QosFlowSetupReqItem."
////              "qosFlowLevelQosPara.NonDynamic5QI.maxDataBurstVol:Flag("
////           << qit->qosFlowLevQosPara.nonDyna5qi.maxDataBurstVolPresent << ")" << endl
////           << "value(" << qit->qosFlowLevQosPara.nonDyna5qi.maxDataBurstVol << ")" << endl
////
////           << "pduSessResSetupReqTransfer:qosFlowLevelQosPara.allocationAndRetentionPri:"<< endl
////           << "priLevelARP("
////           << (int)(qit->qosFlowLevQosPara.alloAndRetPri.priLevelARP) << ")" << endl
////
////           << "preEmptionCap("
////           << (int)(qit->qosFlowLevQosPara.alloAndRetPri.preEmptionCap) << ")" << endl
////
////           << "preEmptionVulabi("
////           << (int)(qit->qosFlowLevQosPara.alloAndRetPri.preEmptionVulabi) << ")" << endl
// 
//           << "gbrQosInfo:Flag("
//           << qit->qosFlowLevQosPara.gBRQosInfoPresent << ")" <<"; "
//           << "maxFlowBitRateDL("
//           << qit->qosFlowLevQosPara.gBRQosInfo.maxFlowBitRateDL << ")"<<"; "
//
//           << "maxFlowBitRateUL("
//           << qit->qosFlowLevQosPara.gBRQosInfo.maxFlowBitRateUL << ")" <<"; "
//
//           << "guaFlowBitRateDL("
//           << qit->qosFlowLevQosPara.gBRQosInfo.guaFlowBitRateDL << ")" <<"; "
//
//           << "guaFlowBitRateUL("
//           << qit->qosFlowLevQosPara.gBRQosInfo.guaFlowBitRateUL << ")"<<"; "
//
//           << "gbrQosInfo.notiControl:Flag("
//           << qit->qosFlowLevQosPara.gBRQosInfo.notiControlPresent << ") "
//           << "value(" << (int)(qit->qosFlowLevQosPara.gBRQosInfo.notiControl) << ")" <<"; "
//
//           << "gbrQosInfo.maxPacketLossRateDL:Flag("
//           << qit->qosFlowLevQosPara.gBRQosInfo.maxPacketLossRateDL << ") "
//           << "value(" << qit->qosFlowLevQosPara.gBRQosInfo.maxPacketLossRateDL << ")" <<"; "
//
//           << "gbrQosInfo.maxPacketLossRateUL:Flag("
//           << qit->qosFlowLevQosPara.gBRQosInfo.maxPacketLossRateUL << ") "
//           << "value(" << qit->qosFlowLevQosPara.gBRQosInfo.maxPacketLossRateUL << ")"<<"; "
//
//           << "refQosAttr:Flag("
//           << qit->qosFlowLevQosPara.refQosAttrPresent << ") "
//           << "value(" << (int)(qit->qosFlowLevQosPara.refQosAttr) << ")" <<"; "
//
//           << "addQosFlowInfo:Flag("
//           << qit->qosFlowLevQosPara.addQosFlowInfoPresent << ") "
//           << "value(" << (int)(qit->qosFlowLevQosPara.addQosFlowInfo) << ")" <<"; "
//           << "eRABId:Flag("
//           << qit->eRABIdPresent<< ") "
//           << "value(" << qit->eRABId << ") }" << endl;
//
//                   index++;
//     }
    return oss.str();
}


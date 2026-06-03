#include <iostream>
#include <sstream>
#include <string>

#include "pduSessionResourceSetupRequestTransferCodec.h"

using namespace std;

PduSesResSetupReqTransferCodec::PduSesResSetupReqTransferCodec() {
  reset();
}

PduSesResSetupReqTransferCodec::~PduSesResSetupReqTransferCodec() {}

bool 
PduSesResSetupReqTransferCodec::isSessAMBRPrst(){
    return isSessAMBRPrst_m;
}

void 
PduSesResSetupReqTransferCodec::setSessAmbr(long ulBitRate, long dlBitRate){
    sessAmbrULBitRate_m = ulBitRate;
    sessAmbrDLBitRate_m = dlBitRate;
    isSessAMBRPrst_m = true;
}

long 
PduSesResSetupReqTransferCodec::getSessAmbrUl(){
    return sessAmbrULBitRate_m;
}

long 
PduSesResSetupReqTransferCodec::getSessAmbrDl(){
    return sessAmbrDLBitRate_m;
}

bool
PduSesResSetupReqTransferCodec::isNtwkInstancePrst(){
    return isNtwkInstancePrst_m;
}

void 
PduSesResSetupReqTransferCodec::setNtwkInstance(Uint64 ntwkInstanceVal){
    ntwkInstance_m = ntwkInstanceVal;
    isNtwkInstancePrst_m = true;
}

Uint64 
PduSesResSetupReqTransferCodec::getNtwkInstance(){
    return ntwkInstance_m;
}

void 
PduSesResSetupReqTransferCodec::setUpGtpTunnel(GtpTunnelInfo& tunnel)
{
   upGtpTunnel_m = tunnel;
}

GtpTunnelInfo& 
PduSesResSetupReqTransferCodec::getUpGtpTunnel()
{
   return upGtpTunnel_m;
}

void
PduSesResSetupReqTransferCodec::addUpTransLayerInfoList(GtpTunnel& item)
{
    isAddTransLayerInfoListPrst_m = true;
    addUpTransportLayerInfoList_m.push_back(item);
}

AddUpTransportLayerInfoList&
PduSesResSetupReqTransferCodec::getUpTransLayerInfoList()
{
    return addUpTransportLayerInfoList_m;
}

bool
PduSesResSetupReqTransferCodec::isAddUpTransLayerInfoListPrst()
{
    return isAddTransLayerInfoListPrst_m;
}

bool 
PduSesResSetupReqTransferCodec::isDataFwNotPossiblePrst(){
    return isDataFwNotPossible_m;
}
void 
PduSesResSetupReqTransferCodec::setDataFwNotPssble(Uint8 value){
    isDataFwNotPossible_m = true;
    dataFwNotPssble_m = value;
}
Uint8 
PduSesResSetupReqTransferCodec::getDataFwNotPssble(){
    return dataFwNotPssble_m;
}

void 
PduSesResSetupReqTransferCodec::setPduSessType(Uint8 value){
    pduSessType_m = value;
}
Uint8 PduSesResSetupReqTransferCodec::getPduSessType(){
    return pduSessType_m;
}

bool PduSesResSetupReqTransferCodec::isSecIndPrst(){
    return isSecInd_m;
}
void 
PduSesResSetupReqTransferCodec::setSecInd(Uint8 intInd, Uint8 confdInd){
    intProtectInd_m = intInd;
    confdProtecInd_m = confdInd;
    isSecInd_m = true;
}
Uint8 
PduSesResSetupReqTransferCodec::getSecIndProtectInd(){
    return intProtectInd_m;
}
Uint8 
PduSesResSetupReqTransferCodec::getSecConfdProtectInd(){
    return confdProtecInd_m;
}

bool 
PduSesResSetupReqTransferCodec::isMaxPrtDataRatePrst()
{
    return isMaxPrtDataRate_m;
}

void 
PduSesResSetupReqTransferCodec::setMaxPrtDataRate(Uint8 maxPrtData)
{
    maxPrtDataRate_m = maxPrtData;
    isMaxPrtDataRate_m = true;
}

Uint8 
PduSesResSetupReqTransferCodec::getMaxPrtDataRate()
{
    return maxPrtDataRate_m;
}

bool 
PduSesResSetupReqTransferCodec::isMaxPrtDataRateDlPrst()
{
    return maxIntProDataRateDownlinkPrst_m;
}

void 
PduSesResSetupReqTransferCodec::setMaxPrtDataRateDl(Uint8 maxPrtDataDl)
{
    maxIntProDataRateDownlink_m = maxPrtDataDl;
    maxIntProDataRateDownlinkPrst_m = true;
}

Uint8 
PduSesResSetupReqTransferCodec::getMaxPrtDataRateDl()
{
    return maxIntProDataRateDownlink_m;
}

void 
PduSesResSetupReqTransferCodec::addQosFlowSetupReqList(QosFlowSetupReqItem& item){
    qosFlowSetupReqList_m.push_back(item);
}
QosFlowSetupReqList& 
PduSesResSetupReqTransferCodec::getQosFlowSetupReqList(){
    return qosFlowSetupReqList_m;
}

void
PduSesResSetupReqTransferCodec::setCommonNetworkInstance(std::string instance)
{
    cmmnNtwrkInstancePrst_m = true;
    cmmnNtwrkInstance_m = instance;

}

std::string
PduSesResSetupReqTransferCodec::getCommonNetworkInstance()
{
    return cmmnNtwrkInstance_m;
}

bool
PduSesResSetupReqTransferCodec::isCommonNetworkInstancePrst()
{
    return cmmnNtwrkInstancePrst_m;
}


bool 
PduSesResSetupReqTransferCodec::reset(){
    msgPdu_m.reset_v();
    isSessAMBRPrst_m = false;
    sessAmbrULBitRate_m = 0;
    sessAmbrDLBitRate_m = 0;

    upGtpTunnel_m.ipType = ipv4;  // ipv4 address
    memset(upGtpTunnel_m.transportLayerAddr, 0, sizeof(upGtpTunnel_m.transportLayerAddr));
    memset(upGtpTunnel_m.gtpTeid,0,sizeof(upGtpTunnel_m.gtpTeid));

    isAddTransLayerInfoListPrst_m = false;
    addUpTransportLayerInfoList_m.clear();

    isDataFwNotPossible_m = false;
    dataFwNotPssble_m = 0;
  
    pduSessType_m = 0;

    isSecInd_m = false;
    intProtectInd_m = 0;
    confdProtecInd_m = 0;
    isMaxPrtDataRate_m =false;
    maxPrtDataRate_m = 0;
    maxIntProDataRateDownlinkPrst_m = false;
    maxIntProDataRateDownlink_m = 0;

    ntwkInstance_m = 0;
    isNtwkInstancePrst_m = false;

   qosFlowSetupReqList_m.clear();
}

//bool PduSesResSetupReqTransferCodec::encode(NgapOssCtxt* ctxt, MsgBuffer msgBuffer)
MsgBuffer PduSesResSetupReqTransferCodec::encode(NgapOssCtxt* ctxt)
{
    MsgBuffer msgBuffer;
    msgBuffer.length = 0;
    msgBuffer.value = NULL;

    // pdu session type, TODO, 4 stand for ?
    if (pduSessType_m > 4) {
        std::cout << "pdu session type is invalid" <<pduSessType_m<<std::endl;
        return msgBuffer;
    }
    *msgPdu_m.getPduSessTypeEncodePtr() = (PDUSessionType)pduSessType_m;

    // up transportlayer information
    UPTransportLayerInformation * upTransInfo = msgPdu_m.getUpTransLayerInfoEncodePtr();
    if(upGtpTunnel_m.ipType == ipv4) {
        upTransInfo->u.gTPTunnel.transportLayerAddress.length = 32; // 32 bit for ipv4
        memcpy(upTransInfo->u.gTPTunnel.transportLayerAddress.value, upGtpTunnel_m.transportLayerAddr, 4); //for ipv4
    }else if(upGtpTunnel_m.ipType == ipv6){
        upTransInfo->u.gTPTunnel.transportLayerAddress.length = 128; // 16*8=128 bit for ipv6
        memcpy(upTransInfo->u.gTPTunnel.transportLayerAddress.value, upGtpTunnel_m.transportLayerAddr, 16); //for ipv6
    }else{
        std::cout << "unsupported ip address" << std::endl;
    }

    upTransInfo->u.gTPTunnel.gTP_TEID.length = sizeof(upGtpTunnel_m.gtpTeid);
    memcpy(upTransInfo->u.gTPTunnel.gTP_TEID.value, upGtpTunnel_m.gtpTeid, sizeof(upGtpTunnel_m.gtpTeid));

    //Additional Ul Ng-U Up TNL Information
    if (isAddTransLayerInfoListPrst_m) {
        unsigned int num = addUpTransportLayerInfoList_m.size();
        msgPdu_m.setNumofUpTransLayerInfoListItem_m_v(num);
        for (int i = 0;i < num ; i++) {
            UPTransportLayerInformationItem* item = msgPdu_m.getUpTransLayerInfoItemEncodePtr(i);
            GtpTunnelInfo* addGtpTunnelInfo = &addUpTransportLayerInfoList_m[i];

            if(addGtpTunnelInfo->ipType == ipv4) {
                item->nGU_UP_TNLInformation.u.gTPTunnel.transportLayerAddress.length = 32; // 32 bit for ipv4
                memcpy(item->nGU_UP_TNLInformation.u.gTPTunnel.transportLayerAddress.value, addGtpTunnelInfo->transportLayerAddr, 4); //for ipv4
            }
            else if (addGtpTunnelInfo->ipType == ipv6) {
                item->nGU_UP_TNLInformation.u.gTPTunnel.transportLayerAddress.length = 128; // 16*8=128 bit for ipv6
                memcpy(item->nGU_UP_TNLInformation.u.gTPTunnel.transportLayerAddress.value, addGtpTunnelInfo->transportLayerAddr, 16); //for ipv6
            }
            else {
                std::cout << "unsupported ip address" << std::endl;
            }
            item->nGU_UP_TNLInformation.u.gTPTunnel.gTP_TEID.length = sizeof(addGtpTunnelInfo->gtpTeid);
            memcpy(item->nGU_UP_TNLInformation.u.gTPTunnel.gTP_TEID.value, addGtpTunnelInfo->gtpTeid, sizeof(addGtpTunnelInfo->gtpTeid));
        }
    }


    // qos flow setup request list
    unsigned int num = qosFlowSetupReqList_m.size();
    msgPdu_m.setNumofQosFlowSetupReqListItem_m_v(num);
    for (int i=0; i < num; i++){
        QosFlowSetupRequestItem* item = msgPdu_m.getQosFlowSetupReqItemEncodePtr(i);
        QosFlowSetupReqItem* qosFlow = &qosFlowSetupReqList_m[i];

        // qfi
        item->qosFlowIdentifier = qosFlow->qosFlowInd;

        // qos flow parameter - qos characteristic
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
        item->qosFlowLevelQosParameters.allocationAndRetentionPriority.priorityLevelARP =
            qosFlow->qosFlowLevQosPara.arp.priorityLevel;
        item->qosFlowLevelQosParameters.allocationAndRetentionPriority.pre_emptionCapability =
            (Pre_emptionCapability)(qosFlow->qosFlowLevQosPara.arp.emptionCapability);
        item->qosFlowLevelQosParameters.allocationAndRetentionPriority.pre_emptionVulnerability =
            (Pre_emptionVulnerability)(qosFlow->qosFlowLevQosPara.arp.emptionVulnerability);

       // qos flow parameter - gbr qos info
        if(qosFlow->qosFlowLevQosPara.gBRQosInfoPresent){
            item->qosFlowLevelQosParameters.bit_mask
                |= gBR_QosInformation_present;

            item->qosFlowLevelQosParameters.gBR_QosInformation.maximumFlowBitRateDL =
                qosFlow->qosFlowLevQosPara.gBRQosInfo.maxFlowBitRateDL;
            item->qosFlowLevelQosParameters.gBR_QosInformation.maximumFlowBitRateUL =
                qosFlow->qosFlowLevQosPara.gBRQosInfo.maxFlowBitRateUL;
            item->qosFlowLevelQosParameters.gBR_QosInformation.guaranteedFlowBitRateDL =
                qosFlow->qosFlowLevQosPara.gBRQosInfo.guaFlowBitRateDL;
            item->qosFlowLevelQosParameters.gBR_QosInformation.guaranteedFlowBitRateUL =
                qosFlow->qosFlowLevQosPara.gBRQosInfo.guaFlowBitRateUL;

            if(qosFlow->qosFlowLevQosPara.gBRQosInfo.notiControlPresent){
                item->qosFlowLevelQosParameters.gBR_QosInformation.bit_mask
                    |= notificationControl_present;
                item->qosFlowLevelQosParameters.gBR_QosInformation.notificationControl =
                    (NotificationControl)(qosFlow->qosFlowLevQosPara.gBRQosInfo.notiControl);
            }
            if(qosFlow->qosFlowLevQosPara.gBRQosInfo.maxPacketLossRateDLPresent){
                item->qosFlowLevelQosParameters.gBR_QosInformation.bit_mask
                    |= maximumPacketLossRateDL_present;
                item->qosFlowLevelQosParameters.gBR_QosInformation.maximumPacketLossRateDL =
                    qosFlow->qosFlowLevQosPara.gBRQosInfo.maxPacketLossRateDL;
            }
            if(qosFlow->qosFlowLevQosPara.gBRQosInfo.maxPacketLossRateULPresent){
                item->qosFlowLevelQosParameters.gBR_QosInformation.bit_mask
                    |= maximumPacketLossRateUL_present;
                item->qosFlowLevelQosParameters.gBR_QosInformation.maximumPacketLossRateUL =
                    qosFlow->qosFlowLevQosPara.gBRQosInfo.maxPacketLossRateUL;
            }
        }

        // qos flow parameter - refQosAttrPresent
        if(qosFlow->qosFlowLevQosPara.refQosAttrPresent){
            item->qosFlowLevelQosParameters.bit_mask
                |= reflectiveQosAttribute_present;
            item->qosFlowLevelQosParameters.reflectiveQosAttribute =
                (ReflectiveQosAttribute)(qosFlow->qosFlowLevQosPara.refQosAttr);
        }

        // qos flow parameter - addQosFlowInfo
        if(qosFlow->qosFlowLevQosPara.addQosFlowInfoPresent){
            item->qosFlowLevelQosParameters.bit_mask
                |= additionalQosFlowInformation_present;
            item->qosFlowLevelQosParameters.additionalQosFlowInformation =
                (AdditionalQosFlowInformation)(qosFlow->qosFlowLevQosPara.addQosFlowInfo);
        }

        // qos flow e-rab id
        if(qosFlow->eRABIdPresent){
            item->e_RAB_ID = (long)qosFlow->eRABId;
            item->bit_mask |= QosFlowSetupRequestItem_e_RAB_ID_present;
        }

    }

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

    // data forwarding
    if(isDataFwNotPossible_m){
        *msgPdu_m.getDataFwdNotPssbleEncodePtr() = (DataForwardingNotPossible)dataFwNotPssble_m;
    }

    // security indication
    if (isSecInd_m){
        msgPdu_m.getSecIndEncodePtr()->confidentialityProtectionIndication
            = (ConfidentialityProtectionIndication)confdProtecInd_m;
        msgPdu_m.getSecIndEncodePtr()->integrityProtectionIndication
            = (IntegrityProtectionIndication)intProtectInd_m;
        if(isMaxPrtDataRate_m)
        {
            msgPdu_m.getSecIndEncodePtr()->bit_mask = maximumIntegrityProtectedDataRate_UL_present;
            msgPdu_m.getSecIndEncodePtr()->maximumIntegrityProtectedDataRate_UL =
                (MaximumIntegrityProtectedDataRate)maxPrtDataRate_m;
        }
        if (maxIntProDataRateDownlinkPrst_m)
        {
            MaximumIntegrityProtectedDataRate* maxDataRateDlPtr = msgPdu_m.getMaxIntProDataRateEncodePtr();
            msgPdu_m.getSecIndEncodePtr()->bit_mask = SecurityIndication_iE_Extensions_present;
            *maxDataRateDlPtr = MaximumIntegrityProtectedDataRate(maxIntProDataRateDownlink_m);
        }
    }

    //common network instance
    if (cmmnNtwrkInstancePrst_m) {
      CommonNetworkInstance *commonNetworkInstancePtr = msgPdu_m.getCommonNetworkInstanceEncodePtr();
      commonNetworkInstancePtr->length = cmmnNtwrkInstance_m.length();
      memcpy(commonNetworkInstancePtr->value, cmmnNtwrkInstance_m.c_str(), commonNetworkInstancePtr->length);
    }

    msgPdu_m.encodeOssToPer(ctxt);

    OssBuf *perBuf = msgPdu_m.getPerBufRef();

    msgBuffer.length = perBuf->length;
    msgBuffer.value = perBuf->value;

    return msgBuffer;
}

bool 
PduSesResSetupReqTransferCodec::decode(NgapOssCtxt *ctxt, MsgBuffer msgBuf){

    reset();

    OssBuf ossbuf;
    ossbuf.length = msgBuf.length;
    ossbuf.value = msgBuf.value;

    msgPdu_m.setPerBufRef_v(ossbuf);

    if (0 != msgPdu_m.decodePerToOss(ctxt)) {
      return false;
    }

    // decode pdu session type
    pduSessType_m = *(msgPdu_m.getPduSessTypeDecodedPtr());

    //network instance
    if(msgPdu_m.isNetworkInstancePrst())
    {
        ntwkInstance_m = *(msgPdu_m.getNetworkInstanceDecodedPtr());
        isNtwkInstancePrst_m = true;
    }

    // up transportlayer information
    UPTransportLayerInformation * transLayerinfo = msgPdu_m.getUpTransLayerInfoDecodedPtr();
    if (transLayerinfo->choice != gTPTunnel_chosen){
        std::cout << "unsupported transport layer, only gtp tunnel supported."<< endl;
        return false;
    }

    unsigned short addrLen = transLayerinfo->u.gTPTunnel.transportLayerAddress.length;
    if (addrLen == 32){
        upGtpTunnel_m.ipType = ipv4;
        memcpy(upGtpTunnel_m.transportLayerAddr, transLayerinfo->u.gTPTunnel.transportLayerAddress.value, 4);
    } else if(addrLen == 128 ){
        upGtpTunnel_m.ipType = ipv6;
        memcpy(upGtpTunnel_m.transportLayerAddr, transLayerinfo->u.gTPTunnel.transportLayerAddress.value, 16);
    }else{
        std::cout << "unsupported ip address" << std::endl;
        return false;
    }
    if (transLayerinfo->u.gTPTunnel.gTP_TEID.length == 4){
        memcpy(upGtpTunnel_m.gtpTeid, transLayerinfo->u.gTPTunnel.gTP_TEID.value,sizeof(Uint32));
    }else{
        std::cout << "invalid gtp teid" << std::endl;
        return false;
    }

    // add up transportlayer information list
    if(msgPdu_m.isUpTransLayerInfoListPrst()){
        for (int i = 0;i < msgPdu_m.getNumofUpTransLayerInfoListItem(); i++) {
            UPTransportLayerInformationItem* item = msgPdu_m.getUpTransLayerInfoItemDecodedPtr(i);
            GtpTunnel addGtpTunnel;

            if (item->nGU_UP_TNLInformation.choice != gTPTunnel_chosen) {
                std::cout << "unsupported transport layer, only gtp tunnel supported." << endl;
                return false;
            }
            unsigned short addrItemLen = item->nGU_UP_TNLInformation.u.gTPTunnel.transportLayerAddress.length;
            if (addrItemLen == 32) {
                addGtpTunnel.ipType = ipv4;
                memcpy(addGtpTunnel.transportLayerAddr, item->nGU_UP_TNLInformation.u.gTPTunnel.transportLayerAddress.value, 4);
            }
            else if (addrItemLen == 128) {
                addGtpTunnel.ipType = ipv6;
                memcpy(addGtpTunnel.transportLayerAddr, item->nGU_UP_TNLInformation.u.gTPTunnel.transportLayerAddress.value, 16);
            }
            else {
                std::cout << "unsupported ip address" << std::endl;
                return false;
            }
            if (item->nGU_UP_TNLInformation.u.gTPTunnel.gTP_TEID.length == 4) {
                memcpy(addGtpTunnel.gtpTeid, item->nGU_UP_TNLInformation.u.gTPTunnel.gTP_TEID.value,sizeof(Uint32));
            }
            else
            {
                std::cout << "invalid gtp teid" << std::endl;
                return false;
            }
            //push back the qos flow
            addUpTransportLayerInfoList_m.push_back(addGtpTunnel);
        }
        isAddTransLayerInfoListPrst_m = true;
    }

    // qos flow list
    for(int i=0; i<msgPdu_m.getNumofQosFlowSetupReqListItem(); i++){
        QosFlowSetupRequestItem* item = msgPdu_m.getQosFlowSetupReqItemDecodedPtr(i);
        QosFlowSetupReqItem qosFlow;

        // qfi
        qosFlow.qosFlowInd = item->qosFlowIdentifier;

        // erab id
        if (item->bit_mask & QosFlowSetupRequestItem_e_RAB_ID_present){
            qosFlow.eRABId =item->e_RAB_ID;
            qosFlow.eRABIdPresent = true;
        }

        // qos parameter - arp
        qosFlow.qosFlowLevQosPara.arp.priorityLevel =
            item->qosFlowLevelQosParameters.allocationAndRetentionPriority.priorityLevelARP;
        qosFlow.qosFlowLevQosPara.arp.emptionCapability =
            item->qosFlowLevelQosParameters.allocationAndRetentionPriority.pre_emptionCapability;
        qosFlow.qosFlowLevQosPara.arp.emptionVulnerability =
            item->qosFlowLevelQosParameters.allocationAndRetentionPriority.pre_emptionVulnerability;

        // qos parameter - qos characteristics
        if(item->qosFlowLevelQosParameters.qosCharacteristics.choice == nonDynamic5QI_chosen){
            qosFlow.qosFlowLevQosPara.qosCharacter.isDynamic = false;

            qosFlow.qosFlowLevQosPara.qosCharacter.nonDynamic5qi.fiveQI =
                item->qosFlowLevelQosParameters.qosCharacteristics.u.nonDynamic5QI.fiveQI;

            if(item->qosFlowLevelQosParameters.qosCharacteristics.u.nonDynamic5QI.bit_mask
                & priorityLevelQos_present){
                qosFlow.qosFlowLevQosPara.qosCharacter.nonDynamic5qi.isPriorityLevelQosPrst = true;
                qosFlow.qosFlowLevQosPara.qosCharacter.nonDynamic5qi.priorityLevelQos =
                item->qosFlowLevelQosParameters.qosCharacteristics.u.nonDynamic5QI.priorityLevelQos;    
            }

            if(item->qosFlowLevelQosParameters.qosCharacteristics.u.nonDynamic5QI.bit_mask
                & NonDynamic5QIDescriptor_averagingWindow_present){
                qosFlow.qosFlowLevQosPara.qosCharacter.nonDynamic5qi.isAveWindowPrst = true;
                qosFlow.qosFlowLevQosPara.qosCharacter.nonDynamic5qi.averagingWindow =
                item->qosFlowLevelQosParameters.qosCharacteristics.u.nonDynamic5QI.averagingWindow; 
            }

            if(item->qosFlowLevelQosParameters.qosCharacteristics.u.nonDynamic5QI.bit_mask
                & NonDynamic5QIDescriptor_maximumDataBurstVolume_present){
                qosFlow.qosFlowLevQosPara.qosCharacter.nonDynamic5qi.isMaxDataBusrtVolPrst = true;
                qosFlow.qosFlowLevQosPara.qosCharacter.nonDynamic5qi.maximumDataBurstVolume =
                item->qosFlowLevelQosParameters.qosCharacteristics.u.nonDynamic5QI.maximumDataBurstVolume;  
            }
        }
        else if(item->qosFlowLevelQosParameters.qosCharacteristics.choice == dynamic5QI_chosen){
            qosFlow.qosFlowLevQosPara.qosCharacter.isDynamic = true;

            qosFlow.qosFlowLevQosPara.qosCharacter.dynamic5qi.priorityLevelQos =
                item->qosFlowLevelQosParameters.qosCharacteristics.u.dynamic5QI.priorityLevelQos;

            qosFlow.qosFlowLevQosPara.qosCharacter.dynamic5qi.packetDelayBudget =
                item->qosFlowLevelQosParameters.qosCharacteristics.u.dynamic5QI.packetDelayBudget;

            qosFlow.qosFlowLevQosPara.qosCharacter.dynamic5qi.packetErrorRate.pERExponent =
                item->qosFlowLevelQosParameters.qosCharacteristics.u.dynamic5QI.packetErrorRate.pERExponent;            

            qosFlow.qosFlowLevQosPara.qosCharacter.dynamic5qi.packetErrorRate.pERScalar =
                item->qosFlowLevelQosParameters.qosCharacteristics.u.dynamic5QI.packetErrorRate.pERScalar;

            if(item->qosFlowLevelQosParameters.qosCharacteristics.u.dynamic5QI.bit_mask
                & fiveQI_present){
                qosFlow.qosFlowLevQosPara.qosCharacter.dynamic5qi.is5QIPrst = true;
                qosFlow.qosFlowLevQosPara.qosCharacter.dynamic5qi.fiveQI =
                    item->qosFlowLevelQosParameters.qosCharacteristics.u.dynamic5QI.fiveQI;
            }

            if(item->qosFlowLevelQosParameters.qosCharacteristics.u.dynamic5QI.bit_mask
                & delayCritical_present){
                qosFlow.qosFlowLevQosPara.qosCharacter.dynamic5qi.isDelayCriticalPrst = true;
                qosFlow.qosFlowLevQosPara.qosCharacter.dynamic5qi.delayCritical =
                    item->qosFlowLevelQosParameters.qosCharacteristics.u.dynamic5QI.delayCritical;
            }

            if(item->qosFlowLevelQosParameters.qosCharacteristics.u.dynamic5QI.bit_mask
                & Dynamic5QIDescriptor_averagingWindow_present){
                qosFlow.qosFlowLevQosPara.qosCharacter.dynamic5qi.isAveWindowPrst = true;
                qosFlow.qosFlowLevQosPara.qosCharacter.dynamic5qi.averagingWindow =
                    item->qosFlowLevelQosParameters.qosCharacteristics.u.dynamic5QI.averagingWindow;
            }

            if(item->qosFlowLevelQosParameters.qosCharacteristics.u.dynamic5QI.bit_mask
                & Dynamic5QIDescriptor_maximumDataBurstVolume_present){
                qosFlow.qosFlowLevQosPara.qosCharacter.dynamic5qi.isMaxDataBusrtVolPrst = true;
                qosFlow.qosFlowLevQosPara.qosCharacter.dynamic5qi.maximumDataBurstVolume =
                    item->qosFlowLevelQosParameters.qosCharacteristics.u.dynamic5QI.maximumDataBurstVolume;
            }

        }
        else{
            std::cout << "unsupported qos parameters." << endl;
        }

        // qos parameter - GBR_QosInformation
        if(item->qosFlowLevelQosParameters.bit_mask 
            & gBR_QosInformation_present){
            qosFlow.qosFlowLevQosPara.gBRQosInfoPresent = true;

            qosFlow.qosFlowLevQosPara.gBRQosInfo.maxFlowBitRateUL =
                item->qosFlowLevelQosParameters.gBR_QosInformation.maximumFlowBitRateUL;
            qosFlow.qosFlowLevQosPara.gBRQosInfo.maxFlowBitRateDL =
                item->qosFlowLevelQosParameters.gBR_QosInformation.maximumFlowBitRateDL;
            qosFlow.qosFlowLevQosPara.gBRQosInfo.guaFlowBitRateUL =
                item->qosFlowLevelQosParameters.gBR_QosInformation.guaranteedFlowBitRateUL;
            qosFlow.qosFlowLevQosPara.gBRQosInfo.guaFlowBitRateDL =
                item->qosFlowLevelQosParameters.gBR_QosInformation.guaranteedFlowBitRateDL;

            if(item->qosFlowLevelQosParameters.gBR_QosInformation.bit_mask
                & notificationControl_present){
                qosFlow.qosFlowLevQosPara.gBRQosInfo.notiControlPresent = true;
                qosFlow.qosFlowLevQosPara.gBRQosInfo.notiControl =
                    item->qosFlowLevelQosParameters.gBR_QosInformation.notificationControl;
            }
            if(item->qosFlowLevelQosParameters.gBR_QosInformation.bit_mask
                & maximumPacketLossRateDL_present){
                qosFlow.qosFlowLevQosPara.gBRQosInfo.maxPacketLossRateDLPresent= true;
                qosFlow.qosFlowLevQosPara.gBRQosInfo.maxPacketLossRateDL =
                    item->qosFlowLevelQosParameters.gBR_QosInformation.maximumPacketLossRateDL;
            }
            if(item->qosFlowLevelQosParameters.gBR_QosInformation.bit_mask
                & maximumPacketLossRateUL_present){
                qosFlow.qosFlowLevQosPara.gBRQosInfo.maxPacketLossRateULPresent= true;
                qosFlow.qosFlowLevQosPara.gBRQosInfo.maxPacketLossRateUL =
                    item->qosFlowLevelQosParameters.gBR_QosInformation.maximumPacketLossRateUL;
            }
        }

        // qos parameter - ReflectiveQosAttribute
        if(item->qosFlowLevelQosParameters.bit_mask 
            & reflectiveQosAttribute_present){
            qosFlow.qosFlowLevQosPara.refQosAttrPresent = true;
            qosFlow.qosFlowLevQosPara.refQosAttr = 
                item->qosFlowLevelQosParameters.reflectiveQosAttribute;
        }

        // qos parameter - addtional QosInformation
        if(item->qosFlowLevelQosParameters.bit_mask 
            & additionalQosFlowInformation_present){
            qosFlow.qosFlowLevQosPara.addQosFlowInfoPresent = true;
            qosFlow.qosFlowLevQosPara.addQosFlowInfo = 
                item->qosFlowLevelQosParameters.additionalQosFlowInformation;
        }

        //push back the qos flow 
        qosFlowSetupReqList_m.push_back(qosFlow);
    }

    // session ambr
    if(msgPdu_m.isSessAmbrPrst()){
        isSessAMBRPrst_m = true;
        PDUSessionAggregateMaximumBitRate * ambrPtr = msgPdu_m.getSessAmbrDecodedPtr();
        sessAmbrULBitRate_m = ambrPtr->pDUSessionAggregateMaximumBitRateUL;
        sessAmbrDLBitRate_m = ambrPtr->pDUSessionAggregateMaximumBitRateDL;
    }

    // DataFwdNotPssblePrst
    if(msgPdu_m.isDataFwdNotPssblePrst()){
        dataFwNotPssble_m = *(msgPdu_m.getDataFwdNotPssbleDecodedPtr());
        isDataFwNotPossible_m = true;
    }

    // security indication
    if(msgPdu_m.isSecIndPrst()){
        isSecInd_m = true;
        intProtectInd_m = msgPdu_m.getSecIndDecodedPtr()->integrityProtectionIndication;
        confdProtecInd_m = msgPdu_m.getSecIndDecodedPtr()->confidentialityProtectionIndication;
        if(msgPdu_m.getSecIndDecodedPtr()->bit_mask & maximumIntegrityProtectedDataRate_UL_present)
        {
            isMaxPrtDataRate_m = true;
            maxPrtDataRate_m = msgPdu_m.getSecIndDecodedPtr()->maximumIntegrityProtectedDataRate_UL;
        }
        if(msgPdu_m.getSecIndDecodedPtr()->bit_mask & SecurityIndication_iE_Extensions_present)
        {
            maxIntProDataRateDownlinkPrst_m = true;
            SecurityIndicationIeExtension* secIndPtr = msgPdu_m.getSecIndIeExtenDecodedPtr();
            maxIntProDataRateDownlink_m = *(secIndPtr->value.extensionValue.decoded.pdu_MaximumIntegrityProtectedDataRate);
        }
    }

    // common network instance
    if(msgPdu_m.isCommonNetworkInstancePrst()){
        CommonNetworkInstance *cmmnNtwrkInstancePtr = msgPdu_m.getCommonNetworkInstanceDecodedPtr();
        cmmnNtwrkInstance_m.assign((char *)(cmmnNtwrkInstancePtr->value), cmmnNtwrkInstancePtr->length);
        cmmnNtwrkInstancePrst_m = true;
    }
    
    return true;
}


std::string 
PduSesResSetupReqTransferCodec::toString(){
  ostringstream oss;

  oss << "PduSesResSetupReqTransfer: {"
      << "SessAmbrPrst(" << isSessAMBRPrst_m << "),"
      << "Uplink("<<sessAmbrULBitRate_m<<"),"
      << "Downlink("<<sessAmbrDLBitRate_m << "),"
      << "Up Tnl Information:{},"
      << "TransportLayerAddress:"
      << (int)upGtpTunnel_m.transportLayerAddr[0]
      << (int)upGtpTunnel_m.transportLayerAddr[1]
      << (int)upGtpTunnel_m.transportLayerAddr[2]
      << (int)upGtpTunnel_m.transportLayerAddr[3]
      << (int)upGtpTunnel_m.transportLayerAddr[4]
      << (int)upGtpTunnel_m.transportLayerAddr[5]
      << (int)upGtpTunnel_m.transportLayerAddr[6]
      << (int)upGtpTunnel_m.transportLayerAddr[7]
      << (int)upGtpTunnel_m.transportLayerAddr[8]
      << (int)upGtpTunnel_m.transportLayerAddr[9]
      << (int)upGtpTunnel_m.transportLayerAddr[10]
      << (int)upGtpTunnel_m.transportLayerAddr[11]
      << (int)upGtpTunnel_m.transportLayerAddr[12]
      << (int)upGtpTunnel_m.transportLayerAddr[13]
      << (int)upGtpTunnel_m.transportLayerAddr[14]
      << (int)upGtpTunnel_m.transportLayerAddr[15]
      << (int)upGtpTunnel_m.transportLayerAddr[16]
      << (int)upGtpTunnel_m.transportLayerAddr[17]
      << (int)upGtpTunnel_m.transportLayerAddr[18]
      << (int)upGtpTunnel_m.transportLayerAddr[19]
      << ", Teid:"
      << (int)upGtpTunnel_m.gtpTeid[0]
      << (int)upGtpTunnel_m.gtpTeid[1]
      << (int)upGtpTunnel_m.gtpTeid[2]
      << (int)upGtpTunnel_m.gtpTeid[3] << endl
      << "Add Up Tnl Information List:Flag(" << isAddTransLayerInfoListPrst_m << ")" <<endl;

      vector<GtpTunnel>::iterator qAddIt = addUpTransportLayerInfoList_m.begin();
      int qAddItIndex = 0;
      for (;qAddIt != addUpTransportLayerInfoList_m.end(); ++qAddIt ) {
          
          oss << "addTransportLayerAddress:"
              << (int)qAddIt->transportLayerAddr[0]
              << (int)qAddIt->transportLayerAddr[1]
              << (int)qAddIt->transportLayerAddr[2]
              << (int)qAddIt->transportLayerAddr[3]
              << (int)qAddIt->transportLayerAddr[4]
              << (int)qAddIt->transportLayerAddr[5]
              << (int)qAddIt->transportLayerAddr[6]
              << (int)qAddIt->transportLayerAddr[7]
              << (int)qAddIt->transportLayerAddr[8]
              << (int)qAddIt->transportLayerAddr[9]
              << (int)qAddIt->transportLayerAddr[10]
              << (int)qAddIt->transportLayerAddr[11]
              << (int)qAddIt->transportLayerAddr[12]
              << (int)qAddIt->transportLayerAddr[13]
              << (int)qAddIt->transportLayerAddr[14]
              << (int)qAddIt->transportLayerAddr[15]
              << (int)qAddIt->transportLayerAddr[16]
              << (int)qAddIt->transportLayerAddr[17]
              << (int)qAddIt->transportLayerAddr[18]
              << (int)qAddIt->transportLayerAddr[19] <<endl
              << "addTeid:"
              << (int)qAddIt->gtpTeid[0]
              << (int)qAddIt->gtpTeid[1]
              << (int)qAddIt->gtpTeid[2]
              << (int)qAddIt->gtpTeid[3] <<endl;
      }  

  oss << "Common Network Instance:Flag(" << cmmnNtwrkInstancePrst_m << "),value(" << cmmnNtwrkInstance_m << ")" << endl
      << ", DatFwNotPssblePrst(" << isDataFwNotPossible_m << "),"
      << "value("<<(int)dataFwNotPssble_m<<"),"
      << "Pdu Session Type(" << (int)pduSessType_m << "),"
      << "SecIndPrst(" << isSecInd_m << "),"
      << "int("<<(int)intProtectInd_m<<"),"
      << "confd("<<(int)confdProtecInd_m << "),"
      << "maxPrtDataRate:Prst " <<isMaxPrtDataRate_m
      << " value: ("<<(int)maxPrtDataRate_m << "),"
      << "network Instance:Prst " << isNtwkInstancePrst_m
      << " value: (" << ntwkInstance_m  << ")," << endl
      << "QosFlowSetupReqList: length("<<qosFlowSetupReqList_m.size()<<")" << endl;
  
       vector<QosFlowSetupReqItem>::iterator qit = qosFlowSetupReqList_m.begin();
       int index = 0;
       for (; qit != qosFlowSetupReqList_m.end(); ++qit) {
       oss << index<<":{ qosFlowInd("<< qit->qosFlowInd << ")," << "isDynamic:"
           << qit->qosFlowLevQosPara.qosCharacter.isDynamic <<"; "
//           << "qosFlowLevelQosPara.QosCharacter.nonDynamic5QI.5QI("
//           << "pduSessResSetupReqTransfer:QosFlowSetupReqItem."
//              "qosFlowLevelQosPara.NonDynamic5QI.priLevelQos:Flag("
//           << qit->qosFlowLevQosPara.nonDyna5qi.plQosPrst << ")" << endl
//           << "value(" << qit->qosFlowLevQosPara.nonDyna5qi.priLevelQos << ")" << endl
//           << "pduSessResSetupReqTransfer:QosFlowSetupReqItem."
//              "qosFlowLevelQosPara.NonDynamic5QI.aveWin:Flag("
//           << qit->qosFlowLevQosPara.nonDyna5qi.aveWinPresent << ")" << endl
//           << "value(" << qit->qosFlowLevQosPara.nonDyna5qi.aveWin << ")" << endl
//           << "pduSessResSetupReqTransfer:QosFlowSetupReqItem."
//              "qosFlowLevelQosPara.NonDynamic5QI.maxDataBurstVol:Flag("
//           << qit->qosFlowLevQosPara.nonDyna5qi.maxDataBurstVolPresent << ")" << endl
//           << "value(" << qit->qosFlowLevQosPara.nonDyna5qi.maxDataBurstVol << ")" << endl
//
//           << "pduSessResSetupReqTransfer:qosFlowLevelQosPara.allocationAndRetentionPri:"<< endl
//           << "priLevelARP("
//           << (int)(qit->qosFlowLevQosPara.alloAndRetPri.priLevelARP) << ")" << endl
//
//           << "preEmptionCap("
//           << (int)(qit->qosFlowLevQosPara.alloAndRetPri.preEmptionCap) << ")" << endl
//
//           << "preEmptionVulabi("
//           << (int)(qit->qosFlowLevQosPara.alloAndRetPri.preEmptionVulabi) << ")" << endl
 
           << "gbrQosInfo:Flag("
           << qit->qosFlowLevQosPara.gBRQosInfoPresent << ")" <<"; "
           << "maxFlowBitRateDL("
           << qit->qosFlowLevQosPara.gBRQosInfo.maxFlowBitRateDL << ")"<<"; "

           << "maxFlowBitRateUL("
           << qit->qosFlowLevQosPara.gBRQosInfo.maxFlowBitRateUL << ")" <<"; "

           << "guaFlowBitRateDL("
           << qit->qosFlowLevQosPara.gBRQosInfo.guaFlowBitRateDL << ")" <<"; "

           << "guaFlowBitRateUL("
           << qit->qosFlowLevQosPara.gBRQosInfo.guaFlowBitRateUL << ")"<<"; "

           << "gbrQosInfo.notiControl:Flag("
           << qit->qosFlowLevQosPara.gBRQosInfo.notiControlPresent << ") "
           << "value(" << (int)(qit->qosFlowLevQosPara.gBRQosInfo.notiControl) << ")" <<"; "

           << "gbrQosInfo.maxPacketLossRateDL:Flag("
           << qit->qosFlowLevQosPara.gBRQosInfo.maxPacketLossRateDL << ") "
           << "value(" << qit->qosFlowLevQosPara.gBRQosInfo.maxPacketLossRateDL << ")" <<"; "

           << "gbrQosInfo.maxPacketLossRateUL:Flag("
           << qit->qosFlowLevQosPara.gBRQosInfo.maxPacketLossRateUL << ") "
           << "value(" << qit->qosFlowLevQosPara.gBRQosInfo.maxPacketLossRateUL << ")"<<"; "

           << "refQosAttr:Flag("
           << qit->qosFlowLevQosPara.refQosAttrPresent << ") "
           << "value(" << (int)(qit->qosFlowLevQosPara.refQosAttr) << ")" <<"; "

           << "addQosFlowInfo:Flag("
           << qit->qosFlowLevQosPara.addQosFlowInfoPresent << ") "
           << "value(" << (int)(qit->qosFlowLevQosPara.addQosFlowInfo) << ")" <<"; "
           << "eRABId:Flag("
           << qit->eRABIdPresent<< ") "
           << "value(" << qit->eRABId << ") }" << endl;

           index++;
     }
    return oss.str();
}


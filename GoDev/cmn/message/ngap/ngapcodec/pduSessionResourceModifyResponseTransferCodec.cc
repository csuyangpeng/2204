#include <iostream>
#include <sstream>
#include <string>

#include "pduSessionResourceModifyResponseTransferCodec.h"

using namespace std;

PduSessResMdfyRespTransferCodec::PduSessResMdfyRespTransferCodec() {
  reset();
}

PduSessResMdfyRespTransferCodec::~PduSessResMdfyRespTransferCodec() {}

void PduSessResMdfyRespTransferCodec::setDlGtpTunnel(GtpTunnelInfo& tunnel)
{
    dlGtpTunnelPrst_m = true;
    dlGtpTunnel_m = tunnel;
}
GtpTunnelInfo& PduSessResMdfyRespTransferCodec::getDlGtpTunnel()
{
    return dlGtpTunnel_m;
}
bool PduSessResMdfyRespTransferCodec::isDlGtpTunnelPrst()
{
    return dlGtpTunnelPrst_m;
}

void PduSessResMdfyRespTransferCodec::setUlGtpTunnel(GtpTunnelInfo& tunnel)
{
    ulGtpTunnelPrst_m = true;
    ulGtpTunnel_m = tunnel;
}
GtpTunnelInfo& PduSessResMdfyRespTransferCodec::getUlGtpTunnel()
{
    return ulGtpTunnel_m;
}
bool PduSessResMdfyRespTransferCodec::isUlGtpTunnelPrst()
{
    return ulGtpTunnelPrst_m;
}



bool PduSessResMdfyRespTransferCodec::isQosFlowTNLInfoListPrst()
{
    return AddQosFlowTNLListPrst_m;
}

void PduSessResMdfyRespTransferCodec::addQosFlowTNLInfoList(AddQosFlowPerTNLInfo& item)
{
    AddQosFlowTNLListPrst_m = true;
    return addQosFlowTNLInfoList_m.push_back(item);
}

AddQosFlowPerTNLInfoList& PduSessResMdfyRespTransferCodec::getQosFlowTNLInfoList()
{
    return addQosFlowTNLInfoList_m;
}

bool PduSessResMdfyRespTransferCodec::isQosFlowAddOrMdfyRespListPrst()
{
    return isQosFlowAddOrMdfyRespListPrst_m;
}
void PduSessResMdfyRespTransferCodec::addQosFlowAddOrMdfyRespList(Uint32 item)
{
    isQosFlowAddOrMdfyRespListPrst_m = true;
    qosFlowAddOrMdfyRespList_m.push_back(item);
}
std::vector<unsigned int>& PduSessResMdfyRespTransferCodec::getQosFlowAddOrMdfyRespList()
{
    return qosFlowAddOrMdfyRespList_m;
}

bool PduSessResMdfyRespTransferCodec::isFailedAddOrMdfyListPrst()
{
    return isFailedAddOrMdfyListPrst_m;
}
void PduSessResMdfyRespTransferCodec::addQosFlowFailedAddOrMdfyList(QosFlowCodecItem& item)
{
    isFailedAddOrMdfyListPrst_m = true;
    qosFlowFailedAddOrMdfyList_m.push_back(item);
}
QosFlowCodecList& PduSessResMdfyRespTransferCodec::getQosFlowFailedAddOrMdfyList()
{
    return qosFlowFailedAddOrMdfyList_m;
}

bool PduSessResMdfyRespTransferCodec::reset(){
    msgPdu_m.reset_v();

    dlGtpTunnelPrst_m = false;
    dlGtpTunnel_m.ipType = ipv4;  // ipv4 address
    memset(dlGtpTunnel_m.transportLayerAddr, 0, sizeof(dlGtpTunnel_m.transportLayerAddr));
    memset(dlGtpTunnel_m.gtpTeid,0,sizeof(dlGtpTunnel_m.gtpTeid));

    ulGtpTunnelPrst_m = false;
    ulGtpTunnel_m.ipType = ipv4;  // ipv4 address
    memset(ulGtpTunnel_m.transportLayerAddr, 0, sizeof(ulGtpTunnel_m.transportLayerAddr));
    memset(ulGtpTunnel_m.gtpTeid,0,sizeof(ulGtpTunnel_m.gtpTeid));

    AddQosFlowTNLListPrst_m =false;
    addQosFlowTNLInfoList_m.clear();

    isQosFlowAddOrMdfyRespListPrst_m = false;
    qosFlowAddOrMdfyRespList_m.clear();
  
    isFailedAddOrMdfyListPrst_m = false;
    qosFlowFailedAddOrMdfyList_m.clear();
}

//bool PduSessResMdfyRespTransferCodec::encode(NgapOssCtxt* ctxt, MsgBuffer msgBuffer)
MsgBuffer PduSessResMdfyRespTransferCodec::encode(NgapOssCtxt *ctxt)
{
    MsgBuffer msgBuffer;
    msgBuffer.length = 0;
    msgBuffer.value = NULL;

    //dl gtp tunnel
    if (dlGtpTunnelPrst_m) {
        UPTransportLayerInformation* gtpTunnelInfo = msgPdu_m.getDlUpTNLInformationEncodePtr();
        if(dlGtpTunnel_m.ipType == ipv4) {
            gtpTunnelInfo->u.gTPTunnel.transportLayerAddress.length = 32; // 32 bit for ipv4
            memcpy(gtpTunnelInfo->u.gTPTunnel.transportLayerAddress.value, 
                   dlGtpTunnel_m.transportLayerAddr, 4); //for ipv4     
        }else if(dlGtpTunnel_m.ipType == ipv6){
            gtpTunnelInfo->u.gTPTunnel.transportLayerAddress.length = 128; // 16*8=128 bit for ipv6
            memcpy(gtpTunnelInfo->u.gTPTunnel.transportLayerAddress.value, 
                   dlGtpTunnel_m.transportLayerAddr, 16); //for ipv6
        }else{
            std::cout << "unsupported ip address" << std::endl;
        }
        gtpTunnelInfo->u.gTPTunnel.gTP_TEID.length = sizeof(dlGtpTunnel_m.gtpTeid);
        memcpy(gtpTunnelInfo->u.gTPTunnel.gTP_TEID.value, 
               dlGtpTunnel_m.gtpTeid, sizeof(dlGtpTunnel_m.gtpTeid));
    }

    //ul gtp tunnel
    if (ulGtpTunnelPrst_m) {
        UPTransportLayerInformation* gtpTunnelInfo = msgPdu_m.getUlUpTNLInformationEncodePtr();
        if(ulGtpTunnel_m.ipType == ipv4) {
            gtpTunnelInfo->u.gTPTunnel.transportLayerAddress.length = 32; // 32 bit for ipv4
            memcpy(gtpTunnelInfo->u.gTPTunnel.transportLayerAddress.value, 
                   ulGtpTunnel_m.transportLayerAddr, 4); //for ipv4     
        }else if(ulGtpTunnel_m.ipType == ipv6){
            gtpTunnelInfo->u.gTPTunnel.transportLayerAddress.length = 128; // 16*8=128 bit for ipv6
            memcpy(gtpTunnelInfo->u.gTPTunnel.transportLayerAddress.value, 
                   ulGtpTunnel_m.transportLayerAddr, 16); //for ipv6
        }else{
            std::cout << "unsupported ip address" << std::endl;
        }
        gtpTunnelInfo->u.gTPTunnel.gTP_TEID.length = sizeof(ulGtpTunnel_m.gtpTeid);
        memcpy(gtpTunnelInfo->u.gTPTunnel.gTP_TEID.value, 
               ulGtpTunnel_m.gtpTeid, sizeof(ulGtpTunnel_m.gtpTeid));
    }

    //additional qos flow per tnl information list
    if (AddQosFlowTNLListPrst_m) {
        unsigned int num = addQosFlowTNLInfoList_m.size();
        msgPdu_m.setNumofQosFlowPerTNLInformationListItem(num);
        for(int i = 0;i <num; i++) {
            Uint32 numofAssQosFlow = addQosFlowTNLInfoList_m[i].assQosFlowList.size();
            msgPdu_m.setNumofAddAssQosFlowListItem(numofAssQosFlow,i);
        }
        for(int i = 0; i <num; i++) {
            QosFlowPerTNLInformationItem * qosFlowPerTNLInfoPtr = msgPdu_m.getQosFlowPerTNLInformationItemEncodePtr(i);
            if (NULL == qosFlowPerTNLInfoPtr) {
                std::cout << "NULL Poninter for qosFlowPerTNLInfoPtr" << std::endl;
                return msgBuffer;
            }
            if(addQosFlowTNLInfoList_m[i].upTransportLayerInfo.ipType == ipv4) {
                qosFlowPerTNLInfoPtr->qosFlowPerTNLInformation.uPTransportLayerInformation.u.gTPTunnel.transportLayerAddress.length = 32; // 32 bit for ipv4
                memcpy(qosFlowPerTNLInfoPtr->qosFlowPerTNLInformation.uPTransportLayerInformation.u.gTPTunnel.transportLayerAddress.value, 
                    addQosFlowTNLInfoList_m[i].upTransportLayerInfo.transportLayerAddr, 
                    4); //for ipv4
            }else if(addQosFlowTNLInfoList_m[i].upTransportLayerInfo.ipType == ipv6){
                qosFlowPerTNLInfoPtr->qosFlowPerTNLInformation.uPTransportLayerInformation.u.gTPTunnel.transportLayerAddress.length = 128; // 16*8=128 bit for ipv6
                memcpy(qosFlowPerTNLInfoPtr->qosFlowPerTNLInformation.uPTransportLayerInformation.u.gTPTunnel.transportLayerAddress.value, 
                    addQosFlowTNLInfoList_m[i].upTransportLayerInfo.transportLayerAddr, 
                    16); //for ipv6
            }else{
                std::cout << "unsupported ip address" << std::endl;
            }
            
            qosFlowPerTNLInfoPtr->qosFlowPerTNLInformation.uPTransportLayerInformation.u.gTPTunnel.gTP_TEID.length = sizeof(addQosFlowTNLInfoList_m[i].upTransportLayerInfo.gtpTeid);
            memcpy(qosFlowPerTNLInfoPtr->qosFlowPerTNLInformation.uPTransportLayerInformation.u.gTPTunnel.gTP_TEID.value, 
                addQosFlowTNLInfoList_m[i].upTransportLayerInfo.gtpTeid, sizeof(addQosFlowTNLInfoList_m[i].upTransportLayerInfo.gtpTeid));
            Uint32 numofAssQosFlow = addQosFlowTNLInfoList_m[i].assQosFlowList.size();
            for (int j = 0;j < numofAssQosFlow;j++) {
                AssociatedQosFlowItem* assQosFlowItemPtr = msgPdu_m.getAddAssQosFlowItemEncodePtr(i,j);
                if (NULL == assQosFlowItemPtr) {
                    std::cout << "NULL Pointer for assQosFlowItemPtr" << std::endl;
                    return  msgBuffer;
                }
                assQosFlowItemPtr->qosFlowIdentifier = addQosFlowTNLInfoList_m[i].assQosFlowList[j].qosFlowInd;
                if(addQosFlowTNLInfoList_m[i].assQosFlowList[j].qosFlowMapIndPrst)
                {
                    assQosFlowItemPtr->bit_mask = qosFlowMappingIndication_present;
                    assQosFlowItemPtr->qosFlowMappingIndication = (_enum1)addQosFlowTNLInfoList_m[i].assQosFlowList[j].qosFlowMapInd;
                }
            }
        }
    }

    // QosFlowAddOrMdfyRespList
    if (isQosFlowAddOrMdfyRespListPrst_m) {
        unsigned int num = qosFlowAddOrMdfyRespList_m.size();
        msgPdu_m.setNumofQosFlowAddOrModRespListItem(num);
        for (int i=0; i < num; i++) {
            QosFlowAddOrModifyResponseItem* item = msgPdu_m.getQosFlowAddOrModRespItemEncodePtr(i);
            // qfi
            item->qosFlowIdentifier = qosFlowAddOrMdfyRespList_m[i];
        }
    }

    // QosFlowFailedAddOrMdfyList
    if (isFailedAddOrMdfyListPrst_m) {
        unsigned int num = qosFlowFailedAddOrMdfyList_m.size();
        msgPdu_m.setNumofQosFlowToRelListItem(num);
        for (int i=0; i < num; i++){
            QosFlowWithCauseItem* item = msgPdu_m.getQosFlowToRelItemEncodePtr(i);
            QosFlowCodecItem* qosFlow = &qosFlowFailedAddOrMdfyList_m[i];
        
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
            }
        }
    }

    msgPdu_m.encodeOssToPer(ctxt);

    OssBuf *perBuf = msgPdu_m.getPerBufRef();

    msgBuffer.length = perBuf->length;
    msgBuffer.value = perBuf->value;

    return msgBuffer;
}

bool 
PduSessResMdfyRespTransferCodec::decode(NgapOssCtxt *ctxt, MsgBuffer msgBuf){

    reset();
    
    OssBuf ossbuf;
    ossbuf.length = msgBuf.length;
    ossbuf.value = msgBuf.value;
    
    msgPdu_m.setPerBufRef_v(ossbuf);
    
    if (0 != msgPdu_m.decodePerToOss(ctxt)) {
      return false;
    }

    //dl gtp tunnel
    if (msgPdu_m.isDlUpTNLInformationPrst()) {
        dlGtpTunnelPrst_m = true;
    
        UPTransportLayerInformation * transLayerinfo 
            = msgPdu_m.getDlUpTNLInformationDecodedPtr();

        if (transLayerinfo->choice != gTPTunnel_chosen) {
            std::cout << "unsupported transport layer, only gtp tunnel supported."<< endl;
            return false;
        }
    
        unsigned short addrLen = transLayerinfo->u.gTPTunnel.transportLayerAddress.length;
        if (addrLen == 32){
            dlGtpTunnel_m.ipType = ipv4;
            memcpy(dlGtpTunnel_m.transportLayerAddr, transLayerinfo->u.gTPTunnel.transportLayerAddress.value, 4);
        } else if(addrLen == 128 ){
            dlGtpTunnel_m.ipType = ipv6;
            memcpy(dlGtpTunnel_m.transportLayerAddr, transLayerinfo->u.gTPTunnel.transportLayerAddress.value, 16);
        }else{
            std::cout << "unsupported ip address" << std::endl;
            return false;
        }
        if (transLayerinfo->u.gTPTunnel.gTP_TEID.length == 4){
            memcpy(dlGtpTunnel_m.gtpTeid, transLayerinfo->u.gTPTunnel.gTP_TEID.value,sizeof(Uint32));
        }else{
            std::cout << "invalid gtp teid" << std::endl;
            return false;
        }
    }

    //ul gtp tunnel
    if (msgPdu_m.isUlUpTNLInformationPrst()) {
        ulGtpTunnelPrst_m = true;

        UPTransportLayerInformation * transLayerinfo 
            = msgPdu_m.getUlUpTNLInformationDecodedPtr();

        if (transLayerinfo->choice != gTPTunnel_chosen) {
            std::cout << "unsupported transport layer, only gtp tunnel supported."<< endl;
            return false;
        }
    
        unsigned short addrLen = transLayerinfo->u.gTPTunnel.transportLayerAddress.length;
        if (addrLen == 32){
            ulGtpTunnel_m.ipType = ipv4;
            memcpy(ulGtpTunnel_m.transportLayerAddr, transLayerinfo->u.gTPTunnel.transportLayerAddress.value, 4);
        } else if(addrLen == 128 ){
            ulGtpTunnel_m.ipType = ipv6;
            memcpy(ulGtpTunnel_m.transportLayerAddr, transLayerinfo->u.gTPTunnel.transportLayerAddress.value, 16);
        }else{
            std::cout << "unsupported ip address" << std::endl;
            return false;
        }
        if (transLayerinfo->u.gTPTunnel.gTP_TEID.length == 4){
            memcpy(ulGtpTunnel_m.gtpTeid, transLayerinfo->u.gTPTunnel.gTP_TEID.value,sizeof(Uint32));
        }else{
            std::cout << "invalid gtp teid" << std::endl;
            return false;
        }
    }

    // additional dl qos flow perTNL
    if (msgPdu_m.isQosFlowPerTNLInformationListPrst()) {
        AddQosFlowTNLListPrst_m = true;
        for(int i = 0; i < msgPdu_m.getNumofQosFlowPerTNLInformationListItem();i++) {
            QosFlowPerTNLInformationItem* item = msgPdu_m.getQosFlowPerTNLInformationItemDecodedPtr(i);
            AddQosFlowPerTNLInfo addQosFlowPerTNLInfo_m;

            if (item->qosFlowPerTNLInformation.uPTransportLayerInformation.choice != gTPTunnel_chosen) {
                std::cout << "unsupported transport layer, only gtp tunnel supported."<< endl;
                return false;
            }
            unsigned short addrLen = item->qosFlowPerTNLInformation.uPTransportLayerInformation.u.gTPTunnel.transportLayerAddress.length;
            if (addrLen == 32){
                addQosFlowPerTNLInfo_m.upTransportLayerInfo.ipType = ipv4;
                memcpy(addQosFlowPerTNLInfo_m.upTransportLayerInfo.transportLayerAddr, item->qosFlowPerTNLInformation.uPTransportLayerInformation.u.gTPTunnel.transportLayerAddress.value, 4);
            } else if(addrLen == 128 ){
                addQosFlowPerTNLInfo_m.upTransportLayerInfo.ipType = ipv6;
                memcpy(addQosFlowPerTNLInfo_m.upTransportLayerInfo.transportLayerAddr, item->qosFlowPerTNLInformation.uPTransportLayerInformation.u.gTPTunnel.transportLayerAddress.value, 16);
            }else{
                std::cout << "unsupported ip address" << std::endl;
                return false;
            }
            if (item->qosFlowPerTNLInformation.uPTransportLayerInformation.u.gTPTunnel.gTP_TEID.length == 4){
                memcpy(addQosFlowPerTNLInfo_m.upTransportLayerInfo.gtpTeid, item->qosFlowPerTNLInformation.uPTransportLayerInformation.u.gTPTunnel.gTP_TEID.value,sizeof(Uint32));
            }else{
                std::cout << "invalid gtp teid" << std::endl;
                return false;
            }
            
            for(int j = 0; j<int(msgPdu_m.getNumofAddAssQosFlowListItem(i)); j++){
                AssociatedQosFlowItem* Assitem = msgPdu_m.getAddAssQosFlowItemDecodedPtr(item, j);
                AssQosFlowItem tempAssQosFlowItem;
                AssQosFlowItem assQosFlowItem_m;
                tempAssQosFlowItem.qosFlowInd = Assitem->qosFlowIdentifier;
                
                if(Assitem->bit_mask & qosFlowMappingIndication_present)
                {
                    tempAssQosFlowItem.qosFlowMapInd = (Uint8)Assitem->qosFlowMappingIndication;
                    tempAssQosFlowItem.qosFlowMapIndPrst = true;
                }
                addQosFlowPerTNLInfo_m.assQosFlowList.push_back(tempAssQosFlowItem);
            }
            addQosFlowTNLInfoList_m.push_back(addQosFlowPerTNLInfo_m);
        }
    }

    //QosFlowAddOrMdfyRespList
    if (msgPdu_m.isQosFlowAddOrModifyRespListPrst()) {
        isQosFlowAddOrMdfyRespListPrst_m = true;
        for(int i=0; i<msgPdu_m.getNumofQosFlowAddOrModRespListItem(); i++){
        QosFlowAddOrModifyResponseItem* item = msgPdu_m.getQosFlowAddOrModRespItemDecodedPtr(i);
        //push back the qos flow 
        qosFlowAddOrMdfyRespList_m.push_back(item->qosFlowIdentifier);
        }
    }
    
    //QosFlowFailedAddOrMdfyList
    if (msgPdu_m.isQosFlowToRelListPrst()) {
        isFailedAddOrMdfyListPrst_m = true;
        for(int i=0; i<msgPdu_m.getNumofQosFlowToRelListItem(); i++){
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
            }
        
            //push back the qos flow 
            qosFlowFailedAddOrMdfyList_m.push_back(qosFlow);
        }
    }
    return true;
}


std::string 
PduSessResMdfyRespTransferCodec::toString(){
  ostringstream oss;

  oss << "PduSessResMdfyRespTransferCodec: " << endl
      << "dl gtp tunnel:(Prst{" << dlGtpTunnelPrst_m << "}"
      << "TransportLayerAddress:"
      << (int)dlGtpTunnel_m.transportLayerAddr[0]
      << (int)dlGtpTunnel_m.transportLayerAddr[1]
      << (int)dlGtpTunnel_m.transportLayerAddr[2]
      << (int)dlGtpTunnel_m.transportLayerAddr[3]
      << (int)dlGtpTunnel_m.transportLayerAddr[4]
      << (int)dlGtpTunnel_m.transportLayerAddr[5]
      << (int)dlGtpTunnel_m.transportLayerAddr[6]
      << (int)dlGtpTunnel_m.transportLayerAddr[7]
      << (int)dlGtpTunnel_m.transportLayerAddr[8]
      << (int)dlGtpTunnel_m.transportLayerAddr[9]
      << (int)dlGtpTunnel_m.transportLayerAddr[10]
      << (int)dlGtpTunnel_m.transportLayerAddr[11]
      << (int)dlGtpTunnel_m.transportLayerAddr[12]
      << (int)dlGtpTunnel_m.transportLayerAddr[13]
      << (int)dlGtpTunnel_m.transportLayerAddr[14]
      << (int)dlGtpTunnel_m.transportLayerAddr[15]
      << (int)dlGtpTunnel_m.transportLayerAddr[16]
      << (int)dlGtpTunnel_m.transportLayerAddr[17]
      << (int)dlGtpTunnel_m.transportLayerAddr[18]
      << (int)dlGtpTunnel_m.transportLayerAddr[19]
      << ", Teid:"
      << (int)dlGtpTunnel_m.gtpTeid[0]
      << (int)dlGtpTunnel_m.gtpTeid[1]
      << (int)dlGtpTunnel_m.gtpTeid[2]
      << (int)dlGtpTunnel_m.gtpTeid[3] << ")" << endl
      << "Additional Ng-U Up Tnl Information:{" << AddQosFlowTNLListPrst_m << "}";
      vector <AddQosFlowPerTNLInfo>::iterator AddIt;
      for(AddIt = addQosFlowTNLInfoList_m.begin();AddIt <addQosFlowTNLInfoList_m.end() ;AddIt++)
      {
          oss << "TransportLayerAddress:"
              << (int)AddIt->upTransportLayerInfo.transportLayerAddr[0]
              << (int)AddIt->upTransportLayerInfo.transportLayerAddr[1]
              << (int)AddIt->upTransportLayerInfo.transportLayerAddr[2]
              << (int)AddIt->upTransportLayerInfo.transportLayerAddr[3]
              << (int)AddIt->upTransportLayerInfo.transportLayerAddr[4]
              << (int)AddIt->upTransportLayerInfo.transportLayerAddr[5]
              << (int)AddIt->upTransportLayerInfo.transportLayerAddr[6]
              << (int)AddIt->upTransportLayerInfo.transportLayerAddr[7]
              << (int)AddIt->upTransportLayerInfo.transportLayerAddr[8]
              << (int)AddIt->upTransportLayerInfo.transportLayerAddr[9]
              << (int)AddIt->upTransportLayerInfo.transportLayerAddr[10]
              << (int)AddIt->upTransportLayerInfo.transportLayerAddr[11]
              << (int)AddIt->upTransportLayerInfo.transportLayerAddr[12]
              << (int)AddIt->upTransportLayerInfo.transportLayerAddr[13]
              << (int)AddIt->upTransportLayerInfo.transportLayerAddr[14]
              << (int)AddIt->upTransportLayerInfo.transportLayerAddr[15]
              << (int)AddIt->upTransportLayerInfo.transportLayerAddr[16]
              << (int)AddIt->upTransportLayerInfo.transportLayerAddr[17]
              << (int)AddIt->upTransportLayerInfo.transportLayerAddr[18]
              << (int)AddIt->upTransportLayerInfo.transportLayerAddr[19]
              << ", Teid:"
              << (int)AddIt->upTransportLayerInfo.gtpTeid[0]
              << (int)AddIt->upTransportLayerInfo.gtpTeid[1]
              << (int)AddIt->upTransportLayerInfo.gtpTeid[2]
              << (int)AddIt->upTransportLayerInfo.gtpTeid[3] << ")" << endl;
      }
      for(int i=0;i<qosFlowAddOrMdfyRespList_m.size();i++)//size()容器中实际数据个数 
      {
         oss << qosFlowAddOrMdfyRespList_m[i] << endl;
      }
  return oss.str();
}

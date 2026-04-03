#include <iostream>
#include <sstream>
#include <string>

#include "pduSessionResourceSetupResponseTransferCodec.h"

using namespace std;

PduSesResSetupRespTransferCodec::PduSesResSetupRespTransferCodec() {
  reset();
}

PduSesResSetupRespTransferCodec::~PduSesResSetupRespTransferCodec() {}


void PduSesResSetupRespTransferCodec::setGtpTunnel(GtpTunnelInfo& tunnel)
{
    gtpTunnel_m = tunnel;
}
GtpTunnelInfo& PduSesResSetupRespTransferCodec::getGtpTunnel()
{
    return gtpTunnel_m;
}
void PduSesResSetupRespTransferCodec::addQosFlowIndList(AssQosFlowItem& val)
{
    qfiList_m.push_back(val);
}
QosFlowIndList& PduSesResSetupRespTransferCodec::getQosFlowIndList()
{
    return qfiList_m;
}

bool PduSesResSetupRespTransferCodec::isQosFlowTNLInfoListPrst()
{
    return AddQosFlowTNLListPrst_m;
}

void PduSesResSetupRespTransferCodec::addQosFlowTNLInfoList(AddQosFlowPerTNLInfo& item)
{
    AddQosFlowTNLListPrst_m = true;
    return addQosFlowTNLInfoList_m.push_back(item);
}

AddQosFlowPerTNLInfoList& PduSesResSetupRespTransferCodec::getQosFlowTNLInfoList()
{
    return addQosFlowTNLInfoList_m;
}

bool PduSesResSetupRespTransferCodec::isSecResultPrst()
{
    return isSecResult_m;
}
void PduSesResSetupRespTransferCodec::setSecResult(Uint8 intInd, Uint8 confdInd)
{
    isSecResult_m = true;
    intProtectRst_m = intInd;
    confdProtecRst_m = confdInd;
}
Uint8 PduSesResSetupRespTransferCodec::getIndPrctResult()
{
    return intProtectRst_m;
}
Uint8 PduSesResSetupRespTransferCodec::getConfdPrctResult()
{
    return confdProtecRst_m;
}

bool PduSesResSetupRespTransferCodec::isFailedSetupListPrst()
{
    return isFailedSetupListPrst_m;
}
void PduSesResSetupRespTransferCodec::addQosFlowFailedSetupList(QosFlowCodecItem& item)
{
    isFailedSetupListPrst_m = true;
    qosFlowFailedSetupList_m.push_back(item);
}
QosFlowCodecList& PduSesResSetupRespTransferCodec::getQosFlowFailedSetupList()
{
    return qosFlowFailedSetupList_m;
}

bool PduSesResSetupRespTransferCodec::reset(){
    msgPdu_m.reset_v();

    gtpTunnel_m.ipType = ipv4;  // ipv4 address
    memset(gtpTunnel_m.transportLayerAddr, 0, sizeof(gtpTunnel_m.transportLayerAddr));
    memset(gtpTunnel_m.gtpTeid,0,sizeof(gtpTunnel_m.gtpTeid));

    qfiList_m.clear();

    isSecResult_m = false;
    intProtectRst_m = 0;
    confdProtecRst_m = 0;

    AddQosFlowTNLListPrst_m =false;
    addQosFlowTNLInfoList_m.clear();
  
    isFailedSetupListPrst_m = false;
    qosFlowFailedSetupList_m.clear();
}

//bool PduSesResSetupRespTransferCodec::encode(NgapOssCtxt* ctxt, MsgBuffer msgBuffer)
MsgBuffer PduSesResSetupRespTransferCodec::encode(NgapOssCtxt *ctxt)
{
    MsgBuffer msgBuffer;
    msgBuffer.length = 0;
    msgBuffer.value = NULL;

    // qos flow per tul information
    QosFlowPerTNLInformation* qosFlowInfo = msgPdu_m.getQosFlowInfoEncodePtr();
    if(gtpTunnel_m.ipType == ipv4) {
        qosFlowInfo->uPTransportLayerInformation.u.gTPTunnel.transportLayerAddress.length = 32; // 32 bit for ipv4
        memcpy(qosFlowInfo->uPTransportLayerInformation.u.gTPTunnel.transportLayerAddress.value, 
               gtpTunnel_m.transportLayerAddr, 
               4); //for ipv4       
    }else if(gtpTunnel_m.ipType == ipv6){
        qosFlowInfo->uPTransportLayerInformation.u.gTPTunnel.transportLayerAddress.length = 128; // 16*8=128 bit for ipv6
        memcpy(qosFlowInfo->uPTransportLayerInformation.u.gTPTunnel.transportLayerAddress.value, 
               gtpTunnel_m.transportLayerAddr, 
               16); //for ipv6
    }else{
        std::cout << "unsupported ip address" << std::endl;
    }
    
    qosFlowInfo->uPTransportLayerInformation.u.gTPTunnel.gTP_TEID.length = sizeof(gtpTunnel_m.gtpTeid);
    memcpy(qosFlowInfo->uPTransportLayerInformation.u.gTPTunnel.gTP_TEID.value, 
           gtpTunnel_m.gtpTeid, sizeof(gtpTunnel_m.gtpTeid));

    unsigned int num = qfiList_m.size();
    msgPdu_m.setNumofAssQosFlowListItem_m_v(num);
    for(int i=0; i<num; i++){
        AssociatedQosFlowItem* item = msgPdu_m.getAssQosFlowItemEncodePtr(i);
        item->qosFlowIdentifier = qfiList_m[i].qosFlowInd;
        if(qfiList_m[i].qosFlowMapIndPrst)
        {
            item->bit_mask = qosFlowMappingIndication_present;
            item->qosFlowMappingIndication = (_enum1)qfiList_m[i].qosFlowMapInd;
        }
    }

    //additional qos flow per tnl information list
    if (AddQosFlowTNLListPrst_m) {
        unsigned int num = addQosFlowTNLInfoList_m.size();
        msgPdu_m.setNumOfQosFlowPerTNLInformationListItem_v(num);
        for(int i = 0;i <num; i++) {
            Uint32 numofAssQosFlow = addQosFlowTNLInfoList_m[i].assQosFlowList.size();
            msgPdu_m.setNumofAddAssQosFlowListItem_m_v(numofAssQosFlow,i);
        }
        for(int i = 0; i <num; i++) {
            QosFlowPerTNLInformationItem * qosFlowPerTNLInfoPtr = msgPdu_m.getQosFlowPerTNLInformationListItemEncodePtr(i);
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

           
    // security result
    if (isSecResult_m){
        msgPdu_m.getSecurityResultEncodePtr()->confidentialityProtectionResult
            = (ConfidentialityProtectionResult)confdProtecRst_m;
        msgPdu_m.getSecurityResultEncodePtr()->integrityProtectionResult
            = (IntegrityProtectionResult)intProtectRst_m;
    }
    
    // QosFlowFailedSetupList
    num = qosFlowFailedSetupList_m.size();
    msgPdu_m.setNumofQosFlowListItem_m_v(num);
    for (int i=0; i < num; i++){
        QosFlowWithCauseItem* item = msgPdu_m.getQosFlowItemEncodePtr(i);
        QosFlowCodecItem* qosFlow = &qosFlowFailedSetupList_m[i];

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

    msgPdu_m.encodeOssToPer(ctxt);

    OssBuf *perBuf = msgPdu_m.getPerBufRef();

    msgBuffer.length = perBuf->length;
    msgBuffer.value = perBuf->value;

    return msgBuffer;
}

bool 
PduSesResSetupRespTransferCodec::decode(NgapOssCtxt *ctxt, MsgBuffer msgBuf){

    reset();
    
    OssBuf ossbuf;
    ossbuf.length = msgBuf.length;
    ossbuf.value = msgBuf.value;
    
    msgPdu_m.setPerBufRef_v(ossbuf);
    
    if (0 != msgPdu_m.decodePerToOss(ctxt)) {
      return false;
    }

    // qos flow per tul information
    UPTransportLayerInformation * transLayerinfo 
        = &(msgPdu_m.getQosFlowInfoDecodedPtr()->uPTransportLayerInformation);

    if (transLayerinfo->choice != gTPTunnel_chosen) {
        std::cout << "unsupported transport layer, only gtp tunnel supported."<< endl;
        return false;
    }
    
    unsigned short addrLen = transLayerinfo->u.gTPTunnel.transportLayerAddress.length;
    if (addrLen == 32){
        gtpTunnel_m.ipType = ipv4;
        memcpy(gtpTunnel_m.transportLayerAddr, transLayerinfo->u.gTPTunnel.transportLayerAddress.value, 4);
    } else if(addrLen == 128 ){
        gtpTunnel_m.ipType = ipv6;
        memcpy(gtpTunnel_m.transportLayerAddr, transLayerinfo->u.gTPTunnel.transportLayerAddress.value, 16);
    }else{
        std::cout << "unsupported ip address" << std::endl;
        return false;
    }
    if (transLayerinfo->u.gTPTunnel.gTP_TEID.length == 4){
        memcpy(gtpTunnel_m.gtpTeid, transLayerinfo->u.gTPTunnel.gTP_TEID.value,sizeof(Uint32));
    }else{
        std::cout << "invalid gtp teid" << std::endl;
        return false;
    }

    // additional dl qos flow perTNL
    if (msgPdu_m.isQosFlowPerTNLInformationListPrst()) {
        AddQosFlowTNLListPrst_m = true;
        for(int i = 0; i < msgPdu_m.getNumOfQosFlowPerTNLInformationListItem();i++) {
            QosFlowPerTNLInformationItem* item = msgPdu_m.getQosFlowPerTNLInformationListItemDecodedPtr(i);
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

    
    for(int i=0; i<int(msgPdu_m.getNumofAssQosFlowListItem()); i++){
        AssociatedQosFlowItem* item = msgPdu_m.getAssQosFlowItemDecodedPtr(i);
        AssQosFlowItem tempAssQosFlowItem;
        tempAssQosFlowItem.qosFlowInd = item->qosFlowIdentifier;
        if(item->bit_mask & qosFlowMappingIndication_present)
        {
            tempAssQosFlowItem.qosFlowMapInd = (Uint8)item->qosFlowMappingIndication;
            tempAssQosFlowItem.qosFlowMapIndPrst = true;
        }
        qfiList_m.push_back(tempAssQosFlowItem);
    }

    // security indication
    if(msgPdu_m.isSecResultPrst()){
        isSecResult_m = true;
        intProtectRst_m = msgPdu_m.getSecurityResultDecodedPtr()->integrityProtectionResult;
        confdProtecRst_m = msgPdu_m.getSecurityResultDecodedPtr()->confidentialityProtectionResult;
    }

    // FailedSetupList
    if(msgPdu_m.isQosFlowFailedSetupListPrst()){
        
        isFailedSetupListPrst_m = true;

        for(int i=0; i<msgPdu_m.getNumofQosFlowListItem(); i++){
            QosFlowWithCauseItem* item = msgPdu_m.getQosFlowItemDecodedPtr(i);
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
            qosFlowFailedSetupList_m.push_back(qosFlow);
        }
    }

    return true;
}


std::string 
PduSesResSetupRespTransferCodec::toString(){
  ostringstream oss;

  oss << "PduSesResSetupRespTransferCodec: " << endl
      << "securityResult:Prst " <<isSecResult_m <<endl
      << "intProtectRst: " <<(int)intProtectRst_m <<endl
      << "Associated Qos Flow List:" <<endl;
      vector<AssQosFlowItem>::iterator it;
      for (it=qfiList_m.begin(); it<qfiList_m.end() ; it++)
      {
          oss << "qos flow identifier: " << it->qosFlowInd << endl
              <<"qos flow mapping indication:prst " <<it->qosFlowMapIndPrst << endl
              <<"value: " <<(int)it->qosFlowMapInd <<endl;
      }
     oss << "Additional Ng-U Up Tnl Information:{" << AddQosFlowTNLListPrst_m << "}";
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

  return oss.str();
}



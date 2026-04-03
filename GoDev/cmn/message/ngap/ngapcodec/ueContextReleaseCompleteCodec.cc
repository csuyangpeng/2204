#include <iostream>
#include <sstream>
#include <string>

#include "ueContextReleaseCompleteCodec.h"

using namespace std;

UeContextReleaseCompleteCodec::UeContextReleaseCompleteCodec() 
{
    reset();
}

UeContextReleaseCompleteCodec::~UeContextReleaseCompleteCodec() 
{
}

void UeContextReleaseCompleteCodec::setRanUeNgapId(Uint32 id) 
{
    ranUeNgapId_m = id;
}

Uint32 UeContextReleaseCompleteCodec::getRanUeNgapId() 
{
    return ranUeNgapId_m;
}

void UeContextReleaseCompleteCodec::setAmfUeNgapId(Uint64 id) 
{
    amfUeNgapId_m = id;
}

Uint64 UeContextReleaseCompleteCodec::getAmfUeNgapId() 
{
    return amfUeNgapId_m;
}


void 
UeContextReleaseCompleteCodec::setUserLocationInfo(UserLocationInfo& info)
{
    isUliPrst_m = true;
    userLocationInfo_m = info;
}

UserLocationInfo&
UeContextReleaseCompleteCodec::getUserLocationInfo()
{
    return userLocationInfo_m;
}

bool 
UeContextReleaseCompleteCodec::isUliPrst()
{
    return isUliPrst_m;
}

void 
UeContextReleaseCompleteCodec::setRecommandCellList(RecommandCellList& cell)
{
    isRecommandInfoPrst_m = true;
    recommandCellList_m = cell;
}

RecommandCellList& 
UeContextReleaseCompleteCodec::getRecommandCellList()
{
    return recommandCellList_m;
}

void 
UeContextReleaseCompleteCodec::setRecommandRanNodeList(RecommandRanNodeList& ranNode)
{
    isRecommandInfoPrst_m = true;
    recommandRanNodeList_m = ranNode;
}

RecommandRanNodeList& 
UeContextReleaseCompleteCodec::getRecommandRanNodeList()
{
    return recommandRanNodeList_m;
}
bool 
UeContextReleaseCompleteCodec::isRecommandInfoPrst()
{
    return isRecommandInfoPrst_m;
}

void UeContextReleaseCompleteCodec::addPduSessResListRelCpl(PduSessResRelCplItem& Item) 
{
    pduSessResListRelCpl_m.push_back(Item);
}

PduSessResRelCplList& 
UeContextReleaseCompleteCodec::getPduSessResListRelCpl() 
{
    return pduSessResListRelCpl_m;
}

bool UeContextReleaseCompleteCodec::reset() 
{
    msgPdu_m.reset_v();

    amfUeNgapId_m = 0;
    ranUeNgapId_m = 0;
    isUliPrst_m = false;
    memset(&userLocationInfo_m,0xff,sizeof(userLocationInfo_m));
    memset(&recommandCellList_m,0,sizeof(recommandCellList_m));
    isRecommandInfoPrst_m = false;
    memset(&recommandRanNodeList_m,0,sizeof(recommandRanNodeList_m));
    pduSessResListRelCpl_m.clear();
}

//bool UeContextReleaseCompleteCodec::encode(NgapOssCtxt* ctxt, MsgBuffer msgBuffer)
MsgBuffer UeContextReleaseCompleteCodec::encode(NgapOssCtxt *ctxt)
{
    MsgBuffer msgBuffer;
    msgBuffer.length = 0;
    msgBuffer.value = NULL;

//    cout << "Codec encode Debug" << endl;

    // encode RAN UE NGAP ID
    RAN_UE_NGAP_ID* ranUeIdPtr = msgPdu_m.getRanUeNgapIdEncodePtr();
    *ranUeIdPtr = ranUeNgapId_m;

    // encode AMF UE NGAP ID
    AMF_UE_NGAP_ID* amfUeIdPtr = msgPdu_m.getAmfUeNgapIdEncodePtr();
    *amfUeIdPtr = amfUeNgapId_m;


    //encode Uer Location Information
    if (isUliPrst_m)
    {
        UserLocationInformation* uliPtr=msgPdu_m.getUseLocationInformationEncodePtr();

        uliPtr->u.userLocationInformationNR.nR_CGI.pLMNIdentity.length = sizeof(userLocationInfo_m.nRCgi.plmn);
        memcpy(uliPtr->u.userLocationInformationNR.nR_CGI.pLMNIdentity.value,
                userLocationInfo_m.nRCgi.plmn,
                sizeof(PlmnId));
        uliPtr->u.userLocationInformationNR.nR_CGI.nRCellIdentity.length = NRCellIdentityBitLen;
        memcpy(uliPtr->u.userLocationInformationNR.nR_CGI.nRCellIdentity.value,
                userLocationInfo_m.nRCgi.nrCellId,
                sizeof(userLocationInfo_m.nRCgi.nrCellId));
                
        uliPtr->u.userLocationInformationNR.tAI.pLMNIdentity.length=sizeof(userLocationInfo_m.tai.plmn);
        memcpy(uliPtr->u.userLocationInformationNR.tAI.pLMNIdentity.value,
                userLocationInfo_m.tai.plmn,sizeof(userLocationInfo_m.tai.plmn));
        uliPtr->u.userLocationInformationNR.tAI.tAC.length = sizeof(userLocationInfo_m.tai.tac);
        memcpy(uliPtr->u.userLocationInformationEUTRA.tAI.tAC.value,
                userLocationInfo_m.tai.tac,sizeof(userLocationInfo_m.tai.tac));
        
        if(userLocationInfo_m.isTimeStampPrst)
        {
           uliPtr->u.userLocationInformationNR.bit_mask = UserLocationInformationNR_timeStamp_present;
           uliPtr->u.userLocationInformationNR.timeStamp.length = sizeof(userLocationInfo_m.timeStamp);
           memcpy(uliPtr->u.userLocationInformationNR.timeStamp.value,
                userLocationInfo_m.timeStamp,sizeof(userLocationInfo_m.timeStamp));
        }
        if(userLocationInfo_m.psCellInfoPrst)
        {
            uliPtr->u.userLocationInformationNR.bit_mask = UserLocationInformationNR_iE_Extensions_present;
            NGRAN_CGI *psCellInfoPtr = msgPdu_m.getNgRanCgiEncodePtr();
            psCellInfoPtr->u.nR_CGI.pLMNIdentity.length = sizeof(userLocationInfo_m.psCellInfo.plmn);
            memcpy(psCellInfoPtr->u.nR_CGI.pLMNIdentity.value,
                   userLocationInfo_m.psCellInfo.plmn,
                   sizeof(PlmnId));
            psCellInfoPtr->u.nR_CGI.nRCellIdentity.length = NRCellIdentityBitLen;
            memcpy(psCellInfoPtr->u.nR_CGI.nRCellIdentity.value,
                    userLocationInfo_m.psCellInfo.nrCellId,
                    sizeof(userLocationInfo_m.psCellInfo.nrCellId));
        }

    }

    //encode Cell and Ran Node Page
    //InfoOnRecommendedCellsAndRANNodesForPaging* cellAndRanNodePagePtr = 
    //    msgPdu_m.getInfoOnRecommendedCellsAndRANNodesForPagingEncodePtr();
    if (isRecommandInfoPrst_m)
    {
        Uint32 numOfRecommendedCellItem_ma = recommandCellList_m.size();
        msgPdu_m.setNumOfRecommendedCellItem_v(numOfRecommendedCellItem_ma);
        for(int i =0;i<numOfRecommendedCellItem_ma;i++)
        {
            RecommendedCellItem* cellPtr = msgPdu_m.getRecommendedCellItemEncodePtr(i);
            if (NULL == cellPtr)
            {
                std::cout << "NULL Pointer for RecommendedCellItem" << std::endl;
                return msgBuffer;
            }

            if(recommandCellList_m[i].timeStayedPrst)
            {
                cellPtr->bit_mask = ExpectedUEMovingTrajectoryItem_timeStayedInCell_present;
                cellPtr->timeStayedInCell = recommandCellList_m[i].timeStayedInCell;
            }

            cellPtr->nGRAN_CGI.u.nR_CGI.pLMNIdentity.length = 
                sizeof(recommandCellList_m[i].nrcgi.plmn);
            memcpy(cellPtr->nGRAN_CGI.u.nR_CGI.pLMNIdentity.value,
                recommandCellList_m[i].nrcgi.plmn,
                    cellPtr->nGRAN_CGI.u.nR_CGI.pLMNIdentity.length);

            cellPtr->nGRAN_CGI.u.nR_CGI.nRCellIdentity.length = NRCellIdentityBitLen;
            memcpy(cellPtr->nGRAN_CGI.u.nR_CGI.nRCellIdentity.value,
                recommandCellList_m[i].nrcgi.nrCellId,
                    sizeof(recommandCellList_m[i].nrcgi.nrCellId));
        }


        Uint32 numOfRecommendedRANNodeItem_ma = recommandRanNodeList_m.size();
        msgPdu_m.setNumOfRecommendedRANNodeItem_v(numOfRecommendedRANNodeItem_ma);
        for(int i =0;i<numOfRecommendedRANNodeItem_ma;i++)
        {
            RecommendedRANNodeItem* ranNodePtr = msgPdu_m.getRecommendedRANNodeItemEncodePtr(i);
            if (NULL == ranNodePtr)
            {
                std::cout << "NULL Pointer for RecommendedRANNode" << std::endl;
                return msgBuffer;
            }
            ranNodePtr->aMFPagingTarget.choice = recommandRanNodeList_m[i].TargetChoice;
            switch(ranNodePtr->aMFPagingTarget.choice)
            {
                case globalRANNodeID_chosen:
                    ranNodePtr->aMFPagingTarget.u.globalRANNodeID.choice = globalGNB_ID_chosen;
                    ranNodePtr->aMFPagingTarget.u.globalRANNodeID.u.globalGNB_ID.bit_mask = 0;
                    ranNodePtr->aMFPagingTarget.u.globalRANNodeID.u.globalGNB_ID
                        .pLMNIdentity.length = sizeof(recommandRanNodeList_m[i].ggnbId.plmnid);
                    memcpy(ranNodePtr->aMFPagingTarget.u.globalRANNodeID.u.globalGNB_ID.pLMNIdentity.value,
                           recommandRanNodeList_m[i].ggnbId.plmnid, 
                           sizeof(recommandRanNodeList_m[i].ggnbId.plmnid));

                    ranNodePtr->aMFPagingTarget.u.globalRANNodeID.u.globalGNB_ID.gNB_ID.choice = gNB_ID_chosen;
                    ranNodePtr->aMFPagingTarget.u.globalRANNodeID.u.globalGNB_ID.gNB_ID.u.gNB_ID
                        .length = recommandRanNodeList_m[i].ggnbId.gnbId.len;
                    memcpy(ranNodePtr->aMFPagingTarget.u.globalRANNodeID.u.globalGNB_ID.gNB_ID.u.gNB_ID.value,
                        recommandRanNodeList_m[i].ggnbId.gnbId.val, 
                        sizeof(recommandRanNodeList_m[i].ggnbId.gnbId.val));
                    break;

                case tAI_chosen:
                    ranNodePtr->aMFPagingTarget.u.tAI.pLMNIdentity.length = 
                        sizeof(recommandRanNodeList_m[i].tai.plmn);

                    memcpy(ranNodePtr->aMFPagingTarget.u.tAI.pLMNIdentity.value,
                        recommandRanNodeList_m[i].tai.plmn,
                            ranNodePtr->aMFPagingTarget.u.tAI.pLMNIdentity.length);

                    ranNodePtr->aMFPagingTarget.u.tAI.tAC.length = 
                        sizeof(recommandRanNodeList_m[i].tai.tac);

                    memcpy(ranNodePtr->aMFPagingTarget.u.tAI.tAC.value,
                        recommandRanNodeList_m[i].tai.tac,
                            ranNodePtr->aMFPagingTarget.u.tAI.tAC.length);
                    break;

                default:
                    break;
            }

        }
    }

    
    //encode PDUSessionResourceListCxtRelCpl
    Uint32 num = pduSessResListRelCpl_m.size();
    msgPdu_m.setNumofPduSessResListItemCxtRelCpl_v(num);

    for (int i = 0; i < num; i++) 
    {
        PDUSessionResourceItemCxtRelCpl *pduPtr = 
            msgPdu_m.getPduSessResItemCxtRelCplEncodePtr(i);
        if (NULL == pduPtr) 
        {
            return msgBuffer;
        }
        PduSessResRelCplItem* pduSessItem = &(pduSessResListRelCpl_m[i]);
        pduPtr->pDUSessionID = pduSessItem ->pduSessId;

        if (pduSessItem->pduSessResRelRespTransferPrst) {
            //pdu session resource release complete
            pduPtr->bit_mask = PDUSessionResourceItemCxtRelCpl_iE_Extensions_present;
            pduPtr->iE_Extensions->value.extensionValue.encoded.length
              = pduSessItem->pduSessResRelRespTransfer.length();
            memcpy(pduPtr->iE_Extensions->value.extensionValue.encoded.value,
                   pduSessItem->pduSessResRelRespTransfer.c_str(),
                   pduPtr->iE_Extensions->value.extensionValue.encoded.length);
        }
    }

    msgPdu_m.encodeOssToPer(ctxt);

    OssBuf* perBuf = msgPdu_m.getPerBufRef();

    msgBuffer.length = perBuf->length;
    msgBuffer.value = perBuf->value;

    return msgBuffer;
}

bool UeContextReleaseCompleteCodec::decode(NgapOssCtxt* ctxt,MsgBuffer msgBuf) 
{
    reset();

    OssBuf ossbuf;
    ossbuf.length = msgBuf.length;
    ossbuf.value = msgBuf.value;

    msgPdu_m.setPerBufRef_v(ossbuf);

    if (0 != msgPdu_m.decodePerToOss(ctxt)) 
    {
        cout << "Failed to decode Per to OSS" << endl;
        return false;
    }

    // decode RAN UE NGAP ID
    ranUeNgapId_m = *(msgPdu_m.getRanUeNgapIdDecodedPtr());

    // decode amf ue ngap id
    amfUeNgapId_m = *(msgPdu_m.getAmfUeNgapIdDecodedPtr());

    //decode User Location Information
    if(msgPdu_m.isUserLocInfoPrst()
        &&  (msgPdu_m.getUseLocationInformationDecodedPtr()->choice == userLocationInformationNR_chosen))
    {
       Uint32 length=msgPdu_m.getUseLocationInformationDecodedPtr()->u.userLocationInformationNR.nR_CGI.pLMNIdentity.length;
       memcpy(userLocationInfo_m.nRCgi.plmn,
              msgPdu_m.getUseLocationInformationDecodedPtr()->u.userLocationInformationNR.nR_CGI.pLMNIdentity.value,
              length);
           
       memcpy(userLocationInfo_m.nRCgi.nrCellId,
              msgPdu_m.getUseLocationInformationDecodedPtr()->u.userLocationInformationNR.nR_CGI.nRCellIdentity.value,
              maxNrCellIdLen);
       
      length=msgPdu_m.getUseLocationInformationDecodedPtr()->u.userLocationInformationNR.tAI.pLMNIdentity.length;
      memcpy(userLocationInfo_m.tai.plmn,
              msgPdu_m.getUseLocationInformationDecodedPtr()->u.userLocationInformationNR.tAI.pLMNIdentity.value,
              length);
       
      length=msgPdu_m.getUseLocationInformationDecodedPtr()->u.userLocationInformationNR.tAI.tAC.length;
      memcpy(userLocationInfo_m.tai.tac,
              msgPdu_m.getUseLocationInformationDecodedPtr()->u.userLocationInformationNR.tAI.tAC.value,
              length);
       
       if(msgPdu_m.getUseLocationInformationDecodedPtr()->u.userLocationInformationNR.bit_mask 
          == UserLocationInformationNR_timeStamp_present)
       {
          length=msgPdu_m.getUseLocationInformationDecodedPtr()->u.userLocationInformationNR.timeStamp.length;
          memcpy(userLocationInfo_m.timeStamp,
                 msgPdu_m.getUseLocationInformationDecodedPtr()->u.userLocationInformationNR.timeStamp.value,
                 length);
          userLocationInfo_m.isTimeStampPrst = true;
       }
       if(msgPdu_m.getUseLocationInformationDecodedPtr()->u.userLocationInformationNR.bit_mask &
           UserLocationInformationNR_iE_Extensions_present)
           {
                UserLocationInfomationNRIeExtension* psCellInfoPtr = msgPdu_m.getIeExtenIeTypeDecodedPtr();
          
                Uint32 length=psCellInfoPtr->value.extensionValue.decoded.pdu_NGRAN_CGI->u.nR_CGI.pLMNIdentity.length;
                memcpy(userLocationInfo_m.psCellInfo.plmn,
                    psCellInfoPtr->value.extensionValue.decoded.pdu_NGRAN_CGI->u.nR_CGI.pLMNIdentity.value,
                    length);
                memcpy(userLocationInfo_m.psCellInfo.nrCellId,
                    psCellInfoPtr->value.extensionValue.decoded.pdu_NGRAN_CGI->u.nR_CGI.nRCellIdentity.value,
                    maxNrCellIdLen);
          }

      isUliPrst_m = true;
    }

    //decode Info On Recommended Cells And RAN Nodes For Paging
    if(msgPdu_m.isRecommendInfoPrst())
    {
        for(Uint32 i = 0; i < msgPdu_m.getNumOfRecommendedCellItem(); i++)
        {
            RecommendedCellItem *cellPtr = 
                msgPdu_m.getRecommendedCellItemDecodedPtr(i);
            if(NULL==cellPtr)
            {
                return false;
            }
            RecommandCellItem cellItem_m;

            memcpy(cellItem_m.nrcgi.plmn,
                cellPtr->nGRAN_CGI.u.nR_CGI.pLMNIdentity.value,
                    cellPtr->nGRAN_CGI.u.nR_CGI.pLMNIdentity.length);

            memcpy(cellItem_m.nrcgi.nrCellId,
                cellPtr->nGRAN_CGI.u.nR_CGI.nRCellIdentity.value,
                    cellPtr->nGRAN_CGI.u.nR_CGI.nRCellIdentity.length);    

            if(cellPtr->bit_mask == ExpectedUEMovingTrajectoryItem_timeStayedInCell_present)
            {
                cellItem_m.timeStayedPrst=true;
                cellItem_m.timeStayedInCell = cellPtr->timeStayedInCell;
            }

            recommandCellList_m.push_back(cellItem_m);
            isRecommandInfoPrst_m = true;
        }

        for(Uint32 i = 0; i < msgPdu_m.getNumOfRecommendedRANNodeItem(); i++)
        {
            RecommendedRANNodeItem *ranNodePtr = 
                msgPdu_m.getRecommendedRANNodeItemDecodedPtr(i);
            if(NULL==ranNodePtr)
            {
                return false;
            }
            
            RecommandRanNodeItem ranNodeItem_m;

            ranNodeItem_m.TargetChoice = ranNodePtr->aMFPagingTarget.choice;
            switch(ranNodeItem_m.TargetChoice)
            {
                case globalRANNodeID_chosen:
                    memcpy(ranNodeItem_m.ggnbId.plmnid, 
                           ranNodePtr->aMFPagingTarget.u.globalRANNodeID.u.globalGNB_ID.pLMNIdentity.value, 
                           ranNodePtr->aMFPagingTarget.u.globalRANNodeID.u.globalGNB_ID.pLMNIdentity.length);
        
                    ranNodeItem_m.ggnbId.gnbId.len = 
                        ranNodePtr->aMFPagingTarget.u.globalRANNodeID.u.globalGNB_ID.gNB_ID.u.gNB_ID.length;
                    memcpy(ranNodeItem_m.ggnbId.gnbId.val, 
                           ranNodePtr->aMFPagingTarget.u.globalRANNodeID.u.globalGNB_ID.gNB_ID.u.gNB_ID.value, 
                           sizeof(ranNodePtr->aMFPagingTarget.u.globalRANNodeID.u.globalGNB_ID.gNB_ID.u.gNB_ID.value));
                    break;

                case tAI_chosen:
                    memcpy(ranNodeItem_m.tai.tac,
                        ranNodePtr->aMFPagingTarget.u.tAI.tAC.value,
                            ranNodePtr->aMFPagingTarget.u.tAI.tAC.length);

                    memcpy(ranNodeItem_m.tai.plmn,
                        ranNodePtr->aMFPagingTarget.u.tAI.pLMNIdentity.value,
                            ranNodePtr->aMFPagingTarget.u.tAI.pLMNIdentity.length);
                    break;

                default:
                    break;
            }

            recommandRanNodeList_m.push_back(ranNodeItem_m);
            isRecommandInfoPrst_m = true;
        }
    }

    //decode PDUSessionResourceListCxtRelCpl
    for (int i = 0; i < msgPdu_m.getNumofPduSessResListItemCxtRelCpl();i++) 
    {
        PDUSessionResourceItemCxtRelCpl *pduPtr =
            msgPdu_m.getPduSessResItemCxtRelCplDecodedPtr(i);
        if (NULL == pduPtr) 
        {
            return false;
        }

        PduSessResRelCplItem pduSessResRelItem_m;
        //psi
        pduSessResRelItem_m.pduSessId = pduPtr->pDUSessionID;
        if (pduPtr->bit_mask == PDUSessionResourceItemCxtRelCpl_iE_Extensions_present){
            pduSessResRelItem_m.pduSessResRelRespTransferPrst = true;

            OssBuf outputBuf = ctxt->getSecondaryCodecBuf();
            PDUSessionResourceReleaseResponseTransfer* transferPdu;
            transferPdu = pduPtr->iE_Extensions->value.extensionValue.decoded.pdu_PDUSessionResourceItemCxtRelCpl_ExtIEs_EXTENSION->decoded;
            unsigned int result = 
                ctxt->encodeOssToPer(PDUSessionResourceReleaseResponseTransfer_PDU, transferPdu, &outputBuf);
            if(result != 0){
                std::cout<< "failed to encode transfer" << std::endl;
                return false;
            }
            pduSessResRelItem_m.pduSessResRelRespTransfer.assign(
                (char *)(outputBuf.value),
                outputBuf.length);
        }
        pduSessResListRelCpl_m.push_back(pduSessResRelItem_m);
    }

    return true;
}

string UeContextReleaseCompleteCodec::toString(){
   ostringstream oss;

   oss << "UeContextReleaseCompleteCodec:" << endl
       << "1.AMF UE NGAP ID:(" << amfUeNgapId_m << ")" << endl
       << "2.RAN UE NGAP ID:(" << ranUeNgapId_m << ")" << endl
       << "3.User Location Information:" << endl
       << "NrCgi" << "plmn" << "NrCellid"
       << "Tai" << "plmn" << "tac"
       << "TimeStampPrst"  << "TimeStamp"
       << "4.Cell And Ran Node Page:" << endl
       << "5.PDU Session Resource Release Complete List:" << endl;
    return oss.str(); 
}


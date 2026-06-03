#include <iostream>
#include <sstream>
#include <string>

#include "pduSessionResourceModifyResponseCodec.h"

using namespace std;

PduSessResModifyRespCodec::PduSessResModifyRespCodec()
{
    reset();
}

PduSessResModifyRespCodec::~PduSessResModifyRespCodec()
{
}

void PduSessResModifyRespCodec::setRanUeNgapId(Uint32 id)
{
    ranUeNgapId_m = id;
}

Uint32
PduSessResModifyRespCodec::getRanUeNgapId()
{
    return ranUeNgapId_m;
}

void PduSessResModifyRespCodec::setAmfUeNgapId(Uint64 id)
{
    amfUeNgapId_m = id;
}

Uint64
PduSessResModifyRespCodec::getAmfUeNgapId()
{
    return amfUeNgapId_m;
}

void 
PduSessResModifyRespCodec::addPduSessResModifyRespList(PduSessResModifyRespItem &item)
{
    pduSessResModRespListPrst_m = true;
    pduSessResModRespList_m.push_back(item);
}

PduSessResModifyRespList& 
PduSessResModifyRespCodec::getPduSessResModifyRespList()
{
    return pduSessResModRespList_m;
}

bool 
PduSessResModifyRespCodec::isPduSessResMdfyRespListPrst()
{
    return pduSessResModRespListPrst_m;
}

void 
PduSessResModifyRespCodec::addPduSessResFailedModifyList(PduSessResFailedMdfyRespItem &item)
{
    pduSessResFailModRespListPrst_m = true;
    pduSessResFailedModList_m.push_back(item);
}

PduSessResFailedModifyList& 
PduSessResModifyRespCodec::getPduSessResFailedModifyList()
{
    return pduSessResFailedModList_m;
}

bool 
PduSessResModifyRespCodec::isPduSessResFailMdfyRespListPrst()
{
    return pduSessResFailModRespListPrst_m;
}

void 
PduSessResModifyRespCodec::setUserLocationInfo(UserLocationInfo& info)
{
    userLocInfoPrst_m = true;
    userLocationInfo_m = info;
}

UserLocationInfo&
PduSessResModifyRespCodec::getUserLocationInfo()
{
    return userLocationInfo_m;
}

bool
PduSessResModifyRespCodec::isUserLocInfoPrst()
{
    return userLocInfoPrst_m;
}

bool 
PduSessResModifyRespCodec::reset()
{
    msgPdu_m.reset_v();

    amfUeNgapId_m = 0;
    ranUeNgapId_m = 0;
    
    pduSessResModRespList_m.clear();
	pduSessResModRespListPrst_m = false;
    pduSessResFailedModList_m.clear();
	pduSessResFailModRespListPrst_m = false;

    memset(&userLocationInfo_m,0xff,sizeof(userLocationInfo_m));
    userLocInfoPrst_m = false;
}


//bool PduSessResModifyRespCodec::encode(NgapOssCtxt* ctxt, MsgBuffer msgBuffer)
MsgBuffer PduSessResModifyRespCodec::encode(NgapOssCtxt *ctxt)
{
    MsgBuffer msgBuffer;
    msgBuffer.length = 0;
    msgBuffer.value = NULL;
    
    // encode RAN UE NGAP ID
    RAN_UE_NGAP_ID *ranUeIdPtr = msgPdu_m.getRanUeNgapIdEncodePtr();
    *ranUeIdPtr = ranUeNgapId_m;
    
    // encode AMF UE NGAP ID
    AMF_UE_NGAP_ID *amfUeIdPtr = msgPdu_m.getAmfUeNgapIdEncodePtr();
    *amfUeIdPtr = amfUeNgapId_m;

    // encode pduSessResModRespList_m
    if (pduSessResModRespListPrst_m) 
	{
	    int num = pduSessResModRespList_m.size();
        msgPdu_m.setNumofPduSessResModifyItemModRes_v(num); 
        for (int i = 0; i < num; i++) {
            PDUSessionResourceModifyItemModRes *pduPtr =
            msgPdu_m.getPduSessResModifyItemModResEncodePtr(i);
            if (NULL == pduPtr) {
                return msgBuffer;
            }
            PduSessResModifyRespItem* pduSessItem = &(pduSessResModRespList_m[i]);
      
            //psi
            pduPtr->pDUSessionID = pduSessItem->pduSessId;
            pduPtr->pDUSessionResourceModifyResponseTransfer.encoded.length
                = pduSessItem->pduSessResMdfyRespTransfer.length();
            memcpy(pduPtr->pDUSessionResourceModifyResponseTransfer.encoded.value,
                   pduSessItem->pduSessResMdfyRespTransfer.c_str(),
                   pduPtr->pDUSessionResourceModifyResponseTransfer.encoded.length);
        }
	}
    // encode pdu session resource failed to setup list
    if (pduSessResFailModRespListPrst_m) 
	{
        int num = pduSessResFailedModList_m.size();
        msgPdu_m.setNumofFailedToModifyItem_v(num);
        for (int i = 0; i < num; i++) {
            PDUSessionResourceFailedToModifyItemModRes *pduPtr =
            msgPdu_m.getPduSessResFailedModifyItemModResEncodePtr(i);
            if (NULL == pduPtr) {
                return msgBuffer;
            }
        PduSessResFailedMdfyRespItem* pduSessFailItem = &(pduSessResFailedModList_m[i]);
      
        //psi
        pduPtr->pDUSessionID = pduSessFailItem->pduSessId;
        pduPtr->pDUSessionResourceModifyUnsuccessfulTransfer.encoded.length
            = pduSessFailItem->pduSessResFailedMdfyRespTransfer.length();
        memcpy(pduPtr->pDUSessionResourceModifyUnsuccessfulTransfer.encoded.value,
               pduSessFailItem->pduSessResFailedMdfyRespTransfer.c_str(),
               pduPtr->pDUSessionResourceModifyUnsuccessfulTransfer.encoded.length);
        }
	}
	

    //Encode Uer Location Information
    if(userLocInfoPrst_m){
        UserLocationInformation* uliPtr=msgPdu_m.getUseLocationInformationEncodePtr();

        uliPtr->u.userLocationInformationNR.nR_CGI.pLMNIdentity.length =
            sizeof(userLocationInfo_m.nRCgi.plmn);
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
           uliPtr->u.userLocationInformationNR.bit_mask |= UserLocationInformationNR_timeStamp_present;
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
    
    
    msgPdu_m.encodeOssToPer(ctxt);
    
    OssBuf *perBuf = msgPdu_m.getPerBufRef();
    
    msgBuffer.length = perBuf->length;
    msgBuffer.value = perBuf->value;
    
    return msgBuffer;

}

bool 
PduSessResModifyRespCodec::decode(NgapOssCtxt *ctxt, MsgBuffer msgBuf)
{
    reset();
    
    OssBuf ossbuf;
    ossbuf.length = msgBuf.length;
    ossbuf.value = msgBuf.value;
    
    msgPdu_m.setPerBufRef_v(ossbuf);
    
    if (0 != msgPdu_m.decodePerToOss(ctxt)) {
        return false;
    }
    
    // decode RAN UE NGAP ID
    ranUeNgapId_m = *(msgPdu_m.getRanUeNgapIdDecodedPtr());
    
    // decode amf ue ngap id
    amfUeNgapId_m = *(msgPdu_m.getAmfUeNgapIdDecodedPtr());
  
    // decode PDUSessionResourceModifyListModRes
    for (int i = 0; i < msgPdu_m.getNumofPduSessResModifyItemModRes();i++) {
        PDUSessionResourceModifyItemModRes *pduPtr =
            msgPdu_m.getPduSessResModifyItemModResDecodedPtr(i);
        if (NULL == pduPtr) {
          return false;
        }
        
        PduSessResModifyRespItem item;
        //psi
        item.pduSessId =pduPtr->pDUSessionID;
        //if (pduPtr->bit_mask & pDUSessionResourceModifyResponseTransfer_present) {
            item.pduSessResMdfyRespTransferPrst = true;
            //pdu session resource modify response transfer
            OssBuf outputBuf = ctxt->getSecondaryCodecBuf();
            PDUSessionResourceModifyResponseTransfer* transferPdu;
            transferPdu = pduPtr->pDUSessionResourceModifyResponseTransfer.decoded;
            unsigned int result =
                ctxt->encodeOssToPer(PDUSessionResourceModifyResponseTransfer_PDU, transferPdu, &outputBuf);
            if (result != 0){
                std::cout << "failed to encode transfer"<< std::endl;
                return false;
            }
            item.pduSessResMdfyRespTransfer.assign(
                (char *)(outputBuf.value),
                outputBuf.length);
        //}

        pduSessResModRespList_m.push_back(item);
		pduSessResModRespListPrst_m = true;
    }

    // decode PDUSessionResourceFailedToModifyListModRes
    for (int i = 0; i < msgPdu_m.getNumofFailedToModifyItem();i++) {
        PDUSessionResourceFailedToModifyItemModRes *pduPtr =
              msgPdu_m.getPduSessResFailedModifyItemModResDecodedPtr(i);
        if (NULL == pduPtr) {
            return false;
        }
        PduSessResFailedMdfyRespItem item;
        //psi
        item.pduSessId =pduPtr->pDUSessionID;

        //transfer
        OssBuf outputBuf = ctxt->getSecondaryCodecBuf();
        
        PDUSessionResourceModifyUnsuccessfulTransfer* transferPdu;
        transferPdu = pduPtr->pDUSessionResourceModifyUnsuccessfulTransfer.decoded;
        unsigned int result =
              ctxt->encodeOssToPer(PDUSessionResourceModifyUnsuccessfulTransfer_PDU, transferPdu, &outputBuf);
        if (result != 0){
              std::cout << "failed to encode transfer"<< std::endl;
              return false;
        }
        item.pduSessResFailedMdfyRespTransfer.assign(
                (char *)(outputBuf.value),
                outputBuf.length);
      
        pduSessResFailedModList_m.push_back(item);
		pduSessResFailModRespListPrst_m = true;
    }

    //decode User Location Information
    if(msgPdu_m.isUseLocationInformationPresent()){
        if (msgPdu_m.getUseLocationInformationDecodedPtr()->choice == userLocationInformationNR_chosen)
        {
           Uint32 length=msgPdu_m.getUseLocationInformationDecodedPtr()->u.userLocationInformationNR.nR_CGI.pLMNIdentity.length;
           memcpy(userLocationInfo_m.nRCgi.plmn,
                  msgPdu_m.getUseLocationInformationDecodedPtr()->u.userLocationInformationNR.nR_CGI.pLMNIdentity.value,
                  length);
       
           //length=msgPdu_m.getUseLocationInformationDecodedPtr()->u.userLocationInformationNR.nR_CGI.nRCellIdentity.length;
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
   
           if(msgPdu_m.getUseLocationInformationDecodedPtr()->u.userLocationInformationNR.bit_mask &
            UserLocationInformationNR_timeStamp_present)
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

        }
        userLocInfoPrst_m = true;
    }

    return true;
}


std::string PduSessResModifyRespCodec::toString()
{
    ostringstream oss;
    
    oss << "PduSessResModifyRespCodec:" << endl
        << "1.AMF UE NGAP ID:(" << amfUeNgapId_m << ")" << endl
        << "2.RAN UE NGAP ID:(" << ranUeNgapId_m << ")" << endl
        << "3.userLocationInformation~~~~"<<endl
        <<"nrcgi~~~~"<<endl
        <<"plmnid: "<<userLocationInfo_m.nRCgi.plmn[0]
        <<userLocationInfo_m.nRCgi.plmn[1]<<userLocationInfo_m.nRCgi.plmn[2]<<endl
        <<"nrcellId: "<<userLocationInfo_m.nRCgi.nrCellId[0]<<userLocationInfo_m.nRCgi.nrCellId[1]
        <<userLocationInfo_m.nRCgi.nrCellId[2]<<userLocationInfo_m.nRCgi.nrCellId[3]
        <<userLocationInfo_m.nRCgi.nrCellId[4]<<endl
        <<"tai:"
        <<"Tac:("<<userLocationInfo_m.tai.tac[0]
        <<userLocationInfo_m.tai.tac[1]
        <<userLocationInfo_m.tai.tac[2]<<")"<<endl
        <<"PLMNIdentity:("<<userLocationInfo_m.tai.plmn[0]
        <<userLocationInfo_m.tai.plmn[1]
        <<userLocationInfo_m.tai.plmn[2] <<")" <<endl
        <<"timeStamp:prst " <<userLocationInfo_m.isTimeStampPrst <<endl
        <<" value: "<<userLocationInfo_m.timeStamp[0]
        <<userLocationInfo_m.timeStamp[1]
        <<userLocationInfo_m.timeStamp[2]
        <<userLocationInfo_m.timeStamp[3]<<endl
        << "PDU Session Resource Modify Response List:" << endl;

//    vector<PduSessResModifyRespItem>::iterator it;
//    for(it = pduSessResModRespList_m.begin(); it<pduSessResModRespList_m.end(); it++)
//    {
//      oss<< "psi:" << int(it->pduSessId) << endl
//         << "transfer:" << it->pduSessResMdfyRespTransfer << endl;
//    }
    oss << "PDU Session Resource Failed To Modify Response List: " << endl;
    vector <PduSessResFailedMdfyRespItem>::iterator failIt;
    for(failIt = pduSessResFailedModList_m.begin(); failIt<pduSessResFailedModList_m.end(); failIt++)
    {
        oss << "psi:" << int(failIt->pduSessId) << endl;
    }
    return oss.str(); 

}


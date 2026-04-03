#include <iostream>
#include <sstream>
#include <string>

#include "pduSessionResourceReleaseResponseCodec.h"

using namespace std;

PduSessResRelRespCodec::PduSessResRelRespCodec() 
{
    reset();
}

PduSessResRelRespCodec::~PduSessResRelRespCodec()
{
}

void PduSessResRelRespCodec::setRanUeNgapId(Uint32 id) 
{
    ranUeNgapId_m = id;
}

Uint32 PduSessResRelRespCodec::getRanUeNgapId() 
{
    return ranUeNgapId_m;
}

void PduSessResRelRespCodec::setAmfUeNgapId(Uint64 id) 
{
    amfUeNgapId_m = id;
}

Uint64 PduSessResRelRespCodec::getAmfUeNgapId() 
{
    return amfUeNgapId_m;
}

void 
PduSessResRelRespCodec::setUserLocationInfo(UserLocationInfo& info)
{
    userLocInfoPrst_m = true;
    userLocationInfo_m = info;
}

UserLocationInfo&
PduSessResRelRespCodec::getUserLocationInfo()
{
    return userLocationInfo_m;
}

bool
PduSessResRelRespCodec::isUserLocInfoPrst()
{
    return userLocInfoPrst_m;
}

void 
PduSessResRelRespCodec::addPduSessResRelRespList(PduSessResRelRespItem& val) 
{
    pduSessResRelRespList_m.push_back(val);
}

PduSessResRelRespList& 
PduSessResRelRespCodec::getPduSessResRelRespList() 
{
    return pduSessResRelRespList_m;
}

bool PduSessResRelRespCodec::reset() 
{
    msgPdu_m.reset_v();

    amfUeNgapId_m = 0;
    ranUeNgapId_m = 0;

    memset(&userLocationInfo_m,0xff,sizeof(userLocationInfo_m));
    userLocInfoPrst_m = false;

    pduSessResRelRespList_m.clear();
}

//bool PduSessResRelRespCodec::encode(NgapOssCtxt* ctxt, MsgBuffer msgBuffer)
MsgBuffer PduSessResRelRespCodec::encode(NgapOssCtxt *ctxt)
{
    MsgBuffer msgBuffer;
    msgBuffer.length = 0;
    msgBuffer.value = NULL;

    // encode RAN UE NGAP ID
    RAN_UE_NGAP_ID* ranUeIdPtr = msgPdu_m.getRanUeNgapIdEncodePtr();
    *ranUeIdPtr = ranUeNgapId_m;

    // encode AMF UE NGAP ID
    AMF_UE_NGAP_ID* amfUeIdPtr = msgPdu_m.getAmfUeNgapIdEncodePtr();
    *amfUeIdPtr = amfUeNgapId_m;

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

    //encode PDUSessionResourceReleasedListRelRes 
    Uint32 num = pduSessResRelRespList_m.size();
    msgPdu_m.setNumofPduSessResListItemRelResp_v(num);

    for (int i = 0; i < num; i++) {
        PDUSessionResourceReleasedItemRelRes *pduPtr = 
            msgPdu_m.getPduSessResItemRelRespEncodePtr(i);
        if (NULL == pduPtr) {
            return msgBuffer;
        }

        PduSessResRelRespItem *itemPtr = &(pduSessResRelRespList_m[i]);
        
        pduPtr->pDUSessionID = itemPtr->pduSessId;

        //pdu session resource setup transfer
        pduPtr->pDUSessionResourceReleaseResponseTransfer.encoded.length
             = itemPtr->pduSessResRelRespTransfer.length();
        memcpy(pduPtr->pDUSessionResourceReleaseResponseTransfer.encoded.value,
               itemPtr->pduSessResRelRespTransfer.c_str(),
               itemPtr->pduSessResRelRespTransfer.length());
    }

    msgPdu_m.encodeOssToPer(ctxt);

    OssBuf* perBuf = msgPdu_m.getPerBufRef();

    msgBuffer.length = perBuf->length;
    msgBuffer.value = perBuf->value;

    return msgBuffer;
}

bool PduSessResRelRespCodec::decode(NgapOssCtxt* ctxt,MsgBuffer msgBuf) 
{
    reset();

    OssBuf ossbuf;
    ossbuf.length = msgBuf.length;
    ossbuf.value = msgBuf.value;

    msgPdu_m.setPerBufRef_v(ossbuf);

    if (0 != msgPdu_m.decodePerToOss(ctxt)) {
        cout << "Failed to decode Per to OSS" << endl;
        return false;
    }
    // decode RAN UE NGAP ID
    ranUeNgapId_m = *(msgPdu_m.getRanUeNgapIdDecodedPtr());
    // decode amf ue ngap id
    amfUeNgapId_m = *(msgPdu_m.getAmfUeNgapIdDecodedPtr());

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

    //decode PDUSessionResourceReleasedListRelRes
    for (int i = 0; i < msgPdu_m.getNumofPduSessResListItemRelResp();i++) {
        PDUSessionResourceReleasedItemRelRes *pduPtr =
            msgPdu_m.getPduSessResItemRelRespDecodedPtr(i);
        if (NULL == pduPtr) {
            return false;
        }
          
        PduSessResRelRespItem item;
        item.pduSessId = pduPtr->pDUSessionID;
        // Pdu Session Resource Release Response Transfer
        OssBuf outputBuf = ctxt->getSecondaryCodecBuf();
        PDUSessionResourceReleaseResponseTransfer* transferPdu;
        transferPdu = pduPtr->pDUSessionResourceReleaseResponseTransfer.decoded;
        unsigned int result =
            ctxt->encodeOssToPer(PDUSessionResourceReleaseResponseTransfer_PDU, transferPdu, &outputBuf);
        if (result != 0){
            std::cout << "failed to encode transfer"<< std::endl;
            return false;
        }
        item.pduSessResRelRespTransfer.assign(
            (char *)(outputBuf.value),
            outputBuf.length);
    
        pduSessResRelRespList_m.push_back(item);
    }
  return true;
}

string PduSessResRelRespCodec::toString(){
   ostringstream oss;

   oss << "PduSessResRelRespCodec:" << endl
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
       << "PDU Session Resource Release Response List:" << endl;

    vector<PduSessResRelRespItem>::iterator it;
    for(it = pduSessResRelRespList_m.begin(); it<pduSessResRelRespList_m.end(); it++)
    {
      oss<< "psi:" << int(it->pduSessId) << endl
         << "transfer:" << it->pduSessResRelRespTransfer << endl;
    }
   return oss.str(); 
}
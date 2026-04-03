#include <iostream>
#include <sstream>
#include <string>

#include "uplinkNASTransportCodec.h"

using namespace std;

UplinkNASTransportCodec::UplinkNASTransportCodec()
{
    reset();
}

UplinkNASTransportCodec::~UplinkNASTransportCodec()
{
}

void 
UplinkNASTransportCodec::setRanUeNgapId(Uint32 id)
{
  ranUeNgapId_m = id;
}

Uint32 
UplinkNASTransportCodec::getRanUeNgapId()
{
  return ranUeNgapId_m;
}

void 
UplinkNASTransportCodec::setAmfUeNgapId(Uint64 id)
{
  amfUeNgapId_m = id;
}

Uint64 
UplinkNASTransportCodec::getAmfUeNgapId()
{
  return amfUeNgapId_m;
}


void 
UplinkNASTransportCodec::setNasPdu(std::string naspdu)
{
  nasPdu_m = naspdu;
}

std::string
UplinkNASTransportCodec::getNasPdu()
{
  return nasPdu_m;
}

void 
UplinkNASTransportCodec::setUserLocationInfo(UserLocationInfo& info)
{
    userLocationInfo_m = info;
}

UserLocationInfo&
UplinkNASTransportCodec::getUserLocationInfo()
{
    return userLocationInfo_m;
}


bool UplinkNASTransportCodec::reset()
{
  msgPdu_m.reset_v();

  amfUeNgapId_m = 0;
  ranUeNgapId_m = 0;
  nasPdu_m.clear();
  memset(&userLocationInfo_m,0xff,sizeof(userLocationInfo_m));
}

//bool UplinkNASTransportCodec::encode(NgapOssCtxt* ctxt, MsgBuffer msgBuffer)
MsgBuffer UplinkNASTransportCodec::encode(NgapOssCtxt *ctxt)
{
    MsgBuffer msgBuffer;
    msgBuffer.length = 0;
    msgBuffer.value = NULL;

    //encode RAN UE NGAP ID
    RAN_UE_NGAP_ID* ranUeIdPtr=msgPdu_m.getRanUeNgapIdEncodePtr();
   *ranUeIdPtr= ranUeNgapId_m;

    //encode AMF UE NGAP ID
    AMF_UE_NGAP_ID* amfUeIdPtr=msgPdu_m.getAmfUeNgapIdEncodePtr();
   *amfUeIdPtr= amfUeNgapId_m;
  
    //encode NAS PDU
    NAS_PDU* nasPduPtr=msgPdu_m.getNasPduEncodePtr();
    nasPduPtr->length = nasPdu_m.length();
    memcpy(nasPduPtr->value, nasPdu_m.c_str(), nasPduPtr->length); 

    //Encode Uer Location Information
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

    msgPdu_m.encodeOssToPer(ctxt);

    OssBuf* perBuf = msgPdu_m.getPerBufRef();

    msgBuffer.length = perBuf->length;
    msgBuffer.value = perBuf->value;

    return msgBuffer;
}

bool UplinkNASTransportCodec::decode(NgapOssCtxt* ctxt, MsgBuffer msgBuf)
{
    reset();

    OssBuf ossbuf;
    ossbuf.length = msgBuf.length;
    ossbuf.value = msgBuf.value;
    
    msgPdu_m.setPerBufRef_v(ossbuf);

    if (0 != msgPdu_m.decodePerToOss(ctxt)){
       cout << "Failed to decode Per to OSS" << endl;
       return false;
    }

    //decode RAN UE NGAP ID
    ranUeNgapId_m=*(msgPdu_m.getRanUeNgapIdDecodedPtr());

    //decode amf ue ngap id
    amfUeNgapId_m=*(msgPdu_m.getAmfUeNgapIdDecodedPtr());

    //decode NAS PDU
    NAS_PDU* nasPduPtr=msgPdu_m.getNasPduDecodedPtr();
    nasPdu_m.assign((char*)(nasPduPtr->value),nasPduPtr->length);

    //decode User Location Information
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
   
       if(msgPdu_m.getUseLocationInformationDecodedPtr()->u.userLocationInformationNR.bit_mask & UserLocationInformationNR_timeStamp_present)
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
    else
    {
       //don't supported yet
       cout <<"don't supported yet."<<endl;
    }

    return true;
}

string UplinkNASTransportCodec::toString()
{
    ostringstream oss;

    oss <<"UplinkNASTransportCodec:" << endl
        <<"RAN UE NGAP ID:(" << ranUeNgapId_m << ")"<<endl
        <<"AMF UE NGAP ID:(" << amfUeNgapId_m << ")"<<endl
        <<"NAS PDU:("<<nasPdu_m <<")"<< endl
        <<"userLocationInformation~~~~"<<endl
        <<"nrcgi:"<<endl
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
        <<"timeStamp:prst " <<userLocationInfo_m.isTimeStampPrst 
        <<" value: "<<userLocationInfo_m.timeStamp[0]
        <<userLocationInfo_m.timeStamp[1]
        <<userLocationInfo_m.timeStamp[2]
        <<userLocationInfo_m.timeStamp[3]<<endl;

   return oss.str();
}



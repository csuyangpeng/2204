#include <iostream>
#include <sstream>
#include <string>

#include "initialUEMessageCodec.h"

using namespace std;

InitialUEMessageCodec::InitialUEMessageCodec()
{
    reset();
}

InitialUEMessageCodec::~InitialUEMessageCodec()
{
}

void 
InitialUEMessageCodec::setRanUeNgapId(Uint32 id)
{
    ranUeNgapId_m = id;
}

Uint32 
InitialUEMessageCodec::getRanUeNgapId()
{
    return ranUeNgapId_m;
}

void 
InitialUEMessageCodec::setNasPdu(std::string naspdu)
{
    nasPdu_m = naspdu;
}

std::string
InitialUEMessageCodec::getNasPdu()
{
    return nasPdu_m;
}

void 
InitialUEMessageCodec::setUserLocationInfo(UserLocationInfo& info)
{
    userLocationInfo_m = info;
}

UserLocationInfo&
InitialUEMessageCodec::getUserLocationInfo()
{
    return userLocationInfo_m;
}

void
InitialUEMessageCodec::setRrcEstablishmentCause(Uint8 cause)
{
    rrcEstablishmentCause_m = cause;
}

Uint8 
InitialUEMessageCodec::getRrcEstablishmentCause()
{
    return rrcEstablishmentCause_m;
}

void
InitialUEMessageCodec::setAmfSetId(Uint8* Id)
{
    isAmfSetIdPrst_m = true;
    memcpy(amfSetId_m,Id,2);
}

Uint8*
InitialUEMessageCodec::getAmfSetId()
{
    return amfSetId_m;
}

bool
InitialUEMessageCodec::isAmfSetPresent()
{
    return isAmfSetIdPrst_m;
}

void 
InitialUEMessageCodec::set5gStmsi(Stmsi5G& stmsi)
{
    isStmsi5gPrst_m = true;
    stmsi_m = stmsi;
}
Stmsi5G&
InitialUEMessageCodec::get5gStmsi()
{
    return stmsi_m;
}
bool
InitialUEMessageCodec::is5gStmsiPresent()
{
    return isStmsi5gPrst_m;
}

void 
InitialUEMessageCodec::setUeContextRequest(Uint16 ueCtxtReq)
{
    isUeContextRequestPrst_m = true;
    ueContextRequest_m = ueCtxtReq;
}

Uint16 
InitialUEMessageCodec::getUeContextRequest()
{
    return ueContextRequest_m;
}
bool 
InitialUEMessageCodec::isUeContextRequestPresent()
{
    return isUeContextRequestPrst_m;
}

void 
InitialUEMessageCodec::addAllowedNssai(SNssai& nssai)
{
    isAllowedNssaiPrst_m = true;
    allowedNssaiList_m.push_back(nssai);
}

AllowedNssaiList&
InitialUEMessageCodec::getAllowedNssaiList()
{
    return allowedNssaiList_m;
}
bool
InitialUEMessageCodec::isAllowedNssaiPresent()
{
    return isAllowedNssaiPrst_m;
}

bool InitialUEMessageCodec::reset()
{
    msgPdu_m.reset_v();

    ranUeNgapId_m = 0;
    nasPdu_m.clear();
    memset(&userLocationInfo_m,0xff,sizeof(userLocationInfo_m));

    rrcEstablishmentCause_m=0;

    memset(&stmsi_m,0xff,sizeof(stmsi_m));
    isStmsi5gPrst_m = false;
    
    memset(&amfSetId_m,0xff,sizeof(amfSetId_m));
    isAmfSetIdPrst_m = false;

    ueContextRequest_m = 0;
    isUeContextRequestPrst_m = false;

    allowedNssaiList_m.clear();
    isAllowedNssaiPrst_m = false;

    return true;
}

//bool InitialUEMessageCodec::encode(NgapOssCtxt* ctxt, MsgBuffer msgBuffer)
MsgBuffer InitialUEMessageCodec::encode(NgapOssCtxt* ctxt)
{
    MsgBuffer msgBuffer;
    msgBuffer.length = 0;
    msgBuffer.value = NULL;

    //encode RAN UE NGAP ID
    RAN_UE_NGAP_ID* ranUeIdPtr=msgPdu_m.getRanUeNgapIdEncodePtr();
    *ranUeIdPtr= ranUeNgapId_m;

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

    //encode RRC Establishment Cause
    RRCEstablishmentCause* rrcEstabCausePtr=msgPdu_m.getRrcEstablishmentCauseEncodePtr();
    *rrcEstabCausePtr = (RRCEstablishmentCause)rrcEstablishmentCause_m;
    
    //encode AMF Set ID
    if(isAmfSetIdPrst_m)
    {
       AMFSetID* amfSetIdPtr=msgPdu_m.getAmfSetIdEncodePtr();
       amfSetIdPtr->length = amfSetIdBitLen;
       memcpy(amfSetIdPtr->value, amfSetId_m, sizeof(amfSetId_m)); 
    }

    //encode UE Context Request
    if(isUeContextRequestPrst_m)
    {
       UEContextRequest* ueContextRequestPtr=msgPdu_m.getUeContextRequestEncodePtr();
       *ueContextRequestPtr= UEContextRequest(ueContextRequest_m);
    }

    //5G Stmsi
    if(isStmsi5gPrst_m)
    {
       FiveG_S_TMSI* tmsi5gPtr=msgPdu_m.get5gStmsiEncodePtr();
       tmsi5gPtr->aMFSetID.length=amfSetIdBitLen;
       memcpy(tmsi5gPtr->aMFSetID.value,stmsi_m.setId,sizeof(stmsi_m.setId));

       tmsi5gPtr->aMFPointer.length=amfPointerBitLen;
       memcpy(tmsi5gPtr->aMFPointer.value,&stmsi_m.pointer,sizeof(stmsi_m.pointer));

       tmsi5gPtr->fiveG_TMSI.length = sizeof(stmsi_m.tmsi);
       memcpy(tmsi5gPtr->fiveG_TMSI.value, stmsi_m.tmsi, tmsi5gPtr->fiveG_TMSI.length);
    }
    
    //Encode Allowed NSSAI
    if(isAllowedNssaiPrst_m)
    {
       unsigned int numofAllowedItems=allowedNssaiList_m.size();
       msgPdu_m.setNumofAllowedNssaiItem_v(numofAllowedItems);
       
       for(int i=0;i<numofAllowedItems;i++)
       {
           AllowedNSSAI_Item* itemPtr = msgPdu_m.getAllowedNssaiItemEncodePtr(i);
           if(itemPtr != NULL)
           {
              itemPtr->s_NSSAI.sST.length= sstSize_c;
              memcpy(itemPtr->s_NSSAI.sST.value, allowedNssaiList_m[i].sst, sstSize_c);

              if(allowedNssaiList_m[i].sdPresent)
              {
                 itemPtr->s_NSSAI.bit_mask = sD_present;
                 itemPtr->s_NSSAI.sD.length= sdSize_c;
                 memcpy(itemPtr->s_NSSAI.sD.value, allowedNssaiList_m[i].sd, sdSize_c);
              }
           }
           else
           {
               //logger
               std::cout << "failed to getAllowedNssai_ItemEncodePtr" << std::endl;
           }
       }
    }

    msgPdu_m.encodeOssToPer(ctxt);

    OssBuf* perBuf = msgPdu_m.getPerBufRef();

    msgBuffer.length = perBuf->length;
    msgBuffer.value = perBuf->value;

    return msgBuffer;
}

bool InitialUEMessageCodec::decode(NgapOssCtxt* ctxt, MsgBuffer msgBuf)
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
   
       if(msgPdu_m.getUseLocationInformationDecodedPtr()->u.userLocationInformationNR.bit_mask == UserLocationInformationNR_timeStamp_present)
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

    //decode RRC Establishment Cause
    RRCEstablishmentCause* rrcEstablishCausePtr=msgPdu_m.getRrcEstablishmentCauseDecodedPtr();
    rrcEstablishmentCause_m=*rrcEstablishCausePtr;
    
    //decode AMF Set ID
    if(msgPdu_m.isAmfSetIdPresent())
    {
       AMFSetID* amfSetIdPtr=msgPdu_m.getAmfSetIdDecodedPtr();
       memcpy(amfSetId_m, amfSetIdPtr->value, amfSetIdLen);
       isAmfSetIdPrst_m = true;
    }
    
    //decode UE Context Request
    if(msgPdu_m.isUeContextRequestPresent())
    {
       ueContextRequest_m=*(msgPdu_m.getUeContextRequestDecodedPtr());
       isUeContextRequestPrst_m = true;
    }

    //decode FiveG S TMSI
    if(msgPdu_m.is5gStmsiPresent())
    {
       memcpy(stmsi_m.setId, 
            msgPdu_m.get5gStmsiDecodedPtr()->aMFSetID.value,
            amfSetIdLen);

       stmsi_m.setId[1] = stmsi_m.setId[1] >> 6;  // 6 LSB pad bits
       Uint8 temp = 0;
       temp = stmsi_m.setId[0] << 2 ;
       stmsi_m.setId[1] = stmsi_m.setId[1] | temp;
       stmsi_m.setId[0] = stmsi_m.setId[0] >> 6;

       stmsi_m.pointer=*(msgPdu_m.get5gStmsiDecodedPtr()->aMFPointer.value); //uint8
       stmsi_m.pointer = stmsi_m.pointer >> 2;  // 2 LSB pad bits

       memcpy(stmsi_m.tmsi,
           msgPdu_m.get5gStmsiDecodedPtr()->fiveG_TMSI.value,
           msgPdu_m.get5gStmsiDecodedPtr()->fiveG_TMSI.length);

       isStmsi5gPrst_m = true;
    }
    
    //decode Allowed NSSAI
    if(msgPdu_m.isAllowedNssaiPresent())
    {
       for(int i=0; i<msgPdu_m.getNumofAllowedNssaiItem(); i++)
       {
           AllowedNSSAI_Item* allowedNssaiPtr=msgPdu_m.getAllowedNssaiItemDecodedPtr(i);
        
           SNssai snssai;
        
           memcpy(snssai.sst, allowedNssaiPtr->s_NSSAI.sST.value, allowedNssaiPtr->s_NSSAI.sST.length);
   
           if(allowedNssaiPtr->s_NSSAI.bit_mask == sD_present)
           {
              memcpy(snssai.sd, allowedNssaiPtr->s_NSSAI.sD.value, allowedNssaiPtr->s_NSSAI.sD.length);
              snssai.sdPresent = true;
           }
        
           allowedNssaiList_m.push_back(snssai);
       }
       isAllowedNssaiPrst_m = true;
    }

    return true;
}

string InitialUEMessageCodec::toString()
{
    ostringstream oss;

    oss <<"InitialUEMessageCodec:" << endl
        <<"RAN UE NGAP ID:(" << ranUeNgapId_m << ")"<<endl
        << "NAS PDU:("<< nasPdu_m << ")" << endl
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
        <<userLocationInfo_m.timeStamp[3]<<endl
        <<"rrc estabilishment cause: "<<(int)rrcEstablishmentCause_m<<endl
        <<"5g s tmsi:prst " <<isStmsi5gPrst_m <<endl
        <<"setId: "<<stmsi_m.setId[0]<<stmsi_m.setId[1]<<endl
        <<"pointer: "<<stmsi_m.pointer<<endl
        <<"tmsi: "<<stmsi_m.tmsi[0]<<stmsi_m.tmsi[1]<<stmsi_m.tmsi[2]<<stmsi_m.tmsi[3]<<endl
        <<"amf set id:prst: "<<isAmfSetIdPrst_m <<" value " <<amfSetId_m[0] <<amfSetId_m[1]<<endl
        <<"ue context request:prst "<<isUeContextRequestPrst_m<< " value " <<(int)ueContextRequest_m<<endl
        <<"allow nssai list:prst "<<isAllowedNssaiPrst_m<< " value "<<endl;
        vector<SNssai>::iterator it;
        for (it=allowedNssaiList_m.begin(); it<allowedNssaiList_m.end() ; it++)
        {
            oss<< "        SD Flag: " << it->sdPresent << ", SD : " << it->sd[0]<< it->sd[1]<< it->sd[2]<<endl
               << "        SST :" << it->sst[0] << endl;
        }

   return oss.str();
}


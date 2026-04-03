#include <iostream>
#include <sstream>
#include <string>

#include "downlinkNASTransportCodec.h"

using namespace std;

DownlinkNASTransportCodec::DownlinkNASTransportCodec()
{
    reset();
}

DownlinkNASTransportCodec::~DownlinkNASTransportCodec()
{
}

void 
DownlinkNASTransportCodec::setRanUeNgapId(Uint32 id)
{
  ranUeNgapId_m = id;
}

Uint32 
DownlinkNASTransportCodec::getRanUeNgapId()
{
  return ranUeNgapId_m;
}

void 
DownlinkNASTransportCodec::setAmfUeNgapId(Uint64 id)
{
  amfUeNgapId_m = id;
}

Uint64 
DownlinkNASTransportCodec::getAmfUeNgapId()
{
  return amfUeNgapId_m;
}


void 
DownlinkNASTransportCodec::setNasPdu(std::string naspdu)
{
  nasPdu_m = naspdu;
}

std::string
DownlinkNASTransportCodec::getNasPdu()
{
  return nasPdu_m;
}

void 
DownlinkNASTransportCodec::setRanPagingPriority(Uint16 ranPagingPrt)
{
   ranPagingPriorityPresent_m = true;
   ranPagingPriority_m = ranPagingPrt;
}

Uint16
DownlinkNASTransportCodec::getRanPagingPriority()
{
  return ranPagingPriority_m;
}

bool 
DownlinkNASTransportCodec::isRanPagingPriorityPresent()
{
  return ranPagingPriority_m;
}


void
DownlinkNASTransportCodec::setOldAmfName(std::string name)
{
   oldAmfNamePresent_m = true;
   oldAmfName_m = name;
}

std::string 
DownlinkNASTransportCodec::getOldAmfName()
{
  return oldAmfName_m;
}

bool 
DownlinkNASTransportCodec::isOldAmfNamePresent()
{
  return oldAmfNamePresent_m;
}

void 
DownlinkNASTransportCodec::setIndexToRfsp(long index)
{
   indexToRfspPresent_m = true;
   indexToRfsp_m = index;
}

long
DownlinkNASTransportCodec::getIndexToRfsp()
{
  return indexToRfsp_m;
}

bool
DownlinkNASTransportCodec::isIndexToRfspPresent()
{
  return indexToRfspPresent_m;
}

void 
DownlinkNASTransportCodec::setUeAmbr(UeAmbr& ambr)
{
   ueAmbrPresent_m = true;
   ueAmbr_m = ambr;
}

UeAmbr&
DownlinkNASTransportCodec::getUeAmbr()
{
   return ueAmbr_m;
}

bool
DownlinkNASTransportCodec::isUeAmbrPresent()
{
   return ueAmbrPresent_m;
}

void 
DownlinkNASTransportCodec::addAllowedNssai(SNssai& nssai)
{
    allowedNssaiPrst_m = true;
    allowedNssaiList_m.push_back(nssai);
}

AllowedNssaiList& DownlinkNASTransportCodec::getAllowedNssaiList()
{
  return allowedNssaiList_m;
}

bool DownlinkNASTransportCodec::isAllowedNssaiPresent()
{
    return allowedNssaiPrst_m;
}

bool DownlinkNASTransportCodec::reset()
{
  msgPdu_m.reset_v();

  amfUeNgapId_m = 0;
  ranUeNgapId_m = 0;
  nasPdu_m.clear();
  
  oldAmfName_m.clear();
  oldAmfNamePresent_m = false;

  ranPagingPriority_m = 0;
  ranPagingPriorityPresent_m = false;

  indexToRfsp_m = 0;
  indexToRfspPresent_m = false;
  
  memset(&ueAmbr_m,0,sizeof(ueAmbr_m));
  ueAmbrPresent_m = false;
  
  allowedNssaiList_m.clear();
  allowedNssaiPrst_m = false;

  return true;
}

MsgBuffer DownlinkNASTransportCodec::encode(NgapOssCtxt* ctxt)
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

    //Encode old Amf Name
    if(oldAmfNamePresent_m)
    {
       AMFName* oldAmfNamePtr = msgPdu_m.getOldAmfNameEncodePtr();
       oldAmfNamePtr->length = oldAmfName_m.length();
       memcpy(oldAmfNamePtr->value, oldAmfName_m.c_str(), oldAmfNamePtr->length);
    }

    //encode ran paging priority
    if(ranPagingPriorityPresent_m)
    {
        RANPagingPriority* ranPagingPriorityPtr = msgPdu_m.getRanPagingPriorityEncodePtr();
        *ranPagingPriorityPtr = ranPagingPriority_m;
    }
  
    //encode index to rfsp
    if(indexToRfspPresent_m)
    {
       IndexToRFSP* index = msgPdu_m.getIndexToRfspEncodePtr();
       *index = indexToRfsp_m;
    }

    //encode Ue ambr
    if(ueAmbrPresent_m)
    {
       UEAggregateMaximumBitRate* ueAmbrPtr = msgPdu_m.getUeAmbrEncodePtr();
       ueAmbrPtr->uEAggregateMaximumBitRateDL = ueAmbr_m.downlink;
       ueAmbrPtr->uEAggregateMaximumBitRateUL = ueAmbr_m.uplink;
    }
  
    //Encode Allowed NSSAI
    if(allowedNssaiPrst_m)
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

bool DownlinkNASTransportCodec::decode(NgapOssCtxt* ctxt, MsgBuffer msgBuf)
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

    //decode old amf name
    if(msgPdu_m.isOldAmfNamePresent())
    {
       AMFName* amfNamePtr = msgPdu_m.getOldAmfNameDecodedPtr();
       oldAmfName_m.assign((char*)(amfNamePtr->value),amfNamePtr->length);
       oldAmfNamePresent_m = true;
    }
  
    //decode ran paging priority
    if(msgPdu_m.isRanPagingPriorityPresent())
    {
       RANPagingPriority* ranPagingPriorityPtr = msgPdu_m.getRanPagingPriorityDecodedPtr();
       ranPagingPriority_m = *ranPagingPriorityPtr;
       ranPagingPriorityPresent_m = true;
    }
  
    //decode index to rfsp
    if(msgPdu_m.isIndexToRfspPresent())
    {
       IndexToRFSP* indexPtr = msgPdu_m.getIndexToRfspDecodedPtr();
       indexToRfsp_m = *indexPtr;
       indexToRfspPresent_m = true;
    }

    //decode Ue Ambr
    if(msgPdu_m.isUeAmbrPresent())
    {
       ueAmbr_m.uplink = msgPdu_m.getUeAmbrDecodedPtr()->uEAggregateMaximumBitRateUL;
       ueAmbr_m.downlink = msgPdu_m.getUeAmbrDecodedPtr()->uEAggregateMaximumBitRateDL;
       ueAmbrPresent_m = true;
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

       allowedNssaiPrst_m = true;
    }

    return true;
}

string DownlinkNASTransportCodec::toString()
{
    ostringstream oss;

    oss <<"DownlinkNASTransportCodec:" << endl
        <<"RAN UE NGAP ID:(" << ranUeNgapId_m << ")"<<endl
        <<"AMF UE NGAP ID:(" << amfUeNgapId_m << ")"<<endl
        <<"Old Amf Name, prst(" << oldAmfNamePresent_m <<") value(" <<oldAmfName_m <<")" << endl
        <<"NasPdu: "<<nasPdu_m<<endl
        <<"RanPagingPriority, prst(" << ranPagingPriorityPresent_m <<") value(" << ranPagingPriority_m <<")" << endl
        <<"IndexToRfsp, prst(" << indexToRfspPresent_m <<") value(" << indexToRfsp_m <<")" << endl
        <<"UeAmbr, prst(" << ueAmbrPresent_m <<") value( downlink - " << ueAmbr_m.downlink <<", uplink - " << ueAmbr_m.uplink << ")" << endl
		<<"allow nssai list:prst "<<allowedNssaiPrst_m<< " value "<<endl;
		vector<SNssai>::iterator it;
		for (it=allowedNssaiList_m.begin(); it<allowedNssaiList_m.end() ; it++)
		{
			oss<< "		   SD Flag: " << it->sdPresent << ", SD : " << it->sd[0]<< it->sd[1]<< it->sd[2]<<endl
			   << "		   SST :" << it->sst[0] << endl;
		}
   return oss.str();
}


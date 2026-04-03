#include <iostream>
#include "ngSetupRequestCodec.h"

using namespace std;

NgSetupRequestCodec::NgSetupRequestCodec()
{
    reset();
}

NgSetupRequestCodec::~NgSetupRequestCodec()
{
}

void NgSetupRequestCodec::setRanNodeName(std::string name)
{
    rNodeNamePresent_m = true;
    ranNodeName_m = name;
}
bool NgSetupRequestCodec::getRNodeNamePrst()
{
    return rNodeNamePresent_m;
}
std::string NgSetupRequestCodec::getRanNodeName()
{
    return ranNodeName_m;
}

void NgSetupRequestCodec::setPagingDrx(Uint8 pdrx)
{
    pagingDrx_m = pdrx;
}
Uint8 NgSetupRequestCodec::getPagingDrx()
{
    return pagingDrx_m;
}

void NgSetupRequestCodec::setGgnbId(GGnbId& ggnbid)
{
    gGnbId_m = ggnbid;
}
GGnbId& NgSetupRequestCodec::getGgnbId()
{
    return gGnbId_m;
}

void NgSetupRequestCodec::addSupTAList(SupTAItem& stai)
{
    supTAList_m.push_back(stai);
}
SupTAList& NgSetupRequestCodec::getSupTAList()
{
    return supTAList_m;
}

void NgSetupRequestCodec::setUeRetentionInfo(Uint8 ueRetentionInfo)
{
    UeRetentionInfoPrst_m = true;
    UeRetentionInfo_m = ueRetentionInfo;
}
bool NgSetupRequestCodec::getUeRetentionInfoPrst()
{
    return UeRetentionInfoPrst_m;
}
Uint8 NgSetupRequestCodec::getUeRetentionInfo()
{
    return UeRetentionInfo_m;
}

bool NgSetupRequestCodec::reset()
{
    msgPdu_m.reset_v();

    rNodeNamePresent_m = false;
    ranNodeName_m.clear();

    pagingDrx_m = 0;

    memset(gGnbId_m.plmnid,0xff,sizeof(gGnbId_m.plmnid));
    gGnbId_m.gnbId.len = 0;
    memset(gGnbId_m.gnbId.val, 0xff, sizeof(gGnbId_m.gnbId.val));

    supTAList_m.clear();

	UeRetentionInfo_m = 0;
    return true;
}

//bool NgSetupRequestCodec::encode(NgapOssCtxt* ctxt, MsgBuffer msgBuffer)
MsgBuffer NgSetupRequestCodec::encode(NgapOssCtxt* ctxt)
{
    cout << "NGSetupRequestCodec::encode()" << endl;

    MsgBuffer msgBuffer;
    msgBuffer.length = 0;
    msgBuffer.value = NULL;

   //encode Global GNB ID
   GlobalRANNodeID* gNBIDPtr = msgPdu_m.getGlobalRanNodeIdEncodePtr();
   gNBIDPtr->choice = globalGNB_ID_chosen;
   gNBIDPtr->u.globalGNB_ID.bit_mask = 0;
   gNBIDPtr->u.globalGNB_ID.pLMNIdentity.length = sizeof(gGnbId_m.plmnid);
   memcpy(gNBIDPtr->u.globalGNB_ID.pLMNIdentity.value,
          gGnbId_m.plmnid,
          sizeof(gGnbId_m.plmnid));
   gNBIDPtr->u.globalGNB_ID.gNB_ID.choice = gNB_ID_chosen;
   gNBIDPtr->u.globalGNB_ID.gNB_ID.u.gNB_ID.length = gGnbId_m.gnbId.len;
   memcpy(gNBIDPtr->u.globalGNB_ID.gNB_ID.u.gNB_ID.value,gGnbId_m.gnbId.val, sizeof(gGnbId_m.gnbId.val));

   //encode Paging DRX
   PagingDRX * pagingDrxPtr = msgPdu_m.getPagingDrxEncodePtr();
   *pagingDrxPtr = (PagingDRX)(pagingDrx_m);

   //encode RanNodeName
   if(rNodeNamePresent_m == true)
   {
      RANNodeName* rnnPtr = msgPdu_m.getRanNodeNameEncodePtr();
      rnnPtr->length = ranNodeName_m.length();
      memcpy(rnnPtr->value, ranNodeName_m.c_str(), rnnPtr->length);
   }

   //encode Supported TA List
   Uint32 numofSupportedTAItems = supTAList_m.size();
   msgPdu_m.setNumOfSupportedTAListItem_v(numofSupportedTAItems);
   for(int i=0; i<numofSupportedTAItems; i++)
   {
      Uint32 numofBPlmnItems = supTAList_m[i].bplmnList.size();
      msgPdu_m.setNumOfBPlmns_v(numofBPlmnItems, i);

      for(int j=0; j< numofBPlmnItems; j++)
      {
        Uint32 numofSliceItems = supTAList_m[i].bplmnList[j].ssList.size();
        msgPdu_m.setNumOfSlices_v(numofSliceItems, i, j);
      }
   }

    for(int i = 0; i < numofSupportedTAItems; i++)
    {
       SupportedTAItem * taiItemPtr = msgPdu_m.getSupportedTAListItemEncodePtr(i);
       if (NULL == taiItemPtr)
       {
          std::cout << "NULL Pointer for taiItemPtr" << std::endl;
          return msgBuffer;
       }
       taiItemPtr->tAC.length = sizeof(supTAList_m[i].tac);
       memcpy(taiItemPtr->tAC.value, supTAList_m[i].tac, taiItemPtr->tAC.length);

       Uint32 numofBPlmnItems = supTAList_m[i].bplmnList.size();
       for(int j = 0; j < numofBPlmnItems; j++)
       {
          BroadcastPLMNItem* plmnItemPtr = msgPdu_m.getBPlmnsEncodePtr(i,j);
          if(NULL == plmnItemPtr)
          {
              std::cout << "NULL Pointer for plmnItemPtr" << std::endl;
              return msgBuffer;
          }
          Uint32 length = sizeof(supTAList_m[i].bplmnList[j].plmnid);
          plmnItemPtr->pLMNIdentity.length = length;
          memcpy(plmnItemPtr->pLMNIdentity.value, supTAList_m[i].bplmnList[j].plmnid, length);

          Uint32 numofSliceItems = supTAList_m[i].bplmnList[j].ssList.size();
          for(int k = 0; k < numofSliceItems; k++)
          {
             SliceSupportItem* sliceItemPtr = msgPdu_m.getSliceSupportListEncodePtr(i,j,k);
             if(NULL == sliceItemPtr)
             {
                 std::cout << "NULL Pointer for sliceItemPtr" << std::endl;
                 return msgBuffer;
             }

             bool sdPrst = supTAList_m[i].bplmnList[j].ssList[k].sdPresent;
             if(sdPrst == true)
             {
                sliceItemPtr->s_NSSAI.bit_mask = sD_present;
                Uint32 length = sizeof(supTAList_m[i].bplmnList[j].ssList[k].sd);
                sliceItemPtr->s_NSSAI.sD.length = length;
                memcpy(sliceItemPtr->s_NSSAI.sD.value, supTAList_m[i].bplmnList[j].ssList[k].sd, length);
             }

             Uint32 length = sizeof(supTAList_m[i].bplmnList[j].ssList[k].sst);
             sliceItemPtr->s_NSSAI.sST.length = length;
             memcpy(sliceItemPtr->s_NSSAI.sST.value, supTAList_m[i].bplmnList[j].ssList[k].sst, length);
          }
       }
    }

    //encode Ue Retention Information
    if(UeRetentionInfoPrst_m == true)
	{
	    UERetentionInformation * ueRetentionInfoPtr = msgPdu_m.getUeRetentionInfoEncodePtr();
		*ueRetentionInfoPtr = (UERetentionInformation)(UeRetentionInfo_m);
	}
    msgPdu_m.encodeOssToPer(ctxt);

    OssBuf* perBuf = msgPdu_m.getPerBufRef();

    msgBuffer.length = perBuf->length;
    msgBuffer.value = perBuf->value;

    return msgBuffer;
}

// Uint32 NgSetupRequestCodec::decodePerToOss(NgapOssCtxt* ctxt, MsgBuffer msgBuf)
// {
//     reset();

//     OssBuf ossbuf;
//     ossbuf.length = msgBuf.length;
//     ossbuf.value = msgBuf.value;
//     msgPdu_m.setPerBufRef_v(ossbuf);

//     return msgPdu_m.decodePerToOss(ctxt);
// }

bool NgSetupRequestCodec::decode(NgapOssCtxt* ctxt, MsgBuffer msgBuf)
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

    //decode global GNB ID
    Uint32 length = msgPdu_m.getGlobalRanNodeIdDecodedPtr()->u.globalGNB_ID.pLMNIdentity.length;

    memcpy(gGnbId_m.plmnid,
           msgPdu_m.getGlobalRanNodeIdDecodedPtr()->u.globalGNB_ID.pLMNIdentity.value,
           length);

    gGnbId_m.gnbId.len = msgPdu_m.getGlobalRanNodeIdDecodedPtr()->u.globalGNB_ID.gNB_ID.u.gNB_ID.length;
    memcpy(gGnbId_m.gnbId.val,
           msgPdu_m.getGlobalRanNodeIdDecodedPtr()->u.globalGNB_ID.gNB_ID.u.gNB_ID.value,
           sizeof(msgPdu_m.getGlobalRanNodeIdDecodedPtr()->u.globalGNB_ID.gNB_ID.u.gNB_ID.value));

  // decode rname node name
  rNodeNamePresent_m = msgPdu_m.isRanNodeNamePresent();
  if (true == rNodeNamePresent_m)
  {
      RANNodeName* rNNPtr = msgPdu_m.getRanNodeNameDecodedPtr();
      ranNodeName_m.assign(rNNPtr->value,rNNPtr->length);
  }

  // decode paging drx
  pagingDrx_m = *(msgPdu_m.getPagingDrxDecodedPtr());

  // decode supported TA List
   for(Uint32 i = 0; i < msgPdu_m.getNumOfSupportedTAListItem(); i++)
   {
      SupportedTAItem* taiItemPtr = msgPdu_m.getSupportedTAListItemDecodedPtr(i);
      if (NULL == taiItemPtr)
      {
         return false;
      }

      SupTAItem supTAItem_m;
      memcpy(supTAItem_m.tac,taiItemPtr->tAC.value,taiItemPtr->tAC.length);

      for(Uint32 j = 0; j < msgPdu_m.getNumOfBPlmns(i); j++)
      {
         BroadcastPLMNItem* bplmnlistPtr = msgPdu_m.getBPlmnsDecodedPtr(taiItemPtr, j);
         if (NULL == bplmnlistPtr)
         {
            return false;
         }

         BPlmnItem bplmnItem;
         memcpy(bplmnItem.plmnid,bplmnlistPtr->pLMNIdentity.value, bplmnlistPtr->pLMNIdentity.length);

         for(Uint32 k = 0; k < msgPdu_m.getNumOfSlices(i,j); k++)
         {
            SliceSupportItem* sliceItemPtr = msgPdu_m.getSliceSupportListDecodedPtr(bplmnlistPtr, k);
            if (NULL == sliceItemPtr)
            {
               return false;
            }

            SNssai snssai;
            memset(&snssai,0,sizeof(SNssai));
            if (sliceItemPtr->s_NSSAI.bit_mask == sD_present) //#       define      sD_present 0x80
            {
                snssai.sdPresent = true;
                memcpy(snssai.sd, sliceItemPtr->s_NSSAI.sD.value, sliceItemPtr->s_NSSAI.sD.length);
            }
            memcpy(snssai.sst, sliceItemPtr->s_NSSAI.sST.value, sliceItemPtr->s_NSSAI.sST.length);

            bplmnItem.ssList.push_back(snssai);
         }
         supTAItem_m.bplmnList.push_back(bplmnItem);
      }
      supTAList_m.push_back(supTAItem_m);
   }

  //decode Ue Retention Information
  UeRetentionInfoPrst_m = msgPdu_m.isUeRetentionInfoPresent();
  if(true == UeRetentionInfoPrst_m)
  {
      UeRetentionInfo_m = *(msgPdu_m.getUeRetentionInfoDecodedPtr());
  }

   return true;
}

void NgSetupRequestCodec::dumpMessage()
{
    cout <<"Dump NgSetupRequestCodec:" << endl;
    cout <<"ranNodeName: flag("<<rNodeNamePresent_m<<") value("<<ranNodeName_m<<")"<<endl;
    cout <<"pagingDrx: "<<(int)pagingDrx_m<<endl;
	cout <<"ue Retention Information: flag(" <<UeRetentionInfoPrst_m<<"),value("<<(int)UeRetentionInfo_m<<")"<<endl;
    cout <<"Goble GNB ID: "
         <<"  Plmn: 0x" << gGnbId_m.plmnid[0]<< gGnbId_m.plmnid[1]<< gGnbId_m.plmnid[2]<< endl;
    cout <<"  gnbId len(" << gGnbId_m.gnbId.len << ")  val (" << gGnbId_m.gnbId.val[0] << gGnbId_m.gnbId.val[1] << gGnbId_m.gnbId.val[2] << ") "<<endl;
    cout <<"Supported TAI List: "<< endl;
    vector<SupTAItem>::iterator it;
    for (it=supTAList_m.begin(); it<supTAList_m.end() ; it++)
    {
        cout << "  TAC: " << it->tac[0]<< it->tac[1]<< it->tac[2] << endl;
        cout <<"  BPlmnList: " << endl;
        BPlmnList bplmnList = it->bplmnList;
        vector<BPlmnItem>::iterator bIt;
        for(bIt=bplmnList.begin(); bIt<bplmnList.end(); bIt++)
        {
            cout <<"     plmnid : "<< bIt->plmnid[0]<< bIt->plmnid[1]<< bIt->plmnid[2]<<endl;
            SliceSupList ssList = bIt->ssList;
            vector<SNssai>::iterator it;
            for (it=ssList.begin(); it<ssList.end(); it++)
            {
                cout << "          SD Flag: " << it->sdPresent << ", SD : " << it->sd[0]<< it->sd[1]<< it->sd[2];
                cout << "          SST :" << it->sst[0] << endl;
            }
        }
    }
}


#include <iostream>
#include <sstream>
#include <string>

#include "ngSetupResponseCodec.h"

using namespace std;

NgSetupResponseCodec::NgSetupResponseCodec()
{
    reset();
}

NgSetupResponseCodec::~NgSetupResponseCodec()
{
}

void 
NgSetupResponseCodec::setAmfName(std::string name)
{
	amfName_m = name;
}
std::string 
NgSetupResponseCodec::getAmfName()
{
	return amfName_m;
}

void 
NgSetupResponseCodec::setRelativeAmfCapacity(Uint16 cap)
{
	relativeAmfCapacity_m = cap;
}
Uint16 
NgSetupResponseCodec::getRelativeAmfCapacity()
{
	return relativeAmfCapacity_m;
}

void 
NgSetupResponseCodec::addServedGuamiList(ServedGuamiItem& sguami)
{
	serGuamiList_m.push_back(sguami);	
}
ServedGuamiList& 
NgSetupResponseCodec::getServedGuamiList()
{
	return serGuamiList_m;
}

void 
NgSetupResponseCodec::addPlmnList(BPlmnItem& plmnItem)
{
	plmnSupportist_m.push_back(plmnItem);
}
BPlmnList& 
NgSetupResponseCodec::getPlmnList()
{
	return plmnSupportist_m;
}

void NgSetupResponseCodec::setUeRetentionInfo(Uint8 ueRetentionInfo)
{
    UeRetentionInfoPrst_m = true;
    UeRetentionInfo_m = ueRetentionInfo;
}
bool NgSetupResponseCodec::getUeRetentionInfoPrst()
{
    return UeRetentionInfoPrst_m;
}
Uint8 NgSetupResponseCodec::getUeRetentionInfo()
{
    return UeRetentionInfo_m;
}

    
bool NgSetupResponseCodec::reset()
{
   msgPdu_m.reset_v();

	amfName_m.clear();
	
	relativeAmfCapacity_m = 0;
	
	serGuamiList_m.clear();
	
	plmnSupportist_m.clear();

	UeRetentionInfo_m = 0;

    return true;
}

MsgBuffer NgSetupResponseCodec::encode(NgapOssCtxt* ctxt)
//bool NgSetupResponseCodec::encode(NgapOssCtxt* ctxt, MsgBuffer msgBuffer)
{
    MsgBuffer msgBuffer;
    msgBuffer.length = 0;
    msgBuffer.value = NULL;

   //encode AmfName
   AMFName* amfNamePtr = msgPdu_m.getAmfNameEncodePtr();
   amfNamePtr->length = amfName_m.length();
   memcpy(amfNamePtr->value, amfName_m.c_str(), amfNamePtr->length);   

   //encode Relative AMF capability
   RelativeAMFCapacity* amfCapPtr = msgPdu_m.getRelAmfCapacityEncodePtr();
   *amfCapPtr = relativeAmfCapacity_m;

   //encode ServedGuamiList
   //ServedGUAMIList* serGuamiListPtr = msgPdu_m.getServedGUAMIListEncodePtr();
   unsigned int numofSerGuamiItems = serGuamiList_m.size();
   msgPdu_m.setNumofServedGUAMIListItem_v(numofSerGuamiItems);
   for(int i=0;i<numofSerGuamiItems;i++)
   {
		ServedGUAMIItem* itemPtr = msgPdu_m.getServedGUAMIListItemEncodedPtr(i);
		if (itemPtr == NULL)
		{
			//logger
			std::cout << "failed to getServedGUAMIListItemEncodedPtr" << std::endl;
			return msgBuffer;
		}

		if(serGuamiList_m[i].backupAmfNamePrst)
		{
			itemPtr->bit_mask=ServedGUAMIItem_backupAMFName_present;
			itemPtr->backupAMFName.length = serGuamiList_m[i].backupAmfName.length();
			memcpy(itemPtr->backupAMFName.value, serGuamiList_m[i].backupAmfName.c_str(), itemPtr->backupAMFName.length); 
		}
		itemPtr->gUAMI.pLMNIdentity.length = plmnSize_c;
		memcpy(itemPtr->gUAMI.pLMNIdentity.value, serGuamiList_m[i].plmnId, plmnSize_c);

        //amf region encode
		itemPtr->gUAMI.aMFRegionID.length = amfReginIdBitLen;
		memcpy(itemPtr->gUAMI.aMFRegionID.value, &(serGuamiList_m[i].amfId.regionId), 
			sizeof(serGuamiList_m[i].amfId.regionId));

        //set id encode, 10 bit, the second byte should left shift 6 bit
		itemPtr->gUAMI.aMFSetID.length = amfSetIdBitLen;
                serGuamiList_m[i].amfId.setId[0] = (serGuamiList_m[i].amfId.setId[0] << 6) | (serGuamiList_m[i].amfId.setId[1] >> 2);
                serGuamiList_m[i].amfId.setId[1] = serGuamiList_m[i].amfId.setId[1] << 6;
        //serGuamiList_m[i].amfId.setId[1] = serGuamiList_m[i].amfId.setId[1] << 6;
	    memcpy(itemPtr->gUAMI.aMFSetID.value, serGuamiList_m[i].amfId.setId,
			   sizeof(serGuamiList_m[i].amfId.setId));

		itemPtr->gUAMI.aMFPointer.length = amfPointerBitLen;
		serGuamiList_m[i].amfId.pointer = serGuamiList_m[i].amfId.pointer << 2;
		memcpy(itemPtr->gUAMI.aMFPointer.value,&(serGuamiList_m[i].amfId.pointer),
				sizeof(serGuamiList_m[i].amfId.pointer));
   }

   //encode PlmnList
   Uint32 numofPlmnItems = plmnSupportist_m.size();
   msgPdu_m.setNumofPlmnSupportListItem_m_v(numofPlmnItems);
   for(int i=0; i<numofPlmnItems; i++)
   {
      Uint32 numofSliceItems = plmnSupportist_m[i].ssList.size();
      msgPdu_m.setNumofSlices_v(numofSliceItems, i);
   }

   for(int i = 0; i < numofPlmnItems; i++)
   {
	   PLMNSupportItem * plmnSupItemPtr = msgPdu_m.getPlmnSupportItemEncodePtr(i);
       if (NULL == plmnSupItemPtr)
       {
          //TODO, logger here
          std::cout << "failed to getPlmnSupportItemEncodePtr(" <<i<<")"<< std::endl;
          return msgBuffer;
       }

	   plmnSupItemPtr->pLMNIdentity.length = plmnSize_c;
	   memcpy(plmnSupItemPtr->pLMNIdentity.value, plmnSupportist_m[i].plmnid,
	   	      sizeof(plmnSupportist_m[i].plmnid));

	   Uint32 numofSliceItems = plmnSupportist_m[i].ssList.size();
       for(int j = 0 ; j < numofSliceItems; j++)
       {
          SliceSupportItem* sliceItemPtr = msgPdu_m.getSliceSupportListEncodePtr(i,j);
          if(NULL == sliceItemPtr)
          {
             //TODO， logger here
             std::cout << "NULL Pointer for sliceItemPtr" << std::endl;
             return msgBuffer;
          }
             
          bool sdPrst = plmnSupportist_m[i].ssList[j].sdPresent;
          if(sdPrst == true)
          {
                sliceItemPtr->s_NSSAI.bit_mask = sD_present;
                Uint32 length = sizeof(plmnSupportist_m[i].ssList[j].sd);
                sliceItemPtr->s_NSSAI.sD.length = length;
                memcpy(sliceItemPtr->s_NSSAI.sD.value, plmnSupportist_m[i].ssList[j].sd, length);
          }
             
          Uint32 length = sizeof(plmnSupportist_m[i].ssList[j].sst);
          sliceItemPtr->s_NSSAI.sST.length = length;
          memcpy(sliceItemPtr->s_NSSAI.sST.value, plmnSupportist_m[i].ssList[j].sst, length);          
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

bool NgSetupResponseCodec::decode(NgapOssCtxt* ctxt, MsgBuffer msgBuf)
{
	
    reset();

    OssBuf ossbuf;
    ossbuf.length = msgBuf.length;
    ossbuf.value = msgBuf.value;
    
    msgPdu_m.setPerBufRef_v(ossbuf);
	std::cout << "Codec decode test~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~"<< std::endl;
    if (0 != msgPdu_m.decodePerToOss(ctxt)){
		cout << "Failed to decode Per to OSS" << endl;
		return false;
	}

    //decode AMF Name
	AMFName* amfNamePtr = msgPdu_m.getAmfNameDecodedPtr();
	amfName_m.assign(amfNamePtr->value,amfNamePtr->length);

    //decode Relative AMF capability
	relativeAmfCapacity_m = *(msgPdu_m.getRelAmfCapacityDecodedPtr());

    //decode ServedGuamiList
    for(Uint32 i = 0; i < msgPdu_m.getNumofServedGUAMIListItem(); i++)
    {
       ServedGUAMIItem* guamiItemPtr = msgPdu_m.getServedGUAMIListItemDecodedPtr(i);
	   ServedGuamiItem guami;

       if (guamiItemPtr->bit_mask == ServedGUAMIItem_backupAMFName_present)
       {

       		guami.backupAmfNamePrst = true;
			guami.backupAmfName.assign((char*)(guamiItemPtr->backupAMFName.value),guamiItemPtr->backupAMFName.length);

       }

	   memcpy(guami.plmnId,
	   	      guamiItemPtr->gUAMI.pLMNIdentity.value,
	   	      sizeof(guamiItemPtr->gUAMI.pLMNIdentity.value));

	   guami.amfId.regionId = *(guamiItemPtr->gUAMI.aMFRegionID.value);

	   memcpy(guami.amfId.setId, guamiItemPtr->gUAMI.aMFSetID.value,
	   			2);

	   guami.amfId.pointer =  *(guamiItemPtr->gUAMI.aMFPointer.value);
	   
	   serGuamiList_m.push_back(guami);
     }

	//decode plmn list
   for(Uint32 i = 0; i < msgPdu_m.getNumofPlmnSupportListItem(); i++)
   {
      PLMNSupportItem * plmnSupItemPtr = 
	  	msgPdu_m.getPlmnSupportItemDecodedPtr(i);

      BPlmnItem plmnItem;
	  memcpy(plmnItem.plmnid, plmnSupItemPtr->pLMNIdentity.value,
	  	plmnSupItemPtr->pLMNIdentity.length);

      for(Uint32 j = 0; j < msgPdu_m.getNumofSlices(i); j++)
      {
         SliceSupportItem* sliceItemPtr = 
		 	msgPdu_m.getSliceSupportListDecodedPtr(plmnSupItemPtr, j);
         if (NULL == sliceItemPtr)
         {
            return false;
         }

         SNssai snssai;
         if (sliceItemPtr->s_NSSAI.bit_mask == sD_present) //#       define      sD_present 0x80
         {
            snssai.sdPresent = true;
            memcpy(snssai.sd, sliceItemPtr->s_NSSAI.sD.value, sizeof(sliceItemPtr->s_NSSAI.sD.value));
         }
		 else
         {
            snssai.sdPresent = false;
		 }
         memcpy(snssai.sst, sliceItemPtr->s_NSSAI.sST.value, sliceItemPtr->s_NSSAI.sST.length);

         plmnItem.ssList.push_back(snssai);
      }
  
      plmnSupportist_m.push_back(plmnItem);
   }

  //decode Ue Retention Information
  UeRetentionInfoPrst_m = msgPdu_m.isUeRetentionInfoPresent();
  if(true == UeRetentionInfoPrst_m)
  {
      UeRetentionInfo_m = *(msgPdu_m.getUeRetentionInfoDecodedPtr());
  }

   return true;
}

string NgSetupResponseCodec::toString()
{
	ostringstream oss;

    oss <<"NgSetupResponseCodec:" << endl
        <<"AmfName: ("<<amfName_m<<")"<<endl
        <<"RelativeAmfCapacity: "<<relativeAmfCapacity_m<<endl
        <<"ue Retention Information: flag(" <<UeRetentionInfoPrst_m<<"),value("<<(int)UeRetentionInfo_m<<")"<<endl
        <<"Served Guami List: "<< endl;
    vector<ServedGuamiItem>::iterator it;
    for (it=serGuamiList_m.begin(); it<serGuamiList_m.end() ; it++)
    {
   		oss <<"  BackupAmfName: prst  " <<it->backupAmfNamePrst <<" value "<<it->backupAmfName<<endl;
        oss <<"  PlmnId: " <<hex<<int(it->plmnId[0])<<int(it->plmnId[1])<<int(it->plmnId[2])<<endl;
		oss <<"  Amf Region ID:"<<int(it->amfId.regionId) << endl;
		oss <<"  Amf Set ID:"<<hex<<int(it->amfId.setId[0])<<int(it->amfId.setId[1])<< endl;
		oss <<"  Amf Pointer:"<<int(it->amfId.pointer) << endl;
     }

     oss <<"Plmn Support List:"<< endl;
	 vector<BPlmnItem>::iterator bIt;
     for(bIt=plmnSupportist_m.begin(); bIt<plmnSupportist_m.end(); bIt++)
     {
        oss <<"   plmnid : "<< hex<<int(bIt->plmnid[0])<<int(bIt->plmnid[1])<<int(bIt->plmnid[2])<<endl;
        SliceSupList ssList = bIt->ssList;
        vector<SNssai>::iterator it;
        for (it=ssList.begin(); it<ssList.end(); it++)
        {
           oss << "      SD Flag: " << it->sdPresent << ", SD : " << hex << int(it->sd[0])<< int(it->sd[1])<< int(it->sd[2]);
           oss << "      SST :" << int(it->sst[0]) << endl;
        }
     }

	 return oss.str();
}


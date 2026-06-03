#include <iostream>
#include <sstream>
#include <string>

#include "pagingCodec.h"

using namespace std;

PagingCodec::PagingCodec() 
{
    reset();
}

PagingCodec::~PagingCodec()
{
}

void 
PagingCodec::setUePagingIdentity(UePagingIdentity& ident)
{
    uePagingIdentity_m = ident;
}

UePagingIdentity& 
PagingCodec::getUePagingIdentity()
{
    return uePagingIdentity_m;
}

void 
PagingCodec::setPagingDrx(Uint8 pdrx)
{
    pagingDrxPrst_m = true;
    pagingDrx_m = pdrx;
}

Uint8 
PagingCodec::getPagingDrx()
{
    return pagingDrx_m;
}

bool  
PagingCodec::isPagingDrxPrst()
{
    return pagingDrxPrst_m;
}

void 
PagingCodec::addTaiPagingList(TaiType& pTai)
{
    taiPagingList_m.push_back(pTai);
}

TaiPagingList& 
PagingCodec::getTaiPagingList()
{
    return taiPagingList_m;
}

void 
PagingCodec::setPagingPriority(Uint8 pPriority)
{
    pagingPriorityPrst_m = true;
    pagingPriority_m = pPriority;
}

Uint8 
PagingCodec::getPagingPriority()
{
    return pagingPriority_m;
}

bool  
PagingCodec::isPagingPriorityPrst()
{
    return pagingPriorityPrst_m;
}

UeRadioCapForPaging& 
PagingCodec::getUeRadioCapForPaging()
{
    return ueRadioCapForPaging_m;
}

void 
PagingCodec::setUeRadioCapForPaging(UeRadioCapForPaging& val)
{
    ueRadioCapForPagingPrst_m=true;
    ueRadioCapForPaging_m = val;
}

bool 
PagingCodec::isUeRadioCapForPagingPrst()
{
    return ueRadioCapForPagingPrst_m;
}

void 
PagingCodec::setPagingOrigin(Uint8 pOrigin)
{
    pagingOriginPrst_m = true;
    pagingOrigin_m = pOrigin;
}

Uint8 
PagingCodec::getPagingOrigin()
{
    return pagingOrigin_m;
}

bool  
PagingCodec::isPagingOriginPrst()
{
    return pagingOriginPrst_m;
}

void 
PagingCodec::setAssDateForPaging(AssDateForPaging& pDate)
{
    assDataForPagingPrst_m = true;
    assDataForPaging_m = pDate;
}

AssDateForPaging& 
PagingCodec::getAssDateForPaging()
{
    return assDataForPaging_m;
}

bool  
PagingCodec::isAssDataForPagingPrst()
{
    return assDataForPagingPrst_m;
}

bool PagingCodec::reset() 
{
    msgPdu_m.reset_v();

    memset(&uePagingIdentity_m, 0, sizeof(uePagingIdentity_m));

    pagingDrx_m = 0;
    pagingDrxPrst_m = false;
    pagingPriority_m = 0;
    pagingPriorityPrst_m = false;
    pagingOrigin_m = 0;
    pagingOriginPrst_m = false;

    taiPagingList_m.clear();

    ueRadioCapForPaging_m.ueRadioCapForPagingNr.clear();
	ueRadioCapForPaging_m.ueRadioCapForPagingNrPrst = false;
    ueRadioCapForPaging_m.ueRadioCapForPagingEutra.clear();
	ueRadioCapForPaging_m.ueRadioCapForPagingEutraPrst = false;
    ueRadioCapForPagingPrst_m = false;

    memset(&assDataForPaging_m, 0, sizeof(assDataForPaging_m));
    assDataForPagingPrst_m = false;
}

//bool PagingCodec::encode(NgapOssCtxt* ctxt, MsgBuffer msgBuffer)
MsgBuffer PagingCodec::encode(NgapOssCtxt* ctxt) 
{
    //cout << "PagingCodec::encode()" << endl;
    MsgBuffer msgBuffer;
    msgBuffer.length = 0;
    msgBuffer.value = NULL;

    //encode UEPagingIdentity
    UEPagingIdentity* uePagingIdentityPtr = msgPdu_m.getUePagingIdentityEncodePtr();
    uePagingIdentityPtr->u.fiveG_S_TMSI.aMFSetID.length = amfSetIdBitLen;
    //uePagingIdentity_m.swapSetId = uePagingIdentity_m.setId[1];
    //uePagingIdentity_m.setId[1] = uePagingIdentity_m.setId[0] << 6;
    //uePagingIdentity_m.setId[0] = uePagingIdentity_m.swapSetId;

    //uePagingIdentity_m.setId[1] = uePagingIdentity_m.setId[1] << 6;
    uePagingIdentity_m.setId[0] = (uePagingIdentity_m.setId[0] << 6) | (uePagingIdentity_m.setId[1] >> 2);
    uePagingIdentity_m.setId[1] = uePagingIdentity_m.setId[1] << 6;
    
    memcpy(uePagingIdentityPtr->u.fiveG_S_TMSI.aMFSetID.value,
           uePagingIdentity_m.setId,
           sizeof(uePagingIdentity_m.setId));

    uePagingIdentityPtr->u.fiveG_S_TMSI.aMFPointer.length = amfPointerBitLen;
    uePagingIdentity_m.pointer = uePagingIdentity_m.pointer << 2;
    memcpy(uePagingIdentityPtr->u.fiveG_S_TMSI.aMFPointer.value,
           &uePagingIdentity_m.pointer,
           sizeof(uePagingIdentity_m.pointer));


    uePagingIdentityPtr->u.fiveG_S_TMSI.fiveG_TMSI.length = sizeof(uePagingIdentity_m.tmsi);
    memcpy(uePagingIdentityPtr->u.fiveG_S_TMSI.fiveG_TMSI.value, 
           uePagingIdentity_m.tmsi,
           uePagingIdentityPtr->u.fiveG_S_TMSI.fiveG_TMSI.length);

    //encode PagingDrx
    if(pagingDrxPrst_m)
    {
        PagingDRX* pagingDrxPtr = msgPdu_m.getPagingDrxEncodePtr();
        *pagingDrxPtr = (PagingDRX)(pagingDrx_m);
    }

    //encode TAIListForPaging
    unsigned int numofPagingTaiItems = taiPagingList_m.size();
    msgPdu_m.setNumofTAIListForPagingItem_v(numofPagingTaiItems);
    for (int i = 0; i < numofPagingTaiItems; i++) 
    {
        TAIListForPagingItem* taiPagingItemPtr = 
            msgPdu_m.getTAIListForPagingItemEncodePtr(i); 
        if (NULL == taiPagingItemPtr)
        {
            std::cout << "NULL Pointer for taiPagingItemPtr" << std::endl;
            return msgBuffer;
        }
        taiPagingItemPtr->tAI.pLMNIdentity.length = sizeof(taiPagingList_m[i].plmn);

        memcpy(taiPagingItemPtr->tAI.pLMNIdentity.value,
               taiPagingList_m[i].plmn,
               taiPagingItemPtr->tAI.pLMNIdentity.length);

        taiPagingItemPtr->tAI.tAC.length = sizeof(taiPagingList_m[i].tac);

        memcpy(taiPagingItemPtr->tAI.tAC.value,
               taiPagingList_m[i].tac,
               taiPagingItemPtr->tAI.tAC.length);
	}

    //encode PagingPriority
    if(pagingPriorityPrst_m)
    {
        PagingPriority* pagingPriorityPtr = msgPdu_m.getPagePriorityEncodePtr();
        *pagingPriorityPtr = (PagingPriority)(pagingPriority_m);
    }

    //encode UERadioCapabilityForPaging
    if(ueRadioCapForPagingPrst_m) {
        UERadioCapabilityForPaging* ueRadioCapForPagingPtr = 
            msgPdu_m.getUeRadioCapForPagingEncodePtr();
		if(ueRadioCapForPaging_m.ueRadioCapForPagingNrPrst) {
			ueRadioCapForPagingPtr->bit_mask |= uERadioCapabilityForPagingOfNR_present;
            ueRadioCapForPagingPtr->uERadioCapabilityForPagingOfNR
			    .length = ueRadioCapForPaging_m.ueRadioCapForPagingNr.length();
            memcpy(ueRadioCapForPagingPtr->uERadioCapabilityForPagingOfNR.value, 
                   ueRadioCapForPaging_m.ueRadioCapForPagingNr.c_str(), 
                   ueRadioCapForPagingPtr->uERadioCapabilityForPagingOfNR.length);
		}
		if(ueRadioCapForPaging_m.ueRadioCapForPagingEutraPrst) {
			ueRadioCapForPagingPtr->bit_mask |= uERadioCapabilityForPagingOfEUTRA_present;
            ueRadioCapForPagingPtr->uERadioCapabilityForPagingOfEUTRA
			    .length = ueRadioCapForPaging_m.ueRadioCapForPagingEutra.length();
            memcpy(ueRadioCapForPagingPtr->uERadioCapabilityForPagingOfEUTRA.value, 
                   ueRadioCapForPaging_m.ueRadioCapForPagingEutra.c_str(), 
                   ueRadioCapForPagingPtr->uERadioCapabilityForPagingOfEUTRA.length);
		}

    }

    //encode PagingOrigin
    if(pagingOriginPrst_m)
    {
        PagingOrigin* pagingOriginPtr = msgPdu_m.getPagingOriginEncodePtr();
        *pagingOriginPtr = (PagingOrigin)(pagingOrigin_m);
    }

    //encode AssistanceDataForPaging
    if(assDataForPagingPrst_m)
    {
        AssistanceDataForPaging* assDataForPagingPtr = msgPdu_m.getAssDataForPagingEncodePtr();

        if(assDataForPaging_m.recommandCellListPrst)
        {
            assDataForPagingPtr->bit_mask = assistanceDataForRecommendedCells_present;

            Uint32 numOfRecommendedCellItem_ma = assDataForPaging_m.recommandCellList.size();
            msgPdu_m.setNumOfRecommendedCellItem_v(numOfRecommendedCellItem_ma);
            for(int i =0;i<numOfRecommendedCellItem_ma;i++)
            {
                RecommendedCellItem* cellPtr = msgPdu_m.getRecommendedCellItemEncodePtr(i);
                if (NULL == cellPtr)
                {
                    std::cout << "NULL Pointer for cellPtr" << std::endl;
                    return msgBuffer;
                }

                if(assDataForPaging_m.recommandCellList[i].timeStayedPrst)
                {
                    cellPtr->bit_mask = ExpectedUEMovingTrajectoryItem_timeStayedInCell_present;
                    cellPtr->timeStayedInCell = assDataForPaging_m.recommandCellList[i].timeStayedInCell;
                }

                cellPtr->nGRAN_CGI.u.nR_CGI.pLMNIdentity.length = 
                    sizeof(assDataForPaging_m.recommandCellList[i].nrcgi.plmn);
                memcpy(cellPtr->nGRAN_CGI.u.nR_CGI.pLMNIdentity.value,
                       assDataForPaging_m.recommandCellList[i].nrcgi.plmn,
                       cellPtr->nGRAN_CGI.u.nR_CGI.pLMNIdentity.length);

                cellPtr->nGRAN_CGI.u.nR_CGI.nRCellIdentity.length = NRCellIdentityBitLen;
                memcpy(cellPtr->nGRAN_CGI.u.nR_CGI.nRCellIdentity.value,
                       assDataForPaging_m.recommandCellList[i].nrcgi.nrCellId,
                       sizeof(assDataForPaging_m.recommandCellList[i].nrcgi.nrCellId));
            }
        }

        if(assDataForPaging_m.pagingAttemptInfoPrst)
        {
            assDataForPagingPtr->bit_mask |= pagingAttemptInformation_present;

            assDataForPagingPtr->pagingAttemptInformation.pagingAttemptCount = 
                assDataForPaging_m.pagingAttemptInfo.count;

            assDataForPagingPtr->pagingAttemptInformation.intendedNumberOfPagingAttempts = 
                assDataForPaging_m.pagingAttemptInfo.intendedNum;

            if(assDataForPaging_m.pagingAttemptInfo.nextAreaScopePrst)
            {   
                assDataForPagingPtr->pagingAttemptInformation.bit_mask |= 
                    nextPagingAreaScope_present;

                assDataForPagingPtr->pagingAttemptInformation.nextPagingAreaScope = 
                    (NextPagingAreaScope)(assDataForPaging_m.pagingAttemptInfo.nextAreaScope);
            }
        }
    }
    msgPdu_m.encodeOssToPer(ctxt);

    OssBuf* perBuf = msgPdu_m.getPerBufRef();

    msgBuffer.length = perBuf->length;
    msgBuffer.value = perBuf->value;

    return msgBuffer;
}

bool PagingCodec::decode(NgapOssCtxt* ctxt,MsgBuffer msgBuf) 
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

    cout << "Start to retriev the decoded info from OSS" << endl;

    //decode UEPagingIdentity
    memcpy(uePagingIdentity_m.setId, 
           msgPdu_m.getUePagingIdentityDecodedPtr()->u.fiveG_S_TMSI.aMFSetID.value,
           amfSetIdLen);
    Uint8 swapSetId = 0;
    swapSetId = uePagingIdentity_m.setId[0];
    uePagingIdentity_m.setId[0] = uePagingIdentity_m.setId[1] >> 6;
    uePagingIdentity_m.setId[1] = swapSetId;

    uePagingIdentity_m.pointer = 
        *(msgPdu_m.getUePagingIdentityDecodedPtr()->u.fiveG_S_TMSI.aMFPointer.value);
    uePagingIdentity_m.pointer = uePagingIdentity_m.pointer >> 2;

    memcpy(uePagingIdentity_m.tmsi,
		   msgPdu_m.getUePagingIdentityDecodedPtr()->u.fiveG_S_TMSI.fiveG_TMSI.value,
		   msgPdu_m.getUePagingIdentityDecodedPtr()->u.fiveG_S_TMSI.fiveG_TMSI.length);

    //decode PagingDRX
    if(msgPdu_m.isPagingDrxPresent())
    {
        pagingDrx_m = *(msgPdu_m.getPagingDrxDecodedPtr());
        pagingDrxPrst_m = true;
    }

    //decode TAIListForPaging
	for (int i = 0; i < msgPdu_m.getNumofTAIListForPagingItem();i++) 
    {
		TAIListForPagingItem *taiPtr =
			msgPdu_m.getTAIListForPagingItemDecodedPtr(i);
		if (NULL == taiPtr) {
			return false;
		}
		  
		TaiType taiItem;
    
        memcpy(taiItem.tac,
               taiPtr->tAI.tAC.value,
               taiPtr->tAI.tAC.length);

        memcpy(taiItem.plmn,
               taiPtr->tAI.pLMNIdentity.value,
               taiPtr->tAI.pLMNIdentity.length);

        taiPagingList_m.push_back(taiItem);
	}

    //decode PagingPriority
    if(msgPdu_m.isPagePriorityPresent())
    {
        pagingPriority_m = *(msgPdu_m.getPagePriorityDecodedPtr());
        pagingPriorityPrst_m = true;
    }

    //decode UERadioCapabilityForPaging
    if(msgPdu_m.isUeRadioCapabilityForPagingPresent())
    {
        if(msgPdu_m.getUeRadioCapForPagingDecodedPtr()->bit_mask & uERadioCapabilityForPagingOfNR_present)
       	{
            UERadioCapabilityForPaging* ueRadioCapForPagingPtr = 
                msgPdu_m.getUeRadioCapForPagingDecodedPtr();
			ueRadioCapForPaging_m.ueRadioCapForPagingNrPrst = true;
            ueRadioCapForPaging_m.ueRadioCapForPagingNr
				.assign((char*)(ueRadioCapForPagingPtr->uERadioCapabilityForPagingOfNR.value), 
                                ueRadioCapForPagingPtr->uERadioCapabilityForPagingOfNR.length);
        }
        if(msgPdu_m.getUeRadioCapForPagingDecodedPtr()->bit_mask & uERadioCapabilityForPagingOfEUTRA_present)
       	{
            UERadioCapabilityForPaging* ueRadioCapForPagingPtr = 
                msgPdu_m.getUeRadioCapForPagingDecodedPtr();
			ueRadioCapForPaging_m.ueRadioCapForPagingEutraPrst = true;
            ueRadioCapForPaging_m.ueRadioCapForPagingEutra
				.assign((char*)(ueRadioCapForPagingPtr->uERadioCapabilityForPagingOfEUTRA.value), 
                                ueRadioCapForPagingPtr->uERadioCapabilityForPagingOfEUTRA.length);
        }

        ueRadioCapForPagingPrst_m = true;
    }

    //decode PagingOrigin
    if(msgPdu_m.isPagingOriginPresent())
    {
        pagingOrigin_m = *(msgPdu_m.getPagingOriginDecodedPtr());
        pagingOriginPrst_m = true;
    }

    //decode AssistanceDataForPaging
    if(msgPdu_m.isAssDataForPagingPresent())
    {
        if(msgPdu_m.getAssDataForPagingDecodedPtr()->bit_mask & pagingAttemptInformation_present)
        {
            assDataForPaging_m.pagingAttemptInfoPrst = true;

            assDataForPaging_m.pagingAttemptInfo.count = msgPdu_m
                .getAssDataForPagingDecodedPtr()->pagingAttemptInformation.pagingAttemptCount;

            assDataForPaging_m.pagingAttemptInfo.intendedNum = msgPdu_m
                .getAssDataForPagingDecodedPtr()->pagingAttemptInformation.intendedNumberOfPagingAttempts;

            if(msgPdu_m.getAssDataForPagingDecodedPtr()->pagingAttemptInformation.bit_mask == 
                nextPagingAreaScope_present)
            {
                assDataForPaging_m.pagingAttemptInfo.nextAreaScopePrst = true;

                assDataForPaging_m.pagingAttemptInfo.nextAreaScope = msgPdu_m
                    .getAssDataForPagingDecodedPtr()->pagingAttemptInformation.nextPagingAreaScope;
            }
        }
        if(msgPdu_m.getAssDataForPagingDecodedPtr()->bit_mask & 
            assistanceDataForRecommendedCells_present)
        {
            assDataForPaging_m.recommandCellListPrst = true;

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

                if(cellPtr->bit_mask == RecommendedCellItem_timeStayedInCell_present)
                {
                    cellItem_m.timeStayedPrst = true;
                    cellItem_m.timeStayedInCell = cellPtr->timeStayedInCell;
                }
				assDataForPaging_m.recommandCellList.push_back(cellItem_m);
            }
        }
        assDataForPagingPrst_m = true;
    }

    return true;
}


string PagingCodec::toString(){
   ostringstream oss;

   oss << "PagingCodec~~~~" << endl
       << "UEPagingIdentity" << endl
       << "AmfPointer:(" << int(uePagingIdentity_m.pointer) << "),"<< endl
       << "AMFSetID:(" << int(uePagingIdentity_m.setId[0]) << int(uePagingIdentity_m.setId[1])
       << "),Tmsi:(" << uePagingIdentity_m.tmsi[0] << uePagingIdentity_m.tmsi[1]
       << uePagingIdentity_m.tmsi[2] << uePagingIdentity_m.tmsi[3] << ")" << endl
       << "Paging Drx:Prst " << pagingDrxPrst_m <<" value: " <<(int)pagingDrx_m <<endl
	   <<" Tai List: "<<endl;
	   vector<TaiType>::iterator it;
	   for (it = taiPagingList_m.begin(); it < taiPagingList_m.end() ; it++)
	   {
			oss<< "		   plmn: " << it->plmn[0]<< it->plmn[1]<< it->plmn[2]<<endl
			   << "		   tac :" << it->tac[0] << it->tac[1] << it->tac[2] << endl;
		}
   oss << "Paging Priority:" <<(int)pagingPriority_m <<endl
	   << "UeRadioCapabilityForPaging:Prst" <<ueRadioCapForPagingPrst_m<< endl
	   << "Nr:Prst" << ueRadioCapForPaging_m.ueRadioCapForPagingNrPrst << endl
	   << " value " << ueRadioCapForPaging_m.ueRadioCapForPagingNr <<endl
	   << "Eutra:Prst" << ueRadioCapForPaging_m.ueRadioCapForPagingEutraPrst << endl
	   << " value " << ueRadioCapForPaging_m.ueRadioCapForPagingEutra <<endl
       << "Assistance Data for Paging:Prst " << assDataForPagingPrst_m <<endl
       << "Assistance Data for Recommend Cell:Prst " <<assDataForPaging_m.recommandCellListPrst <<endl;
	   vector<RecommandCellItem>::iterator rit;
	   for (rit = assDataForPaging_m.recommandCellList.begin(); rit < assDataForPaging_m.recommandCellList.end() ; rit++)
	   {
			oss<<"timeStayedPrst: " << rit->timeStayedPrst 
				<<" value:" << rit->timeStayedInCell << endl
				<<"nrcgi~~~~"<<endl
				<<"plmnid: "<<rit->nrcgi.plmn[0]
				<<rit->nrcgi.plmn[1]<<rit->nrcgi.plmn[2]<<endl
				<<"nrcellId: "<<rit->nrcgi.nrCellId[0]<<rit->nrcgi.nrCellId[1]
				<<rit->nrcgi.nrCellId[2]<<rit->nrcgi.nrCellId[3]
				<<rit->nrcgi.nrCellId[4]<<endl;
	    }

   oss << " Paging Attempt Information:Prst " <<assDataForPaging_m.pagingAttemptInfoPrst <<endl
       << "count:" <<(int)assDataForPaging_m.pagingAttemptInfo.count <<endl
       << "num:" <<(int)assDataForPaging_m.pagingAttemptInfo.intendedNum <<endl
       << "scope:Prst " <<assDataForPaging_m.pagingAttemptInfo.nextAreaScopePrst
       << " value:" <<(int)assDataForPaging_m.pagingAttemptInfo.nextAreaScope <<endl
       << "Paging Origin:Prst " << pagingOriginPrst_m <<" value: " <<(int)pagingOrigin_m <<endl
       << endl;
    return oss.str(); 
}

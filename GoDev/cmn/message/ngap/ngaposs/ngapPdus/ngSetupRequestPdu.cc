#include <string.h>
#include <iostream>

#include "ngSetupRequestPdu.h"

NGSetupRequestPdu::NGSetupRequestPdu()
{
    reset_v();
}

NGSetupRequestPdu::~NGSetupRequestPdu()
{
}

void
NGSetupRequestPdu::reset_v()
{
    NGSetupRequest_mp = NULL;

    //Global ENB ID
    globalRanNodeId_mp = NULL;
    globalRanNodeId_m.choice = globalGNB_ID_chosen;
    globalRanNodeId_m.u.globalGNB_ID.bit_mask = 0;
	globalRanNodeId_m.u.globalGNB_ID.iE_Extensions = NULL;
	globalRanNodeId_m.u.globalGNB_ID.pLMNIdentity.length = 0;
	memset(globalRanNodeId_m.u.globalGNB_ID.pLMNIdentity.value, 0xff,
		   sizeof(globalRanNodeId_m.u.globalGNB_ID.pLMNIdentity.value));
	globalRanNodeId_m.u.globalGNB_ID.gNB_ID.choice = gNB_ID_chosen;
	globalRanNodeId_m.u.globalGNB_ID.gNB_ID.u.gNB_ID.length = 0;
    memset(gnbIdValue_ma, 0xff, sizeof(gnbIdValue_ma));
    globalRanNodeId_m.u.globalGNB_ID.gNB_ID.u.gNB_ID.value = gnbIdValue_ma;
	
    //GNB Name
    ranNodeName_mp = NULL;
    ranNodeName_m.length = 0;
    memset(ranNodeNameValue_ma, 0xff, sizeof(ranNodeNameValue_ma));
    ranNodeName_m.value = ranNodeNameValue_ma;
    ranNodeNamePresent_m = false;

    //Paging DRX
    pagingDrx_m = v32;
    pagingDrx_mp = NULL;

    //Supported TAs
    supportedTAList_m = NULL;
    supportedTAList_mp = NULL;
    numOfSupportedTAListItem_m = 0;
    for (unsigned int i = 0; i < maxnoofTACs; i++)
    {
        supportedTAList_ma[i].next = NULL;
        supportedTAList_ma[i].value.bit_mask = 0;
        supportedTAList_ma[i].value.iE_Extensions = NULL;
        supportedTAList_ma[i].value.tAC.length = 0;
        memset(supportedTAList_ma[i].value.tAC.value, 0xff,
               sizeof(supportedTAList_ma[i].value.tAC.value));

			   
		numofBroadcastPLMNItem_m[i] = 0;
        supportedTAList_ma[i].value.broadcastPLMNList = NULL;
        for (unsigned int j = 0; j < maxnoofBPLMNs; j++)
        {
            broadcastPLMNItem_ma[i][j].next = NULL;
            broadcastPLMNItem_ma[i][j].value.bit_mask = 0;
			broadcastPLMNItem_ma[i][j].value.iE_Extensions = NULL;
			broadcastPLMNItem_ma[i][j].value.pLMNIdentity.length = 0;
			memset(broadcastPLMNItem_ma[i][j].value.pLMNIdentity.value,0xff,
				   sizeof(broadcastPLMNItem_ma[i][j].value.pLMNIdentity.value));

		    numOfSliceSupportItem_m[i][j] = 0;
			for(unsigned int k = 0; k < maxnoofSlices_c; k++) //maxnoofSliceItems
			{
			    sliceSupportItem_ma[i][j][k].next =  NULL;
				sliceSupportItem_ma[i][j][k].value.bit_mask = 0;
				sliceSupportItem_ma[i][j][k].value.iE_Extensions = NULL;
				sliceSupportItem_ma[i][j][k].value.s_NSSAI.bit_mask = 0;
				sliceSupportItem_ma[i][j][k].value.s_NSSAI.iE_Extensions = NULL;
				sliceSupportItem_ma[i][j][k].value.s_NSSAI.sD.length = 0;
				memset(sliceSupportItem_ma[i][j][k].value.s_NSSAI.sD.value,0,
					   sizeof(sliceSupportItem_ma[i][j][k].value.s_NSSAI.sD.value));
				sliceSupportItem_ma[i][j][k].value.s_NSSAI.sST.length = 0;
				memset(sliceSupportItem_ma[i][j][k].value.s_NSSAI.sST.value,0xff,
					   sizeof(sliceSupportItem_ma[i][j][k].value.s_NSSAI.sST.value));
			}
        }
    }

	//Ue Retention Information
	UeRetentionInfo_mp = NULL;
	UeRetentionInfoPresent_m = false;
	UeRetentionInfo_m = ues_retained;
}

//----------------------------------------------------------------------------
// This is a public API. Refer to NGSetupRequestPdu.h
//----------------------------------------------------------------------------
unsigned int
NGSetupRequestPdu::encodeOssToPer(NgapOssCtxt* ctxt)
{
    if(NULL == ctxt)
	{
	    return -1;
    }
	
    ProtocolIeType *nextIe_p = NULL;

    ProtocolIeType pagingDrxIe;
    pagingDrxIe.next = nextIe_p;
    pagingDrxIe.value.id = id_DefaultPagingDRX;
    pagingDrxIe.value.criticality = ignore;
    pagingDrxIe.value.value.pduNum =PDU_NGSetupRequestIEs_Value_PagingDRX;
    pagingDrxIe.value.value.decoded.pdu_PagingDRX = &pagingDrx_m;
    pagingDrxIe.value.value.encoded.length = 0;
    pagingDrxIe.value.value.encoded.value = NULL;
    nextIe_p = &pagingDrxIe;

    ProtocolIeType supportedTAsIe;
    if (numOfSupportedTAListItem_m > maxnoofTACs ||
        numOfSupportedTAListItem_m <= 0)
    {
        return defaultFailureCode_c;
    }
    else
    {
        struct SupportedTAList_ *nextSupportedTas_p = NULL;
        for (unsigned int i = 0; i < numOfSupportedTAListItem_m; i++)
        {
            struct BroadcastPLMNList_ *bplmns_p = NULL;
            for (unsigned int j = 0; j < numofBroadcastPLMNItem_m[i]; j++)
            {

				struct SliceSupportList_ *slices_p = NULL;
				for(unsigned int k=0; k < numOfSliceSupportItem_m[i][j]; k++)
				{
					sliceSupportItem_ma[i][j][k].next = slices_p;
					slices_p = &sliceSupportItem_ma[i][j][k];
				}
				broadcastPLMNItem_ma[i][j].value.tAISliceSupportList = slices_p;
                broadcastPLMNItem_ma[i][j].next = bplmns_p;
                bplmns_p = &(broadcastPLMNItem_ma[i][j]);
            }

            supportedTAList_ma[i].value.broadcastPLMNList = bplmns_p;

            supportedTAList_ma[i].next = nextSupportedTas_p;
			nextSupportedTas_p = &supportedTAList_ma[i];
        }
        supportedTAList_m = nextSupportedTas_p;

        supportedTAsIe.value.value.pduNum = PDU_NGSetupRequestIEs_Value_SupportedTAList;
        supportedTAsIe.value.value.decoded.pdu_SupportedTAList = &supportedTAList_m;
        supportedTAsIe.value.value.encoded.length = 0;
        supportedTAsIe.value.value.encoded.value = NULL;
    }
    supportedTAsIe.next = nextIe_p;
    supportedTAsIe.value.id = id_SupportedTAList;
    supportedTAsIe.value.criticality = reject;
    nextIe_p = &supportedTAsIe;

    ProtocolIeType ranNameIe;
    if (ranNodeNamePresent_m == true)
    {
        ranNameIe.next = nextIe_p;
        ranNameIe.value.id = id_RANNodeName;
        ranNameIe.value.criticality = ignore;
        ranNameIe.value.value.pduNum = PDU_NGSetupRequestIEs_Value_RANNodeName;
        ranNameIe.value.value.decoded.pdu_RANNodeName = &ranNodeName_m;
        ranNameIe.value.value.encoded.length = 0;
        ranNameIe.value.value.encoded.value = NULL;
		
        nextIe_p = &ranNameIe;
    }

    ProtocolIeType globalgnbIdIe;
    globalgnbIdIe.next = nextIe_p;
    globalgnbIdIe.value.id = id_GlobalRANNodeID;
    globalgnbIdIe.value.criticality = reject;
    globalgnbIdIe.value.value.pduNum = PDU_NGSetupRequestIEs_Value_GlobalRANNodeID;
    globalgnbIdIe.value.value.decoded.pdu_GlobalRANNodeID = &globalRanNodeId_m;
    globalgnbIdIe.value.value.encoded.length = 0;
    globalgnbIdIe.value.value.encoded.value = NULL;
	nextIe_p = &globalgnbIdIe;

	ProtocolIeType ueRetentionInfoIe;
	if (UeRetentionInfoPresent_m == true)
    {
	    ueRetentionInfoIe.next = nextIe_p;
        ueRetentionInfoIe.value.id = id_UERetentionInformation;
        ueRetentionInfoIe.value.criticality = ignore;
        ueRetentionInfoIe.value.value.pduNum = PDU_NGSetupRequestIEs_Value_UERetentionInformation;
        ueRetentionInfoIe.value.value.decoded.pdu_UERetentionInformation = &UeRetentionInfo_m;
        ueRetentionInfoIe.value.value.encoded.length = 0;
        ueRetentionInfoIe.value.value.encoded.value = NULL;
	    nextIe_p = &ueRetentionInfoIe;
	}

    NGSetupRequest msg;
    msg.protocolIEs = nextIe_p;
	
    NGAP_PDU ngApPdu;

    ngApPdu.choice = initiatingMessage_chosen;
    ngApPdu.u.initiatingMessage.procedureCode = id_NGSetup;
    ngApPdu.u.initiatingMessage.criticality = reject;
    ngApPdu.u.initiatingMessage.value.pduNum = PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_NGSetupRequest;
    ngApPdu.u.initiatingMessage.value.decoded.pdu_NGSetupRequest = &msg;
    ngApPdu.u.initiatingMessage.value.encoded.length = 0;
    ngApPdu.u.initiatingMessage.value.encoded.value = NULL;

    //Convert Structures to PER
    OssBuf outputBuf;
    unsigned int result;

	ctxt->reset_v();
    outputBuf = ctxt->getPrimaryCodecBuf();

    result = ctxt->encodeOssToPer(NGAP_PDU_PDU, &ngApPdu, &outputBuf);

    if (result == encodeSuccessCode_c)
    {
        perBufRef_m = outputBuf;
    }

    return (result);
}

void
NGSetupRequestPdu::getDecodedOssDataPtrs_v()
{
    ProtocolIeType *currIe_p;
    reset_v();

    ngapPdu_mp = reinterpret_cast<NGAP_PDU *> (ossBufRef_m.value);
    if (ngapPdu_mp == NULL)
    {
        return;
    }

    NGSetupRequest_mp = 
             ngapPdu_mp->u.initiatingMessage.value.decoded.pdu_NGSetupRequest;

    if (NGSetupRequest_mp == NULL)
    {
        return;
    }

    currIe_p = NGSetupRequest_mp->protocolIEs;

    while (currIe_p != NULL)
    {
        if (currIe_p->value.id == id_GlobalRANNodeID
            && currIe_p->value.value.decoded.pdu_GlobalRANNodeID != NULL)
        {
            globalRanNodeId_mp = 
                      currIe_p->value.value.decoded.pdu_GlobalRANNodeID;
        }
        else if (currIe_p->value.id == id_RANNodeName
                 && currIe_p->value.value.decoded.pdu_RANNodeName != NULL)
        {
            ranNodeName_mp = currIe_p->value.value.decoded.pdu_RANNodeName;
            ranNodeNamePresent_m = true;
        }
        else if (currIe_p->value.id == id_SupportedTAList
                 && currIe_p->value.value.decoded.pdu_SupportedTAList != NULL)
        {
            supportedTAList_mp = 
                     currIe_p->value.value.decoded.pdu_SupportedTAList;

            SupportedTAList curr_p = NULL;
            BroadcastPLMNList bplmns_p = NULL;
			SliceSupportList  slices_p = NULL;				
            //curr_p = static_cast<SupportedTAList>(*supportedTAList_mp);
			curr_p = *supportedTAList_mp;
            while (curr_p != NULL)
            {
                bplmns_p = curr_p->value.broadcastPLMNList;
                while (bplmns_p != NULL)
                {
					slices_p = bplmns_p->value.tAISliceSupportList;
					while(slices_p != NULL)
					{
						numOfSliceSupportItem_m[numOfSupportedTAListItem_m][numofBroadcastPLMNItem_m[numOfSupportedTAListItem_m]]++;
						slices_p = slices_p->next;
					}
					
                    numofBroadcastPLMNItem_m[numOfSupportedTAListItem_m]++;
                    bplmns_p = bplmns_p->next;
                }
                numOfSupportedTAListItem_m++;
                curr_p = curr_p->next;
            }
        }
        else if (currIe_p->value.id == id_DefaultPagingDRX
                 && currIe_p->value.value.decoded.pdu_PagingDRX != NULL)
        {
            pagingDrx_mp = currIe_p->value.value.decoded.pdu_PagingDRX;
        }
		else if (currIe_p->value.id == id_UERetentionInformation
			     && currIe_p->value.value.decoded.pdu_UERetentionInformation != NULL)
		{
		    UeRetentionInfo_mp = currIe_p->value.value.decoded.pdu_UERetentionInformation;
			UeRetentionInfoPresent_m = true;
		}

        currIe_p = currIe_p->next;
    }
}

NGSetupRequest *
NGSetupRequestPdu::getNgSetupRequestDecodedPtr() const
{
    return NGSetupRequest_mp;
}

GlobalRANNodeID * 
NGSetupRequestPdu::getGlobalRanNodeIdDecodedPtr() const
{
    return globalRanNodeId_mp;
}

GlobalRANNodeID * 
NGSetupRequestPdu::getGlobalRanNodeIdEncodePtr()
{
	return &globalRanNodeId_m;
}

RANNodeName * 
NGSetupRequestPdu::getRanNodeNameDecodedPtr() const
{
    return ranNodeName_mp;
}

RANNodeName * 
NGSetupRequestPdu::getRanNodeNameEncodePtr()
{
    ranNodeNamePresent_m = true;
    return &ranNodeName_m;
}

bool 
NGSetupRequestPdu::isRanNodeNamePresent() const
{
    return ranNodeNamePresent_m;
}

PagingDRX * 
NGSetupRequestPdu::getPagingDrxDecodedPtr() const
{
    return pagingDrx_mp;
}

PagingDRX * 
NGSetupRequestPdu::getPagingDrxEncodePtr()
{
    return &pagingDrx_m;
}

SupportedTAList * 
NGSetupRequestPdu::getSupportedTAListDecodedPtr() const
{
    return supportedTAList_mp;
}

SupportedTAItem * 
NGSetupRequestPdu::getSupportedTAListItemDecodedPtr(unsigned int index) const
{
    SupportedTAItem *item_p = NULL;
	
    SupportedTAList supportedTAs_p = NULL;

    if (supportedTAList_mp != NULL)
    {
        supportedTAs_p = *(supportedTAList_mp);
    }

    if (supportedTAs_p != NULL
        && (index >= 0 && index < numOfSupportedTAListItem_m))
    {
        unsigned int i = 0;
        for (; i < index && supportedTAs_p != NULL; i++)
        {
            supportedTAs_p = supportedTAs_p->next;
        }
        if (supportedTAs_p != NULL && i == index)
        {
            item_p = &(supportedTAs_p->value);
        }
    }

    return item_p;
}

SupportedTAList * 
NGSetupRequestPdu::getSupportedTAListEncodePtr()
{
    return &supportedTAList_m;
}

SupportedTAItem * 
NGSetupRequestPdu::getSupportedTAListItemEncodePtr(unsigned int index)
{
    if (index >= 0 && index < maxnoofTACs)
    {
        return &(supportedTAList_ma[index].value);
    }
    else
    {
        return NULL;
    }

}

void 
NGSetupRequestPdu::setNumOfSupportedTAListItem_v(unsigned int num)
{
	numOfSupportedTAListItem_m = num;
}
unsigned int 
NGSetupRequestPdu::getNumOfSupportedTAListItem() const
{
   return numOfSupportedTAListItem_m;
}

BroadcastPLMNItem*  
NGSetupRequestPdu::getBPlmnsDecodedPtr(SupportedTAItem * taiPtr,  
											unsigned int indBPLMNs) const
{
    BroadcastPLMNItem* item_p = NULL;
    BroadcastPLMNList bplmns_p = NULL;
    if (taiPtr != NULL)
    {
        bplmns_p = taiPtr->broadcastPLMNList;
        if (bplmns_p != NULL
            && (indBPLMNs >= 0 && indBPLMNs < maxnoofBPLMNs))
        {
            unsigned int i = 0;
            for (; i < indBPLMNs && bplmns_p != NULL; i++)
            {
                bplmns_p = bplmns_p->next;
            }

		    if(bplmns_p != NULL  && i == indBPLMNs)
		    {
				item_p = &(bplmns_p->value);
		    }
        }
    }

    return item_p;

}

BroadcastPLMNItem* 
NGSetupRequestPdu::getBPlmnsEncodePtr(unsigned int indTACs, 
										           unsigned int indBPLMNs)
{
    if ((indTACs >= 0 && indTACs < maxnoofTACs) &&
        (indBPLMNs >= 0 && indBPLMNs < maxnoofBPLMNs))
    {
        return &(broadcastPLMNItem_ma[indTACs][indBPLMNs].value);
    }
    else
    {
        return NULL;
    }
}
												   
void 
NGSetupRequestPdu::setNumOfBPlmns_v(unsigned int num, unsigned int index)
{
    if (index >= 0 && index < maxnoofTACs)
    {
        numofBroadcastPLMNItem_m[index] = num;
    }
}

unsigned int 
NGSetupRequestPdu::getNumOfBPlmns(unsigned int index) const
{
    if (index >= 0 && index < maxnoofTACs)
    {
        return numofBroadcastPLMNItem_m[index];
    }
    else
    {
        return 0;
    }
}

SliceSupportItem* 
NGSetupRequestPdu::getSliceSupportListDecodedPtr(BroadcastPLMNItem * bplmns_p,		   
														 unsigned int indSlices) const
{
    SliceSupportItem* item_p = NULL;
	SliceSupportList slices_p = NULL;
	if (bplmns_p != NULL)
	{
        slices_p = bplmns_p->tAISliceSupportList;
		if (slices_p != NULL && (indSlices >=0 && indSlices < maxnoofSlices_c))  //maxnoofSliceItems
		{
            unsigned int i = 0;
		    for(; i< indSlices && slices_p != NULL; i++)
		    {
		    	slices_p = slices_p->next;
		    }
			
			if(slices_p != NULL  && i == indSlices)
		    {
				item_p = &(slices_p->value);
		    }
		}
	}
    return item_p;
}

SliceSupportItem* 
NGSetupRequestPdu::getSliceSupportListEncodePtr(unsigned int indTACs, 
											  unsigned int indBPLMNs, 
											  unsigned int indSlices)
{

    if ((indTACs >= 0 && indTACs < maxnoofTACs) &&
        (indBPLMNs >= 0 && indBPLMNs < maxnoofBPLMNs) &&
        (indSlices >= 0 && indSlices < maxnoofSlices_c)) //maxnoofSliceItems
    {
        return &(sliceSupportItem_ma[indTACs][indBPLMNs][indSlices].value);
    }
    else
    {
        return NULL;
    }

}

void 
NGSetupRequestPdu::setNumOfSlices_v(unsigned int num, unsigned int indexTac, unsigned int indexPlmns)
{
    if ((indexTac >= 0 && indexTac < maxnoofTACs) &&
		(indexPlmns >= 0 && indexPlmns < maxnoofBPLMNs))
    {
        numOfSliceSupportItem_m[indexTac][indexPlmns] = num;
    }
}

unsigned int 
NGSetupRequestPdu::getNumOfSlices(unsigned int indexTac, unsigned int indexPlmns) const
{
    if ((indexTac >= 0 && indexTac < maxnoofTACs) &&
		(indexPlmns >= 0 && indexPlmns < maxnoofBPLMNs))

    {
        return numOfSliceSupportItem_m[indexTac][indexPlmns];
    }
    else
    {
        return 0;
    }
}

UERetentionInformation *
NGSetupRequestPdu::getUeRetentionInfoDecodedPtr() const
{
    return UeRetentionInfo_mp;
}

UERetentionInformation *
NGSetupRequestPdu::getUeRetentionInfoEncodePtr()
{
    UeRetentionInfoPresent_m = true;
    return &UeRetentionInfo_m;
}

bool
NGSetupRequestPdu::isUeRetentionInfoPresent() const
{
    return UeRetentionInfoPresent_m;
}


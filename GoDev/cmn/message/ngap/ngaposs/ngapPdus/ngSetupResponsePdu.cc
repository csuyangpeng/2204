#include <string.h>
#include <iostream>

#include "ngSetupResponsePdu.h"

NGSetupResponsePdu::NGSetupResponsePdu()
{
    reset_v();
}

NGSetupResponsePdu::~NGSetupResponsePdu()
{
}

void
NGSetupResponsePdu::reset_v()
{
    NGSetupResponse_mp = NULL;

    //AMF Name 
    amfName_mp = NULL;
    amfName_m.length = 0;
    memset(amfNameValue_ma, 0xff, sizeof(amfNameValue_ma));
    amfName_m.value = amfNameValue_ma;

	//Served GUAMI List
	servedGUAMIList_m = NULL;
	servedGUAMIList_mp = NULL;
	numOfSrvGUAMIListItem_m = 0;
	for (unsigned int i = 0; i < maxnoofServedGUAMIs_c; i++)
	{
		 servedGUAMIList_ma[i].next = NULL;
		 servedGUAMIList_ma[i].value.bit_mask = 0;
		 servedGUAMIList_ma[i].value.iE_Extensions = NULL;
		 servedGUAMIList_ma[i].value.backupAMFName.length = 0;
		 memset(backupAmfNameValue_ma, 0xff, sizeof(backupAmfNameValue_ma));
		 servedGUAMIList_ma[i].value.backupAMFName.value = backupAmfNameValue_ma;
		 servedGUAMIList_ma[i].value.gUAMI.bit_mask = 0;
		 servedGUAMIList_ma[i].value.gUAMI.pLMNIdentity.length = 0;
		 memset(servedGUAMIList_ma[i].value.gUAMI.pLMNIdentity.value, 0xff,
		 		sizeof(servedGUAMIList_ma[i].value.gUAMI.pLMNIdentity.value));
		 servedGUAMIList_ma[i].value.gUAMI.aMFRegionID.length = 0;
		 servedGUAMIList_ma[i].value.gUAMI.aMFRegionID.value = &(guamiAmfIdentifier_ma[i].regionId);
	     servedGUAMIList_ma[i].value.gUAMI.aMFSetID.length = 0;
		 servedGUAMIList_ma[i].value.gUAMI.aMFSetID.value = guamiAmfIdentifier_ma[i].setId;
		 servedGUAMIList_ma[i].value.gUAMI.aMFPointer.length = 0;
		 servedGUAMIList_ma[i].value.gUAMI.aMFPointer.value = &(guamiAmfIdentifier_ma[i].pointer);
		 servedGUAMIList_ma[i].value.gUAMI.iE_Extensions = NULL;	 
	}		 	

    //relative AMF Capacity
    relAmfCapacity_m = 0;
	relAmfCapacity_mp = NULL;

    //Supported TAs
    plmnSupportList_m = NULL;
    plmnSupportList_mp = NULL;
	numOfPlmnSupportListItem_m = 0;

    for (unsigned int i = 0; i < maxnoofPLMNs; i++)
    {
        plmnSupportList_ma[i].next = NULL;
		plmnSupportList_ma[i].value.bit_mask = 0;
		plmnSupportList_ma[i].value.iE_Extensions = NULL;
		plmnSupportList_ma[i].value.pLMNIdentity.length = 0;
		memset(plmnSupportList_ma[i].value.pLMNIdentity.value, 0xff,
			sizeof(plmnSupportList_ma[i].value.pLMNIdentity.value));
		plmnSupportList_ma[i].value.sliceSupportList = NULL;

		numOfSliceSupportItem_m[i] = 0;
        for (unsigned int j = 0; j < maxnoofSlices_c; j++)
        {
		    sliceSupportItem_ma[i][j].next =  NULL;
			sliceSupportItem_ma[i][j].value.bit_mask = 0;
			sliceSupportItem_ma[i][j].value.iE_Extensions = NULL;
			sliceSupportItem_ma[i][j].value.s_NSSAI.bit_mask = 0;
			sliceSupportItem_ma[i][j].value.s_NSSAI.iE_Extensions = NULL;
			sliceSupportItem_ma[i][j].value.s_NSSAI.sD.length = 0;
			memset(sliceSupportItem_ma[i][j].value.s_NSSAI.sD.value,0,
				   sizeof(sliceSupportItem_ma[i][j].value.s_NSSAI.sD.value));
			sliceSupportItem_ma[i][j].value.s_NSSAI.sST.length = 0;
			memset(sliceSupportItem_ma[i][j].value.s_NSSAI.sST.value,0,
				   sizeof(sliceSupportItem_ma[i][j].value.s_NSSAI.sST.value));
		}
    }
    //Ue Retention Information
    UeRetentionInfo_mp = NULL;
    UeRetentionInfoPresent_m = false;
    UeRetentionInfo_m = ues_retained;
}

unsigned int
NGSetupResponsePdu::encodeOssToPer(NgapOssCtxt* ctxt)
{
    if(NULL == ctxt)
	{
	    return -1;
    }
	
    ProtocolIeType *nextIe_p = NULL;

	//plmn support list
    ProtocolIeType plmnSupportIe;
    if (numOfPlmnSupportListItem_m > maxnoofPLMNs ||
        numOfPlmnSupportListItem_m <= 0)
    {
        return defaultFailureCode_c;
    }
    else
    {
        struct PLMNSupportList_ *nextPlmnSupportItem_p = NULL;
        for (unsigned int i = 0; i < numOfPlmnSupportListItem_m; i++)
        {
			struct SliceSupportList_ *slices_p = NULL;
			for(unsigned int j=0; j < numOfSliceSupportItem_m[i]; j++)
			{
				sliceSupportItem_ma[i][j].next = slices_p;
				slices_p = &sliceSupportItem_ma[i][j];
			}
			plmnSupportList_ma[i].value.sliceSupportList = slices_p;
			
            plmnSupportList_ma[i].next = nextPlmnSupportItem_p;
			nextPlmnSupportItem_p = &plmnSupportList_ma[i];
		}

        plmnSupportList_m = nextPlmnSupportItem_p;
    }
    plmnSupportIe.next = nextIe_p;
    plmnSupportIe.value.id = id_PLMNSupportList;
    plmnSupportIe.value.criticality = reject;
    plmnSupportIe.value.value.pduNum = PDU_NGSetupResponseIEs_Value_PLMNSupportList;
    plmnSupportIe.value.value.decoded.pdu_PLMNSupportList = &plmnSupportList_m;
    plmnSupportIe.value.value.encoded.length = 0;
    plmnSupportIe.value.value.encoded.value = NULL;	
    nextIe_p = &plmnSupportIe;

	//relative amf capacity
    ProtocolIeType relAmfCapacityIe;
    relAmfCapacityIe.next = nextIe_p;
    relAmfCapacityIe.value.id = id_RelativeAMFCapacity;
    relAmfCapacityIe.value.criticality = ignore;
    relAmfCapacityIe.value.value.pduNum =PDU_NGSetupResponseIEs_Value_RelativeAMFCapacity;
    relAmfCapacityIe.value.value.decoded.pdu_RelativeAMFCapacity = &relAmfCapacity_m;
    relAmfCapacityIe.value.value.encoded.length = 0;
    relAmfCapacityIe.value.value.encoded.value = NULL;
    nextIe_p = &relAmfCapacityIe;

	//Served GUAMI List encode
	ProtocolIeType servedGUAMIListIe;
	if(numOfSrvGUAMIListItem_m > maxnoofServedGUAMIs_c ||
		numOfSrvGUAMIListItem_m <= 0)
	{
		return defaultFailureCode_c;
	}
	else
	{
		struct ServedGUAMIList_ *nextServedGUAMI_p = NULL;
		for(int i = 0; i < numOfSrvGUAMIListItem_m; i++)
		{
			servedGUAMIList_ma[i].next = nextServedGUAMI_p;
			nextServedGUAMI_p = &servedGUAMIList_ma[i];
		}

		servedGUAMIList_m = nextServedGUAMI_p;
	}
	servedGUAMIListIe.next = nextIe_p;
	servedGUAMIListIe.value.id = id_ServedGUAMIList;
	servedGUAMIListIe.value.criticality = reject;
	servedGUAMIListIe.value.value.pduNum = PDU_NGSetupResponseIEs_Value_ServedGUAMIList;
	servedGUAMIListIe.value.value.decoded.pdu_ServedGUAMIList = &servedGUAMIList_m;
	servedGUAMIListIe.value.value.encoded.length = 0;
	servedGUAMIListIe.value.value.encoded.value = NULL;
    nextIe_p = &servedGUAMIListIe;

    //Amf Name IE encode
    ProtocolIeType amfNameIe;
    amfNameIe.next = nextIe_p;
    amfNameIe.value.id = id_AMFName;
    amfNameIe.value.criticality = reject;
    amfNameIe.value.value.pduNum = PDU_NGSetupResponseIEs_Value_AMFName;
    amfNameIe.value.value.decoded.pdu_AMFName = &amfName_m;
    amfNameIe.value.value.encoded.length = 0;
    amfNameIe.value.value.encoded.value = NULL;
	nextIe_p = &amfNameIe;

    //Ue Retention Information IE encode
	ProtocolIeType ueRetentionInfoIe;
	if (UeRetentionInfoPresent_m == true)
    {
        ueRetentionInfoIe.next = nextIe_p;
        ueRetentionInfoIe.value.id = id_UERetentionInformation;
        ueRetentionInfoIe.value.criticality = ignore;
        ueRetentionInfoIe.value.value.pduNum = PDU_NGSetupResponseIEs_Value_UERetentionInformation;
        ueRetentionInfoIe.value.value.decoded.pdu_UERetentionInformation = &UeRetentionInfo_m;
        ueRetentionInfoIe.value.value.encoded.length = 0;
        ueRetentionInfoIe.value.value.encoded.value = NULL;
	    nextIe_p = &ueRetentionInfoIe;
	}

    NGSetupResponse msg;
    msg.protocolIEs = nextIe_p;
	
    NGAP_PDU ngApPdu;

    ngApPdu.choice = successfulOutcome_chosen;
    ngApPdu.u.successfulOutcome.procedureCode = id_NGSetup;
    ngApPdu.u.successfulOutcome.criticality = reject;
    ngApPdu.u.successfulOutcome.value.pduNum = PDU_NGAP_ELEMENTARY_PROCEDURES_SuccessfulOutcome_NGSetupResponse;
    ngApPdu.u.successfulOutcome.value.decoded.pdu_NGSetupResponse = &msg;
    ngApPdu.u.successfulOutcome.value.encoded.length = 0;
    ngApPdu.u.successfulOutcome.value.encoded.value = NULL;
	
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
NGSetupResponsePdu::getDecodedOssDataPtrs_v()
{
    ProtocolIeType *currIe_p;
    reset_v();

    ngapPdu_mp = reinterpret_cast<NGAP_PDU *> (ossBufRef_m.value);
    if (ngapPdu_mp == NULL)
    {
        return;
    }

    NGSetupResponse_mp = 
             ngapPdu_mp->u.successfulOutcome.value.decoded.pdu_NGSetupResponse;

    if (NGSetupResponse_mp == NULL)
    {
        return;
    }

    currIe_p = NGSetupResponse_mp->protocolIEs;

    while (currIe_p != NULL)
    {
        if (currIe_p->value.id == id_AMFName
            && currIe_p->value.value.decoded.pdu_AMFName != NULL)
        {
            amfName_mp = 
                      currIe_p->value.value.decoded.pdu_AMFName;
        }
        else if (currIe_p->value.id == id_ServedGUAMIList
                 && currIe_p->value.value.decoded.pdu_ServedGUAMIList != NULL)
        {
            servedGUAMIList_mp = currIe_p->value.value.decoded.pdu_ServedGUAMIList;
			ServedGUAMIList curr_p = *servedGUAMIList_mp;
			while(curr_p != NULL)
			{
				numOfSrvGUAMIListItem_m++;
				curr_p = curr_p->next;
			}	
			
        }
        else if (currIe_p->value.id == id_RelativeAMFCapacity
                 && currIe_p->value.value.decoded.pdu_RelativeAMFCapacity != NULL)
        {
        	relAmfCapacity_mp = currIe_p->value.value.decoded.pdu_RelativeAMFCapacity;
        }
		else if(currIe_p->value.id == id_PLMNSupportList
                 && currIe_p->value.value.decoded.pdu_PLMNSupportList != NULL)
		{
            plmnSupportList_mp = 
                     currIe_p->value.value.decoded.pdu_PLMNSupportList;

			SliceSupportList  slices_p = NULL;				
			PLMNSupportList curr_p = *plmnSupportList_mp;
            while (curr_p != NULL)
            {
                slices_p = curr_p->value.sliceSupportList;
				while(slices_p != NULL)
				{
					numOfSliceSupportItem_m[numOfPlmnSupportListItem_m]++;
					slices_p = slices_p->next;
				}
                numOfPlmnSupportListItem_m++;
                curr_p = curr_p->next;
            }
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

NGSetupResponse *
NGSetupResponsePdu::getNGSetupResponseDecodedPtr() const
{
    return NGSetupResponse_mp;
}

AMFName * 
NGSetupResponsePdu::getAmfNameDecodedPtr() const
{
    return amfName_mp;
}

AMFName * 
NGSetupResponsePdu::getAmfNameEncodePtr()
{
    return &amfName_m;
}

RelativeAMFCapacity* 
NGSetupResponsePdu::getRelAmfCapacityDecodedPtr() const
{
	return relAmfCapacity_mp;
}

RelativeAMFCapacity* 
NGSetupResponsePdu::getRelAmfCapacityEncodePtr()
{
	return &relAmfCapacity_m;
}

ServedGUAMIList* 
NGSetupResponsePdu::getServedGUAMIListDecodedPtr() const
{
	return servedGUAMIList_mp;
}

ServedGUAMIList* 
NGSetupResponsePdu::getServedGUAMIListEncodePtr()
{
	return &servedGUAMIList_m;
}

ServedGUAMIItem* 
NGSetupResponsePdu::getServedGUAMIListItemDecodedPtr(unsigned int index) const
{
    ServedGUAMIItem *item_p = NULL;
	
    ServedGUAMIList servedGUAMIs_p = NULL;

    if (servedGUAMIList_mp != NULL)
    {
        servedGUAMIs_p = *(servedGUAMIList_mp);
    }

    if (servedGUAMIs_p != NULL
        && (index >= 0 && index < numOfSrvGUAMIListItem_m))
    {
        unsigned int i = 0;
        for (; i < index && servedGUAMIs_p != NULL; i++)
        {
            servedGUAMIs_p = servedGUAMIs_p->next;
        }
        if (servedGUAMIs_p != NULL && i == index)
        {
            item_p = &(servedGUAMIs_p->value);
        }
    }

    return item_p;

}

ServedGUAMIItem* 
NGSetupResponsePdu::getServedGUAMIListItemEncodedPtr(unsigned int index)
{
    if (index >= 0 && index < maxnoofServedGUAMIs)
    {
        return &(servedGUAMIList_ma[index].value);
    }
    else
    {
        return NULL;
    }
}

void 
NGSetupResponsePdu::setNumofServedGUAMIListItem_v(unsigned int num)
{
	numOfSrvGUAMIListItem_m = num;
}

unsigned int 
NGSetupResponsePdu::getNumofServedGUAMIListItem() const
{
	return numOfSrvGUAMIListItem_m;
}

PLMNSupportList * 
NGSetupResponsePdu::getPlmnSupportListDecodedPtr() const
{
	return plmnSupportList_mp;
}
PLMNSupportList * 
NGSetupResponsePdu::getPlmnSupportListEncodePtr()
{
	return &plmnSupportList_m;
}

PLMNSupportItem * 
NGSetupResponsePdu::getPlmnSupportItemDecodedPtr(unsigned int index) const
{
    PLMNSupportItem *item_p = NULL;
	
    PLMNSupportList plmnSupportList_p = NULL;

    if (plmnSupportList_mp != NULL)
    {
        plmnSupportList_p = *(plmnSupportList_mp);
    }

    if (plmnSupportList_p != NULL
        && (index >= 0 && index < numOfPlmnSupportListItem_m))
    {
        unsigned int i = 0;
        for (; i < index && plmnSupportList_p != NULL; i++)
        {
            plmnSupportList_p = plmnSupportList_p->next;
        }
        if (plmnSupportList_p != NULL && i == index)
        {
            item_p = &(plmnSupportList_p->value);
        }
    }

    return item_p;

}

PLMNSupportItem * 
NGSetupResponsePdu::getPlmnSupportItemEncodePtr(unsigned int index)
{
    if (index >= 0 && index < maxnoofPLMNs)
    {
        return &(plmnSupportList_ma[index].value);
    }
    else
    {
        return NULL;
    }
}

void 
NGSetupResponsePdu::setNumofPlmnSupportListItem_m_v(unsigned int num)
{
	numOfPlmnSupportListItem_m = num;
}

unsigned int 
NGSetupResponsePdu::getNumofPlmnSupportListItem() const
{
	return numOfPlmnSupportListItem_m;
}

SliceSupportItem* 
NGSetupResponsePdu::getSliceSupportListDecodedPtr(PLMNSupportItem* plmnSupportItemPtr,
												unsigned int indSlices) const
{
    SliceSupportItem* item_p = NULL;
    SliceSupportList sliceList_p = NULL;
    if (plmnSupportItemPtr != NULL)
    {
        sliceList_p = plmnSupportItemPtr->sliceSupportList;
        if (sliceList_p != NULL
            && (indSlices >= 0 && indSlices < maxnoofSlices_c))
        {
            unsigned int i = 0;
            for (; i < indSlices && sliceList_p != NULL; i++)
            {
                sliceList_p = sliceList_p->next;
            }

		    if(sliceList_p != NULL  && i == indSlices)
		    {
				item_p = &(sliceList_p->value);
		    }
        }
    }
	
    return item_p;
}
												
SliceSupportItem* 
NGSetupResponsePdu::getSliceSupportListEncodePtr( 		unsigned int indPlmn, unsigned int indSlices)
{
    if ((indPlmn < maxnoofPLMNs) &&
        (indSlices < maxnoofSlices_c))
    {
        return &(sliceSupportItem_ma[indPlmn][indSlices].value);
    }
    else
    {
        return NULL;
    }

}

void 
NGSetupResponsePdu::setNumofSlices_v(unsigned int num, unsigned int indPlmn)
{
    if (indPlmn < maxnoofPLMNs)
    {
	    numOfSliceSupportItem_m[indPlmn] = num;
    }	
}

unsigned int 
NGSetupResponsePdu::getNumofSlices(unsigned int indPlmn) const
{
    if (indPlmn < maxnoofPLMNs)
    {
		return numOfSliceSupportItem_m[indPlmn];
    }
	else
	{
		return 0;
	}
}

UERetentionInformation *
NGSetupResponsePdu::getUeRetentionInfoDecodedPtr() const
{
    return UeRetentionInfo_mp;
}

UERetentionInformation *
NGSetupResponsePdu::getUeRetentionInfoEncodePtr()
{
    UeRetentionInfoPresent_m = true;
    return &UeRetentionInfo_m;
}

bool
NGSetupResponsePdu::isUeRetentionInfoPresent() const
{
    return UeRetentionInfoPresent_m;
}

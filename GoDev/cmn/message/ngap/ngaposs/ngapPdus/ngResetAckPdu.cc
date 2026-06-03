#include <string.h>
#include <iostream>

#include "ngResetAckPdu.h"

NGResetAckPdu::NGResetAckPdu()
{
    reset_v();
}

NGResetAckPdu::~NGResetAckPdu()
{
}

void
NGResetAckPdu::reset_v()
{
    NGResetAck_mp = NULL;

    //UE Associated Logical NG Connection
    ueAssNGConnList_m = NULL;
    ueAssNGConnList_mp = NULL;
    ueAssNGConnPresent_m = false;
    numOfUEAssNGConnListItem_m = 0;
    for (unsigned int i = 0; i < maxnoofNGConnectionsToReset; i++)
    {
        ueAssNGConnList_ma[i].next = NULL;
        ueAssNGConnList_ma[i].value.bit_mask = 0;
        ueAssNGConnList_ma[i].value.iE_Extensions = NULL;
        ueAssNGConnList_ma[i].value.aMF_UE_NGAP_ID = 0;
        ueAssNGConnList_ma[i].value.rAN_UE_NGAP_ID = 0;
    }
}

unsigned int
NGResetAckPdu::encodeOssToPer(NgapOssCtxt* ctxt)                   
{
    if(NULL == ctxt)
	{
	    return -1;
    }
	
    ProtocolIeType *nextIe_p = NULL;

    //UE Associated Logical NG Connection
    ProtocolIeType ueAssNGConnIe;
    if (ueAssNGConnPresent_m) 
    {
        if (numOfUEAssNGConnListItem_m > maxnoofNGConnectionsToReset ||
            numOfUEAssNGConnListItem_m <= 0)
        {
            return defaultFailureCode_c;
        }
        else
        {
            struct UE_associatedLogicalNG_connectionList_ *nextUeAssNGConn_p = NULL;
            for (unsigned int i = 0; i < numOfUEAssNGConnListItem_m; i++)
            {
                ueAssNGConnList_ma[i].next = nextUeAssNGConn_p;
	    		nextUeAssNGConn_p = &ueAssNGConnList_ma[i];
            }
            ueAssNGConnList_m = nextUeAssNGConn_p;
    
            ueAssNGConnIe.value.value.pduNum = PDU_NGResetAcknowledgeIEs_Value_UE_associatedLogicalNG_connectionList;
            ueAssNGConnIe.value.value.decoded.pdu_UE_associatedLogicalNG_connectionList = &ueAssNGConnList_m;
            ueAssNGConnIe.value.value.encoded.length = 0;
            ueAssNGConnIe.value.value.encoded.value = NULL;
        }
        ueAssNGConnIe.next = nextIe_p;
        ueAssNGConnIe.value.id = id_UE_associatedLogicalNG_connectionList;
        ueAssNGConnIe.value.criticality = ignore;
        nextIe_p = &ueAssNGConnIe;
    }


    //message
    NGResetAcknowledge msg;
    msg.protocolIEs = nextIe_p;

    NGAP_PDU ngApPdu;

	ngApPdu.choice = successfulOutcome_chosen;
    ngApPdu.u.successfulOutcome.procedureCode = id_NGReset;
    ngApPdu.u.successfulOutcome.criticality = reject;
    ngApPdu.u.successfulOutcome.value.pduNum = PDU_NGAP_ELEMENTARY_PROCEDURES_SuccessfulOutcome_NGResetAcknowledge;
    ngApPdu.u.successfulOutcome.value.decoded.pdu_NGResetAcknowledge = &msg;
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
NGResetAckPdu::getDecodedOssDataPtrs_v()
{
    ProtocolIeType *currIe_p;
    reset_v();

    ngapPdu_mp = reinterpret_cast<NGAP_PDU *> (ossBufRef_m.value);
    if (ngapPdu_mp == NULL)
    {
        return;
    }
    
    NGResetAck_mp = ngapPdu_mp->u.successfulOutcome.value.decoded.pdu_NGResetAcknowledge;
    
    if (NGResetAck_mp == NULL)
    {
        return;
    }

    currIe_p = NGResetAck_mp->protocolIEs;

    while (currIe_p != NULL)
    {
        if (currIe_p->value.id == id_UE_associatedLogicalNG_connectionList
            && currIe_p->value.value.decoded.pdu_UE_associatedLogicalNG_connectionList != NULL)
        {
            ueAssNGConnList_mp = 
                     currIe_p->value.value.decoded.pdu_UE_associatedLogicalNG_connectionList;

            UE_associatedLogicalNG_connectionList curr_p = NULL;				
			curr_p = *ueAssNGConnList_mp;
            while (curr_p != NULL)
            {
                numOfUEAssNGConnListItem_m++;
                curr_p = curr_p->next;
            }
            ueAssNGConnPresent_m = true;
        }

        currIe_p = currIe_p->next;
    }
}

NGResetAcknowledge *
NGResetAckPdu::getNGResetAckDecodedPtr() const
{
    return NGResetAck_mp;
}

UE_associatedLogicalNG_connectionList * 
NGResetAckPdu::getUEAssNGConnListDecodedPtr() const
{
    return ueAssNGConnList_mp;
}

UE_associatedLogicalNG_connectionItem * 
NGResetAckPdu::getUEAssNGConnListItemDecodedPtr(unsigned int index) const
{
    UE_associatedLogicalNG_connectionItem *item_p = NULL;
	
    UE_associatedLogicalNG_connectionList ueAssNGConn_p = NULL;

    if (ueAssNGConnList_mp != NULL)
    {
        ueAssNGConn_p = *(ueAssNGConnList_mp);
    }

    if (ueAssNGConn_p != NULL
        && (index >= 0 && index < numOfUEAssNGConnListItem_m))
    {
        unsigned int i = 0;
        for (; i < index && ueAssNGConn_p != NULL; i++)
        {
            ueAssNGConn_p = ueAssNGConn_p->next;
        }
        if (ueAssNGConn_p != NULL && i == index)
        {
            item_p = &(ueAssNGConn_p->value);
        }
    }
    return item_p;
}

bool 
NGResetAckPdu::isUEAssNGConnPresent()
{
    return ueAssNGConnPresent_m;
}

UE_associatedLogicalNG_connectionList * 
NGResetAckPdu::getUEAssNGConnListEncodePtr()
{
    ueAssNGConnPresent_m = true;
    return &ueAssNGConnList_m;
}

UE_associatedLogicalNG_connectionItem * 
NGResetAckPdu::getUEAssNGConnListItemEncodePtr(unsigned int index)
{
    if (index >= 0 && index < maxnoofNGConnectionsToReset)
    {
        ueAssNGConnPresent_m = true;
        return &(ueAssNGConnList_ma[index].value);
    }
    else
    {
        return NULL;
    }

}

void 
NGResetAckPdu::setNumOfUEAssNGConnListItem_v(unsigned int num)
{
	numOfUEAssNGConnListItem_m = num;
}

unsigned int 
NGResetAckPdu::getNumOfUEAssNGConnListItem() const
{
   return numOfUEAssNGConnListItem_m;
}

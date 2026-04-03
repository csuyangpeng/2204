#include <string.h>
#include <iostream>

#include "downlinkNASTransportPdu.h"

DownlinkNASTransportPdu::DownlinkNASTransportPdu()
{
    reset_v();
}

DownlinkNASTransportPdu::~DownlinkNASTransportPdu()
{
}

void DownlinkNASTransportPdu::reset_v()
{
    downlinkNasTransport_mp = NULL;

    //RAN UE NGAP ID
    ranUeNgapId_m=0;
    ranUeNgapId_mp=NULL;

    //AMF UE NGAP ID
    amfUeNgapId_m=0;
    amfUeNgapId_mp=NULL;

    //NAS PDU
    nasPdu_m.length=0;
    memset(nasPdu_ma, 0xff, sizeof(nasPdu_ma));
    nasPdu_m.value=nasPdu_ma;
    nasPdu_mp=NULL;
    
    //RAN Paging Priority
    ranPagingPriority_m=0;
    ranPagingPriority_mp=NULL;
    ranPagingPriorityPresent_m = false;

    //Index To RFSP
    indexToRfsp_m=0;
    indexToRfsp_mp=NULL;
    indexToRfspPresent_m = false;

    //UE Aggregate Maximum Bit Rate
    ueAmbr_m.bit_mask = 0;
    ueAmbr_m.iE_Extensions = NULL;
    ueAmbr_m.uEAggregateMaximumBitRateUL = 0;
    ueAmbr_m.uEAggregateMaximumBitRateDL = 0;
    ueAmbr_mp=NULL;
    ueAmbrPresent_m = false;

    //Old AMF
    oldAmfName_mp = NULL;
    oldAmfName_m.length = 0;
    memset(oldAmfNameValue_ma, 0xff, sizeof(oldAmfNameValue_ma));
    oldAmfName_m.value = oldAmfNameValue_ma;
    oldAmfNamePresent_m = false;    

    //Allowed NSSAI
    allowedNssai_m=NULL;
    allowedNssai_mp=NULL;
    numOfAllowedNassiItem_m=0;
    allowedNssaiPresent_m = false;
    for(unsigned int i = 0; i < maxnoofAllowedS_NSSAIs; i++)
    {
        allowedNssai_ma[i].next=NULL;
        allowedNssai_ma[i].value.bit_mask=0; 
        allowedNssai_ma[i].value.iE_Extensions=NULL;
        allowedNssai_ma[i].value.s_NSSAI.bit_mask=0; 
        allowedNssai_ma[i].value.s_NSSAI.sST.length=0;
        memset(allowedNssai_ma[i].value.s_NSSAI.sST.value, 0xff,
                   sizeof(allowedNssai_ma[i].value.s_NSSAI.sST.value));
        allowedNssai_ma[i].value.s_NSSAI.sD.length=0; 
        memset(allowedNssai_ma[i].value.s_NSSAI.sD.value, 0xff,
                   sizeof(allowedNssai_ma[i].value.s_NSSAI.sD.value));        
        allowedNssai_ma[i].value.s_NSSAI.iE_Extensions=NULL;        
    }

}

unsigned int
DownlinkNASTransportPdu::encodeOssToPer(NgapOssCtxt* ctxt)                   
{
    if(NULL == ctxt)
    {
        return -1;
    }
    
    ProtocolIeType *nextIe_p = NULL;

	//Allowed NSSAI
    ProtocolIeType allowedNssaiIe;
    if(allowedNssaiPresent_m)
    {
       if(numOfAllowedNassiItem_m > maxnoofAllowedS_NSSAIs ||
          numOfAllowedNassiItem_m < 0)
       {
           return defaultFailureCode_c;
       }
       else
       {
           struct AllowedNSSAI_ *nextAllowedNssai_p = NULL;
           for(int i = 0; i < numOfAllowedNassiItem_m; i++)
           {
               allowedNssai_ma[i].next = nextAllowedNssai_p;
               nextAllowedNssai_p = &allowedNssai_ma[i];       
           }
           
           allowedNssai_m = nextAllowedNssai_p;
       }   
       allowedNssaiIe.next =nextIe_p;
       allowedNssaiIe.value.id = id_AllowedNSSAI;
       allowedNssaiIe.value.criticality = reject;
       allowedNssaiIe.value.value.pduNum =PDU_DownlinkNASTransport_IEs_Value_AllowedNSSAI;
       allowedNssaiIe.value.value.decoded.pdu_AllowedNSSAI = &allowedNssai_m;
       allowedNssaiIe.value.value.encoded.length = 0;
       allowedNssaiIe.value.value.encoded.value = NULL;
       nextIe_p = &allowedNssaiIe;
    }
	
	//UE Aggregate Maximum Bit Rate
    ProtocolIeType ueAmbrIe;
    if (ueAmbrPresent_m)
    {
       ueAmbrIe.next =nextIe_p;
       ueAmbrIe.value.id = id_UEAggregateMaximumBitRate;
       ueAmbrIe.value.criticality = ignore;
       ueAmbrIe.value.value.pduNum =PDU_DownlinkNASTransport_IEs_Value_UEAggregateMaximumBitRate;
       ueAmbrIe.value.value.decoded.pdu_UEAggregateMaximumBitRate = &ueAmbr_m;
       ueAmbrIe.value.value.encoded.length = 0;
       ueAmbrIe.value.value.encoded.value = NULL;
       nextIe_p = &ueAmbrIe;
    }
	
	//Index To RFSP 
    ProtocolIeType indexToRfspIe;
    if (indexToRfspPresent_m)
    {
       indexToRfspIe.next =nextIe_p;
       indexToRfspIe.value.id = id_IndexToRFSP;
       indexToRfspIe.value.criticality = ignore;
       indexToRfspIe.value.value.pduNum =PDU_DownlinkNASTransport_IEs_Value_IndexToRFSP;
       indexToRfspIe.value.value.decoded.pdu_IndexToRFSP = &indexToRfsp_m;
       indexToRfspIe.value.value.encoded.length = 0;
       indexToRfspIe.value.value.encoded.value = NULL;
       nextIe_p = &indexToRfspIe;
    }
	
	//NAS PDU
    ProtocolIeType nasPduIe;
    nasPduIe.next =nextIe_p;
    nasPduIe.value.id = id_NAS_PDU;
    nasPduIe.value.criticality = reject;
    nasPduIe.value.value.pduNum =PDU_DownlinkNASTransport_IEs_Value_NAS_PDU;
    nasPduIe.value.value.decoded.pdu_NAS_PDU = &nasPdu_m;
    nasPduIe.value.value.encoded.length = 0;
    nasPduIe.value.value.encoded.value = NULL;
    nextIe_p = &nasPduIe;
	
	//RAN Paging Priority
    ProtocolIeType ranPagingPriorityIe;
    if (ranPagingPriorityPresent_m)
    {
       ranPagingPriorityIe.next =nextIe_p;
       ranPagingPriorityIe.value.id = id_RANPagingPriority;
       ranPagingPriorityIe.value.criticality = ignore;
       ranPagingPriorityIe.value.value.pduNum =PDU_DownlinkNASTransport_IEs_Value_RANPagingPriority;
       ranPagingPriorityIe.value.value.decoded.pdu_RANPagingPriority = &ranPagingPriority_m;
       ranPagingPriorityIe.value.value.encoded.length = 0;
       ranPagingPriorityIe.value.value.encoded.value = NULL;
       nextIe_p = &ranPagingPriorityIe;
    }
	
	//OldAMF
    ProtocolIeType amfNameIe;
    if(oldAmfNamePresent_m)
    {
       amfNameIe.next = nextIe_p;
       amfNameIe.value.id = id_OldAMF;
       amfNameIe.value.criticality = reject;
       amfNameIe.value.value.pduNum = PDU_DownlinkNASTransport_IEs_Value_AMFName;
       amfNameIe.value.value.decoded.pdu_AMFName = &oldAmfName_m;
       amfNameIe.value.value.encoded.length = 0;
       amfNameIe.value.value.encoded.value = NULL;
       nextIe_p = &amfNameIe;  
    }
	
    //RAN_UE_NGAP_ID
    ProtocolIeType ranUeNgapIdIe;
    ranUeNgapIdIe.next =nextIe_p;
    ranUeNgapIdIe.value.id = id_RAN_UE_NGAP_ID;
    ranUeNgapIdIe.value.criticality = reject;
    ranUeNgapIdIe.value.value.pduNum =PDU_DownlinkNASTransport_IEs_Value_RAN_UE_NGAP_ID;
    ranUeNgapIdIe.value.value.decoded.pdu_RAN_UE_NGAP_ID = &ranUeNgapId_m;
    ranUeNgapIdIe.value.value.encoded.length = 0;
    ranUeNgapIdIe.value.value.encoded.value = NULL;
    nextIe_p = &ranUeNgapIdIe;

    //AMF_UE_NGAP_ID
    ProtocolIeType amfUeNgapIdIe;
    amfUeNgapIdIe.next =nextIe_p;
    amfUeNgapIdIe.value.id = id_AMF_UE_NGAP_ID;
    amfUeNgapIdIe.value.criticality = reject;
    amfUeNgapIdIe.value.value.pduNum =PDU_DownlinkNASTransport_IEs_Value_AMF_UE_NGAP_ID;
    amfUeNgapIdIe.value.value.decoded.pdu_AMF_UE_NGAP_ID = &amfUeNgapId_m;
    amfUeNgapIdIe.value.value.encoded.length = 0;
    amfUeNgapIdIe.value.value.encoded.value = NULL;
    nextIe_p = &amfUeNgapIdIe;
    
    // message
    DownlinkNASTransport msg;
    msg.protocolIEs = nextIe_p;

    NGAP_PDU ngApPdu;
    ngApPdu.choice = initiatingMessage_chosen;
    ngApPdu.u.initiatingMessage.procedureCode = id_DownlinkNASTransport;
    ngApPdu.u.initiatingMessage.criticality = ignore;
    ngApPdu.u.initiatingMessage.value.pduNum = PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_DownlinkNASTransport;
    ngApPdu.u.initiatingMessage.value.decoded.pdu_DownlinkNASTransport = &msg;
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

void DownlinkNASTransportPdu::getDecodedOssDataPtrs_v()
{
    ProtocolIeType *currIe_p;
    reset_v();

    ngapPdu_mp = reinterpret_cast<NGAP_PDU *> (ossBufRef_m.value);
    if (ngapPdu_mp == NULL)
    {
        return;
    }
    
    downlinkNasTransport_mp =
            ngapPdu_mp->u.initiatingMessage.value.decoded.pdu_DownlinkNASTransport;
    
    if (downlinkNasTransport_mp == NULL)
    {
        return;
    }

    currIe_p = downlinkNasTransport_mp->protocolIEs;

    while (currIe_p != NULL)
    {
        if (currIe_p->value.id == id_RAN_UE_NGAP_ID
            && currIe_p->value.value.decoded.pdu_RAN_UE_NGAP_ID != NULL)
        {
            ranUeNgapId_mp = 
                      currIe_p->value.value.decoded.pdu_RAN_UE_NGAP_ID;
        }
        else if (currIe_p->value.id == id_AMF_UE_NGAP_ID
            && currIe_p->value.value.decoded.pdu_AMF_UE_NGAP_ID != NULL)
        {
            amfUeNgapId_mp = 
                      currIe_p->value.value.decoded.pdu_AMF_UE_NGAP_ID;
        }
        else if (currIe_p->value.id == id_NAS_PDU
                 && currIe_p->value.value.decoded.pdu_NAS_PDU != NULL)
        {
            nasPdu_mp = currIe_p->value.value.decoded.pdu_NAS_PDU;
        }
        else if (currIe_p->value.id == id_AllowedNSSAI
                 && currIe_p->value.value.decoded.pdu_AllowedNSSAI != NULL)
        {
            allowedNssai_mp = currIe_p->value.value.decoded.pdu_AllowedNSSAI;
            AllowedNSSAI curr_p = *allowedNssai_mp;
            while(curr_p != NULL)
            {
                numOfAllowedNassiItem_m++;
                curr_p = curr_p->next;
            }
			allowedNssaiPresent_m = true;
        }
        else if (currIe_p->value.id == id_OldAMF
            && currIe_p->value.value.decoded.pdu_AMFName != NULL)
        {
            oldAmfName_mp = 
                currIe_p->value.value.decoded.pdu_AMFName;
			oldAmfNamePresent_m = true;
        }
        else if (currIe_p->value.id == id_IndexToRFSP
            && currIe_p->value.value.decoded.pdu_IndexToRFSP != NULL)
        {
            indexToRfsp_mp = 
               currIe_p->value.value.decoded.pdu_IndexToRFSP;
			indexToRfspPresent_m = true;
        }
        else if (currIe_p->value.id == id_RANPagingPriority
            && currIe_p->value.value.decoded.pdu_RANPagingPriority != NULL)
        {
            ranPagingPriority_mp = 
               currIe_p->value.value.decoded.pdu_RANPagingPriority;
			ranPagingPriorityPresent_m = true;
        }
        else if (currIe_p->value.id == id_UEAggregateMaximumBitRate
            && currIe_p->value.value.decoded.pdu_UEAggregateMaximumBitRate != NULL)
        {
            ueAmbr_mp = 
                currIe_p->value.value.decoded.pdu_UEAggregateMaximumBitRate;
			ueAmbrPresent_m = true;
        }
        currIe_p = currIe_p->next;
    }
	
}

DownlinkNASTransport *
DownlinkNASTransportPdu::getDownlinkNasTransportDecodedPtr() const
{
    return downlinkNasTransport_mp;
}

AMFName * 
DownlinkNASTransportPdu::getOldAmfNameDecodedPtr() const
{
    return oldAmfName_mp;
}

AMFName * 
DownlinkNASTransportPdu::getOldAmfNameEncodePtr()
{
    oldAmfNamePresent_m = true;
    return &oldAmfName_m;
}

bool
DownlinkNASTransportPdu::isOldAmfNamePresent()
{
    return oldAmfNamePresent_m;
}

RANPagingPriority * 
DownlinkNASTransportPdu::getRanPagingPriorityDecodedPtr() const
{
    return ranPagingPriority_mp;
}

RANPagingPriority * 
DownlinkNASTransportPdu::getRanPagingPriorityEncodePtr()
{
    ranPagingPriorityPresent_m = true;
    return &ranPagingPriority_m;
}

bool
DownlinkNASTransportPdu::isRanPagingPriorityPresent()
{
    return ranPagingPriorityPresent_m;
}

IndexToRFSP * 
DownlinkNASTransportPdu::getIndexToRfspDecodedPtr() const
{
    return indexToRfsp_mp;
}

IndexToRFSP * 
DownlinkNASTransportPdu::getIndexToRfspEncodePtr()
{
    indexToRfspPresent_m = true;
    return &indexToRfsp_m;
}

bool
DownlinkNASTransportPdu::isIndexToRfspPresent()
{
    return indexToRfspPresent_m;
}

UEAggregateMaximumBitRate  * 
DownlinkNASTransportPdu::getUeAmbrDecodedPtr() const
{
    return ueAmbr_mp;
}

UEAggregateMaximumBitRate  * 
DownlinkNASTransportPdu::getUeAmbrEncodePtr()
{
    ueAmbrPresent_m = true;
    return &ueAmbr_m;
}

bool
DownlinkNASTransportPdu::isUeAmbrPresent()
{
    return ueAmbrPresent_m;
}

RAN_UE_NGAP_ID * 
DownlinkNASTransportPdu::getRanUeNgapIdDecodedPtr() const
{
    return ranUeNgapId_mp;
}

RAN_UE_NGAP_ID * 
DownlinkNASTransportPdu::getRanUeNgapIdEncodePtr()
{
    return &ranUeNgapId_m;
}

AMF_UE_NGAP_ID * 
DownlinkNASTransportPdu::getAmfUeNgapIdDecodedPtr() const
{
    return amfUeNgapId_mp;
}

AMF_UE_NGAP_ID * 
DownlinkNASTransportPdu::getAmfUeNgapIdEncodePtr()
{
    return &amfUeNgapId_m;
}

NAS_PDU * 
DownlinkNASTransportPdu::getNasPduDecodedPtr() const
{
    return nasPdu_mp;
}

NAS_PDU * 
DownlinkNASTransportPdu::getNasPduEncodePtr()
{
    return &nasPdu_m;
}

AllowedNSSAI * 
DownlinkNASTransportPdu::getAllowedNssaiDecodedPtr() const
{
    return allowedNssai_mp;
}

AllowedNSSAI * 
DownlinkNASTransportPdu::getAllowedNssaiEncodePtr()
{
    return &allowedNssai_m;
}

bool
DownlinkNASTransportPdu::isAllowedNssaiPresent()
{
    return allowedNssaiPresent_m;
}

AllowedNSSAI_Item* 
DownlinkNASTransportPdu::getAllowedNssaiItemDecodedPtr(unsigned int index) const
{
    AllowedNSSAI_Item *item_p = NULL;

    AllowedNSSAI allowedNssai_p=NULL;

    if(allowedNssai_mp !=NULL)
    {
        allowedNssai_p=*(allowedNssai_mp);
    }

    if(allowedNssai_p !=NULL
       &&(index >=0) && index < numOfAllowedNassiItem_m)
    {
       unsigned int i = 0;
       for (; i < index && allowedNssai_p != NULL; i++)
       {
        allowedNssai_p = allowedNssai_p->next;
       }
       if (allowedNssai_p != NULL && i == index)
       {
        item_p = &(allowedNssai_p->value);
       }
    }

    return item_p;	   
}

AllowedNSSAI_Item * 
DownlinkNASTransportPdu::getAllowedNssaiItemEncodePtr(unsigned int index)
{
    if (index >= 0 && index < maxnoofAllowedS_NSSAIs)
    {
        allowedNssaiPresent_m = true;
        return &(allowedNssai_ma[index].value);
    }
    else
    {
        return NULL;
    }
}

void 
DownlinkNASTransportPdu::setNumofAllowedNssaiItem_v(unsigned int num)
{
    numOfAllowedNassiItem_m = num;
}
unsigned int 

DownlinkNASTransportPdu::getNumofAllowedNssaiItem() const
{
    return numOfAllowedNassiItem_m;
}


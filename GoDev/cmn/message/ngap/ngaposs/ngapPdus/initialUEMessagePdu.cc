#include <string.h>
#include <iostream>

#include "initialUEMessagePdu.h"

InitialUEMessagePdu::InitialUEMessagePdu()
{
    reset_v();
}

InitialUEMessagePdu::~InitialUEMessagePdu()
{
}

void
InitialUEMessagePdu::reset_v()
{
    InitialUEMessage_mp = NULL;

    //RAN UE NGAP ID
    ranUeNgapId_m=0;
    ranUeNgapId_mp=NULL;
    
    //NAS PDU
    nasPdu_m.length=0;
    memset(nasPdu_ma, 0xff, sizeof(nasPdu_ma));
    nasPdu_m.value=nasPdu_ma;
    nasPdu_mp=NULL;

    //NgRanCgi
    ngranCgi_m.choice = nR_CGI_chosen;
    ngranCgi_m.u.nR_CGI.bit_mask=0;
    ngranCgi_m.u.nR_CGI.iE_Extensions = NULL;
    ngranCgi_m.u.nR_CGI.pLMNIdentity.length=0;
    memset(ngranCgi_m.u.nR_CGI.pLMNIdentity.value, 0xff,
           sizeof(ngranCgi_m.u.nR_CGI.pLMNIdentity.value));
    ngranCgi_m.u.nR_CGI.nRCellIdentity.length=0;
    memset(nrCellId_mb,0xff,sizeof(nrCellId_mb));
    ngranCgi_m.u.nR_CGI.nRCellIdentity.value = nrCellId_mb;
    ngranCgi_m.u.nR_CGI.iE_Extensions=NULL;

    //Ie Extension
    userLocalInfoIe.next = NULL;
    userLocalInfoIe.value.id = id_PSCellInformation;
    userLocalInfoIe.value.criticality = ignore;
    userLocalInfoIe.value.extensionValue.pduNum = 
        PDU_UserLocationInformationEUTRA_ExtIEs_Extension_NGRAN_CGI;
    userLocalInfoIe.value.extensionValue.decoded.pdu_NGRAN_CGI = &ngranCgi_m;
    userLocalInfoIe.value.extensionValue.encoded.length = 0;
    userLocalInfoIe.value.extensionValue.encoded.value = NULL;

    //User Location Information
    //only NR supported currently
    useLocationInformation_m.choice=userLocationInformationNR_chosen;
    useLocationInformation_m.u.userLocationInformationNR.bit_mask=0;
    useLocationInformation_m.u.userLocationInformationNR.iE_Extensions = NULL;
    useLocationInformation_m.u.userLocationInformationNR.nR_CGI.bit_mask=0;
    useLocationInformation_m.u.userLocationInformationNR.nR_CGI.pLMNIdentity.length=0;
    memset(useLocationInformation_m.u.userLocationInformationNR.nR_CGI.pLMNIdentity.value, 0xff,
           sizeof(useLocationInformation_m.u.userLocationInformationNR.nR_CGI.pLMNIdentity.value));
    useLocationInformation_m.u.userLocationInformationNR.nR_CGI.nRCellIdentity.length=0;
    memset(nrCellId_ma,0xff,sizeof(nrCellId_ma));
    useLocationInformation_m.u.userLocationInformationNR.nR_CGI.nRCellIdentity.value = nrCellId_ma;
    useLocationInformation_m.u.userLocationInformationNR.nR_CGI.iE_Extensions=NULL;

    useLocationInformation_m.u.userLocationInformationNR.tAI.bit_mask=0;
    useLocationInformation_m.u.userLocationInformationNR.tAI.pLMNIdentity.length=0;
    memset(useLocationInformation_m.u.userLocationInformationNR.tAI.pLMNIdentity.value, 0xff,
           sizeof(useLocationInformation_m.u.userLocationInformationNR.tAI.pLMNIdentity.value));
    useLocationInformation_m.u.userLocationInformationNR.tAI.tAC.length=0;
    memset(useLocationInformation_m.u.userLocationInformationNR.tAI.tAC.value, 0xff,
           sizeof(useLocationInformation_m.u.userLocationInformationNR.tAI.tAC.value));
    useLocationInformation_m.u.userLocationInformationNR.tAI.iE_Extensions=NULL;

    useLocationInformation_m.u.userLocationInformationNR.timeStamp.length=0;
    memset(useLocationInformation_m.u.userLocationInformationNR.timeStamp.value, 0xff,
           sizeof(useLocationInformation_m.u.userLocationInformationNR.timeStamp.value));

    useLocationInformation_m.u.userLocationInformationNR.iE_Extensions = &userLocalInfoIe;

    useLocationInformation_mp=NULL;

    //FiveG S TMSI
    stmsi_m.bit_mask=0;
    memset(&stmsi5g_ma,0xff,sizeof(stmsi5g_ma));
    stmsi_m.aMFSetID.length=0;
    stmsi_m.aMFSetID.value=stmsi5g_ma.setId;
    stmsi_m.aMFPointer.length=0;
    stmsi_m.aMFPointer.value=&(stmsi5g_ma.pointer);
    stmsi_m.fiveG_TMSI.length=0;
    memset(stmsi_m.fiveG_TMSI.value, 0xff, sizeof(stmsi_m.fiveG_TMSI.value));
    stmsi_mp=NULL;
    stmsiPresent_m = false;

    //UE Context Request
    ueContextRequest_m=requested;
    ueContextRequest_mp=NULL;
    ueContextRequestPresent_m = false;

    //Rrc Establishment Cause
    rrcEstablishmentCause_mp=NULL;
    rrcEstablishmentCause_m=emergency;

    //AMF Set ID
    amfSetId_mp=NULL;
    amfSetId_m.length=0;
    memset(amfSetIdValue_ma, 0xff, sizeof(amfSetIdValue_ma));
    amfSetId_m.value=amfSetIdValue_ma;
    amfSetIdPresent_m =  false;

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
InitialUEMessagePdu::encodeOssToPer(NgapOssCtxt* ctxt)                   
{
    if(NULL == ctxt)
    {
        std::cout <<"Invalid NgapOssCtxt!" << std::endl;                  
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
       allowedNssaiIe.value.value.pduNum =PDU_InitialUEMessage_IEs_Value_AllowedNSSAI;
       allowedNssaiIe.value.value.decoded.pdu_AllowedNSSAI = &allowedNssai_m;
       allowedNssaiIe.value.value.encoded.length = 0;
       allowedNssaiIe.value.value.encoded.value = NULL;
       nextIe_p = &allowedNssaiIe;
    }

    //UE Context Request
    ProtocolIeType ueContextRequestIe;
    if(ueContextRequestPresent_m)
    {
        ueContextRequestIe.next =nextIe_p;
        ueContextRequestIe.value.id = id_UEContextRequest;
        ueContextRequestIe.value.criticality = ignore;
        ueContextRequestIe.value.value.pduNum =PDU_InitialUEMessage_IEs_Value_UEContextRequest;
        ueContextRequestIe.value.value.decoded.pdu_UEContextRequest = &ueContextRequest_m;
        ueContextRequestIe.value.value.encoded.length = 0;
        ueContextRequestIe.value.value.encoded.value = NULL;
        nextIe_p = &ueContextRequestIe;
    }

    //AMF Set ID
    ProtocolIeType amfSetIdIe;
    if(amfSetIdPresent_m)
    {
        amfSetIdIe.next =nextIe_p;
        amfSetIdIe.value.id = id_AMFSetID;
        amfSetIdIe.value.criticality = ignore;
        amfSetIdIe.value.value.pduNum =PDU_InitialUEMessage_IEs_Value_AMFSetID;
        amfSetIdIe.value.value.decoded.pdu_AMFSetID = &amfSetId_m;
        amfSetIdIe.value.value.encoded.length = 0;
        amfSetIdIe.value.value.encoded.value = NULL;
        nextIe_p = &amfSetIdIe;
    }

    //FiveG S TMSI
    ProtocolIeType stmsiIe;
    if(stmsiPresent_m)
    {
        stmsiIe.next =nextIe_p;
        stmsiIe.value.id = id_FiveG_S_TMSI;
        stmsiIe.value.criticality = reject;
        stmsiIe.value.value.pduNum =PDU_InitialUEMessage_IEs_Value_FiveG_S_TMSI;
        stmsiIe.value.value.decoded.pdu_FiveG_S_TMSI = &stmsi_m;
        stmsiIe.value.value.encoded.length = 0;
        stmsiIe.value.value.encoded.value = NULL;
        nextIe_p = &stmsiIe;
    }

    //RRCE Stablishment Cause
    ProtocolIeType rrcEstablishmentCauseIe;
    rrcEstablishmentCauseIe.next =nextIe_p;
    rrcEstablishmentCauseIe.value.id = id_RRCEstablishmentCause;
    rrcEstablishmentCauseIe.value.criticality = ignore;
    rrcEstablishmentCauseIe.value.value.pduNum =PDU_InitialUEMessage_IEs_Value_RRCEstablishmentCause;
    rrcEstablishmentCauseIe.value.value.decoded.pdu_RRCEstablishmentCause = &rrcEstablishmentCause_m;
    rrcEstablishmentCauseIe.value.value.encoded.length = 0;
    rrcEstablishmentCauseIe.value.value.encoded.value = NULL;
    nextIe_p = &rrcEstablishmentCauseIe;

    //User Location Information
    ProtocolIeType useLocationInformationIe;
    useLocationInformationIe.next =nextIe_p;
    useLocationInformationIe.value.id = id_UserLocationInformation;
    useLocationInformationIe.value.criticality = reject;
    useLocationInformationIe.value.value.pduNum =PDU_InitialUEMessage_IEs_Value_UserLocationInformation;
    useLocationInformationIe.value.value.decoded.pdu_UserLocationInformation = &useLocationInformation_m;
    useLocationInformationIe.value.value.encoded.length = 0;
    useLocationInformationIe.value.value.encoded.value = NULL;
    nextIe_p = &useLocationInformationIe;

    //NAS PDU
    ProtocolIeType nasPduIe;
    nasPduIe.next =nextIe_p;
    nasPduIe.value.id = id_NAS_PDU;
    nasPduIe.value.criticality = reject;
    nasPduIe.value.value.pduNum =PDU_InitialUEMessage_IEs_Value_NAS_PDU;
    nasPduIe.value.value.decoded.pdu_NAS_PDU = &nasPdu_m;
    nasPduIe.value.value.encoded.length = 0;
    nasPduIe.value.value.encoded.value = NULL;
    nextIe_p = &nasPduIe;

    //RAN_UE_NGAP_ID
    ProtocolIeType ranUeNgapIdIe;
    ranUeNgapIdIe.next =nextIe_p;
    ranUeNgapIdIe.value.id = id_RAN_UE_NGAP_ID;
    ranUeNgapIdIe.value.criticality = reject;
    ranUeNgapIdIe.value.value.pduNum =PDU_InitialUEMessage_IEs_Value_RAN_UE_NGAP_ID;
    ranUeNgapIdIe.value.value.decoded.pdu_RAN_UE_NGAP_ID = &ranUeNgapId_m;
    ranUeNgapIdIe.value.value.encoded.length = 0;
    ranUeNgapIdIe.value.value.encoded.value = NULL;
    nextIe_p = &ranUeNgapIdIe;
    
    //message
    InitialUEMessage msg;
    msg.protocolIEs = nextIe_p;
    
    NGAP_PDU ngApPdu;

    ngApPdu.choice = initiatingMessage_chosen;
    ngApPdu.u.initiatingMessage.procedureCode = id_InitialUEMessage;
    ngApPdu.u.initiatingMessage.criticality = ignore;
    ngApPdu.u.initiatingMessage.value.pduNum = PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_InitialUEMessage;
    ngApPdu.u.initiatingMessage.value.decoded.pdu_InitialUEMessage = &msg;
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
InitialUEMessagePdu::getDecodedOssDataPtrs_v()
{
    ProtocolIeType *currIe_p;
    reset_v();

    ngapPdu_mp = reinterpret_cast<NGAP_PDU *> (ossBufRef_m.value);
    if (ngapPdu_mp == NULL)
    {
        return;
    }
    
    InitialUEMessage_mp =
            ngapPdu_mp->u.initiatingMessage.value.decoded.pdu_InitialUEMessage;
    
    if (InitialUEMessage_mp == NULL)
    {
        return;
    }

    currIe_p = InitialUEMessage_mp->protocolIEs;

    while (currIe_p != NULL)
    {
        if (currIe_p->value.id == id_RAN_UE_NGAP_ID
            && currIe_p->value.value.decoded.pdu_RAN_UE_NGAP_ID != NULL)
        {
            ranUeNgapId_mp = 
                      currIe_p->value.value.decoded.pdu_RAN_UE_NGAP_ID;
        }
        else if (currIe_p->value.id == id_NAS_PDU
                 && currIe_p->value.value.decoded.pdu_NAS_PDU != NULL)
        {
            nasPdu_mp = currIe_p->value.value.decoded.pdu_NAS_PDU;
        }
        else if (currIe_p->value.id == id_UserLocationInformation
                 && currIe_p->value.value.decoded.pdu_UserLocationInformation != NULL)
        {
            useLocationInformation_mp = currIe_p->value.value.decoded.pdu_UserLocationInformation;
        }
        else if (currIe_p->value.id == id_FiveG_S_TMSI
                 && currIe_p->value.value.decoded.pdu_FiveG_S_TMSI != NULL)
        {
            stmsi_mp = currIe_p->value.value.decoded.pdu_FiveG_S_TMSI;
            stmsiPresent_m = true;
        }
        else if (currIe_p->value.id == id_UEContextRequest
                 && currIe_p->value.value.decoded.pdu_UEContextRequest != NULL)
        {
            ueContextRequest_mp = currIe_p->value.value.decoded.pdu_UEContextRequest;
            ueContextRequestPresent_m = true;
        }
        else if (currIe_p->value.id == id_RRCEstablishmentCause
                 && currIe_p->value.value.decoded.pdu_RRCEstablishmentCause != NULL)
        {
            rrcEstablishmentCause_mp = currIe_p->value.value.decoded.pdu_RRCEstablishmentCause;
        }
        else if (currIe_p->value.id == id_AMFSetID
                 && currIe_p->value.value.decoded.pdu_AMFSetID != NULL)
        {
            amfSetId_mp = currIe_p->value.value.decoded.pdu_AMFSetID;
            amfSetIdPresent_m = true;
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

        currIe_p = currIe_p->next;
    }
}

InitialUEMessage *
InitialUEMessagePdu::getInitialUEMessageDecodedPtr() const
{
    return InitialUEMessage_mp;
}

RAN_UE_NGAP_ID * 
InitialUEMessagePdu::getRanUeNgapIdDecodedPtr() const
{
    return ranUeNgapId_mp;
}

RAN_UE_NGAP_ID * 
InitialUEMessagePdu::getRanUeNgapIdEncodePtr()
{
    return &ranUeNgapId_m;
}

NAS_PDU * 
InitialUEMessagePdu::getNasPduDecodedPtr() const
{
    return nasPdu_mp;
}

NAS_PDU * 
InitialUEMessagePdu::getNasPduEncodePtr()
{
    return &nasPdu_m;
}

UserLocationInformation * 
InitialUEMessagePdu::getUseLocationInformationDecodedPtr() const
{
    return useLocationInformation_mp;
}

UserLocationInformation * 
InitialUEMessagePdu::getUseLocationInformationEncodePtr()
{
    return &useLocationInformation_m;
}

NGRAN_CGI* 
InitialUEMessagePdu::getNgRanCgiEncodePtr()
{
    return &ngranCgi_m;
}

UserLocationInfomationNRIeExtension* 
InitialUEMessagePdu::getIeExtenIeTypeDecodedPtr()
{
    return &userLocalInfoIe;
}

FiveG_S_TMSI * 
InitialUEMessagePdu::get5gStmsiDecodedPtr() const
{                           
    return stmsi_mp;
}

FiveG_S_TMSI * 
InitialUEMessagePdu::get5gStmsiEncodePtr()
{
    stmsiPresent_m = true;
    return &stmsi_m;
}

bool 
InitialUEMessagePdu::is5gStmsiPresent()
{
    return stmsiPresent_m;
}

UEContextRequest * 
InitialUEMessagePdu::getUeContextRequestDecodedPtr() const
{
    return ueContextRequest_mp;
}

 UEContextRequest * 
InitialUEMessagePdu::getUeContextRequestEncodePtr()
{
    ueContextRequestPresent_m = true;
    return &ueContextRequest_m;
}

bool 
InitialUEMessagePdu::isUeContextRequestPresent()
{
    return ueContextRequestPresent_m;
}

RRCEstablishmentCause * 
InitialUEMessagePdu::getRrcEstablishmentCauseDecodedPtr() const
{
    return rrcEstablishmentCause_mp;
}

RRCEstablishmentCause * 
InitialUEMessagePdu::getRrcEstablishmentCauseEncodePtr()
{
    return &rrcEstablishmentCause_m;
}

AMFSetID * 
InitialUEMessagePdu::getAmfSetIdDecodedPtr() const
{
    return amfSetId_mp;
}

AMFSetID * 
InitialUEMessagePdu::getAmfSetIdEncodePtr()
{
    amfSetIdPresent_m = true;
    return &amfSetId_m;
}

bool 
InitialUEMessagePdu::isAmfSetIdPresent()
{
    return amfSetIdPresent_m;
}

AllowedNSSAI * 
InitialUEMessagePdu::getAllowedNssaiDecodedPtr() const
{
    return allowedNssai_mp;
}

AllowedNSSAI * 
InitialUEMessagePdu::getAllowedNssaiEncodePtr()
{
    allowedNssaiPresent_m = true;
    return &allowedNssai_m;
}
bool 
InitialUEMessagePdu::isAllowedNssaiPresent()
{
    return allowedNssaiPresent_m;
}

AllowedNSSAI_Item* 
InitialUEMessagePdu::getAllowedNssaiItemDecodedPtr(unsigned int index) const
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
InitialUEMessagePdu::getAllowedNssaiItemEncodePtr(unsigned int index)
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
InitialUEMessagePdu::setNumofAllowedNssaiItem_v(unsigned int num)
{
    numOfAllowedNassiItem_m = num;
}
unsigned int 
InitialUEMessagePdu::getNumofAllowedNssaiItem() const
{
    return numOfAllowedNassiItem_m;
}

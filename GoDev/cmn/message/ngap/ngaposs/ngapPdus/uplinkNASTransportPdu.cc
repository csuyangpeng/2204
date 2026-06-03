#include <string.h>
#include <iostream>

#include "uplinkNASTransportPdu.h"

UplinkNASTransportPdu::UplinkNASTransportPdu()
{
    reset_v();
}

UplinkNASTransportPdu::~UplinkNASTransportPdu()
{
}

void
UplinkNASTransportPdu::reset_v()
{
    UplinkNASTransport_mp = NULL;
    
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
    
}

unsigned int
UplinkNASTransportPdu::encodeOssToPer(NgapOssCtxt* ctxt)                   
{
    if(NULL == ctxt)
    {
        return -1;
    }
    
    ProtocolIeType *nextIe_p = NULL;

    //User Location Information
    ProtocolIeType useLocationInformationIe;
    useLocationInformationIe.next =nextIe_p;
    useLocationInformationIe.value.id = id_UserLocationInformation;
    useLocationInformationIe.value.criticality = ignore;
    useLocationInformationIe.value.value.pduNum =PDU_UplinkNASTransport_IEs_Value_UserLocationInformation;
    useLocationInformationIe.value.value.decoded.pdu_UserLocationInformation = &useLocationInformation_m;
    useLocationInformationIe.value.value.encoded.length = 0;
    useLocationInformationIe.value.value.encoded.value = NULL;
    nextIe_p = &useLocationInformationIe;

    //NAS PDU
    ProtocolIeType nasPduIe;
    nasPduIe.next =nextIe_p;
    nasPduIe.value.id = id_NAS_PDU;
    nasPduIe.value.criticality = reject;
    nasPduIe.value.value.pduNum =PDU_UplinkNASTransport_IEs_Value_NAS_PDU;
    nasPduIe.value.value.decoded.pdu_NAS_PDU = &nasPdu_m;
    nasPduIe.value.value.encoded.length = 0;
    nasPduIe.value.value.encoded.value = NULL;
    nextIe_p = &nasPduIe;

    //RAN_UE_NGAP_ID
    ProtocolIeType ranUeNgapIdIe;
    ranUeNgapIdIe.next =nextIe_p;
    ranUeNgapIdIe.value.id = id_RAN_UE_NGAP_ID;
    ranUeNgapIdIe.value.criticality = reject;
    ranUeNgapIdIe.value.value.pduNum =PDU_UplinkNASTransport_IEs_Value_RAN_UE_NGAP_ID;
    ranUeNgapIdIe.value.value.decoded.pdu_RAN_UE_NGAP_ID = &ranUeNgapId_m;
    ranUeNgapIdIe.value.value.encoded.length = 0;
    ranUeNgapIdIe.value.value.encoded.value = NULL;
    nextIe_p = &ranUeNgapIdIe;

    //AMF_UE_NGAP_ID
    ProtocolIeType amfUeNgapIdIe;
    amfUeNgapIdIe.next =nextIe_p;
    amfUeNgapIdIe.value.id = id_AMF_UE_NGAP_ID;
    amfUeNgapIdIe.value.criticality = reject;
    amfUeNgapIdIe.value.value.pduNum =PDU_UplinkNASTransport_IEs_Value_AMF_UE_NGAP_ID;
    amfUeNgapIdIe.value.value.decoded.pdu_AMF_UE_NGAP_ID = &amfUeNgapId_m;
    amfUeNgapIdIe.value.value.encoded.length = 0;
    amfUeNgapIdIe.value.value.encoded.value = NULL;
    nextIe_p = &amfUeNgapIdIe;

    //message
    UplinkNASTransport msg;
    msg.protocolIEs = nextIe_p;

    NGAP_PDU ngApPdu;

    ngApPdu.choice = initiatingMessage_chosen;
    ngApPdu.u.initiatingMessage.procedureCode = id_UplinkNASTransport;
    ngApPdu.u.initiatingMessage.criticality = ignore;
    ngApPdu.u.initiatingMessage.value.pduNum = PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_UplinkNASTransport;
    ngApPdu.u.initiatingMessage.value.decoded.pdu_UplinkNASTransport = &msg;
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
UplinkNASTransportPdu::getDecodedOssDataPtrs_v()
{
    ProtocolIeType *currIe_p;
    reset_v();

    ngapPdu_mp = reinterpret_cast<NGAP_PDU *> (ossBufRef_m.value);
    if (ngapPdu_mp == NULL)
    {
        return;
    }
    
    UplinkNASTransport_mp =
            ngapPdu_mp->u.initiatingMessage.value.decoded.pdu_UplinkNASTransport;
    
    if (UplinkNASTransport_mp == NULL)
    {
        return;
    }

    currIe_p = UplinkNASTransport_mp->protocolIEs;

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
        else if (currIe_p->value.id == id_AMF_UE_NGAP_ID
            && currIe_p->value.value.decoded.pdu_AMF_UE_NGAP_ID != NULL)
        {
            amfUeNgapId_mp = 
                      currIe_p->value.value.decoded.pdu_AMF_UE_NGAP_ID;
        }
        else if (currIe_p->value.id == id_UserLocationInformation
                 && currIe_p->value.value.decoded.pdu_UserLocationInformation != NULL)
        {
            useLocationInformation_mp = currIe_p->value.value.decoded.pdu_UserLocationInformation;
        }

        currIe_p = currIe_p->next;
    }
}

UplinkNASTransport *
UplinkNASTransportPdu::getUplinkNASTransportDecodedPtr() const
{
    return UplinkNASTransport_mp;
}

RAN_UE_NGAP_ID * 
UplinkNASTransportPdu::getRanUeNgapIdDecodedPtr() const
{
    return ranUeNgapId_mp;
}

RAN_UE_NGAP_ID * 
UplinkNASTransportPdu::getRanUeNgapIdEncodePtr()
{
    return &ranUeNgapId_m;
}

AMF_UE_NGAP_ID * 
UplinkNASTransportPdu::getAmfUeNgapIdDecodedPtr() const
{
    return amfUeNgapId_mp;
}

AMF_UE_NGAP_ID * 
UplinkNASTransportPdu::getAmfUeNgapIdEncodePtr()
{
    return &amfUeNgapId_m;
}

NAS_PDU * 
UplinkNASTransportPdu::getNasPduDecodedPtr() const
{
    return nasPdu_mp;
}

NAS_PDU * 
UplinkNASTransportPdu::getNasPduEncodePtr()
{
    return &nasPdu_m;
}

UserLocationInformation * 
UplinkNASTransportPdu::getUseLocationInformationDecodedPtr() const
{
    return useLocationInformation_mp;
}

UserLocationInformation * 
UplinkNASTransportPdu::getUseLocationInformationEncodePtr()
{
    return &useLocationInformation_m;
}

NGRAN_CGI* 
UplinkNASTransportPdu::getNgRanCgiEncodePtr()
{
    return &ngranCgi_m;
}

UserLocationInfomationNRIeExtension* 
UplinkNASTransportPdu::getIeExtenIeTypeDecodedPtr()
{
    return &userLocalInfoIe;
}
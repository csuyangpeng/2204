#include <iostream>
#include <sstream>
#include <string>

#include "initialContextSetupRequestCodec.h"

using namespace std;

InitialContextSetupRequestCodec::InitialContextSetupRequestCodec() {
    reset();
}
InitialContextSetupRequestCodec::~InitialContextSetupRequestCodec(){
}
void InitialContextSetupRequestCodec::setRanUeNgapId(Uint32 id) {
    ranUeNgapId_m = id;
}
Uint32 InitialContextSetupRequestCodec::getRanUeNgapId() {
    return ranUeNgapId_m;
}
void InitialContextSetupRequestCodec::setAmfUeNgapId(Uint64 id) {
    amfUeNgapId_m = id;
}
Uint64 InitialContextSetupRequestCodec::getAmfUeNgapId() {
    return amfUeNgapId_m;
}
void InitialContextSetupRequestCodec::setNasPdu(std::string naspdu) {
    nasPduPrst_m=true;
    nasPdu_m = naspdu;
}
std::string InitialContextSetupRequestCodec::getNasPdu() { 
    return nasPdu_m; 
}
bool InitialContextSetupRequestCodec::isNasPduPrst(){
    return nasPduPrst_m;
}
void InitialContextSetupRequestCodec::setOldAmfName(std::string name) {
    amfNamePrst_m=true;
    oldAmfName_m = name;
}
std::string InitialContextSetupRequestCodec::getOldAmfName() {
    return oldAmfName_m;
}
bool InitialContextSetupRequestCodec::isAmfNamePrst(){
    return amfNamePrst_m;
}
void InitialContextSetupRequestCodec::setSecurityKey(std::string security) {
    securityKey_m = security;
}
std::string InitialContextSetupRequestCodec::getSecurityKey() {
    return securityKey_m;
}
void InitialContextSetupRequestCodec::setIndexToRfsp(long index) {
    indexToRFSPPrst_m=true;
    indexToRfsp_m = index;
}
long InitialContextSetupRequestCodec::getIndexToRfsp() { 
    return indexToRfsp_m; 
}
bool InitialContextSetupRequestCodec::isIndexToRFSPPrst(){
    return indexToRFSPPrst_m;
}
void InitialContextSetupRequestCodec::setUeAmbr(UeAmbr& ambr) {
    ueAMBRPrst_m=true;
    ueAmbr_m = ambr;
}
UeAmbr& InitialContextSetupRequestCodec::getUeAmbr() { 
    return ueAmbr_m; 
}
bool InitialContextSetupRequestCodec::isUeAMBRPrst(){
    return ueAMBRPrst_m;
}
void InitialContextSetupRequestCodec::setGuami(Guami& val) {
    guami_m = val;
}
Guami& InitialContextSetupRequestCodec::getGuami() { 
    return guami_m; 
}
void InitialContextSetupRequestCodec::setUeSecurityCap(UeSecurityCap& val) {
    ueSecurityCap_m = val;
}
UeSecurityCap& InitialContextSetupRequestCodec::getUeSecurityCap() {
    return ueSecurityCap_m;
}
void InitialContextSetupRequestCodec::setTraceAct(TraceAct& val) {
    traceActPrst_m=true;
    traceAct_m = val;
}
TraceAct& InitialContextSetupRequestCodec::getTraceAct(){ 
    return traceAct_m; 
}
bool InitialContextSetupRequestCodec::isTraceActPrst(){
    return traceActPrst_m;
}
std::string InitialContextSetupRequestCodec::getUERadioCapability(){
    return ueRadioCap_m;
}
void InitialContextSetupRequestCodec::setUERadioCapability(std::string val){
    ueRadioCapPrst_m=true;
    ueRadioCap_m = val;
}
bool InitialContextSetupRequestCodec::isUeRadioCapPrst(){
    return ueRadioCapPrst_m;
}
void InitialContextSetupRequestCodec::setMaskedImeiSv(std::string val) {
    maskedImeiSvPrst_m=true;
    maskImeiSv_m = val;
}
std::string InitialContextSetupRequestCodec::getMaskedImeiSv() {
    return maskImeiSv_m;
}
bool InitialContextSetupRequestCodec::isMaskedImeiSvPrst(){
    return maskedImeiSvPrst_m;
}
void InitialContextSetupRequestCodec::setEmergeFBInd(EmergeFBReqInd & val){
    emergeFBReqIndPrst_m = true;
    emergeFBReqInd_m.indicator =val.indicator;
    emergeFBReqInd_m.serviceTargetCn = val.serviceTargetCn;
}
EmergeFBReqInd& InitialContextSetupRequestCodec::getEmergeFBInd(){
    return emergeFBReqInd_m;
}
bool InitialContextSetupRequestCodec::isEmergeFallbackPrst(){
    return emergeFBReqIndPrst_m;
}
void InitialContextSetupRequestCodec::setRRCITRR(Uint32 val) {
    rrcReportPrst_m=true;
    rrcITRR_m = val;
}
Uint32 InitialContextSetupRequestCodec::getRRCITRR() { 
    return rrcITRR_m; 
}
bool InitialContextSetupRequestCodec::isRRCReportPrst(){
    return rrcReportPrst_m;
}

UeRadioCapForPaging& InitialContextSetupRequestCodec::getUeRadioCapForPaging()
{
    return ueRadioCapForPaging_m;
}

void InitialContextSetupRequestCodec::setUeRadioCapForPaging(UeRadioCapForPaging& val)
{
    ueRadioCapForPagingPrst_m=true;
    ueRadioCapForPaging_m = val;
}

bool InitialContextSetupRequestCodec::isUeRadioCapForPagingPrst()
{
    return ueRadioCapForPagingPrst_m;
}

//void InitialContextSetupRequestCodec::setCoreNetworkAssInfo(    CoreNetworkAssInfo& val) {
//    coreNetworkPrst_m=true;
//    coreNetworkAssInfo_m = val;
//}
//CoreNetworkAssInfo& InitialContextSetupRequestCodec::getCoreNetworkAssInfo() {
//    return coreNetworkAssInfo_m;
//}
//bool InitialContextSetupRequestCodec::isCoreNetworkPrst(){
//    return coreNetworkPrst_m;
//}
//void InitialContextSetupRequestCodec::setMobRestricList(MobRestricList& val) {
//    mobRestricList_m = val;
//}
//MobRestricList& InitialContextSetupRequestCodec::getMobRestricList(){
//    return mobRestricList_m;
//}

void InitialContextSetupRequestCodec::setRedirForVoiceEpsFallback(Uint8 val) {
    redirVoiceEpsFallbackPrst_m=true;
    redirVoiceEpsFallback_m = val;
}
Uint8 InitialContextSetupRequestCodec::getRedirForVoiceEpsFallback() { 
    return redirVoiceEpsFallback_m; 
}
bool InitialContextSetupRequestCodec::isRedirForVoiceEpsFallbackPrst(){
    return redirVoiceEpsFallbackPrst_m;
}
void InitialContextSetupRequestCodec::addAllowedSNssai(SNssai& val) {
    allowedSNssaiList_m.push_back(val);
}
AllowedNssaiList& InitialContextSetupRequestCodec::getAllowedSNssaiList() {
    return allowedSNssaiList_m;
}
void InitialContextSetupRequestCodec::addPduSessResReqList(    PduSessResSetupReqItem& val) {
    pduSessResReqList_m.push_back(val);
}
PduSessResSetupReqList& InitialContextSetupRequestCodec::getPduSessResReqList() {
    return pduSessResReqList_m;
}

bool InitialContextSetupRequestCodec::reset() 
{
    msgPdu_m.reset_v();

    amfUeNgapId_m = 0;
    ranUeNgapId_m = 0;
    nasPdu_m.clear();
    nasPduPrst_m=false;

    oldAmfName_m.clear();
    amfNamePrst_m=false;

    indexToRfsp_m = 0;

    memset(&ueAmbr_m, 0, sizeof(ueAmbr_m));
    ueAMBRPrst_m=false;

    allowedSNssaiList_m.clear();

    memset(&guami_m, 0, sizeof(guami_m));

    memset(&ueSecurityCap_m, 0, sizeof(ueSecurityCap_m));

    securityKey_m.clear();

    memset(&traceAct_m, 0, sizeof(traceAct_m));
    traceActPrst_m=false;

    ueRadioCap_m.clear();
    ueRadioCapPrst_m=false;

    maskImeiSv_m.clear();
    maskedImeiSvPrst_m=false;

    memset(&emergeFBReqInd_m, 0, sizeof(emergeFBReqInd_m));
    emergeFBReqIndPrst_m=false;

    rrcITRR_m = 0;
    rrcReportPrst_m=false;

    ueRadioCapForPaging_m.ueRadioCapForPagingNr.clear();
    ueRadioCapForPaging_m.ueRadioCapForPagingEutra.clear();
    ueRadioCapForPagingPrst_m=false;

//    memset(&coreNetworkAssInfo_m, 0, sizeof(coreNetworkAssInfo_m));
//    coreNetworkPrst_m=false;
//
//    memset(&mobRestricList_m, 0, sizeof(mobRestricList_m));
    redirVoiceEpsFallback_m = 0;
    redirVoiceEpsFallbackPrst_m=false;


    pduSessResReqList_m.clear();

    return true;
}

MsgBuffer InitialContextSetupRequestCodec::encode(NgapOssCtxt* ctxt) 
//bool InitialContextSetupRequestCodec::encode(NgapOssCtxt* ctxt, MsgBuffer msgBuffer)
{
    MsgBuffer msgBuffer;

    msgBuffer.length = 0;
    msgBuffer.value = NULL;

    // encode RAN UE NGAP ID
    RAN_UE_NGAP_ID* ranUeIdPtr = msgPdu_m.getRanUeNgapIdEncodePtr();
    *ranUeIdPtr = ranUeNgapId_m;

    // encode AMF UE NGAP ID
    AMF_UE_NGAP_ID* amfUeIdPtr = msgPdu_m.getAmfUeNgapIdEncodePtr();
    *amfUeIdPtr = amfUeNgapId_m;

    // mandatory IEs
    // encode Allowed Nssai 
    unsigned int numofAllowedItems = allowedSNssaiList_m.size();
    msgPdu_m.setNumofAllowedNssai_Item_m_v(numofAllowedItems);
    for (int i = 0; i < numofAllowedItems; i++) 
    {
        AllowedNSSAI_Item* itemPtr = msgPdu_m.getAllowedNssai_ItemEncodePtr(i);
        if (itemPtr != NULL) 
        {
            itemPtr->s_NSSAI.sST.length = sstSize_c;
            memcpy(itemPtr->s_NSSAI.sST.value, allowedSNssaiList_m[i].sst, sstSize_c);

            if (allowedSNssaiList_m[i].sdPresent) 
            {
                itemPtr->s_NSSAI.bit_mask = sD_present;
                itemPtr->s_NSSAI.sD.length = sdSize_c;
                memcpy(itemPtr->s_NSSAI.sD.value, allowedSNssaiList_m[i].sd, sdSize_c);
            }
        }
    }

    // encode GUAMI
    GUAMI* gumaiPtr = msgPdu_m.getGuamiEncodePtr();
    gumaiPtr->pLMNIdentity.length = sizeof(guami_m.plmnId);
    memcpy(gumaiPtr->pLMNIdentity.value, guami_m.plmnId, gumaiPtr->pLMNIdentity.length);

    gumaiPtr->aMFRegionID.length = amfReginIdBitLen;
    memcpy(gumaiPtr->aMFRegionID.value, &(guami_m.amfId.regionId), sizeof(guami_m.amfId.regionId));

    gumaiPtr->aMFSetID.length = amfSetIdBitLen;
    guami_m.amfId.setId[1] = guami_m.amfId.setId[1] << 6;
    memcpy(gumaiPtr->aMFSetID.value, guami_m.amfId.setId, sizeof(guami_m.amfId.setId));

    gumaiPtr->aMFPointer.length = amfPointerBitLen;
    guami_m.amfId.pointer = guami_m.amfId.pointer << 2;
    memcpy(gumaiPtr->aMFPointer.value, &(guami_m.amfId.pointer), sizeof(guami_m.amfId.pointer));

    // encode UE Securtity Capability
    UESecurityCapabilities* ueSecCapPtr = msgPdu_m.getUeSecurityCapabilitiesEncodePtr();
    ueSecCapPtr->nRencryptionAlgorithms.length = nRencryptionAlgBitLen;
    memcpy(ueSecCapPtr->nRencryptionAlgorithms.value,
        ueSecurityCap_m.nrEncryptionAlg,
        sizeof(ueSecurityCap_m.nrEncryptionAlg));
    ueSecCapPtr->nRintegrityProtectionAlgorithms.length = nRintegrityProtectionAlgBitLen;
    memcpy(ueSecCapPtr->nRintegrityProtectionAlgorithms.value,
        ueSecurityCap_m.nrIntegrityProAlg,
        sizeof(ueSecurityCap_m.nrIntegrityProAlg));
    ueSecCapPtr->eUTRAencryptionAlgorithms.length = eUTRAencryptionAlgBitLen;
    memcpy(ueSecCapPtr->eUTRAencryptionAlgorithms.value,
        ueSecurityCap_m.eutraEncryptionAlg,
        sizeof(ueSecurityCap_m.eutraEncryptionAlg));
    ueSecCapPtr->eUTRAintegrityProtectionAlgorithms.length = eUTRAintegrityProtectionAlgBitLen;
    memcpy(ueSecCapPtr->eUTRAintegrityProtectionAlgorithms.value,
        ueSecurityCap_m.eutraIntegrityProAlg,
        sizeof(ueSecurityCap_m.eutraIntegrityProAlg));

    // optional IEs
    // encode NAS PDU
    if(nasPduPrst_m){
        NAS_PDU* nasPduPtr = msgPdu_m.getNasPduEncodePtr();
        nasPduPtr->length = nasPdu_m.length();
        memcpy(nasPduPtr->value, nasPdu_m.c_str(), nasPduPtr->length);
    }

    // Encode Amf Name
    if(amfNamePrst_m){
        AMFName* oldAmfNamePtr = msgPdu_m.getAmfNameEncodePtr();
        oldAmfNamePtr->length = oldAmfName_m.length();
        memcpy(oldAmfNamePtr->value, oldAmfName_m.c_str(), oldAmfNamePtr->length);
    }

    // encode index to rfsp
    if(indexToRFSPPrst_m){
        IndexToRFSP* index = msgPdu_m.getIndexToRFSPEncodePtr();
        *index = indexToRfsp_m;
    }

    // encode Ue ambr
    if(ueAMBRPrst_m){
        UEAggregateMaximumBitRate* ueAmbrPtr =
            msgPdu_m.getUeAggregateMaximumBitRateEncodePtr();
        ueAmbrPtr->uEAggregateMaximumBitRateDL = ueAmbr_m.downlink;
        ueAmbrPtr->uEAggregateMaximumBitRateUL = ueAmbr_m.uplink;
    }

    // encode Security Key
    SecurityKey* securityKeyPtr = msgPdu_m.getSecurityKeyEncodePtr();
    securityKeyPtr->length = securityKeyBitLen;
    memcpy(securityKeyPtr->value, securityKey_m.c_str(), securityKey_m.length());

    // encode TraceActivation
    if(traceActPrst_m){
        TraceActivation* tracePtr = msgPdu_m.getTraceActivationEncodePtr();
        tracePtr->nGRANTraceID.length = sizeof(traceAct_m.traceID);
        memcpy(tracePtr->nGRANTraceID.value, &traceAct_m.traceID, sizeof(traceAct_m.traceID));

        tracePtr->interfacesToTrace.length = interfacesToTraceBitLen;
        memcpy(tracePtr->interfacesToTrace.value, &(traceAct_m.interfaces), sizeof(traceAct_m.interfaces));
        memcpy(&(tracePtr->traceDepth), &(traceAct_m.traceDepth), sizeof(traceAct_m.traceDepth));

        tracePtr->traceCollectionEntityIPAddress.length = traceColEntityIPAddrBitLen;
        memcpy(tracePtr->traceCollectionEntityIPAddress.value,
            traceAct_m.traceEntityIPAddr, sizeof(traceAct_m.traceEntityIPAddr));
    }

    // encode UE Radio Capability
    if(ueRadioCapPrst_m) {
        UERadioCapability* ueRadioCapPtr = msgPdu_m.getUERadioCapabilityEncodePtr();
        ueRadioCapPtr->length = ueRadioCap_m.length();
        memcpy(ueRadioCapPtr->value, ueRadioCap_m.c_str(), ueRadioCapPtr->length);
    }

    // encode Masked IMEISVcd
    if(maskedImeiSvPrst_m) {
        MaskedIMEISV* maskedIMEISVPtr = msgPdu_m.getMaskedIMEISVEncodePtr();
        maskedIMEISVPtr->length = maskedIMEISVBitLen;
        memcpy(maskedIMEISVPtr->value, maskImeiSv_m.c_str(), maskImeiSv_m.length());
    }

    // encode Emergency Fallback Indicator
    if(emergeFBReqIndPrst_m) {
        EmergencyFallbackIndicator* emergencyFallbackIndicatorPtr =
            msgPdu_m.getEmergencyFallbackIndicatorEncodePtr();
        emergencyFallbackIndicatorPtr->emergencyFallbackRequestIndicator =
            (EmergencyFallbackRequestIndicator)(emergeFBReqInd_m.indicator);
        emergencyFallbackIndicatorPtr->emergencyServiceTargetCN =
            (EmergencyServiceTargetCN)(emergeFBReqInd_m.serviceTargetCn);
    }

    // encode RRC Inactive Transition Report Request
    if(rrcReportPrst_m) {
        RRCInactiveTransitionReportRequest* ptr 
            = msgPdu_m.getRRCInactiveTransitionReportRequestEncodePtr();
        *ptr = (RRCInactiveTransitionReportRequest)rrcITRR_m;
    }

    // encode UE Radio Capability For Paging
    if(ueRadioCapForPagingPrst_m) {
        UERadioCapabilityForPaging* ueRadioCapForPagingPtr = 

        msgPdu_m.getUeRadioCapabilityForPagingEncodePtr();
        if(ueRadioCapForPaging_m.ueRadioCapForPagingNrPrst) {
            ueRadioCapForPagingPtr->bit_mask |= uERadioCapabilityForPagingOfNR_present;
            ueRadioCapForPagingPtr->uERadioCapabilityForPagingOfNR.length 
                = ueRadioCapForPaging_m.ueRadioCapForPagingNr.length();
            memcpy(ueRadioCapForPagingPtr->uERadioCapabilityForPagingOfNR.value, 
                   ueRadioCapForPaging_m.ueRadioCapForPagingNr.c_str(), 
                   ueRadioCapForPagingPtr->uERadioCapabilityForPagingOfNR.length);
        }
        if(ueRadioCapForPaging_m.ueRadioCapForPagingEutraPrst) {
            ueRadioCapForPagingPtr->bit_mask |= uERadioCapabilityForPagingOfEUTRA_present;
            ueRadioCapForPagingPtr->uERadioCapabilityForPagingOfEUTRA.length 
                = ueRadioCapForPaging_m.ueRadioCapForPagingEutra.length();
            memcpy(ueRadioCapForPagingPtr->uERadioCapabilityForPagingOfEUTRA.value, 
                   ueRadioCapForPaging_m.ueRadioCapForPagingEutra.c_str(), 
                   ueRadioCapForPagingPtr->uERadioCapabilityForPagingOfEUTRA.length);
        }
    }

    // encode 
    if(redirVoiceEpsFallbackPrst_m) {
        RedirectionVoiceFallback* ptr 
            = msgPdu_m.getRedirectionVoiceFallbackEncodePtr();
        *ptr = (RedirectionVoiceFallback)redirVoiceEpsFallback_m;
    }

    //encode PDUSessionResourceSetupListCxtReq 
    Uint32 num = pduSessResReqList_m.size();
    msgPdu_m.setNumofPDUSessionResourceSetupListItemCxtReq_v(num);

    for (int i = 0; i < num; i++) {
        PDUSessionResourceSetupItemCxtReq *pduPtr = 
            msgPdu_m.getPDUSessionResourceSetupItemCxtReqEncodePtr(i);
        if (NULL == pduPtr) {
        std::cout <<"invalid pduPtr"<< std::endl;
            return msgBuffer;
        }

        PduSessResSetupReqItem *itemPtr = &(pduSessResReqList_m[i]);
        
        pduPtr->pDUSessionID = itemPtr->pduSessId;

        if (itemPtr->nasPduPrst) {
            pduPtr->bit_mask = PDUSessionResourceSetupItemCxtReq_nAS_PDU_present;
            pduPtr->nAS_PDU.length =itemPtr->nasPdu.length();
            memcpy(pduPtr->nAS_PDU.value,itemPtr->nasPdu.c_str(), itemPtr->nasPdu.length());
        }

        pduPtr->s_NSSAI.sST.length =sizeof(itemPtr->snssai.sst);
        memcpy(pduPtr->s_NSSAI.sST.value, itemPtr->snssai.sst, pduPtr->s_NSSAI.sST.length);
        if (itemPtr->snssai.sdPresent) {
            pduPtr->s_NSSAI.bit_mask = sD_present;
            pduPtr->s_NSSAI.sD.length = sizeof(itemPtr->snssai.sd);
            memcpy(pduPtr->s_NSSAI.sD.value, itemPtr->snssai.sd, pduPtr->s_NSSAI.sD.length);
        }

        //pdu session resource setup transfer
        pduPtr->pDUSessionResourceSetupRequestTransfer.encoded.length 
          = itemPtr->pduSessResSetupReqTransfer.length();
        memcpy(pduPtr->pDUSessionResourceSetupRequestTransfer.encoded.value,
               itemPtr->pduSessResSetupReqTransfer.c_str(),
               itemPtr->pduSessResSetupReqTransfer.length());   
    }

    msgPdu_m.encodeOssToPer(ctxt);

    OssBuf* perBuf = msgPdu_m.getPerBufRef();

    msgBuffer.length = perBuf->length;
    msgBuffer.value = perBuf->value;

    return msgBuffer;
}

bool InitialContextSetupRequestCodec::decode(NgapOssCtxt* ctxt,MsgBuffer msgBuf) 
{
    reset();

    OssBuf ossbuf;
    ossbuf.length = msgBuf.length;
    ossbuf.value = msgBuf.value;

    msgPdu_m.setPerBufRef_v(ossbuf);

    if (0 != msgPdu_m.decodePerToOss(ctxt)) {
        cout << "Failed to decode Per to OSS" << endl;
        return false;
    }
    // decode RAN UE NGAP ID
    ranUeNgapId_m = *(msgPdu_m.getRanUeNgapIdDecodedPtr());
    // decode amf ue ngap id
    amfUeNgapId_m = *(msgPdu_m.getAmfUeNgapIdDecodedPtr());

    // decode Allowed NSSAI
    for (int i = 0; i < msgPdu_m.getNumofAllowedNssai_Item(); i++) {
        AllowedNSSAI_Item* ptr = msgPdu_m.getAllowedNssai_ItemDecodedPtr(i);

        SNssai snssai;
        memcpy(snssai.sst, ptr->s_NSSAI.sST.value, sizeof(ptr->s_NSSAI.sST.value));

        if (ptr->s_NSSAI.bit_mask == sD_present) {
            memcpy(snssai.sd, ptr->s_NSSAI.sD.value, sizeof(ptr->s_NSSAI.sD.value));
            snssai.sdPresent = true;
        }

        allowedSNssaiList_m.push_back(snssai);
    }

    // decode GUAMI
    memcpy(guami_m.plmnId, 
           msgPdu_m.getGuamiDecodedPtr()->pLMNIdentity.value,
           msgPdu_m.getGuamiDecodedPtr()->pLMNIdentity.length);
    memcpy(guami_m.amfId.setId, 
           msgPdu_m.getGuamiDecodedPtr()->aMFSetID.value,
           sizeof(msgPdu_m.getGuamiDecodedPtr()->aMFSetID.value));
    guami_m.amfId.pointer = *(msgPdu_m.getGuamiDecodedPtr()->aMFPointer.value);
    guami_m.amfId.regionId = *(msgPdu_m.getGuamiDecodedPtr()->aMFRegionID.value);

    // decode UE Securtity Capability
    memcpy(ueSecurityCap_m.nrEncryptionAlg,
           msgPdu_m.getUeSecurityCapabilitiesDecodedPtr()->nRencryptionAlgorithms.value,
           nRencryptionAlgBitLen / bitLen);
    memcpy(ueSecurityCap_m.nrIntegrityProAlg,
           msgPdu_m.getUeSecurityCapabilitiesDecodedPtr()->nRintegrityProtectionAlgorithms.value,
           nRintegrityProtectionAlgBitLen / bitLen);
    memcpy(ueSecurityCap_m.eutraEncryptionAlg,
           msgPdu_m.getUeSecurityCapabilitiesDecodedPtr()->eUTRAencryptionAlgorithms.value,
           eUTRAencryptionAlgBitLen / bitLen);
    memcpy(ueSecurityCap_m.eutraIntegrityProAlg,
           msgPdu_m.getUeSecurityCapabilitiesDecodedPtr()->eUTRAintegrityProtectionAlgorithms.value,
           eUTRAintegrityProtectionAlgBitLen / bitLen);

    // decode Security
    SecurityKey* securityKeyPtr = msgPdu_m.getSecurityKeyDecodedPtr();
    securityKey_m.assign((char*)(securityKeyPtr->value),securityKeyBitLen / bitLen);
    
    //decode PDUSessionResourceSetupListCxtReq  new
    for (int i = 0; i < msgPdu_m.getNumofPDUSessionResourceSetupListItemCxtReq();i++) {
        PDUSessionResourceSetupItemCxtReq *pduPtr =
            msgPdu_m.getPDUSessionResourceSetupItemCxtReqDecodedPtr(i);
        if (NULL == pduPtr) {
            return false;
        }
          
        PduSessResSetupReqItem item;
        item.pduSessId = pduPtr->pDUSessionID;
    
        if (pduPtr->bit_mask ==PDUSessionResourceSetupItemCxtReq_nAS_PDU_present) {
            item.nasPduPrst = true;
            item.nasPdu.assign((char*)(pduPtr->nAS_PDU.value),  pduPtr->nAS_PDU.length);
        }
    
        memcpy(item.snssai.sst, pduPtr->s_NSSAI.sST.value,sizeof(pduPtr->s_NSSAI.sST.value));
        if (pduPtr->s_NSSAI.bit_mask == sD_present) {
            item.snssai.sdPresent = true;
            memcpy(item.snssai.sd, pduPtr->s_NSSAI.sD.value, sizeof(pduPtr->s_NSSAI.sD.value));
        }
    
        //Pdu Session Resource Setup Request Transfer
        OssBuf outputBuf = ctxt->getSecondaryCodecBuf();
        PDUSessionResourceSetupRequestTransfer* transferPdu;
        transferPdu = pduPtr->pDUSessionResourceSetupRequestTransfer.decoded;
        unsigned int result = 
            ctxt->encodeOssToPer(PDUSessionResourceSetupRequestTransfer_PDU, transferPdu, &outputBuf);
        if (result != 0){
            std::cout << "failed to encode transfer"<< std::endl;
            return false;
        }
        item.pduSessResSetupReqTransfer.assign(
            (char *)(outputBuf.value),
            outputBuf.length);    
    
        pduSessResReqList_m.push_back(item);
    }

    // decode NAS PDU
    if(msgPdu_m.isNasPduPresent()) {
        NAS_PDU* nasPduPtr = msgPdu_m.getNasPduDecodedPtr();
        nasPdu_m.assign((char*)(nasPduPtr->value), nasPduPtr->length);
        nasPduPrst_m=true;
    }

    // decode old amf name
    if(msgPdu_m.isAmfNamePresent()) {
        AMFName* amfNamePtr = msgPdu_m.getAmfNameDecodedPtr();
        oldAmfName_m.assign((char*)(amfNamePtr->value), amfNamePtr->length);
        amfNamePrst_m=true;
    }

    // decode index to rfsp
    if(msgPdu_m.isIndexToRFSPPresent()){
        IndexToRFSP* indexPtr = msgPdu_m.getIndexToRFSPDecodedPtr();
        indexToRfsp_m = *indexPtr;
        indexToRFSPPrst_m=true;
    }

    // decode Ue Ambr
    if(msgPdu_m.isUeAMBRPresent()) {
        ueAmbr_m.uplink = 
            msgPdu_m.getUeAggregateMaximumBitRateDecodedPtr()->uEAggregateMaximumBitRateUL;
        ueAmbr_m.downlink = 
            msgPdu_m.getUeAggregateMaximumBitRateDecodedPtr()->uEAggregateMaximumBitRateDL;
        ueAMBRPrst_m=true;
    }

    // decode TraceActivation
    if(msgPdu_m.isTraceActPresent()) {
        memcpy(&(traceAct_m.traceID),
            msgPdu_m.getTraceActivationDecodedPtr()->nGRANTraceID.value,
            msgPdu_m.getTraceActivationDecodedPtr()->nGRANTraceID.length);
        traceAct_m.interfaces =
            *(msgPdu_m.getTraceActivationDecodedPtr()->interfacesToTrace.value);
        traceAct_m.traceDepth =
            (Uint8)(msgPdu_m.getTraceActivationDecodedPtr()->traceDepth);
        memcpy(traceAct_m.traceEntityIPAddr,
            msgPdu_m.getTraceActivationDecodedPtr()->traceCollectionEntityIPAddress.value,
            traceColEntityIPAddrBitLen / bitLen);
        traceActPrst_m=true;
    }

    // decode UE Radio Capability
    if(msgPdu_m.isUeRadioPresent()) {
        UERadioCapability* ueRadioCapPtr = msgPdu_m.getUERadioCapabilityDecodedPtr();
        ueRadioCap_m.assign((char*)(ueRadioCapPtr->value), ueRadioCapPtr->length);
        ueRadioCapPrst_m=true;
    }

    // decode Masked IMEISV
    if(msgPdu_m.isMaskedIMEISVPresent()){
        MaskedIMEISV* ptr = msgPdu_m.getMaskedIMEISVDecodedPtr();

        maskImeiSv_m.assign((char*)(ptr->value),  sizeof((char*)(ptr->value)));
        maskedImeiSvPrst_m=true;
    }

    // decode Emergency Fallback Indicator
    if(msgPdu_m.isEmerFallbackPresent())
    {
        emergeFBReqInd_m.indicator 
            = msgPdu_m.getEmergencyFallbackIndicatorDecodedPtr()->emergencyFallbackRequestIndicator;
        emergeFBReqInd_m.serviceTargetCn
            = msgPdu_m.getEmergencyFallbackIndicatorDecodedPtr()->emergencyFallbackRequestIndicator;
        emergeFBReqIndPrst_m=true;
    }

    // decode RRC Inactive Transition Report Request
    if(msgPdu_m.isRRCReportPresent()){
        rrcITRR_m = *(msgPdu_m.getRRCInactiveTransitionReportRequestDecodedPtr());
        rrcReportPrst_m=true;
    }

    // decode Redirection for Voice Eps FallBack
    if(msgPdu_m.isRedirectionVoiceFallbackPresent()){
        redirVoiceEpsFallback_m = *(msgPdu_m.getRedirectionVoiceFallbackDecodedPtr());
        redirVoiceEpsFallbackPrst_m=true;
    }

    // decode UE Radio Capability For Paging
    if(msgPdu_m.isUeRadioCapabilityForPagingPresent()) {
        if(msgPdu_m.getUeRadioCapabilityForPagingDecodedPtr()->bit_mask & uERadioCapabilityForPagingOfNR_present)
        {
            UERadioCapabilityForPaging* ueRadioCapForPagingPtr = 
                msgPdu_m.getUeRadioCapabilityForPagingDecodedPtr();
            ueRadioCapForPaging_m.ueRadioCapForPagingNrPrst = true;
            ueRadioCapForPaging_m.ueRadioCapForPagingNr
                .assign((char*)(ueRadioCapForPagingPtr->uERadioCapabilityForPagingOfNR.value), 
                                ueRadioCapForPagingPtr->uERadioCapabilityForPagingOfNR.length);
        }
        if(msgPdu_m.getUeRadioCapabilityForPagingDecodedPtr()->bit_mask & uERadioCapabilityForPagingOfEUTRA_present)
        {
            UERadioCapabilityForPaging* ueRadioCapForPagingPtr = 
                msgPdu_m.getUeRadioCapabilityForPagingDecodedPtr();
            ueRadioCapForPaging_m.ueRadioCapForPagingEutraPrst = true;
            ueRadioCapForPaging_m.ueRadioCapForPagingEutra
                .assign((char*)(ueRadioCapForPagingPtr->uERadioCapabilityForPagingOfEUTRA.value), 
                                ueRadioCapForPagingPtr->uERadioCapabilityForPagingOfEUTRA.length);
        }

        ueRadioCapForPagingPrst_m=true;
    }

  return true;
}

string InitialContextSetupRequestCodec::toString(){
   ostringstream oss;

   oss << "InitialContextSetupRequestCodec:" << endl
       << "1.AMF UE NGAP ID:(" << amfUeNgapId_m << ")" << endl
       << "2.RAN UE NGAP ID:(" << ranUeNgapId_m << ")" << endl
       << "3.Old Amf Name:(" << oldAmfName_m << ")" << endl
       << "4.UeAmbr:( downlink - " << ueAmbr_m.downlink << ", uplink - "
       << ueAmbr_m.uplink << ")" << endl
       << "5.GUAMI:(AMFRegionId:(" << guami_m.amfId.regionId << endl
       << "AmfPointer:" << int(guami_m.amfId.pointer) <<endl
       << "),AMFSetID:(" << guami_m.amfId.setId[0] << guami_m.amfId.setId[1]
       << "),PLMNIdentity:(" << guami_m.plmnId[0] << guami_m.plmnId[1]
       << guami_m.plmnId[2] << "))" << endl
       << "RedirectionForVoiceEpsFallback:Prst("<<redirVoiceEpsFallbackPrst_m << ") value (" << redirVoiceEpsFallback_m << ")" << endl
       << "6.UeRadioCapabilityForPaging:Prst" <<ueRadioCapForPagingPrst_m<< endl
       << "Nr:Prst" << ueRadioCapForPaging_m.ueRadioCapForPagingNrPrst << endl
       << " value " << ueRadioCapForPaging_m.ueRadioCapForPagingNr <<endl
       << "Eutra:Prst" << ueRadioCapForPaging_m.ueRadioCapForPagingEutraPrst << endl
       << " value " << ueRadioCapForPaging_m.ueRadioCapForPagingEutra <<endl;
//  vector<SNssai>::iterator it1;
//  for (it1 = allowedSNssaiList_m.begin(); it1 < allowedSNssaiList_m.end(); it1++) {
//    oss << "sD:(" << it1->sd[0] << it1->sd[1] << it1->sd[2] << ")" << endl
//        << "sST:(" << it1->sst[0] << ")" << endl;
//  }
//   oss << "7.UESecurtityCapability:" << endl
//       << "NRencryptionAlgorithms:(" << ueSecurityCap_m.nRencryptionAlg[0]
//       << ueSecurityCap_m.nRencryptionAlg[1]
//       << "),NRintegrityProtectionAlgorithms:("
//       << ueSecurityCap_m.nRintegrityProAlg[0]
//       << ueSecurityCap_m.nRintegrityProAlg[1]
//       << "),EUTRAencryptionAlgorithms:("
//       << ueSecurityCap_m.eUTRAencryptionAlg[0]
//       << ueSecurityCap_m.eUTRAencryptionAlg[1]
//       << "),EUTRAintegrityProtectionAlgorithms:("
//       << ueSecurityCap_m.eUTRAintegrityProAlg[0]
//       << ueSecurityCap_m.eUTRAintegrityProAlg[1] << ")" << endl
//       << "8.SecurityKey value:(" << securityKey_m << ")" << endl
//       << "9.TraceActivation:" << endl
//       << "NGRANTraceID:(" << traceAct_m.nGRANTraceID[0]
//       << traceAct_m.nGRANTraceID[1] << traceAct_m.nGRANTraceID[2]
//       << traceAct_m.nGRANTraceID[3] << traceAct_m.nGRANTraceID[4]
//       << traceAct_m.nGRANTraceID[5] << traceAct_m.nGRANTraceID[6]
//       << traceAct_m.nGRANTraceID[7] << "),InterfacesToTrace:("
//       << traceAct_m.interfacesToTrace << "),TraceDepth:("
//       << traceAct_m.traceDepth << "),TransportLayerAddress:("
//       << traceAct_m.traceCollectionEntityIPAddress << ")" << endl
//       << "10.UERadioCapability:(" << ueRadioCap_m << ")" << endl
//       << "11.IndexToRfsp:(" << indexToRfsp_m << ")" << endl
//       << "12.MaskedIMEISV:(" << maskIMEISV_m << ")" << endl
//       << "13.NAS PDU:(" << nasPdu_m << ")" << endl
//       << "14.EmergencyFallbackIndicator:(EmergencyFallbackRequestIndicator:("
//       << emerFallbackIndicator_m.emerFallbackRequestIndicator
//       << "),EmergencyServiceTargetCN:("
//       << emerFallbackIndicator_m.emerServiceTargetCN << "))" << endl
//       << "15.RRCInactiveTransitionReportRequest:(" << rrcITRR_m << ")" << endl;
//   oss << "PDU Session Resource Setup Request List:" << endl;
//
//
////CoreNetworkInformation new
//   oss << "16.CoreNetworkInformation:"<<endl
//       << "UE Identity IndexValue:("
//       << coreNetworkAssInfo_m.ueIdentityIndexValue[0] << coreNetworkAssInfo_m.ueIdentityIndexValue[1] 
//       << ")"<<endl
//       << "Paging DRX:Flag(" << coreNetworkAssInfo_m.pagingDRXPrst << ")" <<endl
//       << "(" << coreNetworkAssInfo_m.pagingDRX
//       << ")"<<endl
//       << "PeriodicRegistrationUpdateTimer:("<<coreNetworkAssInfo_m.PerRegUpdateTimer<<")"<<endl
//       << "MICOModeIndication:Flag("<<coreNetworkAssInfo_m.mICOModeIndPrst << ")" << endl
//       << "(" <<coreNetworkAssInfo_m.mICOModeIndication
//       << ")"<<endl
//       << "TAIListForInactive:"<<endl;
// vector<TAILFIItem>::iterator it2;
// for(it2=coreNetworkAssInfo_m.taiLFIList.begin();it2<coreNetworkAssInfo_m.taiLFIList.end();it2++)
// {
//   oss <<"Tac:("<<it2->tac[0]<<it2->tac[1]<<it2->tac[2]<<")"<<endl
//       <<"PLMNIdentity:("<<it2->plmnId[0] <<it2->plmnId[1] <<it2->plmnId[2] <<")" <<endl;
// }
//
// oss <<"ExpectedUEBehaviour:Flag(" 
//     << coreNetworkAssInfo_m.expectUEBehaPrst 
//     << ")"<<endl
//
//     <<"ExpectedUEActivityBehaviour:"<<endl
//
//     <<"ExpectedActivityPeriod:Flag(" 
//     << coreNetworkAssInfo_m.expectedUEBeha.expectedUEActBeha.expActPeriodPrst 
//     << ")" <<endl
//
//     <<"(" <<coreNetworkAssInfo_m.expectedUEBeha.expectedUEActBeha.expectedActivityPeriod
//     <<")"<<endl
//
//     <<"ExpectedIdlePeriod:Flag(" 
//     << coreNetworkAssInfo_m.expectedUEBeha.expectedUEActBeha.expIdlePeriodPrst 
//     << ")" <<endl
//
//     <<"(" <<coreNetworkAssInfo_m.expectedUEBeha.expectedUEActBeha.expectedIdlePeriod
//     <<")"<<endl
//
//     <<"SourceOfUEActivityBehaviourInformation:Flag("
//     << coreNetworkAssInfo_m.expectedUEBeha.expectedUEActBeha.sourceOfUEActBehaInfo 
//     << ")" <<endl
//
//     <<"("<<coreNetworkAssInfo_m.expectedUEBeha.expectedUEActBeha.sourceOfUEActBehaInfo
//     <<")"<<endl
//
//     <<"ExpectedHOInterval:Flag("<< coreNetworkAssInfo_m.expectedUEBeha.expHOIntervalPrst 
//     << ")" <<endl
//
//     <<"("<<coreNetworkAssInfo_m.expectedUEBeha.expHOInterval
//     <<")"<<endl
//
//     <<"ExpectedUEMobility:Flag("<< coreNetworkAssInfo_m.expectedUEBeha.expUEMobilityPrst 
//     << ")" <<endl
//
//     <<"(" <<coreNetworkAssInfo_m.expectedUEBeha.expUEMobility 
//     <<")"<<endl;
//
// vector<ExpectedUEMovingTra>::iterator it3;
// for(it3=coreNetworkAssInfo_m.expectedUEBeha.expectedUEMovingTraLists.begin();
//     it3<coreNetworkAssInfo_m.expectedUEBeha.expectedUEMovingTraLists.end();
//     it3++)
// {
//    oss <<"ExpectedUEMovingTra:"<<endl
//        <<"timeStayedInCell:("<< it3->timeStayedPrst << ")" <<endl
//        <<"("<<it3->timeStayedInCell<<")"<<endl
//        <<"PLMNIdentity:(" 
//        << it3->nrcgi_a.plmn[0] 
//        << it3->nrcgi_a.plmn[1] 
//        << it3->nrcgi_a.plmn[2] 
//        << ")"<<endl
//        <<"NRCellIdentity:(" 
//        <<it3->nrcgi_a.nrCellId[0] 
//        <<it3->nrcgi_a.nrCellId[1] 
//        <<it3->nrcgi_a.nrCellId[2] 
//        <<it3->nrcgi_a.nrCellId[3] 
//        <<it3->nrcgi_a.nrCellId[4] 
//        <<")"<<endl;
// }
   return oss.str(); 
}


/////////////////////////////////////////////////////
//encode CoreNetworkAssistanceInformation
/////////////////////////////////////////////////////
//    if(coreNetworkPrst_m) {
//        CoreNetworkAssistanceInformation* cnAssPtr = 
//            msgPdu_m.getCoreNetworkAssistanceInformationEncodePtr();
//
//        cnAssPtr->uEIdentityIndexValue.u.indexLength10.length=uEIdentityIndexValueBitLen;
//        coreNetworkAssInfo_m.ueIdentityIndValue[1] = << 6;
//
//        memcpy(cnAssPtr->uEIdentityIndexValue.u.indexLength10.value,
//            coreNetworkAssInfo_m.ueIdentityIndValue,
//            sizeof(coreNetworkAssInfo_m.ueIdentityIndValue));
//
//        if(coreNetworkAssInfo_m.pagingDrxPrst){
//            cnAssPtr->bit_mask=uESpecificDRX_present;
//            cnAssPtr->uESpecificDRX=(PagingDRX)(coreNetworkAssInfo_m.pagingDrx);
//        }
//
//        cnAssPtr->periodicRegistrationUpdateTimer.length=perRegUpdateTimerBitLen;
//        memcpy(cnAssPtr->periodicRegistrationUpdateTimer.value,
//            &(coreNetworkAssInfo_m.PerRegUpdateTimer),
//                sizeof(coreNetworkAssInfo_m.PerRegUpdateTimer));
//
//        if(coreNetworkAssInfo_m.micoModeIndPrst){
//            cnAssPtr->bit_mask=mICOModeIndication_present;
//            cnAssPtr->mICOModeIndication=
//                (MICOModeIndication)(coreNetworkAssInfo_m.micoModeInd);
//        }
        
//        Uint32 num =coreNetworkAssInfo_m.taiLFIList.size();
//        msgPdu_m.setNumOfTAIListForInactiveItem_v(num);
//        for(int i =0;i<num;i++)
//        {
//            TAIListForInactiveItem* taiListItemPtr = 
//                msgPdu_m.getTAIListForInactiveItemEncodePtr(i); 
//            if (NULL ==taiListItemPtr)
//            {
//                std::cout << "NULL Pointer for taiListItemPtr" << std::endl;
//                return msgBuffer;
//            }
//            taiListItemPtr->tAI.pLMNIdentity.length = 
//                sizeof(coreNetworkAssInfo_m.taiLFIList[i].plmnId);
//
//            memcpy(taiListItemPtr->tAI.pLMNIdentity.value,
//                coreNetworkAssInfo_m.taiLFIList[i].plmnId,
//                    taiListItemPtr->tAI.pLMNIdentity.length);
//
//            taiListItemPtr->tAI.tAC.length = 
//                sizeof(coreNetworkAssInfo_m.taiLFIList[i].tac);
//
//            memcpy(taiListItemPtr->tAI.tAC.value,
//                coreNetworkAssInfo_m.taiLFIList[i].tac,
//                    taiListItemPtr->tAI.tAC.length);
//        }
//
//        if(coreNetworkAssInfo_m.expectUEBehaPrst)
//        {
//            cnAssPtr->bit_mask=expectedUEBehaviour_present;
//
//            if(coreNetworkAssInfo_m.expectedUEBeha.expectUEActBehaPrst)
//            {
//                cnAssPtr->expectedUEBehaviour.bit_mask=expectedUEActivityBehaviour_present;
//                if(coreNetworkAssInfo_m.expectedUEBeha.expectedUEActBeha.expActPeriodPrst)
//                {
//                cnAssPtr->expectedUEBehaviour.expectedUEActivityBehaviour.bit_mask=
//                    expectedActivityPeriod_present;
//                cnAssPtr->expectedUEBehaviour.expectedUEActivityBehaviour.expectedActivityPeriod=
//                    coreNetworkAssInfo_m.expectedUEBeha.expectedUEActBeha.expectedActivityPeriod;
//                }
//
//                if(coreNetworkAssInfo_m.expectedUEBeha.expectedUEActBeha.expIdlePeriodPrst)
//                {
//                    cnAssPtr->expectedUEBehaviour.expectedUEActivityBehaviour.bit_mask=
//                        expectedIdlePeriod_present;
//                    cnAssPtr->expectedUEBehaviour.expectedUEActivityBehaviour.expectedIdlePeriod=
//                        coreNetworkAssInfo_m.expectedUEBeha.expectedUEActBeha.expectedIdlePeriod;
//                }
//
//                if(coreNetworkAssInfo_m.expectedUEBeha.expectedUEActBeha.sourceActBehaPrst)
//                {
//                    cnAssPtr->expectedUEBehaviour.expectedUEActivityBehaviour.bit_mask=
//                        sourceOfUEActivityBehaviourInformation_present;
//                    cnAssPtr->expectedUEBehaviour.expectedUEActivityBehaviour
//                        .sourceOfUEActivityBehaviourInformation=(SourceOfUEActivityBehaviourInformation)
//                            (coreNetworkAssInfo_m.expectedUEBeha.expectedUEActBeha.sourceOfUEActBehaInfo);
//                }
//
//            }
//
//            if(coreNetworkAssInfo_m.expectedUEBeha.expHOIntervalPrst)
//            {
//                cnAssPtr->expectedUEBehaviour.bit_mask=expectedHOInterval_present;
//                cnAssPtr->expectedUEBehaviour.expectedHOInterval=
//                    (ExpectedHOInterval)(coreNetworkAssInfo_m.expectedUEBeha.expHOInterval);
//            }
//
//            if(coreNetworkAssInfo_m.expectedUEBeha.expUEMobilityPrst)
//            {
//                cnAssPtr->expectedUEBehaviour.bit_mask=expectedUEMobility_present;
//                cnAssPtr->expectedUEBehaviour.expectedUEMobility=
//                    (ExpectedUEMobility)(coreNetworkAssInfo_m.expectedUEBeha.expUEMobility);
//            }
//
//            if(coreNetworkAssInfo_m.expectedUEBeha.expUEMovingTraPrst)
//            {
//                cnAssPtr->expectedUEBehaviour.bit_mask=expectedUEMovingTrajectory_present;
//                Uint32  numOfExpectedUEMovingTrajectoryItem_ma=
//                    coreNetworkAssInfo_m.expectedUEBeha.expectedUEMovingTraLists.size();
//                msgPdu_m.setNumOfExpectedUEMovingTrajectoryItem_v(numOfExpectedUEMovingTrajectoryItem_ma);
//                for(int i =0;i<numOfExpectedUEMovingTrajectoryItem_ma;i++)
//                {
//                    ExpectedUEMovingTrajectoryItem* expUEMovingTraItemPtr = 
//                        msgPdu_m.getExpectedUEMovingTrajectoryItemEncodePtr(i); 
//
//                    if(NULL == expUEMovingTraItemPtr)
//                    {
//                        std::cout << "NULL Pointer for expUEMovingTraItemPtr" << std::endl;
//                        return msgBuffer;
//                    }
//
//                    if(coreNetworkAssInfo_m.expectedUEBeha.expectedUEMovingTraLists[i].timeStayedPrst)
//                    {
//                        expUEMovingTraItemPtr->bit_mask = 
//                            ExpectedUEMovingTrajectoryItem_timeStayedInCell_present;
//                        expUEMovingTraItemPtr->timeStayedInCell = 
//                            coreNetworkAssInfo_m.expectedUEBeha.expectedUEMovingTraLists[i].timeStayedInCell;
//                    }
//
//                    expUEMovingTraItemPtr->nGRAN_CGI.u.nR_CGI.pLMNIdentity.length = 
//                        sizeof(coreNetworkAssInfo_m.expectedUEBeha.expectedUEMovingTraLists[i].nrcgi_a.plmn);
//                    memcpy(expUEMovingTraItemPtr->nGRAN_CGI.u.nR_CGI.pLMNIdentity.value,
//                        coreNetworkAssInfo_m.expectedUEBeha.expectedUEMovingTraLists[i].nrcgi_a.plmn,
//                            expUEMovingTraItemPtr->nGRAN_CGI.u.nR_CGI.pLMNIdentity.length);
//
//                    expUEMovingTraItemPtr->nGRAN_CGI.u.nR_CGI.nRCellIdentity.length = NRCellIdentityBitLen;
//                    memcpy(expUEMovingTraItemPtr->nGRAN_CGI.u.nR_CGI.nRCellIdentity.value,
//                        coreNetworkAssInfo_m.expectedUEBeha.expectedUEMovingTraLists[i].nrcgi_a.nrCellId,
//                            sizeof(coreNetworkAssInfo_m.expectedUEBeha.expectedUEMovingTraLists[i].nrcgi_a.nrCellId));
//                }
//            }
//        }
//    }
/////////////////////////////////////////////////////
//encode MobilityRestrictionList
/////////////////////////////////////////////////////
//
// MobilityRestrictionList*
//mobilityRestrictionListPtr=msgPdu_m.getMobilityRestrictionListEncodePtr();
// mobilityRestrictionListPtr->servingPLMN.length=sizeof(mobResList_m.plmnId);
// memcpy(mobilityRestrictionListPtr->servingPLMN.value,mobResList_m.plmnId,mobilityRestrictionListPtr->servingPLMN.length);
//
//
// Uint32 numofEquivalentPLMNs_ma=mobResList_m.equPLMNsList.size();
// msgPdu_m.setNumofPLMNIdentity_v(numofEquivalentPLMNs_ma);
// for(int i=0;i<numofEquivalentPLMNs_ma;i++)
// {
//    PLMNIdentity *plmnIdentityPtr=msgPdu_m.getPLMNIdentityEncodePtr(i);
//    if (NULL == plmnIdentityPtr)
//    {
//       std::cout << "NULL Pointer for plmnIdentityPtr" << std::endl;
//       return msgBuffer;
//    }
//    plmnIdentityPtr->length=sizeof(mobResList_m.equPLMNsList[i].plmnId);
//    memcpy(plmnIdentityPtr->value,mobResList_m.equPLMNsList[i].plmnId,plmnIdentityPtr->length);
// }
//
// Uint32 numofRATRestrictions_ma=mobResList_m.ratRestrictionsItemList.size();
// msgPdu_m.setNumofRATRestrictionsItem_v(numofRATRestrictions_ma);
// for(int i=0;i<numofRATRestrictions_ma;i++)
// {
//    RATRestrictions_Item
//*RATRestrictionsItemPtr=msgPdu_m.getRATRestrictionsItemEncodePtr(i);
//    if(NULL==RATRestrictionsItemPtr)
//    {
//       std::cout << "NULL Pointer for RATRestrictionsItemPtr" << std::endl;
//       return msgBuffer;
//    }
//    RATRestrictionsItemPtr->pLMNIdentity.length=sizeof(mobResList_m.ratRestrictionsItemList[i].plmnId);
//    memcpy(RATRestrictionsItemPtr->pLMNIdentity.value,mobResList_m.ratRestrictionsItemList[i].plmnId,RATRestrictionsItemPtr->pLMNIdentity.length);
//    RATRestrictionsItemPtr->rATRestrictionInformation.length=sizeof(mobResList_m.ratRestrictionsItemList[i].ratResInfo);
//    memcpy(RATRestrictionsItemPtr->rATRestrictionInformation.value,&(mobResList_m.ratRestrictionsItemList[i].ratResInfo),RATRestrictionsItemPtr->rATRestrictionInformation.length);
// }
//
// Uint32
//numofForbiddenAreaInformation_ma=mobResList_m.forbiddenAreaInfoList.size();
// msgPdu_m.setNumofForbiddenAreaInformationItem_v(numofForbiddenAreaInformation_ma);
// for(int i=0;i<numofForbiddenAreaInformation_ma;i++)
// {
//    Uint32
//numofForbiddenTACs_ma=mobResList_m.forbiddenAreaInfoList[i].forbiddenTACsList.size();
//    msgPdu_m.setNumofForbiddenTACs_v(numofForbiddenTACs_ma,i);
// }
// for(int i=0;i<numofForbiddenAreaInformation_ma;i++)
// {
//    ForbiddenAreaInformation_Item
//*forbiddenAreaInformationItemPtr=msgPdu_m.getForbiddenAreaInformationItemEncodePtr(i);
//    if (NULL == forbiddenAreaInformationItemPtr)
//    {
//       std::cout << "NULL Pointer for forbiddenAreaInformationItemPtr" <<
//std::endl; return msgBuffer;
//    }
//    forbiddenAreaInformationItemPtr->pLMNIdentity.length=sizeof(mobResList_m.forbiddenAreaInfoList[i].plmnId);
//    memcpy(forbiddenAreaInformationItemPtr->pLMNIdentity.value,mobResList_m.forbiddenAreaInfoList[i].plmnId,forbiddenAreaInformationItemPtr->pLMNIdentity.length);
//    Uint32
//numofForbiddenTACs_ma=mobResList_m.forbiddenAreaInfoList[i].forbiddenTACsList.size();
//    for(int j=0;j<numofForbiddenTACs_ma;j++)
//    {
//       TAC *tacPtr=msgPdu_m.getForbiddenTACsEncodePtr(i,j);
//       if(NULL == tacPtr)
//       {
//          std::cout << "NULL Pointer for tacPtr" << std::endl;
//          return msgBuffer;
//       }
//       tacPtr->length=sizeof(mobResList_m.forbiddenAreaInfoList[i].forbiddenTACsList[j].tac);
//       memcpy(tacPtr->value,mobResList_m.forbiddenAreaInfoList[i].forbiddenTACsList[j].tac,tacPtr->length);
//    }
// }
//
//
// Uint32 numofserviceAreaInformation_ma=mobResList_m.serviceAreaInfoList.size();
// msgPdu_m.setNumofServiceAreaInformationItem_v(numofserviceAreaInformation_ma);
// for(int i=0;i<numofserviceAreaInformation_ma;i++)
// {
//    Uint32
//numofAllowedTACs_ma=mobResList_m.serviceAreaInfoList[i].allowedTACsList.size();
//    msgPdu_m.setNumofAllowedTACs_v(numofAllowedTACs_ma,i);
//    Uint32
//numofNotAllowedTACs_ma=mobResList_m.serviceAreaInfoList[i].notAllowedTACsList.size();
//    msgPdu_m.setNumofNotAllowedTACs_v(numofNotAllowedTACs_ma,i);
// }
// for(int i=0;i<numofserviceAreaInformation_ma;i++)
// {
//    ServiceAreaInformation_Item
//*ServiceAreaInformationItemPtr=msgPdu_m.getServiceAreaInformationItemEncodePtr(i);
//    if (NULL == ServiceAreaInformationItemPtr)
//    {
//       std::cout << "NULL Pointer for ServiceAreaInformationItemPtr" <<
//std::endl; return msgBuffer;
//    }
//    ServiceAreaInformationItemPtr->pLMNIdentity.length=sizeof(mobResList_m.serviceAreaInfoList[i].plmnId);
//    memcpy(ServiceAreaInformationItemPtr->pLMNIdentity.value,mobResList_m.serviceAreaInfoList[i].plmnId,ServiceAreaInformationItemPtr->pLMNIdentity.length);
//    Uint32
//numofAllowedTACs_ma=mobResList_m.serviceAreaInfoList[i].allowedTACsList.size();
//    for(int j = 0; j < numofAllowedTACs_ma; j++)
//    {
//       TAC *tacPtr=msgPdu_m.getAllowedTACsEncodePtr(i,j);
//       if (NULL == tacPtr)
//       {
//       std::cout << "NULL Pointer for tacPtr" << std::endl;
//       return msgBuffer;
//       }
//       tacPtr->length=sizeof(mobResList_m.serviceAreaInfoList[i].allowedTACsList[j].tac);
//       memcpy(tacPtr->value,mobResList_m.serviceAreaInfoList[i].allowedTACsList[j].tac,tacPtr->length);
//    }
//    Uint32
//numofNotAllowedTACs_ma=mobResList_m.serviceAreaInfoList[i].notAllowedTACsList.size();
//    for(int j = 0; j < numofNotAllowedTACs_ma; j++)
//    {
//       TAC *tacPtr=msgPdu_m.getNotAllowedTACsEncodePtr(i,j);
//       if (NULL == tacPtr)
//       {
//       std::cout << "NULL Pointer for tacPtr" << std::endl;
//       return msgBuffer;
//       }
//       tacPtr->length=sizeof(mobResList_m.serviceAreaInfoList[i].notAllowedTACsList[j].tac);
//       memcpy(tacPtr->value,mobResList_m.serviceAreaInfoList[i].notAllowedTACsList[j].tac,tacPtr->length);
//    }
// }
/////////////////////////////////////////////////////
//decode CoreNetworkAssistanceInformation
/////////////////////////////////////////////////////
//    if(msgPdu_m.isCoreNetworkPresent())
//    {
//        memcpy(coreNetworkAssInfo_m.ueIdentityIndexValue,
//            msgPdu_m.getCoreNetworkAssistanceInformationDecodedPtr()
//                ->uEIdentityIndexValue.u.indexLength10.value,
//                    uEIdentityIndexValueBitLen);
//
//        if(msgPdu_m.getCoreNetworkAssistanceInformationDecodedPtr()
//            ->bit_mask==uESpecificDRX_present)
//        {
//            coreNetworkAssInfo_m.pagingDRXPrst=true;
//            coreNetworkAssInfo_m.pagingDRX = 
//                msgPdu_m.getCoreNetworkAssistanceInformationDecodedPtr()->uESpecificDRX;
//        }
//        memcpy(&(coreNetworkAssInfo_m.PerRegUpdateTimer),
//            msgPdu_m.getCoreNetworkAssistanceInformationDecodedPtr()
//                ->periodicRegistrationUpdateTimer.value,
//                    msgPdu_m.getCoreNetworkAssistanceInformationDecodedPtr()
//                        ->periodicRegistrationUpdateTimer.length);
//
//        if(msgPdu_m.getCoreNetworkAssistanceInformationDecodedPtr()
//            ->bit_mask==mICOModeIndication_present)
//        {
//            coreNetworkAssInfo_m.mICOModeIndPrst = true;
//            coreNetworkAssInfo_m.mICOModeIndication = 
//                msgPdu_m.getCoreNetworkAssistanceInformationDecodedPtr()->mICOModeIndication;
//        }
//
//        if(msgPdu_m.getCoreNetworkAssistanceInformationDecodedPtr()
//            ->bit_mask == expectedUEBehaviour_present)
//        {
//            coreNetworkAssInfo_m.expectUEBehaPrst=true;
//            for(Uint32 i = 0; i < msgPdu_m.getNumOfTAIListForInactiveItem(); i++)
//            {
//                TAIListForInactiveItem *taiListForInactivePtr = 
//                    msgPdu_m.getTAIListForInactiveItemDecodedPtr(i);
//                if(NULL==taiListForInactivePtr)
//                {
//                    return false;
//                }
//                TAILFIItem taiLFIItem_m;
//
//                memcpy(taiLFIItem_m.tac,
//                    taiListForInactivePtr->tAI.tAC.value,
//                        taiListForInactivePtr->tAI.tAC.length);
//
//                memcpy(taiLFIItem_m.plmnId,
//                    taiListForInactivePtr->tAI.pLMNIdentity.value,
//                        taiListForInactivePtr->tAI.pLMNIdentity.length);
//
//                coreNetworkAssInfo_m.taiLFIList.push_back(taiLFIItem_m);
//            }
//
//        if(msgPdu_m.getCoreNetworkAssistanceInformationDecodedPtr()
//            ->expectedUEBehaviour.bit_mask == expectedUEBehaviour_present)
//        {
//            coreNetworkAssInfo_m.expectedUEBeha.expectUEActBehaPrst=true;
//   
//            if(msgPdu_m.getCoreNetworkAssistanceInformationDecodedPtr()
//                ->expectedUEBehaviour.expectedUEActivityBehaviour.bit_mask ==
//                     expectedActivityPeriod_present)
//            {
//                coreNetworkAssInfo_m.expectedUEBeha.expectedUEActBeha.expActPeriodPrst=true;
//                coreNetworkAssInfo_m.expectedUEBeha.expectedUEActBeha.expectedActivityPeriod = 
//                    msgPdu_m.getCoreNetworkAssistanceInformationDecodedPtr()
//                        ->expectedUEBehaviour.expectedUEActivityBehaviour.expectedActivityPeriod;
//            }
//
//            if(msgPdu_m.getCoreNetworkAssistanceInformationDecodedPtr()
//                ->expectedUEBehaviour.expectedUEActivityBehaviour.bit_mask == 
//                    expectedIdlePeriod_present)
//            {
//                coreNetworkAssInfo_m.expectedUEBeha.expectedUEActBeha.expIdlePeriodPrst=true;
//                coreNetworkAssInfo_m.expectedUEBeha.expectedUEActBeha.expectedIdlePeriod = 
//                    msgPdu_m.getCoreNetworkAssistanceInformationDecodedPtr()
//                        ->expectedUEBehaviour.expectedUEActivityBehaviour.expectedIdlePeriod;
//            }
//
//            if(msgPdu_m.getCoreNetworkAssistanceInformationDecodedPtr()
//                ->expectedUEBehaviour.expectedUEActivityBehaviour.bit_mask == 
//                    sourceOfUEActivityBehaviourInformation_present)
//            {
//                coreNetworkAssInfo_m.expectedUEBeha.expectedUEActBeha.sourceActBehaPrst=true;
//                coreNetworkAssInfo_m.expectedUEBeha.expectedUEActBeha.sourceOfUEActBehaInfo = 
//                    msgPdu_m.getCoreNetworkAssistanceInformationDecodedPtr()->expectedUEBehaviour
//                        .expectedUEActivityBehaviour.sourceOfUEActivityBehaviourInformation;
//            }
//
//        }
//
//
//        if(msgPdu_m.getCoreNetworkAssistanceInformationDecodedPtr()
//            ->expectedUEBehaviour.bit_mask == expectedHOInterval_present)
//        {
//            coreNetworkAssInfo_m.expectedUEBeha.expHOIntervalPrst=true;
//            coreNetworkAssInfo_m.expectedUEBeha.expHOInterval = 
//                msgPdu_m.getCoreNetworkAssistanceInformationDecodedPtr()
//                    ->expectedUEBehaviour.expectedHOInterval;
//        }
//
//        if(msgPdu_m.getCoreNetworkAssistanceInformationDecodedPtr()
//            ->expectedUEBehaviour.bit_mask==expectedUEMobility_present)
//        {
//            coreNetworkAssInfo_m.expectedUEBeha.expUEMobilityPrst=true;
//            coreNetworkAssInfo_m.expectedUEBeha.expUEMobility = 
//                msgPdu_m.getCoreNetworkAssistanceInformationDecodedPtr()
//                    ->expectedUEBehaviour.expectedUEMobility;
//        }
//
//        if(msgPdu_m.getCoreNetworkAssistanceInformationDecodedPtr()
//            ->expectedUEBehaviour.bit_mask == expectedUEMovingTrajectory_present)
//        {
//            coreNetworkAssInfo_m.expectedUEBeha.expUEMovingTraPrst=true;
//            for(Uint32 i=0 ;i<msgPdu_m.getNumOfExpectedUEMovingTrajectoryItem();i++)
//            {
//                ExpectedUEMovingTrajectoryItem* ExpectedUEMovingTrajectoryPtr = 
//                    msgPdu_m.getExpectedUEMovingTrajectoryItemDecodedPtr(i);
//                if(NULL==ExpectedUEMovingTrajectoryPtr)
//                {
//                    return false;
//                }
//                ExpectedUEMovingTra ExpectedUEMovingTra_m;
//                memcpy(ExpectedUEMovingTra_m.nrcgi_a.plmn,
//                    ExpectedUEMovingTrajectoryPtr->nGRAN_CGI.u.nR_CGI.pLMNIdentity.value,
//                        ExpectedUEMovingTrajectoryPtr->nGRAN_CGI.u.nR_CGI.pLMNIdentity.length);
//
//                memcpy(ExpectedUEMovingTra_m.nrcgi_a.nrCellId,
//                    ExpectedUEMovingTrajectoryPtr->nGRAN_CGI.u.nR_CGI.nRCellIdentity.value,
//                        ExpectedUEMovingTrajectoryPtr->nGRAN_CGI.u.nR_CGI.nRCellIdentity.length);    
//
//                if(ExpectedUEMovingTrajectoryPtr->bit_mask == 
//                    ExpectedUEMovingTrajectoryItem_timeStayedInCell_present)
//                {
//                    ExpectedUEMovingTra_m.timeStayedPrst=true;
//                    ExpectedUEMovingTra_m.timeStayedInCell = 
//                        ExpectedUEMovingTrajectoryPtr->timeStayedInCell;
//                }
//                coreNetworkAssInfo_m.expectedUEBeha.expectedUEMovingTraLists.push_back(ExpectedUEMovingTra_m);
//            }
//        }
//    }
//    coreNetworkPrst_m=true;
//}
/////////////////////////////////////////////////////
 //decode MobilityRestrictionList
 /////////////////////////////////////////////////////
// MobilityRestrictionList
//*mobilityRestrictionListPtr=msgPdu_m.getMobilityRestrictionListDecodedPtr();
// if(mobilityRestrictionListPtr==NULL)
// {
//    return false;
// }
// MobResList mobResList_p;
// memcpy(mobResList_p.plmnId,mobilityRestrictionListPtr->servingPLMN.value,mobilityRestrictionListPtr->servingPLMN.length);
// for(Uint32 i = 0; i < msgPdu_m.getNumofPLMNIdentity(); i++)
// {
//    PLMNIdentity *plmnIdentityPtr=msgPdu_m.getPLMNIdentityDecodedPtr(i);
//    if (NULL == plmnIdentityPtr)
//    {
//       return false;
//    }
//    EquPLMNsItem equPLMNsItem_m;
//    memcpy(equPLMNsItem_m.plmnId,plmnIdentityPtr->value,plmnIdentityPtr->length);
////      mobResList_p.equPLMNsList.push_back(equPLMNsList_m);
//    mobResList_p.equPLMNsList.push_back(equPLMNsItem_m);
// }
//
// for(Uint32 i = 0; i < msgPdu_m.getNumofRATRestrictionsItem(); i++)
// {
//    RATRestrictions_Item
//*ratRestrictionsItemPtr=msgPdu_m.getRATRestrictionsItemDecodedPtr(i); if (NULL
//== ratRestrictionsItemPtr)
//    {
//       return false;
//    }
//    RATRestrictionsItem ratRestrictionsItemList_m;
//    memcpy(ratRestrictionsItemList_m.plmnId,ratRestrictionsItemPtr->pLMNIdentity.value,ratRestrictionsItemPtr->pLMNIdentity.length);
//    memcpy(&(ratRestrictionsItemList_m.ratResInfo),ratRestrictionsItemPtr->rATRestrictionInformation.value,ratRestrictionsItemPtr->rATRestrictionInformation.length);
//    mobResList_p.ratRestrictionsItemList.push_back(ratRestrictionsItemList_m);
// }
//
// for(Uint32 i = 0; i < msgPdu_m.getNumofForbiddenAreaInformationItem(); i++)
// {
//    ForbiddenAreaInformation_Item
//*forbiddenAreaInformationItemPtr=msgPdu_m.getForbiddenAreaInformationItemDecodedPtr(i);
//    if (NULL == forbiddenAreaInformationItemPtr)
//    {
//       return false;
//    }
//    ForbiddenAreaInfoItem forbiddenAreaInfoList_m;
//    memcpy(forbiddenAreaInfoList_m.plmnId,forbiddenAreaInformationItemPtr->pLMNIdentity.value,forbiddenAreaInformationItemPtr->pLMNIdentity.length);
//    for(Uint32 j = 0; j < msgPdu_m.getNumofForbiddenTACs(i); j++)
//    {
//       TAC
//*tacPtr=msgPdu_m.getForbiddenTACsDecodedPtr(forbiddenAreaInformationItemPtr,j);
//       if (NULL == tacPtr)
//       {
//          return false;
//       }
//       ForbiddenTACsItem forbiddenTACsItem_m;
//       memcpy(forbiddenTACsItem_m.tac,tacPtr->value,tacPtr->length);
//       forbiddenAreaInfoList_m.forbiddenTACsList.push_back(forbiddenTACsItem_m);
//    }
//    mobResList_p.forbiddenAreaInfoList.push_back(forbiddenAreaInfoList_m);
// }
//
// for(Uint32 i = 0; i < msgPdu_m.getNumofServiceAreaInformationItem(); i++)
// {
//    ServiceAreaInformation_Item
//*serviceAreaInformationItemPtr=msgPdu_m.getServiceAreaInformationItemDecodedPtr(i);
//    if (NULL == serviceAreaInformationItemPtr)
//    {
//       return false;
//    }
//    ServiceAreaInfoItem serviceAreaInfoList_m;
//    memcpy(serviceAreaInfoList_m.plmnId,serviceAreaInformationItemPtr->pLMNIdentity.value,serviceAreaInformationItemPtr->pLMNIdentity.length);
//    for(Uint32 j = 0; j < msgPdu_m.getNumofAllowedTACs(j); j++)
//    {
//       TAC
//*tacPtr=msgPdu_m.getAllowedTACsDecodedPtr(serviceAreaInformationItemPtr,j); if
//(NULL == tacPtr)
//       {
//          return false;
//       }
//       AllowedTACsItem allowedTACsList_m;
//       memcpy(allowedTACsList_m.tac,tacPtr->value,tacPtr->length);
//       serviceAreaInfoList_m.allowedTACsList.push_back(allowedTACsList_m);
//    }
//    for(Uint32 j = 0; j < msgPdu_m.getNumofNotAllowedTACs(j); j++)
//    {
//       TAC
//*tacPtr=msgPdu_m.getNotAllowedTACsDecodedPtr(serviceAreaInformationItemPtr,j);
//       if (NULL == tacPtr)
//       {
//          return false;
//       }
//       NotAllowedTACsItem notAllowedTACsList_m;
//       memcpy(notAllowedTACsList_m.tac,tacPtr->value,tacPtr->length);
//       serviceAreaInfoList_m.notAllowedTACsList.push_back(notAllowedTACsList_m);
//    }
//    mobResList_p.serviceAreaInfoList.push_back(serviceAreaInfoList_m);
// }

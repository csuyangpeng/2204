#include <iostream>
#include <string.h>

#include "initialContextSetupRequestPdu.h"

InitialContextSetupRequestPdu::InitialContextSetupRequestPdu()
{
    reset_v();
}

//----------------------------------------------------------------------------
// This is a public API. Refer to InitialContextSetupRequestPdu.h
//----------------------------------------------------------------------------
InitialContextSetupRequestPdu::~InitialContextSetupRequestPdu()
{
}

//----------------------------------------------------------------------------
// This is a public API. Refer to InitialContextSetupRequestPdu.h
//----------------------------------------------------------------------------
void InitialContextSetupRequestPdu::reset_v()
{
    InitialContextSetupRequest_mp = NULL;

    //RAN UE NGAP ID
    ranUeNgapId_m = 0;
    ranUeNgapId_mp = NULL;

    //AMF UE NGAP ID
    amfUeNgapId_m = 0;
    amfUeNgapId_mp = NULL;

    //Old AMF
    amfName_mp = NULL;
    amfName_m.length = 0;
    memset(amfNameValue_ma, 0xff, sizeof(amfNameValue_ma));
    amfName_m.value = amfNameValue_ma;
    amfNamePresent_m=false;


    //UE Aggregate Maximum Bit Rate
    ueAggregateMaximumBitRate_m.bit_mask = 0;
    ueAggregateMaximumBitRate_m.uEAggregateMaximumBitRateDL = 0;
    ueAggregateMaximumBitRate_m.uEAggregateMaximumBitRateUL = 0;
    ueAggregateMaximumBitRate_m.iE_Extensions=NULL;
    ueAggregateMaximumBitRate_mp = NULL;
    ueAMBRPresent_m=false;

    //GUAMI
    guami_mp = NULL;
    guami_m.bit_mask = 0;
    guami_m.pLMNIdentity.length = 0;
    memset(guami_m.pLMNIdentity.value, 0xff, sizeof(guami_m.pLMNIdentity.value));
    memset(&guamiIdentifier_ma, 0xff, sizeof(guamiIdentifier_ma));
    guami_m.aMFSetID.length = 0;
    guami_m.aMFSetID.value = guamiIdentifier_ma.setId;
    guami_m.aMFRegionID.length = 0;
    guami_m.aMFRegionID.value = &(guamiIdentifier_ma.regionId);
    guami_m.aMFPointer.length = 0;
    guami_m.aMFPointer.value = &(guamiIdentifier_ma.pointer);
    guami_m.iE_Extensions=NULL;

    //Allowed NSSAI
    allowedNssai_m = NULL;
    allowedNssai_mp = NULL;
    numOfAllowedNassiItem_m = 0;

    for (unsigned int i = 0; i < maxnoofNSSAIs_c; i++) {
        allowedNssai_ma[i].next = NULL;
        allowedNssai_ma[i].value.bit_mask = 0;
        allowedNssai_ma[i].value.iE_Extensions = NULL;
        allowedNssai_ma[i].value.s_NSSAI.bit_mask = 0;
        allowedNssai_ma[i].value.s_NSSAI.sST.length = 0;
        memset(allowedNssai_ma[i].value.s_NSSAI.sST.value, 0xff,
            sizeof(allowedNssai_ma[i].value.s_NSSAI.sST.value));
        allowedNssai_ma[i].value.s_NSSAI.sD.length = 0;
        memset(allowedNssai_ma[i].value.s_NSSAI.sD.value, 0xff,
            sizeof(allowedNssai_ma[i].value.s_NSSAI.sD.value));
        allowedNssai_ma[i].value.s_NSSAI.iE_Extensions = NULL;
    }

    //Ue Security Capabilities
    ueSecurityCapabilities_mp = NULL;
    ueSecurityCapabilities_m.bit_mask = 0;
    memset(&ueSecurityCapabilitiesValue_ma, 0xff, sizeof(ueSecurityCapabilitiesValue_ma));

    ueSecurityCapabilities_m.nRencryptionAlgorithms.length = 0;
    ueSecurityCapabilities_m.nRencryptionAlgorithms.value = 
        ueSecurityCapabilitiesValue_ma.nrEncryptionAlg;

    ueSecurityCapabilities_m.nRintegrityProtectionAlgorithms.length = 0;
    ueSecurityCapabilities_m.nRintegrityProtectionAlgorithms.value = 
        ueSecurityCapabilitiesValue_ma.nrIntegrityProAlg;

    ueSecurityCapabilities_m.eUTRAencryptionAlgorithms.length = 0;
    ueSecurityCapabilities_m.eUTRAencryptionAlgorithms.value = 
        ueSecurityCapabilitiesValue_ma.eutraEncryptionAlg;

    ueSecurityCapabilities_m.eUTRAintegrityProtectionAlgorithms.length = 0;
    ueSecurityCapabilities_m.eUTRAintegrityProtectionAlgorithms.value = 
        ueSecurityCapabilitiesValue_ma.eutraIntegrityProAlg;

    ueSecurityCapabilities_m.iE_Extensions = NULL;

    //SecurityKey
    securityKey_mp = NULL;
    securityKey_m.length = 0;
    memset(securityKeyValue_ma, 0xff, sizeof(securityKeyValue_ma));
    securityKey_m.value = securityKeyValue_ma;

    //TraceActivation
    traceActivation_mp = NULL;
    traceActivation_m.bit_mask = 0;
    traceActivation_m.nGRANTraceID.length = 0;
    memset(traceActivation_m.nGRANTraceID.value, 0xff, sizeof(traceActivation_m.nGRANTraceID.value));
    memset(&traceActivationValue_ma, 0xff, sizeof(traceActivationValue_ma));
    traceActivation_m.interfacesToTrace.length = 0;
    traceActivation_m.interfacesToTrace.value = &(traceActivationValue_ma.interfaces);
    traceActivation_m.traceDepth = minimum;
    traceActivation_m.traceCollectionEntityIPAddress.length = 0;
    traceActivation_m.traceCollectionEntityIPAddress.value = 
        traceActivationValue_ma.traceEntityIPAddr;
    traceActivation_m.iE_Extensions = NULL;
    traceActPresent_m=false;

    //UERadioCapability
    ueRadioCapability_mp = NULL;
    ueRadioCapability_m.length = 0;
    memset(ueRadioCapabilityValue_ma, 0xff, sizeof(ueRadioCapabilityValue_ma));
    ueRadioCapability_m.value = ueRadioCapabilityValue_ma;
    ueRadioCapPresent_m=false;

    //Index To RFSP
    indexToRFSP_m = 0;
    indexToRFSP_mp = NULL;
    indexToRFSPPresent_m=false;

    //MaskedIMEISV
    maskedIMEISV_mp = NULL;
    maskedIMEISV_m.length = 0;
    memset(maskedIMEISVValue_ma, 0xff, sizeof(maskedIMEISVValue_ma));
    maskedIMEISV_m.value = maskedIMEISVValue_ma;
    maskedIMEISVPresent_m=false;

    //NAS PDU
    nasPdu_m.length = 0;
    memset(nasPdu_ma, 0xff, sizeof(nasPdu_ma));
    nasPdu_m.value = nasPdu_ma;
    nasPdu_mp = NULL;
    nasPduPresent_m=false;

    //EmergencyFallbackIndicator
    emergencyFallbackIndicator_mp = NULL;
    emergencyFallbackIndicator_m.bit_mask = 0;
    emergencyFallbackIndicator_m.emergencyFallbackRequestIndicator = 
        emergency_fallback_requested;
    emergencyFallbackIndicator_m.emergencyServiceTargetCN = fiveGC;
    emergencyFallbackIndicator_m.iE_Extensions = NULL;
    emerFallbackPresent_m=false;

    //RRCInactiveTransitionReportRequest
    rrcInactiveTransitionReportRequest_mp = NULL;
    rrcInactiveTransitionReportRequest_m = subsequent_state_transition_report;
    RRCReportPresent_m=false;

    //UERadioCapabilityForPaging
    ueRadioCapabilityForPaging_mp = NULL;
    ueRadioCapabilityForPaging_m.uERadioCapabilityForPagingOfNR.length = 0;
    memset(ueRadioCapabilityForPagingNr_ma, 0xff, sizeof(ueRadioCapabilityForPagingNr_ma));
    ueRadioCapabilityForPaging_m.uERadioCapabilityForPagingOfNR.value = ueRadioCapabilityForPagingNr_ma;
    ueRadioCapabilityForPaging_m.uERadioCapabilityForPagingOfEUTRA.length = 0;
    memset(ueRadioCapabilityForPagingEutra_ma, 0xff, sizeof(ueRadioCapabilityForPagingEutra_ma));
    ueRadioCapabilityForPaging_m.uERadioCapabilityForPagingOfEUTRA.value = ueRadioCapabilityForPagingEutra_ma;
    ueRadioCapabilityForPagingPresent_m=false;

    //RedirectionVoiceFallback
    redirectionVoiceFallback_mp = NULL;
    redirectionVoiceFallback_m = possible;
    redirectionVoiceFallbackPresent_m=false;

    //CoreNetworkAssistanceInformation
    coreNetworkAssistanceInfo_mp = NULL;
    coreNetworkAssistanceInfo_m.bit_mask = 0;
    coreNetworkAssistanceInfo_m.iE_Extensions = NULL;
    coreNetworkAssistanceInfo_m.uEIdentityIndexValue.choice = indexLength10_chosen;
    coreNetworkAssistanceInfo_m.uEIdentityIndexValue.u.indexLength10.length = 0;
    memset(CoreNetAssInfo_ma.ueIdentityIndValue, 0xff, sizeof(CoreNetAssInfo_ma.ueIdentityIndValue));
    coreNetworkAssistanceInfo_m.uEIdentityIndexValue.u.indexLength10.value = CoreNetAssInfo_ma.ueIdentityIndValue;
    coreNetworkAssistanceInfo_m.uESpecificDRX = v32;
    coreNetworkAssistanceInfo_m.periodicRegistrationUpdateTimer.length = 0;
    CoreNetAssInfo_ma.PerRegUpdateTimer = 0;
    coreNetworkAssistanceInfo_m.periodicRegistrationUpdateTimer.value= &(CoreNetAssInfo_ma.PerRegUpdateTimer);
    coreNetworkAssistanceInfo_m.mICOModeIndication = MICOModeIndication_true;
    numOfTAIListForInactiveItem_m = 0;
    coreNetworkAssistanceInfo_m.tAIListForInactive = NULL;
    for (unsigned int i = 0; i < maxnoofTAIforInactive; i++) {
        taiListForInactive_ma[i].next = NULL;
        taiListForInactive_ma[i].value.bit_mask = 0;
        taiListForInactive_ma[i].value.iE_Extensions=NULL;
        taiListForInactive_ma[i].value.tAI.bit_mask = 0;
        taiListForInactive_ma[i].value.tAI.iE_Extensions=NULL;
        taiListForInactive_ma[i].value.tAI.pLMNIdentity.length = 0;
        memset(taiListForInactive_ma[i].value.tAI.pLMNIdentity.value, 0xff, 
            sizeof(taiListForInactive_ma[i].value.tAI.pLMNIdentity.value));

        taiListForInactive_ma[i].value.tAI.tAC.length = 0;
        memset(taiListForInactive_ma[i].value.tAI.tAC.value, 0xff, 
            sizeof(taiListForInactive_ma[i].value.tAI.tAC.value));
    }
    coreNetworkAssistanceInfo_m.expectedUEBehaviour.bit_mask = 0;
    coreNetworkAssistanceInfo_m.expectedUEBehaviour.iE_Extensions=NULL;
    coreNetworkAssistanceInfo_m.expectedUEBehaviour.expectedUEActivityBehaviour.bit_mask = 0;
    coreNetworkAssistanceInfo_m.expectedUEBehaviour.expectedUEActivityBehaviour.iE_Extensions = NULL;
    coreNetworkAssistanceInfo_m.expectedUEBehaviour.expectedUEActivityBehaviour.expectedActivityPeriod = 0;
    coreNetworkAssistanceInfo_m.expectedUEBehaviour.expectedUEActivityBehaviour.expectedIdlePeriod = 0;
    coreNetworkAssistanceInfo_m.expectedUEBehaviour.expectedUEActivityBehaviour.sourceOfUEActivityBehaviourInformation 
        = subscription_information;
    coreNetworkAssistanceInfo_m.expectedUEBehaviour.expectedHOInterval = sec15;
    coreNetworkAssistanceInfo_m.expectedUEBehaviour.expectedUEMobility = stationary;
    coreNetworkAssistanceInfo_m.expectedUEBehaviour.expectedUEMovingTrajectory = NULL;
    numOfExpectedUEMovingTrajectoryItem_m = 0;
    for (unsigned int i = 0; i < maxnoofCellsUEMovingTrajectory_c; i++) {
        expectedUEMovingTrajectory_ma[i].next = NULL;
        expectedUEMovingTrajectory_ma[i].value.bit_mask = 0;
        expectedUEMovingTrajectory_ma[i].value.iE_Extensions=NULL;
        expectedUEMovingTrajectory_ma[i].value.nGRAN_CGI.choice = nR_CGI_chosen;
        expectedUEMovingTrajectory_ma[i].value.nGRAN_CGI.u.nR_CGI.bit_mask = 0;
        expectedUEMovingTrajectory_ma[i].value.nGRAN_CGI.u.nR_CGI.iE_Extensions=NULL;
        expectedUEMovingTrajectory_ma[i].value.nGRAN_CGI.u.nR_CGI.pLMNIdentity.length = 0;
        memset(expectedUEMovingTrajectory_ma[i].value.nGRAN_CGI.u.nR_CGI.pLMNIdentity.value,
            0xff, 
            sizeof(expectedUEMovingTrajectory_ma[i].value.nGRAN_CGI.u.nR_CGI.pLMNIdentity.value));
        expectedUEMovingTrajectory_ma[i].value.nGRAN_CGI.u.nR_CGI.nRCellIdentity.length = 0;
        expectedUEMovingTrajectory_ma[i].value.nGRAN_CGI.u.nR_CGI.nRCellIdentity.value = &(nrCellmaId[i]);
        expectedUEMovingTrajectory_ma[i].value.timeStayedInCell = 0;
    }
    coreNetworkPresent_m=false;

    //MobilityRestrictionList
    mobilityRestrictionList_mp = NULL;
    mobilityRestrictionList_m.bit_mask = 0;
    mobilityRestrictionList_m.iE_Extensions=NULL;
    mobilityRestrictionList_m.servingPLMN.length = 0;
    memset(mobilityRestrictionList_m.servingPLMN.value, 0xff, 
        sizeof(mobilityRestrictionList_m.servingPLMN.value));
    mobilityRestrictionList_m.equivalentPLMNs = NULL;
    numofEquivalentPLMNs_m = 0;
    for (unsigned int i = 0; i < maxnoofEPLMNs_c; i++) {
        equivalentPLMNs_ma[i].next = NULL;
        equivalentPLMNs_ma[i].value.length = 0;
        memset(equivalentPLMNs_ma[i].value.value, 0xff, 
            sizeof(equivalentPLMNs_ma[i].value.value));
    }
    mobilityRestrictionList_m.rATRestrictions = NULL;
    numofRATRestrictions_m = 0;
    for (unsigned int i = 0; i < maxnoofEPLMNsPlusOne_c; i++) {
        rATRestrictions_ma[i].next = NULL;
        rATRestrictions_ma[i].value.bit_mask = 0;
        rATRestrictions_ma[i].value.iE_Extensions=NULL;
        rATRestrictions_ma[i].value.pLMNIdentity.length = 0;
        memset(rATRestrictions_ma[i].value.pLMNIdentity.value, 0xff, 
            sizeof(rATRestrictions_ma[i].value.pLMNIdentity.value));
        rATRestrictions_ma[i].value.rATRestrictionInformation.length = 0;
        memset(ratResInfo_ma, 0xff, sizeof(ratResInfo_ma));
        rATRestrictions_ma[i].value.rATRestrictionInformation.value = &(ratResInfo_ma[i]);
    }
    mobilityRestrictionList_m.forbiddenAreaInformation = NULL;
    numofForbiddenAreaInformation_m = 0;
    for (unsigned int i = 0; i < maxnoofEPLMNsPlusOne_c; i++) {
        forbiddenAreaInformation_ma[i].next = NULL;
        forbiddenAreaInformation_ma[i].value.bit_mask = 0;
        forbiddenAreaInformation_ma[i].value.iE_Extensions=NULL;
        forbiddenAreaInformation_ma[i].value.pLMNIdentity.length = 0;
        memset(forbiddenAreaInformation_ma[i].value.pLMNIdentity.value, 0xff, 
            sizeof(forbiddenAreaInformation_ma[i].value.pLMNIdentity.value));
        numofForbiddenTACs_m[i] = 0;
        forbiddenAreaInformation_ma[i].value.forbiddenTACs = NULL;
        for (unsigned int j = 0; j < maxnoofForbTACs_c; j++) {
            forbiddenTACs_ma[i][j].next = NULL;
            forbiddenTACs_ma[i][j].value.length = 0;
            memset(forbiddenTACs_ma[i][j].value.value, 0xff, 
                sizeof(forbiddenTACs_ma[i][j].value.value));
        }
    }
    mobilityRestrictionList_m.serviceAreaInformation = NULL;
    numofserviceAreaInformation_m = 0;
    for (unsigned int i = 0; i < maxnoofEPLMNsPlusOne_c; i++) {
        serviceAreaInformation_ma[i].next = NULL;
        serviceAreaInformation_ma[i].value.bit_mask = 0;
        serviceAreaInformation_ma[i].value.iE_Extensions=NULL;
        serviceAreaInformation_ma[i].value.pLMNIdentity.length = 0;
        memset(serviceAreaInformation_ma[i].value.pLMNIdentity.value, 0xff, 
            sizeof(serviceAreaInformation_ma[i].value.pLMNIdentity.value));
        numofAllowedTACs_m[i] = 0;
        serviceAreaInformation_ma[i].value.allowedTACs = NULL;
        for (unsigned int j = 0; j < maxnoofEPLMNsPlusOne_c; j++) {
            allowedTACs_ma[i][j].next = NULL;
            allowedTACs_ma[i][j].value.length = 0;
            memset(allowedTACs_ma[i][j].value.value, 0xff, 
                sizeof(allowedTACs_ma[i][j].value.value));
        }
        numofNotAllowedTACs_m[i] = 0;
        serviceAreaInformation_ma[i].value.notAllowedTACs = NULL;
        for (unsigned int k = 0; k < maxnoofEPLMNsPlusOne_c; k++) {
            notAllowedTACs_ma[i][k].next = NULL;
            notAllowedTACs_ma[i][k].value.length = 0;
            memset(notAllowedTACs_ma[i][k].value.value, 0xff, 
                sizeof(notAllowedTACs_ma[i][k].value.value));
        }
    }

    //PDUSessionResourceSetupListCxtReq
    pDUSessionResourceSetupListCxtReq_mp = NULL;
    pDUSessionResourceSetupListCxtReq_m = NULL;
    numofPDUSessionResourceSetupListItemCxtReq_m = 0;

    memset(pduNasPdus_ma, 0xff, sizeof(pduNasPdus_ma));
    memset(transfers_ma, 0xff, sizeof(transfers_ma));

    for (unsigned int i = 0; i < maxnoofPDUSessions_c; i++) {
        PDUSessionResSetupListCxtReq_ma[i].next = NULL;
        PDUSessionResSetupListCxtReq_ma[i].value.bit_mask = 0;
        PDUSessionResSetupListCxtReq_ma[i].value.iE_Extensions=NULL;
        PDUSessionResSetupListCxtReq_ma[i].value.pDUSessionID = 0;
        PDUSessionResSetupListCxtReq_ma[i].value.nAS_PDU.length = 0;
        PDUSessionResSetupListCxtReq_ma[i].value.nAS_PDU.value = pduNasPdus_ma[i].data;
        PDUSessionResSetupListCxtReq_ma[i].value.s_NSSAI.bit_mask = 0;
        PDUSessionResSetupListCxtReq_ma[i].value.s_NSSAI.iE_Extensions=NULL;
        PDUSessionResSetupListCxtReq_ma[i].value.s_NSSAI.sST.length = 0;
        memset(PDUSessionResSetupListCxtReq_ma[i].value.s_NSSAI.sST.value, 0xff, 
            sizeof(PDUSessionResSetupListCxtReq_ma[i].value.s_NSSAI.sST.value));
        PDUSessionResSetupListCxtReq_ma[i].value.s_NSSAI.sD.length = 0;
        memset(PDUSessionResSetupListCxtReq_ma[i].value.s_NSSAI.sD.value, 0xff, 
            sizeof(PDUSessionResSetupListCxtReq_ma[i].value.s_NSSAI.sD.value));
        PDUSessionResSetupListCxtReq_ma[i].value.pDUSessionResourceSetupRequestTransfer.encoded.length = 0;
        PDUSessionResSetupListCxtReq_ma[i].value.pDUSessionResourceSetupRequestTransfer.encoded.value = transfers_ma[i].data;
        PDUSessionResSetupListCxtReq_ma[i].value.pDUSessionResourceSetupRequestTransfer.decoded = NULL;
    }

    //CNAssistedRANTuning
    cnAssistedRANTuning_mp = NULL;
    cnAssistedRANTuning_m.bit_mask = 0;
    cnAssistedRANTuning_m.iE_Extensions = NULL;
    cnAssistedRANTuning_m.expectedUEBehaviour.bit_mask = 0;
    cnAssistedRANTuning_m.expectedUEBehaviour.iE_Extensions=NULL;
    cnAssistedRANTuning_m.expectedUEBehaviour.expectedUEActivityBehaviour.bit_mask = 0;
    cnAssistedRANTuning_m.expectedUEBehaviour.expectedUEActivityBehaviour.iE_Extensions = NULL;
    cnAssistedRANTuning_m.expectedUEBehaviour.expectedUEActivityBehaviour.expectedActivityPeriod = 0;
    cnAssistedRANTuning_m.expectedUEBehaviour.expectedUEActivityBehaviour.expectedIdlePeriod = 0;
    cnAssistedRANTuning_m.expectedUEBehaviour.expectedUEActivityBehaviour.sourceOfUEActivityBehaviourInformation
        = subscription_information;
    cnAssistedRANTuning_m.expectedUEBehaviour.expectedHOInterval = sec15;
    cnAssistedRANTuning_m.expectedUEBehaviour.expectedUEMobility = stationary;
    cnAssistedRANTuning_m.expectedUEBehaviour.expectedUEMovingTrajectory = NULL;
    numOfExpectedUEMovingTrajectoryItem_mb = 0;
    for (unsigned int i = 0; i < maxnoofCellsUEMovingTrajectory_c; i++) {
        expectedUEMovingTrajectory_mb[i].next = NULL;
        expectedUEMovingTrajectory_mb[i].value.bit_mask = 0;
        expectedUEMovingTrajectory_mb[i].value.iE_Extensions=NULL;
        expectedUEMovingTrajectory_mb[i].value.nGRAN_CGI.choice = nR_CGI_chosen;
        expectedUEMovingTrajectory_mb[i].value.nGRAN_CGI.u.nR_CGI.bit_mask = 0;
        expectedUEMovingTrajectory_mb[i].value.nGRAN_CGI.u.nR_CGI.iE_Extensions=NULL;
        expectedUEMovingTrajectory_mb[i].value.nGRAN_CGI.u.nR_CGI.pLMNIdentity.length = 0;
        memset(expectedUEMovingTrajectory_mb[i].value.nGRAN_CGI.u.nR_CGI.pLMNIdentity.value,
            0xff,
            sizeof(expectedUEMovingTrajectory_mb[i].value.nGRAN_CGI.u.nR_CGI.pLMNIdentity.value));
        expectedUEMovingTrajectory_mb[i].value.nGRAN_CGI.u.nR_CGI.nRCellIdentity.length = 0;
        expectedUEMovingTrajectory_mb[i].value.nGRAN_CGI.u.nR_CGI.nRCellIdentity.value = &(nrCellmbId[i]);
        expectedUEMovingTrajectory_mb[i].value.timeStayedInCell = 0;
    }
    cnAssistedRANTuningPresent_m = false;

    //LocationReportingRequestType
    locationReportRequestType_mp = NULL;
    locationReportRequestType_m.bit_mask = 0;
    locationReportRequestType_m.iE_Extensions = NULL;
    locationReportRequestType_m.eventType = direct;
    locationReportRequestType_m.reportArea = cell;
    numofAreaOfInterestItem_m = 0;
    for (unsigned int i = 0; i < maxnoofAoI_c; i++) {
        areaOfInterestList_ma[i].next = NULL;
        areaOfInterestList_ma[i].value.bit_mask = 0;
        areaOfInterestList_ma[i].value.iE_Extensions = NULL;
        areaOfInterestList_ma[i].value.locationReportingReferenceID = 0;
        areaOfInterestList_ma[i].value.areaOfInterest.bit_mask = 0;
        areaOfInterestList_ma[i].value.areaOfInterest.iE_Extensions = NULL;
        areaOfInterestList_ma[i].value.areaOfInterest.areaOfInterestTAIList = NULL;
        for (unsigned int j = 0;j < maxnoofTAIinAoI_c ; j++) {
            areaOfInterestTAIList_ma[i][j].next = NULL;
            areaOfInterestTAIList_ma[i][j].value.bit_mask = 0;
            areaOfInterestTAIList_ma[i][j].value.iE_Extensions = NULL;
            areaOfInterestTAIList_ma[i][j].value.tAI.bit_mask = 0;
            areaOfInterestTAIList_ma[i][j].value.tAI.pLMNIdentity.length=0;
            memset(areaOfInterestTAIList_ma[i][j].value.tAI.pLMNIdentity.value, 0xff,
                   sizeof(areaOfInterestTAIList_ma[i][j].value.tAI.pLMNIdentity.value));
            areaOfInterestTAIList_ma[i][j].value.tAI.tAC.length=0;
            memset(areaOfInterestTAIList_ma[i][j].value.tAI.tAC.value, 0xff,
                   sizeof(areaOfInterestTAIList_ma[i][j].value.tAI.tAC.value));
            areaOfInterestTAIList_ma[i][j].value.tAI.iE_Extensions=NULL;
        }
        //now only support NR_CGI_Chosen not support eUTRA_CGI_chosen
        areaOfInterestList_ma[i].value.areaOfInterest.areaOfInterestCellList = NULL;
        for (unsigned int j = 0;j < maxnoofCellinAoI_c ; j++) {
            areaOfInterestCellList_ma[i][j].next = NULL;
            areaOfInterestCellList_ma[i][j].value.bit_mask = 0;
            areaOfInterestCellList_ma[i][j].value.iE_Extensions = NULL;
            areaOfInterestCellList_ma[i][j].value.nGRAN_CGI.choice = nR_CGI_chosen;
            areaOfInterestCellList_ma[i][j].value.nGRAN_CGI.u.nR_CGI.bit_mask=0;
            areaOfInterestCellList_ma[i][j].value.nGRAN_CGI.u.nR_CGI.pLMNIdentity.length=0;
            memset(areaOfInterestCellList_ma[i][j].value.nGRAN_CGI.u.nR_CGI.pLMNIdentity.value, 0xff,
                   sizeof(areaOfInterestCellList_ma[i][j].value.nGRAN_CGI.u.nR_CGI.pLMNIdentity.value));
            areaOfInterestCellList_ma[i][j].value.nGRAN_CGI.u.nR_CGI.nRCellIdentity.length=0;
            memset(nrCellId_ma[i][j],0xff,sizeof(nrCellId_ma[i][j]));
            areaOfInterestCellList_ma[i][j].value.nGRAN_CGI.u.nR_CGI.nRCellIdentity.value = nrCellId_ma[i][j];
            areaOfInterestCellList_ma[i][j].value.nGRAN_CGI.u.nR_CGI.iE_Extensions=NULL;
        }
        areaOfInterestList_ma[i].value.areaOfInterest.areaOfInterestRANNodeList = NULL;
        for (unsigned int j = 0;j < maxnoofRANNodeinAoI_c ; j++) {
            areaOfInterestRANNodeList_ma[i][j].next = NULL;
            areaOfInterestRANNodeList_ma[i][j].value.bit_mask = 0;
            areaOfInterestRANNodeList_ma[i][j].value.iE_Extensions = NULL;
            areaOfInterestRANNodeList_ma[i][j].value.globalRANNodeID.choice = globalGNB_ID_chosen;
            areaOfInterestRANNodeList_ma[i][j].value.globalRANNodeID.u.globalGNB_ID.bit_mask = 0;
            areaOfInterestRANNodeList_ma[i][j].value.globalRANNodeID.u.globalGNB_ID.iE_Extensions = NULL;
            areaOfInterestRANNodeList_ma[i][j].value.globalRANNodeID.u.globalGNB_ID.pLMNIdentity.length = 0;
            memset(areaOfInterestRANNodeList_ma[i][j].value.globalRANNodeID.u.globalGNB_ID.pLMNIdentity.value, 0xff,
                   sizeof(areaOfInterestRANNodeList_ma[i][j].value.globalRANNodeID.u.globalGNB_ID.pLMNIdentity.value));
            areaOfInterestRANNodeList_ma[i][j].value.globalRANNodeID.u.globalGNB_ID.gNB_ID.choice = gNB_ID_chosen;
            areaOfInterestRANNodeList_ma[i][j].value.globalRANNodeID.u.globalGNB_ID.gNB_ID.u.gNB_ID.length = 0;
            memset(gnbIdValue_ma[i][j], 0xff, sizeof(gnbIdValue_ma[i][j]));
            areaOfInterestRANNodeList_ma[i][j].value.globalRANNodeID.u.globalGNB_ID.gNB_ID.u.gNB_ID.value = gnbIdValue_ma[i][j];
        }
    }
    locationReportRequestType_m.locationReportingReferenceIDToBeCancelled = 0;
    locationReportRequestTypePresent_m = false;

}

//----------------------------------------------------------------------------
// This is a public API. Refer to InitialContextSetupRequestPdu.h
//----------------------------------------------------------------------------
unsigned int
InitialContextSetupRequestPdu::encodeOssToPer(NgapOssCtxt* ctxt)
{
    if (NULL == ctxt) {
        std::cout << "Invalid NgapOssCtxt!" << std::endl;
        return -1;
    }

    ProtocolIeType* nextIe_p = NULL;

    //LocationReportingRequestType
    ProtocolIeType locationReportRequestTypeIe;
    if(locationReportRequestTypePresent_m)
    {
        if (numofAreaOfInterestItem_m > maxnoofAoI || numofAreaOfInterestItem_m < 0) {
            return defaultFailureCode_c;
        }
        else
        {
            struct AreaOfInterestList_ *nextAreaOfInterest_p = NULL;
            for (int i = 0; i < numofAreaOfInterestItem_m; i++) {
                struct AreaOfInterestTAIList_ *nextAreaOfInterestTAI_p = NULL;
                for (unsigned int j = 0; j <numAreaOfInterestTAIList_m[i]; j++) {
                    areaOfInterestTAIList_ma[i][j].next = nextAreaOfInterestTAI_p;
                    nextAreaOfInterestTAI_p = &areaOfInterestTAIList_ma[i][j];
                }
                struct AreaOfInterestCellList_ *nextAreaOfInterestCell_p = NULL;
                for (unsigned int j = 0; j <numAreaOfInterestCellList_m[i]; j++) {
                    areaOfInterestCellList_ma[i][j].next = nextAreaOfInterestCell_p;
                    nextAreaOfInterestCell_p = &areaOfInterestCellList_ma[i][j];
                }
                struct AreaOfInterestRANNodeList_ *nextAreaOfInterestRANNode_p = NULL;
                for (unsigned int j = 0; j<numAreaOfInterestRANNodeList_m[i]; j++) {
                    areaOfInterestRANNodeList_ma[i][j].next = nextAreaOfInterestRANNode_p;
                    nextAreaOfInterestRANNode_p = &areaOfInterestRANNodeList_ma[i][j];
                }
                areaOfInterestList_ma[i].value.areaOfInterest.areaOfInterestTAIList = nextAreaOfInterestTAI_p;
                areaOfInterestList_ma[i].value.areaOfInterest.areaOfInterestCellList = nextAreaOfInterestCell_p;
                areaOfInterestList_ma[i].value.areaOfInterest.areaOfInterestRANNodeList = nextAreaOfInterestRANNode_p;

                areaOfInterestList_ma[i].next = nextAreaOfInterest_p;
                nextAreaOfInterest_p = &areaOfInterestList_ma[i];
            }
            locationReportRequestType_m.areaOfInterestList = nextAreaOfInterest_p;
            locationReportRequestTypeIe.value.value.pduNum=
                PDU_InitialContextSetupRequestIEs_Value_LocationReportingRequestType;
            locationReportRequestTypeIe.value.value.decoded.pdu_LocationReportingRequestType
                =&locationReportRequestType_m;
            locationReportRequestTypeIe.value.value.encoded.length = 0;
            locationReportRequestTypeIe.value.value.encoded.value = NULL;
        }
        locationReportRequestTypeIe.next=nextIe_p;
        locationReportRequestTypeIe.value.id=id_LocationReportingRequestType;
        locationReportRequestTypeIe.value.criticality=ignore;
        nextIe_p=&locationReportRequestTypeIe;
    }

    //CNAssistedRANTuning
    ProtocolIeType cnAssistedRANTuningIe;
    if(cnAssistedRANTuningPresent_m)
    {
        if(numOfExpectedUEMovingTrajectoryItem_mb > maxnoofCellsUEMovingTrajectory ||
              numOfExpectedUEMovingTrajectoryItem_mb < 0){
            return defaultFailureCode_c;
        }
        else
        {
            struct ExpectedUEMovingTrajectory_ *nextExpectedUEMovingTrajectory_p=NULL;
            for(int i=0;i<numOfExpectedUEMovingTrajectoryItem_mb;i++){
                expectedUEMovingTrajectory_mb[i].next=nextExpectedUEMovingTrajectory_p;
                nextExpectedUEMovingTrajectory_p=&expectedUEMovingTrajectory_mb[i];
            }
            cnAssistedRANTuning_m.expectedUEBehaviour.expectedUEMovingTrajectory=
                nextExpectedUEMovingTrajectory_p;

        cnAssistedRANTuningIe.value.value.pduNum=
            PDU_InitialContextSetupRequestIEs_Value_CNAssistedRANTuning;
        cnAssistedRANTuningIe.value.value.decoded.pdu_CNAssistedRANTuning
            =&cnAssistedRANTuning_m;
        cnAssistedRANTuningIe.value.value.encoded.length = 0;
        cnAssistedRANTuningIe.value.value.encoded.value = NULL;
        }
        cnAssistedRANTuningIe.next=nextIe_p;
        cnAssistedRANTuningIe.value.id=id_CNAssistedRANTuning;
        cnAssistedRANTuningIe.value.criticality=ignore;
        nextIe_p=&cnAssistedRANTuningIe;
    }

    //RedirectionVoiceFallback
    ProtocolIeType redirectionVoiceFallbackIe;
    if(redirectionVoiceFallbackPresent_m)
    {
        redirectionVoiceFallbackIe.next = nextIe_p;
        redirectionVoiceFallbackIe.value.id = id_RedirectionVoiceFallback;
        redirectionVoiceFallbackIe.value.criticality = ignore;
        redirectionVoiceFallbackIe.value.value.pduNum =
            PDU_InitialContextSetupRequestIEs_Value_RedirectionVoiceFallback;
        redirectionVoiceFallbackIe.value.value.decoded.pdu_RedirectionVoiceFallback =
            &redirectionVoiceFallback_m;
        redirectionVoiceFallbackIe.value.value.encoded.length = 0;
        redirectionVoiceFallbackIe.value.value.encoded.value = NULL;
        nextIe_p = &redirectionVoiceFallbackIe;
    }


    //UERadioCapabilityForPaging
    ProtocolIeType ueRadioCapabilityForPagingIe;
    if(ueRadioCapabilityForPagingPresent_m)
    {
        ueRadioCapabilityForPagingIe.next = nextIe_p;
        ueRadioCapabilityForPagingIe.value.id = id_UERadioCapabilityForPaging;
        ueRadioCapabilityForPagingIe.value.criticality = ignore;
        ueRadioCapabilityForPagingIe.value.value.pduNum = 
            PDU_InitialContextSetupRequestIEs_Value_UERadioCapabilityForPaging;
        ueRadioCapabilityForPagingIe.value.value.decoded.pdu_UERadioCapabilityForPaging = 
            &ueRadioCapabilityForPaging_m;
        ueRadioCapabilityForPagingIe.value.value.encoded.length = 0;
        ueRadioCapabilityForPagingIe.value.value.encoded.value = NULL;
        nextIe_p = &ueRadioCapabilityForPagingIe;
    }

    //RRCInactiveTransitionReportRequest
    ProtocolIeType rrcInactiveTransitionReportRequestIe;
    if(RRCReportPresent_m)
    {
        rrcInactiveTransitionReportRequestIe.next = nextIe_p;
        rrcInactiveTransitionReportRequestIe.value.id = id_RRCInactiveTransitionReportRequest;
        rrcInactiveTransitionReportRequestIe.value.criticality = ignore;
        rrcInactiveTransitionReportRequestIe.value.value.pduNum = 
            PDU_InitialContextSetupRequestIEs_Value_RRCInactiveTransitionReportRequest;
        rrcInactiveTransitionReportRequestIe.value.value.decoded.pdu_RRCInactiveTransitionReportRequest 
            = &rrcInactiveTransitionReportRequest_m;
        rrcInactiveTransitionReportRequestIe.value.value.encoded.length = 0;
        rrcInactiveTransitionReportRequestIe.value.value.encoded.value = NULL;
        nextIe_p = &rrcInactiveTransitionReportRequestIe;
    }

    //EmergencyFallbackIndicator
    ProtocolIeType emergencyFallbackIndicatorIe;
    if(emerFallbackPresent_m)
    {
        emergencyFallbackIndicatorIe.next = nextIe_p;
        emergencyFallbackIndicatorIe.value.id = id_EmergencyFallbackIndicator;
        emergencyFallbackIndicatorIe.value.criticality = reject;
        emergencyFallbackIndicatorIe.value.value.pduNum = 
            PDU_InitialContextSetupRequestIEs_Value_EmergencyFallbackIndicator;
        emergencyFallbackIndicatorIe.value.value.decoded.pdu_EmergencyFallbackIndicator 
            = &emergencyFallbackIndicator_m;
        emergencyFallbackIndicatorIe.value.value.encoded.length = 0;
        emergencyFallbackIndicatorIe.value.value.encoded.value = NULL;
        nextIe_p = &emergencyFallbackIndicatorIe;
    }

    //NAS PDU
    ProtocolIeType nasPduIe;
    if(nasPduPresent_m)
    {
        nasPduIe.next = nextIe_p;
        nasPduIe.value.id = id_NAS_PDU;
        nasPduIe.value.criticality = ignore;
        nasPduIe.value.value.pduNum = 
            PDU_InitialContextSetupRequestIEs_Value_NAS_PDU;
        nasPduIe.value.value.decoded.pdu_NAS_PDU = &nasPdu_m;
        nasPduIe.value.value.encoded.length = 0;
        nasPduIe.value.value.encoded.value = NULL;
        nextIe_p = &nasPduIe;
    }

    //MaskedIMEISV
    ProtocolIeType maskedIMEISVIe;
    if(maskedIMEISVPresent_m)
    {
        maskedIMEISVIe.next = nextIe_p;
        maskedIMEISVIe.value.id = id_MaskedIMEISV;
        maskedIMEISVIe.value.criticality = ignore;
        maskedIMEISVIe.value.value.pduNum = 
            PDU_InitialContextSetupRequestIEs_Value_MaskedIMEISV;
        maskedIMEISVIe.value.value.decoded.pdu_MaskedIMEISV = &maskedIMEISV_m;
        maskedIMEISVIe.value.value.encoded.length = 0;
        maskedIMEISVIe.value.value.encoded.value = NULL;
        nextIe_p = &maskedIMEISVIe;
    }

    //Index To RFSP
    ProtocolIeType indexToRFSPIe;
    if(indexToRFSPPresent_m)
    {
        indexToRFSPIe.next = nextIe_p;
        indexToRFSPIe.value.id = id_IndexToRFSP;
        indexToRFSPIe.value.criticality = ignore;
        indexToRFSPIe.value.value.pduNum = 
            PDU_InitialContextSetupRequestIEs_Value_IndexToRFSP;
        indexToRFSPIe.value.value.decoded.pdu_IndexToRFSP = &indexToRFSP_m;
        indexToRFSPIe.value.value.encoded.length = 0;
        indexToRFSPIe.value.value.encoded.value = NULL;
        nextIe_p = &indexToRFSPIe;
    }

    //UERadioCapability
    ProtocolIeType ueRadioCapabilityIe;
    if(ueRadioCapPresent_m)
    {
        ueRadioCapabilityIe.next = nextIe_p;
        ueRadioCapabilityIe.value.id = id_UERadioCapability;
        ueRadioCapabilityIe.value.criticality = ignore;
        ueRadioCapabilityIe.value.value.pduNum = 
            PDU_InitialContextSetupRequestIEs_Value_UERadioCapability;
        ueRadioCapabilityIe.value.value.decoded.pdu_UERadioCapability = &ueRadioCapability_m;
        ueRadioCapabilityIe.value.value.encoded.length = 0;
        ueRadioCapabilityIe.value.value.encoded.value = NULL;
        nextIe_p = &ueRadioCapabilityIe;
    }

    //TraceActivation
    ProtocolIeType traceActivationIe;
    if(traceActPresent_m)
    {
        traceActivationIe.next = nextIe_p;
        traceActivationIe.value.id = id_TraceActivation;
        traceActivationIe.value.criticality = ignore;
        traceActivationIe.value.value.pduNum = 
            PDU_InitialContextSetupRequestIEs_Value_TraceActivation;
        traceActivationIe.value.value.decoded.pdu_TraceActivation = &traceActivation_m;
        traceActivationIe.value.value.encoded.length = 0;
        traceActivationIe.value.value.encoded.value = NULL;
        nextIe_p = &traceActivationIe;
    }

    //SecurityKey
    ProtocolIeType securityKeyIe;
    securityKeyIe.next = nextIe_p;
    securityKeyIe.value.id = id_SecurityKey;
    securityKeyIe.value.criticality = reject;
    securityKeyIe.value.value.pduNum = 
        PDU_InitialContextSetupRequestIEs_Value_SecurityKey;
    securityKeyIe.value.value.decoded.pdu_SecurityKey = &securityKey_m;
    securityKeyIe.value.value.encoded.length = 0;
    securityKeyIe.value.value.encoded.value = NULL;
    nextIe_p = &securityKeyIe;

    //UeSecurityCapabilities
    ProtocolIeType ueSecurityCapabilitiesIe;
    ueSecurityCapabilitiesIe.next = nextIe_p;
    ueSecurityCapabilitiesIe.value.id = id_UESecurityCapabilities;
    ueSecurityCapabilitiesIe.value.criticality = reject;
    ueSecurityCapabilitiesIe.value.value.pduNum = 
        PDU_InitialContextSetupRequestIEs_Value_UESecurityCapabilities;
    ueSecurityCapabilitiesIe.value.value.decoded.pdu_UESecurityCapabilities = &ueSecurityCapabilities_m;
    ueSecurityCapabilitiesIe.value.value.encoded.length = 0;
    ueSecurityCapabilitiesIe.value.value.encoded.value = NULL;
    nextIe_p = &ueSecurityCapabilitiesIe;

    //Allowed NSSAI
    ProtocolIeType allowedNssaiIe;
    if (numOfAllowedNassiItem_m > maxnoofNSSAIs_c || numOfAllowedNassiItem_m <= 0) {
        return defaultFailureCode_c;
    } else {
        struct AllowedNSSAI_* nextAllowedNssai_p = NULL;
        for (int i = 0; i < numOfAllowedNassiItem_m; i++) {
            allowedNssai_ma[i].next = nextAllowedNssai_p;
            nextAllowedNssai_p = &allowedNssai_ma[i];
        }

        allowedNssai_m = nextAllowedNssai_p;
    }
    allowedNssaiIe.next = nextIe_p;
    allowedNssaiIe.value.id = id_AllowedNSSAI;
    allowedNssaiIe.value.criticality = reject;
    allowedNssaiIe.value.value.pduNum = 
        PDU_InitialContextSetupRequestIEs_Value_AllowedNSSAI;
    allowedNssaiIe.value.value.decoded.pdu_AllowedNSSAI = &allowedNssai_m;
    allowedNssaiIe.value.value.encoded.length = 0;
    allowedNssaiIe.value.value.encoded.value = NULL;
    nextIe_p = &allowedNssaiIe;

    //PDUSessionResourceSetupListCxtReq
    ProtocolIeType PDUSessionResourceSetupListCxtReqIe;
    if(numofPDUSessionResourceSetupListItemCxtReq_m > maxnoofPDUSessions ||
          numofPDUSessionResourceSetupListItemCxtReq_m < 0)
    {
        return defaultFailureCode_c;
    }
    else
    {
        struct PDUSessionResourceSetupListCxtReq_ *nextPDUSessionResourceSetupListCxtReq_p=NULL;
        for(unsigned int i=0;i<numofPDUSessionResourceSetupListItemCxtReq_m;i++)
        {
            PDUSessionResSetupListCxtReq_ma[i].next=nextPDUSessionResourceSetupListCxtReq_p;
            nextPDUSessionResourceSetupListCxtReq_p=&(PDUSessionResSetupListCxtReq_ma[i]);
        }

        if(nextPDUSessionResourceSetupListCxtReq_p != NULL)
        {
            pDUSessionResourceSetupListCxtReq_m=nextPDUSessionResourceSetupListCxtReq_p;

            PDUSessionResourceSetupListCxtReqIe.value.value.pduNum=
                PDU_InitialContextSetupRequestIEs_Value_PDUSessionResourceSetupListCxtReq;
            PDUSessionResourceSetupListCxtReqIe.value.value.decoded.pdu_PDUSessionResourceSetupListCxtReq
                =&pDUSessionResourceSetupListCxtReq_m;
            PDUSessionResourceSetupListCxtReqIe.value.value.encoded.length = 0;
            PDUSessionResourceSetupListCxtReqIe.value.value.encoded.value = NULL;
            PDUSessionResourceSetupListCxtReqIe.next=nextIe_p;
            PDUSessionResourceSetupListCxtReqIe.value.id=id_PDUSessionResourceSetupListCxtReq;
            PDUSessionResourceSetupListCxtReqIe.value.criticality=reject;
            nextIe_p=&PDUSessionResourceSetupListCxtReqIe;
        }
    }

    //GUAMI
    ProtocolIeType gUAMIIe;
    gUAMIIe.next = nextIe_p;
    gUAMIIe.value.id = id_GUAMI;
    gUAMIIe.value.criticality = reject;
    gUAMIIe.value.value.pduNum = PDU_InitialContextSetupRequestIEs_Value_GUAMI;
    gUAMIIe.value.value.decoded.pdu_GUAMI = &guami_m;
    gUAMIIe.value.value.encoded.length = 0;
    gUAMIIe.value.value.encoded.value = NULL;
    nextIe_p = &gUAMIIe;

    //CoreNetworkAssistanceInformation
    ProtocolIeType coreNetworkAssistanceInformationIe;
    if(coreNetworkPresent_m)
    {
        if(numOfTAIListForInactiveItem_m > maxnoofTAIforInactive ||
              numOfTAIListForInactiveItem_m < 0){
            return defaultFailureCode_c;
        }
        else if(numOfExpectedUEMovingTrajectoryItem_m > maxnoofCellsUEMovingTrajectory ||
              numOfExpectedUEMovingTrajectoryItem_m < 0){
            return defaultFailureCode_c;
        }
        else
        {
            struct TAIListForInactive_ *nextTAIListForInactive_p=NULL;
            for(int i=0;i<numOfTAIListForInactiveItem_m;i++){
                taiListForInactive_ma[i].next=nextTAIListForInactive_p;
                nextTAIListForInactive_p=&taiListForInactive_ma[i];
            }
            coreNetworkAssistanceInfo_m.tAIListForInactive=nextTAIListForInactive_p;

            struct ExpectedUEMovingTrajectory_ *nextExpectedUEMovingTrajectory_p=NULL;
            for(int i=0;i<numOfExpectedUEMovingTrajectoryItem_m;i++){
                expectedUEMovingTrajectory_ma[i].next=nextExpectedUEMovingTrajectory_p;
                nextExpectedUEMovingTrajectory_p=&expectedUEMovingTrajectory_ma[i];
            }
            coreNetworkAssistanceInfo_m.expectedUEBehaviour.expectedUEMovingTrajectory=
                nextExpectedUEMovingTrajectory_p;

        coreNetworkAssistanceInformationIe.value.value.pduNum=
            PDU_InitialContextSetupRequestIEs_Value_CoreNetworkAssistanceInformationForInactive;
        coreNetworkAssistanceInformationIe.value.value.decoded.pdu_CoreNetworkAssistanceInformationForInactive
            =&coreNetworkAssistanceInfo_m;
        coreNetworkAssistanceInformationIe.value.value.encoded.length = 0;
        coreNetworkAssistanceInformationIe.value.value.encoded.value = NULL;
        }
        coreNetworkAssistanceInformationIe.next=nextIe_p;
        coreNetworkAssistanceInformationIe.value.id=id_CoreNetworkAssistanceInformationForInactive;
        coreNetworkAssistanceInformationIe.value.criticality=ignore;
        nextIe_p=&coreNetworkAssistanceInformationIe;
    }

    //UE Aggregate Maximum Bit Rate
    ProtocolIeType ueAggregateMaximumBitRateIe;
    if(ueAMBRPresent_m)
    {
        ueAggregateMaximumBitRateIe.next = nextIe_p;
        ueAggregateMaximumBitRateIe.value.id = id_UEAggregateMaximumBitRate;
        ueAggregateMaximumBitRateIe.value.criticality = reject;
        ueAggregateMaximumBitRateIe.value.value.pduNum = 
            PDU_InitialContextSetupRequestIEs_Value_UEAggregateMaximumBitRate;
        ueAggregateMaximumBitRateIe.value.value.decoded.pdu_UEAggregateMaximumBitRate = 
            &ueAggregateMaximumBitRate_m;
        ueAggregateMaximumBitRateIe.value.value.encoded.length = 0;
        ueAggregateMaximumBitRateIe.value.value.encoded.value = NULL;
        nextIe_p = &ueAggregateMaximumBitRateIe;
    }

    //OldAMF
    ProtocolIeType amfNameIe;
    if(amfNamePresent_m)
    {
        amfNameIe.next = nextIe_p;
        amfNameIe.value.id = id_OldAMF;
        amfNameIe.value.criticality = reject;
        amfNameIe.value.value.pduNum = 
            PDU_InitialContextSetupRequestIEs_Value_AMFName;
        amfNameIe.value.value.decoded.pdu_AMFName = &amfName_m;
        amfNameIe.value.value.encoded.length = 0;
        amfNameIe.value.value.encoded.value = NULL;
        nextIe_p = &amfNameIe;
    }

    //RAN_UE_NGAP_ID
    ProtocolIeType ranUeNgapIdIe;
    ranUeNgapIdIe.next = nextIe_p;
    ranUeNgapIdIe.value.id = id_RAN_UE_NGAP_ID;
    ranUeNgapIdIe.value.criticality = reject;
    ranUeNgapIdIe.value.value.pduNum = 
        PDU_InitialContextSetupRequestIEs_Value_RAN_UE_NGAP_ID;
    ranUeNgapIdIe.value.value.decoded.pdu_RAN_UE_NGAP_ID = &ranUeNgapId_m;
    ranUeNgapIdIe.value.value.encoded.length = 0;
    ranUeNgapIdIe.value.value.encoded.value = NULL;
    nextIe_p = &ranUeNgapIdIe;

    //AMF_UE_NGAP_ID
    ProtocolIeType amfUeNgapIdIe;
    amfUeNgapIdIe.next = nextIe_p;
    amfUeNgapIdIe.value.id = id_AMF_UE_NGAP_ID;
    amfUeNgapIdIe.value.criticality = reject;
    amfUeNgapIdIe.value.value.pduNum = 
        PDU_InitialContextSetupRequestIEs_Value_AMF_UE_NGAP_ID;
    amfUeNgapIdIe.value.value.decoded.pdu_AMF_UE_NGAP_ID = &amfUeNgapId_m;
    amfUeNgapIdIe.value.value.encoded.length = 0;
    amfUeNgapIdIe.value.value.encoded.value = NULL;
    nextIe_p = &amfUeNgapIdIe;

    //message
    InitialContextSetupRequest msg;
    msg.protocolIEs = nextIe_p;

    NGAP_PDU ngApPdu;

    ngApPdu.choice = initiatingMessage_chosen;
    ngApPdu.u.initiatingMessage.procedureCode = id_InitialContextSetup;
    ngApPdu.u.initiatingMessage.criticality = reject;
    ngApPdu.u.initiatingMessage.value.pduNum = 
        PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_InitialContextSetupRequest;
    ngApPdu.u.initiatingMessage.value.decoded.pdu_InitialContextSetupRequest = &msg;
    ngApPdu.u.initiatingMessage.value.encoded.length = 0;
    ngApPdu.u.initiatingMessage.value.encoded.value = NULL;

    //Convert Structures to PER
    OssBuf outputBuf;
    unsigned int result;

    ctxt->reset_v();
    outputBuf = ctxt->getPrimaryCodecBuf();

    result = ctxt->encodeOssToPer(NGAP_PDU_PDU, &ngApPdu, &outputBuf);

    if (result == encodeSuccessCode_c) {
        perBufRef_m = outputBuf;
    }

    return (result);
}

//----------------------------------------------------------------------------
// This is a public API. Refer to InitialContextSetupRequestPdu.h
//----------------------------------------------------------------------------
void InitialContextSetupRequestPdu::getDecodedOssDataPtrs_v()
{
    ProtocolIeType* currIe_p;
    reset_v();

    ngapPdu_mp = reinterpret_cast<NGAP_PDU*>(ossBufRef_m.value);
    if (ngapPdu_mp == NULL) {
        return;
    }

    InitialContextSetupRequest_mp = 
        ngapPdu_mp->u.initiatingMessage.value.decoded.pdu_InitialContextSetupRequest;

    if (InitialContextSetupRequest_mp == NULL) {
        return;
    }

    currIe_p = InitialContextSetupRequest_mp->protocolIEs;

    while (currIe_p != NULL) {
        if (currIe_p->value.id == id_RAN_UE_NGAP_ID
            && currIe_p->value.value.decoded.pdu_RAN_UE_NGAP_ID != NULL) {
            ranUeNgapId_mp = currIe_p->value.value.decoded.pdu_RAN_UE_NGAP_ID;
        } else if (currIe_p->value.id == id_AMF_UE_NGAP_ID
            && currIe_p->value.value.decoded.pdu_AMF_UE_NGAP_ID != NULL) {
            amfUeNgapId_mp = currIe_p->value.value.decoded.pdu_AMF_UE_NGAP_ID;
        } else if (currIe_p->value.id == id_NAS_PDU
            && currIe_p->value.value.decoded.pdu_NAS_PDU != NULL) {
            nasPdu_mp = currIe_p->value.value.decoded.pdu_NAS_PDU;
            nasPduPresent_m=true;
        } else if (currIe_p->value.id == id_OldAMF
            && currIe_p->value.value.decoded.pdu_AMFName != NULL) {
            amfName_mp = currIe_p->value.value.decoded.pdu_AMFName;
            amfNamePresent_m=true;
        } else if (currIe_p->value.id == id_UEAggregateMaximumBitRate
            && currIe_p->value.value.decoded.pdu_UEAggregateMaximumBitRate != NULL) {
            ueAggregateMaximumBitRate_mp = 
                currIe_p->value.value.decoded.pdu_UEAggregateMaximumBitRate;
            ueAMBRPresent_m=true;
        } else if (currIe_p->value.id == id_GUAMI
            && currIe_p->value.value.decoded.pdu_GUAMI != NULL) {
            guami_mp = currIe_p->value.value.decoded.pdu_GUAMI;
        } else if (currIe_p->value.id == id_AllowedNSSAI
            && currIe_p->value.value.decoded.pdu_AllowedNSSAI != NULL) {
            allowedNssai_mp = currIe_p->value.value.decoded.pdu_AllowedNSSAI;
            AllowedNSSAI curr_p = *allowedNssai_mp;
            while (curr_p != NULL) {
                numOfAllowedNassiItem_m++;
                curr_p = curr_p->next;
            }
        }

        else if (currIe_p->value.id == id_CoreNetworkAssistanceInformationForInactive
            && currIe_p->value.value.decoded.pdu_CoreNetworkAssistanceInformationForInactive != NULL)
        {
            coreNetworkAssistanceInfo_mp =
                      currIe_p->value.value.decoded.pdu_CoreNetworkAssistanceInformationForInactive;
            
            TAIListForInactive taiList_p=NULL;
            taiList_p=coreNetworkAssistanceInfo_mp->tAIListForInactive;
            while(taiList_p!=NULL)
            {
                numOfTAIListForInactiveItem_m++;
                taiList_p=taiList_p->next;
            }

            ExpectedUEMovingTrajectory expUEmovingTra_p=NULL;
            expUEmovingTra_p=
                coreNetworkAssistanceInfo_mp->expectedUEBehaviour.expectedUEMovingTrajectory;
            while(expUEmovingTra_p!=NULL)
            {
                numOfExpectedUEMovingTrajectoryItem_m++;
                expUEmovingTra_p=expUEmovingTra_p->next;
            }
            coreNetworkPresent_m=true;
        }

        else if (currIe_p->value.id == id_PDUSessionResourceSetupListCxtReq
            && currIe_p->value.value.decoded.pdu_PDUSessionResourceSetupListCxtReq != NULL)
        {
            pDUSessionResourceSetupListCxtReq_mp =
                      currIe_p->value.value.decoded.pdu_PDUSessionResourceSetupListCxtReq;

            PDUSessionResourceSetupListCxtReq pduSessResSetup_p=NULL;

            pduSessResSetup_p=*pDUSessionResourceSetupListCxtReq_mp;
            while(pduSessResSetup_p!=NULL)
            {
                numofPDUSessionResourceSetupListItemCxtReq_m ++;
                pduSessResSetup_p=pduSessResSetup_p->next;
            }
        }

        else if (currIe_p->value.id == id_UESecurityCapabilities
            && currIe_p->value.value.decoded.pdu_UESecurityCapabilities != NULL) {
            ueSecurityCapabilities_mp = currIe_p->value.value.decoded.pdu_UESecurityCapabilities;
        } else if (currIe_p->value.id == id_SecurityKey
            && currIe_p->value.value.decoded.pdu_SecurityKey != NULL) {
            securityKey_mp = currIe_p->value.value.decoded.pdu_SecurityKey;
        } else if (currIe_p->value.id == id_TraceActivation
            && currIe_p->value.value.decoded.pdu_TraceActivation != NULL) {
            traceActivation_mp = currIe_p->value.value.decoded.pdu_TraceActivation;
            traceActPresent_m=true;
        } else if (currIe_p->value.id == id_UERadioCapability
            && currIe_p->value.value.decoded.pdu_UERadioCapability != NULL) {
            ueRadioCapability_mp = currIe_p->value.value.decoded.pdu_UERadioCapability;
            ueRadioCapPresent_m=true;
        } else if (currIe_p->value.id == id_IndexToRFSP
            && currIe_p->value.value.decoded.pdu_IndexToRFSP != NULL) {
            indexToRFSP_mp = currIe_p->value.value.decoded.pdu_IndexToRFSP;
            indexToRFSPPresent_m=true;
        } else if (currIe_p->value.id == id_MaskedIMEISV
            && currIe_p->value.value.decoded.pdu_MaskedIMEISV != NULL) {
            maskedIMEISV_mp = currIe_p->value.value.decoded.pdu_MaskedIMEISV;
            maskedIMEISVPresent_m=true;
        } else if (currIe_p->value.id == id_EmergencyFallbackIndicator
            && currIe_p->value.value.decoded.pdu_EmergencyFallbackIndicator != NULL) {
            emergencyFallbackIndicator_mp = currIe_p->value.value.decoded.pdu_EmergencyFallbackIndicator;
            emerFallbackPresent_m=true;
        } else if (currIe_p->value.id == id_RRCInactiveTransitionReportRequest
            && currIe_p->value.value.decoded.pdu_RRCInactiveTransitionReportRequest != NULL) {
            rrcInactiveTransitionReportRequest_mp = 
                currIe_p->value.value.decoded.pdu_RRCInactiveTransitionReportRequest;
            RRCReportPresent_m=true;
        } else if (currIe_p->value.id == id_UERadioCapabilityForPaging
            && currIe_p->value.value.decoded.pdu_UERadioCapabilityForPaging != NULL) {
            ueRadioCapabilityForPaging_mp = currIe_p->value.value.decoded.pdu_UERadioCapabilityForPaging;
            ueRadioCapabilityForPagingPresent_m=true;
        } else if (currIe_p->value.id == id_RedirectionVoiceFallback
            && currIe_p->value.value.decoded.pdu_RedirectionVoiceFallback != NULL) {
            redirectionVoiceFallback_mp = currIe_p->value.value.decoded.pdu_RedirectionVoiceFallback;
            redirectionVoiceFallbackPresent_m=true;
        } else if (currIe_p->value.id == id_CNAssistedRANTuning
            && currIe_p->value.value.decoded.pdu_CNAssistedRANTuning != NULL)
        {
            cnAssistedRANTuning_mp =
                      currIe_p->value.value.decoded.pdu_CNAssistedRANTuning;

            ExpectedUEMovingTrajectory expUEmovingTra_p=NULL;
            expUEmovingTra_p=
                cnAssistedRANTuning_mp->expectedUEBehaviour.expectedUEMovingTrajectory;
            while(expUEmovingTra_p!=NULL)
            {
                numOfExpectedUEMovingTrajectoryItem_mb++;
                expUEmovingTra_p=expUEmovingTra_p->next;
            }
            cnAssistedRANTuningPresent_m=true;
        } else if (currIe_p->value.id == id_LocationReportingRequestType
            && currIe_p->value.value.decoded.pdu_LocationReportingRequestType != NULL)
        {
            locationReportRequestType_mp =
                currIe_p->value.value.decoded.pdu_LocationReportingRequestType;

            AreaOfInterestList areaOfInterestList_p = NULL;
            AreaOfInterestTAIList areaOfInterestTAIList_p = NULL;
            AreaOfInterestCellList areaOfInterestCellList_p = NULL;
            AreaOfInterestRANNodeList areaOfInterestRANNodeList_p = NULL;
            areaOfInterestList_p = locationReportRequestType_mp->areaOfInterestList;
            while (areaOfInterestList_p != NULL) {
                areaOfInterestTAIList_p = areaOfInterestList_p ->value.areaOfInterest.areaOfInterestTAIList;
                while (areaOfInterestTAIList_p != NULL) {
                    numAreaOfInterestTAIList_m[numofAreaOfInterestItem_m]++;
                    areaOfInterestTAIList_p = areaOfInterestTAIList_p->next;
                }
                areaOfInterestCellList_p = areaOfInterestList_p ->value.areaOfInterest.areaOfInterestCellList;
                while (areaOfInterestCellList_p != NULL) {
                    numAreaOfInterestCellList_m[numofAreaOfInterestItem_m]++;
                    areaOfInterestCellList_p = areaOfInterestCellList_p->next;
                }
                areaOfInterestRANNodeList_p = areaOfInterestList_p ->value.areaOfInterest.areaOfInterestRANNodeList;
                while (areaOfInterestRANNodeList_p != NULL) {
                    numAreaOfInterestRANNodeList_m[numofAreaOfInterestItem_m]++;
                    areaOfInterestRANNodeList_p = areaOfInterestRANNodeList_p->next;
                }
                numofAreaOfInterestItem_m++;
                areaOfInterestList_p = areaOfInterestList_p->next;
            }
            locationReportRequestTypePresent_m = true;
        }

        currIe_p = currIe_p->next;
    }
}

//----------------------------------------------------------------------------
// This is a public API. Refer to InitialContextSetupRequestPdu.h
//----------------------------------------------------------------------------
InitialContextSetupRequest*
InitialContextSetupRequestPdu::getInitialContextSetupRequestDecodedPtr() const
{
    return InitialContextSetupRequest_mp;
}

//----------------------------------------------------------------------------
// This is a public API. Refer to InitialContextSetupRequestPdu.h
//----------------------------------------------------------------------------
AMFName*
InitialContextSetupRequestPdu::getAmfNameDecodedPtr() const
{
    return amfName_mp;
}

AMFName*
InitialContextSetupRequestPdu::getAmfNameEncodePtr()
{
    amfNamePresent_m=true;
    return &amfName_m;
}

bool 
InitialContextSetupRequestPdu::isAmfNamePresent()
{
    return amfNamePresent_m;
}

//----------------------------------------------------------------------------
// This is a public API. Refer to InitialContextSetupRequestPdu.h
//----------------------------------------------------------------------------
IndexToRFSP*
InitialContextSetupRequestPdu::getIndexToRFSPDecodedPtr() const
{
    return indexToRFSP_mp;
}

IndexToRFSP*
InitialContextSetupRequestPdu::getIndexToRFSPEncodePtr()
{
    indexToRFSPPresent_m=true;
    return &indexToRFSP_m;
}

bool 
InitialContextSetupRequestPdu::isIndexToRFSPPresent()
{
    return indexToRFSPPresent_m;
}

//----------------------------------------------------------------------------
// This is a public API. Refer to InitialContextSetupRequestPdu.h
//----------------------------------------------------------------------------
UEAggregateMaximumBitRate*
InitialContextSetupRequestPdu::getUeAggregateMaximumBitRateDecodedPtr() const
{
    return ueAggregateMaximumBitRate_mp;
}

UEAggregateMaximumBitRate*
InitialContextSetupRequestPdu::getUeAggregateMaximumBitRateEncodePtr()
{
    ueAMBRPresent_m=true;
    return &ueAggregateMaximumBitRate_m;
}

bool 
InitialContextSetupRequestPdu::isUeAMBRPresent()
{
    return ueAMBRPresent_m;
}


//----------------------------------------------------------------------------
// This is a public API. Refer to InitialContextSetupRequestPdu.h
//----------------------------------------------------------------------------
RAN_UE_NGAP_ID*
InitialContextSetupRequestPdu::getRanUeNgapIdDecodedPtr() const
{
    return ranUeNgapId_mp;
}

RAN_UE_NGAP_ID*
InitialContextSetupRequestPdu::getRanUeNgapIdEncodePtr()
{
    return &ranUeNgapId_m;
}

//----------------------------------------------------------------------------
// This is a public API. Refer to InitialContextSetupRequestPdu.h
//----------------------------------------------------------------------------
AMF_UE_NGAP_ID*
InitialContextSetupRequestPdu::getAmfUeNgapIdDecodedPtr() const
{
    return amfUeNgapId_mp;
}

AMF_UE_NGAP_ID*
InitialContextSetupRequestPdu::getAmfUeNgapIdEncodePtr()
{
    return &amfUeNgapId_m;
}

//----------------------------------------------------------------------------
// This is a public API. Refer to InitialContextSetupRequestPdu.h
//----------------------------------------------------------------------------
NAS_PDU*
InitialContextSetupRequestPdu::getNasPduDecodedPtr() const
{
    return nasPdu_mp;
}

NAS_PDU*
InitialContextSetupRequestPdu::getNasPduEncodePtr()
{
    nasPduPresent_m=true;
    return &nasPdu_m;
}

bool 
InitialContextSetupRequestPdu::isNasPduPresent()
{
    return nasPduPresent_m;
}

//----------------------------------------------------------------------------
// This is a public API. Refer to InitialContextSetupRequestPdu.h
//----------------------------------------------------------------------------
CoreNetworkAssistanceInformationForInactive*
InitialContextSetupRequestPdu::getCoreNetworkAssistanceInformationDecodedPtr() const
{
    return coreNetworkAssistanceInfo_mp;
}

CoreNetworkAssistanceInformationForInactive*
InitialContextSetupRequestPdu::getCoreNetworkAssistanceInformationEncodePtr()
{
    coreNetworkPresent_m=true;
    return &coreNetworkAssistanceInfo_m;
}

bool 
InitialContextSetupRequestPdu::isCoreNetworkPresent()
{
    return coreNetworkPresent_m;
}

TAIListForInactiveItem*
InitialContextSetupRequestPdu::getTAIListForInactiveItemDecodedPtr(unsigned int index) const
{
    TAIListForInactiveItem* item_p = NULL;
    TAIListForInactive TAIListForInactive_a = NULL;
    if (coreNetworkAssistanceInfo_mp->tAIListForInactive != NULL) {
        TAIListForInactive_a = coreNetworkAssistanceInfo_mp->tAIListForInactive;
    }

    if (TAIListForInactive_a != NULL
        && (index >= 0 && index < numOfTAIListForInactiveItem_m)) {
        unsigned int i = 0;
        for (; i < index && TAIListForInactive_a != NULL; i++) {
            TAIListForInactive_a = TAIListForInactive_a->next;
        }
        if (TAIListForInactive_a != NULL && i == index) {
            item_p = &(TAIListForInactive_a->value);
        }
    }

    return item_p;
}

TAIListForInactiveItem*
InitialContextSetupRequestPdu::getTAIListForInactiveItemEncodePtr(unsigned int index)
{
    if (index >= 0 && index < maxnoofTAIforInactive) {
        return &(taiListForInactive_ma[index].value);
    } else {
        return NULL;
    }
}

void InitialContextSetupRequestPdu::setNumOfTAIListForInactiveItem_v(unsigned int num)
{
    numOfTAIListForInactiveItem_m = num;
}

unsigned int
InitialContextSetupRequestPdu::getNumOfTAIListForInactiveItem() const
{
    return numOfTAIListForInactiveItem_m;
}

ExpectedUEMovingTrajectoryItem*
InitialContextSetupRequestPdu::getExpectedUEMovingTrajectoryItemDecodedPtr(unsigned int index) const
{
    ExpectedUEMovingTrajectoryItem* item_p = NULL;
    ExpectedUEMovingTrajectory ExpectedUEMovingTrajectory_a = NULL;
    if (coreNetworkAssistanceInfo_mp->expectedUEBehaviour.expectedUEMovingTrajectory != NULL) {
        ExpectedUEMovingTrajectory_a = 
            coreNetworkAssistanceInfo_mp->expectedUEBehaviour.expectedUEMovingTrajectory;
    }

    if (ExpectedUEMovingTrajectory_a != NULL
        && (index >= 0 && index < numOfExpectedUEMovingTrajectoryItem_m)) {
        unsigned int i = 0;
        for (; i < index && ExpectedUEMovingTrajectory_a != NULL; i++) {
            ExpectedUEMovingTrajectory_a = ExpectedUEMovingTrajectory_a->next;
        }
        if (ExpectedUEMovingTrajectory_a != NULL && i == index) {
            item_p = &(ExpectedUEMovingTrajectory_a->value);
        }
    }

    return item_p;
}

ExpectedUEMovingTrajectoryItem*
InitialContextSetupRequestPdu::getExpectedUEMovingTrajectoryItemEncodePtr(unsigned int index)
{
    if (index >= 0 && index < maxnoofCellsUEMovingTrajectory) {
        return &(expectedUEMovingTrajectory_ma[index].value);
    } else {
        return NULL;
    }
}

void InitialContextSetupRequestPdu::setNumOfExpectedUEMovingTrajectoryItem_v(unsigned int num)
{
    numOfExpectedUEMovingTrajectoryItem_m = num;
}

unsigned int
InitialContextSetupRequestPdu::getNumOfExpectedUEMovingTrajectoryItem() const
{
    return numOfExpectedUEMovingTrajectoryItem_m;
}

//----------------------------------------------------------------------------
// This is a public API. Refer to InitialContextSetupRequestPdu.h
//----------------------------------------------------------------------------
//PDUSessionResourceSetupListCxtReq
PDUSessionResourceSetupListCxtReq*
InitialContextSetupRequestPdu::getPDUSessionResourceSetupListCxtReqDecodedPtr() const
{
    return pDUSessionResourceSetupListCxtReq_mp;
}

PDUSessionResourceSetupListCxtReq*
InitialContextSetupRequestPdu::getPDUSessionResourceSetupListCxtReqEncodePtr()
{
    return &pDUSessionResourceSetupListCxtReq_m;
}

PDUSessionResourceSetupItemCxtReq*
InitialContextSetupRequestPdu::getPDUSessionResourceSetupItemCxtReqDecodedPtr(unsigned int index) const
{
    PDUSessionResourceSetupItemCxtReq* item_p = NULL;

    PDUSessionResourceSetupListCxtReq pduSessResSetupList_p = NULL;

    if (pDUSessionResourceSetupListCxtReq_mp != NULL) {
        pduSessResSetupList_p = *(pDUSessionResourceSetupListCxtReq_mp);
    }

    if (pduSessResSetupList_p != NULL
        && (index >= 0 && index < numofPDUSessionResourceSetupListItemCxtReq_m)) {
        unsigned int i = 0;
        for (; i < index && pduSessResSetupList_p != NULL; i++) {
            pduSessResSetupList_p = pduSessResSetupList_p->next;
        }
        if (pduSessResSetupList_p != NULL && i == index) {
            item_p = &(pduSessResSetupList_p->value);
        }
    }

    return item_p;
}

PDUSessionResourceSetupItemCxtReq*
InitialContextSetupRequestPdu::getPDUSessionResourceSetupItemCxtReqEncodePtr(unsigned int index)
{
    if (index >= 0 && index < maxnoofPDUSessions) {
        return &(PDUSessionResSetupListCxtReq_ma[index].value);
    } else {
        return NULL;
    }
}
void InitialContextSetupRequestPdu::setNumofPDUSessionResourceSetupListItemCxtReq_v(unsigned int num)
{
    numofPDUSessionResourceSetupListItemCxtReq_m = num;
}

unsigned int
InitialContextSetupRequestPdu::getNumofPDUSessionResourceSetupListItemCxtReq() const
{
    return numofPDUSessionResourceSetupListItemCxtReq_m;
}

//----------------------------------------------------------------------------
// This is a public API. Refer to InitialContextSetupRequestPdu.h
//----------------------------------------------------------------------------
MobilityRestrictionList*
InitialContextSetupRequestPdu::getMobilityRestrictionListDecodedPtr() const
{
    return mobilityRestrictionList_mp;
}

MobilityRestrictionList*
InitialContextSetupRequestPdu::getMobilityRestrictionListEncodePtr()
{
    return &mobilityRestrictionList_m;
}

PLMNIdentity*
InitialContextSetupRequestPdu::getPLMNIdentityDecodedPtr(unsigned int index) const
{
    PLMNIdentity* item_p = NULL;
    EquivalentPLMNs EquivalentPLMNs_a = NULL;
    if (mobilityRestrictionList_mp->equivalentPLMNs != NULL) {
        EquivalentPLMNs_a = mobilityRestrictionList_mp->equivalentPLMNs;
    }
    if (EquivalentPLMNs_a != NULL && (index >= 0 && index < numofEquivalentPLMNs_m)) {
        unsigned int i = 0;
        for (; i < index && EquivalentPLMNs_a != NULL; i++) {
            EquivalentPLMNs_a = EquivalentPLMNs_a->next;
        }
        if (EquivalentPLMNs_a != NULL && i == index) {
            item_p = &(EquivalentPLMNs_a->value);
        }
    }

    return item_p;
}

PLMNIdentity*
InitialContextSetupRequestPdu::getPLMNIdentityEncodePtr(unsigned int index)
{
    if (index >= 0 && index < maxnoofEPLMNs) {
        return &(equivalentPLMNs_ma[index].value);
    } else {
        return NULL;
    }
}

void InitialContextSetupRequestPdu::setNumofPLMNIdentity_v(unsigned int num)
{
    numofEquivalentPLMNs_m = num;
}

unsigned int
InitialContextSetupRequestPdu::getNumofPLMNIdentity() const
{
    return numofEquivalentPLMNs_m;
}

RATRestrictions_Item*
InitialContextSetupRequestPdu::getRATRestrictionsItemDecodedPtr(unsigned int index) const
{
    RATRestrictions_Item* item_p = NULL;
    RATRestrictions RATRestrictions_a = NULL;
    if (mobilityRestrictionList_mp->rATRestrictions != NULL) {
        RATRestrictions_a = mobilityRestrictionList_mp->rATRestrictions;
    }
    if (RATRestrictions_a != NULL && (index >= 0 && index < numofRATRestrictions_m)) {
        unsigned int i = 0;
        for (; i < index && RATRestrictions_a != NULL; i++) {
            RATRestrictions_a = RATRestrictions_a->next;
        }
        if (RATRestrictions_a != NULL && i == index) {
            item_p = &(RATRestrictions_a->value);
        }
    }

    return item_p;
}

RATRestrictions_Item*
InitialContextSetupRequestPdu::getRATRestrictionsItemEncodePtr(unsigned int index)
{
    if (index >= 0 && index < maxnoofEPLMNsPlusOne) {
        return &(rATRestrictions_ma[index].value);
    } else {
        return NULL;
    }
}

void InitialContextSetupRequestPdu::setNumofRATRestrictionsItem_v(unsigned int num)
{
    numofRATRestrictions_m = num;
}

unsigned int
InitialContextSetupRequestPdu::getNumofRATRestrictionsItem() const
{
    return numofRATRestrictions_m;
}

ForbiddenAreaInformation_Item*
InitialContextSetupRequestPdu::getForbiddenAreaInformationItemDecodedPtr(unsigned int index) const
{
    ForbiddenAreaInformation_Item* item_p = NULL;
    ForbiddenAreaInformation ForbiddenAreaInformation_a = NULL;
    if (mobilityRestrictionList_mp->forbiddenAreaInformation != NULL) {
        ForbiddenAreaInformation_a = mobilityRestrictionList_mp->forbiddenAreaInformation;
    }
    if (ForbiddenAreaInformation_a != NULL && (index >= 0 && index < numofForbiddenAreaInformation_m)) {
        unsigned int i = 0;
        for (; i < index && ForbiddenAreaInformation_a != NULL; i++) {
            ForbiddenAreaInformation_a = ForbiddenAreaInformation_a->next;
        }
        if (ForbiddenAreaInformation_a != NULL && i == index) {
            item_p = &(ForbiddenAreaInformation_a->value);
        }
    }

    return item_p;
}

ForbiddenAreaInformation_Item*
InitialContextSetupRequestPdu::getForbiddenAreaInformationItemEncodePtr(unsigned int index)
{
    if (index >= 0 && index < maxnoofEPLMNsPlusOne) {
        return &(forbiddenAreaInformation_ma[index].value);
    } else {
        return NULL;
    }
}

void InitialContextSetupRequestPdu::setNumofForbiddenAreaInformationItem_v(unsigned int num)
{
    numofForbiddenAreaInformation_m = num;
}

unsigned int
InitialContextSetupRequestPdu::getNumofForbiddenAreaInformationItem() const
{
    return numofForbiddenAreaInformation_m;
}

ServiceAreaInformation_Item*
InitialContextSetupRequestPdu::getServiceAreaInformationItemDecodedPtr(unsigned int index) const
{
    ServiceAreaInformation_Item* item_p = NULL;
    ServiceAreaInformation ServiceAreaInformation_a = NULL;
    if (mobilityRestrictionList_mp->serviceAreaInformation != NULL) {
        ServiceAreaInformation_a = mobilityRestrictionList_mp->serviceAreaInformation;
    }
    if (ServiceAreaInformation_a != NULL && (index >= 0 && index < numofserviceAreaInformation_m)) {
        unsigned int i = 0;
        for (; i < index && ServiceAreaInformation_a != NULL; i++) {
            ServiceAreaInformation_a = ServiceAreaInformation_a->next;
        }
        if (ServiceAreaInformation_a != NULL && i == index) {
            item_p = &(ServiceAreaInformation_a->value);
        }
    }

    return item_p;
}

ServiceAreaInformation_Item*
InitialContextSetupRequestPdu::getServiceAreaInformationItemEncodePtr(unsigned int index)
{
    if (index >= 0 && index < maxnoofEPLMNsPlusOne) {
        return &(serviceAreaInformation_ma[index].value);
    } else {
        return NULL;
    }
}

void InitialContextSetupRequestPdu::setNumofServiceAreaInformationItem_v(unsigned int num)
{
    numofserviceAreaInformation_m = num;
}

unsigned int
InitialContextSetupRequestPdu::getNumofServiceAreaInformationItem() const
{
    return numofserviceAreaInformation_m;
}

TAC* InitialContextSetupRequestPdu::getForbiddenTACsDecodedPtr(
    ForbiddenAreaInformation_Item* forbidAreaInfo, unsigned int indbidTACs) const
{
    TAC* item_p = NULL;
    ForbiddenTACs forbiddenTACs_p = NULL;
    if (forbidAreaInfo != NULL) {
        forbiddenTACs_p = forbidAreaInfo->forbiddenTACs;
        if (forbiddenTACs_p != NULL && (indbidTACs >= 0 && indbidTACs < maxnoofForbTACs)) {
            unsigned int i = 0;
            for (; i < indbidTACs && forbiddenTACs_p != NULL; i++) {
                forbiddenTACs_p = forbiddenTACs_p->next;
            }

            if (forbiddenTACs_p != NULL && i == indbidTACs) {
                item_p = &(forbiddenTACs_p->value);
            }
        }
    }

    return item_p;
}

TAC* InitialContextSetupRequestPdu::getForbiddenTACsEncodePtr(unsigned int indPlmn, unsigned int indbidTACs)
{
    if ((indPlmn >= 0 && indPlmn < maxnoofEPLMNsPlusOne) && 
        (indbidTACs >= 0 && indbidTACs < maxnoofForbTACs)) {
        return &(forbiddenTACs_ma[indPlmn][indbidTACs].value);
    } else {
        return NULL;
    }
}

void InitialContextSetupRequestPdu::setNumofForbiddenTACs_v(unsigned int num, unsigned int index)
{
    if (index >= 0 && index < maxnoofEPLMNsPlusOne) {
        numofForbiddenTACs_m[index] = num;
    }
}

unsigned int
InitialContextSetupRequestPdu::getNumofForbiddenTACs(unsigned int index) const
{
    if (index >= 0 && index < maxnoofEPLMNsPlusOne) {
        return numofForbiddenTACs_m[index];
    } else {
        return 0;
    }
}

TAC* InitialContextSetupRequestPdu::getAllowedTACsDecodedPtr(
        ServiceAreaInformation_Item* serAreaInfo, unsigned int allowedTACs) const
{
    TAC* item_p = NULL;
    AllowedTACs allowedTACs_p = NULL;
    if (serAreaInfo != NULL) {
        allowedTACs_p = serAreaInfo->allowedTACs;
        if (allowedTACs_p != NULL && (allowedTACs >= 0 && allowedTACs < maxnoofAllowedAreas)) {
            unsigned int i = 0;
            for (; i < allowedTACs && allowedTACs_p != NULL; i++) {
                allowedTACs_p = allowedTACs_p->next;
            }

            if (allowedTACs_p != NULL && i == allowedTACs) {
                item_p = &(allowedTACs_p->value);
            }
        }
    }

    return item_p;
}

TAC* InitialContextSetupRequestPdu::getAllowedTACsEncodePtr(unsigned int indPlmn, unsigned int allowedTACs)
{
    if ((indPlmn >= 0 && indPlmn < maxnoofEPLMNsPlusOne) && 
        (allowedTACs >= 0 && allowedTACs < maxnoofAllowedAreas)) {
        return &(allowedTACs_ma[indPlmn][allowedTACs].value);
    } else {
        return NULL;
    }
}

void InitialContextSetupRequestPdu::setNumofAllowedTACs_v(unsigned int num, unsigned int index)
{
    if (index >= 0 && index < maxnoofEPLMNsPlusOne) {
        numofAllowedTACs_m[index] = num;
    }
}

unsigned int
InitialContextSetupRequestPdu::getNumofAllowedTACs(unsigned int index) const
{
    if (index >= 0 && index < maxnoofEPLMNsPlusOne) {
        return numofAllowedTACs_m[index];
    } else {
        return 0;
    }
}

TAC* InitialContextSetupRequestPdu::getNotAllowedTACsDecodedPtr(
        ServiceAreaInformation_Item* serAreaInfo, unsigned int notallowedTACs) const
{
    TAC* item_p = NULL;
    NotAllowedTACs notAllowedTACs_p = NULL;
    if (serAreaInfo != NULL) {
        notAllowedTACs_p = serAreaInfo->notAllowedTACs;
        if (notAllowedTACs_p != NULL && (notallowedTACs >= 0 && notallowedTACs < maxnoofAllowedAreas)) {
            unsigned int i = 0;
            for (; i < notallowedTACs && notAllowedTACs_p != NULL; i++) {
                notAllowedTACs_p = notAllowedTACs_p->next;
            }

            if (notAllowedTACs_p != NULL && i == notallowedTACs) {
                item_p = &(notAllowedTACs_p->value);
            }
        }
    }
    return item_p;
}

TAC* InitialContextSetupRequestPdu::getNotAllowedTACsEncodePtr(unsigned int indPlmn, unsigned int notallowedTACs)
{
    if ((indPlmn >= 0 && indPlmn < maxnoofEPLMNsPlusOne) && (notallowedTACs >= 0 && notallowedTACs < maxnoofAllowedAreas)) {
        return &(notAllowedTACs_ma[indPlmn][notallowedTACs].value);
    } else {
        return NULL;
    }
}

void InitialContextSetupRequestPdu::setNumofNotAllowedTACs_v(unsigned int num, unsigned int index)
{
    if (index >= 0 && index < maxnoofEPLMNsPlusOne) {
        numofNotAllowedTACs_m[index] = num;
    }
}

unsigned int
InitialContextSetupRequestPdu::getNumofNotAllowedTACs(unsigned int index) const
{
    if (index >= 0 && index < maxnoofEPLMNsPlusOne) {
        return numofNotAllowedTACs_m[index];
    } else {
        return 0;
    }
}

//----------------------------------------------------------------------------
// This is a public API. Refer to InitialContextSetupRequestPdu.h
//----------------------------------------------------------------------------
AllowedNSSAI*
InitialContextSetupRequestPdu::getAllowedNssaiDecodedPtr() const
{
    return allowedNssai_mp;
}

AllowedNSSAI*
InitialContextSetupRequestPdu::getAllowedNssaiEncodePtr()
{
    return &allowedNssai_m;
}

AllowedNSSAI_Item*
InitialContextSetupRequestPdu::getAllowedNssai_ItemDecodedPtr(unsigned int index) const
{
    AllowedNSSAI_Item* item_p = NULL;

    AllowedNSSAI allowedNssai_p = NULL;

    if (allowedNssai_mp != NULL) {
        allowedNssai_p = *(allowedNssai_mp);
    }

    if (allowedNssai_p != NULL
        && (index >= 0) && index < numOfAllowedNassiItem_m) {
        unsigned int i = 0;
        for (; i < index && allowedNssai_p != NULL; i++) {
            allowedNssai_p = allowedNssai_p->next;
        }
        if (allowedNssai_p != NULL && i == index) {
            item_p = &(allowedNssai_p->value);
        }
    }
    return item_p;
}
AllowedNSSAI_Item*
InitialContextSetupRequestPdu::getAllowedNssai_ItemEncodePtr(unsigned int index)
{
    if (index >= 0 && index < maxnoofNSSAIs_c) {
        return &(allowedNssai_ma[index].value);
    } else {
        return NULL;
    }
}

void InitialContextSetupRequestPdu::setNumofAllowedNssai_Item_m_v(unsigned int num)
{
    numOfAllowedNassiItem_m = num;
}
unsigned int
InitialContextSetupRequestPdu::getNumofAllowedNssai_Item() const
{
    return numOfAllowedNassiItem_m;
}

//----------------------------------------------------------------------------
// This is a public API. Refer to InitialContextSetupRequestPdu.h
//----------------------------------------------------------------------------
GUAMI*
InitialContextSetupRequestPdu::getGuamiDecodedPtr() const
{
    return guami_mp;
}

GUAMI*
InitialContextSetupRequestPdu::getGuamiEncodePtr()
{
    return &guami_m;
}

//----------------------------------------------------------------------------
// This is a public API. Refer to InitialContextSetupRequestPdu.h
//----------------------------------------------------------------------------
UESecurityCapabilities*
InitialContextSetupRequestPdu::getUeSecurityCapabilitiesDecodedPtr() const
{
    return ueSecurityCapabilities_mp;
}

UESecurityCapabilities*
InitialContextSetupRequestPdu::getUeSecurityCapabilitiesEncodePtr()
{
    return &ueSecurityCapabilities_m;
}

//----------------------------------------------------------------------------
// This is a public API. Refer to InitialContextSetupRequestPdu.h
//----------------------------------------------------------------------------
SecurityKey*
InitialContextSetupRequestPdu::getSecurityKeyDecodedPtr() const
{
    return securityKey_mp;
}

SecurityKey*
InitialContextSetupRequestPdu::getSecurityKeyEncodePtr()
{
    return &securityKey_m;
}

//----------------------------------------------------------------------------
// This is a public API. Refer to InitialContextSetupRequestPdu.h
//----------------------------------------------------------------------------
TraceActivation*
InitialContextSetupRequestPdu::getTraceActivationDecodedPtr() const
{
    return traceActivation_mp;
}

TraceActivation*
InitialContextSetupRequestPdu::getTraceActivationEncodePtr()
{
    traceActPresent_m=true;
    return &traceActivation_m;
}

bool
InitialContextSetupRequestPdu::isTraceActPresent()
{
    return traceActPresent_m;
}

//----------------------------------------------------------------------------
// This is a public API. Refer to InitialContextSetupRequestPdu.h
//----------------------------------------------------------------------------
UERadioCapability*
InitialContextSetupRequestPdu::getUERadioCapabilityDecodedPtr() const
{
    return ueRadioCapability_mp;
}

UERadioCapability*
InitialContextSetupRequestPdu::getUERadioCapabilityEncodePtr()
{
    ueRadioCapPresent_m=true;
    return &ueRadioCapability_m;
}

bool
InitialContextSetupRequestPdu::isUeRadioPresent()
{
    return ueRadioCapPresent_m;
}

//----------------------------------------------------------------------------
// This is a public API. Refer to InitialContextSetupRequestPdu.h
//----------------------------------------------------------------------------
MaskedIMEISV*
InitialContextSetupRequestPdu::getMaskedIMEISVDecodedPtr() const
{
    return maskedIMEISV_mp;
}

MaskedIMEISV*
InitialContextSetupRequestPdu::getMaskedIMEISVEncodePtr()
{
    maskedIMEISVPresent_m=true;
    return &maskedIMEISV_m;
}

bool
InitialContextSetupRequestPdu::isMaskedIMEISVPresent()
{
    return maskedIMEISVPresent_m;
}

//----------------------------------------------------------------------------
// This is a public API. Refer to InitialContextSetupRequestPdu.h
//----------------------------------------------------------------------------
EmergencyFallbackIndicator*
InitialContextSetupRequestPdu::getEmergencyFallbackIndicatorDecodedPtr() const
{
    return emergencyFallbackIndicator_mp;
}

EmergencyFallbackIndicator*
InitialContextSetupRequestPdu::getEmergencyFallbackIndicatorEncodePtr()
{
    emerFallbackPresent_m=true;
    return &emergencyFallbackIndicator_m;
}

bool
InitialContextSetupRequestPdu::isEmerFallbackPresent()
{
    return emerFallbackPresent_m;
}

//----------------------------------------------------------------------------
// This is a public API. Refer to InitialContextSetupRequestPdu.h
//----------------------------------------------------------------------------
RRCInactiveTransitionReportRequest*
InitialContextSetupRequestPdu::getRRCInactiveTransitionReportRequestDecodedPtr() const
{
    return rrcInactiveTransitionReportRequest_mp;
}

RRCInactiveTransitionReportRequest*
InitialContextSetupRequestPdu::getRRCInactiveTransitionReportRequestEncodePtr()
{
    RRCReportPresent_m=true;
    return &rrcInactiveTransitionReportRequest_m;
}

bool
InitialContextSetupRequestPdu::isRRCReportPresent()
{
    return RRCReportPresent_m;
}

//----------------------------------------------------------------------------
// This is a public API. Refer to InitialContextSetupRequestPdu.h
//----------------------------------------------------------------------------
UERadioCapabilityForPaging*
InitialContextSetupRequestPdu::getUeRadioCapabilityForPagingDecodedPtr() const
{
    return ueRadioCapabilityForPaging_mp;
}

UERadioCapabilityForPaging*
InitialContextSetupRequestPdu::getUeRadioCapabilityForPagingEncodePtr()
{
    ueRadioCapabilityForPagingPresent_m=true;
    return &ueRadioCapabilityForPaging_m;
}

bool
InitialContextSetupRequestPdu::isUeRadioCapabilityForPagingPresent()
{
    return ueRadioCapabilityForPagingPresent_m;
}

//----------------------------------------------------------------------------
// This is a public API. Refer to InitialContextSetupRequestPdu.h
//----------------------------------------------------------------------------
RedirectionVoiceFallback*
InitialContextSetupRequestPdu::getRedirectionVoiceFallbackDecodedPtr() const
{
    return redirectionVoiceFallback_mp;
}

RedirectionVoiceFallback*
InitialContextSetupRequestPdu::getRedirectionVoiceFallbackEncodePtr()
{
    redirectionVoiceFallbackPresent_m=true;
    return &redirectionVoiceFallback_m;
}

bool
InitialContextSetupRequestPdu::isRedirectionVoiceFallbackPresent()
{
    return redirectionVoiceFallbackPresent_m;
}

//----------------------------------------------------------------------------
// This is a public API. Refer to InitialContextSetupRequestPdu.h
//----------------------------------------------------------------------------
CNAssistedRANTuning*
InitialContextSetupRequestPdu::getCNAssistedRANTuningDecodedPtr() const
{
    return cnAssistedRANTuning_mp;
}

CNAssistedRANTuning*
InitialContextSetupRequestPdu::getCNAssistedRANTuningEncodePtr()
{
    cnAssistedRANTuningPresent_m=true;
    return &cnAssistedRANTuning_m;
}

bool
InitialContextSetupRequestPdu::isCNAssistedRANTuningPresent()
{
    return cnAssistedRANTuningPresent_m;
}

ExpectedUEMovingTrajectoryItem*
InitialContextSetupRequestPdu::getExpectedUEMovingTrajectoryItem_maDecodedPtr(unsigned int index) const
{
    ExpectedUEMovingTrajectoryItem* item_p = NULL;
    ExpectedUEMovingTrajectory ExpectedUEMovingTrajectory_b = NULL;
    if (coreNetworkAssistanceInfo_mp->expectedUEBehaviour.expectedUEMovingTrajectory != NULL) {
        ExpectedUEMovingTrajectory_b =
            coreNetworkAssistanceInfo_mp->expectedUEBehaviour.expectedUEMovingTrajectory;
    }

    if (ExpectedUEMovingTrajectory_b != NULL
        && (index >= 0 && index < numOfExpectedUEMovingTrajectoryItem_mb)) {
        unsigned int i = 0;
        for (; i < index && ExpectedUEMovingTrajectory_b != NULL; i++) {
            ExpectedUEMovingTrajectory_b = ExpectedUEMovingTrajectory_b->next;
        }
        if (ExpectedUEMovingTrajectory_b != NULL && i == index) {
            item_p = &(ExpectedUEMovingTrajectory_b->value);
        }
    }

    return item_p;
}

ExpectedUEMovingTrajectoryItem*
InitialContextSetupRequestPdu::getExpectedUEMovingTrajectoryItem_maEncodePtr(unsigned int index)
{
    if (index >= 0 && index < maxnoofCellsUEMovingTrajectory) {
        return &(expectedUEMovingTrajectory_mb[index].value);
    } else {
        return NULL;
    }
}

void InitialContextSetupRequestPdu::setNumOfExpectedUEMovingTrajectoryItem_ma_v(unsigned int num)
{
    numOfExpectedUEMovingTrajectoryItem_mb = num;
}

unsigned int
InitialContextSetupRequestPdu::getNumOfExpectedUEMovingTrajectoryItem_ma() const
{
    return numOfExpectedUEMovingTrajectoryItem_mb;
}

//----------------------------------------------------------------------------
// This is a public API. Refer to InitialContextSetupRequestPdu.h
//----------------------------------------------------------------------------
LocationReportingRequestType*
InitialContextSetupRequestPdu::getLocationReportingRequestTypeDecodedPtr() const
{
    return locationReportRequestType_mp;
}

LocationReportingRequestType*
InitialContextSetupRequestPdu::getLocationReportingRequestTypeEncodePtr()
{
    locationReportRequestTypePresent_m=true;
    return &locationReportRequestType_m;
}

bool
InitialContextSetupRequestPdu::isLocationReportingRequestTypePresent()
{
    return locationReportRequestTypePresent_m;
}

//----------------------------------------------------------------------------
// This is a public API. Refer to InitialContextSetupRequestPdu.h
//----------------------------------------------------------------------------
AreaOfInterestItem*
InitialContextSetupRequestPdu::getAreaOfInterestItemDecodedPtr(unsigned int index) const
{
    AreaOfInterestItem* item_p = NULL;
    AreaOfInterestList AreaOfInterestList_b = NULL;
    if (locationReportRequestType_mp->areaOfInterestList != NULL) {
        AreaOfInterestList_b = locationReportRequestType_mp->areaOfInterestList;
    }

    if (AreaOfInterestList_b != NULL
        && (index >= 0 && index < numofAreaOfInterestItem_m)) {
        unsigned int i = 0;
        for (; i < index && AreaOfInterestList_b != NULL; i++) {
            AreaOfInterestList_b = AreaOfInterestList_b->next;
        }
        if (AreaOfInterestList_b != NULL && i == index) {
            item_p = &(AreaOfInterestList_b->value);
        }
    }

    return item_p;
}

AreaOfInterestItem*
InitialContextSetupRequestPdu::getAreaOfInterestItemEncodePtr(unsigned int index)
{
    if (index >= 0 && index < maxnoofAoI) {
        return &(areaOfInterestList_ma[index].value);
    } else {
        return NULL;
    }
}

void InitialContextSetupRequestPdu::setNumOfAreaOfInterestItem_v(unsigned int num)
{
    numofAreaOfInterestItem_m = num;
}

unsigned int
InitialContextSetupRequestPdu::getNumOfAreaOfInterestItem() const
{
    return numofAreaOfInterestItem_m;
}

//----------------------------------------------------------------------------
// This is a public API. Refer to InitialContextSetupRequestPdu.h
//----------------------------------------------------------------------------
AreaOfInterestTAIItem*
InitialContextSetupRequestPdu::getAreaOfInterestTAIItemDecodedPtr(AreaOfInterestItem * taiPtr,
                                            unsigned int indTai) const
{
    AreaOfInterestTAIItem* item_p = NULL;
    AreaOfInterestTAIList tai_p = NULL;
    if (taiPtr != NULL)
    {
        tai_p = taiPtr->areaOfInterest.areaOfInterestTAIList;
        if (tai_p != NULL
            && (indTai >= 0 && indTai < maxnoofAoI))
        {
            unsigned int i = 0;
            for (; i < indTai && tai_p != NULL; i++)
            {
                tai_p = tai_p->next;
            }

            if(tai_p != NULL  && i == indTai)
            {
                item_p = &(tai_p->value);
            }
        }
    }

    return item_p;

}

AreaOfInterestTAIItem*
InitialContextSetupRequestPdu::getAreaOfInterestTAIItemEncodePtr(unsigned int indTACs,
                                                   unsigned int indTai)
{
    if ((indTACs >= 0 && indTACs < maxnoofTAIinAoI) &&
        (indTai >= 0 && indTai < maxnoofAoI))
    {
        return &(areaOfInterestTAIList_ma[indTACs][indTai].value);
    }
    else
    {
        return NULL;
    }
}

void
InitialContextSetupRequestPdu::setNumOfAreaOfInterestTAIItem_v(unsigned int num, unsigned int index)
{
    if (index >= 0 && index < maxnoofTAIinAoI)
    {
        numAreaOfInterestTAIList_m[index] = num;
    }
}

unsigned int
InitialContextSetupRequestPdu::getNumOfAreaOfInterestTAIItem(unsigned int index) const
{
    if (index >= 0 && index < maxnoofTAIinAoI)
    {
        return numAreaOfInterestTAIList_m[index];
    }
    else
    {
        return 0;
    }
}


//----------------------------------------------------------------------------
// This is a public API. Refer to InitialContextSetupRequestPdu.h
//----------------------------------------------------------------------------
AreaOfInterestCellItem*
InitialContextSetupRequestPdu::getAreaOfInterestCellItemDecodedPtr(AreaOfInterestItem * taiPtr,
                                            unsigned int indCell) const
{
    AreaOfInterestCellItem* item_p = NULL;
    AreaOfInterestCellList cell_p = NULL;
    if (taiPtr != NULL)
    {
        cell_p = taiPtr->areaOfInterest.areaOfInterestCellList;
        if (cell_p != NULL
            && (indCell >= 0 && indCell < maxnoofAoI))
        {
            unsigned int i = 0;
            for (; i < indCell && cell_p != NULL; i++)
            {
                cell_p = cell_p->next;
            }

            if(cell_p != NULL  && i == indCell)
            {
                item_p = &(cell_p->value);
            }
        }
    }

    return item_p;

}

AreaOfInterestCellItem*
InitialContextSetupRequestPdu::getAreaOfInterestCellItemEncodePtr(unsigned int indTACs,
                                                   unsigned int indCell)
{
    if ((indTACs >= 0 && indTACs < maxnoofCellinAoI) &&
        (indCell >= 0 && indCell < maxnoofAoI))
    {
        return &(areaOfInterestCellList_ma[indTACs][indCell].value);
    }
    else
    {
        return NULL;
    }
}

void
InitialContextSetupRequestPdu::setNumOfAreaOfInterestCellItem_v(unsigned int num, unsigned int index)
{
    if (index >= 0 && index < maxnoofCellinAoI)
    {
        numAreaOfInterestCellList_m[index] = num;
    }
}

unsigned int
InitialContextSetupRequestPdu::getNumOfAreaOfInterestCellItem(unsigned int index) const
{
    if (index >= 0 && index < maxnoofCellinAoI)
    {
        return numAreaOfInterestCellList_m[index];
    }
    else
    {
        return 0;
    }
}


//----------------------------------------------------------------------------
// This is a public API. Refer to InitialContextSetupRequestPdu.h
//----------------------------------------------------------------------------
AreaOfInterestRANNodeItem*
InitialContextSetupRequestPdu::getAreaOfInterestRANNodeItemDecodedPtr(AreaOfInterestItem * taiPtr, 
                                            unsigned int indRANNode) const
{
    AreaOfInterestRANNodeItem* item_p = NULL;
    AreaOfInterestRANNodeList ranNode_p = NULL;
    if (taiPtr != NULL)
    {
        ranNode_p = taiPtr->areaOfInterest.areaOfInterestRANNodeList;
        if (ranNode_p != NULL
            && (indRANNode >= 0 && indRANNode < maxnoofAoI))
        {
            unsigned int i = 0;
            for (; i < indRANNode && ranNode_p != NULL; i++)
            {
                ranNode_p = ranNode_p->next;
            }

            if(ranNode_p != NULL  && i == indRANNode)
            {
                item_p = &(ranNode_p->value);
            }
        }
    }

    return item_p;

}

AreaOfInterestRANNodeItem*
InitialContextSetupRequestPdu::getAreaOfInterestRANNodeItemEncodePtr(unsigned int indTACs,
                                                   unsigned int indRANNode)
{
    if ((indTACs >= 0 && indTACs < maxnoofRANNodeinAoI) &&
        (indRANNode >= 0 && indRANNode < maxnoofAoI))
    {
        return &(areaOfInterestRANNodeList_ma[indTACs][indRANNode].value);
    }
    else
    {
        return NULL;
    }
}

void
InitialContextSetupRequestPdu::setNumOfAreaOfInterestRANNodeItem_v(unsigned int num, unsigned int index)
{
    if (index >= 0 && index < maxnoofRANNodeinAoI)
    {
        numAreaOfInterestRANNodeList_m[index] = num;
    }
}

unsigned int
InitialContextSetupRequestPdu::getNumOfAreaOfInterestRANNodeItem(unsigned int index) const
{
    if (index >= 0 && index < maxnoofRANNodeinAoI)
    {
        return numAreaOfInterestRANNodeList_m[index];
    }
    else
    {
        return 0;
    }
}

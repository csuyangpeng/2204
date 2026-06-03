#ifndef __ngApOssInitialContextSetupRequestPdu_h__INCLUDED__
#define __ngApOssInitialContextSetupRequestPdu_h__INCLUDED__

#include "ngapOssPdu.h"
#include <ngapMsgTypes.h>
///////////////////////////////////////////////////////////////////////////////
// @class InitialContextSetupRequestPdu
//-- **************************************************************
//--
//-- Initial Context Setup Request
//--
//-- **************************************************************
//
//
//InitialContextSetupRequestIEs NGAP-PROTOCOL-IES ::= {
//  { ID id-AMF-UE-NGAP-ID                               CRITICALITY reject TYPE AMF-UE-NGAP-ID                                     PRESENCE mandatory  }|
//  { ID id-RAN-UE-NGAP-ID                               CRITICALITY reject TYPE RAN-UE-NGAP-ID                                     PRESENCE mandatory  }|
//  { ID id-OldAMF                                       CRITICALITY reject TYPE AMFName                                            PRESENCE optional       }|
//  { ID id-UEAggregateMaximumBitRate                    CRITICALITY reject TYPE UEAggregateMaximumBitRate                      PRESENCE conditional    }|
//  { ID id-CoreNetworkAssistanceInformationForInactive  CRITICALITY ignore TYPE CoreNetworkAssistanceInformationForInactive                PRESENCE optional       }|
//  { ID id-GUAMI                                        CRITICALITY reject TYPE GUAMI                                              PRESENCE mandatory  }|
//  { ID id-PDUSessionResourceSetupListCxtReq            CRITICALITY reject TYPE PDUSessionResourceSetupListCxtReq          PRESENCE optional       }|
//  { ID id-AllowedNSSAI                                 CRITICALITY ignore TYPE AllowedNSSAI                                       PRESENCE mandatory  }|
//  { ID id-UESecurityCapabilities                       CRITICALITY reject TYPE UESecurityCapabilities                         PRESENCE mandatory  }|
//  { ID id-SecurityKey                                  CRITICALITY reject TYPE SecurityKey                                        PRESENCE mandatory  }|
//  { ID id-TraceActivation                              CRITICALITY ignore TYPE TraceActivation                                    PRESENCE optional       }|
//  { ID id-MobilityRestrictionList                      CRITICALITY ignore TYPE MobilityRestrictionList                            PRESENCE optional       }|
//  { ID id-UERadioCapability                            CRITICALITY ignore TYPE UERadioCapability                                  PRESENCE optional       }|
//  { ID id-IndexToRFSP                                  CRITICALITY ignore TYPE IndexToRFSP                                        PRESENCE optional       }|
//  { ID id-MaskedIMEISV                                 CRITICALITY ignore TYPE MaskedIMEISV                                       PRESENCE optional       }|
//  { ID id-NAS-PDU                                      CRITICALITY ignore TYPE NAS-PDU                                            PRESENCE optional       }|
//  { ID id-EmergencyFallbackIndicator                   CRITICALITY reject TYPE EmergencyFallbackIndicator                     PRESENCE optional       }|
//  { ID id-RRCInactiveTransitionReportRequest           CRITICALITY ignore TYPE RRCInactiveTransitionReportRequest         PRESENCE optional       }|
//  { ID id-UERadioCapabilityForPaging                   CRITICALITY ignore TYPE UERadioCapabilityForPaging                     PRESENCE optional       }|
//  { ID id-RedirectionVoiceFallback                     CRITICALITY ignore TYPE RedirectionVoiceFallback                       PRESENCE optional       }|
//  { ID id-LocationReportingRequestType                 CRITICALITY ignore TYPE LocationReportingRequestType                   PRESENCE optional       }|
//  { ID id-CNAssistedRANTuning                          CRITICALITY ignore TYPE CNAssistedRANTuning                            PRESENCE optional       },
//  ...
//}
///////////////////////////////////////////////////////////////////////////////

class InitialContextSetupRequestPdu:public NgapOssPdu
{
  public:
    InitialContextSetupRequestPdu();
    //virtual
    ~InitialContextSetupRequestPdu();
    
    /////////////////////////////////////////////////////////////
    /// @brief reset_v
    ///  resets all data members to default value
    /////////////////////////////////////////////////////////////
    //virtual
    void reset_v();

    /////////////////////////////////////////////////////////////
    /// @brief encode Oss to PER using ASN.1
    ///  builds and processes the OSS PDU to encode
    ///  it to PER using ASN.1
    /////////////////////////////////////////////////////////////
    //virtual
    unsigned int encodeOssToPer(NgapOssCtxt* ctxt);

    /////////////////////////////////////////////////////////////
    /// @brief get Decoded Oss Data Pointers
    ///  processes the decoded OSS buffer and populates
    ///  the data access pointers in the class for the IEs present in the
    ///  decoded OSS buffer
    /////////////////////////////////////////////////////////////
    //virtual
    void getDecodedOssDataPtrs_v();

    /////////////////////////////////////////////////////////////
    /// @brief get InitialContextSetupRequest pointer
    /// returns the pointer to the
    /// InitialContextSetupRequest in the decoded message
    /////////////////////////////////////////////////////////////
    InitialContextSetupRequest * getInitialContextSetupRequestDecodedPtr() const;

    /////////////////////////////////////////////////////////////
    /// RAN_UE_NGAP_ID
    /////////////////////////////////////////////////////////////
    RAN_UE_NGAP_ID* getRanUeNgapIdDecodedPtr() const;
    RAN_UE_NGAP_ID* getRanUeNgapIdEncodePtr();

    /////////////////////////////////////////////////////////////
    /// AMF_UE_NGAP_ID
    /////////////////////////////////////////////////////////////
    AMF_UE_NGAP_ID* getAmfUeNgapIdDecodedPtr() const;
    AMF_UE_NGAP_ID* getAmfUeNgapIdEncodePtr();

    /////////////////////////////////////////////////////////////
    /// NAS_PDU
    /////////////////////////////////////////////////////////////
    NAS_PDU* getNasPduDecodedPtr() const;
    NAS_PDU* getNasPduEncodePtr();
    bool isNasPduPresent();

    /////////////////////////////////////////////////////////////
    /// AmfName 
    /////////////////////////////////////////////////////////////
    AMFName* getAmfNameDecodedPtr() const;
    AMFName* getAmfNameEncodePtr();
    bool isAmfNamePresent();

    /////////////////////////////////////////////////////////////
    /// UE Aggregate Maximum Bit Rate
    /////////////////////////////////////////////////////////////
    UEAggregateMaximumBitRate* getUeAggregateMaximumBitRateDecodedPtr() const;
    UEAggregateMaximumBitRate* getUeAggregateMaximumBitRateEncodePtr();
    bool isUeAMBRPresent();

    /////////////////////////////////////////////////////////////
    /// Index To RFSP
    /////////////////////////////////////////////////////////////
    IndexToRFSP* getIndexToRFSPDecodedPtr() const;
    IndexToRFSP* getIndexToRFSPEncodePtr();
    bool isIndexToRFSPPresent();
        
    /////////////////////////////////////////////////////////////
    /// Allowed NSSAI
    /////////////////////////////////////////////////////////////
    AllowedNSSAI* getAllowedNssaiDecodedPtr() const;
    AllowedNSSAI* getAllowedNssaiEncodePtr();

    AllowedNSSAI_Item* getAllowedNssai_ItemDecodedPtr(unsigned int index) const;
    AllowedNSSAI_Item* getAllowedNssai_ItemEncodePtr(unsigned int index);

    void setNumofAllowedNssai_Item_m_v(unsigned int num);
    unsigned int getNumofAllowedNssai_Item() const;

    /////////////////////////////////////////////////////////////
    /// CoreNetworkAssistanceInformation
    /////////////////////////////////////////////////////////////
    CoreNetworkAssistanceInformationForInactive * getCoreNetworkAssistanceInformationDecodedPtr() const;
    CoreNetworkAssistanceInformationForInactive * getCoreNetworkAssistanceInformationEncodePtr();
    bool isCoreNetworkPresent();

    TAIListForInactiveItem * getTAIListForInactiveItemDecodedPtr(unsigned int index) const;
    TAIListForInactiveItem * getTAIListForInactiveItemEncodePtr(unsigned int index);

    void setNumOfTAIListForInactiveItem_v(unsigned int num);
    unsigned int getNumOfTAIListForInactiveItem() const;
    
    ExpectedUEMovingTrajectoryItem * getExpectedUEMovingTrajectoryItemDecodedPtr(unsigned int index) const;
    ExpectedUEMovingTrajectoryItem * getExpectedUEMovingTrajectoryItemEncodePtr(unsigned int index);

    void setNumOfExpectedUEMovingTrajectoryItem_v(unsigned int num);
    unsigned int getNumOfExpectedUEMovingTrajectoryItem() const;

    /////////////////////////////////////////////////////////////
    /// PDUSessionResourceSetupListCxtReq
    /////////////////////////////////////////////////////////////
    PDUSessionResourceSetupListCxtReq* getPDUSessionResourceSetupListCxtReqDecodedPtr() const;
    PDUSessionResourceSetupListCxtReq* getPDUSessionResourceSetupListCxtReqEncodePtr();
    
    PDUSessionResourceSetupItemCxtReq* getPDUSessionResourceSetupItemCxtReqDecodedPtr(unsigned int index) const;
    PDUSessionResourceSetupItemCxtReq* getPDUSessionResourceSetupItemCxtReqEncodePtr(unsigned int index);

    void setNumofPDUSessionResourceSetupListItemCxtReq_v(unsigned int num);
    unsigned int getNumofPDUSessionResourceSetupListItemCxtReq() const;

    /////////////////////////////////////////////////////////////
    /// MobilityRestrictionList
    /////////////////////////////////////////////////////////////
    MobilityRestrictionList* getMobilityRestrictionListDecodedPtr() const;
    MobilityRestrictionList* getMobilityRestrictionListEncodePtr();
    
    PLMNIdentity* getPLMNIdentityDecodedPtr(unsigned int index) const;
    PLMNIdentity* getPLMNIdentityEncodePtr(unsigned int index);

    void setNumofPLMNIdentity_v(unsigned int num);
    unsigned int getNumofPLMNIdentity() const;
    
    RATRestrictions_Item* getRATRestrictionsItemDecodedPtr(unsigned int index) const;
    RATRestrictions_Item* getRATRestrictionsItemEncodePtr(unsigned int index);

    void setNumofRATRestrictionsItem_v(unsigned int num);
    unsigned int getNumofRATRestrictionsItem() const;
    
    ForbiddenAreaInformation_Item* getForbiddenAreaInformationItemDecodedPtr(unsigned int index) const;
    ForbiddenAreaInformation_Item* getForbiddenAreaInformationItemEncodePtr(unsigned int index);

    void setNumofForbiddenAreaInformationItem_v(unsigned int num);
    unsigned int getNumofForbiddenAreaInformationItem() const;
    
    ServiceAreaInformation_Item* getServiceAreaInformationItemDecodedPtr(unsigned int index) const;
    ServiceAreaInformation_Item* getServiceAreaInformationItemEncodePtr(unsigned int index);

    void setNumofServiceAreaInformationItem_v(unsigned int num);
    unsigned int getNumofServiceAreaInformationItem() const;

    TAC* getForbiddenTACsDecodedPtr(ForbiddenAreaInformation_Item* forbidAreaInfo, unsigned int indbidTACs) const;
    TAC* getForbiddenTACsEncodePtr(unsigned int indPlmn, unsigned int indbidTACs);

    void setNumofForbiddenTACs_v(unsigned int num, unsigned int index);
    unsigned int getNumofForbiddenTACs(unsigned int index) const;
    
    TAC* getAllowedTACsDecodedPtr(ServiceAreaInformation_Item* serAreaInfo,unsigned int allowedTACs) const;
    TAC* getAllowedTACsEncodePtr(unsigned int indPlmn,unsigned int allowedTACs);
    
    void setNumofAllowedTACs_v(unsigned int num, unsigned int index);
    unsigned int getNumofAllowedTACs(unsigned int index) const;

    TAC* getNotAllowedTACsDecodedPtr(ServiceAreaInformation_Item* serAreaInfo,unsigned int notallowedTACs) const;
    TAC* getNotAllowedTACsEncodePtr(unsigned int indPlmn,unsigned int notallowedTACs);

    void setNumofNotAllowedTACs_v(unsigned int num, unsigned int index);
    unsigned int getNumofNotAllowedTACs(unsigned int index) const;

    /////////////////////////////////////////////////////////////
    /// GUAMI
    /////////////////////////////////////////////////////////////
    GUAMI* getGuamiDecodedPtr() const;
    GUAMI* getGuamiEncodePtr();

    /////////////////////////////////////////////////////////////
    /// UE Security Capabilities
    /////////////////////////////////////////////////////////////
    UESecurityCapabilities* getUeSecurityCapabilitiesDecodedPtr() const;
    UESecurityCapabilities* getUeSecurityCapabilitiesEncodePtr();

    /////////////////////////////////////////////////////////////
    /// SecurityKey
    /////////////////////////////////////////////////////////////
    SecurityKey* getSecurityKeyDecodedPtr() const;
    SecurityKey* getSecurityKeyEncodePtr();

    /////////////////////////////////////////////////////////////
    /// TraceActivation
    /////////////////////////////////////////////////////////////
    TraceActivation* getTraceActivationDecodedPtr() const;
    TraceActivation* getTraceActivationEncodePtr();
    bool isTraceActPresent();

    /////////////////////////////////////////////////////////////
    /// UERadioCapability
    /////////////////////////////////////////////////////////////
    UERadioCapability* getUERadioCapabilityDecodedPtr() const;
    UERadioCapability* getUERadioCapabilityEncodePtr();
    bool isUeRadioPresent();

    /////////////////////////////////////////////////////////////
    /// MaskedIMEISV
    /////////////////////////////////////////////////////////////
    MaskedIMEISV* getMaskedIMEISVDecodedPtr() const;
    MaskedIMEISV* getMaskedIMEISVEncodePtr();
    bool isMaskedIMEISVPresent();

    /////////////////////////////////////////////////////////////
    /// EmergencyFallbackIndicator
    /////////////////////////////////////////////////////////////
    EmergencyFallbackIndicator* getEmergencyFallbackIndicatorDecodedPtr() const;
    EmergencyFallbackIndicator* getEmergencyFallbackIndicatorEncodePtr();
    bool isEmerFallbackPresent();

    /////////////////////////////////////////////////////////////
    /// RRCInactiveTransitionReportRequest
    /////////////////////////////////////////////////////////////
    RRCInactiveTransitionReportRequest* getRRCInactiveTransitionReportRequestDecodedPtr() const;
    RRCInactiveTransitionReportRequest* getRRCInactiveTransitionReportRequestEncodePtr();
    bool isRRCReportPresent();

    /////////////////////////////////////////////////////////////
    /// UERadioCapabilityForPaging
    /////////////////////////////////////////////////////////////
    UERadioCapabilityForPaging* getUeRadioCapabilityForPagingDecodedPtr() const;
    UERadioCapabilityForPaging* getUeRadioCapabilityForPagingEncodePtr();
    bool isUeRadioCapabilityForPagingPresent();

    /////////////////////////////////////////////////////////////
    /// RedirectionVoiceFallback
    /////////////////////////////////////////////////////////////
    RedirectionVoiceFallback* getRedirectionVoiceFallbackDecodedPtr() const;
    RedirectionVoiceFallback* getRedirectionVoiceFallbackEncodePtr();
    bool isRedirectionVoiceFallbackPresent();

    /////////////////////////////////////////////////////////////
    /// CNAssistedRANTuning
    /////////////////////////////////////////////////////////////
    CNAssistedRANTuning * getCNAssistedRANTuningDecodedPtr() const;
    CNAssistedRANTuning * getCNAssistedRANTuningEncodePtr();
    bool isCNAssistedRANTuningPresent();

    ExpectedUEMovingTrajectoryItem * getExpectedUEMovingTrajectoryItem_maDecodedPtr(unsigned int index) const;
    ExpectedUEMovingTrajectoryItem * getExpectedUEMovingTrajectoryItem_maEncodePtr(unsigned int index);

    void setNumOfExpectedUEMovingTrajectoryItem_ma_v(unsigned int num);
    unsigned int getNumOfExpectedUEMovingTrajectoryItem_ma() const;

    /////////////////////////////////////////////////////////////
    /// LocationReportingRequestType
    /////////////////////////////////////////////////////////////
    LocationReportingRequestType * getLocationReportingRequestTypeDecodedPtr() const;
    LocationReportingRequestType * getLocationReportingRequestTypeEncodePtr();
    bool isLocationReportingRequestTypePresent();

    AreaOfInterestItem* getAreaOfInterestItemDecodedPtr(unsigned int index) const;
    AreaOfInterestItem* getAreaOfInterestItemEncodePtr(unsigned int index);
    void setNumOfAreaOfInterestItem_v(unsigned int num);
    unsigned int getNumOfAreaOfInterestItem() const;

    AreaOfInterestTAIItem* getAreaOfInterestTAIItemDecodedPtr(AreaOfInterestItem * taiPtr,unsigned int indTai) const;
    AreaOfInterestTAIItem* getAreaOfInterestTAIItemEncodePtr(unsigned int indTACs,unsigned int indTai);
    void setNumOfAreaOfInterestTAIItem_v(unsigned int num, unsigned int index);
    unsigned int getNumOfAreaOfInterestTAIItem(unsigned int index) const;

    AreaOfInterestCellItem* getAreaOfInterestCellItemDecodedPtr(AreaOfInterestItem * taiPtr,unsigned int indCell) const;
    AreaOfInterestCellItem* getAreaOfInterestCellItemEncodePtr(unsigned int indTACs,unsigned int indCell);
    void setNumOfAreaOfInterestCellItem_v(unsigned int num, unsigned int index);
    unsigned int getNumOfAreaOfInterestCellItem(unsigned int index) const;

    AreaOfInterestRANNodeItem* getAreaOfInterestRANNodeItemDecodedPtr(AreaOfInterestItem * taiPtr,unsigned int indRANNode) const;
    AreaOfInterestRANNodeItem* getAreaOfInterestRANNodeItemEncodePtr(unsigned int indTACs,unsigned int indRANNode);
    void setNumOfAreaOfInterestRANNodeItem_v(unsigned int num, unsigned int index);
    unsigned int getNumOfAreaOfInterestRANNodeItem(unsigned int index) const;

  private:
    typedef InitialContextSetupRequestIe ProtocolIeType;
  
    //OSS structures
    InitialContextSetupRequest *InitialContextSetupRequest_mp; //for encode

    //AMF-UE-NGAP-ID
    AMF_UE_NGAP_ID amfUeNgapId_m;
    AMF_UE_NGAP_ID* amfUeNgapId_mp;

    //RAN-UE-NGAP-ID
    RAN_UE_NGAP_ID ranUeNgapId_m;
    RAN_UE_NGAP_ID* ranUeNgapId_mp;

    //NAS-PDU
    NAS_PDU nasPdu_m;
    unsigned char nasPdu_ma[maxNasPduValue_c];
    NAS_PDU* nasPdu_mp;
    bool nasPduPresent_m;

    //OldAMF
    AMFName amfName_m; 
    char amfNameValue_ma[maxAMFNameValue_c];
    AMFName *amfName_mp;
    bool amfNamePresent_m;  

    //UE Aggregate Maximum Bit Rate
    UEAggregateMaximumBitRate     ueAggregateMaximumBitRate_m; 
    UEAggregateMaximumBitRate*    ueAggregateMaximumBitRate_mp;
    bool ueAMBRPresent_m;

    //Index To RFSP
    IndexToRFSP  indexToRFSP_m; 
    IndexToRFSP* indexToRFSP_mp;
    bool indexToRFSPPresent_m;
    
    //Allowed NSSAI
    AllowedNSSAI  allowedNssai_m;
    struct AllowedNSSAI_ allowedNssai_ma[maxnoofNSSAIs_c];
    unsigned int numOfAllowedNassiItem_m;
    AllowedNSSAI* allowedNssai_mp;

    //GUAMI
    GUAMI guami_m;
    GUAMI* guami_mp;
    AmfIdentifier guamiIdentifier_ma;

    //Ue Security Capabilities
    UESecurityCapabilities ueSecurityCapabilities_m;
    UESecurityCapabilities* ueSecurityCapabilities_mp;
    UeSecurityCap ueSecurityCapabilitiesValue_ma;

    //SecurityKey
    SecurityKey securityKey_m;
    unsigned char securityKeyValue_ma[maxSecurityKeyValue_c];
    SecurityKey* securityKey_mp;

    //TraceActivation
    TraceActivation traceActivation_m;
    TraceAct traceActivationValue_ma;
    TraceActivation* traceActivation_mp;
    bool traceActPresent_m;

    //UERadioCapability
    UERadioCapability ueRadioCapability_m;
    unsigned char ueRadioCapabilityValue_ma[maxUeRadioCapabilityValue_c];
    UERadioCapability* ueRadioCapability_mp;
    bool ueRadioCapPresent_m;

    //MaskedIMEISV
    MaskedIMEISV maskedIMEISV_m;
    unsigned char maskedIMEISVValue_ma[maxMaskedIMEISV_c];
    MaskedIMEISV* maskedIMEISV_mp;
    bool maskedIMEISVPresent_m;

    //EmergencyFallbackIndicator
    EmergencyFallbackIndicator emergencyFallbackIndicator_m; 
    EmergencyFallbackIndicator* emergencyFallbackIndicator_mp; 
    bool emerFallbackPresent_m;

    //RRCInactiveTransitionReportRequest
    RRCInactiveTransitionReportRequest rrcInactiveTransitionReportRequest_m;
    RRCInactiveTransitionReportRequest* rrcInactiveTransitionReportRequest_mp;
    bool RRCReportPresent_m;

    //UERadioCapabilityForPaging
    UERadioCapabilityForPaging ueRadioCapabilityForPaging_m;
    unsigned char ueRadioCapabilityForPagingNr_ma[maxUeRadioCapabilityValue_c];
    unsigned char ueRadioCapabilityForPagingEutra_ma[maxUeRadioCapabilityValue_c];
    UERadioCapabilityForPaging* ueRadioCapabilityForPaging_mp;
    bool  ueRadioCapabilityForPagingPresent_m;

    //CoreNetworkAssistanceInformation
    CoreNetworkAssistanceInformationForInactive coreNetworkAssistanceInfo_m;
    CoreNetworkAssistanceInformationForInactive* coreNetworkAssistanceInfo_mp;
    CoreNetworkAssInfo CoreNetAssInfo_ma;
    bool coreNetworkPresent_m;

    struct TAIListForInactive_ taiListForInactive_ma[maxnoofTAIforInactive_c];
    unsigned int numOfTAIListForInactiveItem_m;

    struct ExpectedUEMovingTrajectory_ expectedUEMovingTrajectory_ma[maxnoofCellsUEMovingTrajectory_c];
    unsigned char nrCellmaId[maxnoofCellsUEMovingTrajectory_c];
    unsigned int numOfExpectedUEMovingTrajectoryItem_m;

    //MobilityRestrictionList
    MobilityRestrictionList mobilityRestrictionList_m;
    MobilityRestrictionList* mobilityRestrictionList_mp;

    struct EquivalentPLMNs_ equivalentPLMNs_ma[maxnoofEPLMNs_c]; 
    unsigned int numofEquivalentPLMNs_m;
 
    struct RATRestrictions_  rATRestrictions_ma[maxnoofEPLMNsPlusOne_c];
    unsigned int numofRATRestrictions_m;
    unsigned char ratResInfo_ma[maxnoofEPLMNsPlusOne_c];

    struct ForbiddenAreaInformation_ forbiddenAreaInformation_ma[maxnoofEPLMNsPlusOne_c];
    unsigned int numofForbiddenAreaInformation_m;
    //ForbiddenTACs
    unsigned int numofForbiddenTACs_m[maxnoofEPLMNsPlusOne_c];
    struct ForbiddenTACs_ forbiddenTACs_ma[maxnoofEPLMNsPlusOne_c][maxnoofForbTACs_c];

    struct ServiceAreaInformation_ serviceAreaInformation_ma[maxnoofEPLMNsPlusOne_c];
    unsigned int numofserviceAreaInformation_m;
    //AllowedTACs
    unsigned int numofAllowedTACs_m[maxnoofEPLMNsPlusOne_c];
    struct AllowedTACs_ allowedTACs_ma[maxnoofEPLMNsPlusOne_c][maxnoofAllowedAreas_c];
    //NotAllowedTACs
    unsigned int numofNotAllowedTACs_m[maxnoofEPLMNsPlusOne_c];
    struct NotAllowedTACs_ notAllowedTACs_ma[maxnoofEPLMNsPlusOne_c][maxnoofAllowedAreas_c];

    //PDUSessionResourceSetupListCxtReq
    PDUSessionResourceSetupListCxtReq pDUSessionResourceSetupListCxtReq_m;
    struct PDUSessionResourceSetupListCxtReq_ PDUSessionResSetupListCxtReq_ma[maxnoofPDUSessions_c];
    unsigned int numofPDUSessionResourceSetupListItemCxtReq_m;
    PDUSessionResourceSetupListCxtReq* pDUSessionResourceSetupListCxtReq_mp;

    NasPdu pduNasPdus_ma[maxnoofPDUSessions_c];
    SmfTransferPdu transfers_ma[maxnoofPDUSessions_c];

    //RedirectionVoiceFallback
    RedirectionVoiceFallback redirectionVoiceFallback_m;
    RedirectionVoiceFallback* redirectionVoiceFallback_mp;
    bool redirectionVoiceFallbackPresent_m;

    //LocationReportingRequestType
    LocationReportingRequestType locationReportRequestType_m;
    LocationReportingRequestType* locationReportRequestType_mp;
    bool locationReportRequestTypePresent_m;

    struct AreaOfInterestList_ areaOfInterestList_ma[maxnoofAoI_c];
    unsigned int numofAreaOfInterestItem_m;

    //AreaOfInterestTAIList
    struct AreaOfInterestTAIList_ areaOfInterestTAIList_ma[maxnoofAoI_c][maxnoofTAIinAoI_c];
    unsigned int numAreaOfInterestTAIList_m[maxnoofAoI_c];

    //AreaOfInterestCellList
    struct AreaOfInterestCellList_ areaOfInterestCellList_ma[maxnoofAoI_c][maxnoofCellinAoI_c];
    unsigned int numAreaOfInterestCellList_m[maxnoofAoI_c];
    unsigned char nrCellId_ma[maxnoofAoI_c][maxnoofCellinAoI_c][maxNrCellIdLength_c];

    //AreaOfInterestRANNodeList
    struct AreaOfInterestRANNodeList_ areaOfInterestRANNodeList_ma[maxnoofAoI_c][maxnoofRANNodeinAoI_c];
    unsigned int numAreaOfInterestRANNodeList_m[maxnoofAoI_c];
    unsigned char gnbIdValue_ma[maxnoofAoI_c][maxnoofRANNodeinAoI_c][maxGnbIdValue_c];

    //CNAssistedRANTuning
    CNAssistedRANTuning cnAssistedRANTuning_m;
    CNAssistedRANTuning* cnAssistedRANTuning_mp;
    CnAssRanParamsTuning CnAssRanParams_ma;
    bool cnAssistedRANTuningPresent_m;

    struct ExpectedUEMovingTrajectory_ expectedUEMovingTrajectory_mb[maxnoofCellsUEMovingTrajectory_c];
    unsigned char nrCellmbId[maxnoofCellsUEMovingTrajectory_c];
    unsigned int numOfExpectedUEMovingTrajectoryItem_mb;
};

#endif
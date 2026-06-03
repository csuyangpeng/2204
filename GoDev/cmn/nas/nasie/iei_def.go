package nasie

type Iei byte

const (
	//registration request msg
	IeiNoncurNativeNasKSI   Iei = 0xC0
	Iei5GmmCapability       Iei = 0x10
	IeiUeSecurityCapability Iei = 0x2E
	IeiRequestNssai         Iei = 0x2F
	IeiLastVisitedRegTAI    Iei = 0x52
	IeiS1UeNwCapability     Iei = 0x17
	IeiUplinkDataStatus     Iei = 0x40
	IeiPduSessionStatus     Iei = 0x50
	IeiMicoIndication       Iei = 0xB0
	IeiUeStatus             Iei = 0x2B
	IeiAdditionalGuti       Iei = 0x77
	IeiAllowedPduSessStatus Iei = 0x25
	IeiUeUsageSetting       Iei = 0x18
	IeiReqDrxParameters     Iei = 0x51
	IeiEpsNasMsgContainer   Iei = 0x70
	IeiLandIndication       Iei = 0x74
	IeiPayloadContainer     Iei = 0x7B
	IeiNetSliceIndication   Iei = 0x90
	IeiFiveGUpdate          Iei = 0x53
	IeiNasMsgContainer      Iei = 0x71

	//registration accept msg
	IeiGuti5G                                   Iei = 0x77
	IeiEquivalentPLMNs                          Iei = 0x4A
	IeiTAIList                                  Iei = 0x54
	IeiAllowedNSSAI                             Iei = 0x15
	IeiRejectedNSSAI                            Iei = 0x11
	IeiConfiguredNSSAI                          Iei = 0x31
	Iei5GSNetworkFeatureSupport                 Iei = 0x21
	IeiPDUSessionStatus                         Iei = 0x50
	IeiPDUSessionReactivationResult             Iei = 0x26
	IeiPDUSessionReactivationResultErrorCause   Iei = 0x72
	IeiLADNInformation                          Iei = 0x79
	IeiMICOIndication                           Iei = 0xB0
	IeiNetworkSlicingIndication                 Iei = 0x90
	IeiServiceAreaList                          Iei = 0x27
	IeiT3512Value                               Iei = 0x5E
	IeiNon3GPPDeRegistrationTimerValue          Iei = 0x5D
	IeiT3346Value                               Iei = 0x5F
	IeiT3502Value                               Iei = 0x16
	IeiEmergencyNumberList                      Iei = 0x34
	IeiExtendedEmergencyNumberList              Iei = 0x7A
	IeiSORTransparentContainer                  Iei = 0x73
	IeiEAPMessage                               Iei = 0x78
	IeiNSSAIInclusionMode                       Iei = 0xA0
	IeiOperatorDefinedAccessCategoryDefinitions Iei = 0x76
	IeiNegotiatedDRXParameters                  Iei = 0x51

	//registration complete msg
	IeiSorTransContainer Iei = 0x73

	//session request msg
	IeiPDUSessionType                        Iei = 0x90
	IeiSSCMode                               Iei = 0xA0
	Iei5GSMCapability                        Iei = 0x28
	IeiMaximumNumberOfSupportedPacketFilters Iei = 0x55
	IeiAlwaysOnPDUSessionRequested           Iei = 0xB0
	IeiSMPDUDNRequestContainer               Iei = 0x39
	IeiExtendedProtocolConfigurationOptions  Iei = 0x7B

	//session accept msg
	IeiSMCause                       Iei = 0x59
	IeiPDUAddress                    Iei = 0x29
	IeiRQTimerValue                  Iei = 0x56
	IeiSNSSAI                        Iei = 0x22
	IeiAlwaysOnPDUSessionIndication  Iei = 0x80
	IeiMappedEPSBearerContexts       Iei = 0x75
	IeiAuthorizedQoSFlowDescriptions Iei = 0x79
	IeiBackOffTimerValue             Iei = 0x37
	IeiAllowedSSCMode                Iei = 0xF0
	IeiFIVEGSMCongestionReAttemptInt Iei = 0x61
	IeiPduSessionId                  Iei = 0x12
	IeiOldPduSessionId               Iei = 0x59
	IeiAdditinalInfo                 Iei = 0x24
	IeiDnn                           Iei = 0x25

	// authentication request message
	IeiAuthRand   Iei = 0x21
	IeiAuthAutn   Iei = 0x20
	IeiAuthEapMsg Iei = 0x78

	// authentication response message
	IeiAuthResStar Iei = 0x2d

	// security mod command
	IeiImsiSv       Iei = 0x77
	IeiAdd5GSecInfo Iei = 0x36
	IeiABBA         Iei = 0x38

	//De-registration request (UE terminated de-registration)
	Iei5GMMCause Iei = 0x58

	//session modification msg
	IeiInterProctMaxDataRate Iei = 0x13
	IeiQosRules              Iei = 0x7A

	IeiAmbr Iei = 0x2A

	IeiSMCongestionReAttempInd Iei = 0x61

	// 24.501 ((V15.1.0 (2018-09))) Table 8.2.10.1.1: UL NAS TRANSPORT message content
	IeiPduSessId             Iei = 0x12
	IeiOldPDUSession         Iei = 0x59
	IeiRequestType           Iei = 0x80
	IeiDNN                   Iei = 0x25
	IeiAdditionalInformation Iei = 0x24

	// 24.501 ((V15.1.0 (2018-09))) Table 8.2.11.1.1: DL NAS TRANSPORT message content
	IeiFiveGMMcause Iei = 0x58
	IeiBackOffTimer Iei = 0x37
)

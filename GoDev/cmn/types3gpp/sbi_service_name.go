package types3gpp

//"nnrf-nfm"	Nnrf_NFManagement Service offered by the NRF
//"nnrf-disc"	Nnrf_NFDiscovery Service offered by the NRF
//"nudm-sdm"	Nudm_SubscriberDataManagement Service offered by the UDM
//"nudm-uecm"	Nudm_UEContextManagement Service offered by the UDM
//"nudm-ueau"	Nudm_UEAuthentication Service offered by the UDM
//"nudm-ee"	Nudm_EventExposure Service offered by the UDM
//"nudm-pp"	Nudm_ParameterProvision Service offered by the UDM
//"namf-comm"	Namf_Communication Service offered by the AMF
//"namf-evts"	Namf_EventExposure Service offered by the AMF
//"namf-mt"	Namf_MT Service offered by the AMF
//"namf-loc"	Namf_Location Service offered by the AMF
//"nsmf-pdusession"	Nsmf_PDUSession Service offered by the SMF
//"nsmf-event-exposure"	Nsmf_EventExposure Service offered by the SMF
//"nausf-auth"	Nausf_UEAuthentication Service offered by the AUSF
//"nausf-sorprotection"	Nausf_SoRProtection Service offered by the AUSF
//"nnef-pfdmanagement"	Nnef_PFDManagement offered by the NEF
//"npcf-am-policy-control"	Npcf_AMPolicyControl Service offered by the PCF
//"npcf-smpolicycontrol"	Npcf_SMPolicyControl Service offered by the PCF
//"npcf-policyauthorization"	Npcf_PolicyAuthorization Service offered by the PCF
//"npcf-bdtpolicycontrol"	Npcf_BDTPolicyControl Service offered by the PCF
//"npcf-eventexposure"	Npcf_EventExposure Service offered by the PCF
//"npcf-ue-policy-control"	Npcf_UEPolicyControl Service offered by the PCF
//"nsmsf-sms"	Nsmsf_SMService Service offered by the SMSF
//"nnssf-nsselection"	Nnssf_NSSelection Service offered by the NSSF
//"nnssf-nssaiavailability"	Nnssf_NSSAIAvailability Service offered by the NSSF
//"nudr-dr"	Nudr_DataRepository Service offered by the UDR
//"nlmf-loc"	Nlmf_Location Service offered by the LMF
//"n5g-eir-eic"	N5g-eir_EquipmentIdentityCheck Service offered by the 5G-EIR
//"nbsf-management"	Nbsf_Management Service offered by the BSF
//"nchf-spendinglimitcontrol"	Nchf_SpendingLimitControl Service offered by the CHF
//"nchf-convergedcharging"	Nchf_Converged_Charging Service offered by the CHF
//"nnwdaf-eventssubscription"	Nnwdaf_EventsSubscription Service offered by the NWDAF
//"nnwdaf-analyticsinfo"	Nnwdaf_AnalyticsInfo Service offered by the NWDAF
//NOTE:	The services defined in this table are those defined by 3GPP NFs in 5GC; however, in order to support custom services offered by standard and custom NFs, the NRF shall also accept the registration of NF Services with other service names.

//29.510 6.1.6.3.11
type ServiceName byte

const (
	NNRF_NFM ServiceName = iota
	NNRF_DISC
	NUDM_SDM
	NUDM_UECM
	NUDM_UEAU
	NUDM_EE
	NUDM_PP
	NAMF_COMM
	NAMF_EVTS
	NAMF_MT
	NAMF_LOC
	NSMF_PDUSESSION
	NSMF_EVENT_EXPOSURE
	NAUSF_AUTH
	NAUSF_SORPROTECTION
	NNEF_PFDMANAGEMENT
	NPCF_AM_POLICY_CONTROL
	NPCF_SMPOLICYCONTROL
	NPCF_POLICYAUTHORIZATION
	NPCF_BDTPOLICYCONTROL
	NPCF_EVENTEXPOSURE
	NPCF_UE_POLICY_CONTROL
	NSMSF_SMS
	NNSSF_NSSELECTION
	NNSSF_NSSAIAVAILABILITY
	NUDR_DR
	NLMF_LOC
	N5G_EIR_EIC
	NBSF_MANAGEMENT
	NCHF_SPENDINGLIMITCONTROL
	NCHF_CONVERGEDCHARGING
	NNWDAF_EVENTSSUBSCRIPTION
	NNWDAF_ANALYTICSINFO
)

func SetFromString(s string) ServiceName {
	switch s {
	case "NNRF_NFM":
		return NNRF_NFM
	case "NNRF_DISC":
		return NNRF_DISC
	case "NUDM_SDM":
		return NUDM_SDM
	case "NUDM_UECM":
		return NUDM_UECM
	case "NUDM_UEAU":
		return NUDM_UEAU
	case "NUDM_EE":
		return NUDM_EE
	case "NUDM_PP":
		return NUDM_PP
	case "NAMF_COMM":
		return NAMF_COMM
	case "NAMF_EVTS":
		return NAMF_EVTS
	case "NAMF_MT":
		return NAMF_MT
	case "NAMF_LOC":
		return NAMF_LOC
	case "NSMF_PDUSESSION":
		return NSMF_PDUSESSION
	case "NSMF_EVENT_EXPOSURE":
		return NSMF_EVENT_EXPOSURE
	case "NAUSF_AUTH":
		return NAUSF_AUTH
	case "NAUSF_SORPROTECTION":
		return NAUSF_SORPROTECTION
	case "NNEF_PFDMANAGEMENT":
		return NNEF_PFDMANAGEMENT
	case "NPCF_AM_POLICY_CONTROL":
		return NPCF_AM_POLICY_CONTROL
	case "NPCF_SMPOLICYCONTROL":
		return NPCF_SMPOLICYCONTROL
	case "NPCF_POLICYAUTHORIZATION":
		return NPCF_POLICYAUTHORIZATION
	case "NPCF_BDTPOLICYCONTROL":
		return NPCF_BDTPOLICYCONTROL
	case "NPCF_EVENTEXPOSURE":
		return NPCF_EVENTEXPOSURE
	case "NPCF_UE_POLICY_CONTROL":
		return NPCF_UE_POLICY_CONTROL
	case "NSMSF_SMS":
		return NSMSF_SMS
	case "NNSSF_NSSELECTION":
		return NNSSF_NSSELECTION
	case "NNSSF_NSSAIAVAILABILITY":
		return NNSSF_NSSAIAVAILABILITY
	case "NUDR_DR":
		return NUDR_DR
	case "NLMF_LOC":
		return NLMF_LOC
	case "N5G_EIR_EIC":
		return N5G_EIR_EIC
	case "NBSF_MANAGEMENT":
		return NBSF_MANAGEMENT
	case "NNWDAF_EVENTSSUBSCRIPTION":
		return NNWDAF_EVENTSSUBSCRIPTION
	case "NCHF_SPENDINGLIMITCONTROL":
		return NCHF_SPENDINGLIMITCONTROL
	case "NCHF_CONVERGEDCHARGING":
		return NCHF_CONVERGEDCHARGING
	case "NNWDAF_ANALYTICSINFO":
		return NNWDAF_ANALYTICSINFO
	}
	return NNRF_NFM
}

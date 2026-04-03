package n11msg

import (
	T "lite5gc/cmn/types3gpp"
	"time"
)

type RequestType byte

const (
	INITIAL_REQUEST RequestType = iota + 1
	EXISTING_PDU_SESSION
	INITIAL_EMERGENCY_REQUEST
	EXISTING_EMERGENCY_PDU_SESSION
	MODIFY_REQUEST
)

type PresenceState byte

const (
	PresenceState_IN_AREA PresenceState = iota + 1
	PresenceState_OUT_OF_AREA
	PresenceState_UNKNOWN
	PresenceState_INACTIVE
)

type HoState byte

const (
	PresenceState_NONE PresenceState = iota + 1
	PresenceState_PREPARING
	PresenceState_PREPARED
	PresenceState_COMPLETED
	PresenceState_CANCELLED
)

type DnnSelectionMode byte

const (
	VERIFIED DnnSelectionMode = iota + 1
	UE_DNN_NOT_VERIFIED
	NW_DNN_NOT_VERIFIED
)

type EpsInterworkingIndication byte

const (
	Eps_NONE EpsInterworkingIndication = iota + 1
	Eps_WITH_N26
	Eps_WITHOUT_N26
)

type UserLocation struct {
	//eutraLocation	EutraLocation
	NrLocation NrLocation
	//n3gaLocation	N3gaLocation
}

type NrLocation struct {
	Tai                      T.TAI
	Ncgi                     Ncgi
	AgeOfLocationInformation int
	UeLocationTimestamp      time.Time
	GeographicalInformation  string
	GeodeticInformation      string
	GlobalGnbId              int
}

type Ncgi struct {
	PlmnId   T.PlmnID
	NrCellId string
}

//3GPP TS 29.502V15.1.0(2018-09)
type UpCnxState byte

const (
	INAVLID UpCnxState = iota
	ACTIVATED
	DEACTIVATED
	ACTIVATING
)

func (p UpCnxState) String() string {
	switch p {
	case ACTIVATED:
		return "activated"
	case DEACTIVATED:
		return "deactivate"
	case ACTIVATING:
		return "activating"
	default:
		return "invalid"
	}
}

//type IeId uint
const (
	Ieid_supi uint = iota
	Ieid_unauthenticatedSupi
	Ieid_pei
	Ieid_gpsi
	Ieid_dnn
	Ieid_sNssai
	Ieid_hplmnSnssai
	Ieid_servingNfId
	Ieid_guami
	Ieid_serviceName
	Ieid_servingNetwork
	Ieid_requestType
	Ieid_n1SmMsg
	Ieid_anType
	Ieid_ratType
	Ieid_presenceInLadn
	Ieid_ueLocation
	Ieid_ueTimeZone
	Ieid_addUeLocation
	Ieid_smContextStatusUri
	Ieid_hSmfUri
	Ieid_oldPduSessionId
	Ieid_pduSessionsActivateList
	Ieid_ueEpsPdnConnection
	Ieid_hoState
	Ieid_additionalHsmfUri
	Ieid_pcfId
	Ieid_nrfUri
	Ieid_supportedFeatures
	Ieid_selMode
	Ieid_backupAmfInfo
	Ieid_traceData
	Ieid_epsInterworkingInd
	Ieid_udmGroupId
	Ieid_RoutingIndicator
	Ieid_IndirectForwardingFlag

	Ieid_upCnxState
	Ieid_n2SmInfo
	Ieid_RanInfo
	Ieid_n2SmInfoType
	Ieid_allocatedEbiList
	Ieid_n1MessageContainer
	Ieid_n2InfoContainer
	Ieid_n2InfoContent
	Ieid_skipInd
	Ieid_lastMsgIndication
	Ieid_pdusessionId
	Ieid_ppi
	Ieid_arp
	Ieid_qi5
	Ieid_n1n2FailureTxfNotifURI
	Ieid_cause
	Ieid_toBeSwitched
	Ieid_failedToBeSwitched
	Ieid_targetServingNfId
	Ieid_dataForwarding
	Ieid_epsBearerSetup
	Ieid_revokeEbiList
	Ieid_release
	Ieid_ngApCause
	Ieid_ngApIeType
	Ieid_Mm5GCauseValue
	Ieid_releaseEbiList
	Ieid_modifiedEbiList
	Ieid_vsmfReleaseOnly
	Ieid_statusInfo
	Ieid_maxvalue
)

type N2SmInfoType string

const (
	PduResSetupReq  N2SmInfoType = "PDU_RES_SETUP_REQ"
	PduResSetupRsp  N2SmInfoType = "PDU_RES_SETUP_RSP"
	PduResSetupFail N2SmInfoType = "PDU_RES_SETUP_FAIL"

	PduResRelCmd N2SmInfoType = "PDU_RES_REL_CMD"
	PduResRelRsp N2SmInfoType = "PDU_RES_REL_RSP"

	PduResModReq  N2SmInfoType = "PDU_RES_MOD_REQ"
	PduResModRsp  N2SmInfoType = "PDU_RES_MOD_RSP"
	PduResModFail N2SmInfoType = "PDU_RES_MOD_FAIL"

	PduResNty            N2SmInfoType = "PDU_RES_NTY"
	PduResNtyRel         N2SmInfoType = "PDU_RES_NTY_REL"
	PduResModInd         N2SmInfoType = "PDU_RES_MOD_IND"
	PduResModCfm         N2SmInfoType = "PDU_RES_MOD_CFM"
	PathSwitchReq        N2SmInfoType = "PATH_SWITCH_REQ"
	PathSwitchSetupFail  N2SmInfoType = "PATH_SWITCH_SETUP_FAIL"
	PathSwitchReqAck     N2SmInfoType = "PATH_SWITCH_REQ_ACK"
	PathSwitchReqFail    N2SmInfoType = "PATH_SWITCH_REQ_FAIL"
	HandoverRequired     N2SmInfoType = "HANDOVER_REQUIRED"
	HandoverCmd          N2SmInfoType = "HANDOVER_CMD"
	HandoverPrepFail     N2SmInfoType = "HANDOVER_PREP_FAIL"
	HandoverReqAck       N2SmInfoType = "HANDOVER_REQ_ACK"
	HandoverResAllocFail N2SmInfoType = "HANDOVER_RES_ALLOC_FAIL"
)

type NgapIeType string

const (
	NgapPduResSetupReq          NgapIeType = "PDU_RES_SETUP_REQ"
	NgapPduResRelCmd            NgapIeType = "PDU_RES_REL_CMD"
	NgapPduResMdReq             NgapIeType = "PDU_RES_MOD_REQ"
	NgapHandoverCmd             NgapIeType = "HANDOVER_CMD"
	NgapHandoverRequired        NgapIeType = "HANDOVER_REQUIRED"
	NgapHandoverPrepFail        NgapIeType = "HANDOVER_PREP_FAIL"
	NgapSrcToTarContainer       NgapIeType = "SRC_TO_TAR_CONTAINER"
	NgapTarTpSrcContainer       NgapIeType = "TAR_TO_SRC_CONTAINER"
	NgapRanStatusTransContainer NgapIeType = "RAN_STATUS_TRANS_CONTAINER"
	NgapSonConfigTransfer       NgapIeType = "SON_CONFIG_TRANSFER"
	NgapNrppaPdu                NgapIeType = "NR_PPA_PDU"
	NgapUeRadioCapability       NgapIeType = "UE_RADIO_CAPABILITY"
)

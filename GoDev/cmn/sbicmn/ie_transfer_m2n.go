package sbicmn

import (
	"lite5gc/cmn/message/n11msg"
	"lite5gc/cmn/message/nasmsg"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types3gpp"
	"lite5gc/openapi/models"
)

func Trans_ModelsToN11_RequestType(modelsData models.RequestType) (n11Data n11msg.RequestType) {
	rlogger.FuncEntry(types.ModCmn, nil)
	switch modelsData {
	case models.RequestType_INITIAL_REQUEST:
		n11Data = n11msg.INITIAL_REQUEST
	case models.RequestType_EXISTING_PDU_SESSION:
		n11Data = n11msg.EXISTING_PDU_SESSION
	case models.RequestType_INITIAL_EMERGENCY_REQUEST:
		n11Data = n11msg.INITIAL_EMERGENCY_REQUEST
	case models.RequestType_EXISTING_EMERGENCY_PDU_SESSION:
		n11Data = n11msg.EXISTING_EMERGENCY_PDU_SESSION
	}
	return n11Data
}

func Trans_ModelsToN11_AnType(modelsData models.AccessType) (n11Data nasmsg.RegistrationResult) {
	rlogger.FuncEntry(types.ModCmn, nil)
	switch modelsData {
	case models.AccessType__3_GPP_ACCESS:
		n11Data = nasmsg.Access3gpp
	case models.AccessType_NON_3_GPP_ACCESS:
		n11Data = nasmsg.AccessNon3gpp
	}
	return n11Data
}

func Trans_ModelsToN11_RatType(modelsData models.RatType) (n11Data types3gpp.RatType) {
	rlogger.FuncEntry(types.ModCmn, nil)
	switch modelsData {
	case models.RatType_NR:
		n11Data = types3gpp.RatType_NR
	case models.RatType_EUTRA:
		n11Data = types3gpp.RatType_EUTRA
	case models.RatType_WLAN:
		n11Data = types3gpp.RatType_WLAN
	case models.RatType_VIRTUAL:
		n11Data = types3gpp.RatType_VIRTUAL
	}
	return n11Data
}

func Trans_ModelsToN11_SelMode(modelsData models.DnnSelectionMode) (n11Data n11msg.DnnSelectionMode) {
	rlogger.FuncEntry(types.ModCmn, nil)
	switch modelsData {
	case models.DnnSelectionMode_VERIFIED:
		n11Data = n11msg.VERIFIED
	case models.DnnSelectionMode_UE_DNN_NOT_VERIFIED:
		n11Data = n11msg.UE_DNN_NOT_VERIFIED
	case models.DnnSelectionMode_NW_DNN_NOT_VERIFIED:
		n11Data = n11msg.NW_DNN_NOT_VERIFIED
	}
	return n11Data
}

func Trans_ModelsToN11_N2SmInfoType(modelsData models.N2SmInfoType) (n11Data n11msg.N2SmInfoType) {
	rlogger.FuncEntry(types.ModCmn, nil)
	switch modelsData {
	case models.N2SmInfoType_PDU_RES_SETUP_REQ:
		n11Data = n11msg.PduResSetupReq
	case models.N2SmInfoType_PDU_RES_SETUP_RSP:
		n11Data = n11msg.PduResSetupRsp
	case models.N2SmInfoType_PDU_RES_SETUP_FAIL:
		n11Data = n11msg.PduResSetupFail
	case models.N2SmInfoType_PDU_RES_REL_CMD:
		n11Data = n11msg.PduResRelCmd
	case models.N2SmInfoType_PDU_RES_REL_RSP:
		n11Data = n11msg.PduResRelRsp
	case models.N2SmInfoType_PDU_RES_MOD_REQ:
		n11Data = n11msg.PduResModReq
	case models.N2SmInfoType_PDU_RES_MOD_RSP:
		n11Data = n11msg.PduResModRsp
	case models.N2SmInfoType_PDU_RES_MOD_FAIL:
		n11Data = n11msg.PduResModFail
	case models.N2SmInfoType_PDU_RES_NTY:
		n11Data = n11msg.PduResNty
	case models.N2SmInfoType_PDU_RES_NTY_REL:
		n11Data = n11msg.PduResNtyRel
	case models.N2SmInfoType_PDU_RES_MOD_IND:
		n11Data = n11msg.PduResModInd
	case models.N2SmInfoType_PDU_RES_MOD_CFM:
		n11Data = n11msg.PduResModCfm
	case models.N2SmInfoType_PATH_SWITCH_REQ:
		n11Data = n11msg.PathSwitchReq
	case models.N2SmInfoType_PATH_SWITCH_SETUP_FAIL:
		n11Data = n11msg.PathSwitchSetupFail
	case models.N2SmInfoType_PATH_SWITCH_REQ_ACK:
		n11Data = n11msg.PathSwitchReqAck
	case models.N2SmInfoType_PATH_SWITCH_REQ_FAIL:
		n11Data = n11msg.PathSwitchReqFail
	case models.N2SmInfoType_HANDOVER_REQUIRED:
		n11Data = n11msg.HandoverRequired
	case models.N2SmInfoType_HANDOVER_CMD:
		n11Data = n11msg.HandoverCmd
	case models.N2SmInfoType_HANDOVER_PREP_FAIL:
		n11Data = n11msg.HandoverPrepFail
	case models.N2SmInfoType_HANDOVER_REQ_ACK:
		n11Data = n11msg.HandoverReqAck
	case models.N2SmInfoType_HANDOVER_RES_ALLOC_FAIL:
		n11Data = n11msg.HandoverResAllocFail
	}
	return n11Data
}

func Trans_ModelsToN11_UpCnxState(modelsData models.UpCnxState) (n11Data n11msg.UpCnxState) {
	rlogger.FuncEntry(types.ModCmn, nil)
	rlogger.Trace(types.ModCmn, rlogger.DEBUG, nil, "models UpCnxState(%s)", modelsData)

	switch modelsData {
	case models.UpCnxState_ACTIVATED:
		n11Data = n11msg.ACTIVATED
	case models.UpCnxState_DEACTIVATED:
		n11Data = n11msg.DEACTIVATED
	case models.UpCnxState_ACTIVATING:
		n11Data = n11msg.ACTIVATING
	default:
		n11Data = n11msg.INAVLID
	}
	return n11Data
}

func Trans_ModelsToN11_N1MessageClass(modelsData models.N1MessageClass) (n11Data n11msg.N1MessageClass) {
	rlogger.FuncEntry(types.ModCmn, nil)
	switch modelsData {
	case models.N1MessageClass__5_GMM:
		n11Data = n11msg.FiveGMM
	case models.N1MessageClass_SM:
		n11Data = n11msg.SM_N1Info
	case models.N1MessageClass_LPP:
		n11Data = n11msg.LPP
	case models.N1MessageClass_SMS:
		n11Data = n11msg.SMS
	case models.N1MessageClass_UPDP:
	}
	return n11Data
}

func Trans_ModelsToN11_NgapIeType(modelsData models.NgapIeType) (n11Data n11msg.NgapIeType) {
	rlogger.FuncEntry(types.ModCmn, nil)
	switch modelsData {
	case models.NgapIeType_PDU_RES_SETUP_REQ:
		n11Data = n11msg.NgapPduResSetupReq
	case models.NgapIeType_PDU_RES_REL_CMD:
		n11Data = n11msg.NgapPduResRelCmd
	case models.NgapIeType_PDU_RES_MOD_REQ:
		n11Data = n11msg.NgapPduResMdReq
	case models.NgapIeType_HANDOVER_CMD:
		n11Data = n11msg.NgapHandoverCmd
	case models.NgapIeType_HANDOVER_REQUIRED:
		n11Data = n11msg.NgapHandoverRequired
	case models.NgapIeType_HANDOVER_PREP_FAIL:
		n11Data = n11msg.NgapHandoverPrepFail
	case models.NgapIeType_SRC_TO_TAR_CONTAINER:
		n11Data = n11msg.NgapSrcToTarContainer
	case models.NgapIeType_TAR_TO_SRC_CONTAINER:
		n11Data = n11msg.NgapTarTpSrcContainer
	case models.NgapIeType_RAN_STATUS_TRANS_CONTAINER:
		n11Data = n11msg.NgapRanStatusTransContainer
	case models.NgapIeType_SON_CONFIG_TRANSFER:
		n11Data = n11msg.NgapSonConfigTransfer
	case models.NgapIeType_NRPPA_PDU:
		n11Data = n11msg.NgapNrppaPdu
	case models.NgapIeType_UE_RADIO_CAPABILITY:
		n11Data = n11msg.NgapUeRadioCapability
	}
	return n11Data
}

func Trans_ModelsToN11_N2InformationClass(modelsData models.N2InformationClass) (n11Data n11msg.N2InformationClass) {
	rlogger.FuncEntry(types.ModCmn, nil)
	switch modelsData {
	case models.N2InformationClass_SM:
		n11Data = n11msg.SM_N2Info
	case models.N2InformationClass_NRP_PA:
		n11Data = n11msg.NRPPa
	case models.N2InformationClass_PWS:
		n11Data = n11msg.PWS
	case models.N2InformationClass_PWS_BCAL:
		n11Data = n11msg.PWS_BCAL
	case models.N2InformationClass_PWS_RF:
		n11Data = n11msg.PWS_RF
	case models.N2InformationClass_RAN:
		n11Data = n11msg.RAN
	}
	return n11Data
}

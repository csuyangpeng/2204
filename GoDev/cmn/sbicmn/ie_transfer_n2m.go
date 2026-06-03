package sbicmn

import (
	"lite5gc/cmn/message/n11msg"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/openapi/models"
)

func Trans_N11ToModels_N1MessageClass(n11Data n11msg.N1MessageClass) (modelsData models.N1MessageClass) {
	rlogger.FuncEntry(types.ModCmn, nil)
	switch n11Data {
	case n11msg.FiveGMM:
		modelsData = models.N1MessageClass__5_GMM
	case n11msg.SM_N1Info:
		modelsData = models.N1MessageClass_SM
	case n11msg.LPP:
		modelsData = models.N1MessageClass_LPP
	case n11msg.SMS:
		modelsData = models.N1MessageClass_SMS
	}
	return modelsData
}

func Trans_N11ToModels_NgapIeType(n11Data n11msg.NgapIeType) (modelsData models.NgapIeType) {
	rlogger.FuncEntry(types.ModCmn, nil)
	switch n11Data {
	case n11msg.NgapPduResSetupReq:
		modelsData = models.NgapIeType_PDU_RES_SETUP_REQ
	case n11msg.NgapPduResRelCmd:
		modelsData = models.NgapIeType_PDU_RES_REL_CMD
	case n11msg.NgapPduResMdReq:
		modelsData = models.NgapIeType_PDU_RES_MOD_REQ
	case n11msg.NgapHandoverCmd:
		modelsData = models.NgapIeType_HANDOVER_CMD
	case n11msg.NgapHandoverRequired:
		modelsData = models.NgapIeType_HANDOVER_REQUIRED
	case n11msg.NgapHandoverPrepFail:
		modelsData = models.NgapIeType_HANDOVER_PREP_FAIL
	case n11msg.NgapSrcToTarContainer:
		modelsData = models.NgapIeType_SRC_TO_TAR_CONTAINER
	case n11msg.NgapTarTpSrcContainer:
		modelsData = models.NgapIeType_TAR_TO_SRC_CONTAINER
	case n11msg.NgapRanStatusTransContainer:
		modelsData = models.NgapIeType_RAN_STATUS_TRANS_CONTAINER
	case n11msg.NgapSonConfigTransfer:
		modelsData = models.NgapIeType_SON_CONFIG_TRANSFER
	case n11msg.NgapNrppaPdu:
		modelsData = models.NgapIeType_NRPPA_PDU
	case n11msg.NgapUeRadioCapability:
		modelsData = models.NgapIeType_UE_RADIO_CAPABILITY
	}
	return modelsData
}

func Trans_N11ToModels_N2InformationClass(n11Data n11msg.N2InformationClass) (modelsData models.N2InformationClass) {
	rlogger.FuncEntry(types.ModCmn, nil)
	rlogger.Trace(types.ModCmn, rlogger.DEBUG, nil, "N2InformationClass: ", n11Data)
	switch n11Data {
	case n11msg.SM_N2Info:
		modelsData = models.N2InformationClass_SM
	case n11msg.NRPPa:
		modelsData = models.N2InformationClass_NRP_PA
	case n11msg.PWS:
		modelsData = models.N2InformationClass_PWS
	case n11msg.PWS_BCAL:
		modelsData = models.N2InformationClass_PWS_BCAL
	case n11msg.PWS_RF:
		modelsData = models.N2InformationClass_PWS_RF
	case n11msg.RAN:
		modelsData = models.N2InformationClass_RAN
	}
	return modelsData
}

func Trans_N11ToModels_N2SmInfoType(n11Data n11msg.N2SmInfoType) (modelsData models.N2SmInfoType) {
	rlogger.FuncEntry(types.ModCmn, nil)
	switch n11Data {
	case n11msg.PduResSetupReq:
		modelsData = models.N2SmInfoType_PDU_RES_SETUP_REQ
	case n11msg.PduResSetupRsp:
		modelsData = models.N2SmInfoType_PDU_RES_SETUP_RSP
	case n11msg.PduResSetupFail:
		modelsData = models.N2SmInfoType_PDU_RES_SETUP_FAIL
	case n11msg.PduResRelCmd:
		modelsData = models.N2SmInfoType_PDU_RES_REL_CMD
	case n11msg.PduResRelRsp:
		modelsData = models.N2SmInfoType_PDU_RES_REL_RSP
	case n11msg.PduResModReq:
		modelsData = models.N2SmInfoType_PDU_RES_MOD_REQ
	case n11msg.PduResModRsp:
		modelsData = models.N2SmInfoType_PDU_RES_MOD_RSP
	case n11msg.PduResModFail:
		modelsData = models.N2SmInfoType_PDU_RES_MOD_FAIL
	case n11msg.PduResNty:
		modelsData = models.N2SmInfoType_PDU_RES_NTY
	case n11msg.PduResNtyRel:
		modelsData = models.N2SmInfoType_PDU_RES_NTY_REL
	case n11msg.PduResModInd:
		modelsData = models.N2SmInfoType_PDU_RES_MOD_IND
	case n11msg.PduResModCfm:
		modelsData = models.N2SmInfoType_PDU_RES_MOD_CFM
	case n11msg.PathSwitchReq:
		modelsData = models.N2SmInfoType_PATH_SWITCH_REQ
	case n11msg.PathSwitchSetupFail:
		modelsData = models.N2SmInfoType_PATH_SWITCH_SETUP_FAIL
	case n11msg.PathSwitchReqAck:
		modelsData = models.N2SmInfoType_PATH_SWITCH_REQ_ACK
	case n11msg.PathSwitchReqFail:
		modelsData = models.N2SmInfoType_PATH_SWITCH_REQ_FAIL
	case n11msg.HandoverRequired:
		modelsData = models.N2SmInfoType_HANDOVER_REQUIRED
	case n11msg.HandoverCmd:
		modelsData = models.N2SmInfoType_HANDOVER_CMD
	case n11msg.HandoverPrepFail:
		modelsData = models.N2SmInfoType_HANDOVER_PREP_FAIL
	case n11msg.HandoverReqAck:
		modelsData = models.N2SmInfoType_HANDOVER_REQ_ACK
	case n11msg.HandoverResAllocFail:
		modelsData = models.N2SmInfoType_HANDOVER_RES_ALLOC_FAIL
	}
	return modelsData
}

func Trans_N11ToModels_UpCnxState(n11Data n11msg.UpCnxState) (modelsData models.UpCnxState) {
	rlogger.FuncEntry(types.ModCmn, nil)
	switch n11Data {
	case n11msg.INAVLID:
	case n11msg.ACTIVATED:
		modelsData = models.UpCnxState_ACTIVATED
	case n11msg.DEACTIVATED:
		modelsData = models.UpCnxState_DEACTIVATED
	case n11msg.ACTIVATING:
		modelsData = models.UpCnxState_ACTIVATING
	}
	return modelsData
}

func Trans_N11ToModels_RequestType(n11Data n11msg.RequestType) (modelsData models.RequestType) {
	rlogger.FuncEntry(types.ModCmn, nil)
	switch n11Data {
	case n11msg.INITIAL_REQUEST:
		modelsData = models.RequestType_INITIAL_REQUEST
	case n11msg.EXISTING_PDU_SESSION:
		modelsData = models.RequestType_EXISTING_PDU_SESSION
	case n11msg.INITIAL_EMERGENCY_REQUEST:
		modelsData = models.RequestType_INITIAL_EMERGENCY_REQUEST
	case n11msg.EXISTING_EMERGENCY_PDU_SESSION:
		modelsData = models.RequestType_EXISTING_EMERGENCY_PDU_SESSION
	}
	return modelsData
}

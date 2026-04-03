package sbicmn

import (
	"encoding/hex"
	"lite5gc/cmn/message/n11msg"
	"lite5gc/cmn/nas/nasie"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/openapi/models"
)

//amf send Nsmf msg, trans n11Data to modelsData
func Trans_N11ToModels_SmContextReleaseDataFormat(n11Data n11msg.ReleaseSMContextRequestData) (modelsData models.SmContextReleaseData) {
	rlogger.FuncEntry(types.ModCmn, nil)
	modelsData.NgApCause = &models.NgApCause{}
	modelsData.NgApCause.Value = int32(n11Data.NgApCause.Value)
	modelsData.NgApCause.Group = int32(n11Data.NgApCause.Type)
	return modelsData
}

//amf send Nsmf msg, trans n11Data to modelsData
func Trans_N11ToModels_SmContextCreateDataFormat(n11Data n11msg.SmContextCreateData) (modelsData models.SmContextCreateData) {
	rlogger.FuncEntry(types.ModCmn, nil)

	modelsData.Supi = n11Data.Supi.String()
	modelsData.PduSessionId = int32(n11Data.PduSessionId)
	modelsData.Dnn = n11Data.Dnn.String()

	snssai := models.Snssai{}
	snssai.Sst = int32(n11Data.SNssai.Sst)
	if n11Data.SNssai.Ind == nasie.SstSd {
		//sd := []byte{}
		//copy(sd[:], n11Data.SNssai.Sd[:])
		//snssai.Sd = hex.EncodeToString(sd)
		snssai.Sd = hex.EncodeToString(n11Data.SNssai.Sd[:])
	}
	modelsData.SNssai = &snssai
	rlogger.Trace(types.ModCmn, rlogger.INFO, nil, "snssai:%v", snssai)

	//modelsData.ServingNfId = configure.AmfConf.Service.AmfInstanceId
	modelsData.Guami = &models.Guami{}
	modelsData.Guami.AmfId = n11Data.Guami.AmfId.String()
	modelsData.Guami.PlmnId = &models.PlmnId{}
	modelsData.Guami.PlmnId.Mcc = n11Data.Guami.PlmnId.GetMccString()
	modelsData.Guami.PlmnId.Mnc = n11Data.Guami.PlmnId.GetMncString()
	modelsData.ServingNetwork = &models.PlmnId{}
	modelsData.ServingNetwork.Mnc = n11Data.ServingNetwork.GetMncString()
	modelsData.ServingNetwork.Mcc = n11Data.ServingNetwork.GetMccString()

	modelsData.RequestType = models.RequestType_INITIAL_REQUEST

	modelsData.AnType = models.AccessType__3_GPP_ACCESS

	modelsData.N1SmMsg = &models.RefToBinaryData{}

	modelsData.N1SmMsg.ContentId = hex.EncodeToString(n11Data.N1SmMsg)
	modelsData.RatType = models.RatType_NR
	//modelsData.UeLocation
	//modelsData.SmContextStatusUri
	modelsData.SelMode = models.DnnSelectionMode_VERIFIED
	return modelsData
}

//amf send Nsmf msg, trans n11Data to modelsData
func Trans_N11ToModels_SmContextModifyDataFormat(n11Data n11msg.UpdateSMContextRequestData) (data models.UpdateSmContextRequest) {
	rlogger.FuncEntry(types.ModCmn, nil)

	rlogger.Trace(types.ModCmn, rlogger.ERROR, nil, "n11Data(%v)", n11Data)
	modelsData := models.SmContextUpdateData{}

	if n11Data.IeFlags.Test(n11msg.Ieid_n2SmInfoType) {
		modelsData.N2SmInfoType = Trans_N11ToModels_N2SmInfoType(n11Data.N2SmInfoType)
	}

	if n11Data.IeFlags.Test(n11msg.Ieid_n2SmInfo) {
		modelsData.N2SmInfo = &models.RefToBinaryData{ContentId: "ngapie"}
		data.BinaryDataN2SmInformation = []byte(n11Data.N2SmInfo)
		rlogger.Trace(types.ModCmn, rlogger.INFO, nil, "data n2 info:(%v)", data.BinaryDataN2SmInformation)
	}
	if n11Data.IeFlags.Test(n11msg.Ieid_upCnxState) {
		modelsData.UpCnxState = Trans_N11ToModels_UpCnxState(n11Data.UpCnxState)
	}

	if n11Data.IeFlags.Test(n11msg.Ieid_release) {
		modelsData.Release = n11Data.Release
	}

	if n11Data.IeFlags.Test(n11msg.Ieid_n1SmMsg) {
		modelsData.N1SmMsg = &models.RefToBinaryData{ContentId: "n1msg"}
		data.BinaryDataN1SmMessage = n11Data.N1SmMsg
		rlogger.Trace(types.ModCmn, rlogger.INFO, nil, "data n1 msg:(%v)", data.BinaryDataN1SmMessage)
	}
	data.JsonData = &modelsData
	rlogger.Trace(types.ModCmn, rlogger.ERROR, nil, "modelsData(%v)", modelsData)

	return data
}

//func Trans_N11ToModels_SmContextUpdatedDataReq(n11Data *n11msg.UpdateSMContextResponseData) (*models.SmContextUpdatedData) {
func Trans_N11ToModels_SmContextUpdatedDataReq(n11Data *n11msg.UpdateSMContextResponseData) *models.UpdateSmContextResponse {
	rlogger.FuncEntry(types.ModCmn, nil)

	respData := &models.UpdateSmContextResponse{}
	respData.JsonData = &models.SmContextUpdatedData{}

	if n11Data.IeFlags.Test(n11msg.Ieid_n1SmMsg) {
		respData.JsonData.N1SmMsg = &models.RefToBinaryData{}
		respData.JsonData.N1SmMsg.ContentId = "n1msg"
		respData.BinaryDataN1SmMessage = []byte(n11Data.N1SmMsg)
		rlogger.Trace(types.ModCmn, rlogger.DEBUG, nil, "respData.BinaryDataN1SmMessage: (%v)", respData.BinaryDataN1SmMessage)
	}

	if n11Data.IeFlags.Test(n11msg.Ieid_n2SmInfoType) {
		respData.JsonData.N2SmInfoType = Trans_N11ToModels_N2SmInfoType(n11Data.N2SmInfoType)
		if n11Data.IeFlags.Test(n11msg.Ieid_n2SmInfo) {
			respData.JsonData.N2SmInfo = &models.RefToBinaryData{}
			respData.JsonData.N2SmInfo.ContentId = "ngapie"
			respData.BinaryDataN2SmInformation = []byte(n11Data.N2SmInfo)
		}
	}

	if n11Data.IeFlags.Test(n11msg.Ieid_upCnxState) {
		respData.JsonData.UpCnxState = Trans_N11ToModels_UpCnxState(n11Data.UpCnxState)
		rlogger.Trace(types.ModCmn, rlogger.DEBUG, nil, "UpCnxState: (%v)", respData.JsonData.UpCnxState)
	}

	return respData
}

func Trans_N11ToModels_N1N2MsgTransferReqFormat(n11Data *n11msg.N1N2MessageTransferReqData) *models.N1N2MessageTransferReqData {
	rlogger.FuncEntry(types.ModCmn, nil)

	modelData := &models.N1N2MessageTransferReqData{}

	if n11Data.IeFlags.Test(n11msg.Ieid_n1MessageContainer) {
		modelData.N1MessageContainer = &models.N1MessageContainer{}
		modelData.N1MessageContainer.N1MessageContent = &models.RefToBinaryData{}

		modelData.N1MessageContainer.N1MessageContent.ContentId = "n1msg"
		modelData.N1MessageContainer.N1MessageClass =
			Trans_N11ToModels_N1MessageClass(n11Data.N1MessageContainer.N1MsgClass)
	}

	if n11Data.IeFlags.Test(n11msg.Ieid_pdusessionId) {
		modelData.PduSessionId = int32(n11Data.SessionId)
	}

	if n11Data.IeFlags.Test(n11msg.Ieid_n2InfoContainer) {
		modelData.N2InfoContainer = &models.N2InfoContainer{}
		modelData.N2InfoContainer.SmInfo = &models.N2SmInformation{}
		modelData.N2InfoContainer.SmInfo.N2InfoContent = &models.N2InfoContent{}
		modelData.N2InfoContainer.SmInfo.N2InfoContent.NgapData = &models.RefToBinaryData{}

		if n11Data.N2InfoContainer.N2InforClass != 0 {
			modelData.N2InfoContainer.N2InformationClass = Trans_N11ToModels_N2InformationClass(n11Data.N2InfoContainer.N2InforClass)
		}

		if n11Data.N2InfoContainer.IeFlags.Test(n11msg.Ieid_n2SmInfo) {
			if n11Data.N2InfoContainer.SmInfo.IeFlags.Test(n11msg.Ieid_n2InfoContent) {
				modelData.N2InfoContainer.SmInfo.N2InfoContent.NgapData.ContentId = "ngapie"
				modelData.N2InfoContainer.SmInfo.N2InfoContent.NgapIeType =
					Trans_N11ToModels_NgapIeType(n11Data.N2InfoContainer.SmInfo.N2InfoCont.NgapIeType)
			}
		}

		if n11Data.N2InfoContainer.SmInfo.IeFlags.Test(n11msg.Ieid_pdusessionId) {
			modelData.N2InfoContainer.SmInfo.PduSessionId = int32(n11Data.N2InfoContainer.SmInfo.PduSessionId)
		}

	}

	if n11Data.IeFlags.Test(n11msg.Ieid_ppi) {
		modelData.Ppi = int32(n11Data.Ppi)
	}

	return modelData
}

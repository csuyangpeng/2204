package smsender

import (
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/sbicmn"
	"lite5gc/cmn/types"
	"lite5gc/openapi/models"
	"lite5gc/smf/smfcontext/gctxt"
	"net/http"
)

func SendSbiRespMsgSMContextCreatedData(pduSessCtxt *gctxt.PduSessContext,
	sbimsg *sbicmn.SbiHandlerMessage) error {
	rlogger.FuncEntry(types.ModuleSmfSM, nil)
	if pduSessCtxt == nil || sbimsg == nil {
		rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, nil, "invalid sbiMsg parameter")
		return types.ErrInputParaNil
	}
	postResp := models.PostSmContextsResponse{}
	postResp.JsonData = &models.SmContextCreatedData{}
	postResp.JsonData.PduSessionId = int32(pduSessCtxt.PduSessionId)
	respMsg := sbicmn.NewSbiHandlerResponseMessage()
	respMsg.HTTPResponse.Body = postResp
	respMsg.HTTPResponse.Status = http.StatusCreated
	sbimsg.ResponseChan <- respMsg
	return nil
}

func SendSbiRespMsgUpdateSmCtxtResponse(respData *models.UpdateSmContextResponse,
	sbimsg *sbicmn.SbiHandlerMessage) {
	rlogger.FuncEntry(types.ModuleSmfSM, nil)
	if sbimsg == nil {
		rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, nil, "invalid sbiMsg parameter")
		return
	}
	respMsg := sbicmn.NewSbiHandlerResponseMessage()
	respMsg.HTTPResponse.Body = respData
	respMsg.HTTPResponse.Status = http.StatusOK
	sbimsg.ResponseChan <- respMsg
	return
}

func SendSbiRespMsgUpdateSmCtxtResponseWithN1N2Info(respData *models.UpdateSmContextResponse,
	sbimsg *sbicmn.SbiHandlerMessage) error {
	rlogger.FuncEntry(types.ModuleSmfSM, nil)
	if respData == nil || sbimsg == nil {
		rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, nil, "invalid input parameter, respData is nil")
		return types.ErrInputParaNil
	}
	respMsg := sbicmn.NewSbiHandlerResponseMessage()
	respMsg.HTTPResponse.Body = respData
	respMsg.HTTPResponse.Status = http.StatusOK
	sbimsg.ResponseChan <- respMsg
	return nil
}

func SendSbiRespMsgReleaseSMContextResponse(sbimsg *sbicmn.SbiHandlerMessage) {
	rlogger.FuncEntry(types.ModuleSmfSM, nil)
	if sbimsg == nil {
		rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, nil, "invalid input parameter, respData is nil")
		return
	}
	respMsg := sbicmn.NewSbiHandlerResponseMessage()
	respMsg.HTTPResponse.Status = http.StatusNoContent
	sbimsg.ResponseChan <- respMsg
	return
}

func SendSbiReqMsgN1N2MsgTransferReq(msgData *models.N1N2MessageTransferReqData,
	sbimsg *sbicmn.SbiHandlerMessage) {
	rlogger.FuncEntry(types.ModuleSmfSM, nil)
	if msgData == nil || sbimsg == nil {
		rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, nil, "invalid input parameter, respData is nil")
		return
	}
	respMsg := &sbicmn.SbiHandlerResponseMessage{}
	respMsg.HTTPResponse.Body = msgData
	sbimsg.ResponseChan <- respMsg
	return
}

package mmsender

import (
	"context"
	"fmt"
	"lite5gc/cmn/message/n11msg"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/router"
	"lite5gc/cmn/routeragent"
	"lite5gc/cmn/sbicmn"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types3gpp"
	"lite5gc/openapi/models"
	"net/http"
)

func SendSbiRespMsgN1N2MessageTransferRspData(sbimsg *sbicmn.SbiHandlerMessage) error {
	rlogger.FuncEntry(types.ModuleAmfNas, nil)
	respData := models.N1N2MessageTransferRspData{}
	respData.Cause = models.N1N2MessageTransferCause_N1_N2_TRANSFER_INITIATED
	// todo send more data to amf
	respMsg := sbicmn.NewSbiHandlerResponseMessage()
	respMsg.HTTPResponse.Body = respData
	respMsg.HTTPResponse.Status = http.StatusOK
	rlogger.Trace(types.ModuleAmfNas, rlogger.INFO, nil, "respMsg(%v) ", respMsg)
	sbimsg.ResponseChan <- respMsg
	return nil
}

func SendReleaseSMCtxtSBIMsg(ctxt context.Context,
	psi nas.PduSessID,
	imsi types3gpp.Imsi,
	releaseSmCtxtReq n11msg.ReleaseSMContextRequestData) {
	rlogger.FuncEntry(types.ModuleAmfNas, nil)

	smMessage := sbicmn.SbiMessage{}
	var postsm = &sbicmn.SbiPostReleaseSmContext{}
	reqData := sbicmn.Trans_N11ToModels_SmContextReleaseDataFormat(releaseSmCtxtReq)
	postsm.ReqData = &reqData
	postsm.SmContextRef = imsi.AddIMSIPrefix() + "-" + fmt.Sprint(psi)
	rlogger.Trace(types.ModuleAmfNas, rlogger.INFO, nil, "psi:(%v), postsm.SmContextRef(%s)", psi, postsm.SmContextRef)
	smMessage.MsgData = postsm
	smMessage.MsgType = sbicmn.PduSessReleaseSMContextReq
	smMessage.ScInstId = 1
	err := routeragent.SendIpcMessage(ctxt, router.SbiGR, smMessage.ScInstId, &smMessage)
	if err != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, nil, "failed to send sbi msg to udm. error(%s)", err)
		return
	}
}

func SendUpdateSMCtxtSBIMsg(ctxt context.Context,
	psi nas.PduSessID,
	imsi types3gpp.Imsi,
	updateSmCtxtReq n11msg.UpdateSMContextRequestData) {
	rlogger.FuncEntry(types.ModuleAmfNas, nil)

	smMessage := sbicmn.SbiMessage{}
	var postsm = &sbicmn.SbiPostModifySmContext{}
	reqData := sbicmn.Trans_N11ToModels_SmContextModifyDataFormat(updateSmCtxtReq)
	postsm.ReqData = &reqData
	postsm.SmContextRef = imsi.AddIMSIPrefix() + "-" + fmt.Sprint(psi)
	rlogger.Trace(types.ModuleAmfNas, rlogger.INFO, nil, "psi:(%v), postsm.SmContextRef(%s)", psi, postsm.SmContextRef)
	smMessage.MsgData = postsm
	smMessage.MsgType = sbicmn.PduSessUpdateSMContextReq
	smMessage.ScInstId = 1
	err := routeragent.SendIpcMessage(ctxt, router.SbiGR, smMessage.ScInstId, &smMessage)
	if err != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, nil, "failed to send sbi msg to udm. error(%s)", err)
		return
	}
}

func SendCreateSMCtxtSBIMsg(ctxt context.Context, imsi types3gpp.Imsi, createSmCtxtReq n11msg.SmContextCreateData) {
	rlogger.FuncEntry(types.ModuleAmfNas, nil)

	smMessage := sbicmn.SbiMessage{}
	var postsm = &sbicmn.SbiPostCreateSmContext{}
	reqData := sbicmn.Trans_N11ToModels_SmContextCreateDataFormat(createSmCtxtReq)
	rlogger.Trace(types.ModuleAmfNas, rlogger.INFO, nil, "createSmCtxtReq msg:%v", createSmCtxtReq)
	rlogger.Trace(types.ModuleAmfNas, rlogger.INFO, nil, "reqData snssai msg:%v", reqData.SNssai)
	rlogger.Trace(types.ModuleAmfNas, rlogger.INFO, nil, "reqData msg:%v", reqData)
	postsm.ReqData = &reqData
	postsm.Supi = imsi.AddIMSIPrefix()
	smMessage.MsgData = postsm
	smMessage.MsgType = sbicmn.PduSessCreateSMContextReq
	smMessage.ScInstId = 1
	err := routeragent.SendIpcMessage(ctxt, router.SbiGR, smMessage.ScInstId, &smMessage)
	if err != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, nil, "failed to send sbi msg to udm. error(%s)", err)
		return
	}
}

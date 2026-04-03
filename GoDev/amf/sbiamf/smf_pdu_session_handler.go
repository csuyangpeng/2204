package sbiamf

import (
	"fmt"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/sbicmn"
	"lite5gc/cmn/types"
	"lite5gc/openapi"
	"lite5gc/openapi/models"
)

func (p *SbiAmf) HandlePostCreateSmContext(scid uint32, msg *sbicmn.SbiPostCreateSmContext) error {
	rlogger.FuncEntry(types.ModuleAmfSbi, nil)
	//invoke the api here
	postSmContexts := models.PostSmContextsRequest{JsonData: msg.ReqData}
	rlogger.Trace(types.ModuleAmfSbi, rlogger.INFO, nil, "request SmContextCreateData:%v", msg.ReqData)

	resp, res, err := p.nsmfPDUSessionApiClt.SMContextsCollectionApi.PostSmContexts(nil, postSmContexts)
	rlogger.Trace(types.ModuleAmfSbi, rlogger.INFO, nil, "get smf sbi resp(%v,%v,%v)", resp, res, err)
	if err != nil {
		var problemDetails models.ProblemDetails
		problemDetails.Cause = err.(openapi.GenericOpenAPIError).Model().(models.ProblemDetails).Cause
		rlogger.Trace(types.ModuleAmfSbi, rlogger.ERROR, nil,
			"failed to PostCreateSmContext, the http status code(%d), return problem is(%+v)",
			res.StatusCode,
			problemDetails)
		//udm_message.SendHttpResponseMessage(httpChannel, nil, res.StatusCode, problemDetails)
		return fmt.Errorf("failed to PostCreateSmContext, return (%+v)")
	}

	// send back
	msg.RespData = resp.JsonData
	respMsg := &sbicmn.SbiMessage{
		MsgType:  sbicmn.PduSessCreateSMContextResp,
		ScInstId: scid,
		MsgData:  msg}
	p.SendMsg2SC(scid, respMsg)

	return nil
}

func (p *SbiAmf) HandlePostModifySmContext(scid uint32, msg *sbicmn.SbiPostModifySmContext) error {
	rlogger.FuncEntry(types.ModuleAmfSbi, nil)
	rlogger.Trace(types.ModuleAmfSbi, rlogger.INFO, nil,
		"SmContextRef(%s)", msg.SmContextRef,
		"ReqData(%s)", msg.ReqData.String())
	resp, res, err := p.nsmfPDUSessionApiClt.IndividualSMContextApi.UpdateSmContext(nil,
		msg.SmContextRef,
		*msg.ReqData)
	rlogger.Trace(types.ModuleAmfSbi, rlogger.INFO, nil, "get udm sbi resp(%v,%v,%v)", resp, res, err)
	if err != nil {
		var problemDetails models.ProblemDetails
		problemDetails.Cause = err.(openapi.GenericOpenAPIError).Model().(models.ProblemDetails).Cause
		rlogger.Trace(types.ModuleAmfSbi, rlogger.ERROR, nil,
			"failed to PostModifySmContext, the http status code(%d), return problem is(%+v)",
			res.StatusCode,
			problemDetails)
		//udm_message.SendHttpResponseMessage(httpChannel, nil, res.StatusCode, problemDetails)
		return fmt.Errorf("failed to PostModifySmContext, return (%+v)")
	}

	// send back
	msg.RespData = &resp
	respMsg := &sbicmn.SbiMessage{
		MsgType:  sbicmn.PduSessUpdateSMContextResp,
		ScInstId: scid,
		MsgData:  msg}
	p.SendMsg2SC(scid, respMsg)
	return nil
}

func (p *SbiAmf) HandlePostReleaseSmContext(scid uint32, msg *sbicmn.SbiPostReleaseSmContext) error {
	rlogger.FuncEntry(types.ModuleAmfSbi, nil)
	//invoke the api here
	releaseSmContexts := models.ReleaseSmContextRequest{JsonData: msg.ReqData}
	rlogger.Trace(types.ModuleAmfSbi, rlogger.INFO, nil, "release SmContexts", msg.ReqData)

	res, err := p.nsmfPDUSessionApiClt.IndividualSMContextApi.ReleaseSmContext(nil,
		msg.SmContextRef,
		releaseSmContexts)
	rlogger.Trace(types.ModuleAmfSbi, rlogger.INFO, nil, "get udm sbi resp(%v,%v)", res, err)
	if err != nil {
		var problemDetails models.ProblemDetails
		problemDetails.Cause = err.(openapi.GenericOpenAPIError).Model().(models.ProblemDetails).Cause
		rlogger.Trace(types.ModuleAmfSbi, rlogger.ERROR, nil,
			"failed to PostReleaseSmContext, the http status code(%d), return problem is(%+v)",
			res.StatusCode,
			problemDetails)
		//udm_message.SendHttpResponseMessage(httpChannel, nil, res.StatusCode, problemDetails)
		return fmt.Errorf("failed to PostReleaseSmContext, return (%+v)")
	}

	// send back the
	respMsg := &sbicmn.SbiMessage{
		MsgType:  sbicmn.PduSessReleaseSMContextResp,
		ScInstId: scid,
		MsgData:  msg}
	p.SendMsg2SC(scid, respMsg)
	return nil
}

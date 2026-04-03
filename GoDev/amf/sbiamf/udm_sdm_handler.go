package sbiamf

import (
	"fmt"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/sbicmn"
	"lite5gc/cmn/types"
	"lite5gc/openapi"
	"lite5gc/openapi/models"
)

func (p *SbiAmf) HandlePostSdmSubscription(scid uint32, msg *sbicmn.SbiPostSdmSubscription) error {
	rlogger.FuncEntry(types.ModuleAmfSbi, nil)

	//invoke the api here
	resp, res, err := p.nudmSdmApiClt.SubscriptionCreationApi.Subscribe(nil, msg.Supi, *msg.Data)
	rlogger.Trace(types.ModuleAmfSbi, rlogger.INFO, nil, "get udm sbi resp(%v,%v,%v)", resp, res, err)
	if err != nil {
		var problemDetails models.ProblemDetails
		problemDetails.Cause = err.(openapi.GenericOpenAPIError).Model().(models.ProblemDetails).Cause
		rlogger.Trace(types.ModuleAmfSbi, rlogger.ERROR, nil,
			"failed to post SdmSubscription, the http status code(%d), return problem is(%+v)",
			res.StatusCode,
			problemDetails)
		//udm_message.SendHttpResponseMessage(httpChannel, nil, res.StatusCode, problemDetails)
		return fmt.Errorf("failed to post SdmSubscription, return (%+v)")
	}
	rlogger.Trace(types.ModuleAmfSbi, rlogger.DEBUG, nil, "Sdm Subscription Success")

	return nil
}

func (p *SbiAmf) HandleGetAMDataRequest(scid uint32, msg *sbicmn.SbiGetAmDataMsg) error {
	rlogger.FuncEntry(types.ModuleAmfSbi, nil)

	//invoke the api here
	resp, res, err := p.nudmSdmApiClt.AccessAndMobilitySubscriptionDataRetrievalApi.GetAmData(nil, msg.Supi, nil)
	rlogger.Trace(types.ModuleAmfSbi, rlogger.INFO, nil, "get udm sbi resp(%v,%v,%v)", resp, res, err)
	if err != nil {
		var problemDetails models.ProblemDetails
		problemDetails.Cause = err.(openapi.GenericOpenAPIError).Model().(models.ProblemDetails).Cause
		rlogger.Trace(types.ModuleAmfSbi, rlogger.ERROR, nil,
			"failed to GetAmData, the http status code(%d), return problem is(%+v)",
			res.StatusCode,
			problemDetails)
		//udm_message.SendHttpResponseMessage(httpChannel, nil, res.StatusCode, problemDetails)
		// send back the
		respMsg := &sbicmn.SbiMessage{MsgType: sbicmn.NudmFailMsg, ScInstId: scid}
		respData := &sbicmn.SbiHandleFailMsg{}
		respData.Supi = msg.Supi
		respData.Cause = fmt.Sprintf("code(%v),event(%s)",res.StatusCode,"get auth data")
		respMsg.MsgData = respData
		p.SendMsg2SC(scid, respMsg)
		return fmt.Errorf("failed to GetAmData, return (%+v)")
	}

	// send back the
	respMsg := &sbicmn.SbiMessage{MsgType: sbicmn.GetAmDataMsgResponse, ScInstId: scid}
	respData := &sbicmn.SbiGetAmDataMsg{}
	respData.Supi = msg.Supi
	respData.Data = &resp
	respMsg.MsgData = respData
	p.SendMsg2SC(scid, respMsg)

	return nil
}

func (p *SbiAmf) HandleGetSmfSelDataRequest(scid uint32, msg *sbicmn.SbiGetSmfSelDataMsg) error {
	rlogger.FuncEntry(types.ModuleAmfSbi, nil)

	//invoke the api here
	resp, res, err := p.nudmSdmApiClt.SMFSelectionSubscriptionDataRetrievalApi.GetSmfSelectData(nil, msg.Supi, nil)
	rlogger.Trace(types.ModuleAmfSbi, rlogger.INFO, nil, "get udm sbi resp(%v,%v,%v)", resp, res, err)
	if err != nil {
		var problemDetails models.ProblemDetails
		problemDetails.Cause = err.(openapi.GenericOpenAPIError).Model().(models.ProblemDetails).Cause
		rlogger.Trace(types.ModuleAmfSbi, rlogger.ERROR, nil,
			"failed to GetSmfSelData, the http status code(%d), return problem is(%+v)",
			res.StatusCode,
			problemDetails)
		//udm_message.SendHttpResponseMessage(httpChannel, nil, res.StatusCode, problemDetails)
		return fmt.Errorf("failed to GetAmData, return (%+v)")
	}

	// send back the
	respMsg := &sbicmn.SbiMessage{MsgType: sbicmn.GetSmfSelDataMsgResponse, ScInstId: scid}
	respData := &sbicmn.SbiGetSmfSelDataMsg{}
	respData.Supi = msg.Supi
	respData.Data = &resp
	respMsg.MsgData = respData
	p.SendMsg2SC(scid, respMsg)

	return nil
}

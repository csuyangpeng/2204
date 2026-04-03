package sbiamf

import (
	"fmt"
	"lite5gc/amf/sc/mobmgnt/mmutils"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/sbicmn"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types/configure"
	"lite5gc/openapi"
	"lite5gc/openapi/models"
)

func (p *SbiAmf) HandleGetAuthDataRequest(scid uint32, msg *sbicmn.SbiGetAuthDataMsg) error {
	rlogger.FuncEntry(types.ModuleAmfSbi, nil)
	//invoke the api here
	var authenticationInfoRequest models.AuthenticationInfoRequest
	snName := mmutils.GenerateSnName(configure.AmfConf.PlmnList.List[0])
	authenticationInfoRequest.ServingNetworkName = snName
	authenticationInfoRequest.AusfInstanceId = "1"
	rlogger.Trace(types.ModuleAmfSbi, rlogger.INFO, nil, "msg.Supi(%s)", msg.Supi)

	resp, res, err := p.nudmAuthApiClt.GenerateAuthDataApi.GenerateAuthData(nil, msg.Supi, authenticationInfoRequest)
	rlogger.Trace(types.ModuleAmfSbi, rlogger.INFO, nil, "get udm sbi resp(%v,%v,%v)", resp, res, err)
	if err != nil {
		var problemDetails models.ProblemDetails
		problemDetails.Cause = err.(openapi.GenericOpenAPIError).Model().(models.ProblemDetails).Cause
		rlogger.Trace(types.ModuleAmfSbi, rlogger.ERROR, nil,
			"failed to GetAuthData, the http status code(%d), return problem is(%+v)",
			res.StatusCode,
			problemDetails)
		//udm_message.SendHttpResponseMessage(httpChannel, nil, res.StatusCode, problemDetails)
		// send back the
		respMsg := &sbicmn.SbiMessage{MsgType: sbicmn.NudmFailMsg, ScInstId: scid}
		respData := &sbicmn.SbiHandleFailMsg{}
		respData.Supi = msg.Supi
		respData.Cause = fmt.Sprintf("code(%v),event(%s)", res.StatusCode, "get auth data")
		respMsg.MsgData = respData
		p.SendMsg2SC(scid, respMsg)
		return fmt.Errorf("failed to GetAuthData, return (%+v)", err)
	}

	// send back the
	respMsg := &sbicmn.SbiMessage{MsgType: sbicmn.GetAuthDataMsgResponse, ScInstId: scid}
	respData := &sbicmn.SbiGetAuthDataMsg{}
	respData.Supi = msg.Supi
	respData.Data = &resp
	respMsg.MsgData = respData
	p.SendMsg2SC(scid, respMsg)

	return nil
}

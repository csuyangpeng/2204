package sbismf

import (
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/sbicmn"
	"lite5gc/cmn/types"
	"fmt"
)

func (p *SbiSmf) handleGetSMDataRequest(scid uint32, msg *sbicmn.SbiGetSmDataMsg) error {
	rlogger.FuncEntry(types.ModuleSmfSbi, nil)
	if msg != nil {
		//invoke the api here
		resp, res, err := p.nudmSdmApiClt.SessionManagementSubscriptionDataRetrievalApi.GetSmData(nil, msg.Supi, nil)
		rlogger.Trace(types.ModuleSmfSbi, rlogger.ERROR, nil, "resp:(%v)(%v)(%v) ", resp, res, err)
		if err != nil {
			//var problemDetails models.ProblemDetails
			//problemDetails.Cause = err.(nudm_sdm.GenericOpenAPIError).Model().(models.ProblemDetails).Cause
			//rlogger.Trace(types.ModuleSmfSbi, rlogger.ERROR, nil,
			//	"failed to GetSmData, the http status code(%d), return problem is(%+v)",
			//	res.StatusCode,
			//	problemDetails)
			//udm_message.SendHttpResponseMessage(httpChannel, nil, res.StatusCode, problemDetails)
			return fmt.Errorf("failed to GetSmData, return (%+v)", err)
		}

		// send back the
		respMsg := &sbicmn.SbiMessage{MsgType: sbicmn.GetSmDataMsgResponse, ScInstId: scid}
		respData := &sbicmn.SbiGetSmDataMsg{}
		respData.Supi = msg.Supi
		respData.Psi = msg.Psi
		respData.Data = &resp
		respMsg.MsgData = respData
		SendScMsg(respMsg)
	} else {
		rlogger.Trace(types.ModuleSmfSbi, rlogger.ERROR, nil, "input para is nil")
		return types.ErrInputParaNil
	}
	return nil
}

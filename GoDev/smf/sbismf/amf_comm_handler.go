package sbismf

import (
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/sbicmn"
	"lite5gc/cmn/types"
	"fmt"
)

func (p *SbiSmf) handlePostN1N2MsgTransferRequest(scid uint32, msg *sbicmn.SbiPostN1N2MsgTransferMsg) error {
	rlogger.FuncEntry(types.ModuleSmfSbi, nil)
	if msg != nil {
		//invoke the api here
		resp, res, err := p.namfCommApiClt.N1N2MessageCollectionDocumentApi.N1N2MessageTransfer(nil, msg.Supi, *msg.ReqData)
		rlogger.Trace(types.ModuleSmfSbi, rlogger.INFO, nil, "resp: ", resp, res, err)
		if err != nil {
			//var problemDetails models.ProblemDetails
			//problemDetails.Cause = err.(namf_comm.GenericOpenAPIError).Model().(models.ProblemDetails).Cause
			//rlogger.Trace(types.ModuleSmfSbi, rlogger.ERROR, nil,
			//	"failed to N1N2MessageTransfer , the http status code(%d), return problem is(%+v)",
			//	res.StatusCode,
			//	problemDetails)
			//udm_message.SendHttpResponseMessage(httpChannel, nil, res.StatusCode, problemDetails)
			return fmt.Errorf("failed to N1N2MessageTransfer, return (%+v)", err)
		}

		// send back the
		respMsg := &sbicmn.SbiMessage{MsgType: sbicmn.N1N2MessageTransferResp, ScInstId: scid}
		respData := &sbicmn.SbiPostN1N2MsgTransferMsg{}
		respData.Supi = msg.Supi
		respData.Psi = msg.Psi
		respData.RespData = &resp
		respMsg.MsgData = respData
		SendScMsg(respMsg)
	} else {
		rlogger.Trace(types.ModuleSmfSbi, rlogger.ERROR, nil, "input msg is nil")
	}

	return nil
}

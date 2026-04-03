package sbiamf

import (
	"fmt"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/sbicmn"
	"lite5gc/cmn/types"
	"lite5gc/openapi"
	"lite5gc/openapi/models"
)

func (p *SbiAmf) HandlePostAmfRegistration(msg *sbicmn.SbiPostAmf3gppAccessRegistration) error {
	rlogger.FuncEntry(types.ModuleAmfSbi, nil)
	rlogger.Trace(types.ModuleAmfSbi, rlogger.INFO, nil, "msg.Supi(%s), *msg.Data(%v) ", msg.Supi, *msg.Data)
	//invoke the api here
	resp, res, err := p.nudmUeCtxtApiClt.AMFRegistrationFor3GPPAccessApi.Registration(nil, msg.Supi, *msg.Data)
	rlogger.Trace(types.ModuleAmfSbi, rlogger.INFO, nil, "get udm sbi resp(%v,%v,%v)", resp, res, err)
	if err != nil {
		var problemDetails models.ProblemDetails
		problemDetails.Cause = err.(openapi.GenericOpenAPIError).Model().(models.ProblemDetails).Cause
		rlogger.Trace(types.ModuleAmfSbi, rlogger.ERROR, nil,
			"failed to post AMFRegistrationFor3GPPAccess, the http status code(%d), return problem is(%+v)",
			res.StatusCode,
			problemDetails)
		//udm_message.SendHttpResponseMessage(httpChannel, nil, res.StatusCode, problemDetails)
		return fmt.Errorf("failed to post AMFRegistrationFor3GPPAccess")
	}
	rlogger.Trace(types.ModuleAmfSbi, rlogger.DEBUG, nil, "AMF Registration For 3GPPAccess Success")

	return nil
}

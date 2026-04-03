package procedure

import (
	"context"
	"lite5gc/amf/context/gctxt"
	"lite5gc/amf/context/prcdctxt"
	"lite5gc/amf/sc/naslayer/mmsender"
	"lite5gc/cmn/nas/nasie"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types/statetypes"
)

func HandleIdentityReqMsg(ctxt context.Context) {
	rlogger.FuncEntry(types.ModuleAmfMM, ctxt)

	//get ue context with amf ngap id
	ueContext, ok := ctxt.Value(types.UeContextCK).(*gctxt.UeContext)
	if !ok {
		rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ctxt, "type assertion failed for ue context")
		return
	}

	//set the sc instance id in amf id
	err := gctxt.AddIndexUeContext(gctxt.AmfUeNgApId(ueContext.GetAmfUeNgapId()), ueContext)
	if err != nil {
		rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueContext.GetImsiPtr(),
			"failed to add amf ue ngap id for ue context table")
		return
	}

	err = mmsender.SendIdentityRequest(ctxt, ueContext, nasie.Suci)
	if err != nil {
		rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueContext.GetImsiPtr(),
			"failed to send identity response msg")
		return
	}

	// 5gc\amf\sc\naslayer\ingress\handleIdentityResponse.go line126
	ueRegistPrcdCtxt, ok := ueContext.GetProcCtxt().(*prcdctxt.RegistrationPrcdCtxt)
	if !ok {
		rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueContext.GetImsiPtr(),
			"failed to get RegistrationPrcdCtxt, and :", ueContext.GetProcCtxt())
		return
	} else {
		err = ueRegistPrcdCtxt.SetNextState(statetype.StateRegisterWfIdentityResp)
		if err != nil {
			rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueContext.GetImsiPtr(),
				"failed to set state")
			return
		}
	}
}

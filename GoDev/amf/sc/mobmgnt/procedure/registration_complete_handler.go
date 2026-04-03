package procedure

import (
	"context"
	"lite5gc/amf/context/gctxt"
	"lite5gc/amf/context/prcdctxt"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types3gpp"
)

// HandleRegistrationCmpMsg process the registration complete message
func HandleRegistrationCmpMsg(ctxt context.Context) {

	rlogger.FuncEntry(types.ModuleAmfMM, ctxt)

	//get the ueCtxt
	ueCtxt, ok := ctxt.Value(types.UeContextCK).(*gctxt.UeContext)
	if !ok {
		rlogger.Trace(types.ModuleAmfMM, rlogger.DEBUG, ctxt, "no ue context.")
		return
	}

	//get the procedure context
	prcdCtxt, ok := ueCtxt.GetProcCtxt().(*prcdctxt.RegistrationPrcdCtxt)
	if !ok {
		rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"failed to get RegistrationPrcdCtxt, and :", ueCtxt.GetProcCtxt())
		return
	}

	// print the status changes
	RMState := ueCtxt.GetRmState()
	PrcdState := prcdCtxt.GetCurrentState()
	rlogger.Trace(types.ModuleAmfMM, rlogger.INFO, ueCtxt.GetImsiPtr(),
		"RM State is (%s), Procedure State is (%s).", RMState, PrcdState)

	//cancel t3550
	err := gctxt.CancelPrcdTimer(ueCtxt, ctxt, gctxt.T3550)
	if err != nil {
		rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"failed to cancel t3550")
		return
	}
	rlogger.Trace(types.ModuleAmfMM, rlogger.DEBUG, ctxt,
		"cancel t3550 timer for register accept msg")

	RegisterPrcdFinished(ctxt)

	// 若initial ue message的FOR(follow on pending)为0，注册完成后需要进行一轮an release流程
	if prcdCtxt.ForPending == false {
		rlogger.Trace(types.ModuleAmfMM, rlogger.DEBUG, ueCtxt.GetImsiPtr(),
			"for_pending is false. trigger the an_release procedure")
		err := StartAnRelPrcd(ctxt, ueCtxt, types3gpp.Nas_normal_release)
		if err != nil {
			rlogger.Trace(types.ModuleAmfMM, rlogger.DEBUG, ueCtxt.GetImsiPtr(),
				"fail to send ue context release command msg, err:", err)
			return
		}
	} else {
		rlogger.Trace(types.ModuleAmfMM, rlogger.DEBUG, ueCtxt.GetImsiPtr(),
			"for_pending is true. registration is success")
	}
	return
}

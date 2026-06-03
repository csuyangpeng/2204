package procedure

import (
	"lite5gc/amf/context/gctxt"
	"lite5gc/amf/context/prcdctxt"
	"lite5gc/amf/sc/naslayer/mmsender"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"

	"context"
	"lite5gc/cmn/types/statetypes"
	"lite5gc/cmn/types3gpp"
)

func HandleUeContextRelRequestMsg(ctxt context.Context) {
	rlogger.FuncEntry(types.ModuleAmfMM,ctxt)

	ueCtxt, ok := ctxt.Value(types.UeContextCK).(*gctxt.UeContext)
	if !ok {
		rlogger.Trace(types.ModuleAmfMM, rlogger.DEBUG, ctxt, "no ue context.")
		return
	}

	//get the prcdCtxt from ueCtxt
	prcdCtxt, ok := ueCtxt.GetProcCtxt().(*prcdctxt.AnReleasePrcdCtxt)
	if !ok {
		rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"failed to get AnReleasePrcdCtxt, and :", ueCtxt.GetProcCtxt())
		return
	}

	relCause := types3gpp.Nas_normal_release
	//send ue context release command msg to ran
	err := mmsender.SendUeContextReleaseCmd(ctxt, relCause)
	if err != nil {
		rlogger.Trace(types.ModuleAmfMM, rlogger.DEBUG, ueCtxt.GetImsiPtr(),
			"fail to send ue context release command msg to RAN,err:", err)
		return
	}

	// change the state in FSM
	err = prcdCtxt.SetNextState(statetype.StateAnRelWfRelCmp)
	if err != nil {
		rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"failed to set state")
		return
	}
	return

}

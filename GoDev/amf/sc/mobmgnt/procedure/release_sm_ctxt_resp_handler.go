package procedure

import (
	"lite5gc/amf/context/gctxt"
	"lite5gc/amf/context/prcdctxt"
	"lite5gc/amf/sc/naslayer/mmsender"
	"lite5gc/cmn/rlogger"

	"context"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types/statetypes"
	"lite5gc/cmn/types3gpp"
)

func HandleRelSmCtxtResp4Deregist(ctxt context.Context) {
	rlogger.FuncEntry(types.ModuleAmfMM, ctxt)

	//get the ueCtxt from ctxt
	ueCtxt, ok := ctxt.Value(types.UeContextCK).(*gctxt.UeContext)
	if !ok {
		rlogger.Trace(types.ModuleAmfMM, rlogger.DEBUG, ctxt, "no ue context.")
		return
	}

	// get procedure context
	procCtxt, ok := ueCtxt.GetProcCtxt().(*prcdctxt.DeRegistrationPrcdCtxt)
	if !ok {
		rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"failed to get DeRegistrationPrcdCtxt, and :", ueCtxt.GetProcCtxt())
		handleCause4DeregistReq(ueCtxt, nas.SystemFailure)
		return
	}

	// number of release sm context msg that amf send to smf before
	// for multiple messages, only if amf received all msg , amf can trigger the fsm
	counter := procCtxt.GetCounter()

	if counter-1 > 0 {
		counter-- // continue to wait for the rest of the message

		// reset the state
		err := procCtxt.SetNextState(statetype.StateDeRegisterWfRelSmCtxtResp)
		if err != nil {
			rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
				"failed to set state")
			return
		}

		procCtxt.SetCounter(counter)

		return
	} else {
		if procCtxt.DeRegistrationType.SwithOff == true {
			// send an release command to gnb
			relCause := types3gpp.Nas_deregister
			//send ue context command to ue
			err := mmsender.SendUeContextReleaseCmd(ctxt, relCause)
			if err != nil {
				rlogger.Trace(types.ModuleAmfMM, rlogger.DEBUG, ueCtxt.GetImsiPtr(),
					"failed to send ue context command msg")
			}
		} else {
			err := HandleTriggerSendDeregiserAccetpAndUeRelCmdMsg(ctxt)
			if err != nil {
				rlogger.Trace(types.ModuleAmfMM, rlogger.INFO, ueCtxt.GetImsiPtr(),
					"failed to handle trigger send dereigster accept message.")
			}
		}
		//set status for FSM
		err := procCtxt.SetNextState(statetype.StateDeRegisterWfUeCtxtRelCmp)
		if err != nil {
			rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
				"failed to set state")
			return
		}
	}
}

package procedure

import (
	"context"
	"lite5gc/amf/context/gctxt"
	"lite5gc/amf/context/prcdctxt"
	"lite5gc/amf/sc/mobmgnt/mmutils"
	"lite5gc/amf/sc/naslayer/mmsender"
	"lite5gc/cmn/message/n11msg"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types/statetypes"
	"lite5gc/cmn/types3gpp"
	"time"
)

func HandleDeRegistrationRequest(ctxt context.Context) {
	rlogger.FuncEntry(types.ModuleAmfMM, nil)

	//get the ueCtxt from ctxt
	ueCtxt, ok := ctxt.Value(types.UeContextCK).(*gctxt.UeContext)
	if !ok {
		rlogger.Trace(types.ModuleAmfMM, rlogger.WARN, nil, "no ue context")
		handleCause4DeregistReq(ueCtxt, nas.SystemFailure)
		return
	}

	//must registered before de-registered
	if ueCtxt.GetRmState() != types.StateRmRegistered {
		rlogger.Trace(types.ModuleAmfMM, rlogger.WARN, ueCtxt.GetImsiPtr(),
			"rm status expect (Registered), but is (%s)", ueCtxt.GetRmState())
		handleCause4DeregistReq(ueCtxt, nas.SystemFailure)
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

	//Determine if GUTI is assigned by this AMF
	isAssigned, err := mmutils.Check5GGuti(procCtxt.Guti5g)
	if !isAssigned || err != nil {
		rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"guti not assigned by AMF, err(%s)", err)
		handleCause4DeregistReq(ueCtxt, nas.IllegalUE)
		//TODO start identity request procedure
		return
	}

	//check Pdu Session in Ue Ctxt
	if len(ueCtxt.GetPDUSessionCtxts()) != 0 {
		rlogger.Trace(types.ModuleAmfMM, rlogger.INFO, ueCtxt.GetImsiPtr(),
			"there are (%d) session resource need to release", len(ueCtxt.GetPDUSessionCtxts()))
		procCtxt.SetCounter(len(ueCtxt.GetPDUSessionCtxts()))
		// has session established , need to send msg to SMF to release related resource
		for _, v := range ueCtxt.GetPDUSessionCtxts() {
			releaseSmCtxtReq := n11msg.ReleaseSMContextRequestData{}
			releaseSmCtxtReq.NgApCause.Type = types3gpp.CT_Nas
			releaseSmCtxtReq.NgApCause.Value = types3gpp.Nas_deregister
			mmsender.SendReleaseSMCtxtSBIMsg(ctxt, v.Psi, ueCtxt.GetImsi(), releaseSmCtxtReq)
		}

		//set status of FSM
		err = procCtxt.SetNextState(statetype.StateDeRegisterWfRelSmCtxtResp)
		if err != nil {
			rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
				"failed to set state")
			return
		}
	} else {
		err := HandleTriggerSendDeregiserAccetpAndUeRelCmdMsg(ctxt)
		if err != nil {
			rlogger.Trace(types.ModuleAmfMM, rlogger.INFO, ueCtxt.GetImsiPtr(),
				"failed to handle trigger send dereigster accept message.")
		}
		//set status for FSM
		err = procCtxt.SetNextState(statetype.StateDeRegisterWfUeCtxtRelCmp)
		if err != nil {
			rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
				"failed to set state")
			return
		}
	}
	return
}

func handleCause4DeregistReq(uectxt *gctxt.UeContext, cause nas.Mm5gCause) {
	rlogger.FuncEntry(types.ModuleAmfMM, uectxt)
}

func HandleTriggerSendDeregiserAccetpAndUeRelCmdMsg(ctxt context.Context) error {
	rlogger.FuncEntry(types.ModuleAmfMM, ctxt)

	//send deregister accept
	err := mmsender.SendDeregisterAcceptMsg(ctxt)
	if err != nil {
		rlogger.Trace(types.ModuleAmfMM, rlogger.DEBUG, nil,
			"failed to send deregister accept msg ")
		//TODO fix error scenario
		return err
	}

	//TODO fix the sequence of these two messages
	time.Sleep(time.Millisecond * 20)

	relCause := types3gpp.Nas_deregister

	//send ue context command to ue
	err = mmsender.SendUeContextReleaseCmd(ctxt, relCause)
	if err != nil {
		rlogger.Trace(types.ModuleAmfMM, rlogger.DEBUG, nil,
			"failed to send ue context command msg")
		return err
	}
	return nil
}

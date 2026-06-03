package procedure

import (
	"lite5gc/amf/context/gctxt"
	"lite5gc/amf/context/prcdctxt"
	"lite5gc/amf/sc/naslayer/mmsender"
	"lite5gc/cmn/rlogger"

	"context"
	"lite5gc/cmn/message/n11msg"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types/statetypes"
	"lite5gc/cmn/types3gpp"
)

func HandleUeContextRelCompleteMsg(ctxt context.Context) {
	rlogger.FuncEntry(types.ModuleAmfMM, ctxt)

	//get the imsi
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

	var psiList []uint8
	if len(prcdCtxt.ReleasePsiList) > 0 {
		psiList = append(psiList, prcdCtxt.ReleasePsiList...)
	} else {
		psiList = append(psiList, ueCtxt.GetPsiList(types3gpp.SessActived)...)
	}
	rlogger.Trace(types.ModuleAmfMM, rlogger.DEBUG, ueCtxt.GetImsiPtr(), "PSI List %v", psiList)

	if len(psiList) > 0 {
		//send update sm context request to SMF
		for _, psi := range psiList {
			smCtxtUpdateData := n11msg.UpdateSMContextRequestData{}
			smCtxtUpdateData.UpCnxState = n11msg.DEACTIVATED
			smCtxtUpdateData.IeFlags.Set(n11msg.Ieid_upCnxState)
			mmsender.SendUpdateSMCtxtSBIMsg(ctxt, nas.PduSessID(psi), ueCtxt.GetImsi(), smCtxtUpdateData)
		}
		//set counter
		prcdCtxt.SetCounter(len(psiList))
		//change the state in FSM
		err := prcdCtxt.SetNextState(statetype.StateAnRelWfUpSmCtxtAck)
		if err != nil {
			rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
				"failed to set state")
			return
		}
		return
	} else {
		rlogger.Trace(types.ModuleAmfMM, rlogger.DEBUG, ueCtxt.GetImsiPtr(),
			"no session to release, an release is done")
	}
}

func HandleUeCtxtRelCmp4Deregist(ctxt context.Context) {
	//get the imsi
	ueCtxt, ok := ctxt.Value(types.UeContextCK).(*gctxt.UeContext)
	if !ok {
		rlogger.Trace(types.ModuleAmfMM, rlogger.DEBUG, ctxt, "no ue context.")
		return
	}
	ReleaseN2ConnId(ueCtxt)
	DeregisterPrcdFinished(ctxt)
}

func ReleaseN2ConnId(ueCtxt *gctxt.UeContext) {
	rlogger.FuncEntry(types.ModuleAmfMM, ueCtxt)

	rlogger.Trace(types.ModuleAmfMM, rlogger.DEBUG, nil,
		"delete key(%d) AmfUeNgApId-N2ConnCtxt MapTable.", ueCtxt.GetAmfUeNgapId())

	//release N2 connection
	err := gctxt.DeleteN2ConnContext(gctxt.AmfUeNgApId(ueCtxt.GetAmfUeNgapId()))
	if err != nil {
		rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(), "failed to delete UeContext, error(%s)", err)
	}
}

func HandleUeCtxtReqBySwitchOff(ctxt context.Context) {
	rlogger.FuncEntry(types.ModuleAmfMM, ctxt)

	// release pdu session on smf
	ueCtxt, ok := ctxt.Value(types.UeContextCK).(*gctxt.UeContext)
	if !ok {
		rlogger.Trace(types.ModuleAmfMM, rlogger.DEBUG, nil, "no ue context.")
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
		err := procCtxt.SetNextState(statetype.StateDeRegisterWfRelSmCtxtResp)
		if err != nil {
			rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
				"failed to set state")
			return
		}
	} else {
		rlogger.Trace(types.ModuleAmfMM, rlogger.INFO, ueCtxt.GetImsiPtr(),
			"there is no session resource need to release,send ue context cmd directly")
		// send an release command to gnb
		relCause := types3gpp.Nas_deregister
		//send ue context command to ue
		err := mmsender.SendUeContextReleaseCmd(ctxt, relCause)
		if err != nil {
			rlogger.Trace(types.ModuleAmfMM, rlogger.DEBUG, ueCtxt.GetImsiPtr(),
				"failed to send ue context command msg")
		}
		//set status for FSM
		err = procCtxt.SetNextState(statetype.StateDeRegisterWfUeCtxtRelCmp)
		if err != nil {
			rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
				"failed to set state")
			return
		}
	}
}

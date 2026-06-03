package procedure

import (
	"lite5gc/amf/context/gctxt"
	"lite5gc/amf/context/prcdctxt"
	"lite5gc/amf/sc/naslayer/mmsender"
	"lite5gc/cmn/message/nasmsg"
	"lite5gc/cmn/rlogger"

	"context"
	"lite5gc/cmn/message/n11msg"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types/statetypes"
	"lite5gc/cmn/types3gpp"
)

func HandleSessRelReqMsg(ctxt context.Context) {

	rlogger.FuncEntry(types.ModuleAmfMM, ctxt)

	//get the ueCtxt
	ueCtxt, ok := ctxt.Value(types.UeContextCK).(*gctxt.UeContext)
	if !ok {
		rlogger.Trace(types.ModuleAmfMM, rlogger.DEBUG, ctxt, "no ue context.")
		return
	}
	// get the UplinkNasTransportMsg
	ulNasMsg, ok := ctxt.Value(types.UplinkMsgCK).(*nasmsg.UplinkNasTransportMsg)
	if !ok {
		rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"failed to get uplink msg for session release request")
		return
	}

	//get the prcdCtxt from ueCtxt
	prcdCtxt, ok := ueCtxt.GetProcCtxt().(*prcdctxt.SessionReleasePrcdCtxt)
	if !ok {
		rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"failed to get SessionReleasePrcdCtxt, and :", ueCtxt.GetProcCtxt())
		return
	}

	//create create session context and send to smf with N11 if
	updateSmCtxtReq := n11msg.UpdateSMContextRequestData{}

	updateSmCtxtReq.N1SmMsg = ulNasMsg.PayloadContainer.PayloadContainerEntry[0].ContainerContents
	updateSmCtxtReq.IeFlags.Set(n11msg.Ieid_n1SmMsg)

	updateSmCtxtReq.Release = true
	updateSmCtxtReq.IeFlags.Set(n11msg.Ieid_release)

	updateSmCtxtReq.UpCnxState = n11msg.DEACTIVATED
	updateSmCtxtReq.IeFlags.Set(n11msg.Ieid_upCnxState)

	rlogger.Trace(types.ModuleAmfMM, rlogger.DEBUG, ueCtxt.GetImsiPtr(), "prcdCtxt.Psi ", prcdCtxt.Psi)

	mmsender.SendUpdateSMCtxtSBIMsg(ctxt, prcdCtxt.Psi, ueCtxt.GetImsi(), updateSmCtxtReq)

	//set the next state
	err := prcdCtxt.SetNextState(statetype.StateSessRelWfUpSmCtxtResp)
	if err != nil {
		rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"failed to set state")
		return
	}
}

func HandleSessUpdateSmCtxtResp(ctxt context.Context) {

	rlogger.FuncEntry(types.ModuleAmfMM, ctxt)

	//get the ueCtxt
	ueCtxt, ok := ctxt.Value(types.UeContextCK).(*gctxt.UeContext)
	if !ok {
		rlogger.Trace(types.ModuleAmfMM, rlogger.DEBUG, ctxt, "no ue context.")
		return
	}

	msgData, ok := ctxt.Value(types.SmfN11MsgDataCK).(n11msg.UpdateSMContextResponseData)
	if !ok {
		rlogger.Trace(types.ModuleAmfMM, rlogger.DEBUG, ueCtxt.GetImsiPtr(), "no N11MsgData.")
		return
	}

	prcdCtxt, ok := ueCtxt.GetProcCtxt().(*prcdctxt.SessionReleasePrcdCtxt)
	if !ok {
		rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"failed to get SessionReleasePrcdCtxt, and :", ueCtxt.GetProcCtxt())
		return
	}

	pduSessCtxt := ueCtxt.GetPduSessCtxt(prcdCtxt.Psi)
	pduSessCtxt.Status = types3gpp.SessDeactive

	//send session release accetp & n2 resource release request to ue
	err := mmsender.SendN2ResRelReq(ctxt, msgData)
	if err != nil {
		rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(), "failed to send service accept msg")
	}

	//clear n2 sm info in procedure context
	prcdCtxt.N2SmInfo = make(map[uint32][]byte)

	//set next status
	err = prcdCtxt.SetNextState(statetype.StateSessRelWfUpSmCtxtReqSec)
	if err != nil {
		rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"failed to set state")
		return
	}

}

func UpdateSmCtxt2ndRespMsg(ctxt context.Context) {

	rlogger.FuncEntry(types.ModuleAmfMM, ctxt)

	//get the ueCtxt
	ueCtxt, ok := ctxt.Value(types.UeContextCK).(*gctxt.UeContext)
	if !ok {
		rlogger.Trace(types.ModuleAmfMM, rlogger.DEBUG, ctxt, "no ue context.")
		return
	}
	// get the UplinkNasTransportMsg
	ulNasMsg, ok := ctxt.Value(types.UplinkMsgCK).(*nasmsg.UplinkNasTransportMsg)
	if !ok {
		rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"failed to get uplink msg for session release request")
	}

	//get the prcdCtxt from ueCtxt
	prcdCtxt, ok := ueCtxt.GetProcCtxt().(*prcdctxt.SessionReleasePrcdCtxt)
	if !ok {
		rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"failed to get SessionReleasePrcdCtxt, and :", ueCtxt.GetProcCtxt())
		return
	}

	//create create session context and send to smf with N11 if
	updateSmCtxtReq := n11msg.UpdateSMContextRequestData{}

	//set the n1 sm message
	updateSmCtxtReq.N1SmMsg = ulNasMsg.PayloadContainer.PayloadContainerEntry[0].ContainerContents
	updateSmCtxtReq.IeFlags.Set(n11msg.Ieid_n1SmMsg)

	updateSmCtxtReq.Release = true
	updateSmCtxtReq.IeFlags.Set(n11msg.Ieid_release)

	updateSmCtxtReq.UpCnxState = n11msg.DEACTIVATED
	updateSmCtxtReq.IeFlags.Set(n11msg.Ieid_upCnxState)

	mmsender.SendUpdateSMCtxtSBIMsg(ctxt, prcdCtxt.Psi, ueCtxt.GetImsi(), updateSmCtxtReq)

	//set next status
	err := prcdCtxt.SetNextState(statetype.StateSessRelWfUpSmCtxtRespSec)
	if err != nil {
		rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"failed to set state")
		return
	}
}

func HandleSessRelEnd(ctxt context.Context) {
	rlogger.FuncEntry(types.ModuleAmfMM, ctxt)

	SessionReleasePrcdFinished(ctxt)
}

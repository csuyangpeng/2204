package procedure

import (
	"context"
	"lite5gc/amf/context/gctxt"
	"lite5gc/amf/context/prcdctxt"
	nasegress "lite5gc/amf/sc/naslayer/mmsender"
	"lite5gc/amf/sc/naslayer/nassecurity"
	"lite5gc/amf/sc/ngaplayer/ngapsender"
	"lite5gc/amf/sc/statistics"
	"lite5gc/amf/sc/utils"
	"lite5gc/cmn/message/nasmsg"
	"lite5gc/cmn/message/n11msg"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/nas/nasie"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types/statetypes"
	"lite5gc/oam/pm"
)

func HandleSessMod1stUpSmCtxtReqMsg(ctxt context.Context) {

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
	prcdCtxt, ok := ueCtxt.GetProcCtxt().(*prcdctxt.SessionModificationPrcdCtxt)
	if !ok {
		rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"failed to get SessionModificationPrcdCtxt, and :", ueCtxt.GetProcCtxt())
		return
	}

	//create create session context and send to smf with N11 if
	updateSmCtxtReq := n11msg.UpdateSMContextRequestData{}

	//set the n1 sm message
	updateSmCtxtReq.N1SmMsg = ulNasMsg.PayloadContainer.PayloadContainerEntry[0].ContainerContents
	updateSmCtxtReq.IeFlags.Set(n11msg.Ieid_n1SmMsg)

	nasegress.SendUpdateSMCtxtSBIMsg(ctxt, prcdCtxt.Psi, ueCtxt.GetImsi(), updateSmCtxtReq)

	//set the next state
	err := prcdCtxt.SetNextState(statetype.StateSessModWf1stUpSmCtxtResp)
	if err != nil {
		rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"failed to set state")
		return
	}
}

func HandleSessModN2SessReqMsg(ctxt context.Context) {

	rlogger.FuncEntry(types.ModuleAmfMM, ctxt)

	//get the ueCtxt
	ueCtxt, ok := ctxt.Value(types.UeContextCK).(*gctxt.UeContext)
	if !ok {
		rlogger.Trace(types.ModuleAmfMM, rlogger.DEBUG, ctxt, "no ue context.")
		return
	}
	msgData, ok := ctxt.Value(types.SmfN11MsgDataCK).(n11msg.UpdateSMContextResponseData)
	if !ok {
		rlogger.Trace(types.ModuleAmfMM, rlogger.DEBUG, ueCtxt.GetImsiPtr(), "no UpdateSMContextResponseData.")
		return
	}

	//get the prcdCtxt from ueCtxt
	prcdCtxt, ok := ueCtxt.GetProcCtxt().(*prcdctxt.SessionModificationPrcdCtxt)
	if !ok {
		rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"failed to get SessionModificationPrcdCtxt, and :", ueCtxt.GetProcCtxt())
		return
	}
	// get nas message
	var nasMsgData []byte
	if msgData.IeFlags.Test(n11msg.Ieid_n1SmMsg) {
		nasMsgData = append(nasMsgData, []byte(msgData.N1SmMsg)...)
	} else {
		rlogger.Trace(types.ModuleAmfMM, rlogger.DEBUG, ueCtxt.GetImsiPtr(), "no n1 message container")
	}

	// get sc nagp egress layer
	//  get the scnglayer
	sender, ok := ctxt.Value(types.NgapSenderCK).(*ngapsender.NgapSender)
	if !ok {
		rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"failed to get ngap layer .")
		return
	}

	dlNasTransMsgData := &nasmsg.DownLinkNasTransportMsg{}
	dlNasTransMsgData.PayloadType = nasie.N1SmInformation
	dlNasTransMsgData.PayloadContainer.PayloadContainerEntry = make([]nasie.PayloadContainerENTRY, 1)
	dlNasTransMsgData.PayloadContainer.PayloadContainerEntry[0].ContainerContents =
		append(dlNasTransMsgData.PayloadContainer.PayloadContainerEntry[0].ContainerContents, nasMsgData...)

	dlNasData, err := dlNasTransMsgData.Encode()
	if err != nil {
		rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"failed to encode downlink nas transport message.")
		return
	}

	//add sec nas header
	// set security protected header
	dlNasMsg, err := utils.EncodeSecPrctNasMsg(ueCtxt, nas.IntegrityPrtctCipher, dlNasData)
	if err != nil {
		rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(), "failed to add security header")
		return
	}

	//msg counter
	pm.PegCounter(statistics.DLNasTransportCounter)

	err = sender.SendPduSessResModRequest(ueCtxt, dlNasMsg, msgData.N2SmInfo, 5)
	if err != nil {
		rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"failed to send pdu session res modify request message")
		return
	}

	//increate nas counter
	nassecurity.UpdateDownlinkNasCounter(ueCtxt)
	//set the next state
	err = prcdCtxt.SetNextState(statetype.StateSessModWfN2SessResp)
	if err != nil {
		rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"failed to set state")
		return
	}
}

func HandleSessMod2ndUpSmCtxtReqMsg(ctxt context.Context) {

	rlogger.FuncEntry(types.ModuleAmfMM, ctxt)
}

func HandleSessModN2NASUplinkTransMsg(ctxt context.Context) {

	rlogger.FuncEntry(types.ModuleAmfMM, ctxt)
}

func HandleSessMod3rdUpSmCtxtReqMsg(ctxt context.Context) {

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
	prcdCtxt, ok := ueCtxt.GetProcCtxt().(*prcdctxt.SessionModificationPrcdCtxt)
	if !ok {
		rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"failed to get SessionModificationPrcdCtxt, and :", ueCtxt.GetProcCtxt())
		return
	}

	//create create session context and send to smf with N11 if
	updateSmCtxtReq := n11msg.UpdateSMContextRequestData{}

	//set the n1 sm message
	updateSmCtxtReq.N1SmMsg = ulNasMsg.PayloadContainer.PayloadContainerEntry[0].ContainerContents
	updateSmCtxtReq.IeFlags.Set(n11msg.Ieid_n1SmMsg)

	nasegress.SendUpdateSMCtxtSBIMsg(ctxt, prcdCtxt.Psi, ueCtxt.GetImsi(), updateSmCtxtReq)
	//set the next state
	err := prcdCtxt.SetNextState(statetype.StateSessModWf3rdUpSmCtxtResp)
	if err != nil {
		rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"failed to set state")
		return
	}
}

func HandleSessModEnd(ctxt context.Context) {
	rlogger.FuncEntry(types.ModuleAmfMM, ctxt)
	//get the ueCtxt
	ueCtxt, ok := ctxt.Value(types.UeContextCK).(*gctxt.UeContext)
	if !ok {
		rlogger.Trace(types.ModuleAmfMM, rlogger.DEBUG, ctxt, "no ue context.")
		return
	}
	ueCtxt.SetProcCtxt(nil)
	rlogger.Trace(types.ModuleAmfMM, rlogger.INFO, ueCtxt.GetImsiPtr(),
		"session modification procedure finished in AMF")
}

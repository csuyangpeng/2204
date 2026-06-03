package mmsender

import (
	"context"
	"lite5gc/amf/context/gctxt"
	"lite5gc/amf/sc/ngaplayer/ngapsender"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types3gpp"
)

func SendUeContextReleaseCmd(ctxt context.Context, relCause types3gpp.CauseValue) error {
	rlogger.FuncEntry(types.ModuleAmfNas, ctxt)

	//get the ueCtxt from ctxt
	ueCtxt, ok := ctxt.Value(types.UeContextCK).(*gctxt.UeContext)
	if !ok {
		rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, ctxt, "no ue context")
		return types.ErrFailFindUeCtxt
	}

	//get gnb instance id for sending ngap msg
	gnbInstId, err := gctxt.GetGnbInstIdByAmfUeNgapId(ueCtxt.GetAmfUeNgapId())
	if err != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, nil, "failed to get gnb instance id")
		return err
	}

	//  get the scnglayer
	sender, ok := ctxt.Value(types.NgapSenderCK).(*ngapsender.NgapSender)
	if !ok {
		rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"failed to get ngap layer .")
		return types.ErrFailFindNgapLayer
	}

	// send the message in scng layer
	err = sender.SendUEContextReleaseCommand(ueCtxt, gnbInstId, relCause)
	if err != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.WARN, ctxt, "failed to send UEContext Release Command message,err :", err)
		return err
	}

	return nil
}

func SendUeContextReleaseCmdSim(ctxt context.Context, relCause types3gpp.CauseValue) error {
	rlogger.FuncEntry(types.ModuleAmfNas, ctxt)

	//get the ueCtxt from ctxt
	ueCtxt, ok := ctxt.Value(types.UeContextCK).(*gctxt.UeContext)
	if !ok {
		rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, ctxt, "no ue context")
		return types.ErrFailFindUeCtxt
	}

	//get n2 connection context
	n2Conn, err := gctxt.GetN2ConnContext((gctxt.AmfUeNgApId)(ueCtxt.GetAmfUeNgapId()))
	if err != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, ctxt, "failed to get n2 connection data,err :", err)
		return err
	}

	//  get the scnglayer
	sender, ok := ctxt.Value(types.NgapSenderCK).(*ngapsender.NgapSender)
	if !ok {
		rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"failed to get ngap layer .")
		return types.ErrFailFindNgapLayer
	}

	// send the message in scng layer
	err = sender.SendUEContextReleaseCommandSim(n2Conn, relCause)
	if err != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.WARN, ctxt, "failed to send UEContext Release Command message,err :", err)
		return err
	}

	return nil
}

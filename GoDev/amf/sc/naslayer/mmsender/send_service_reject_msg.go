package mmsender

import (
	"context"
	"fmt"
	"lite5gc/amf/context/gctxt"
	"lite5gc/amf/sc/naslayer/nassecurity"
	"lite5gc/amf/sc/ngaplayer/ngapsender"
	"lite5gc/amf/sc/statistics"
	"lite5gc/cmn/message/nasmsg"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/oam/pm"
)

func SendServiceRejectMsg(ctxt context.Context, cause nas.Mm5gCause) error {
	rlogger.FuncEntry(types.ModuleAmfNas, ctxt)

	ueCtxt, ok := ctxt.Value(types.UeContextCK).(*gctxt.UeContext)
	if !ok {
		rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, ctxt,
			"no ue context")
		return types.ErrFailFindUeCtxt
	}

	rlogger.FuncEntry(types.ModuleAmfNas, ueCtxt)

	srmsg := nasmsg.ServiceRejectMsg{}
	srmsg.FiveGSRCause = cause
	// encode nas message
	bytes, err := srmsg.Encode()
	if err != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, ueCtxt, "failed to encode service reject message")
		return fmt.Errorf("failed to encode service reject message")
	}
	rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, ueCtxt, "encode service reject nas msg:", bytes)

	// add the security header
	procNasMsg := VerifyAndBuildSecProtectNasMsg(ueCtxt, bytes)

	//msg counter
	pm.PegCounter(statistics.ServiceRejectCounter)

	//  get the scnglayer
	sender, ok := ctxt.Value(types.NgapSenderCK).(*ngapsender.NgapSender)
	if !ok {
		rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"failed to get ngap layer .")
		return types.ErrFailFindNgapLayer
	}

	//get n2 connection context
	gnbInstId, err := gctxt.GetGnbInstIdByAmfUeNgapId(ueCtxt.GetAmfUeNgapId())
	if err != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, ueCtxt, "failed to get gnbInstId by AmfUeNgapId")
		return fmt.Errorf("failed to get gnbInstId by AmfUeNgapId")
	}

	// send the message in scng layer
	err = sender.SendInitialContextSetupRequest(ueCtxt, gnbInstId, procNasMsg, nil) //todo
	if err != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.WARN, ueCtxt, "failed to send service reject message")
		return fmt.Errorf("failed to send service reject message")
	}

	nassecurity.UpdateDownlinkNasCounter(ueCtxt)

	//set state for ue
	err = ueCtxt.SetCmState(types.CmIdle)
	if err != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.WARN, ueCtxt, "failed to set CM state")
		return fmt.Errorf("failed to set CM state")
	}
	//set the prcd ctxt to nil
	ueCtxt.SetProcCtxt(nil)

	return nil
}

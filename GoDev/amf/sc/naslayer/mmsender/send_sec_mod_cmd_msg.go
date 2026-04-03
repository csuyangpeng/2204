package mmsender

import (
	"context"
	"fmt"
	"lite5gc/amf/context/gctxt"
	"lite5gc/amf/sc/naslayer/nassecurity"
	"lite5gc/amf/sc/utils"
	"lite5gc/cmn/message/nasmsg"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types3gpp"
)

func SendSecurityModeCommandMsg(ctxt context.Context, selAlgo types3gpp.SelNasSecAlgo) error {
	rlogger.FuncEntry(types.ModuleAmfNas, nil)

	//get the ueCtxt from ctxt
	ueCtxt, ok := ctxt.Value(types.UeContextCK).(*gctxt.UeContext)
	if !ok {
		rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, ctxt,
			"failed to get ue context")
		return types.ErrFailFindUeCtxt
	}

	// derive the keys
	err := nassecurity.DeriveSecurityKeys(ueCtxt)
	if err != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"failed to derive security keys")
		return fmt.Errorf("failed to derive security keys")
	}

	//apply new security context here.
	ueCtxt.UpdateSecCtxt()

	//send authentication request message to UE
	msg := nasmsg.SecurityModeCommandMsg{}
	msg.SelectNasSecAlg = selAlgo
	msg.NgKSI = ueCtxt.NgKsi
	msg.UeSecCap = ueCtxt.SecurityCtxt.UeSecCapablity
	encBuf, err := msg.Encode()
	if err != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"failed to encode security mode command message")
		return err
	}

	// set security protected header
	encBuffer, err := utils.EncodeSecPrctNasMsg(ueCtxt, nas.IntegrityPrtctNewSecCtxt, encBuf)
	if err != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"failed to add security header")
		return err
	}

	err, encNgapMsg := SendDownLinkNasMsg(ctxt, ueCtxt, encBuffer)
	if err != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"fail to send downlink nas ngap msg")
		return types.ErrFailSendNgapMsg
	}

	// update downlink nas counter
	nassecurity.UpdateDownlinkNasCounter(ueCtxt)

	// start t3560 timer for authentication request message
	err = StartNasTimer(ctxt, ueCtxt, gctxt.T3560, encNgapMsg)
	if err != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"failed to start nas timer for timer(%s), gctxt.T3560, error(%s)", gctxt.T3560, err)
		return fmt.Errorf("failed to start nas timer for timer(%s), gctxt.T3560, error(%s)", gctxt.T3560, err)
	}

	return nil
}

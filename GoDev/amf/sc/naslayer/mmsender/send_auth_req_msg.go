package mmsender

import (
	"context"
	"fmt"
	"lite5gc/amf/context/gctxt"
	"lite5gc/amf/sc/naslayer/nassecurity"
	"lite5gc/amf/sc/statistics"
	"lite5gc/cmn/message/nasmsg"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/oam/pm"
)

func SendAuthenticationRequestMsg(ctxt context.Context) error {
	rlogger.FuncEntry(types.ModuleAmfNas, nil)

	// get ue context
	ueCtxt, ok := ctxt.Value(types.UeContextCK).(*gctxt.UeContext)
	if !ok {
		rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, nil,
			"fail to get ue context")
		return types.ErrFailFindUeCtxt
	}

	//send authentication request message
	msg := nasmsg.AuthenticationRequestMsg{}
	msg.NgKSI = ueCtxt.NgKsi
	msg.Abba = ueCtxt.TempSecCtxt.Abba[:]
	msg.Rand = ueCtxt.TempSecCtxt.AuthVector.Rand
	msg.IeFlags.Set(nasmsg.IeidAuthreqRand)
	msg.Autn = ueCtxt.TempSecCtxt.AuthVector.Autn
	msg.IeFlags.Set(nasmsg.IeidAuthreqAutn)
	rlogger.Trace(types.ModuleAmfNas, rlogger.INFO, ueCtxt.GetImsiPtr(),
		"AuthenticationRequestMsg (%x)", msg)

	encBuf, err := msg.Encode()
	if err != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"failed to encode authentication request message")
		return types.ErrFailEncodeNasMsg
	}

	procNasMsg := VerifyAndBuildSecProtectNasMsg(ueCtxt, encBuf)

	//msg counter
	pm.PegCounter(statistics.AuthenticationRequestCounter)

	err, encNgapMsg := SendDownLinkNasMsg(ctxt, ueCtxt, procNasMsg)
	if err != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"fail to send downlink nas ngap msg")
		return types.ErrFailSendNgapMsg
	}

	if ueCtxt.IsSecCtxtEstb && (ueCtxt.SecHdrTypeIncomingMsg != nas.PlainNasMsg) {
		nassecurity.UpdateDownlinkNasCounter(ueCtxt)
	}

	// start t3560 timer for authentication request message
	err = StartNasTimer(ctxt, ueCtxt, gctxt.T3560, encNgapMsg)
	if err != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"failed to start T3560 nas timer, error(%s)", err)
		return fmt.Errorf("failed to start T3560 nas timer, error(%s)", err)
	}

	return nil
}

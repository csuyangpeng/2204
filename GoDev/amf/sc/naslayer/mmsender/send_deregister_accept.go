package mmsender

import (
	"golang.org/x/net/context"
	"lite5gc/amf/context/gctxt"
	"lite5gc/amf/sc/naslayer/nassecurity"
	"lite5gc/amf/sc/statistics"
	"lite5gc/cmn/message/nasmsg"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/oam/pm"
)

func SendDeregisterAcceptMsg(ctxt context.Context) error {

	rlogger.FuncEntry(types.ModuleAmfNas, ctxt)

	// get ue context
	ueCtxt, ok := ctxt.Value(types.UeContextCK).(*gctxt.UeContext)
	if !ok {
		rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, nil,
			"fail to get ue context")
		return types.ErrFailFindUeCtxt
	}

	var deregisterAcceptMsg nasmsg.DeRegistrationAcceptMsg
	encBuf := deregisterAcceptMsg.Encode()

	// add the security header
	procNasMsg := VerifyAndBuildSecProtectNasMsg(ueCtxt, encBuf)

	err, _ := SendDownLinkNasMsg(ctxt, ueCtxt, procNasMsg)
	if err != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"fail to send downlink nas ngap msg")
		return types.ErrFailSendNgapMsg
	}

	nassecurity.UpdateDownlinkNasCounter(ueCtxt)

	//msg counter
	pm.PegCounter(statistics.DeregistrationAcceptUeCounter)

	return nil
}

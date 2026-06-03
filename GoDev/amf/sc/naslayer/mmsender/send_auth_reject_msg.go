package mmsender

import (
	"context"
	"lite5gc/amf/context/gctxt"
	"lite5gc/amf/sc/statistics"
	"lite5gc/cmn/message/nasmsg"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/oam/pm"
)

func SendAuthenticationRejectMsg(ctxt context.Context) error {
	rlogger.FuncEntry(types.ModuleAmfNas, nil)

	// get ue context
	ueCtxt, ok := ctxt.Value(types.UeContextCK).(*gctxt.UeContext)
	if !ok {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, nil,
			"fail to get ue context")
		return types.ErrFailFindUeCtxt
	}

	// send authentication reject message to UE
	msg := nasmsg.AuthenticationRejectMsg{}
	encBuf, err := msg.Encode()
	if err != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"fail to encode authentication reject message")
		return types.ErrFailEncodeNasMsg
	}

	err, _ = SendDownLinkNasMsg(ctxt, ueCtxt, encBuf)
	if err != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"fail to send downlink nas ngap msg")
		return types.ErrFailSendNgapMsg
	}

	//msg counter
	pm.PegCounter(statistics.AuthenticationRejectCounter)

	// procedure context set to nil
	ueCtxt.SetProcCtxt(nil)

	return nil
}

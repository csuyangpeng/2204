package mmsender

import (
	"context"
	"fmt"
	"lite5gc/amf/context/gctxt"
	"lite5gc/amf/sc/statistics"
	"lite5gc/cmn/message/nasmsg"
	"lite5gc/cmn/nas/nasie"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/oam/pm"
)

func SendIdentityRequest(ctxt context.Context, ueCtxt *gctxt.UeContext, idtype nasie.IdentityType) error {
	rlogger.FuncEntry(types.ModuleAmfNas, ueCtxt)

	idReqMsg := nasmsg.IdentityRequestMsg{}
	idReqMsg.IdentityType = idtype
	// encode nas message
	msgBytes, err := idReqMsg.Encode()
	if err != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"failed to encode identity request message,error(%s)", err)
		return types.ErrFailEncodeNasMsg
	}

	err, encNgapMsg := SendDownLinkNasMsg(ctxt, ueCtxt, msgBytes)
	if err != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"fail to send downlink nas ngap msg")
		return types.ErrFailSendNgapMsg
	}

	//msg counter
	pm.PegCounter(statistics.IdentifyRequestCounter)

	// start t3570 timer
	err = StartNasTimer(ctxt, ueCtxt, gctxt.T3570, encNgapMsg)
	if err != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"failed to start nas timer T3570, error(%s)", err)
		return fmt.Errorf("failed to start nas timer T3570, error(%s)", err)
	}

	return nil
}

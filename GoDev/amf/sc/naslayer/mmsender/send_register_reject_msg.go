package mmsender

import (
	"context"
	"fmt"
	"lite5gc/amf/context/gctxt"
	"lite5gc/amf/sc/mobmgnt/mmutils"
	"lite5gc/amf/sc/statistics"
	"lite5gc/cmn/message/nasmsg"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/oam/pm"
)

func SendRegisterRejectMsg(ctxt context.Context, cause nas.Mm5gCause) error {
	rlogger.FuncEntry(types.ModuleAmfNas, ctxt)

	ueCtxt, ok := ctxt.Value(types.UeContextCK).(*gctxt.UeContext)
	if !ok {
		rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, ctxt,
			"fail to get ue context")
		return types.ErrFailFindUeCtxt
	}

	rejectMsg := nasmsg.RegistrationRejectMsg{}
	rejectMsg.MMCause = cause
	// encode nas message
	bytes, err := rejectMsg.Encode()
	if err != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, ueCtxt,
			"failed to encode registration accept message")
		return fmt.Errorf("failed to encode registration accept message")
	}

	// add the security header
	procNasMsg := VerifyAndBuildSecProtectNasMsg(ueCtxt, bytes)

	err, _ = SendDownLinkNasMsg(ctxt, ueCtxt, procNasMsg)
	if err != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"fail to send downlink nas ngap msg")
		return types.ErrFailSendNgapMsg
	}

	//set state for ue
	err = ueCtxt.SetRmState(types.StateRmDeRegistered)
	if err != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, nil,
			"fail to set RM state")
		return fmt.Errorf("fail to set RM state")
	}

	//release ue ctxt
	err = mmutils.ReleaseUECtxt(ueCtxt)
	if err != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, nil, "fail to release ue context")
		return fmt.Errorf("fail to release ue context")
	}

	//msg counter
	pm.PegCounter(statistics.RegistrationRejectCounter)

	return nil
}

package naslayer

import (
	"context"
	"fmt"
	"lite5gc/amf/context/gctxt"
	"lite5gc/amf/context/prcdctxt"
	"lite5gc/amf/sc/statemgr"
	"lite5gc/cmn/message/nasmsg"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types/statetypes"
)

func HandlePduSessRelRequestMsg(ctx context.Context,
	ueCtxt *gctxt.UeContext,
	ulNasMsg *nasmsg.UplinkNasTransportMsg) error {

	rlogger.FuncEntry(types.ModuleAmfNas, ueCtxt)

	if ulNasMsg.OptIeBitSet.Test(nasmsg.IeidUplinknastransPdusessid) == false {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, ueCtxt, "missing psi. ul nas transport:%s", ulNasMsg)
		return fmt.Errorf("missing PSI for pdu session establishment")
		//reject msg todo
	}

	if !ueCtxt.IsPduSessExist(ulNasMsg.PduSessId) {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, ueCtxt, "pdu session is not exist(%d)", ulNasMsg.PduSessId)
		return fmt.Errorf("psi(%d) duplicated", ulNasMsg.PduSessId)
	}

	//allocate service request prcd ctxt
	sessRelPrcdCtxt, ok := ueCtxt.GetProcCtxt().(*prcdctxt.SessionReleasePrcdCtxt)
	if ok {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"failed to get SessionReleasePrcdCtxt, and :", ueCtxt.GetProcCtxt())
		return types.ErrFailGetProcedureCtxt
	} else {
		//allocate a new procedure context
		sessRelPrcdCtxt = prcdctxt.NewSessionRelPrcdCtxt() //StateSessRelStart
	}

	if ulNasMsg.OptIeBitSet.Test(nasmsg.IeidUplinknastransPdusessid) {
		sessRelPrcdCtxt.Psi = ulNasMsg.PduSessId
	}

	ueCtxt.SetProcCtxt(sessRelPrcdCtxt)
	ctx = context.WithValue(ctx, types.UeContextCK, ueCtxt)
	ctx = context.WithValue(ctx, types.UplinkMsgCK, ulNasMsg)

	// trigger FSM
	err := statemgr.TriggerFsm(ctx,
		statemgr.SessionRelease,
		sessRelPrcdCtxt.GetCurrentState(),
		statetype.EventSessRelReq)
	if err != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, nil, "fail to trigger fsm")
		return types.ErrFailTriggerFsm
	}
	return nil
}

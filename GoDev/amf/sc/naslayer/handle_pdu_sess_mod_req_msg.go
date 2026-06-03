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

func HandlePduSessModRequestMsg(ctx context.Context,
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
	var sessModPrcdCtxt *prcdctxt.SessionModificationPrcdCtxt

	sessModPrcdCtxt, ok := ueCtxt.GetProcCtxt().(*prcdctxt.SessionModificationPrcdCtxt)
	if ok {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"failed to get SessionModificationPrcdCtxt, and :", ueCtxt.GetProcCtxt())
		return types.ErrFailGetProcedureCtxt
	} else {
		//allocate a new procedure context
		sessModPrcdCtxt = prcdctxt.NewSessionModPrcdCtxt()
	}

	if ulNasMsg.OptIeBitSet.Test(nasmsg.IeidUplinknastransPdusessid) {
		sessModPrcdCtxt.Psi = ulNasMsg.PduSessId
	}

	ueCtxt.SetProcCtxt(sessModPrcdCtxt)
	ctx = context.WithValue(ctx, types.UeContextCK, ueCtxt)
	ctx = context.WithValue(ctx, types.UplinkMsgCK, ulNasMsg)

	// trigger FSM
	rlogger.Trace(types.ModuleAmfNas, rlogger.INFO, ueCtxt,
		"sessModPrcdCtxt.GetCurrentState()", sessModPrcdCtxt.GetCurrentState())
	err := statemgr.TriggerFsm(ctx,
		statemgr.SessionModify,
		sessModPrcdCtxt.GetCurrentState(),
		statetype.EventSessModWf1stUpSmCtxtReq)
	if err != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, nil, "fail to trigger fsm")
		return types.ErrFailTriggerFsm
	}

	return nil
}

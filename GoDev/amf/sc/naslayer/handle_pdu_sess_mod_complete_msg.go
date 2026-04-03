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

func HandleSessionModCompleteMsg(ctx context.Context,
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

	sessModPrcdCtxt, ok := ueCtxt.GetProcCtxt().(*prcdctxt.SessionModificationPrcdCtxt)
	if !ok {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"failed to get SessionModificationPrcdCtxt, and :", ueCtxt.GetProcCtxt())
		return types.ErrFailGetProcedureCtxt
	}

	ctx = context.WithValue(ctx, types.UeContextCK, ueCtxt)
	ctx = context.WithValue(ctx, types.UplinkMsgCK, ulNasMsg)
	rlogger.Trace(types.ModuleAmfNas, rlogger.INFO, ueCtxt, "sessModPrcdCtxt.GetCurrentState()", sessModPrcdCtxt.GetCurrentState())
	// trigger FSM
	err := statemgr.TriggerFsm(ctx,
		statemgr.SessionModify,
		statetype.StateSessModWfN2NASUplinkTrans,
		statetype.EventSessModWf3rdUpSmCtxtReq)
	if err != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, nil, "fail to trigger fsm")
		return types.ErrFailTriggerFsm
	}

	return nil
}

package ngaplayer

import (
	"context"
	"lite5gc/amf/context/gctxt"
	"lite5gc/amf/context/prcdctxt"
	"lite5gc/amf/sc/statemgr"
	"lite5gc/cmn/message/ngap/ngapmsg"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/timermgr"
	"lite5gc/cmn/types"
	statetype "lite5gc/cmn/types/statetypes"
)

func (p *LayerMgr) handleInitialContextSetupRespMsg(ctx context.Context, msgBuf []byte) error {
	rlogger.FuncEntry(types.ModuleAmfNgap, nil)

	message := ngapmsg.NewInitialContextSetupRespMsg()
	message.SetOssCodecCtxt(p.ossCtxt.GetOssCtxtPtr_m())
	err := message.Decode(msgBuf)
	if err != nil {
		rlogger.Trace(types.ModuleAmfNgap, rlogger.ERROR, nil,
			"failed to decode initial context setup response message")
		return err
	}

	ueCtxt, err := gctxt.GetUeContext(gctxt.AmfUeNgApId(message.AmfUeNGAPId))
	if err != nil {
		rlogger.Trace(types.ModuleAmfNgap, rlogger.ERROR, nil,
			"failed to find the ue context with AmfUeNGAPId(%d)",
			message.AmfUeNGAPId)
		return err
	}
	ctx = context.WithValue(ctx, types.UeContextCK, ueCtxt)

	//get procedure context
	prcdCtxt := ueCtxt.GetProcCtxt()

	//get the timer Mgr
	timerMgr, ok := ctx.Value(types.ScTimerMgrCK).(*timermgr.TimerMgr)
	if !ok {
		rlogger.Trace(types.ModuleAmfNgap, rlogger.ERROR, ueCtxt.GetImsiPtr(), "failed to get timer manager")
		return types.ErrFailFindTimerMgr
	}

	switch prcdCtxt.(type) {
	case *prcdctxt.RegistrationPrcdCtxt:
		pc := prcdCtxt.(*prcdctxt.RegistrationPrcdCtxt)
		rlogger.Trace(types.ModuleAmfNgap, rlogger.INFO, ueCtxt.GetImsiPtr(),
			"get in RegistrationPrcdCtxt", pc.TimerId)
		timerMgr.CancelTimer(pc.TimerId)
		//trigger the FSM
		err := statemgr.TriggerFsm(ctx,
			statemgr.Register,
			prcdCtxt.GetCurrentState(),
			statetype.EventRegisterInitCtxtSetupResp)
		if err != nil {
			rlogger.Trace(types.ModuleAmfNgap, rlogger.ERROR, ueCtxt.GetImsiPtr(), "fail to trigger fsm")
			return types.ErrFailTriggerFsm
		}
	case *prcdctxt.ServiceRequestPrcdCtxt:
		//store the message into procedure context
		pc := prcdCtxt.(*prcdctxt.ServiceRequestPrcdCtxt)
		rlogger.Trace(types.ModuleAmfNgap, rlogger.INFO, ueCtxt.GetImsiPtr(),
			"get in ServiceRequestPrcdCtxt", pc.TimerId)
		timerMgr.CancelTimer(pc.TimerId)
		for _, pdu := range message.PduSessResSetupRespList {
			pc.PduSessResSetupRespList = append(pc.PduSessResSetupRespList, *pdu)
		}
		for _, failedPdu := range message.PduSessResFailedToSetupList {
			pc.PduSessResFailedToSetupList = append(pc.PduSessResFailedToSetupList, *failedPdu)
		}
		rlogger.Trace(types.ModuleAmfNgap, rlogger.INFO, ueCtxt.GetImsiPtr(),
			"len of pdu session", len(ueCtxt.GetPDUSessionCtxts()))
		//trigger the FSM
		if len(ueCtxt.GetPDUSessionCtxts()) == 0 {
			err := statemgr.TriggerFsm(ctx,
				statemgr.ServiceRequest,
				prcdCtxt.GetCurrentState(),
				statetype.EventSrvReqInitCtxtSetupRespNoPsi)
			if err != nil {
				rlogger.Trace(types.ModuleAmfNgap, rlogger.ERROR, ueCtxt.GetImsiPtr(), "fail to trigger fsm")
				return types.ErrFailTriggerFsm
			}
		} else {
			err := statemgr.TriggerFsm(ctx,
				statemgr.ServiceRequest,
				prcdCtxt.GetCurrentState(),
				statetype.EventSrvReqInitCtxtSetupResp)
			if err != nil {
				rlogger.Trace(types.ModuleAmfNgap, rlogger.ERROR, ueCtxt.GetImsiPtr(), "fail to trigger fsm")
				return types.ErrFailTriggerFsm
			}
		}
	default:
		rlogger.Trace(types.ModuleAmfNgap, rlogger.DEBUG, ueCtxt.GetImsiPtr(), "currently unsupported procedure")
	}

	return nil
}

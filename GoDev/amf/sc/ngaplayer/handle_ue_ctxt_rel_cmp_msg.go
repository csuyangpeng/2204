package ngaplayer

import (
	"context"
	"lite5gc/amf/context/gctxt"
	"lite5gc/amf/context/prcdctxt"
	"lite5gc/amf/sc/statemgr"
	"lite5gc/cmn/message/ngap/ngapmsg"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	statetype "lite5gc/cmn/types/statetypes"
)

func (p *LayerMgr) handleUeCtxtReleaseCmpMsg(ctx context.Context, msgBuf []byte) error {
	rlogger.FuncEntry(types.ModuleAmfNgap, nil)

	message := ngapmsg.NewUeContextReleaseCmpMsg()
	message.SetOssCodecCtxt(p.ossCtxt.GetOssCtxtPtr_m())
	err := message.Decode(msgBuf)
	if err != nil {
		rlogger.Trace(types.ModuleAmfNgap, rlogger.ERROR, nil,
			"failed to decode ue context release complete message")
		return err
	}

	ueCtxt, err := gctxt.GetUeContext(gctxt.AmfUeNgApId(message.AmfUeNgapId))
	if err != nil {
		rlogger.Trace(types.ModuleAmfNgap, rlogger.ERROR, nil,
			"failed to find the ue context with AmfUeNgapId(%d)", message.AmfUeNgapId)
		return err
	}

	ctx = context.WithValue(ctx, types.UeContextCK, ueCtxt)

	//get procedure context
	prcdCtxt := ueCtxt.GetProcCtxt()
	if prcdCtxt != nil {
		rlogger.Trace(types.ModuleAmfNgap, rlogger.DEBUG, ueCtxt.GetImsiPtr(),
			"current prcd status(%s)", prcdCtxt.GetCurrentState())
		switch prcdCtxt.(type) {
		case *prcdctxt.DeRegistrationPrcdCtxt:
			//TODO store the message into procedure context
			//trigger the FSM
			err := statemgr.TriggerFsm(ctx,
				statemgr.DeRegister,
				prcdCtxt.GetCurrentState(),
				statetype.EventDeRegisterUeCtxtRelCmp)
			if err != nil {
				rlogger.Trace(types.ModuleAmfNgap, rlogger.ERROR, nil, "fail to trigger fsm")
				return types.ErrFailTriggerFsm
			}
		case *prcdctxt.AnReleasePrcdCtxt:
			//trigger the FSM
			err := statemgr.TriggerFsm(ctx,
				statemgr.AnRelease,
				prcdCtxt.GetCurrentState(),
				statetype.EventAnRelUeCtxtRelCmp)
			if err != nil {
				rlogger.Trace(types.ModuleAmfNgap, rlogger.ERROR, nil, "fail to trigger fsm")
				return types.ErrFailTriggerFsm
			}
		default:
			rlogger.Trace(types.ModuleAmfNgap, rlogger.DEBUG, nil, "igonor ue context release complete message")
		}
	}
	return nil
}

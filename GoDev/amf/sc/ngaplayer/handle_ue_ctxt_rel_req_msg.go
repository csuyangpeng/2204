package ngaplayer

import (
	"context"
	"fmt"
	"lite5gc/amf/context/gctxt"
	"lite5gc/amf/context/prcdctxt"
	"lite5gc/amf/context/prcdctxt/prcdutils"
	"lite5gc/amf/sc/statemgr"
	"lite5gc/cmn/message/ngap/ngapmsg"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	statetype "lite5gc/cmn/types/statetypes"
)

func (p *LayerMgr) handleUeCtxtReleaseReqMsg(ctx context.Context, msgBuf []byte) error {
	rlogger.FuncEntry(types.ModuleAmfNgap, nil)

	message := ngapmsg.NewUeContextReleaseReqMsg()
	message.SetOssCodecCtxt(p.ossCtxt.GetOssCtxtPtr_m())
	err := message.Decode(msgBuf)
	if err != nil {
		rlogger.Trace(types.ModuleAmfNgap, rlogger.ERROR, nil,
			"failed to decode ue context release request message")
		return err
	}

	ueCtxt, err := gctxt.GetUeContext(gctxt.AmfUeNgApId(message.AmfUeNgapId))
	if err != nil {
		rlogger.Trace(types.ModuleAmfNgap, rlogger.ERROR, nil,
			"failed to find the ue context with RanNgapId(%d)", message.RanUeNgapId)
		return err
	}

	if ueCtxt.GetAmfUeNgapId() != message.AmfUeNgapId {
		rlogger.Trace(types.ModuleAmfNgap, rlogger.DEBUG, ueCtxt.GetImsiPtr(),
			"Mismatch. AmfUeNgapId(core-%d,msg-%d), RanUeNgapId(core-%d,msg-%d",
			ueCtxt.GetAmfUeNgapId(), message.AmfUeNgapId,
			ueCtxt.GetRanUeNgapId(), message.RanUeNgapId)
		return fmt.Errorf("UeNgapId Mismatch")
	}

	// allocated procedure context
	anRelPrcdCtxt, ok := ueCtxt.GetProcCtxt().(*prcdctxt.AnReleasePrcdCtxt)
	if ok {
		rlogger.Trace(types.ModuleAmfNgap, rlogger.DEBUG, ueCtxt.GetImsiPtr(),
			"failed to get AnReleasePrcdCtxt, and :", ueCtxt.GetProcCtxt())
		// TODO there is a procedure unFinished.
		// TODO cancel the current procedure and process initial registration request
		//return fmt.Errorf("there are currently processes in progress")
	} else {
		err = prcdutils.CancelAmfTimer(ctx, ueCtxt)
		if err != nil {
			rlogger.Trace(types.ModuleAmfNgap, rlogger.DEBUG, ueCtxt.GetImsiPtr(), "fail to cancel timer")
		}
		//allocate the procedure context
		anRelPrcdCtxt = prcdctxt.NewAnReleasePrcdCtxtUETrigger()
		ueCtxt.SetProcCtxt(anRelPrcdCtxt)
	}

	ctx = context.WithValue(ctx, types.UeContextCK, ueCtxt)

	//store the info from message into procedure context
	anRelPrcdCtxt.ReleasePsiList = append(anRelPrcdCtxt.ReleasePsiList, message.PduSessIdList...)
	anRelPrcdCtxt.SetCounter(len(message.PduSessIdList))

	err = statemgr.TriggerFsm(ctx,
		statemgr.AnRelease,
		anRelPrcdCtxt.GetCurrentState(),
		statetype.EventAnRelUeCtxtRelReq)
	if err != nil {
		rlogger.Trace(types.ModuleAmfNgap, rlogger.ERROR, nil, "fail to trigger fsm")
		return types.ErrFailTriggerFsm
	}

	return nil
}

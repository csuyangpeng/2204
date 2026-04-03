package n11layer

import (
	"context"
	"fmt"
	"lite5gc/amf/context/gctxt"
	"lite5gc/amf/context/prcdctxt"
	"lite5gc/amf/sc/statemgr"
	"lite5gc/amf/sc/statistics"
	"lite5gc/cmn/message/n11msg"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/sbicmn"
	"lite5gc/cmn/types"
	statetype "lite5gc/cmn/types/statetypes"
	"lite5gc/cmn/types3gpp"
	"lite5gc/oam/pm"
)

func HandleUpdateSmCtxtResponse(ctx context.Context, smResp *sbicmn.SbiPostModifySmContext) error {
	rlogger.FuncEntry(types.ModuleAmfSbi, nil)

	//get the ueCtxt from ctxt
	imsi, psi, err := nas.GetSmfKeys(smResp.SmContextRef)
	if err != nil {
		rlogger.Trace(types.ModuleAmfSbi, rlogger.ERROR, nil, "failed to get imsi and psi from url(%s)", smResp.SmContextRef)
		return fmt.Errorf("failed to get imsi and psi from url(%s)", smResp.SmContextRef)
	}

	ueCtxt, err := gctxt.GetUeContext(gctxt.ImsiKey(imsi.GetValue()))
	if err != nil {
		rlogger.Trace(types.ModuleAmfSbi, rlogger.ERROR, nil, "failed to find the ue context with AmfUeNgApId %s,error %s",
			imsi, err)
		return fmt.Errorf("failed to find the ue context with AmfUeNgApId %s,error %s",
			imsi, err)
	}
	// store the ue context in context
	ctx = context.WithValue(ctx, types.UeContextCK, ueCtxt)

	//get the context Mgr
	stateMgr, ok := ctx.Value(types.StateMgrCK).(*statemgr.StateMgr)
	if !ok {
		rlogger.Trace(types.ModuleAmfSbi, rlogger.ERROR, ueCtxt, "get state manager failed")
		return fmt.Errorf("get state manager failed")
	}

	var pduStatus types3gpp.PduSessStatus
	updateSmRespMsg := sbicmn.Trans_ModelsToN11_SmContextUpdatedDataFormat(smResp.RespData)
	if updateSmRespMsg.IeFlags.Test(n11msg.Ieid_upCnxState) {
		pduStatus = types3gpp.PduSessStatus(updateSmRespMsg.UpCnxState)
		rlogger.Trace(types.ModuleAmfSbi, rlogger.DEBUG, ueCtxt, "upCnxState(%s) in updateSmCtxtResp message", pduStatus)
	}

	//get procedure context
	prcdCtxt := ueCtxt.GetProcCtxt()
	switch prcdCtxt.(type) {
	case *prcdctxt.AnReleasePrcdCtxt:
		rlogger.Trace(types.ModuleAmfSbi, rlogger.DEBUG, ueCtxt, "prcdctxt is AnReleasePrcdCtxt")
		//store procedure ctxt
		pctxt := prcdCtxt.(*prcdctxt.AnReleasePrcdCtxt)
		pctxt.UpCnxState = pduStatus
		//pctxt.AmfSmCtxtId = msg.AmfSmCtxtId

		//trigger the FSM
		stateMgr.AnReleaseFsm.Bfsm.SetState(prcdCtxt.GetCurrentState())
		stateMgr.AnReleaseFsm.Bfsm.Event(statetype.EventAnRelUpdateSmCtxtAck, ctx)
	case *prcdctxt.ServiceRequestPrcdCtxt:
		rlogger.Trace(types.ModuleAmfSbi, rlogger.DEBUG, ueCtxt, "prcdctxt is ServiceRequestPrcdCtxt")
		//store procedure ctxt
		pctxt := prcdCtxt.(*prcdctxt.ServiceRequestPrcdCtxt)
		pctxt.UpCnxState = pduStatus
		//pctxt.AmfSmCtxtId = msg.AmfSmCtxtId
		if updateSmRespMsg.IeFlags.Test(n11msg.Ieid_n2SmInfo) {
			pctxt.N2SmInfo[uint32(psi)] = []byte(updateSmRespMsg.N2SmInfo)
			rlogger.Trace(types.ModuleAmfSbi, rlogger.DEBUG, ueCtxt, "counter,pctxt.Order, n2smInfo:",
				pctxt.GetCounter(), pctxt.Order, pctxt.N2SmInfo)
		}
		//trigger the FSM
		stateMgr.ServiceRequestFsm.Bfsm.SetState(prcdCtxt.GetCurrentState())
		if pctxt.Order == prcdctxt.SerReqUpDataSmCtxtRespFirst {
			stateMgr.ServiceRequestFsm.Bfsm.Event(statetype.EventSrvReqUpdateSmCtxtResp, ctx)
		} else if pctxt.Order == prcdctxt.SerReqUpDataSmCtxtRespSecond {
			stateMgr.ServiceRequestFsm.Bfsm.Event(statetype.EventSrvReqUpdateSmCtxtRespSec, ctx)
		} else {
			rlogger.Trace(types.ModuleAmfSbi, rlogger.DEBUG, ueCtxt, "invalid order(%d)", pctxt.Order)
		}
	case *prcdctxt.SessionReleasePrcdCtxt:
		rlogger.Trace(types.ModuleAmfSbi, rlogger.DEBUG, ueCtxt, "prcdctxt is SessionReleasePrcdCtxt")
		//store procedure ctxt
		pctxt := prcdCtxt.(*prcdctxt.SessionReleasePrcdCtxt)
		pctxt.UpCnxState = pduStatus
		//pctxt.AmfSmCtxtId = msg.AmfSmCtxtId

		ctx = context.WithValue(ctx, types.SmfN11MsgDataCK, updateSmRespMsg)

		stateMgr.SessRelRequestFsm.Bfsm.SetState(prcdCtxt.GetCurrentState())
		switch prcdCtxt.GetCurrentState() {
		case statetype.StateSessRelWfUpSmCtxtResp:
			rlogger.Trace(types.ModuleAmfSbi, rlogger.DEBUG, ueCtxt, "1st UpSmCtxt Resp")
			stateMgr.SessRelRequestFsm.Bfsm.Event(statetype.EventSessRelUpSmCtxtResp, ctx)
		case  statetype.StateSessRelWfUpSmCtxtReqSec:
			rlogger.Trace(types.ModuleAmfSbi, rlogger.DEBUG, ueCtxt, "2nd UpSmCtxt Req Sec")
			stateMgr.SessRelRequestFsm.Bfsm.Event(statetype.EventSessRelUpSmCtxtReqSec, ctx)
		case statetype.StateSessRelWfUpSmCtxtRespSec:
			rlogger.Trace(types.ModuleAmfSbi, rlogger.DEBUG, ueCtxt, "2nd UpSmCtxt Resp")
			stateMgr.SessRelRequestFsm.Bfsm.Event(statetype.EventSessRelEnd, ctx)
		default:
			rlogger.Trace(types.ModuleAmfSbi, rlogger.DEBUG, ueCtxt, "unknown prcdCtxt statetype",prcdCtxt.GetCurrentState())
		}
	case *prcdctxt.SessionModificationPrcdCtxt:
		rlogger.Trace(types.ModuleAmfSbi, rlogger.DEBUG, ueCtxt, "prcdctxt is Session mod PrcdCtxt")
		//store procedure ctxt
		pctxt := prcdCtxt.(*prcdctxt.SessionModificationPrcdCtxt)
		pctxt.UpCnxState = pduStatus
		//pctxt.AmfSmCtxtId = msg.AmfSmCtxtId

		//msgData := msg.MsgData.(n11msg.UpdateSMContextResponseData)
		ctx = context.WithValue(ctx, types.SmfN11MsgDataCK, updateSmRespMsg)

		stateMgr.SessModRequestFsm.Bfsm.SetState(prcdCtxt.GetCurrentState())
		rlogger.Trace(types.ModuleAmfSbi, rlogger.DEBUG, ueCtxt, "prcdCtxt.GetCurrentState()", prcdCtxt.GetCurrentState())
		switch prcdCtxt.GetCurrentState() {
		case statetype.StateSessModWf1stUpSmCtxtResp:
			stateMgr.SessModRequestFsm.Bfsm.Event(statetype.EventSessModWfN2SessReq, ctx)
		case statetype.StateSessModWf3rdUpSmCtxtResp:
			stateMgr.SessModRequestFsm.Bfsm.Event(statetype.EventSessModEnd, ctx)
		default:
			rlogger.Trace(types.ModuleAmfSbi, rlogger.DEBUG, ueCtxt, "do nothing for UpSmCtxtRespSec in session release")
		}
	default:
		rlogger.Trace(types.ModuleAmfSbi, rlogger.DEBUG, ueCtxt, "prcdCtxt:(%v)", prcdCtxt)
		pduSessCtxt := ueCtxt.GetPduSessCtxt(nas.PduSessID(psi))

		pduSessCtxt.Status = types3gpp.PduSessStatus(types3gpp.SessActived)
		ueCtxt.PduSessStatus = ueCtxt.GetUePsiStatus()

		//仅在该UE建立首个会话时设置即可，后面的会话无需重复设置
		if ueCtxt.GetCmState() == types.CmIdle {
			ueCtxt.SetCmState(types.CmConnected)
			pm.PegCounter(statistics.ActiveUserCounter)
		}

		rlogger.Trace(types.ModuleAmfSbi, rlogger.DEBUG, ueCtxt, "set the pdu status to (%s)", pduSessCtxt.Status)
		rlogger.Trace(types.ModuleAmfSbi, rlogger.DEBUG, ueCtxt, "PDU Session psi( %d ) Established!", pduSessCtxt.Psi)

	}

	return nil
}


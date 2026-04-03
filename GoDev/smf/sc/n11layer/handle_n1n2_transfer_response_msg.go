package n11layer

import (
	"context"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/sbicmn"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types/statetypes"
	"lite5gc/cmn/types3gpp"
	"lite5gc/smf/sc/statemgr"
	"lite5gc/smf/smfcontext/gctxt"
	"lite5gc/smf/smfcontext/prcdctxt"
	"strings"
)

func HandlePostN1N2MsgTransferResp(ctxt context.Context,
	msgData *sbicmn.SbiPostN1N2MsgTransferMsg) error {
	rlogger.FuncEntry(types.ModuleSmfN11, msgData)
	if msgData != nil {
		imsiStr := strings.TrimPrefix(msgData.Supi, "imsi-")
		imsi := &types3gpp.Imsi{}
		err := imsi.StoreImsiString(imsiStr, types3gpp.CheckMncLen(imsiStr))
		if err != nil {
			rlogger.Trace(types.ModuleSmfN11, rlogger.ERROR, nil, "failed to get imsi")
			return types.ErrFailGetImsi
		}

		ueCtxt, err := gctxt.GetUeContext(gctxt.ImsiKey(imsi.GetValue()))
		if err != nil {
			// no Ue context, create a new ue context
			rlogger.Trace(types.ModuleSmfN11, rlogger.ERROR, nil,
				"failed to get ue context with imsi(%s)", imsi)
			return types.ErrFailFindUeCtxt
		}

		pduSessCtxt, ok := ueCtxt.PduSessCtxts[msgData.Psi]
		if !ok {
			rlogger.Trace(types.ModuleSmfN11, rlogger.ERROR, ueCtxt.GetImsiPtr(),
				"failed to get pdu session context with psi(%d)", msgData.Psi)
			return types.ErrFailFindSessionCtxt
		}

		// get procedure context
		procCtxt := pduSessCtxt.GetPrcdCtxt()
		switch procCtxt.(type) {
		case *prcdctxt.PduSessionEstbPrcdCtxt:
			pCtxt := procCtxt.(*prcdctxt.PduSessionEstbPrcdCtxt)
			//store info in procedure context
			msg := sbicmn.Trans_ModelsToN11_N1N2MsgTransferRespFormat(msgData.RespData)
			pCtxt.Cause = msg.Cause
			pCtxt.SupFeatures = msg.SupportedFeatures

			ctxt = context.WithValue(ctxt, types.SmfPduSessCtxtCK, pduSessCtxt)

			err := statemgr.TriggerSmfFsm(ctxt,
				statemgr.SessionESTB,
				pCtxt.GetCurrentState(),
				statetype.EventPduSessEstbN1N2MsgTransferResp)
			if err != nil {
				rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, nil, "fail to trigger fsm")
				return types.ErrFailTriggerFsm
			}
		default:
			rlogger.Trace(types.ModuleSmfN11, rlogger.DEBUG, ueCtxt.GetImsiPtr(),
				"unknown procedure for n1n2 msg transfer response")
			//TODO handle for other scenarios
		}
	} else {
		rlogger.Trace(types.ModuleSmfN11, rlogger.ERROR, nil, "input para is nil")
		return types.ErrInputParaNil
	}
	return nil
}

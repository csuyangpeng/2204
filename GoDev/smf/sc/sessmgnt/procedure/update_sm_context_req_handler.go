package procedure

import (
	"context"
	"lite5gc/cmn/message/n11msg"
	"lite5gc/cmn/message/ngap/ngapmsg"
	"lite5gc/cmn/message/pfcp"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types/statetypes"
	"lite5gc/smf/sc/sessmgnt"
	"lite5gc/smf/sc/sessmgnt/smsender"
	"lite5gc/smf/smfcontext/gctxt"
	"lite5gc/smf/smfcontext/prcdctxt"
)

//for session estb
func HandleUpdateSmCtxtRequestMsg4Estb(ctxt context.Context) {
	rlogger.FuncEntry(types.ModuleSmfSM, ctxt)

	// pdu session context
	pduSessCtxt, ok := ctxt.Value(types.SmfPduSessCtxtCK).(*gctxt.PduSessContext)
	if !ok {
		rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, ctxt, "failed to get pdu session context")
		return
	}

	// sm ctxt update msg data
	smCtxtUpdateData := ctxt.Value(types.SmfUpdateSmCtxtReqMsgCK).(*n11msg.UpdateSMContextRequestData)

	sessMgnt, ok := ctxt.Value(types.SessMgntCK).(*sessmgnt.SessMGMT)
	if !ok {
		rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, pduSessCtxt, "failed to get session mangement")
		return
	}

	pCtxt, ok := pduSessCtxt.GetPrcdCtxt().(*prcdctxt.PduSessionEstbPrcdCtxt)
	if !ok {
		rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, pduSessCtxt,
			"prcdctxt is not PduSessionEstbPrcdCtxt, but:",
			gctxt.PrcdCtxt2String(pduSessCtxt.GetPrcdCtxt()))
		return
	}
	pduState := n11msg.INAVLID
	// SMF to UPF: send N4 session establish update request msg to UPF
	// The SMF provides AN Tunnel Info to the UPF as well as the corresponding forwarding rules.
	if smCtxtUpdateData.IeFlags.Test(n11msg.Ieid_n2SmInfo) {

		var resourceResp ngapmsg.PduSessResSetupRespTransfer
		resourceResp.SetOssCodecCtxt(sessMgnt.GetOssCtxt().GetOssCtxtPtr_m())
		resourceResp.Decode([]byte(smCtxtUpdateData.N2SmInfo))

		switch smCtxtUpdateData.N2SmInfoType {
		case n11msg.PduResSetupRsp:
			smsender.SendPduResModifyReq2Upf(ctxt, pduSessCtxt, pfcp.PFCP_Session_Modification_Request, &resourceResp)
			pduState = n11msg.ACTIVATED

		case n11msg.PduResSetupFail:
			if smCtxtUpdateData.UpCnxState == n11msg.DEACTIVATED {
				smsender.SendPfcpSessDelReqMsg(ctxt, pduSessCtxt, pfcp.PFCP_Session_Deletion_Request)
				pduState = n11msg.DEACTIVATED
			}
			return
		default:
			rlogger.Trace(types.ModuleSmfSM, rlogger.DEBUG, pduSessCtxt, "not process yet")
			return
		}
	}

	pCtxt.PduState = pduState
	// set the next state
	err := pCtxt.SetNextState(statetype.StatePduSessEstbWfN4SessModResp)
	if err != nil {
		rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, pduSessCtxt.IMSI, "fail to set state in sess estb proc")
		return
	}

	return
}

func HandleUpdateSmCtxtRequestMsg4SerReqPrcd(ctxt context.Context) {
	rlogger.FuncEntry(types.ModuleSmfSM, ctxt)

	// pdu session context
	pduSessCtxt, ok := ctxt.Value(types.SmfPduSessCtxtCK).(*gctxt.PduSessContext)
	if !ok {
		rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, ctxt, "failed to get pdu session context")
		return
	}

	smsender.SendActivatePfcpSessModRequest(ctxt, pduSessCtxt)

	// pdu sess estb procedure context
	pCtxt, ok := pduSessCtxt.GetPrcdCtxt().(*prcdctxt.AnRelSerReqPrcdCtxt)
	if !ok {
		rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, pduSessCtxt,
			"the prcdctxt is not AnRelSerReqPrcdCtxt,but:",
			gctxt.PrcdCtxt2String(pduSessCtxt.GetPrcdCtxt()))
		return
	}
	err := pCtxt.SetNextState(statetype.StatePduSessSerReqWfN4ModResp)
	if err != nil {
		rlogger.Trace(types.ModuleSmfNas, rlogger.ERROR, nil, "fail to set state")
		return
	}
}

func HandlerUpdataSmCtxtRequestMsg4AnRelease(ctxt context.Context) {
	rlogger.FuncEntry(types.ModuleSmfSM, ctxt)

	// pdu session context
	pduSessCtxt, ok := ctxt.Value(types.SmfPduSessCtxtCK).(*gctxt.PduSessContext)
	if !ok {
		rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, ctxt, "failed to get pdu session context")
		return
	}

	// sm ctxt update msg data
	smCtxtUpdateData := ctxt.Value(types.SmfUpdateSmCtxtReqMsgCK).(*n11msg.UpdateSMContextRequestData)
	pduSessCtxt.UpCnxState = smCtxtUpdateData.UpCnxState

	smsender.DeactivePduSession(ctxt, pduSessCtxt, pfcp.PFCP_Session_Modification_Request)
}

func HandleUpdateSmCtxtRequestSecMsg4SerReqPrcd(ctxt context.Context) {
	rlogger.FuncEntry(types.ModuleSmfSM, ctxt)

	// pdu session context
	pduSessCtxt, ok := ctxt.Value(types.SmfPduSessCtxtCK).(*gctxt.PduSessContext)
	if !ok {
		rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, ctxt, "failed to get pdu session context")
		return
	}

	pCtxt, ok := pduSessCtxt.GetPrcdCtxt().(*prcdctxt.AnRelSerReqPrcdCtxt)
	if !ok {
		rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, pduSessCtxt,
			"the prcdctxt is not AnRelSerReqPrcdCtxt,but:",
			gctxt.PrcdCtxt2String(pduSessCtxt.GetPrcdCtxt()))
		return
	}

	// sm ctxt update msg data
	smCtxtUpdateData := ctxt.Value(types.SmfUpdateSmCtxtReqMsgCK).(*n11msg.UpdateSMContextRequestData)

	if smCtxtUpdateData.IeFlags.Test(n11msg.Ieid_n2SmInfo) {

		sessMgnt, ok := ctxt.Value(types.SessMgntCK).(*sessmgnt.SessMGMT)
		if !ok {
			rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, pduSessCtxt, "failed to get session mangement")
			return
		}

		var resourceResp ngapmsg.PduSessResSetupRespTransfer

		resourceResp.SetOssCodecCtxt(sessMgnt.GetOssCtxt().GetOssCtxtPtr_m())

		resourceResp.Decode([]byte(smCtxtUpdateData.N2SmInfo))

		switch smCtxtUpdateData.N2SmInfoType {
		case n11msg.PduResSetupRsp:
			pduSessCtxt.UpCnxState = n11msg.ACTIVATED
			smsender.SendPduResModifyReq2Upf(ctxt, pduSessCtxt, pfcp.PFCP_Session_Modification_Request, &resourceResp)

			err := pCtxt.SetNextState(statetype.StatePduSessSerReqWfN4ModRespSec)
			if err != nil {
				rlogger.Trace(types.ModuleSmfNas, rlogger.ERROR, nil, "fail to set state")
				return
			}
		case n11msg.PduResSetupFail:
			if smCtxtUpdateData.UpCnxState == n11msg.DEACTIVATED {
				smsender.DeactivePduSession(ctxt, pduSessCtxt, pfcp.PFCP_Session_Deletion_Request)
			}
			err := pCtxt.SetNextState(statetype.StatePduSessSerReqWfN4ModRespSec)
			if err != nil {
				rlogger.Trace(types.ModuleSmfNas, rlogger.ERROR, nil, "fail to set state")
				return
			}
		default:
			rlogger.Trace(types.ModuleSmfSM, rlogger.DEBUG, pduSessCtxt, "not process yet")
		}
	} else {
		rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, ctxt, "no n2sminfo ie in update sm context reqeust message.")
	}
}

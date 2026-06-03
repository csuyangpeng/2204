package naslayer

import (
	"bytes"
	"context"
	"lite5gc/cmn/message/pfcp"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/nas/nasie"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/sbicmn"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types/statetypes"
	"lite5gc/smf/sc/sessmgnt/smutility"
	"lite5gc/smf/smfcontext/gctxt"
	"lite5gc/smf/smfcontext/prcdctxt"
)

func (p *NasMgr) HandleSmPduSessModRequest(ctx context.Context,
	header *nas.SmNasMessageHeader,
	plainNasMsg *bytes.Reader) nas.Sm5gCause {

	rlogger.FuncEntry(types.ModuleSmfNas, nil)

	if plainNasMsg == nil || header == nil {
		rlogger.Trace(types.ModuleSmfNas, rlogger.ERROR, nil, "input para is nil")
		return nas.OtherValue
	}
	cause := nas.SuccessNoReason
	// get ue context
	ueCtxt, ok := ctx.Value(types.UeContextCK).(*gctxt.UeContext)
	if !ok {
		rlogger.Trace(types.ModuleSmfNas, rlogger.ERROR, nil, "failed to get ue context")
		return nas.OtherValue
	}
	//decode the pdu session establish request message
	p.PduSessModRequest.Reset()
	p.PduSessModRequest.MsgHeader.PduSessionID = header.PduSessionID
	p.PduSessModRequest.MsgHeader.PrcdTransactionID = header.PrcdTransactionID
	p.PduSessModRequest.MsgHeader.MessageType = nas.PduSessionModRequest
	err := p.PduSessModRequest.Decode(plainNasMsg)
	if err != nil {
		rlogger.Trace(types.ModuleSmfNas, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"failed to decode pdu session modification message, err: ", err)
		return nas.OtherValue
	}

	//get the pdu session ctxt
	pduSessCtxt, ok := ctx.Value(types.SmfPduSessCtxtCK).(*gctxt.PduSessContext)
	if !ok {
		rlogger.Trace(types.ModuleSmfNas, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"failed to get pdu session context")
		return nas.OtherValue
	}

	sbimsg, _ := ctx.Value(types.SmfSbiHandlerMsgCK).(*sbicmn.SbiHandlerMessage)

	// check pti
	pti := header.PrcdTransactionID
	if pti < nas.MinPTI || pti > nas.MaxPTI {
		rlogger.Trace(types.ModuleSmfNas, rlogger.ERROR, ueCtxt.GetImsiPtr(), "pti is out of range")
		return nas.InvalidPTIValue
	}

	//create a pdu session modification prcd ctxt
	pCtxt := prcdctxt.NewPduSessionModPrcdCtxt(header.PduSessionID)
	pduSessCtxt.SetPrcdCtxt(pCtxt)
	pCtxt.IMSI = ueCtxt.IMSI
	pCtxt.DNN = pduSessCtxt.DNN
	pCtxt.SNSSAI = pduSessCtxt.SNSSAI
	pCtxt.UEIP = pduSessCtxt.UEIP
	pCtxt.SbiMessage = sbimsg

	msgData := &gctxt.ScN4MsgData{}
	msgData.IMSI = ueCtxt.IMSI
	//msgData.DNN = pduSessCtxt.DNN
	msgData.SNSSAI = pduSessCtxt.SNSSAI
	//msgData.UEIP = pduSessCtxt.UEIP

	switch p.PduSessModRequest.RequestQosRules.QoSRules[0].RuleOprCode {
	case nasie.CreateNewQoSRule:
		msgData.PFCPParameters, err = TransCreateQosRulesToPFCPParams(p.PduSessModRequest)
		if err != nil {
			rlogger.Trace(types.ModuleSmfNas, rlogger.INFO, ueCtxt.GetImsiPtr(),
				"fail to Trans CreateQosRules To PFCPParams")
			return nas.OtherValue
		}
		rlogger.Trace(types.ModuleSmfNas, rlogger.INFO, ueCtxt.GetImsiPtr(),
			"create pduSessCtxt.PFCPParameters", msgData.PFCPParameters)
	case nasie.DeleteExistingQoSRule:
		msgData.PFCPParameters = TransDeleteQosRulesToPFCPParams(p.PduSessModRequest)
		rlogger.Trace(types.ModuleSmfNas, rlogger.INFO, ueCtxt.GetImsiPtr(),
			"delete pduSessCtxt.PFCPParameters", msgData.PFCPParameters)
	default:
	}

	err = smutility.SendMsg2PfcpNode(ctx, pduSessCtxt, msgData, pfcp.PFCP_Session_Modification_Request)
	if err != nil {
		rlogger.Trace(types.ModuleSmfNas, rlogger.ERROR, ueCtxt.GetImsiPtr(), "fail to send msg to pfcp node")
		return nas.OtherValue
	}

	// set the next state
	err = pCtxt.SetNextState(statetype.StatePduSessModWf1stN4SessModResp)
	if err != nil {
		rlogger.Trace(types.ModuleSmfNas, rlogger.ERROR, ueCtxt.GetImsiPtr(), "fail to set state")
		return nas.OtherValue
	}
	return cause
}

package sbismf

import (
	"context"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/router"
	"lite5gc/cmn/sbicmn"
	"lite5gc/cmn/types"
)

func (p *SbiSmf) HandleSmfScMsg(ctx context.Context, msg *router.DataMsg) error {
	rlogger.FuncEntry(types.ModuleSmfSbi, ctx)
	if msg != nil {
		go func(ctx context.Context, msg router.IpcMsgData) {
			err := p.HandlerIncomingSmfScMsg(ctx, msg)
			if err != nil {
				rlogger.Trace(types.ModuleSmfSbi, rlogger.ERROR, ctx, "failed to HandlerIncomingSmfScMsg, err(%s) ", err)
			}
		}(ctx, msg.MsgData)
	} else {
		rlogger.Trace(types.ModuleSmfSbi, rlogger.ERROR, nil, "input para is nil")
		return types.ErrInputParaNil
	}
	return nil
}

func (p *SbiSmf) HandlerIncomingSmfScMsg(ctx context.Context, msg router.IpcMsgData) (err error) {
	rlogger.FuncEntry(types.ModuleSmfSbi, ctx)
	message, ok := msg.(*sbicmn.SbiMessage)
	if !ok {
		rlogger.Trace(types.ModuleSmfSbi, rlogger.ERROR, ctx, "invalid IPC message - SmfSc2SbiMsg")
		return
	}
	switch message.MsgType {
	case sbicmn.GetSmDataMsgRequest:
		rlogger.Trace(types.ModuleSmfSbi, rlogger.DEBUG, ctx, "GetSmDataMsgRequest msg")
		getSmMsgReq, ok := message.MsgData.(*sbicmn.SbiGetSmDataMsg)
		if !ok {
			rlogger.Trace(types.ModuleSmfSbi, rlogger.ERROR, ctx, "incorrect message GetSmDataMsgRequest")
			return
		}
		err := p.handleGetSMDataRequest(message.ScInstId, getSmMsgReq)
		if err != nil {
			rlogger.Trace(types.ModuleSmfSbi, rlogger.DEBUG, ctx, "failed to handle GetSmDataMsgRequest, error(%s)", err)
		}
	case sbicmn.N1N2MessageTransferReq:
		rlogger.Trace(types.ModuleSmfSbi, rlogger.DEBUG, ctx, "N1N2MessageTransferRequest msg")
		n1n2MsgTransfer, ok := message.MsgData.(*sbicmn.SbiPostN1N2MsgTransferMsg)
		if !ok {
			rlogger.Trace(types.ModuleSmfSbi, rlogger.ERROR, ctx, "incorrect message N1N2MessageTransferRequest")
			return
		}
		err := p.handlePostN1N2MsgTransferRequest(message.ScInstId, n1n2MsgTransfer)
		if err != nil {
			rlogger.Trace(types.ModuleSmfSbi, rlogger.DEBUG, ctx, "failed to handle N1N2MessageTransferRequest, error(%s)", err)
		}
	default:
	}

	return nil
}

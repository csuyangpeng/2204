package sbiamf

import (
	"context"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/router"
	"lite5gc/cmn/sbicmn"
	"lite5gc/cmn/types"
)

func (p *SbiAmf) HandleScMsg(ctx context.Context, msg *router.DataMsg) error {
	rlogger.FuncEntry(types.ModuleAmfSbi, ctx)

	go func(ctx context.Context, msg router.IpcMsgData) {
		err := p.HandlerIncomingScMsg(ctx, msg)
		if err != nil {
			rlogger.Trace(types.ModuleAmfSbi, rlogger.ERROR, ctx, "failed to HandlerIncomingScMsg, err(%s) ", err)
		}
	}(ctx, msg.MsgData)

	return nil
}

func (p *SbiAmf) HandlerIncomingScMsg(ctx context.Context, msg router.IpcMsgData) (err error) {
	rlogger.FuncEntry(types.ModuleAmfSbi, ctx)
	message, ok := msg.(*sbicmn.SbiMessage)
	if !ok {
		rlogger.Trace(types.ModuleAmfSbi, rlogger.ERROR, ctx, "invalid IPC message - Sc2SbiMsg")
		return
	}

	switch message.MsgType {
	case sbicmn.GetAmDataMsgRequest:
		rlogger.Trace(types.ModuleAmfSbi, rlogger.DEBUG, ctx, "GetAmDataMsgRequest msg")
		getAmdMsgReq, ok := message.MsgData.(*sbicmn.SbiGetAmDataMsg)
		if !ok {
			rlogger.Trace(types.ModuleAmfSbi, rlogger.ERROR, ctx, "incorrect message GetAmDataMsgRequest")
			return
		}
		err := p.HandleGetAMDataRequest(message.ScInstId, getAmdMsgReq)
		if err != nil {
			rlogger.Trace(types.ModuleAmfSbi, rlogger.DEBUG, ctx, "failed to handleGetAMDataRequest, error(%s)", err)
		}
	case sbicmn.GetSmfSelDataMsgRequest:
		rlogger.Trace(types.ModuleAmfSbi, rlogger.DEBUG, ctx, "GetSmfSelectDataMsgRequest msg")
		getSmfSelMsgReq, ok := message.MsgData.(*sbicmn.SbiGetSmfSelDataMsg)
		if !ok {
			rlogger.Trace(types.ModuleAmfSbi, rlogger.ERROR, ctx, "incorrect message GetSmfSelectDataMsgRequest")
			return
		}
		err := p.HandleGetSmfSelDataRequest(message.ScInstId, getSmfSelMsgReq)
		if err != nil {
			rlogger.Trace(types.ModuleAmfSbi, rlogger.DEBUG, ctx, "failed to handle GetSmfSelectDataMsgRequest, error(%s)", err)
		}
	case sbicmn.GetAuthDataMsgRequest:
		rlogger.Trace(types.ModuleAmfSbi, rlogger.DEBUG, ctx, "GetAuthDataMsgRequest msg")
		getAuthMsgReq, ok := message.MsgData.(*sbicmn.SbiGetAuthDataMsg)
		if !ok {
			rlogger.Trace(types.ModuleAmfSbi, rlogger.ERROR, ctx, "incorrect message GetAuthDataMsgRequest")
			return
		}
		err := p.HandleGetAuthDataRequest(message.ScInstId, getAuthMsgReq)
		if err != nil {
			rlogger.Trace(types.ModuleAmfSbi, rlogger.DEBUG, ctx, "failed to GetAuthDataMsgRequest, error(%s)", err)
		}
	case sbicmn.PostAmf3gppAccessRegistration:
		rlogger.Trace(types.ModuleAmfSbi, rlogger.DEBUG, ctx, "PostAmf3gppAccessRegistration msg")
		postAmfRegMsg, ok := message.MsgData.(*sbicmn.SbiPostAmf3gppAccessRegistration)
		if !ok {
			rlogger.Trace(types.ModuleAmfSbi, rlogger.ERROR, ctx, "incorrect message SbiPostAmf3gppAccessRegistration")
			return
		}
		err := p.HandlePostAmfRegistration(postAmfRegMsg)
		if err != nil {
			rlogger.Trace(types.ModuleAmfSbi, rlogger.DEBUG, ctx, "failed to SbiPostAmf3gppAccessRegistration, error(%s)", err)
		}
	case sbicmn.PostSdmSubscription:
		rlogger.Trace(types.ModuleAmfSbi, rlogger.DEBUG, ctx, "PostSdmSubscription msg")
		postSdmSubMsg, ok := message.MsgData.(*sbicmn.SbiPostSdmSubscription)
		if !ok {
			rlogger.Trace(types.ModuleAmfSbi, rlogger.ERROR, ctx, "incorrect message PostSdmSubscription")
			return
		}
		err := p.HandlePostSdmSubscription(message.ScInstId, postSdmSubMsg)
		if err != nil {
			rlogger.Trace(types.ModuleAmfSbi, rlogger.DEBUG, ctx, "failed to PostSdmSubscription, error(%s)", err)
		}
	case sbicmn.PduSessCreateSMContextReq:
		rlogger.Trace(types.ModuleAmfSbi, rlogger.DEBUG, ctx, "PostCreateSmContext msg")
		postCreatSmMsgReq, ok := message.MsgData.(*sbicmn.SbiPostCreateSmContext)
		if !ok {
			rlogger.Trace(types.ModuleAmfSbi, rlogger.ERROR, ctx, "incorrect message PostCreateSmContext")
			return
		}
		err := p.HandlePostCreateSmContext(message.ScInstId, postCreatSmMsgReq)
		if err != nil {
			rlogger.Trace(types.ModuleAmfSbi, rlogger.DEBUG, ctx, "failed to HandlePostCreateSmContext, error(%s)", err)
		}
	case sbicmn.PduSessUpdateSMContextReq:
		rlogger.Trace(types.ModuleAmfSbi, rlogger.DEBUG, ctx, "PostModifySmContext msg")
		postModifySmMsgReq, ok := message.MsgData.(*sbicmn.SbiPostModifySmContext)
		if !ok {
			rlogger.Trace(types.ModuleAmfSbi, rlogger.ERROR, ctx, "incorrect message PostModifySmContext")
			return
		}
		err := p.HandlePostModifySmContext(message.ScInstId, postModifySmMsgReq)
		if err != nil {
			rlogger.Trace(types.ModuleAmfSbi, rlogger.DEBUG, ctx, "failed to HandlePostModifySmContext, error(%s)", err)
		}
	case sbicmn.PduSessReleaseSMContextReq:
		rlogger.Trace(types.ModuleAmfSbi, rlogger.DEBUG, ctx, "PostReleaseSmContext msg")
		postReleaseSmMsgReq, ok := message.MsgData.(*sbicmn.SbiPostReleaseSmContext)
		if !ok {
			rlogger.Trace(types.ModuleAmfSbi, rlogger.ERROR, ctx, "incorrect message PostReleaseSmContext")
			return
		}
		err := p.HandlePostReleaseSmContext(message.ScInstId, postReleaseSmMsgReq)
		if err != nil {
			rlogger.Trace(types.ModuleAmfSbi, rlogger.DEBUG, ctx, "failed to HandlePostReleaseSmContext, error(%s)", err)
		}
	default:
	}

	return nil
}

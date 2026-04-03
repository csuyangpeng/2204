package mmsender

import (
	"context"
	"lite5gc/amf/context/gctxt"
	"lite5gc/amf/sc/ngaplayer/ngapsender"
	"lite5gc/cmn/message/n11msg"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
)

func SendN2ResRelReq(ctxt context.Context, msgData n11msg.UpdateSMContextResponseData) error {
	rlogger.FuncEntry(types.ModuleAmfNas, ctxt)

	var msgbytes []byte
	if msgData.IeFlags.Test(n11msg.Ieid_n1SmMsg) {
		msgbytes = []byte(msgData.N1SmMsg)
	}

	// get ue context
	ueCtxt, ok := ctxt.Value(types.UeContextCK).(*gctxt.UeContext)
	if !ok {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, nil,
			"fail to get ue context")
		return types.ErrFailFindUeCtxt
	}

	//get gnb instance id for sending ngap msg
	gnbInstId, err := gctxt.GetGnbInstIdByAmfUeNgapId(ueCtxt.GetAmfUeNgapId())
	if err != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"fail to get GnbInstId by AmfUeNgapId(%d)", ueCtxt.GetAmfUeNgapId())
		return types.ErrFailFindGnbInstId
	}

	//  get the scnglayer
	sender, ok := ctxt.Value(types.NgapSenderCK).(*ngapsender.NgapSender)
	if !ok {
		rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"failed to get ngap layer .")
		return types.ErrFailFindNgapLayer
	}

	// send the message in scng layer
	err = sender.SendN2ResourceReleaseCommand(ueCtxt, gnbInstId, msgbytes)
	if err != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.WARN, ctxt,
			"failed to send PDU session resource release command message,err :", err)
		return types.ErrFailSendNgapMsg
	}
	return nil
}

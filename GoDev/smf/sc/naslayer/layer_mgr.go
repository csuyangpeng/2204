package naslayer

import (
	"bytes"
	"context"
	"fmt"
	"lite5gc/cmn/message/nasmsg"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/smf/sc/sessmgnt/smsender"
)

type NasMgr struct {
	PduSessEstbRequest nasmsg.PduSessionEstbRequestMsg
	PduSessRelRequest  nasmsg.PduSessionReleaseRequestMsg
	PduSessModRequest  nasmsg.PduSessionModifyRequestMsg
	PdqSessModComplete nasmsg.PduSessionModifyCompleteMsg
	PduSessModReject   nasmsg.PduSessionModifyRejectMsg

	PduSessionEstbReqRecvCounter uint32
}

// HandleIncomingNasMsg process nas initial ue message
func (p *NasMgr) HandleIncomingNasMsg(ctx context.Context, msgData []byte) error {
	rlogger.FuncEntry(types.ModuleSmfNas, ctx)

	rlogger.Trace(types.ModuleSmfNas, rlogger.DEBUG, ctx, "smf Receive Nas Msg:%v", msgData)

	nasMsg := bytes.NewReader(msgData)
	// 1. decode extract the EPD
	epd, err := nas.GetEpd(nasMsg)
	if err != nil {
		rlogger.Trace(types.ModuleSmfNas, rlogger.DEBUG, ctx, "failed to decode nas epd")
		return fmt.Errorf("failed to decode nas epd")
	}

	switch epd {
	case nas.Epd5gsSessMgntMsg:
		// decode sm header
		smHeader := &nas.SmNasMessageHeader{}
		err := smHeader.Decode(nasMsg)
		if err != nil {
			rlogger.Trace(types.ModuleSmfNas, rlogger.DEBUG, ctx, "failed to decode nas session management header")
			return fmt.Errorf("decode nas sm header failed")
		}
		// store the smHeader in context
		ctx = context.WithValue(ctx, types.SmHeaderCK, smHeader)
		switch smHeader.MessageType {
		case nas.PduSessEstabishRequest:
			cause := p.HandleSmPduSessEstbRequest(ctx, smHeader, nasMsg)
			smsender.ProcessPduSessEstbReqCause(ctx, cause)
		case nas.PduSessionRelRequest:
			cause := p.HandleSmPduSessRelRequest(ctx, smHeader, nasMsg)
			smsender.ProcessPduSessRelReqCause(ctx, smHeader.PduSessionID, cause)
		case nas.PduSessionRelComplete:
			p.HandleSmPduSessRelComplete(ctx, smHeader)
		case nas.PduSessionModRequest:
			cause := p.HandleSmPduSessModRequest(ctx, smHeader, nasMsg)
			smsender.ProcessPduSessModReqCause(ctx, smHeader.PduSessionID, cause)
		case nas.PduSessionModComplete:
			p.HandleSmPduSessModComplete(ctx, smHeader)
		default:
			rlogger.Trace(types.ModuleSmfNas, rlogger.DEBUG, ctx, "unsupported header (%s)", smHeader)
		}
	default:
		rlogger.Trace(types.ModuleSmfNas, rlogger.DEBUG, ctx, "unsupported epd (%x)", epd)
	}

	return nil
}

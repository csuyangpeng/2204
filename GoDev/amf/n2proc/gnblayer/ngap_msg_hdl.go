package gnblayer

import (
	codec "lite5gc/cmn/message/ngap/ngapcodec"
	"lite5gc/cmn/message/ngap/ngapmsg"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types3gpp"
)

// HandleSctpMsg process all the sctp message from Ran Node side
func (p *GnbLayer) HandleSctpMsg(msgBuf *types.MsgBuf) error {
	rlogger.FuncEntry(types.ModuleAmfN2Proc, nil)

	msgHeader := codec.NewNgapCodec()
	msgHeader.SetEncBuffer(&msgBuf.Buffer[0], msgBuf.MsgLen)
	msgHeader.DecodeHeader(p.ossCtxt.GetOssCtxtPtr_m())

	//dispatch message here
	prcdCode := types3gpp.ProcedureCode(msgHeader.GetProcedureCode())
	msgType := types3gpp.MsgType(msgHeader.GetMsgType())
	rlogger.Trace(types.ModuleAmfN2Proc, rlogger.DEBUG, nil,
		"get the header of incoming ngap message: type(%s), procedure(%s)",
		msgType,
		prcdCode)

	switch prcdCode {
	case types3gpp.NGSetup:
		switch msgType {
		case types3gpp.InitiatingMessage:
			p.handleNgSetupRequestMsg(msgBuf)
		case types3gpp.SuccessfulOutcome:
			p.handleNgSetupResponseMsg(msgBuf)
		case types3gpp.UnsuccessfulOutcome:
			p.handleNgSetupFailureMsg(msgBuf)
		default:
			rlogger.Trace(types.ModuleAmfN2Proc, rlogger.ERROR, nil, "invalid message type(%s)", msgType)
		}
	case types3gpp.NGReset:
		p.handleNgResetMsg(msgBuf)
	default:
		rlogger.Trace(types.ModuleAmfN2Proc, rlogger.DEBUG, nil, "%s, send to sc goroutine", prcdCode)
		err := p.HandleUeMessages(prcdCode, msgType, msgBuf)
		if err != nil {
			rlogger.Trace(types.ModuleAmfN2Proc, rlogger.ERROR, nil, "failed to handle message(%s:%s) on ngap",
				prcdCode, msgType)
		}
	}
	return nil
}

func (p *GnbLayer) HandleUeMessages(prcdCode types3gpp.ProcedureCode,
	msgType types3gpp.MsgType, msgBuf *types.MsgBuf) error {
	rlogger.FuncEntry(types.ModuleAmfN2Proc, nil)

	var err error
	var amfUeNgapId uint64

	switch prcdCode {
	case types3gpp.InitialUEMessage:
		if msgType == types3gpp.InitiatingMessage {
			//msg := ngapmsg.NewInitialUeMessage()
			//msg.SetOssCodecCtxt(p.ossCtxt.GetOssCtxtPtr_m())
			//err := msg.Decode(msgBuf.Buffer)
			//if err != nil {
			//	rlogger.Trace(types.ModuleAmfN2Proc, rlogger.ERROR, nil, "failed to decode %s;%s", prcdCode, msgType)
			//	return err
			//}
			amfUeNgapId = types3gpp.InvalidAmfNgapId
		} else {
			rlogger.Trace(types.ModuleAmfN2Proc, rlogger.ERROR, nil,
				"invalid msg type(%s) for initial ue message.", msgType)
		}
	case types3gpp.UplinkNASTransport:
		if msgType == types3gpp.InitiatingMessage {
			msg := ngapmsg.NewUplinkNasTransportMessage()
			msg.SetOssCodecCtxt(p.ossCtxt.GetOssCtxtPtr_m())
			err := msg.Decode(msgBuf.Buffer)
			if err != nil {
				rlogger.Trace(types.ModuleAmfN2Proc, rlogger.ERROR, nil,
					"failed to decode %s;%s", prcdCode, msgType)
				return err
			}
			amfUeNgapId = msg.AmfUeNgapId
		} else {
			rlogger.Trace(types.ModuleAmfN2Proc, rlogger.ERROR, nil,
				"invalid msg type(%s) for uplinkn nas message.", msgType)
		}
	case types3gpp.PDUSessionResourceSetup:
		if msgType == types3gpp.InitiatingMessage {
			msg := ngapmsg.NewPduSessResSetupReqMsg()
			msg.SetOssCodecCtxt(p.ossCtxt.GetOssCtxtPtr_m())
			err := msg.Decode(msgBuf.Buffer)
			if err != nil {
				rlogger.Trace(types.ModuleAmfN2Proc, rlogger.ERROR, nil,
					"failed to decode %s;%s", prcdCode, msgType)
				return err
			}
			amfUeNgapId = msg.AmfUeNgApId
		} else if msgType == types3gpp.SuccessfulOutcome {
			msg := ngapmsg.NewPduSessResSetupRespMsg()
			msg.SetOssCodecCtxt(p.ossCtxt.GetOssCtxtPtr_m())
			err := msg.Decode(msgBuf.Buffer)
			if err != nil {
				rlogger.Trace(types.ModuleAmfN2Proc, rlogger.ERROR, nil,
					"failed to decode %s;%s", prcdCode, msgType)
				return err
			}
			amfUeNgapId = msg.AmfUeNGAPId
		} else {
			rlogger.Trace(types.ModuleAmfN2Proc, rlogger.ERROR, nil,
				"invalid msg type(%s) for pdu sess resource setup message.", msgType)
		}
	case types3gpp.InitialContextSetup:
		if msgType == types3gpp.InitiatingMessage {
			msg := ngapmsg.NewInitialContextSetupReqMsg()
			msg.SetOssCodecCtxt(p.ossCtxt.GetOssCtxtPtr_m())
			err := msg.Decode(msgBuf.Buffer)
			if err != nil {
				rlogger.Trace(types.ModuleAmfN2Proc, rlogger.ERROR, nil,
					"failed to decode %s;%s", prcdCode, msgType)
				return err
			}
			amfUeNgapId = msg.AmfUeNgapId
		} else if msgType == types3gpp.SuccessfulOutcome {
			msg := ngapmsg.NewInitialContextSetupRespMsg()
			msg.SetOssCodecCtxt(p.ossCtxt.GetOssCtxtPtr_m())
			err := msg.Decode(msgBuf.Buffer[:msgBuf.MsgLen])
			if err != nil {
				rlogger.Trace(types.ModuleAmfN2Proc, rlogger.ERROR, nil,
					"failed to decode %s;%s", prcdCode, msgType)
				return err
			}
			amfUeNgapId = msg.AmfUeNGAPId
		} else {
			rlogger.Trace(types.ModuleAmfN2Proc, rlogger.ERROR, nil,
				"invalid msg type(%s) for pdu sess resource setup message.", msgType)
		}
	case types3gpp.UEContextReleaseRequest:
		if msgType == types3gpp.InitiatingMessage {
			msg := ngapmsg.NewUeContextReleaseReqMsg()
			msg.SetOssCodecCtxt(p.ossCtxt.GetOssCtxtPtr_m())
			err := msg.Decode(msgBuf.Buffer)
			if err != nil {
				rlogger.Trace(types.ModuleAmfN2Proc, rlogger.ERROR, nil,
					"failed to decode %s;%s", prcdCode, msgType)
				return err
			}
			amfUeNgapId = msg.AmfUeNgapId
		} else {
			rlogger.Trace(types.ModuleAmfN2Proc, rlogger.ERROR, nil,
				"invalid msg type(%s) for initial ue message.", msgType)
		}
	case types3gpp.UEContextRelease:
		if msgType == types3gpp.InitiatingMessage {
			msg := ngapmsg.NewUeContextReleaseCmpMsg()
			msg.SetOssCodecCtxt(p.ossCtxt.GetOssCtxtPtr_m())
			err := msg.Decode(msgBuf.Buffer)
			if err != nil {
				rlogger.Trace(types.ModuleAmfN2Proc, rlogger.ERROR, nil,
					"failed to decode %s;%s", prcdCode, msgType)
				return err
			}
			amfUeNgapId = msg.AmfUeNgapId
		} else if msgType == types3gpp.SuccessfulOutcome {
			msg := ngapmsg.NewUeContextReleaseCmpMsg()
			msg.SetOssCodecCtxt(p.ossCtxt.GetOssCtxtPtr_m())
			err := msg.Decode(msgBuf.Buffer)
			if err != nil {
				rlogger.Trace(types.ModuleAmfN2Proc, rlogger.ERROR, nil,
					"failed to decode %s;%s", prcdCode, msgType)
				return err
			}
			amfUeNgapId = msg.AmfUeNgapId
		} else {
			rlogger.Trace(types.ModuleAmfN2Proc, rlogger.ERROR, nil,
				"invalid msg type(%s) for pdu sess resource setup message.", msgType)
		}
	case types3gpp.UERadioCapabilityInfoIndication:
		if msgType == types3gpp.InitiatingMessage {
			msg := ngapmsg.NewUeRadioCapInfoIndMsg()
			msg.SetOssCodecCtxt(p.ossCtxt.GetOssCtxtPtr_m())
			err := msg.Decode(msgBuf.Buffer)
			if err != nil {
				rlogger.Trace(types.ModuleAmfN2Proc, rlogger.ERROR, nil, "failed to decode %s;%s", prcdCode, msgType)
				return err
			}
			amfUeNgapId = msg.AmfUeNgapId
		} else {
			rlogger.Trace(types.ModuleAmfN2Proc, rlogger.ERROR, nil,
				"invalid msg type(%s) for initial ue message.", msgType)
		}
	case types3gpp.PDUSessionResourceRelease:
		if msgType == types3gpp.SuccessfulOutcome {
			msg := ngapmsg.NewPduSessResRelRespMsg()
			msg.SetOssCodecCtxt(p.ossCtxt.GetOssCtxtPtr_m())
			err := msg.Decode(msgBuf.Buffer)
			if err != nil {
				rlogger.Trace(types.ModuleAmfN2Proc, rlogger.ERROR, nil, "failed to decode %s;%s", prcdCode, msgType)
				return err
			}
			amfUeNgapId = msg.AmfUeNGAPId
		} else {
			rlogger.Trace(types.ModuleAmfN2Proc, rlogger.ERROR, nil,
				"invalid msg type(%s) for initial ue message.", msgType)
		}
	default:
		rlogger.Trace(types.ModuleAmfN2Proc, rlogger.ERROR, nil,
			"unsupported message %s", prcdCode)
	}
	//construct Ngap 2 sc message
	msg := &types3gpp.Gnb2AmfScMsg{}
	msg.MsgType = msgType
	msg.PrcdCode = prcdCode
	msg.GnbInfo = p.gnbInfo
	msg.NgapMsg = msgBuf.Buffer[:msgBuf.MsgLen] //original ngap message

	var scInstId uint32
	if amfUeNgapId == types3gpp.InvalidAmfNgapId {
		scInstId = types3gpp.InvalidInstId
	} else {
		//bit 40-32 is the sc instance id
		scInstId = uint32(amfUeNgapId>>32) & 0x000000FF
	}

	p.SendMsg2AmfSC(scInstId, msg)

	return err
}

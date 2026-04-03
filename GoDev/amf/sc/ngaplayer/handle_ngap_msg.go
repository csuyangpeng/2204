package ngaplayer

import (
	"context"
	"fmt"
	"lite5gc/amf/sc/statistics"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types3gpp"
	"lite5gc/oam/pm"
)

func (p *LayerMgr) HandlerIncomingMsg(ctx context.Context, msg *types3gpp.Gnb2AmfScMsg) (err error) {
	rlogger.FuncEntry(types.ModuleAmfNgap, nil)

	//dispatch message here
	switch msg.PrcdCode {
	case types3gpp.InitialUEMessage:
		//msg counter
		//pm.PegCounter(statistics.InitialUEMessageCounter)
		err = p.handleInitialUeMsg(ctx, msg.GnbInfo, msg.NgapMsg)
		if err != nil {
			rlogger.Trace(types.ModuleAmfNgap, rlogger.ERROR, nil,
				"failed to handleInitialUeMsg, err(%s) ", err)
			return fmt.Errorf("failed to handle initial ue msg, err(%s) ", err)
		}
	case types3gpp.UplinkNASTransport:
		//msg counter
		pm.PegCounter(statistics.UpLinkNASTransportCounter)
		err = p.handleUpLinkNasTransportMsg(ctx, msg.NgapMsg)
		if err != nil {
			rlogger.Trace(types.ModuleAmfNgap, rlogger.ERROR, nil,
				"failed to handle upLink nas transport msg, err(%s) ", err)
			return fmt.Errorf("failed to handleUpLinkNasTransportMsg, err(%s) ", err)
		}
	case types3gpp.UEContextReleaseRequest:
		//msg counter
		pm.PegCounter(statistics.UEContextReleaseRequestCounter)
		err = p.handleUeCtxtReleaseReqMsg(ctx, msg.NgapMsg)
		if err != nil {
			rlogger.Trace(types.ModuleAmfNgap, rlogger.ERROR, nil,
				"failed to handleUeCtxtReleaseReqMsg, err(%s) ", err)
			return fmt.Errorf("failed to handle ue ctxt release request msg, err(%s) ", err)
		}
	case types3gpp.UEContextRelease:
		//msg counter
		pm.PegCounter(statistics.UEContextReleaseCompleteCounter)
		err = p.handleUeCtxtReleaseCmpMsg(ctx, msg.NgapMsg)
		if err != nil {
			rlogger.Trace(types.ModuleAmfNgap, rlogger.ERROR, nil,
				"failed to handleUeCtxtReleaseCmpMsg, err(%s) ", err)
			return fmt.Errorf("failed to handle ue ctxt release complete msg, err(%s) ", err)
		}
	case types3gpp.InitialContextSetup:
		//msg counter
		pm.PegCounter(statistics.InitialContextSetupResponseCounter)
		switch msg.MsgType {
		case types3gpp.InitiatingMessage:
			rtStr := fmt.Sprintf("invalid message in amf (%s,%s)", msg.PrcdCode, msg.MsgType)
			rlogger.Trace(types.ModuleAmfNgap, rlogger.DEBUG, nil, "%s", rtStr)
			return fmt.Errorf("%s", rtStr)
		case types3gpp.SuccessfulOutcome:
			err = p.handleInitialContextSetupRespMsg(ctx, msg.NgapMsg)
			if err != nil {
				rlogger.Trace(types.ModuleAmfNgap, rlogger.ERROR, nil,
					"failed to handleInitialContextSetupRespMsg, err(%s) ", err)
				return fmt.Errorf("failed to handle initial context setup response msg, err(%s) ", err)
			}
		default:
			rlogger.Trace(types.ModuleAmfNgap, rlogger.ERROR, nil, "invalid message type")
		}
	//case ngapmsg.InitialContextSetupFailure:
	case types3gpp.PDUSessionResourceSetup:
		//msg counter
		pm.PegCounter(statistics.PDUSessionResourceSetupRequestCounter)
		switch msg.MsgType {
		case types3gpp.InitiatingMessage:
			rtStr := fmt.Sprintf("invalid message in amf (%s,%s)", msg.PrcdCode, msg.MsgType)
			rlogger.Trace(types.ModuleAmfNgap, rlogger.DEBUG, nil, "%s", rtStr)
			return fmt.Errorf("%s", rtStr)
		case types3gpp.SuccessfulOutcome:
			err = p.handlePduSessResSetupResponseMsg(ctx, msg.NgapMsg)
			if err != nil {
				rlogger.Trace(types.ModuleAmfNgap, rlogger.ERROR, nil,
					"failed to handle pdu session resource setup response msg, err(%s) ", err)
				return fmt.Errorf("failed to handlePduSessResSetupResponseMsg, err(%s) ", err)
			}
		default:
			rlogger.Trace(types.ModuleAmfNgap, rlogger.ERROR, nil, "invalid message type")
		}
	case types3gpp.PDUSessionResourceRelease:
		err = p.handleN2ResRelRespMsg()
		if err != nil {
			rlogger.Trace(types.ModuleAmfNgap, rlogger.ERROR, nil,
				"failed to N2ResRelRespMsg, err(%s) ", err)
			return fmt.Errorf("failed to N2ResRelRespMsg, err(%s) ", err)
		}
	case types3gpp.UERadioCapabilityInfoIndication:
		//msg counter
		pm.PegCounter(statistics.UERadioCapabilityInfoIndicationCounter)
		rlogger.Trace(types.ModuleAmfNgap, rlogger.DEBUG, nil,
			"received ue radio capability info indication msg, do nothing")
		err = p.handleUeRadioCapInfoIndMsg(ctx, msg.NgapMsg)
		if err != nil {
			rlogger.Trace(types.ModuleAmfNgap, rlogger.ERROR, nil,
				"failed to handle ue radio capability info indication msg, err(%s) ", err)
			return fmt.Errorf("failed to handle ue radio capability info indication msg, err(%s) ", err)
		}
	case types3gpp.MaxProcedureCode:
		// reset message or ta table update message
		switch msg.MsgType {
		case types3gpp.SctpShutdown:
			p.handleGnbShutdownEvent(ctx)
		case types3gpp.TaTblUpdate:
			p.handleTatblUpdateEvent(ctx)
		}
	default:
		rlogger.Trace(types.ModuleAmfNgap, rlogger.ERROR, nil, "invalid procedure code (%d)", msg.PrcdCode)
	}
	return
}

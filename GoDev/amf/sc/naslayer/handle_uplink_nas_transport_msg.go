package naslayer

import (
	"bytes"
	"context"
	"fmt"
	"lite5gc/amf/context/gctxt"
	"lite5gc/amf/sc/statistics"
	"lite5gc/cmn/message/nasmsg"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/nas/nasie"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
)

func (p *NasMgr) HandleULNasTransportMsg(ctx context.Context,
	n2connData *gctxt.N2ConnCtxt, plainNasMsg *bytes.Reader) error {
	td := []interface{}{ctx, n2connData}
	rlogger.FuncEntry(types.ModuleAmfNas, td)

	p.upLinkNasTransport.Reset()

	//decode the message
	err := p.upLinkNasTransport.Decode(plainNasMsg)
	if err != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, td, "failed to decode "+
			"UL Nas Transport, error(%s)", err)
		return fmt.Errorf("failed to decode nas message, error(%s)", err)
	}

	//get ue context with amf ngap id
	ueCtxt, ok := ctx.Value(types.UeContextCK).(*gctxt.UeContext)
	if !ok {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, td, "failed to get ue context in UL Nas Transport Msg")
		return types.ErrFailFindUeCtxt
	}

	if ueCtxt.GetRmState() != types.StateRmRegistered {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, td, "ue not registered yet, please send register request first")
		return fmt.Errorf("ue not registered yet, please send register request first")
	}

	switch p.upLinkNasTransport.PayloadType {
	case nasie.N1SmInformation:
		n1SmMsgCont := p.upLinkNasTransport.PayloadContainer.PayloadContainerEntry[0].ContainerContents
		//get the message type
		msgType := nas.SmMsgType(n1SmMsgCont[3])

		switch msgType {
		case nas.PduSessEstabishRequest:
			//msg counter
			statistics.PduSessEstablishRequestCounter.Inc(1)
			if p.upLinkNasTransport.OptIeBitSet.Test(nasmsg.IeidUplinknastransRequesttype) == true {
				switch p.upLinkNasTransport.RequestType {
				case nasie.InitialRequest:
					err := HandlePduSessEstabRequestMsg(ctx, ueCtxt, &(p.upLinkNasTransport))
					if err != nil {
						rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, td, "send create sm ctxt req message failed.error(%s)", err)
						return fmt.Errorf("send create sm ctxt req msg failed. error(%s)", err)
					}
				case nasie.ExistingPduSess:
				case nasie.InitialEmergency:
				case nasie.ExistingEmergency:
				case nasie.ModifyRequest:
				case nasie.ReservedReqType:
				default:
					rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, td, "unsupported request type(%s)", p.upLinkNasTransport.RequestType)
					return fmt.Errorf("unsupported request type(%s)", p.upLinkNasTransport.RequestType)
				}
			}
		case nas.PduSessionRelRequest:
			err := HandlePduSessRelRequestMsg(ctx, ueCtxt, &(p.upLinkNasTransport))
			if err != nil {
				rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, td, "send update sm ctxt req message failed.error(%s)", err)
				return fmt.Errorf("send update sm ctxt req msg failed. error(%s)", err)
			}
		case nas.PduSessionRelComplete:
			err := HandleSessionRelCompleteMsg(ctx, ueCtxt, &(p.upLinkNasTransport))
			if err != nil {
				rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, td, "send update sm ctxt req message failed.error(%s)", err)
				return fmt.Errorf("send update sm ctxt req msg failed. error(%s)", err)
			}
		case nas.PduSessionModRequest:
			err := HandlePduSessModRequestMsg(ctx, ueCtxt, &(p.upLinkNasTransport))
			if err != nil {
				rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, td, "send update sm ctxt req message failed.error(%s)", err)
				return fmt.Errorf("send update sm ctxt req msg failed. error(%s)", err)
			}
		case nas.PduSessionModComplete:
			err := HandleSessionModCompleteMsg(ctx, ueCtxt, &(p.upLinkNasTransport))
			if err != nil {
				rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, td, "send update sm ctxt req message failed.error(%s)", err)
				return fmt.Errorf("send update sm ctxt req msg failed. error(%s)", err)
			}
		default:
			rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, td, "unsupported sm msgtype(%s) currently.", msgType)
			return fmt.Errorf("unsupported sm msg type(%s)", msgType)
		}
	default:
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, td, "unsuppported playload type(%s) currently", p.upLinkNasTransport.PayloadType)
		return fmt.Errorf("unspported payload type(%s)", p.upLinkNasTransport.PayloadType)
	}

	return nil
}

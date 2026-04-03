package ngaplayer

import (
	"context"
	"fmt"
	"lite5gc/amf/context/gctxt"
	"lite5gc/amf/sc/naslayer"
	"lite5gc/cmn/message/ngap/ngapmsg"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
)

func (p *LayerMgr) handleUpLinkNasTransportMsg(ctx context.Context, msgBuf []byte) error {
	rlogger.FuncEntry(types.ModuleAmfNgap, nil)

	// decode the ngap message
	message := ngapmsg.NewUplinkNasTransportMessage()
	message.SetOssCodecCtxt(p.ossCtxt.GetOssCtxtPtr_m())
	err := message.Decode(msgBuf)
	if err != nil {
		rlogger.Trace(types.ModuleAmfNgap, rlogger.ERROR, nil,
			"failed to decode uplink nas transport message")
		return err
	}

	nasLayer, ok := ctx.Value(types.NasLayerCK).(*naslayer.NasMgr)
	if !ok {
		rlogger.Trace(types.ModuleAmfNgap, rlogger.WARN, nil, "fail to find nas layer")
		return types.ErrFailFindNasLayer
	}

	amfUeNgApId := gctxt.AmfUeNgApId(message.AmfUeNgapId)

	//get ue ctxt
	ueCtxt, err := gctxt.GetUeContext(gctxt.AmfUeNgApId(amfUeNgApId))
	if err != nil {
		rlogger.Trace(types.ModuleAmfNgap, rlogger.ERROR, nil,
			"failed tngap_msg_hdl.goo find the ue context with AmfUeNgApId (%x),  AmfUeIdUeCtxtTable(%s)",
			amfUeNgApId,
			gctxt.DumpAmfUeIdUeCtxtTable())
		return err
	}
	ctx = context.WithValue(ctx, types.UeContextCK, ueCtxt)

	// get n2 connection data
	n2ConnCtxt, err := gctxt.GetN2ConnContext(amfUeNgApId)
	if err != nil {
		rlogger.Trace(types.ModuleAmfNgap, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"failed to get ne conn context, err(%s) ", err)
		return fmt.Errorf("failed to get n2 connection context, err(%s) ", err)
	}

	err = nasLayer.HandleIncomingNasMsg(ctx, n2ConnCtxt, message.NasPdu)
	if err != nil {
		return fmt.Errorf("nas layer failed to handle uplink nas transport msg, err(%s) ", err)
	}

	return nil
}

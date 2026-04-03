package ngaplayer

import (
	"context"
	"fmt"
	"lite5gc/amf/context/gctxt"
	"lite5gc/amf/sc/naslayer"
	"lite5gc/cmn/idmgr"
	"lite5gc/cmn/message/ngap/ngapmsg"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types3gpp"
)

func (p *LayerMgr) handleInitialUeMsg(ctx context.Context, gnbInfo types3gpp.GnbInfo, msgBuf []byte) error {
	rlogger.FuncEntry(types.ModuleAmfNgap, nil)

	// decode the ngap message
	message := ngapmsg.NewInitialUeMessage()
	message.SetOssCodecCtxt(p.ossCtxt.GetOssCtxtPtr_m())
	err := message.Decode(msgBuf)
	if err != nil {
		rlogger.Trace(types.ModuleAmfNgap, rlogger.ERROR, nil,
			"failed to decode initial ue message")
		return err
	}

	nasLayer, ok := ctx.Value(types.NasLayerCK).(*naslayer.NasMgr)
	if !ok {
		rlogger.Trace(types.ModuleAmfNgap, rlogger.WARN, nil, "fail to find nas layer")
		return types.ErrFailFindNasLayer
	}

	// allocated and save amfUeNgapId
	ngapId, err := idmgr.GetInst().BorrowID(string(types.AMFUeNgapId))
	if err != nil {
		rlogger.Trace(types.ModuleAmfNgap, rlogger.ERROR, nil, "failed to borrow amf ue ngap id")
		return fmt.Errorf("failed to borrow amf ue ngap id")
	}
	amfUeNgapId := uint64(nasLayer.ScInstId) << 32
	amfUeNgapId = amfUeNgapId | uint64(ngapId)

	n2ConnCtxt, err := gctxt.CreateN2ConnCtxt(gctxt.AmfUeNgApId(amfUeNgapId))
	if err != nil {
		rlogger.Trace(types.ModuleAmfNgap, rlogger.ERROR, nil, "failed to create n2 connection context")
		return fmt.Errorf("failed to create n2 connection data, err(%s)", err)
	}

	n2ConnCtxt.GnbInfo = gnbInfo
	n2ConnCtxt.GnbConnID = message.RanUeNgapId
	n2ConnCtxt.AmfUeNgapID = gctxt.AmfUeNgApId(amfUeNgapId)

	rlogger.Trace(types.ModuleAmfNgap, rlogger.INFO, nil,
		"get from message ran_ue_ngap_id(%x)", n2ConnCtxt.GnbConnID)

	// if ue context request ie is included, initial context setup procedure should be triggered
	if message.IsUeContextRequestPrst {
		n2ConnCtxt.NeedInitCtxtSetupPrcd = true
	}

	err = nasLayer.HandleIncomingNasMsg(ctx, n2ConnCtxt, message.NasPdu)
	if err != nil {
		return fmt.Errorf("nas layer failed to handle intial ue message")
	}
	return nil
}

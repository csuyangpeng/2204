package ngapsender

import (
	"fmt"
	"lite5gc/amf/context/gctxt"
	"lite5gc/amf/sc/statistics"
	"lite5gc/cmn/message/ngap/ngapmsg"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/oam/pm"
)

func (p *NgapSender) SendDownlinkNasTransport(ueCtxt *gctxt.UeContext, gnbInstId uint32, nasmsg []byte) (error, []byte) {
	rlogger.FuncEntry(types.ModuleAmfNgap, nil)

	dlNasTransMsg := ngapmsg.NewDownlinkNasTransportMessage()
	dlNasTransMsg.SetOssCodecCtxt(p.ossCtxt.GetOssCtxtPtr_m())

	dlNasTransMsg.AmfUeNgapId = ueCtxt.GetAmfUeNgapId()

	//for debug
	dlNasTransMsg.RanUeNgapId = ueCtxt.GetRanUeNgapId()
	rlogger.Trace(types.ModuleAmfNgap, rlogger.INFO, nil,
		"send downlink nas transport with amf_ue_ngap_id(%x),ran_ue_ngap_id(%x)",
		dlNasTransMsg.AmfUeNgapId,dlNasTransMsg.RanUeNgapId)

	var encodeNasMsg []byte
	if nasmsg != nil {
		encodeNasMsg = append(encodeNasMsg, nasmsg[:]...)
		dlNasTransMsg.NasPdu = encodeNasMsg
	} else {
		rlogger.Trace(types.ModuleAmfNgap, rlogger.ERROR, nil, "internal error, invalid nasmsg")
		return fmt.Errorf("internal error: invalid nas message"), nil
	}

	msgBuf := types.MsgBuf{}
	msgBuf.Buffer = dlNasTransMsg.Encode()
	msgBuf.MsgLen = len(msgBuf.Buffer)
	ngapInstId := gnbInstId
	err := p.SendNgapMsg(ngapInstId, msgBuf)
	if err != nil {
		rlogger.Trace(types.ModuleAmfNgap, rlogger.ERROR, ueCtxt.GetImsiPtr(), types.ErrFailSendNgapMsg)
		return types.ErrFailSendNgapMsg, nil
	}

	//msg counter
	pm.PegCounter(statistics.DownLinkNASTransportCounter)

	return nil, msgBuf.Buffer
}

func (p *NgapSender) ReSendDownlinkNasTransport(gnbInstId uint32, ngapmsg []byte) error {
	rlogger.FuncEntry(types.ModuleAmfNgap, nil)

	msgBuf := types.MsgBuf{}
	msgBuf.Buffer = ngapmsg
	msgBuf.MsgLen = len(msgBuf.Buffer)
	ngapInstId := gnbInstId
	err := p.SendNgapMsg(ngapInstId, msgBuf)
	if err != nil {
		rlogger.Trace(types.ModuleAmfNgap, rlogger.ERROR, nil, types.ErrFailSendNgapMsg)
		return types.ErrFailSendNgapMsg
	}

	//msg counter
	pm.PegCounter(statistics.DownLinkNASTransportCounter)

	return nil
}

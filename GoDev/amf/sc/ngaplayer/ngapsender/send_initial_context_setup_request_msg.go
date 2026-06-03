package ngapsender

import (
	"lite5gc/amf/context/gctxt"
	"lite5gc/amf/sc/statistics"
	"lite5gc/amf/security/seaf"
	"lite5gc/cmn/message/ngap/ngapmsg"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types/configure"
	"lite5gc/cmn/types3gpp"
	"lite5gc/oam/pm"
)

func (p *NgapSender) SendInitialContextSetupRequest(ueCtxt *gctxt.UeContext,
	gnbInstId uint32,
	nasMsg []byte,
	setupReqList []*types3gpp.PduSessResSetupReqItem) error {

	rlogger.FuncEntry(types.ModuleAmfNgap, nil)

	msg := ngapmsg.NewInitialContextSetupReqMsg()
	msg.SetOssCodecCtxt(p.ossCtxt.GetOssCtxtPtr_m())
	//mandatory IE
	msg.AmfUeNgapId = ueCtxt.GetAmfUeNgapId()
	msg.RanUeNgapId = ueCtxt.GetRanUeNgapId()
	msg.Guami = configure.GetTypesGuami()
	for _, v := range ueCtxt.AllowedNssai {
		msg.AddAllowedNssai(&v)
	}
	msg.UeSecCap = ueCtxt.SecurityCtxt.UeSecCapablity
	//Security key
	kgnb, err := seaf.GetKgnb(ueCtxt)
	if err != nil {
		rlogger.Trace(types.ModuleAmfNgap, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"failed to get kgnb, error(%s), set to default", err)
		kgnb = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	}
	for i, v := range kgnb {
		msg.SecKey[i] = v
	}
	if len(setupReqList) > 0 {
		for _, v := range setupReqList {
			msg.AddPduSessResSetupReqItem(v)
		}
	}
	// Optional IE
	msg.UeAmbr.Downlink = ueCtxt.AccMobSubsData.SubsUeAmbr.Downlink
	msg.UeAmbr.Uplink = ueCtxt.AccMobSubsData.SubsUeAmbr.Uplink
	msg.OptFlags.Set(ngapmsg.ICSR_Ueambr)

	if nasMsg != nil {
		msg.NasPdu = make([]byte, len(nasMsg))
		copy(msg.NasPdu, nasMsg)
		msg.OptFlags.Set(ngapmsg.ICSR_Nas)
	}

	rlogger.Trace(types.ModuleAmfNgap, rlogger.DEBUG, ueCtxt.GetImsiPtr(), "InitialContextSetupReqMsg: ", msg)

	// send ngap msg
	msgBuf := types.MsgBuf{}
	msgBuf.Buffer = msg.Encode()
	msgBuf.MsgLen = len(msgBuf.Buffer)
	ngapInstId := gnbInstId
	err = p.SendNgapMsg(ngapInstId, msgBuf)
	if err != nil {
		rlogger.Trace(types.ModuleAmfNgap,rlogger.ERROR,ueCtxt.GetImsiPtr(),types.ErrFailSendNgapMsg)
		return types.ErrFailSendNgapMsg
	}

	//msg counter
	pm.PegCounter(statistics.InitialContextSetupRequestCounter)

	return nil
}

package gctxt

import (
	"lite5gc/cmn/message/n11msg"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/nas/nasie"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types3gpp"
	"lite5gc/smf/smfcontext/prcdctxt"
	"net"
)

type ScN4MsgData struct {
	SEID           uint64
	DNN            types3gpp.Apn
	SNSSAI         nasie.SNssai
	UEIP           net.IP
	IMSI           types3gpp.Imsi
	PFCPParameters SmfToPFCPParameters

	// pfcp请求返回值
	Cause N4Cause
}

type PduSessContext struct {
	// N4 Session ID
	SEID uint64
	IMSI types3gpp.Imsi

	//SMF要保存服务这个UE的AMF ID, 对应CreateSMContextRequestMsg 中的 ServingNfId
	AmfID string

	PduSessionId nas.PduSessID
	DNN          types3gpp.Apn
	SNSSAI       nasie.SNssai
	UEIP         net.IP
	DnnConfig    *nasie.DNNConfiguration
	SessionAmbr  nasie.SessionAmbr

	IsAuthed bool //true - authenticated, false - unauthenticated

	//procedure ctxt interface type
	prcCtxt         prcdctxt.Base
	UPInactiveTimer uint32
	UpCnxState      n11msg.UpCnxState
	DirectRespFlag  bool
}

func NewPduSessContext(psi nas.PduSessID) *PduSessContext {
	rlogger.FuncEntry(types.ModuleSmfCtxt, nil)
	return &PduSessContext{PduSessionId: psi}
}

func (p *PduSessContext) SetPrcdCtxt(ctxt prcdctxt.Base) {
	rlogger.Trace(types.ModuleSmfCtxt, rlogger.DEBUG, nil,
		"current prcd(%s), set to (%s)", PrcdCtxt2String(p.prcCtxt), PrcdCtxt2String(ctxt))
	p.prcCtxt = ctxt
}

func PrcdCtxt2String(ctxt prcdctxt.Base) string {
	var rt string
	switch ctxt.(type) {
	case nil:
		rt = "nil"
	case *prcdctxt.PduSessionEstbPrcdCtxt:
		rt = "Pdu Sess Estb Prcd Ctxt"
	case *prcdctxt.PduSessionModPrcdCtxt:
		rt = "Pdu Sess Mod Prcd Ctxt"
	case *prcdctxt.PduSessionRelReqPrcdCtxt:
		rt = "Pdu Sess Rel Prcd Ctxt"
	case *prcdctxt.AnRelSerReqPrcdCtxt:
		rt = "AnRel SerReq Prcedure ctxt"
	default:
		rt = "Unknown Prcd Ctxt"
	}
	return rt
}

func (p *PduSessContext) GetPrcdCtxt() prcdctxt.Base {
	return p.prcCtxt
}
